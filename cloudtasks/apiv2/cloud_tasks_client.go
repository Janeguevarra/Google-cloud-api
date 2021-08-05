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

package cloudtasks

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
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListQueues         []gax.CallOption
	GetQueue           []gax.CallOption
	CreateQueue        []gax.CallOption
	UpdateQueue        []gax.CallOption
	DeleteQueue        []gax.CallOption
	PurgeQueue         []gax.CallOption
	PauseQueue         []gax.CallOption
	ResumeQueue        []gax.CallOption
	GetIamPolicy       []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
	ListTasks          []gax.CallOption
	GetTask            []gax.CallOption
	CreateTask         []gax.CallOption
	DeleteTask         []gax.CallOption
	RunTask            []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudtasks.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudtasks.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudtasks.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListQueues: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetQueue: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateQueue: []gax.CallOption{},
		UpdateQueue: []gax.CallOption{},
		DeleteQueue: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		PurgeQueue:  []gax.CallOption{},
		PauseQueue:  []gax.CallOption{},
		ResumeQueue: []gax.CallOption{},
		GetIamPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SetIamPolicy: []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListTasks: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetTask: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateTask: []gax.CallOption{},
		DeleteTask: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RunTask: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from Cloud Tasks API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListQueues(context.Context, *taskspb.ListQueuesRequest, ...gax.CallOption) *QueueIterator
	GetQueue(context.Context, *taskspb.GetQueueRequest, ...gax.CallOption) (*taskspb.Queue, error)
	CreateQueue(context.Context, *taskspb.CreateQueueRequest, ...gax.CallOption) (*taskspb.Queue, error)
	UpdateQueue(context.Context, *taskspb.UpdateQueueRequest, ...gax.CallOption) (*taskspb.Queue, error)
	DeleteQueue(context.Context, *taskspb.DeleteQueueRequest, ...gax.CallOption) error
	PurgeQueue(context.Context, *taskspb.PurgeQueueRequest, ...gax.CallOption) (*taskspb.Queue, error)
	PauseQueue(context.Context, *taskspb.PauseQueueRequest, ...gax.CallOption) (*taskspb.Queue, error)
	ResumeQueue(context.Context, *taskspb.ResumeQueueRequest, ...gax.CallOption) (*taskspb.Queue, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	ListTasks(context.Context, *taskspb.ListTasksRequest, ...gax.CallOption) *TaskIterator
	GetTask(context.Context, *taskspb.GetTaskRequest, ...gax.CallOption) (*taskspb.Task, error)
	CreateTask(context.Context, *taskspb.CreateTaskRequest, ...gax.CallOption) (*taskspb.Task, error)
	DeleteTask(context.Context, *taskspb.DeleteTaskRequest, ...gax.CallOption) error
	RunTask(context.Context, *taskspb.RunTaskRequest, ...gax.CallOption) (*taskspb.Task, error)
}

// Client is a client for interacting with Cloud Tasks API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Cloud Tasks allows developers to manage the execution of background
// work in their applications.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
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

// ListQueues lists queues.
//
// Queues are returned in lexicographical order.
func (c *Client) ListQueues(ctx context.Context, req *taskspb.ListQueuesRequest, opts ...gax.CallOption) *QueueIterator {
	return c.internalClient.ListQueues(ctx, req, opts...)
}

// GetQueue gets a queue.
func (c *Client) GetQueue(ctx context.Context, req *taskspb.GetQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	return c.internalClient.GetQueue(ctx, req, opts...)
}

// CreateQueue creates a queue.
//
// Queues created with this method allow tasks to live for a maximum of 31
// days. After a task is 31 days old, the task will be deleted regardless of whether
// it was dispatched or not.
//
// WARNING: Using this method may have unintended side effects if you are
// using an App Engine queue.yaml or queue.xml file to manage your queues.
// Read
// Overview of Queue Management and
// queue.yaml (at https://cloud.google.com/tasks/docs/queue-yaml) before using
// this method.
func (c *Client) CreateQueue(ctx context.Context, req *taskspb.CreateQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	return c.internalClient.CreateQueue(ctx, req, opts...)
}

// UpdateQueue updates a queue.
//
// This method creates the queue if it does not exist and updates
// the queue if it does exist.
//
// Queues created with this method allow tasks to live for a maximum of 31
// days. After a task is 31 days old, the task will be deleted regardless of whether
// it was dispatched or not.
//
// WARNING: Using this method may have unintended side effects if you are
// using an App Engine queue.yaml or queue.xml file to manage your queues.
// Read
// Overview of Queue Management and
// queue.yaml (at https://cloud.google.com/tasks/docs/queue-yaml) before using
// this method.
func (c *Client) UpdateQueue(ctx context.Context, req *taskspb.UpdateQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	return c.internalClient.UpdateQueue(ctx, req, opts...)
}

