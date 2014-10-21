// Copyright 2014 Google Inc. All Rights Reserved.
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

package storage_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/golang/oauth2/google"

	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

// TODO(jbd): Remove after Go 1.4.
// Related to https://codereview.appspot.com/107320046
func TestA(t *testing.T) {}

func Example_auth() {
	// Initialize an authorized transport with Google Developers Console
	// JSON key. Read the google package examples to learn more about
	// different authorization flows you can use.
	// http://godoc.org/github.com/golang/oauth2/google
	conf, err := google.NewServiceAccountJSONConfig(
		"/path/to/json/keyfile.json",
		storage.ScopeFullControl)
	if err != nil {
		log.Fatal(err)
	}

	ctx := cloud.NewContext("project-id", &http.Client{Transport: conf.NewTransport()})
	_ = ctx // Use the context (see other examples)
}

func Example_listObjects() {
	// see the auth example how to initiate a context.
	ctx := cloud.NewContext("project-id", &http.Client{Transport: nil})

	var query *storage.Query
	for {
		// If you are using this package on App Engine Managed VMs runtime,
		// you can init a bucket client with your app's default bucket name.
		// See http://godoc.org/google.golang.org/appengine/file#DefaultBucketName.
		objects, err := storage.List(ctx, "bucketname", query)
		if err != nil {
			log.Fatal(err)
		}
		for _, obj := range objects.Results {
			log.Printf("object name: %s, size: %v", obj.Name, obj.Size)
		}
		// if there are more results, objects.Next
		// will be non-nil.
		query = objects.Next
		if query == nil {
			break
		}
	}

	log.Println("paginated through all object items in the bucket you specified.")
}

func Example_readObjects() {
	// see the auth example how to initiate a context.
	ctx := cloud.NewContext("project-id", &http.Client{Transport: nil})

	rc, err := storage.NewReader(ctx, "bucketname", "filename1")
	if err != nil {
		log.Fatal(err)
	}
	slurp, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("file contents:", slurp)
}

func Example_writeObjects() {
	// see the auth example how to initiate a context.
	ctx := cloud.NewContext("project-id", &http.Client{Transport: nil})

	wc := storage.NewWriter(ctx, "bucketname", "filename1", &storage.Object{
		ContentType: "text/plain",
	})
	if _, err := wc.Write([]byte("hello world")); err != nil {
		log.Fatal(err)
	}
	if err := wc.Close(); err != nil {
		log.Fatal(err)
	}

	o, err := wc.Object()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("updated object:", o)
}

func Example_touchObjects() {
	// see the auth example how to initiate a context.
	ctx := cloud.NewContext("project-id", &http.Client{Transport: nil})

	o, err := storage.Put(ctx, "bucketname", "filename", &storage.Object{
		ContentType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("touched new file:", o)
}

func Example_copyObjects() {
	// see the auth example how to initiate a context.
	ctx := cloud.NewContext("project-id", &http.Client{Transport: nil})

	o, err := storage.Copy(ctx, "bucketname", "file1", &storage.Object{
		Name:   "file2",
		Bucket: "yet-another-bucketname",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("copied file:", o)
}
