// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by cloud.google.com/go/internal/gapicgen/gensnippets. DO NOT EDIT.

// [START analyticsadmin_v1alpha_generated_AnalyticsAdminService_ListAndroidAppDataStreams_sync]

package main

import (
	"context"

	admin "cloud.google.com/go/analytics/admin/apiv1alpha"
	"google.golang.org/api/iterator"
	adminpb "google.golang.org/genproto/googleapis/analytics/admin/v1alpha"
)

func main() {
	ctx := context.Background()
	c, err := admin.NewAnalyticsAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.ListAndroidAppDataStreamsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/analytics/admin/v1alpha#ListAndroidAppDataStreamsRequest.
	}
	it := c.ListAndroidAppDataStreams(ctx, req)
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

// [END analyticsadmin_v1alpha_generated_AnalyticsAdminService_ListAndroidAppDataStreams_sync]
