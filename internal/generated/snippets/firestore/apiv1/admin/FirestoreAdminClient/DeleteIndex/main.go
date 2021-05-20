// Copyright 2021 Google LLC
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

// [START firestore_v1_generated_FirestoreAdmin_DeleteIndex_sync]

package main

import (
	"context"

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	adminpb "google.golang.org/genproto/googleapis/firestore/admin/v1"
)

func main() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.DeleteIndexRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteIndex(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END firestore_v1_generated_FirestoreAdmin_DeleteIndex_sync]
