// Copyright 2015 Google Inc. All Rights Reserved.
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

package bigquery

import (
	"fmt"

	bq "google.golang.org/api/bigquery/v2"
)

type queryDestination interface {
	customizeQueryDst(conf *bq.JobConfigurationQuery, projectID string)
}

type querySource interface {
	customizeQuerySrc(conf *bq.JobConfigurationQuery, projectID string)
}

type queryOption interface {
	customizeQuery(conf *bq.JobConfigurationQuery, projectID string)
}

// UseQueryCache returns an Option that causes results to be fetched from the query cache if they are available.
// The query cache is a best-effort cache that is flushed whenever tables in the query are modified.
// Cached results are only available when TableID is unspecified in the query's destination Table.
// For more information, see https://cloud.google.com/bigquery/querying-data?#querycaching
func UseQueryCache() Option { return useQueryCache{} }

type useQueryCache struct{}

func (opt useQueryCache) implementsOption() {}

func (opt useQueryCache) customizeQuery(conf *bq.JobConfigurationQuery, projectID string) {
	conf.UseQueryCache = true
}

func query(job *bq.Job, dst Destination, src Source, projectID string, options ...Option) error {
	payload := &bq.JobConfigurationQuery{}

	d := dst.(queryDestination)
	s := src.(querySource)

	d.customizeQueryDst(payload, projectID)
	s.customizeQuerySrc(payload, projectID)

	for _, opt := range options {
		o, ok := opt.(queryOption)
		if !ok {
			return fmt.Errorf("option not applicable to dst/src pair: %#v", opt)
		}
		o.customizeQuery(payload, projectID)
	}

	job.Configuration = &bq.JobConfiguration{
		Query: payload,
	}
	return nil
}
