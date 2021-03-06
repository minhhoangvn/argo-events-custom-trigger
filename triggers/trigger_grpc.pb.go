// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package triggers

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TriggerClient is the client API for Trigger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TriggerClient interface {
	// FetchResource fetches the resource to be triggered.
	FetchResource(ctx context.Context, in *FetchResourceRequest, opts ...grpc.CallOption) (*FetchResourceResponse, error)
	// Execute executes the requested trigger resource.
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
	// ApplyPolicy applies policies on the trigger execution result.
	ApplyPolicy(ctx context.Context, in *ApplyPolicyRequest, opts ...grpc.CallOption) (*ApplyPolicyResponse, error)
}

type triggerClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerClient(cc grpc.ClientConnInterface) TriggerClient {
	return &triggerClient{cc}
}

var triggerFetchResourceStreamDesc = &grpc.StreamDesc{
	StreamName: "FetchResource",
}

func (c *triggerClient) FetchResource(ctx context.Context, in *FetchResourceRequest, opts ...grpc.CallOption) (*FetchResourceResponse, error) {
	out := new(FetchResourceResponse)
	err := c.cc.Invoke(ctx, "/triggers.Trigger/FetchResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var triggerExecuteStreamDesc = &grpc.StreamDesc{
	StreamName: "Execute",
}

func (c *triggerClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/triggers.Trigger/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var triggerApplyPolicyStreamDesc = &grpc.StreamDesc{
	StreamName: "ApplyPolicy",
}

func (c *triggerClient) ApplyPolicy(ctx context.Context, in *ApplyPolicyRequest, opts ...grpc.CallOption) (*ApplyPolicyResponse, error) {
	out := new(ApplyPolicyResponse)
	err := c.cc.Invoke(ctx, "/triggers.Trigger/ApplyPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TriggerService is the service API for Trigger service.
// Fields should be assigned to their respective handler implementations only before
// RegisterTriggerService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type TriggerService struct {
	// FetchResource fetches the resource to be triggered.
	FetchResource func(context.Context, *FetchResourceRequest) (*FetchResourceResponse, error)
	// Execute executes the requested trigger resource.
	Execute func(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	// ApplyPolicy applies policies on the trigger execution result.
	ApplyPolicy func(context.Context, *ApplyPolicyRequest) (*ApplyPolicyResponse, error)
}

func (s *TriggerService) fetchResource(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.FetchResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/triggers.Trigger/FetchResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.FetchResource(ctx, req.(*FetchResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *TriggerService) execute(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/triggers.Trigger/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *TriggerService) applyPolicy(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ApplyPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/triggers.Trigger/ApplyPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ApplyPolicy(ctx, req.(*ApplyPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterTriggerService registers a service implementation with a gRPC server.
func RegisterTriggerService(s grpc.ServiceRegistrar, srv *TriggerService) {
	srvCopy := *srv
	if srvCopy.FetchResource == nil {
		srvCopy.FetchResource = func(context.Context, *FetchResourceRequest) (*FetchResourceResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method FetchResource not implemented")
		}
	}
	if srvCopy.Execute == nil {
		srvCopy.Execute = func(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
		}
	}
	if srvCopy.ApplyPolicy == nil {
		srvCopy.ApplyPolicy = func(context.Context, *ApplyPolicyRequest) (*ApplyPolicyResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ApplyPolicy not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "triggers.Trigger",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "FetchResource",
				Handler:    srvCopy.fetchResource,
			},
			{
				MethodName: "Execute",
				Handler:    srvCopy.execute,
			},
			{
				MethodName: "ApplyPolicy",
				Handler:    srvCopy.applyPolicy,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "protoc/trigger.proto",
	}

	s.RegisterService(&sd, nil)
}
