// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for FirewallPolicy.
type FirewallPolicyServer struct{}

// ProtoToFirewallPolicy converts a FirewallPolicy resource from its proto representation.
func ProtoToFirewallPolicy(p *computepb.ComputeFirewallPolicy) *compute.FirewallPolicy {
	obj := &compute.FirewallPolicy{
		Name:              dcl.StringOrNil(p.Name),
		Id:                dcl.StringOrNil(p.Id),
		CreationTimestamp: dcl.StringOrNil(p.CreationTimestamp),
		Description:       dcl.StringOrNil(p.Description),
		Fingerprint:       dcl.StringOrNil(p.Fingerprint),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		SelfLinkWithId:    dcl.StringOrNil(p.SelfLinkWithId),
		RuleTupleCount:    dcl.Int64OrNil(p.RuleTupleCount),
		ShortName:         dcl.StringOrNil(p.ShortName),
		Parent:            dcl.StringOrNil(p.Parent),
	}
	return obj
}

// FirewallPolicyToProto converts a FirewallPolicy resource to its proto representation.
func FirewallPolicyToProto(resource *compute.FirewallPolicy) *computepb.ComputeFirewallPolicy {
	p := &computepb.ComputeFirewallPolicy{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Id:                dcl.ValueOrEmptyString(resource.Id),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Fingerprint:       dcl.ValueOrEmptyString(resource.Fingerprint),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		SelfLinkWithId:    dcl.ValueOrEmptyString(resource.SelfLinkWithId),
		RuleTupleCount:    dcl.ValueOrEmptyInt64(resource.RuleTupleCount),
		ShortName:         dcl.ValueOrEmptyString(resource.ShortName),
		Parent:            dcl.ValueOrEmptyString(resource.Parent),
	}

	return p
}

// ApplyFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicy Apply() method.
func (s *FirewallPolicyServer) applyFirewallPolicy(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeFirewallPolicyRequest) (*computepb.ComputeFirewallPolicy, error) {
	p := ProtoToFirewallPolicy(request.GetResource())
	res, err := c.ApplyFirewallPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallPolicyToProto(res)
	return r, nil
}

// ApplyFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicy Apply() method.
func (s *FirewallPolicyServer) ApplyComputeFirewallPolicy(ctx context.Context, request *computepb.ApplyComputeFirewallPolicyRequest) (*computepb.ComputeFirewallPolicy, error) {
	cl, err := createConfigFirewallPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFirewallPolicy(ctx, cl, request)
}

// DeleteFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicy Delete() method.
func (s *FirewallPolicyServer) DeleteComputeFirewallPolicy(ctx context.Context, request *computepb.DeleteComputeFirewallPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallPolicy(ctx, ProtoToFirewallPolicy(request.GetResource()))

}

// ListComputeFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicyList() method.
func (s *FirewallPolicyServer) ListComputeFirewallPolicy(ctx context.Context, request *computepb.ListComputeFirewallPolicyRequest) (*computepb.ListComputeFirewallPolicyResponse, error) {
	cl, err := createConfigFirewallPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallPolicy(ctx, ProtoToFirewallPolicy(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeFirewallPolicy
	for _, r := range resources.Items {
		rp := FirewallPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeFirewallPolicyResponse{Items: protos}, nil
}

func createConfigFirewallPolicy(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
