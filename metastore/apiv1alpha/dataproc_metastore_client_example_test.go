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

package metastore_test

import (
	"context"

	metastore "cloud.google.com/go/metastore/apiv1alpha"
	"google.golang.org/api/iterator"
	metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"
)

func ExampleNewDataprocMetastoreClient() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleDataprocMetastoreClient_ListServices() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.ListServicesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListServices(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDataprocMetastoreClient_GetService() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.GetServiceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_CreateService() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.CreateServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_UpdateService() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.UpdateServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_DeleteService() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.DeleteServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDataprocMetastoreClient_ListMetadataImports() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.ListMetadataImportsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListMetadataImports(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDataprocMetastoreClient_GetMetadataImport() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.GetMetadataImportRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetMetadataImport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_CreateMetadataImport() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.CreateMetadataImportRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateMetadataImport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_UpdateMetadataImport() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.UpdateMetadataImportRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateMetadataImport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
