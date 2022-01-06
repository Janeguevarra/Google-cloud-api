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

package accessapproval_test

import (
	"context"

	accessapproval "cloud.google.com/go/accessapproval/apiv1"
	"google.golang.org/api/iterator"
	accessapprovalpb "google.golang.org/genproto/googleapis/cloud/accessapproval/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_ListApprovalRequests() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.ListApprovalRequestsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/accessapproval/v1#ListApprovalRequestsMessage.
	}
	it := c.ListApprovalRequests(ctx, req)
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

func ExampleClient_GetApprovalRequest() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.GetApprovalRequestMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/accessapproval/v1#GetApprovalRequestMessage.
	}
	resp, err := c.GetApprovalRequest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ApproveApprovalRequest() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.ApproveApprovalRequestMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/accessapproval/v1#ApproveApprovalRequestMessage.
	}
	resp, err := c.ApproveApprovalRequest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DismissApprovalRequest() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.DismissApprovalRequestMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/accessapproval/v1#DismissApprovalRequestMessage.
	}
	resp, err := c.DismissApprovalRequest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetAccessApprovalSettings() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.GetAccessApprovalSettingsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/accessapproval/v1#GetAccessApprovalSettingsMessage.
	}
	resp, err := c.GetAccessApprovalSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateAccessApprovalSettings() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.UpdateAccessApprovalSettingsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/accessapproval/v1#UpdateAccessApprovalSettingsMessage.
	}
	resp, err := c.UpdateAccessApprovalSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteAccessApprovalSettings() {
	ctx := context.Background()
	c, err := accessapproval.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &accessapprovalpb.DeleteAccessApprovalSettingsMessage{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/accessapproval/v1#DeleteAccessApprovalSettingsMessage.
	}
	err = c.DeleteAccessApprovalSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