// DeleteQueue deletes a queue.
//
// This command will delete the queue even if it has tasks in it.
//
// Note: If you delete a queue, a queue with the same name can’t be created
// for 7 days.
//
// WARNING: Using this method may have unintended side effects if you are
// using an App Engine queue.yaml or queue.xml file to manage your queues.
// Read
// Overview of Queue Management and
// queue.yaml (at https://cloud.google.com/tasks/docs/queue-yaml) before using
// this method.
func (c *Client) DeleteQueue(ctx context.Context, req *taskspb.DeleteQueueRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteQueue(ctx, req, opts...)
}

// PurgeQueue purges a queue by deleting all of its tasks.
//
// All tasks created before this method is called are permanently deleted.
//
// Purge operations can take up to one minute to take effect. Tasks
// might be dispatched before the purge takes effect. A purge is irreversible.
func (c *Client) PurgeQueue(ctx context.Context, req *taskspb.PurgeQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	return c.internalClient.PurgeQueue(ctx, req, opts...)
}

// PauseQueue pauses the queue.
//
// If a queue is paused then the system will stop dispatching tasks
// until the queue is resumed via
// ResumeQueue. Tasks can still be added
// when the queue is paused. A queue is paused if its
// state is PAUSED.
func (c *Client) PauseQueue(ctx context.Context, req *taskspb.PauseQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	return c.internalClient.PauseQueue(ctx, req, opts...)
}

