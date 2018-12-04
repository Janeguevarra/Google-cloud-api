// Copyright 2018 Google LLC
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

// AUTO-GENERATED CODE. DO NOT EDIT.

package storage

import (
	"context"
	"fmt"
	"time"

	gax "github.com/googleapis/gax-go"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// BigQueryStorageCallOptions contains the retry settings for each method of BigQueryStorageClient.
type BigQueryStorageCallOptions struct {
	CreateReadSession             []gax.CallOption
	ReadRows                      []gax.CallOption
	BatchCreateReadSessionStreams []gax.CallOption
	FinalizeStream                []gax.CallOption
	SplitReadStream               []gax.CallOption
}

func defaultBigQueryStorageClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("bigquerystorage.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultBigQueryStorageCallOptions() *BigQueryStorageCallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &BigQueryStorageCallOptions{
		CreateReadSession:             retry[[2]string{"default", "idempotent"}],
		ReadRows:                      retry[[2]string{"default", "idempotent"}],
		BatchCreateReadSessionStreams: retry[[2]string{"default", "idempotent"}],
		FinalizeStream:                retry[[2]string{"default", "idempotent"}],
		SplitReadStream:               retry[[2]string{"default", "idempotent"}],
	}
}

// BigQueryStorageClient is a client for interacting with BigQuery Storage API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type BigQueryStorageClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	bigQueryStorageClient storagepb.BigQueryStorageClient

	// The call options for this service.
	CallOptions *BigQueryStorageCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBigQueryStorageClient creates a new big query storage client.
//
// BigQuery storage API.
//
// The BigQuery storage API can be used to read data stored in BigQuery.
func NewBigQueryStorageClient(ctx context.Context, opts ...option.ClientOption) (*BigQueryStorageClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultBigQueryStorageClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &BigQueryStorageClient{
		conn:        conn,
		CallOptions: defaultBigQueryStorageCallOptions(),

		bigQueryStorageClient: storagepb.NewBigQueryStorageClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *BigQueryStorageClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BigQueryStorageClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BigQueryStorageClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateReadSession creates a new read session. A read session divides the contents of a
// BigQuery table into one or more streams, which can then be used to read
// data from the table. The read session also specifies properties of the
// data to be read, such as a list of columns or a push-down filter describing
// the rows to be returned.
//
// A particular row can be read by at most one stream. When the caller has
// reached the end of each stream in the session, then all the data in the
// table has been read.
//
// Read sessions automatically expire 24 hours after they are created and do
// not require manual clean-up by the caller.
func (c *BigQueryStorageClient) CreateReadSession(ctx context.Context, req *storagepb.CreateReadSessionRequest, opts ...gax.CallOption) (*storagepb.ReadSession, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "table_reference.project_id", req.GetTableReference().GetProjectId(), "table_reference.dataset_id", req.GetTableReference().GetDatasetId()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateReadSession[0:len(c.CallOptions.CreateReadSession):len(c.CallOptions.CreateReadSession)], opts...)
	var resp *storagepb.ReadSession
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.CreateReadSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReadRows reads rows from the table in the format prescribed by the read session.
// Each response contains one or more table rows, up to a maximum of 10 MiB
// per response; read requests which attempt to read individual rows larger
// than this will fail.
//
// Each request also returns a set of stream statistics reflecting the
// estimated total number of rows in the read stream. This number is computed
// based on the total table size and the number of active streams in the read
// session, and may change as other streams continue to read data.
func (c *BigQueryStorageClient) ReadRows(ctx context.Context, req *storagepb.ReadRowsRequest, opts ...gax.CallOption) (storagepb.BigQueryStorage_ReadRowsClient, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "read_position.stream.name", req.GetReadPosition().GetStream().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ReadRows[0:len(c.CallOptions.ReadRows):len(c.CallOptions.ReadRows)], opts...)
	var resp storagepb.BigQueryStorage_ReadRowsClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.ReadRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchCreateReadSessionStreams creates additional streams for a ReadSession. This API can be used to
// dynamically adjust the parallelism of a batch processing task upwards by
// adding additional workers.
func (c *BigQueryStorageClient) BatchCreateReadSessionStreams(ctx context.Context, req *storagepb.BatchCreateReadSessionStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCreateReadSessionStreamsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session.name", req.GetSession().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchCreateReadSessionStreams[0:len(c.CallOptions.BatchCreateReadSessionStreams):len(c.CallOptions.BatchCreateReadSessionStreams)], opts...)
	var resp *storagepb.BatchCreateReadSessionStreamsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.BatchCreateReadSessionStreams(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FinalizeStream triggers the graceful termination of a single stream in a ReadSession. This
// API can be used to dynamically adjust the parallelism of a batch processing
// task downwards without losing data.
//
// This API does not delete the stream -- it remains visible in the
// ReadSession, and any data processed by the stream is not released to other
// streams. However, no additional data will be assigned to the stream once
// this call completes. Callers must continue reading data on the stream until
// the end of the stream is reached so that data which has already been
// assigned to the stream will be processed.
//
// This method will return an error if there are no other live streams
// in the Session, or if SplitReadStream() has been called on the given
// Stream.
func (c *BigQueryStorageClient) FinalizeStream(ctx context.Context, req *storagepb.FinalizeStreamRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "stream.name", req.GetStream().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.FinalizeStream[0:len(c.CallOptions.FinalizeStream):len(c.CallOptions.FinalizeStream)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.bigQueryStorageClient.FinalizeStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SplitReadStream splits a given read stream into two Streams. These streams are referred to
// as the primary and the residual of the split. The original stream can still
// be read from in the same manner as before. Both of the returned streams can
// also be read from, and the total rows return by both child streams will be
// the same as the rows read from the original stream.
//
// Moreover, the two child streams will be allocated back to back in the
// original Stream. Concretely, it is guaranteed that for streams Original,
// Primary, and Residual, that Original[0-j] = Primary[0-j] and
// Original[j-n] = Residual[0-m] once the streams have been read to
// completion.
//
// This method is guaranteed to be idempotent.
func (c *BigQueryStorageClient) SplitReadStream(ctx context.Context, req *storagepb.SplitReadStreamRequest, opts ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "original_stream.name", req.GetOriginalStream().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SplitReadStream[0:len(c.CallOptions.SplitReadStream):len(c.CallOptions.SplitReadStream)], opts...)
	var resp *storagepb.SplitReadStreamResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.SplitReadStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
