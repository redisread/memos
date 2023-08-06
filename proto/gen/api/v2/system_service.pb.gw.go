// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v2/system_service.proto

/*
Package apiv2 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package apiv2

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_SystemService_GetSystemInfo_0(ctx context.Context, marshaler runtime.Marshaler, client SystemServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSystemInfoRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetSystemInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SystemService_GetSystemInfo_0(ctx context.Context, marshaler runtime.Marshaler, server SystemServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSystemInfoRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetSystemInfo(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterSystemServiceHandlerServer registers the http handlers for service SystemService to "mux".
// UnaryRPC     :call SystemServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterSystemServiceHandlerFromEndpoint instead.
func RegisterSystemServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server SystemServiceServer) error {

	mux.Handle("GET", pattern_SystemService_GetSystemInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/memos.api.v2.SystemService/GetSystemInfo", runtime.WithHTTPPathPattern("/api/v2/system/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SystemService_GetSystemInfo_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SystemService_GetSystemInfo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterSystemServiceHandlerFromEndpoint is same as RegisterSystemServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSystemServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSystemServiceHandler(ctx, mux, conn)
}

// RegisterSystemServiceHandler registers the http handlers for service SystemService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSystemServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSystemServiceHandlerClient(ctx, mux, NewSystemServiceClient(conn))
}

// RegisterSystemServiceHandlerClient registers the http handlers for service SystemService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SystemServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SystemServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SystemServiceClient" to call the correct interceptors.
func RegisterSystemServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SystemServiceClient) error {

	mux.Handle("GET", pattern_SystemService_GetSystemInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/memos.api.v2.SystemService/GetSystemInfo", runtime.WithHTTPPathPattern("/api/v2/system/info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SystemService_GetSystemInfo_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SystemService_GetSystemInfo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SystemService_GetSystemInfo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v2", "system", "info"}, ""))
)

var (
	forward_SystemService_GetSystemInfo_0 = runtime.ForwardResponseMessage
)
