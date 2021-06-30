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

package dataflow

import (
	"context"
	"math"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dataflowpb "google.golang.org/genproto/googleapis/dataflow/v1beta3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newMessagesV1Beta3ClientHook clientHook

// MessagesV1Beta3CallOptions contains the retry settings for each method of MessagesV1Beta3Client.
type MessagesV1Beta3CallOptions struct {
	ListJobMessages []gax.CallOption
}

func defaultMessagesV1Beta3GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMessagesV1Beta3CallOptions() *MessagesV1Beta3CallOptions {
	return &MessagesV1Beta3CallOptions{
		ListJobMessages: []gax.CallOption{},
	}
}

// internalMessagesV1Beta3Client is an interface that defines the methods availaible from Dataflow API.
type internalMessagesV1Beta3Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListJobMessages(context.Context, *dataflowpb.ListJobMessagesRequest, ...gax.CallOption) *JobMessageIterator
}

// MessagesV1Beta3Client is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Dataflow Messages API is used for monitoring the progress of
// Dataflow jobs.
type MessagesV1Beta3Client struct {
	// The internal transport-dependent client.
	internalClient internalMessagesV1Beta3Client

	// The call options for this service.
	CallOptions *MessagesV1Beta3CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MessagesV1Beta3Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MessagesV1Beta3Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *MessagesV1Beta3Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListJobMessages request the job status.
//
// To request the status of a job, we recommend using
// projects.locations.jobs.messages.list with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.messages.list is not recommended, as you can only request
// the status of jobs that are running in us-central1.
func (c *MessagesV1Beta3Client) ListJobMessages(ctx context.Context, req *dataflowpb.ListJobMessagesRequest, opts ...gax.CallOption) *JobMessageIterator {
	return c.internalClient.ListJobMessages(ctx, req, opts...)
}

// messagesV1Beta3GRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type messagesV1Beta3GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing MessagesV1Beta3Client
	CallOptions **MessagesV1Beta3CallOptions

	// The gRPC API client.
	messagesV1Beta3Client dataflowpb.MessagesV1Beta3Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMessagesV1Beta3Client creates a new messages v1 beta3 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Dataflow Messages API is used for monitoring the progress of
// Dataflow jobs.
func NewMessagesV1Beta3Client(ctx context.Context, opts ...option.ClientOption) (*MessagesV1Beta3Client, error) {
	clientOpts := defaultMessagesV1Beta3GRPCClientOptions()
	if newMessagesV1Beta3ClientHook != nil {
		hookOpts, err := newMessagesV1Beta3ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MessagesV1Beta3Client{CallOptions: defaultMessagesV1Beta3CallOptions()}

	c := &messagesV1Beta3GRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		messagesV1Beta3Client: dataflowpb.NewMessagesV1Beta3Client(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *messagesV1Beta3GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *messagesV1Beta3GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *messagesV1Beta3GRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *messagesV1Beta3GRPCClient) ListJobMessages(ctx context.Context, req *dataflowpb.ListJobMessagesRequest, opts ...gax.CallOption) *JobMessageIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListJobMessages[0:len((*c.CallOptions).ListJobMessages):len((*c.CallOptions).ListJobMessages)], opts...)
	it := &JobMessageIterator{}
	req = proto.Clone(req).(*dataflowpb.ListJobMessagesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.JobMessage, string, error) {
		var resp *dataflowpb.ListJobMessagesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.messagesV1Beta3Client.ListJobMessages(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetJobMessages(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

// JobMessageIterator manages a stream of *dataflowpb.JobMessage.
type JobMessageIterator struct {
	items    []*dataflowpb.JobMessage
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataflowpb.JobMessage, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobMessageIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobMessageIterator) Next() (*dataflowpb.JobMessage, error) {
	var item *dataflowpb.JobMessage
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobMessageIterator) bufLen() int {
	return len(it.items)
}

func (it *JobMessageIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