// ResumeQueue resume a queue.
//
// This method resumes a queue after it has been
// PAUSED or
// DISABLED. The state of a queue is stored
// in the queue’s state; after calling this method it
// will be set to RUNNING.
//
// WARNING: Resuming many high-QPS queues at the same time can
// lead to target overloading. If you are resuming high-QPS
// queues, follow the 500/50/5 pattern described in
// Managing Cloud Tasks Scaling
// Risks (at https://cloud.google.com/tasks/docs/manage-cloud-task-scaling).
func (c *Client) ResumeQueue(ctx context.Context, req *taskspb.ResumeQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	return c.internalClient.ResumeQueue(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a Queue.
// Returns an empty policy if the resource exists and does not have a policy
// set.
//
// Authorization requires the following
// Google IAM (at https://cloud.google.com/iam) permission on the specified
// resource parent:
//
//   cloudtasks.queues.getIamPolicy
func (c *Client) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy for a Queue. Replaces any existing
// policy.
//
// Note: The Cloud Console does not check queue-level IAM permissions yet.
// Project-level permissions are required to use the Cloud Console.
//
// Authorization requires the following
// Google IAM (at https://cloud.google.com/iam) permission on the specified
// resource parent:
//
//   cloudtasks.queues.setIamPolicy
func (c *Client) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on a Queue.
// If the resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building permission-aware
// UIs and command-line tools, not for authorization checking. This operation
// may “fail open” without warning.
func (c *Client) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// ListTasks lists the tasks in a queue.
//
// By default, only the BASIC view is retrieved
// due to performance considerations;
// response_view controls the
// subset of information which is returned.
//
// The tasks may be returned in any order. The ordering may change at any
// time.
func (c *Client) ListTasks(ctx context.Context, req *taskspb.ListTasksRequest, opts ...gax.CallOption) *TaskIterator {
	return c.internalClient.ListTasks(ctx, req, opts...)
}

// GetTask gets a task.
func (c *Client) GetTask(ctx context.Context, req *taskspb.GetTaskRequest, opts ...gax.CallOption) (*taskspb.Task, error) {
	return c.internalClient.GetTask(ctx, req, opts...)
}

// CreateTask creates a task and adds it to a queue.
//
// Tasks cannot be updated after creation; there is no UpdateTask command.
//
//   The maximum task size is 100KB.
func (c *Client) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest, opts ...gax.CallOption) (*taskspb.Task, error) {
	return c.internalClient.CreateTask(ctx, req, opts...)
}

// DeleteTask deletes a task.
//
// A task can be deleted if it is scheduled or dispatched. A task
// cannot be deleted if it has executed successfully or permanently
// failed.
func (c *Client) DeleteTask(ctx context.Context, req *taskspb.DeleteTaskRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteTask(ctx, req, opts...)
}

// RunTask forces a task to run now.
//
// When this method is called, Cloud Tasks will dispatch the task, even if
// the task is already running, the queue has reached its RateLimits or
// is PAUSED.
//
// This command is meant to be used for manual debugging. For
// example, RunTask can be used to retry a failed
// task after a fix has been made or to manually force a task to be
// dispatched now.
//
// The dispatched task is returned. That is, the task that is returned
// contains the status after the task is dispatched but
// before the task is received by its target.
//
// If Cloud Tasks receives a successful response from the task’s
// target, then the task will be deleted; otherwise the task’s
// schedule_time will be reset to the time that
// RunTask was called plus the retry delay specified
// in the queue’s RetryConfig.
//
// RunTask returns
// NOT_FOUND when it is called on a
// task that has already succeeded or permanently failed.
func (c *Client) RunTask(ctx context.Context, req *taskspb.RunTaskRequest, opts ...gax.CallOption) (*taskspb.Task, error) {
	return c.internalClient.RunTask(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Cloud Tasks API over gRPC transport.
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
	client taskspb.CloudTasksClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new cloud tasks client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Cloud Tasks allows developers to manage the execution of background
// work in their applications.
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
		client:           taskspb.NewCloudTasksClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

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

func (c *gRPCClient) ListQueues(ctx context.Context, req *taskspb.ListQueuesRequest, opts ...gax.CallOption) *QueueIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListQueues[0:len((*c.CallOptions).ListQueues):len((*c.CallOptions).ListQueues)], opts...)
	it := &QueueIterator{}
	req = proto.Clone(req).(*taskspb.ListQueuesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*taskspb.Queue, string, error) {
		resp := &taskspb.ListQueuesResponse{}
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
			resp, err = c.client.ListQueues(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetQueues(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetQueue(ctx context.Context, req *taskspb.GetQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetQueue[0:len((*c.CallOptions).GetQueue):len((*c.CallOptions).GetQueue)], opts...)
	var resp *taskspb.Queue
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetQueue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateQueue(ctx context.Context, req *taskspb.CreateQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateQueue[0:len((*c.CallOptions).CreateQueue):len((*c.CallOptions).CreateQueue)], opts...)
	var resp *taskspb.Queue
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateQueue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateQueue(ctx context.Context, req *taskspb.UpdateQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "queue.name", url.QueryEscape(req.GetQueue().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateQueue[0:len((*c.CallOptions).UpdateQueue):len((*c.CallOptions).UpdateQueue)], opts...)
	var resp *taskspb.Queue
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateQueue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteQueue(ctx context.Context, req *taskspb.DeleteQueueRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteQueue[0:len((*c.CallOptions).DeleteQueue):len((*c.CallOptions).DeleteQueue)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteQueue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) PurgeQueue(ctx context.Context, req *taskspb.PurgeQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PurgeQueue[0:len((*c.CallOptions).PurgeQueue):len((*c.CallOptions).PurgeQueue)], opts...)
	var resp *taskspb.Queue
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.PurgeQueue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) PauseQueue(ctx context.Context, req *taskspb.PauseQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PauseQueue[0:len((*c.CallOptions).PauseQueue):len((*c.CallOptions).PauseQueue)], opts...)
	var resp *taskspb.Queue
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.PauseQueue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ResumeQueue(ctx context.Context, req *taskspb.ResumeQueueRequest, opts ...gax.CallOption) (*taskspb.Queue, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ResumeQueue[0:len((*c.CallOptions).ResumeQueue):len((*c.CallOptions).ResumeQueue)], opts...)
	var resp *taskspb.Queue
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ResumeQueue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListTasks(ctx context.Context, req *taskspb.ListTasksRequest, opts ...gax.CallOption) *TaskIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTasks[0:len((*c.CallOptions).ListTasks):len((*c.CallOptions).ListTasks)], opts...)
	it := &TaskIterator{}
	req = proto.Clone(req).(*taskspb.ListTasksRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*taskspb.Task, string, error) {
		resp := &taskspb.ListTasksResponse{}
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
			resp, err = c.client.ListTasks(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTasks(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetTask(ctx context.Context, req *taskspb.GetTaskRequest, opts ...gax.CallOption) (*taskspb.Task, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTask[0:len((*c.CallOptions).GetTask):len((*c.CallOptions).GetTask)], opts...)
	var resp *taskspb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetTask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest, opts ...gax.CallOption) (*taskspb.Task, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateTask[0:len((*c.CallOptions).CreateTask):len((*c.CallOptions).CreateTask)], opts...)
	var resp *taskspb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateTask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteTask(ctx context.Context, req *taskspb.DeleteTaskRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTask[0:len((*c.CallOptions).DeleteTask):len((*c.CallOptions).DeleteTask)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteTask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) RunTask(ctx context.Context, req *taskspb.RunTaskRequest, opts ...gax.CallOption) (*taskspb.Task, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RunTask[0:len((*c.CallOptions).RunTask):len((*c.CallOptions).RunTask)], opts...)
	var resp *taskspb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.RunTask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueueIterator manages a stream of *taskspb.Queue.
type QueueIterator struct {
	items    []*taskspb.Queue
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
	InternalFetch func(pageSize int, pageToken string) (results []*taskspb.Queue, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *QueueIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *QueueIterator) Next() (*taskspb.Queue, error) {
	var item *taskspb.Queue
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *QueueIterator) bufLen() int {
	return len(it.items)
}

func (it *QueueIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TaskIterator manages a stream of *taskspb.Task.
type TaskIterator struct {
	items    []*taskspb.Task
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
	InternalFetch func(pageSize int, pageToken string) (results []*taskspb.Task, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TaskIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TaskIterator) Next() (*taskspb.Task, error) {
	var item *taskspb.Task
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TaskIterator) bufLen() int {
	return len(it.items)
}

func (it *TaskIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
