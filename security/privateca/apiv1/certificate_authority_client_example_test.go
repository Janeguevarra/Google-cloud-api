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

package privateca_test

import (
	"context"

	privateca "cloud.google.com/go/security/privateca/apiv1"
	"google.golang.org/api/iterator"
	privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
)

func ExampleNewCertificateAuthorityClient() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleCertificateAuthorityClient_CreateCertificate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.CreateCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.GetCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificates() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.ListCertificatesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCertificates(ctx, req)
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

func ExampleCertificateAuthorityClient_RevokeCertificate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.RevokeCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RevokeCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_UpdateCertificate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.UpdateCertificateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ActivateCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.ActivateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ActivateCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_CreateCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.CreateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_DisableCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.DisableCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DisableCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_EnableCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.EnableCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.EnableCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_FetchCertificateAuthorityCsr() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.FetchCertificateAuthorityCsrRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.FetchCertificateAuthorityCsr(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.GetCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateAuthorities() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.ListCertificateAuthoritiesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCertificateAuthorities(ctx, req)
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

func ExampleCertificateAuthorityClient_UndeleteCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.UndeleteCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UndeleteCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_DeleteCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.DeleteCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCertificateAuthority() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.UpdateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_CreateCaPool() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.CreateCaPoolRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateCaPool(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCaPool() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.UpdateCaPoolRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateCaPool(ctx, req)
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

func ExampleCertificateAuthorityClient_GetCaPool() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.GetCaPoolRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCaPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCaPools() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.ListCaPoolsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCaPools(ctx, req)
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

func ExampleCertificateAuthorityClient_DeleteCaPool() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.DeleteCaPoolRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteCaPool(ctx, req)
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

func ExampleCertificateAuthorityClient_FetchCaCerts() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.FetchCaCertsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.FetchCaCerts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificateRevocationList() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.GetCertificateRevocationListRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificateRevocationList(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateRevocationLists() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.ListCertificateRevocationListsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCertificateRevocationLists(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCertificateRevocationList() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.UpdateCertificateRevocationListRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateCertificateRevocationList(ctx, req)
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

func ExampleCertificateAuthorityClient_CreateCertificateTemplate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.CreateCertificateTemplateRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateCertificateTemplate(ctx, req)
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

func ExampleCertificateAuthorityClient_DeleteCertificateTemplate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.DeleteCertificateTemplateRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteCertificateTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCertificateAuthorityClient_GetCertificateTemplate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.GetCertificateTemplateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCertificateTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateTemplates() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.ListCertificateTemplatesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCertificateTemplates(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCertificateTemplate() {
	// import privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"

	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &privatecapb.UpdateCertificateTemplateRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateCertificateTemplate(ctx, req)
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
