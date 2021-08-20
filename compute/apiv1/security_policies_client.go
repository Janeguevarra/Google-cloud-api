// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newSecurityPoliciesClientHook clientHook

// SecurityPoliciesCallOptions contains the retry settings for each method of SecurityPoliciesClient.
type SecurityPoliciesCallOptions struct {
	AddRule                         []gax.CallOption
	Delete                          []gax.CallOption
	Get                             []gax.CallOption
	GetRule                         []gax.CallOption
	Insert                          []gax.CallOption
	List                            []gax.CallOption
	ListPreconfiguredExpressionSets []gax.CallOption
	Patch                           []gax.CallOption
	PatchRule                       []gax.CallOption
	RemoveRule                      []gax.CallOption
}

// internalSecurityPoliciesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalSecurityPoliciesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddRule(context.Context, *computepb.AddRuleSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeleteSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetSecurityPolicyRequest, ...gax.CallOption) (*computepb.SecurityPolicy, error)
	GetRule(context.Context, *computepb.GetRuleSecurityPolicyRequest, ...gax.CallOption) (*computepb.SecurityPolicyRule, error)
	Insert(context.Context, *computepb.InsertSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListSecurityPoliciesRequest, ...gax.CallOption) *SecurityPolicyIterator
	ListPreconfiguredExpressionSets(context.Context, *computepb.ListPreconfiguredExpressionSetsSecurityPoliciesRequest, ...gax.CallOption) (*computepb.SecurityPoliciesListPreconfiguredExpressionSetsResponse, error)
	Patch(context.Context, *computepb.PatchSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	PatchRule(context.Context, *computepb.PatchRuleSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
	RemoveRule(context.Context, *computepb.RemoveRuleSecurityPolicyRequest, ...gax.CallOption) (*Operation, error)
}

// SecurityPoliciesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The SecurityPolicies API.
type SecurityPoliciesClient struct {
	// The internal transport-dependent client.
	internalClient internalSecurityPoliciesClient

	// The call options for this service.
	CallOptions *SecurityPoliciesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SecurityPoliciesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SecurityPoliciesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SecurityPoliciesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddRule inserts a rule into a security policy.
func (c *SecurityPoliciesClient) AddRule(ctx context.Context, req *computepb.AddRuleSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AddRule(ctx, req, opts...)
}

// Delete deletes the specified policy.
func (c *SecurityPoliciesClient) Delete(ctx context.Context, req *computepb.DeleteSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get list all of the ordered rules present in a single specified policy.
func (c *SecurityPoliciesClient) Get(ctx context.Context, req *computepb.GetSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetRule gets a rule at the specified priority.
func (c *SecurityPoliciesClient) GetRule(ctx context.Context, req *computepb.GetRuleSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicyRule, error) {
	return c.internalClient.GetRule(ctx, req, opts...)
}

// Insert creates a new policy in the specified project using the data included in the request.
func (c *SecurityPoliciesClient) Insert(ctx context.Context, req *computepb.InsertSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List list all the policies that have been configured for the specified project.
func (c *SecurityPoliciesClient) List(ctx context.Context, req *computepb.ListSecurityPoliciesRequest, opts ...gax.CallOption) *SecurityPolicyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// ListPreconfiguredExpressionSets gets the current list of preconfigured Web Application Firewall (WAF) expressions.
func (c *SecurityPoliciesClient) ListPreconfiguredExpressionSets(ctx context.Context, req *computepb.ListPreconfiguredExpressionSetsSecurityPoliciesRequest, opts ...gax.CallOption) (*computepb.SecurityPoliciesListPreconfiguredExpressionSetsResponse, error) {
	return c.internalClient.ListPreconfiguredExpressionSets(ctx, req, opts...)
}

// Patch patches the specified policy with the data included in the request. This cannot be used to be update the rules in the policy. Please use the per rule methods like addRule, patchRule, and removeRule instead.
func (c *SecurityPoliciesClient) Patch(ctx context.Context, req *computepb.PatchSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// PatchRule patches a rule at the specified priority.
func (c *SecurityPoliciesClient) PatchRule(ctx context.Context, req *computepb.PatchRuleSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.PatchRule(ctx, req, opts...)
}

// RemoveRule deletes a rule at the specified priority.
func (c *SecurityPoliciesClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.RemoveRule(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type securityPoliciesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSecurityPoliciesRESTClient creates a new security policies rest client.
//
// The SecurityPolicies API.
func NewSecurityPoliciesRESTClient(ctx context.Context, opts ...option.ClientOption) (*SecurityPoliciesClient, error) {
	clientOpts := append(defaultSecurityPoliciesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &securityPoliciesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &SecurityPoliciesClient{internalClient: c, CallOptions: &SecurityPoliciesCallOptions{}}, nil
}

func defaultSecurityPoliciesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *securityPoliciesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *securityPoliciesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *securityPoliciesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddRule inserts a rule into a security policy.
func (c *securityPoliciesRESTClient) AddRule(ctx context.Context, req *computepb.AddRuleSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyRuleResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/%v/addRule", req.GetProject(), req.GetSecurityPolicy())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Delete deletes the specified policy.
func (c *securityPoliciesRESTClient) Delete(ctx context.Context, req *computepb.DeleteSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/%v", req.GetProject(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get list all of the ordered rules present in a single specified policy.
func (c *securityPoliciesRESTClient) Get(ctx context.Context, req *computepb.GetSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicy, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/%v", req.GetProject(), req.GetSecurityPolicy())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.SecurityPolicy{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// GetRule gets a rule at the specified priority.
func (c *securityPoliciesRESTClient) GetRule(ctx context.Context, req *computepb.GetRuleSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicyRule, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/%v/getRule", req.GetProject(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.Priority != nil {
		params.Add("priority", fmt.Sprintf("%v", req.GetPriority()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.SecurityPolicyRule{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a new policy in the specified project using the data included in the request.
func (c *securityPoliciesRESTClient) Insert(ctx context.Context, req *computepb.InsertSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List list all the policies that have been configured for the specified project.
func (c *securityPoliciesRESTClient) List(ctx context.Context, req *computepb.ListSecurityPoliciesRequest, opts ...gax.CallOption) *SecurityPolicyIterator {
	it := &SecurityPolicyIterator{}
	req = proto.Clone(req).(*computepb.ListSecurityPoliciesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.SecurityPolicy, string, error) {
		resp := &computepb.SecurityPolicyList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ListPreconfiguredExpressionSets gets the current list of preconfigured Web Application Firewall (WAF) expressions.
func (c *securityPoliciesRESTClient) ListPreconfiguredExpressionSets(ctx context.Context, req *computepb.ListPreconfiguredExpressionSetsSecurityPoliciesRequest, opts ...gax.CallOption) (*computepb.SecurityPoliciesListPreconfiguredExpressionSetsResponse, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/listPreconfiguredExpressionSets", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.SecurityPoliciesListPreconfiguredExpressionSetsResponse{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Patch patches the specified policy with the data included in the request. This cannot be used to be update the rules in the policy. Please use the per rule methods like addRule, patchRule, and removeRule instead.
func (c *securityPoliciesRESTClient) Patch(ctx context.Context, req *computepb.PatchSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/%v", req.GetProject(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// PatchRule patches a rule at the specified priority.
func (c *securityPoliciesRESTClient) PatchRule(ctx context.Context, req *computepb.PatchRuleSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSecurityPolicyRuleResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/%v/patchRule", req.GetProject(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.Priority != nil {
		params.Add("priority", fmt.Sprintf("%v", req.GetPriority()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// RemoveRule deletes a rule at the specified priority.
func (c *securityPoliciesRESTClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleSecurityPolicyRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/securityPolicies/%v/removeRule", req.GetProject(), req.GetSecurityPolicy())

	params := url.Values{}
	if req != nil && req.Priority != nil {
		params.Add("priority", fmt.Sprintf("%v", req.GetPriority()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// SecurityPolicyIterator manages a stream of *computepb.SecurityPolicy.
type SecurityPolicyIterator struct {
	items    []*computepb.SecurityPolicy
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.SecurityPolicy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SecurityPolicyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SecurityPolicyIterator) Next() (*computepb.SecurityPolicy, error) {
	var item *computepb.SecurityPolicy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SecurityPolicyIterator) bufLen() int {
	return len(it.items)
}

func (it *SecurityPolicyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
