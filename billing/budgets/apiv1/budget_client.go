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

package budgets

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
	budgetspb "google.golang.org/genproto/googleapis/cloud/billing/budgets/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newBudgetClientHook clientHook

// BudgetCallOptions contains the retry settings for each method of BudgetClient.
type BudgetCallOptions struct {
	CreateBudget []gax.CallOption
	UpdateBudget []gax.CallOption
	GetBudget    []gax.CallOption
	ListBudgets  []gax.CallOption
	DeleteBudget []gax.CallOption
}

func defaultBudgetGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("billingbudgets.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("billingbudgets.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://billingbudgets.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBudgetCallOptions() *BudgetCallOptions {
	return &BudgetCallOptions{
		CreateBudget: []gax.CallOption{},
		UpdateBudget: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetBudget: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListBudgets: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteBudget: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalBudgetClient is an interface that defines the methods availaible from Cloud Billing Budget API.
type internalBudgetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateBudget(context.Context, *budgetspb.CreateBudgetRequest, ...gax.CallOption) (*budgetspb.Budget, error)
	UpdateBudget(context.Context, *budgetspb.UpdateBudgetRequest, ...gax.CallOption) (*budgetspb.Budget, error)
	GetBudget(context.Context, *budgetspb.GetBudgetRequest, ...gax.CallOption) (*budgetspb.Budget, error)
	ListBudgets(context.Context, *budgetspb.ListBudgetsRequest, ...gax.CallOption) *BudgetIterator
	DeleteBudget(context.Context, *budgetspb.DeleteBudgetRequest, ...gax.CallOption) error
}

// BudgetClient is a client for interacting with Cloud Billing Budget API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// BudgetService stores Cloud Billing budgets, which define a
// budget plan and rules to execute as we track spend against that plan.
type BudgetClient struct {
	// The internal transport-dependent client.
	internalClient internalBudgetClient

	// The call options for this service.
	CallOptions *BudgetCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BudgetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BudgetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BudgetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateBudget creates a new budget. See
// Quotas and limits (at https://cloud.google.com/billing/quotas)
// for more information on the limits of the number of budgets you can create.
func (c *BudgetClient) CreateBudget(ctx context.Context, req *budgetspb.CreateBudgetRequest, opts ...gax.CallOption) (*budgetspb.Budget, error) {
	return c.internalClient.CreateBudget(ctx, req, opts...)
}

// UpdateBudget updates a budget and returns the updated budget.
//
// WARNING: There are some fields exposed on the Google Cloud Console that
// aren’t available on this API. Budget fields that are not exposed in
// this API will not be changed by this method.
func (c *BudgetClient) UpdateBudget(ctx context.Context, req *budgetspb.UpdateBudgetRequest, opts ...gax.CallOption) (*budgetspb.Budget, error) {
	return c.internalClient.UpdateBudget(ctx, req, opts...)
}

// GetBudget returns a budget.
//
// WARNING: There are some fields exposed on the Google Cloud Console that
// aren’t available on this API. When reading from the API, you will not
// see these fields in the return value, though they may have been set
// in the Cloud Console.
func (c *BudgetClient) GetBudget(ctx context.Context, req *budgetspb.GetBudgetRequest, opts ...gax.CallOption) (*budgetspb.Budget, error) {
	return c.internalClient.GetBudget(ctx, req, opts...)
}

// ListBudgets returns a list of budgets for a billing account.
//
// WARNING: There are some fields exposed on the Google Cloud Console that
// aren’t available on this API. When reading from the API, you will not
// see these fields in the return value, though they may have been set
// in the Cloud Console.
func (c *BudgetClient) ListBudgets(ctx context.Context, req *budgetspb.ListBudgetsRequest, opts ...gax.CallOption) *BudgetIterator {
	return c.internalClient.ListBudgets(ctx, req, opts...)
}

// DeleteBudget deletes a budget. Returns successfully if already deleted.
func (c *BudgetClient) DeleteBudget(ctx context.Context, req *budgetspb.DeleteBudgetRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteBudget(ctx, req, opts...)
}

// budgetGRPCClient is a client for interacting with Cloud Billing Budget API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type budgetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BudgetClient
	CallOptions **BudgetCallOptions

	// The gRPC API client.
	budgetClient budgetspb.BudgetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBudgetClient creates a new budget service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// BudgetService stores Cloud Billing budgets, which define a
// budget plan and rules to execute as we track spend against that plan.
func NewBudgetClient(ctx context.Context, opts ...option.ClientOption) (*BudgetClient, error) {
	clientOpts := defaultBudgetGRPCClientOptions()
	if newBudgetClientHook != nil {
		hookOpts, err := newBudgetClientHook(ctx, clientHookParams{})
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
	client := BudgetClient{CallOptions: defaultBudgetCallOptions()}

	c := &budgetGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		budgetClient:     budgetspb.NewBudgetServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *budgetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *budgetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *budgetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *budgetGRPCClient) CreateBudget(ctx context.Context, req *budgetspb.CreateBudgetRequest, opts ...gax.CallOption) (*budgetspb.Budget, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateBudget[0:len((*c.CallOptions).CreateBudget):len((*c.CallOptions).CreateBudget)], opts...)
	var resp *budgetspb.Budget
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.budgetClient.CreateBudget(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *budgetGRPCClient) UpdateBudget(ctx context.Context, req *budgetspb.UpdateBudgetRequest, opts ...gax.CallOption) (*budgetspb.Budget, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "budget.name", url.QueryEscape(req.GetBudget().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateBudget[0:len((*c.CallOptions).UpdateBudget):len((*c.CallOptions).UpdateBudget)], opts...)
	var resp *budgetspb.Budget
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.budgetClient.UpdateBudget(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *budgetGRPCClient) GetBudget(ctx context.Context, req *budgetspb.GetBudgetRequest, opts ...gax.CallOption) (*budgetspb.Budget, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetBudget[0:len((*c.CallOptions).GetBudget):len((*c.CallOptions).GetBudget)], opts...)
	var resp *budgetspb.Budget
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.budgetClient.GetBudget(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *budgetGRPCClient) ListBudgets(ctx context.Context, req *budgetspb.ListBudgetsRequest, opts ...gax.CallOption) *BudgetIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListBudgets[0:len((*c.CallOptions).ListBudgets):len((*c.CallOptions).ListBudgets)], opts...)
	it := &BudgetIterator{}
	req = proto.Clone(req).(*budgetspb.ListBudgetsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*budgetspb.Budget, string, error) {
		resp := &budgetspb.ListBudgetsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.budgetClient.ListBudgets(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetBudgets(), resp.GetNextPageToken(), nil
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

func (c *budgetGRPCClient) DeleteBudget(ctx context.Context, req *budgetspb.DeleteBudgetRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteBudget[0:len((*c.CallOptions).DeleteBudget):len((*c.CallOptions).DeleteBudget)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.budgetClient.DeleteBudget(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// BudgetIterator manages a stream of *budgetspb.Budget.
type BudgetIterator struct {
	items    []*budgetspb.Budget
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
	InternalFetch func(pageSize int, pageToken string) (results []*budgetspb.Budget, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *BudgetIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *BudgetIterator) Next() (*budgetspb.Budget, error) {
	var item *budgetspb.Budget
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *BudgetIterator) bufLen() int {
	return len(it.items)
}

func (it *BudgetIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
