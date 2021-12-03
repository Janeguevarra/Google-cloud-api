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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newNodeGroupsClientHook clientHook

// NodeGroupsCallOptions contains the retry settings for each method of NodeGroupsClient.
type NodeGroupsCallOptions struct {
	AddNodes           []gax.CallOption
	AggregatedList     []gax.CallOption
	Delete             []gax.CallOption
	DeleteNodes        []gax.CallOption
	Get                []gax.CallOption
	GetIamPolicy       []gax.CallOption
	Insert             []gax.CallOption
	List               []gax.CallOption
	ListNodes          []gax.CallOption
	Patch              []gax.CallOption
	SetIamPolicy       []gax.CallOption
	SetNodeTemplate    []gax.CallOption
	TestIamPermissions []gax.CallOption
}

// internalNodeGroupsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalNodeGroupsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddNodes(context.Context, *computepb.AddNodesNodeGroupRequest, ...gax.CallOption) (*Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListNodeGroupsRequest, ...gax.CallOption) *NodeGroupsScopedListPairIterator
	Delete(context.Context, *computepb.DeleteNodeGroupRequest, ...gax.CallOption) (*Operation, error)
	DeleteNodes(context.Context, *computepb.DeleteNodesNodeGroupRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetNodeGroupRequest, ...gax.CallOption) (*computepb.NodeGroup, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyNodeGroupRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertNodeGroupRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListNodeGroupsRequest, ...gax.CallOption) *NodeGroupIterator
	ListNodes(context.Context, *computepb.ListNodesNodeGroupsRequest, ...gax.CallOption) *NodeGroupNodeIterator
	Patch(context.Context, *computepb.PatchNodeGroupRequest, ...gax.CallOption) (*Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyNodeGroupRequest, ...gax.CallOption) (*computepb.Policy, error)
	SetNodeTemplate(context.Context, *computepb.SetNodeTemplateNodeGroupRequest, ...gax.CallOption) (*Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsNodeGroupRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// NodeGroupsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The NodeGroups API.
type NodeGroupsClient struct {
	// The internal transport-dependent client.
	internalClient internalNodeGroupsClient

	// The call options for this service.
	CallOptions *NodeGroupsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *NodeGroupsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *NodeGroupsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *NodeGroupsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddNodes adds specified number of nodes to the node group.
func (c *NodeGroupsClient) AddNodes(ctx context.Context, req *computepb.AddNodesNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AddNodes(ctx, req, opts...)
}

// AggregatedList retrieves an aggregated list of node groups. Note: use nodeGroups.listNodes for more details about each group.
func (c *NodeGroupsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListNodeGroupsRequest, opts ...gax.CallOption) *NodeGroupsScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified NodeGroup resource.
func (c *NodeGroupsClient) Delete(ctx context.Context, req *computepb.DeleteNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DeleteNodes deletes specified nodes from the node group.
func (c *NodeGroupsClient) DeleteNodes(ctx context.Context, req *computepb.DeleteNodesNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.DeleteNodes(ctx, req, opts...)
}

// Get returns the specified NodeGroup. Get a list of available NodeGroups by making a list() request. Note: the “nodes” field should not be used. Use nodeGroups.listNodes instead.
func (c *NodeGroupsClient) Get(ctx context.Context, req *computepb.GetNodeGroupRequest, opts ...gax.CallOption) (*computepb.NodeGroup, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *NodeGroupsClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyNodeGroupRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert creates a NodeGroup resource in the specified project using the data included in the request.
func (c *NodeGroupsClient) Insert(ctx context.Context, req *computepb.InsertNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of node groups available to the specified project. Note: use nodeGroups.listNodes for more details about each group.
func (c *NodeGroupsClient) List(ctx context.Context, req *computepb.ListNodeGroupsRequest, opts ...gax.CallOption) *NodeGroupIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// ListNodes lists nodes in the node group.
func (c *NodeGroupsClient) ListNodes(ctx context.Context, req *computepb.ListNodesNodeGroupsRequest, opts ...gax.CallOption) *NodeGroupNodeIterator {
	return c.internalClient.ListNodes(ctx, req, opts...)
}

// Patch updates the specified node group.
func (c *NodeGroupsClient) Patch(ctx context.Context, req *computepb.PatchNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *NodeGroupsClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyNodeGroupRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// SetNodeTemplate updates the node template of the node group.
func (c *NodeGroupsClient) SetNodeTemplate(ctx context.Context, req *computepb.SetNodeTemplateNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetNodeTemplate(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *NodeGroupsClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsNodeGroupRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type nodeGroupsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewNodeGroupsRESTClient creates a new node groups rest client.
//
// The NodeGroups API.
func NewNodeGroupsRESTClient(ctx context.Context, opts ...option.ClientOption) (*NodeGroupsClient, error) {
	clientOpts := append(defaultNodeGroupsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &nodeGroupsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &NodeGroupsClient{internalClient: c, CallOptions: &NodeGroupsCallOptions{}}, nil
}

func defaultNodeGroupsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *nodeGroupsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *nodeGroupsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *nodeGroupsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddNodes adds specified number of nodes to the node group.
func (c *nodeGroupsRESTClient) AddNodes(ctx context.Context, req *computepb.AddNodesNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNodeGroupsAddNodesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v/addNodes", req.GetProject(), req.GetZone(), req.GetNodeGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// AggregatedList retrieves an aggregated list of node groups. Note: use nodeGroups.listNodes for more details about each group.
func (c *nodeGroupsRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListNodeGroupsRequest, opts ...gax.CallOption) *NodeGroupsScopedListPairIterator {
	it := &NodeGroupsScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListNodeGroupsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]NodeGroupsScopedListPair, string, error) {
		resp := &computepb.NodeGroupAggregatedList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/nodeGroups", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.IncludeAllScopes != nil {
			params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp

		elems := make([]NodeGroupsScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, NodeGroupsScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key })

		return elems, resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Delete deletes the specified NodeGroup resource.
func (c *nodeGroupsRESTClient) Delete(ctx context.Context, req *computepb.DeleteNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v", req.GetProject(), req.GetZone(), req.GetNodeGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// DeleteNodes deletes specified nodes from the node group.
func (c *nodeGroupsRESTClient) DeleteNodes(ctx context.Context, req *computepb.DeleteNodesNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNodeGroupsDeleteNodesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v/deleteNodes", req.GetProject(), req.GetZone(), req.GetNodeGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get returns the specified NodeGroup. Get a list of available NodeGroups by making a list() request. Note: the “nodes” field should not be used. Use nodeGroups.listNodes instead.
func (c *nodeGroupsRESTClient) Get(ctx context.Context, req *computepb.GetNodeGroupRequest, opts ...gax.CallOption) (*computepb.NodeGroup, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v", req.GetProject(), req.GetZone(), req.GetNodeGroup())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.NodeGroup{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *nodeGroupsRESTClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyNodeGroupRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v/getIamPolicy", req.GetProject(), req.GetZone(), req.GetResource())

	params := url.Values{}
	if req != nil && req.OptionsRequestedPolicyVersion != nil {
		params.Add("optionsRequestedPolicyVersion", fmt.Sprintf("%v", req.GetOptionsRequestedPolicyVersion()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Policy{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// Insert creates a NodeGroup resource in the specified project using the data included in the request.
func (c *nodeGroupsRESTClient) Insert(ctx context.Context, req *computepb.InsertNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNodeGroupResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups", req.GetProject(), req.GetZone())

	params := url.Values{}
	params.Add("initialNodeCount", fmt.Sprintf("%v", req.GetInitialNodeCount()))
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves a list of node groups available to the specified project. Note: use nodeGroups.listNodes for more details about each group.
func (c *nodeGroupsRESTClient) List(ctx context.Context, req *computepb.ListNodeGroupsRequest, opts ...gax.CallOption) *NodeGroupIterator {
	it := &NodeGroupIterator{}
	req = proto.Clone(req).(*computepb.ListNodeGroupsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.NodeGroup, string, error) {
		resp := &computepb.NodeGroupList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups", req.GetProject(), req.GetZone())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ListNodes lists nodes in the node group.
func (c *nodeGroupsRESTClient) ListNodes(ctx context.Context, req *computepb.ListNodesNodeGroupsRequest, opts ...gax.CallOption) *NodeGroupNodeIterator {
	it := &NodeGroupNodeIterator{}
	req = proto.Clone(req).(*computepb.ListNodesNodeGroupsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.NodeGroupNode, string, error) {
		resp := &computepb.NodeGroupsListNodes{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v/listNodes", req.GetProject(), req.GetZone(), req.GetNodeGroup())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch updates the specified node group.
func (c *nodeGroupsRESTClient) Patch(ctx context.Context, req *computepb.PatchNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNodeGroupResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v", req.GetProject(), req.GetZone(), req.GetNodeGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *nodeGroupsRESTClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyNodeGroupRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetZoneSetPolicyRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v/setIamPolicy", req.GetProject(), req.GetZone(), req.GetResource())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Policy{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// SetNodeTemplate updates the node template of the node group.
func (c *nodeGroupsRESTClient) SetNodeTemplate(ctx context.Context, req *computepb.SetNodeTemplateNodeGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNodeGroupsSetNodeTemplateRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v/setNodeTemplate", req.GetProject(), req.GetZone(), req.GetNodeGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *nodeGroupsRESTClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsNodeGroupRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTestPermissionsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/nodeGroups/%v/testIamPermissions", req.GetProject(), req.GetZone(), req.GetResource())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.TestPermissionsResponse{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// NodeGroupIterator manages a stream of *computepb.NodeGroup.
type NodeGroupIterator struct {
	items    []*computepb.NodeGroup
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.NodeGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NodeGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NodeGroupIterator) Next() (*computepb.NodeGroup, error) {
	var item *computepb.NodeGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NodeGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *NodeGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// NodeGroupNodeIterator manages a stream of *computepb.NodeGroupNode.
type NodeGroupNodeIterator struct {
	items    []*computepb.NodeGroupNode
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.NodeGroupNode, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NodeGroupNodeIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NodeGroupNodeIterator) Next() (*computepb.NodeGroupNode, error) {
	var item *computepb.NodeGroupNode
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NodeGroupNodeIterator) bufLen() int {
	return len(it.items)
}

func (it *NodeGroupNodeIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// NodeGroupsScopedListPair is a holder type for string/*computepb.NodeGroupsScopedList map entries
type NodeGroupsScopedListPair struct {
	Key   string
	Value *computepb.NodeGroupsScopedList
}

// NodeGroupsScopedListPairIterator manages a stream of NodeGroupsScopedListPair.
type NodeGroupsScopedListPairIterator struct {
	items    []NodeGroupsScopedListPair
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []NodeGroupsScopedListPair, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NodeGroupsScopedListPairIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NodeGroupsScopedListPairIterator) Next() (NodeGroupsScopedListPair, error) {
	var item NodeGroupsScopedListPair
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NodeGroupsScopedListPairIterator) bufLen() int {
	return len(it.items)
}

func (it *NodeGroupsScopedListPairIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
