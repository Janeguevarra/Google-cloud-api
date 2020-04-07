// Copyright 2020 Google LLC
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

package errorreporting

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newErrorStatsClientHook clientHook

// ErrorStatsCallOptions contains the retry settings for each method of ErrorStatsClient.
type ErrorStatsCallOptions struct {
	ListGroupStats []gax.CallOption
	ListEvents     []gax.CallOption
	DeleteEvents   []gax.CallOption
}

func defaultErrorStatsClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("clouderrorreporting.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultErrorStatsCallOptions() *ErrorStatsCallOptions {
	return &ErrorStatsCallOptions{
		ListGroupStats: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListEvents: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteEvents: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// ErrorStatsClient is a client for interacting with Stackdriver Error Reporting API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ErrorStatsClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	errorStatsClient clouderrorreportingpb.ErrorStatsServiceClient

	// The call options for this service.
	CallOptions *ErrorStatsCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewErrorStatsClient creates a new error stats service client.
//
// An API for retrieving and managing error statistics as well as data for
// individual events.
func NewErrorStatsClient(ctx context.Context, opts ...option.ClientOption) (*ErrorStatsClient, error) {
	clientOpts := defaultErrorStatsClientOptions()

	if newErrorStatsClientHook != nil {
		hookOpts, err := newErrorStatsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &ErrorStatsClient{
		connPool:    connPool,
		CallOptions: defaultErrorStatsCallOptions(),

		errorStatsClient: clouderrorreportingpb.NewErrorStatsServiceClient(connPool),
	}
	c.SetGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ErrorStatsClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ErrorStatsClient) Close() error {
	return c.connPool.Close()
}

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ErrorStatsClient) SetGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListGroupStats lists the specified groups.
func (c *ErrorStatsClient) ListGroupStats(ctx context.Context, req *clouderrorreportingpb.ListGroupStatsRequest, opts ...gax.CallOption) *ErrorGroupStatsIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListGroupStats[0:len(c.CallOptions.ListGroupStats):len(c.CallOptions.ListGroupStats)], opts...)
	it := &ErrorGroupStatsIterator{}
	req = proto.Clone(req).(*clouderrorreportingpb.ListGroupStatsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*clouderrorreportingpb.ErrorGroupStats, string, error) {
		var resp *clouderrorreportingpb.ListGroupStatsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.errorStatsClient.ListGroupStats(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.ErrorGroupStats, resp.NextPageToken, nil
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
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// ListEvents lists the specified events.
func (c *ErrorStatsClient) ListEvents(ctx context.Context, req *clouderrorreportingpb.ListEventsRequest, opts ...gax.CallOption) *ErrorEventIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListEvents[0:len(c.CallOptions.ListEvents):len(c.CallOptions.ListEvents)], opts...)
	it := &ErrorEventIterator{}
	req = proto.Clone(req).(*clouderrorreportingpb.ListEventsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*clouderrorreportingpb.ErrorEvent, string, error) {
		var resp *clouderrorreportingpb.ListEventsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.errorStatsClient.ListEvents(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.ErrorEvents, resp.NextPageToken, nil
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
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// DeleteEvents deletes all error events of a given project.
func (c *ErrorStatsClient) DeleteEvents(ctx context.Context, req *clouderrorreportingpb.DeleteEventsRequest, opts ...gax.CallOption) (*clouderrorreportingpb.DeleteEventsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteEvents[0:len(c.CallOptions.DeleteEvents):len(c.CallOptions.DeleteEvents)], opts...)
	var resp *clouderrorreportingpb.DeleteEventsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.errorStatsClient.DeleteEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ErrorEventIterator manages a stream of *clouderrorreportingpb.ErrorEvent.
type ErrorEventIterator struct {
	items    []*clouderrorreportingpb.ErrorEvent
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
	InternalFetch func(pageSize int, pageToken string) (results []*clouderrorreportingpb.ErrorEvent, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ErrorEventIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ErrorEventIterator) Next() (*clouderrorreportingpb.ErrorEvent, error) {
	var item *clouderrorreportingpb.ErrorEvent
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ErrorEventIterator) bufLen() int {
	return len(it.items)
}

func (it *ErrorEventIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ErrorGroupStatsIterator manages a stream of *clouderrorreportingpb.ErrorGroupStats.
type ErrorGroupStatsIterator struct {
	items    []*clouderrorreportingpb.ErrorGroupStats
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
	InternalFetch func(pageSize int, pageToken string) (results []*clouderrorreportingpb.ErrorGroupStats, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ErrorGroupStatsIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ErrorGroupStatsIterator) Next() (*clouderrorreportingpb.ErrorGroupStats, error) {
	var item *clouderrorreportingpb.ErrorGroupStats
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ErrorGroupStatsIterator) bufLen() int {
	return len(it.items)
}

func (it *ErrorGroupStatsIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
