// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"
)

// A Converter converts HTTP requests and responses to the Request and Response types
// of this package, while removing or redacting information.
type Converter struct {
	// These all apply to both headers and trailers.
	RedactHeaders         []tRegexp // replace matching headers with "REDACTED"
	RemoveRequestHeaders  []tRegexp // remove matching headers in requests
	RemoveResponseHeaders []tRegexp // remove matching headers in responses
}

// A regexp that can be marshaled to and from text.
type tRegexp struct {
	*regexp.Regexp
}

func (r tRegexp) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *tRegexp) UnmarshalText(b []byte) error {
	var err error
	r.Regexp, err = regexp.Compile(string(b))
	return err
}

func (c *Converter) registerRemoveRequestHeaders(pat string) {
	c.RemoveRequestHeaders = append(c.RemoveRequestHeaders, pattern(pat))
}

func (c *Converter) registerRedactHeaders(pat string) {
	c.RedactHeaders = append(c.RedactHeaders, pattern(pat))
}

var (
	defaultRemoveRequestHeaders = []string{
		"Authorization", // not only is it secret, but it is probably missing on replay
		"Proxy-Authorization",
		"Connection",
		"Content-Type", // because it may contain a random multipart boundary
		"Date",
		"Host",
		"Transfer-Encoding",
		"Via",
		"X-Forwarded-*",
		// Google-specific
		"X-Cloud-Trace-Context", // OpenCensus traces have a random ID
		"X-Goog-Api-Client",     // can differ for, e.g., different Go versions
	}

	defaultRemoveBothHeaders = []string{
		// Google-specific
		// GFEs scrub X-Google- and X-GFE- headers from requests and responses.
		// Drop them from recordings made by users inside Google.
		// http://g3doc/gfe/g3doc/gfe3/design/http_filters/google_header_filter
		// (internal Google documentation).
		"X-Google-*",
		"X-Gfe-*",
	}

	defaultRedactHeaders = []string{
		// Google-specific
		// Used by Cloud Storage for customer-supplied encryption.
		"X-Goog-*Encryption-Key",
	}
)

func defaultConverter() *Converter {
	c := &Converter{}
	for _, h := range defaultRedactHeaders {
		c.registerRedactHeaders(h)
	}
	for _, h := range defaultRemoveRequestHeaders {
		c.registerRemoveRequestHeaders(h)
	}
	for _, h := range defaultRemoveBothHeaders {
		c.registerRemoveRequestHeaders(h)
		c.RemoveResponseHeaders = append(c.RemoveResponseHeaders, pattern(h))
	}
	return c
}

// Convert a pattern into a regexp.
// A pattern is like a literal regexp anchored on both ends, with only one
// non-literal character: "*", which matches zero or more characters.
func pattern(p string) tRegexp {
	q := regexp.QuoteMeta(p)
	q = "^" + strings.Replace(q, `\*`, `.*`, -1) + "$"
	// q must be a legal regexp.
	return tRegexp{regexp.MustCompile(q)}
}

func (c *Converter) convertRequest(req *http.Request) (*Request, error) {
	body, err := snapshotBody(&req.Body)
	if err != nil {
		return nil, err
	}
	mediaType, parts, err := parseRequestBody(req.Header.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}
	return &Request{
		Method:    req.Method,
		URL:       req.URL.String(),
		Header:    scrubHeaders(req.Header, c.RedactHeaders, c.RemoveRequestHeaders),
		MediaType: mediaType,
		BodyParts: parts,
		Trailer:   scrubHeaders(req.Trailer, c.RedactHeaders, c.RemoveRequestHeaders),
	}, nil
}

// parseRequestBody parses the Content-Type header, reads the body, and splits it into
// parts if necessary. It returns the media type and the body parts.
func parseRequestBody(contentType string, body []byte) (string, [][]byte, error) {
	if contentType == "" {
		// No content-type header. There should not be a body.
		if len(body) != 0 {
			return "", nil, errors.New("no Content-Type, but body")
		}
		return "", nil, nil
	}
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return "", nil, err
	}
	var parts [][]byte
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				return "", nil, err
			}
			part, err := ioutil.ReadAll(p)
			if err != nil {
				return "", nil, err
			}
			// TODO(jba): care about part headers?
			parts = append(parts, part)
		}
	} else {
		parts = [][]byte{body}
	}
	return mediaType, parts, nil
}

func (c *Converter) convertResponse(res *http.Response) (*Response, error) {
	data, err := snapshotBody(&res.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		ProtoMajor: res.ProtoMajor,
		ProtoMinor: res.ProtoMinor,
		Header:     scrubHeaders(res.Header, c.RedactHeaders, c.RemoveResponseHeaders),
		Body:       data,
		Trailer:    scrubHeaders(res.Trailer, c.RedactHeaders, c.RemoveResponseHeaders),
	}, nil
}

func snapshotBody(body *io.ReadCloser) ([]byte, error) {
	data, err := ioutil.ReadAll(*body)
	if err != nil {
		return nil, err
	}
	(*body).Close()
	*body = ioutil.NopCloser(bytes.NewReader(data))
	return data, nil
}

// Copy headers, redacting some and removing others.
func scrubHeaders(hs http.Header, redact, remove []tRegexp) http.Header {
	rh := http.Header{}
	for k, v := range hs {
		switch {
		case match(k, redact):
			rh.Set(k, "REDACTED")
		case match(k, remove):
			// skip
		default:
			rh[k] = v
		}
	}
	return rh
}

func match(s string, res []tRegexp) bool {
	for _, re := range res {
		if re.MatchString(s) {
			return true
		}
	}
	return false
}
