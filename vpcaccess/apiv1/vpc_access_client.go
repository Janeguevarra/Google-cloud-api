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

package vpcaccess

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	vpcaccesspb "google.golang.org/genproto/googleapis/cloud/vpcaccess/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateConnector []gax.CallOption
	GetConnector    []gax.CallOption
	ListConnectors  []gax.CallOption
	DeleteConnector []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("vpcaccess.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("vpcaccess.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://vpcaccess.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateConnector: []gax.CallOption{},
		GetConnector:    []gax.CallOption{},
		ListConnectors:  []gax.CallOption{},
		DeleteConnector: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from Serverless VPC Access API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateConnector(context.Context, *vpcaccesspb.CreateConnectorRequest, ...gax.CallOption) (*CreateConnectorOperation, error)
	CreateConnectorOperation(name string) *CreateConnectorOperation
	GetConnector(context.Context, *vpcaccesspb.GetConnectorRequest, ...gax.CallOption) (*vpcaccesspb.Connector, error)
	ListConnectors(context.Context, *vpcaccesspb.ListConnectorsRequest, ...gax.CallOption) *ConnectorIterator
	DeleteConnector(context.Context, *vpcaccesspb.DeleteConnectorRequest, ...gax.CallOption) (*DeleteConnectorOperation, error)
	DeleteConnectorOperation(name string) *DeleteConnectorOperation
}

// Client is a client for interacting with Serverless VPC Access API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Serverless VPC Access API allows users to create and manage connectors for
// App Engine, Cloud Functions and Cloud Run to have internal connections to
// Virtual Private Cloud networks.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateConnector creates a Serverless VPC Access connector, returns an operation.
func (c *Client) CreateConnector(ctx context.Context, req *vpcaccesspb.CreateConnectorRequest, opts ...gax.CallOption) (*CreateConnectorOperation, error) {
	return c.internalClient.CreateConnector(ctx, req, opts...)
}

// CreateConnectorOperation returns a new CreateConnectorOperation from a given name.
// The name must be that of a previously created CreateConnectorOperation, possibly from a different process.
func (c *Client) CreateConnectorOperation(name string) *CreateConnectorOperation {
	return c.internalClient.CreateConnectorOperation(name)
}

// GetConnector gets a Serverless VPC Access connector. Returns NOT_FOUND if the resource
// does not exist.
func (c *Client) GetConnector(ctx context.Context, req *vpcaccesspb.GetConnectorRequest, opts ...gax.CallOption) (*vpcaccesspb.Connector, error) {
	return c.internalClient.GetConnector(ctx, req, opts...)
}

// ListConnectors lists Serverless VPC Access connectors.
func (c *Client) ListConnectors(ctx context.Context, req *vpcaccesspb.ListConnectorsRequest, opts ...gax.CallOption) *ConnectorIterator {
	return c.internalClient.ListConnectors(ctx, req, opts...)
}

// DeleteConnector deletes a Serverless VPC Access connector. Returns NOT_FOUND if the
// resource does not exist.
func (c *Client) DeleteConnector(ctx context.Context, req *vpcaccesspb.DeleteConnectorRequest, opts ...gax.CallOption) (*DeleteConnectorOperation, error) {
	return c.internalClient.DeleteConnector(ctx, req, opts...)
}

// DeleteConnectorOperation returns a new DeleteConnectorOperation from a given name.
// The name must be that of a previously created DeleteConnectorOperation, possibly from a different process.
func (c *Client) DeleteConnectorOperation(name string) *DeleteConnectorOperation {
	return c.internalClient.DeleteConnectorOperation(name)
}

// gRPCClient is a client for interacting with Serverless VPC Access API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client vpcaccesspb.VpcAccessServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new vpc access service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Serverless VPC Access API allows users to create and manage connectors for
// App Engine, Cloud Functions and Cloud Run to have internal connections to
// Virtual Private Cloud networks.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           vpcaccesspb.NewVpcAccessServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) CreateConnector(ctx context.Context, req *vpcaccesspb.CreateConnectorRequest, opts ...gax.CallOption) (*CreateConnectorOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateConnector[0:len((*c.CallOptions).CreateConnector):len((*c.CallOptions).CreateConnector)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateConnector(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) GetConnector(ctx context.Context, req *vpcaccesspb.GetConnectorRequest, opts ...gax.CallOption) (*vpcaccesspb.Connector, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetConnector[0:len((*c.CallOptions).GetConnector):len((*c.CallOptions).GetConnector)], opts...)
	var resp *vpcaccesspb.Connector
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetConnector(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListConnectors(ctx context.Context, req *vpcaccesspb.ListConnectorsRequest, opts ...gax.CallOption) *ConnectorIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListConnectors[0:len((*c.CallOptions).ListConnectors):len((*c.CallOptions).ListConnectors)], opts...)
	it := &ConnectorIterator{}
	req = proto.Clone(req).(*vpcaccesspb.ListConnectorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*vpcaccesspb.Connector, string, error) {
		var resp *vpcaccesspb.ListConnectorsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListConnectors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetConnectors(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) DeleteConnector(ctx context.Context, req *vpcaccesspb.DeleteConnectorRequest, opts ...gax.CallOption) (*DeleteConnectorOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteConnector[0:len((*c.CallOptions).DeleteConnector):len((*c.CallOptions).DeleteConnector)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DeleteConnector(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateConnectorOperation manages a long-running operation from CreateConnector.
type CreateConnectorOperation struct {
	lro *longrunning.Operation
}

// CreateConnectorOperation returns a new CreateConnectorOperation from a given name.
// The name must be that of a previously created CreateConnectorOperation, possibly from a different process.
func (c *gRPCClient) CreateConnectorOperation(name string) *CreateConnectorOperation {
	return &CreateConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateConnectorOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*vpcaccesspb.Connector, error) {
	var resp vpcaccesspb.Connector
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateConnectorOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*vpcaccesspb.Connector, error) {
	var resp vpcaccesspb.Connector
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateConnectorOperation) Metadata() (*vpcaccesspb.OperationMetadata, error) {
	var meta vpcaccesspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateConnectorOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateConnectorOperation) Name() string {
	return op.lro.Name()
}

// DeleteConnectorOperation manages a long-running operation from DeleteConnector.
type DeleteConnectorOperation struct {
	lro *longrunning.Operation
}

// DeleteConnectorOperation returns a new DeleteConnectorOperation from a given name.
// The name must be that of a previously created DeleteConnectorOperation, possibly from a different process.
func (c *gRPCClient) DeleteConnectorOperation(name string) *DeleteConnectorOperation {
	return &DeleteConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteConnectorOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteConnectorOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteConnectorOperation) Metadata() (*vpcaccesspb.OperationMetadata, error) {
	var meta vpcaccesspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteConnectorOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteConnectorOperation) Name() string {
	return op.lro.Name()
}

// ConnectorIterator manages a stream of *vpcaccesspb.Connector.
type ConnectorIterator struct {
	items    []*vpcaccesspb.Connector
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
	InternalFetch func(pageSize int, pageToken string) (results []*vpcaccesspb.Connector, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ConnectorIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ConnectorIterator) Next() (*vpcaccesspb.Connector, error) {
	var item *vpcaccesspb.Connector
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ConnectorIterator) bufLen() int {
	return len(it.items)
}

func (it *ConnectorIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
