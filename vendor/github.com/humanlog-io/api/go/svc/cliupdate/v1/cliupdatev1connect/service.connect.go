// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/cliupdate/v1/service.proto

package cliupdatev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/humanlog-io/api/go/svc/cliupdate/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UpdateServiceName is the fully-qualified name of the UpdateService service.
	UpdateServiceName = "svc.cliupdate.v1.UpdateService"
)

// UpdateServiceClient is a client for the svc.cliupdate.v1.UpdateService service.
type UpdateServiceClient interface {
	GetNextUpdate(context.Context, *connect_go.Request[v1.GetNextUpdateRequest]) (*connect_go.Response[v1.GetNextUpdateResponse], error)
}

// NewUpdateServiceClient constructs a client for the svc.cliupdate.v1.UpdateService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUpdateServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UpdateServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &updateServiceClient{
		getNextUpdate: connect_go.NewClient[v1.GetNextUpdateRequest, v1.GetNextUpdateResponse](
			httpClient,
			baseURL+"/svc.cliupdate.v1.UpdateService/GetNextUpdate",
			opts...,
		),
	}
}

// updateServiceClient implements UpdateServiceClient.
type updateServiceClient struct {
	getNextUpdate *connect_go.Client[v1.GetNextUpdateRequest, v1.GetNextUpdateResponse]
}

// GetNextUpdate calls svc.cliupdate.v1.UpdateService.GetNextUpdate.
func (c *updateServiceClient) GetNextUpdate(ctx context.Context, req *connect_go.Request[v1.GetNextUpdateRequest]) (*connect_go.Response[v1.GetNextUpdateResponse], error) {
	return c.getNextUpdate.CallUnary(ctx, req)
}

// UpdateServiceHandler is an implementation of the svc.cliupdate.v1.UpdateService service.
type UpdateServiceHandler interface {
	GetNextUpdate(context.Context, *connect_go.Request[v1.GetNextUpdateRequest]) (*connect_go.Response[v1.GetNextUpdateResponse], error)
}

// NewUpdateServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUpdateServiceHandler(svc UpdateServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/svc.cliupdate.v1.UpdateService/GetNextUpdate", connect_go.NewUnaryHandler(
		"/svc.cliupdate.v1.UpdateService/GetNextUpdate",
		svc.GetNextUpdate,
		opts...,
	))
	return "/svc.cliupdate.v1.UpdateService/", mux
}

// UnimplementedUpdateServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUpdateServiceHandler struct{}

func (UnimplementedUpdateServiceHandler) GetNextUpdate(context.Context, *connect_go.Request[v1.GetNextUpdateRequest]) (*connect_go.Response[v1.GetNextUpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("svc.cliupdate.v1.UpdateService.GetNextUpdate is not implemented"))
}
