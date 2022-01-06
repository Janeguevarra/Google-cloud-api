// Copyright 2022 Google LLC
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

package dataflow_test

import (
	"context"

	dataflow "cloud.google.com/go/dataflow/apiv1beta3"
	dataflowpb "google.golang.org/genproto/googleapis/dataflow/v1beta3"
)

func ExampleNewTemplatesClient() {
	ctx := context.Background()
	c, err := dataflow.NewTemplatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleTemplatesClient_CreateJobFromTemplate() {
	ctx := context.Background()
	c, err := dataflow.NewTemplatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataflowpb.CreateJobFromTemplateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/dataflow/v1beta3#CreateJobFromTemplateRequest.
	}
	resp, err := c.CreateJobFromTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTemplatesClient_LaunchTemplate() {
	ctx := context.Background()
	c, err := dataflow.NewTemplatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataflowpb.LaunchTemplateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/dataflow/v1beta3#LaunchTemplateRequest.
	}
	resp, err := c.LaunchTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTemplatesClient_GetTemplate() {
	ctx := context.Background()
	c, err := dataflow.NewTemplatesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dataflowpb.GetTemplateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/dataflow/v1beta3#GetTemplateRequest.
	}
	resp, err := c.GetTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
