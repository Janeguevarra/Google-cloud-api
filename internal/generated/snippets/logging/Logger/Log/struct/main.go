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

// [START logging_generated_logging_Logger_Log_struct]

package main

import (
	"context"

	"cloud.google.com/go/logging"
)

func main() {
	type MyEntry struct {
		Name  string
		Count int
	}

	ctx := context.Background()
	client, err := logging.NewClient(ctx, "my-project")
	if err != nil {
		// TODO: Handle error.
	}
	lg := client.Logger("my-log")
	lg.Log(logging.Entry{Payload: MyEntry{Name: "Bob", Count: 3}})
}

// [END logging_generated_logging_Logger_Log_struct]
