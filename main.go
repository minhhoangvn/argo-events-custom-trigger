package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ghodss/yaml"
	"github.com/minhhoangvn/argo-events-custom-trigger/triggers"
	"google.golang.org/grpc"
)

const (
	port = ":6969"
)

// server is used to implement helloworld.GreeterServer.
type customTrigger struct {
}

func (trigger *customTrigger) FetchResource(ctx context.Context, in *triggers.FetchResourceRequest) (*triggers.FetchResourceResponse, error) {
	fmt.Printf("Custom Trigger FetchResource")
	var resource map[string]string
	if err := yaml.Unmarshal(in.Resource, &resource); err != nil {
		return nil, err
	}

	fmt.Println("the resource to fetch is: \n", resource)

	return &triggers.FetchResourceResponse{
		Resource: nil,
	}, nil
}

func (trigger *customTrigger) Execute(ctx context.Context, in *triggers.ExecuteRequest) (*triggers.ExecuteResponse, error) {
	fmt.Printf("Custom Trigger Execute")
	var resource map[string]string
	if err := yaml.Unmarshal(in.Resource, &resource); err != nil {
		return nil, err
	}

	fmt.Printf("the resource to execute is: %s\n", string(in.Resource))
	return &triggers.ExecuteResponse{
		Response: []byte("success"),
	}, nil
}

func (trigger *customTrigger) ApplyPolicy(ctx context.Context, in *triggers.ApplyPolicyRequest) (*triggers.ApplyPolicyResponse, error) {
	fmt.Printf("Custom Trigger ApplyPolicy")
	return &triggers.ApplyPolicyResponse{
		Success: true,
		Message: "success",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	trigger := customTrigger{}
	triggers.RegisterTriggerService(grpcServer, &triggers.TriggerService{FetchResource: trigger.FetchResource, ApplyPolicy: trigger.ApplyPolicy, Execute: trigger.Execute})
	fmt.Printf("Starting trigger server at port 6969\n")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
