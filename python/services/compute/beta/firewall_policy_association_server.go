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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for FirewallPolicyAssociation.
type FirewallPolicyAssociationServer struct{}

// ProtoToFirewallPolicyAssociation converts a FirewallPolicyAssociation resource from its proto representation.
func ProtoToFirewallPolicyAssociation(p *betapb.ComputeBetaFirewallPolicyAssociation) *beta.FirewallPolicyAssociation {
	obj := &beta.FirewallPolicyAssociation{
		Name:             dcl.StringOrNil(p.Name),
		AttachmentTarget: dcl.StringOrNil(p.AttachmentTarget),
		FirewallPolicy:   dcl.StringOrNil(p.FirewallPolicy),
		ShortName:        dcl.StringOrNil(p.ShortName),
	}
	return obj
}

// FirewallPolicyAssociationToProto converts a FirewallPolicyAssociation resource to its proto representation.
func FirewallPolicyAssociationToProto(resource *beta.FirewallPolicyAssociation) *betapb.ComputeBetaFirewallPolicyAssociation {
	p := &betapb.ComputeBetaFirewallPolicyAssociation{
		Name:             dcl.ValueOrEmptyString(resource.Name),
		AttachmentTarget: dcl.ValueOrEmptyString(resource.AttachmentTarget),
		FirewallPolicy:   dcl.ValueOrEmptyString(resource.FirewallPolicy),
		ShortName:        dcl.ValueOrEmptyString(resource.ShortName),
	}

	return p
}

// ApplyFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociation Apply() method.
func (s *FirewallPolicyAssociationServer) applyFirewallPolicyAssociation(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaFirewallPolicyAssociationRequest) (*betapb.ComputeBetaFirewallPolicyAssociation, error) {
	p := ProtoToFirewallPolicyAssociation(request.GetResource())
	res, err := c.ApplyFirewallPolicyAssociation(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallPolicyAssociationToProto(res)
	return r, nil
}

// ApplyFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociation Apply() method.
func (s *FirewallPolicyAssociationServer) ApplyComputeBetaFirewallPolicyAssociation(ctx context.Context, request *betapb.ApplyComputeBetaFirewallPolicyAssociationRequest) (*betapb.ComputeBetaFirewallPolicyAssociation, error) {
	cl, err := createConfigFirewallPolicyAssociation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFirewallPolicyAssociation(ctx, cl, request)
}

// DeleteFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociation Delete() method.
func (s *FirewallPolicyAssociationServer) DeleteComputeBetaFirewallPolicyAssociation(ctx context.Context, request *betapb.DeleteComputeBetaFirewallPolicyAssociationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallPolicyAssociation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallPolicyAssociation(ctx, ProtoToFirewallPolicyAssociation(request.GetResource()))

}

// ListComputeBetaFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociationList() method.
func (s *FirewallPolicyAssociationServer) ListComputeBetaFirewallPolicyAssociation(ctx context.Context, request *betapb.ListComputeBetaFirewallPolicyAssociationRequest) (*betapb.ListComputeBetaFirewallPolicyAssociationResponse, error) {
	cl, err := createConfigFirewallPolicyAssociation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallPolicyAssociation(ctx, ProtoToFirewallPolicyAssociation(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaFirewallPolicyAssociation
	for _, r := range resources.Items {
		rp := FirewallPolicyAssociationToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaFirewallPolicyAssociationResponse{Items: protos}, nil
}

func createConfigFirewallPolicyAssociation(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
