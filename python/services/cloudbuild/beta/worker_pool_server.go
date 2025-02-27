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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuild/beta/cloudbuild_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
)

// Server implements the gRPC interface for WorkerPool.
type WorkerPoolServer struct{}

// ProtoToWorkerPoolStateEnum converts a WorkerPoolStateEnum enum from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolStateEnum(e betapb.CloudbuildBetaWorkerPoolStateEnum) *beta.WorkerPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudbuildBetaWorkerPoolStateEnum_name[int32(e)]; ok {
		e := beta.WorkerPoolStateEnum(n[len("CloudbuildBetaWorkerPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkerPoolWorkerConfig converts a WorkerPoolWorkerConfig resource from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolWorkerConfig(p *betapb.CloudbuildBetaWorkerPoolWorkerConfig) *beta.WorkerPoolWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkerPoolWorkerConfig{
		MachineType:  dcl.StringOrNil(p.MachineType),
		DiskSizeGb:   dcl.Int64OrNil(p.DiskSizeGb),
		NoExternalIP: dcl.Bool(p.NoExternalIp),
	}
	return obj
}

// ProtoToWorkerPoolNetworkConfig converts a WorkerPoolNetworkConfig resource from its proto representation.
func ProtoToCloudbuildBetaWorkerPoolNetworkConfig(p *betapb.CloudbuildBetaWorkerPoolNetworkConfig) *beta.WorkerPoolNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkerPoolNetworkConfig{
		PeeredNetwork: dcl.StringOrNil(p.PeeredNetwork),
	}
	return obj
}

// ProtoToWorkerPool converts a WorkerPool resource from its proto representation.
func ProtoToWorkerPool(p *betapb.CloudbuildBetaWorkerPool) *beta.WorkerPool {
	obj := &beta.WorkerPool{
		Name:          dcl.StringOrNil(p.Name),
		State:         ProtoToCloudbuildBetaWorkerPoolStateEnum(p.GetState()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:    dcl.StringOrNil(p.GetDeleteTime()),
		WorkerConfig:  ProtoToCloudbuildBetaWorkerPoolWorkerConfig(p.GetWorkerConfig()),
		NetworkConfig: ProtoToCloudbuildBetaWorkerPoolNetworkConfig(p.GetNetworkConfig()),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	return obj
}

// WorkerPoolStateEnumToProto converts a WorkerPoolStateEnum enum to its proto representation.
func CloudbuildBetaWorkerPoolStateEnumToProto(e *beta.WorkerPoolStateEnum) betapb.CloudbuildBetaWorkerPoolStateEnum {
	if e == nil {
		return betapb.CloudbuildBetaWorkerPoolStateEnum(0)
	}
	if v, ok := betapb.CloudbuildBetaWorkerPoolStateEnum_value["WorkerPoolStateEnum"+string(*e)]; ok {
		return betapb.CloudbuildBetaWorkerPoolStateEnum(v)
	}
	return betapb.CloudbuildBetaWorkerPoolStateEnum(0)
}

// WorkerPoolWorkerConfigToProto converts a WorkerPoolWorkerConfig resource to its proto representation.
func CloudbuildBetaWorkerPoolWorkerConfigToProto(o *beta.WorkerPoolWorkerConfig) *betapb.CloudbuildBetaWorkerPoolWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaWorkerPoolWorkerConfig{
		MachineType:  dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:   dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		NoExternalIp: dcl.ValueOrEmptyBool(o.NoExternalIP),
	}
	return p
}

// WorkerPoolNetworkConfigToProto converts a WorkerPoolNetworkConfig resource to its proto representation.
func CloudbuildBetaWorkerPoolNetworkConfigToProto(o *beta.WorkerPoolNetworkConfig) *betapb.CloudbuildBetaWorkerPoolNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaWorkerPoolNetworkConfig{
		PeeredNetwork: dcl.ValueOrEmptyString(o.PeeredNetwork),
	}
	return p
}

// WorkerPoolToProto converts a WorkerPool resource to its proto representation.
func WorkerPoolToProto(resource *beta.WorkerPool) *betapb.CloudbuildBetaWorkerPool {
	p := &betapb.CloudbuildBetaWorkerPool{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		State:         CloudbuildBetaWorkerPoolStateEnumToProto(resource.State),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:    dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:    dcl.ValueOrEmptyString(resource.DeleteTime),
		WorkerConfig:  CloudbuildBetaWorkerPoolWorkerConfigToProto(resource.WorkerConfig),
		NetworkConfig: CloudbuildBetaWorkerPoolNetworkConfigToProto(resource.NetworkConfig),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) applyWorkerPool(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudbuildBetaWorkerPoolRequest) (*betapb.CloudbuildBetaWorkerPool, error) {
	p := ProtoToWorkerPool(request.GetResource())
	res, err := c.ApplyWorkerPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkerPoolToProto(res)
	return r, nil
}

// ApplyWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Apply() method.
func (s *WorkerPoolServer) ApplyCloudbuildBetaWorkerPool(ctx context.Context, request *betapb.ApplyCloudbuildBetaWorkerPoolRequest) (*betapb.CloudbuildBetaWorkerPool, error) {
	cl, err := createConfigWorkerPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyWorkerPool(ctx, cl, request)
}

// DeleteWorkerPool handles the gRPC request by passing it to the underlying WorkerPool Delete() method.
func (s *WorkerPoolServer) DeleteCloudbuildBetaWorkerPool(ctx context.Context, request *betapb.DeleteCloudbuildBetaWorkerPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkerPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkerPool(ctx, ProtoToWorkerPool(request.GetResource()))

}

// ListCloudbuildBetaWorkerPool handles the gRPC request by passing it to the underlying WorkerPoolList() method.
func (s *WorkerPoolServer) ListCloudbuildBetaWorkerPool(ctx context.Context, request *betapb.ListCloudbuildBetaWorkerPoolRequest) (*betapb.ListCloudbuildBetaWorkerPoolResponse, error) {
	cl, err := createConfigWorkerPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkerPool(ctx, ProtoToWorkerPool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudbuildBetaWorkerPool
	for _, r := range resources.Items {
		rp := WorkerPoolToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListCloudbuildBetaWorkerPoolResponse{Items: protos}, nil
}

func createConfigWorkerPool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
