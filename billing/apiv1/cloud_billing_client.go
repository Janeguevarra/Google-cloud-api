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

package billing

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
	billingpb "google.golang.org/genproto/googleapis/cloud/billing/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCloudBillingClientHook clientHook

// CloudBillingCallOptions contains the retry settings for each method of CloudBillingClient.
type CloudBillingCallOptions struct {
	GetBillingAccount        []gax.CallOption
	ListBillingAccounts      []gax.CallOption
	UpdateBillingAccount     []gax.CallOption
	CreateBillingAccount     []gax.CallOption
	ListProjectBillingInfo   []gax.CallOption
	GetProjectBillingInfo    []gax.CallOption
	UpdateProjectBillingInfo []gax.CallOption
	GetIamPolicy             []gax.CallOption
	SetIamPolicy             []gax.CallOption
	TestIamPermissions       []gax.CallOption
}

func defaultCloudBillingClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("cloudbilling.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCloudBillingCallOptions() *CloudBillingCallOptions {
	return &CloudBillingCallOptions{
		GetBillingAccount: []gax.CallOption{
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
		ListBillingAccounts: []gax.CallOption{
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
		UpdateBillingAccount: []gax.CallOption{
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
		CreateBillingAccount: []gax.CallOption{},
		ListProjectBillingInfo: []gax.CallOption{
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
		GetProjectBillingInfo: []gax.CallOption{
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
		UpdateProjectBillingInfo: []gax.CallOption{
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
		GetIamPolicy: []gax.CallOption{
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
		SetIamPolicy: []gax.CallOption{
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
		TestIamPermissions: []gax.CallOption{
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

// CloudBillingClient is a client for interacting with Cloud Billing API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type CloudBillingClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	cloudBillingClient billingpb.CloudBillingClient

	// The call options for this service.
	CallOptions *CloudBillingCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCloudBillingClient creates a new cloud billing client.
//
// Retrieves GCP Console billing accounts and associates them with projects.
func NewCloudBillingClient(ctx context.Context, opts ...option.ClientOption) (*CloudBillingClient, error) {
	clientOpts := defaultCloudBillingClientOptions()

	if newCloudBillingClientHook != nil {
		hookOpts, err := newCloudBillingClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &CloudBillingClient{
		connPool:    connPool,
		CallOptions: defaultCloudBillingCallOptions(),

		cloudBillingClient: billingpb.NewCloudBillingClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CloudBillingClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CloudBillingClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CloudBillingClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// GetBillingAccount gets information about a billing account. The current authenticated user
// must be a viewer of the billing
// account (at https://cloud.google.com/billing/docs/how-to/billing-access).
func (c *CloudBillingClient) GetBillingAccount(ctx context.Context, req *billingpb.GetBillingAccountRequest, opts ...gax.CallOption) (*billingpb.BillingAccount, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetBillingAccount[0:len(c.CallOptions.GetBillingAccount):len(c.CallOptions.GetBillingAccount)], opts...)
	var resp *billingpb.BillingAccount
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.GetBillingAccount(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListBillingAccounts lists the billing accounts that the current authenticated user has
// permission to
// view (at https://cloud.google.com/billing/docs/how-to/billing-access).
func (c *CloudBillingClient) ListBillingAccounts(ctx context.Context, req *billingpb.ListBillingAccountsRequest, opts ...gax.CallOption) *BillingAccountIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListBillingAccounts[0:len(c.CallOptions.ListBillingAccounts):len(c.CallOptions.ListBillingAccounts)], opts...)
	it := &BillingAccountIterator{}
	req = proto.Clone(req).(*billingpb.ListBillingAccountsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*billingpb.BillingAccount, string, error) {
		var resp *billingpb.ListBillingAccountsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.cloudBillingClient.ListBillingAccounts(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.BillingAccounts, resp.NextPageToken, nil
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

// UpdateBillingAccount updates a billing account’s fields.
// Currently the only field that can be edited is display_name.
// The current authenticated user must have the billing.accounts.update
// IAM permission, which is typically given to the
// administrator (at https://cloud.google.com/billing/docs/how-to/billing-access)
// of the billing account.
func (c *CloudBillingClient) UpdateBillingAccount(ctx context.Context, req *billingpb.UpdateBillingAccountRequest, opts ...gax.CallOption) (*billingpb.BillingAccount, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateBillingAccount[0:len(c.CallOptions.UpdateBillingAccount):len(c.CallOptions.UpdateBillingAccount)], opts...)
	var resp *billingpb.BillingAccount
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.UpdateBillingAccount(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateBillingAccount creates a billing account.
// This method can only be used to create
// billing subaccounts (at https://cloud.google.com/billing/docs/concepts)
// by GCP resellers.
// When creating a subaccount, the current authenticated user must have the
// billing.accounts.update IAM permission on the master account, which is
// typically given to billing account
// administrators (at https://cloud.google.com/billing/docs/how-to/billing-access).
// This method will return an error if the master account has not been
// provisioned as a reseller account.
func (c *CloudBillingClient) CreateBillingAccount(ctx context.Context, req *billingpb.CreateBillingAccountRequest, opts ...gax.CallOption) (*billingpb.BillingAccount, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateBillingAccount[0:len(c.CallOptions.CreateBillingAccount):len(c.CallOptions.CreateBillingAccount)], opts...)
	var resp *billingpb.BillingAccount
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.CreateBillingAccount(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListProjectBillingInfo lists the projects associated with a billing account. The current
// authenticated user must have the billing.resourceAssociations.list IAM
// permission, which is often given to billing account
// viewers (at https://cloud.google.com/billing/docs/how-to/billing-access).
func (c *CloudBillingClient) ListProjectBillingInfo(ctx context.Context, req *billingpb.ListProjectBillingInfoRequest, opts ...gax.CallOption) *ProjectBillingInfoIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListProjectBillingInfo[0:len(c.CallOptions.ListProjectBillingInfo):len(c.CallOptions.ListProjectBillingInfo)], opts...)
	it := &ProjectBillingInfoIterator{}
	req = proto.Clone(req).(*billingpb.ListProjectBillingInfoRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*billingpb.ProjectBillingInfo, string, error) {
		var resp *billingpb.ListProjectBillingInfoResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.cloudBillingClient.ListProjectBillingInfo(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.ProjectBillingInfo, resp.NextPageToken, nil
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

// GetProjectBillingInfo gets the billing information for a project. The current authenticated user
// must have permission to view the
// project (at https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo).
func (c *CloudBillingClient) GetProjectBillingInfo(ctx context.Context, req *billingpb.GetProjectBillingInfoRequest, opts ...gax.CallOption) (*billingpb.ProjectBillingInfo, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetProjectBillingInfo[0:len(c.CallOptions.GetProjectBillingInfo):len(c.CallOptions.GetProjectBillingInfo)], opts...)
	var resp *billingpb.ProjectBillingInfo
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.GetProjectBillingInfo(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateProjectBillingInfo sets or updates the billing account associated with a project. You specify
// the new billing account by setting the billing_account_name in the
// ProjectBillingInfo resource to the resource name of a billing account.
// Associating a project with an open billing account enables billing on the
// project and allows charges for resource usage. If the project already had a
// billing account, this method changes the billing account used for resource
// usage charges.
//
// Note: Incurred charges that have not yet been reported in the transaction
// history of the GCP Console might be billed to the new billing
// account, even if the charge occurred before the new billing account was
// assigned to the project.
//
// The current authenticated user must have ownership privileges for both the
// project (at https://cloud.google.com/docs/permissions-overview#h.bgs0oxofvnoo) and the billing
// account (at https://cloud.google.com/billing/docs/how-to/billing-access).
//
// You can disable billing on the project by setting the
// billing_account_name field to empty. This action disassociates the
// current billing account from the project. Any billable activity of your
// in-use services will stop, and your application could stop functioning as
// expected. Any unbilled charges to date will be billed to the previously
// associated account. The current authenticated user must be either an owner
// of the project or an owner of the billing account for the project.
//
// Note that associating a project with a closed billing account will have
// much the same effect as disabling billing on the project: any paid
// resources used by the project will be shut down. Thus, unless you wish to
// disable billing, you should always call this method with the name of an
// open billing account.
func (c *CloudBillingClient) UpdateProjectBillingInfo(ctx context.Context, req *billingpb.UpdateProjectBillingInfoRequest, opts ...gax.CallOption) (*billingpb.ProjectBillingInfo, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateProjectBillingInfo[0:len(c.CallOptions.UpdateProjectBillingInfo):len(c.CallOptions.UpdateProjectBillingInfo)], opts...)
	var resp *billingpb.ProjectBillingInfo
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.UpdateProjectBillingInfo(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetIamPolicy gets the access control policy for a billing account.
// The caller must have the billing.accounts.getIamPolicy permission on the
// account, which is often given to billing account
// viewers (at https://cloud.google.com/billing/docs/how-to/billing-access).
func (c *CloudBillingClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetIamPolicy[0:len(c.CallOptions.GetIamPolicy):len(c.CallOptions.GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SetIamPolicy sets the access control policy for a billing account. Replaces any existing
// policy.
// The caller must have the billing.accounts.setIamPolicy permission on the
// account, which is often given to billing account
// administrators (at https://cloud.google.com/billing/docs/how-to/billing-access).
func (c *CloudBillingClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SetIamPolicy[0:len(c.CallOptions.SetIamPolicy):len(c.CallOptions.SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TestIamPermissions tests the access control policy for a billing account. This method takes
// the resource and a set of permissions as input and returns the subset of
// the input permissions that the caller is allowed for that resource.
func (c *CloudBillingClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.TestIamPermissions[0:len(c.CallOptions.TestIamPermissions):len(c.CallOptions.TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudBillingClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BillingAccountIterator manages a stream of *billingpb.BillingAccount.
type BillingAccountIterator struct {
	items    []*billingpb.BillingAccount
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
	InternalFetch func(pageSize int, pageToken string) (results []*billingpb.BillingAccount, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *BillingAccountIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *BillingAccountIterator) Next() (*billingpb.BillingAccount, error) {
	var item *billingpb.BillingAccount
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *BillingAccountIterator) bufLen() int {
	return len(it.items)
}

func (it *BillingAccountIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ProjectBillingInfoIterator manages a stream of *billingpb.ProjectBillingInfo.
type ProjectBillingInfoIterator struct {
	items    []*billingpb.ProjectBillingInfo
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
	InternalFetch func(pageSize int, pageToken string) (results []*billingpb.ProjectBillingInfo, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ProjectBillingInfoIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ProjectBillingInfoIterator) Next() (*billingpb.ProjectBillingInfo, error) {
	var item *billingpb.ProjectBillingInfo
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ProjectBillingInfoIterator) bufLen() int {
	return len(it.items)
}

func (it *ProjectBillingInfoIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
