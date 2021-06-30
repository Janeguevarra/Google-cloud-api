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

package retail

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newCatalogClientHook clientHook

// CatalogCallOptions contains the retry settings for each method of CatalogClient.
type CatalogCallOptions struct {
	ListCatalogs  []gax.CallOption
	UpdateCatalog []gax.CallOption
}

func defaultCatalogGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("retail.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("retail.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://retail.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCatalogCallOptions() *CatalogCallOptions {
	return &CatalogCallOptions{
		ListCatalogs: []gax.CallOption{
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
		UpdateCatalog: []gax.CallOption{
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

// internalCatalogClient is an interface that defines the methods availaible from Retail API.
type internalCatalogClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListCatalogs(context.Context, *retailpb.ListCatalogsRequest, ...gax.CallOption) *CatalogIterator
	UpdateCatalog(context.Context, *retailpb.UpdateCatalogRequest, ...gax.CallOption) (*retailpb.Catalog, error)
}

// CatalogClient is a client for interacting with Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing catalog configuration.
type CatalogClient struct {
	// The internal transport-dependent client.
	internalClient internalCatalogClient

	// The call options for this service.
	CallOptions *CatalogCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CatalogClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CatalogClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CatalogClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListCatalogs lists all the Catalogs associated with
// the project.
func (c *CatalogClient) ListCatalogs(ctx context.Context, req *retailpb.ListCatalogsRequest, opts ...gax.CallOption) *CatalogIterator {
	return c.internalClient.ListCatalogs(ctx, req, opts...)
}

// UpdateCatalog updates the Catalogs.
func (c *CatalogClient) UpdateCatalog(ctx context.Context, req *retailpb.UpdateCatalogRequest, opts ...gax.CallOption) (*retailpb.Catalog, error) {
	return c.internalClient.UpdateCatalog(ctx, req, opts...)
}

// catalogGRPCClient is a client for interacting with Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type catalogGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CatalogClient
	CallOptions **CatalogCallOptions

	// The gRPC API client.
	catalogClient retailpb.CatalogServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCatalogClient creates a new catalog service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing catalog configuration.
func NewCatalogClient(ctx context.Context, opts ...option.ClientOption) (*CatalogClient, error) {
	clientOpts := defaultCatalogGRPCClientOptions()
	if newCatalogClientHook != nil {
		hookOpts, err := newCatalogClientHook(ctx, clientHookParams{})
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
	client := CatalogClient{CallOptions: defaultCatalogCallOptions()}

	c := &catalogGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		catalogClient:    retailpb.NewCatalogServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *catalogGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *catalogGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *catalogGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *catalogGRPCClient) ListCatalogs(ctx context.Context, req *retailpb.ListCatalogsRequest, opts ...gax.CallOption) *CatalogIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListCatalogs[0:len((*c.CallOptions).ListCatalogs):len((*c.CallOptions).ListCatalogs)], opts...)
	it := &CatalogIterator{}
	req = proto.Clone(req).(*retailpb.ListCatalogsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*retailpb.Catalog, string, error) {
		var resp *retailpb.ListCatalogsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.catalogClient.ListCatalogs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCatalogs(), resp.GetNextPageToken(), nil
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

func (c *catalogGRPCClient) UpdateCatalog(ctx context.Context, req *retailpb.UpdateCatalogRequest, opts ...gax.CallOption) (*retailpb.Catalog, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "catalog.name", url.QueryEscape(req.GetCatalog().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateCatalog[0:len((*c.CallOptions).UpdateCatalog):len((*c.CallOptions).UpdateCatalog)], opts...)
	var resp *retailpb.Catalog
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.catalogClient.UpdateCatalog(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CatalogIterator manages a stream of *retailpb.Catalog.
type CatalogIterator struct {
	items    []*retailpb.Catalog
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
	InternalFetch func(pageSize int, pageToken string) (results []*retailpb.Catalog, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CatalogIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CatalogIterator) Next() (*retailpb.Catalog, error) {
	var item *retailpb.Catalog
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CatalogIterator) bufLen() int {
	return len(it.items)
}

func (it *CatalogIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
