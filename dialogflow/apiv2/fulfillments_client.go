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

package dialogflow

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newFulfillmentsClientHook clientHook

// FulfillmentsCallOptions contains the retry settings for each method of FulfillmentsClient.
type FulfillmentsCallOptions struct {
	GetFulfillment    []gax.CallOption
	UpdateFulfillment []gax.CallOption
}

func defaultFulfillmentsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFulfillmentsCallOptions() *FulfillmentsCallOptions {
	return &FulfillmentsCallOptions{
		GetFulfillment: []gax.CallOption{
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
		UpdateFulfillment: []gax.CallOption{
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

// internalFulfillmentsClient is an interface that defines the methods availaible from Dialogflow API.
type internalFulfillmentsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFulfillment(context.Context, *dialogflowpb.GetFulfillmentRequest, ...gax.CallOption) (*dialogflowpb.Fulfillment, error)
	UpdateFulfillment(context.Context, *dialogflowpb.UpdateFulfillmentRequest, ...gax.CallOption) (*dialogflowpb.Fulfillment, error)
}

// FulfillmentsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Fulfillments.
type FulfillmentsClient struct {
	// The internal transport-dependent client.
	internalClient internalFulfillmentsClient

	// The call options for this service.
	CallOptions *FulfillmentsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FulfillmentsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FulfillmentsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FulfillmentsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFulfillment retrieves the fulfillment.
func (c *FulfillmentsClient) GetFulfillment(ctx context.Context, req *dialogflowpb.GetFulfillmentRequest, opts ...gax.CallOption) (*dialogflowpb.Fulfillment, error) {
	return c.internalClient.GetFulfillment(ctx, req, opts...)
}

// UpdateFulfillment updates the fulfillment.
func (c *FulfillmentsClient) UpdateFulfillment(ctx context.Context, req *dialogflowpb.UpdateFulfillmentRequest, opts ...gax.CallOption) (*dialogflowpb.Fulfillment, error) {
	return c.internalClient.UpdateFulfillment(ctx, req, opts...)
}

// fulfillmentsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type fulfillmentsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FulfillmentsClient
	CallOptions **FulfillmentsCallOptions

	// The gRPC API client.
	fulfillmentsClient dialogflowpb.FulfillmentsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFulfillmentsClient creates a new fulfillments client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Fulfillments.
func NewFulfillmentsClient(ctx context.Context, opts ...option.ClientOption) (*FulfillmentsClient, error) {
	clientOpts := defaultFulfillmentsGRPCClientOptions()
	if newFulfillmentsClientHook != nil {
		hookOpts, err := newFulfillmentsClientHook(ctx, clientHookParams{})
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
	client := FulfillmentsClient{CallOptions: defaultFulfillmentsCallOptions()}

	c := &fulfillmentsGRPCClient{
		connPool:           connPool,
		disableDeadlines:   disableDeadlines,
		fulfillmentsClient: dialogflowpb.NewFulfillmentsClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *fulfillmentsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *fulfillmentsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *fulfillmentsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *fulfillmentsGRPCClient) GetFulfillment(ctx context.Context, req *dialogflowpb.GetFulfillmentRequest, opts ...gax.CallOption) (*dialogflowpb.Fulfillment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFulfillment[0:len((*c.CallOptions).GetFulfillment):len((*c.CallOptions).GetFulfillment)], opts...)
	var resp *dialogflowpb.Fulfillment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.fulfillmentsClient.GetFulfillment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *fulfillmentsGRPCClient) UpdateFulfillment(ctx context.Context, req *dialogflowpb.UpdateFulfillmentRequest, opts ...gax.CallOption) (*dialogflowpb.Fulfillment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "fulfillment.name", url.QueryEscape(req.GetFulfillment().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateFulfillment[0:len((*c.CallOptions).UpdateFulfillment):len((*c.CallOptions).UpdateFulfillment)], opts...)
	var resp *dialogflowpb.Fulfillment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.fulfillmentsClient.UpdateFulfillment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
