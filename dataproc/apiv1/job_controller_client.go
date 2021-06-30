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

package dataproc

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
	dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newJobControllerClientHook clientHook

// JobControllerCallOptions contains the retry settings for each method of JobControllerClient.
type JobControllerCallOptions struct {
	SubmitJob            []gax.CallOption
	SubmitJobAsOperation []gax.CallOption
	GetJob               []gax.CallOption
	ListJobs             []gax.CallOption
	UpdateJob            []gax.CallOption
	CancelJob            []gax.CallOption
	DeleteJob            []gax.CallOption
}

func defaultJobControllerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataproc.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataproc.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataproc.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultJobControllerCallOptions() *JobControllerCallOptions {
	return &JobControllerCallOptions{
		SubmitJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SubmitJobAsOperation: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListJobs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CancelJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
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

// internalJobControllerClient is an interface that defines the methods availaible from Cloud Dataproc API.
type internalJobControllerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SubmitJob(context.Context, *dataprocpb.SubmitJobRequest, ...gax.CallOption) (*dataprocpb.Job, error)
	SubmitJobAsOperation(context.Context, *dataprocpb.SubmitJobRequest, ...gax.CallOption) (*SubmitJobAsOperationOperation, error)
	SubmitJobAsOperationOperation(name string) *SubmitJobAsOperationOperation
	GetJob(context.Context, *dataprocpb.GetJobRequest, ...gax.CallOption) (*dataprocpb.Job, error)
	ListJobs(context.Context, *dataprocpb.ListJobsRequest, ...gax.CallOption) *JobIterator
	UpdateJob(context.Context, *dataprocpb.UpdateJobRequest, ...gax.CallOption) (*dataprocpb.Job, error)
	CancelJob(context.Context, *dataprocpb.CancelJobRequest, ...gax.CallOption) (*dataprocpb.Job, error)
	DeleteJob(context.Context, *dataprocpb.DeleteJobRequest, ...gax.CallOption) error
}

// JobControllerClient is a client for interacting with Cloud Dataproc API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The JobController provides methods to manage jobs.
type JobControllerClient struct {
	// The internal transport-dependent client.
	internalClient internalJobControllerClient

	// The call options for this service.
	CallOptions *JobControllerCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *JobControllerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *JobControllerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *JobControllerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SubmitJob submits a job to a cluster.
func (c *JobControllerClient) SubmitJob(ctx context.Context, req *dataprocpb.SubmitJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	return c.internalClient.SubmitJob(ctx, req, opts...)
}

// SubmitJobAsOperation submits job to a cluster.
func (c *JobControllerClient) SubmitJobAsOperation(ctx context.Context, req *dataprocpb.SubmitJobRequest, opts ...gax.CallOption) (*SubmitJobAsOperationOperation, error) {
	return c.internalClient.SubmitJobAsOperation(ctx, req, opts...)
}

// SubmitJobAsOperationOperation returns a new SubmitJobAsOperationOperation from a given name.
// The name must be that of a previously created SubmitJobAsOperationOperation, possibly from a different process.
func (c *JobControllerClient) SubmitJobAsOperationOperation(name string) *SubmitJobAsOperationOperation {
	return c.internalClient.SubmitJobAsOperationOperation(name)
}

// GetJob gets the resource representation for a job in a project.
func (c *JobControllerClient) GetJob(ctx context.Context, req *dataprocpb.GetJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	return c.internalClient.GetJob(ctx, req, opts...)
}

// ListJobs lists regions/{region}/jobs in a project.
func (c *JobControllerClient) ListJobs(ctx context.Context, req *dataprocpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	return c.internalClient.ListJobs(ctx, req, opts...)
}

// UpdateJob updates a job in a project.
func (c *JobControllerClient) UpdateJob(ctx context.Context, req *dataprocpb.UpdateJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	return c.internalClient.UpdateJob(ctx, req, opts...)
}

// CancelJob starts a job cancellation request. To access the job resource
// after cancellation, call
// regions/{region}/jobs.list (at https://cloud.google.com/dataproc/docs/reference/rest/v1/projects.regions.jobs/list)
// or
// regions/{region}/jobs.get (at https://cloud.google.com/dataproc/docs/reference/rest/v1/projects.regions.jobs/get).
func (c *JobControllerClient) CancelJob(ctx context.Context, req *dataprocpb.CancelJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	return c.internalClient.CancelJob(ctx, req, opts...)
}

// DeleteJob deletes the job from the project. If the job is active, the delete fails,
// and the response returns FAILED_PRECONDITION.
func (c *JobControllerClient) DeleteJob(ctx context.Context, req *dataprocpb.DeleteJobRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteJob(ctx, req, opts...)
}

// jobControllerGRPCClient is a client for interacting with Cloud Dataproc API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type jobControllerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing JobControllerClient
	CallOptions **JobControllerCallOptions

	// The gRPC API client.
	jobControllerClient dataprocpb.JobControllerClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewJobControllerClient creates a new job controller client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The JobController provides methods to manage jobs.
func NewJobControllerClient(ctx context.Context, opts ...option.ClientOption) (*JobControllerClient, error) {
	clientOpts := defaultJobControllerGRPCClientOptions()
	if newJobControllerClientHook != nil {
		hookOpts, err := newJobControllerClientHook(ctx, clientHookParams{})
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
	client := JobControllerClient{CallOptions: defaultJobControllerCallOptions()}

	c := &jobControllerGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		jobControllerClient: dataprocpb.NewJobControllerClient(connPool),
		CallOptions:         &client.CallOptions,
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
func (c *jobControllerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *jobControllerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *jobControllerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *jobControllerGRPCClient) SubmitJob(ctx context.Context, req *dataprocpb.SubmitJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 900000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SubmitJob[0:len((*c.CallOptions).SubmitJob):len((*c.CallOptions).SubmitJob)], opts...)
	var resp *dataprocpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobControllerClient.SubmitJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobControllerGRPCClient) SubmitJobAsOperation(ctx context.Context, req *dataprocpb.SubmitJobRequest, opts ...gax.CallOption) (*SubmitJobAsOperationOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 900000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SubmitJobAsOperation[0:len((*c.CallOptions).SubmitJobAsOperation):len((*c.CallOptions).SubmitJobAsOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobControllerClient.SubmitJobAsOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &SubmitJobAsOperationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *jobControllerGRPCClient) GetJob(ctx context.Context, req *dataprocpb.GetJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 900000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "job_id", url.QueryEscape(req.GetJobId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetJob[0:len((*c.CallOptions).GetJob):len((*c.CallOptions).GetJob)], opts...)
	var resp *dataprocpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobControllerClient.GetJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobControllerGRPCClient) ListJobs(ctx context.Context, req *dataprocpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListJobs[0:len((*c.CallOptions).ListJobs):len((*c.CallOptions).ListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*dataprocpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataprocpb.Job, string, error) {
		var resp *dataprocpb.ListJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.jobControllerClient.ListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

func (c *jobControllerGRPCClient) UpdateJob(ctx context.Context, req *dataprocpb.UpdateJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 900000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "job_id", url.QueryEscape(req.GetJobId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateJob[0:len((*c.CallOptions).UpdateJob):len((*c.CallOptions).UpdateJob)], opts...)
	var resp *dataprocpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobControllerClient.UpdateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobControllerGRPCClient) CancelJob(ctx context.Context, req *dataprocpb.CancelJobRequest, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 900000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "job_id", url.QueryEscape(req.GetJobId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CancelJob[0:len((*c.CallOptions).CancelJob):len((*c.CallOptions).CancelJob)], opts...)
	var resp *dataprocpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobControllerClient.CancelJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobControllerGRPCClient) DeleteJob(ctx context.Context, req *dataprocpb.DeleteJobRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 900000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "job_id", url.QueryEscape(req.GetJobId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteJob[0:len((*c.CallOptions).DeleteJob):len((*c.CallOptions).DeleteJob)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.jobControllerClient.DeleteJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SubmitJobAsOperationOperation manages a long-running operation from SubmitJobAsOperation.
type SubmitJobAsOperationOperation struct {
	lro *longrunning.Operation
}

// SubmitJobAsOperationOperation returns a new SubmitJobAsOperationOperation from a given name.
// The name must be that of a previously created SubmitJobAsOperationOperation, possibly from a different process.
func (c *jobControllerGRPCClient) SubmitJobAsOperationOperation(name string) *SubmitJobAsOperationOperation {
	return &SubmitJobAsOperationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *SubmitJobAsOperationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	var resp dataprocpb.Job
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
func (op *SubmitJobAsOperationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Job, error) {
	var resp dataprocpb.Job
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
func (op *SubmitJobAsOperationOperation) Metadata() (*dataprocpb.JobMetadata, error) {
	var meta dataprocpb.JobMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *SubmitJobAsOperationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *SubmitJobAsOperationOperation) Name() string {
	return op.lro.Name()
}

// JobIterator manages a stream of *dataprocpb.Job.
type JobIterator struct {
	items    []*dataprocpb.Job
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataprocpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*dataprocpb.Job, error) {
	var item *dataprocpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
