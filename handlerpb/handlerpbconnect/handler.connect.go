// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: handlerpb/handler.proto

package handlerpbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	handlerpb "reduction.dev/reduction-handler/handlerpb"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// HandlerName is the fully-qualified name of the Handler service.
	HandlerName = "dev.reduction.handler.Handler"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HandlerProcessEventBatchProcedure is the fully-qualified name of the Handler's ProcessEventBatch
	// RPC.
	HandlerProcessEventBatchProcedure = "/dev.reduction.handler.Handler/ProcessEventBatch"
	// HandlerKeyEventBatchProcedure is the fully-qualified name of the Handler's KeyEventBatch RPC.
	HandlerKeyEventBatchProcedure = "/dev.reduction.handler.Handler/KeyEventBatch"
)

// HandlerClient is a client for the dev.reduction.handler.Handler service.
type HandlerClient interface {
	ProcessEventBatch(context.Context, *connect.Request[handlerpb.ProcessEventBatchRequest]) (*connect.Response[handlerpb.ProcessEventBatchResponse], error)
	KeyEventBatch(context.Context, *connect.Request[handlerpb.KeyEventBatchRequest]) (*connect.Response[handlerpb.KeyEventBatchResponse], error)
}

// NewHandlerClient constructs a client for the dev.reduction.handler.Handler service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHandlerClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) HandlerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	handlerMethods := handlerpb.File_handlerpb_handler_proto.Services().ByName("Handler").Methods()
	return &handlerClient{
		processEventBatch: connect.NewClient[handlerpb.ProcessEventBatchRequest, handlerpb.ProcessEventBatchResponse](
			httpClient,
			baseURL+HandlerProcessEventBatchProcedure,
			connect.WithSchema(handlerMethods.ByName("ProcessEventBatch")),
			connect.WithClientOptions(opts...),
		),
		keyEventBatch: connect.NewClient[handlerpb.KeyEventBatchRequest, handlerpb.KeyEventBatchResponse](
			httpClient,
			baseURL+HandlerKeyEventBatchProcedure,
			connect.WithSchema(handlerMethods.ByName("KeyEventBatch")),
			connect.WithClientOptions(opts...),
		),
	}
}

// handlerClient implements HandlerClient.
type handlerClient struct {
	processEventBatch *connect.Client[handlerpb.ProcessEventBatchRequest, handlerpb.ProcessEventBatchResponse]
	keyEventBatch     *connect.Client[handlerpb.KeyEventBatchRequest, handlerpb.KeyEventBatchResponse]
}

// ProcessEventBatch calls dev.reduction.handler.Handler.ProcessEventBatch.
func (c *handlerClient) ProcessEventBatch(ctx context.Context, req *connect.Request[handlerpb.ProcessEventBatchRequest]) (*connect.Response[handlerpb.ProcessEventBatchResponse], error) {
	return c.processEventBatch.CallUnary(ctx, req)
}

// KeyEventBatch calls dev.reduction.handler.Handler.KeyEventBatch.
func (c *handlerClient) KeyEventBatch(ctx context.Context, req *connect.Request[handlerpb.KeyEventBatchRequest]) (*connect.Response[handlerpb.KeyEventBatchResponse], error) {
	return c.keyEventBatch.CallUnary(ctx, req)
}

// HandlerHandler is an implementation of the dev.reduction.handler.Handler service.
type HandlerHandler interface {
	ProcessEventBatch(context.Context, *connect.Request[handlerpb.ProcessEventBatchRequest]) (*connect.Response[handlerpb.ProcessEventBatchResponse], error)
	KeyEventBatch(context.Context, *connect.Request[handlerpb.KeyEventBatchRequest]) (*connect.Response[handlerpb.KeyEventBatchResponse], error)
}

// NewHandlerHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHandlerHandler(svc HandlerHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	handlerMethods := handlerpb.File_handlerpb_handler_proto.Services().ByName("Handler").Methods()
	handlerProcessEventBatchHandler := connect.NewUnaryHandler(
		HandlerProcessEventBatchProcedure,
		svc.ProcessEventBatch,
		connect.WithSchema(handlerMethods.ByName("ProcessEventBatch")),
		connect.WithHandlerOptions(opts...),
	)
	handlerKeyEventBatchHandler := connect.NewUnaryHandler(
		HandlerKeyEventBatchProcedure,
		svc.KeyEventBatch,
		connect.WithSchema(handlerMethods.ByName("KeyEventBatch")),
		connect.WithHandlerOptions(opts...),
	)
	return "/dev.reduction.handler.Handler/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HandlerProcessEventBatchProcedure:
			handlerProcessEventBatchHandler.ServeHTTP(w, r)
		case HandlerKeyEventBatchProcedure:
			handlerKeyEventBatchHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHandlerHandler returns CodeUnimplemented from all methods.
type UnimplementedHandlerHandler struct{}

func (UnimplementedHandlerHandler) ProcessEventBatch(context.Context, *connect.Request[handlerpb.ProcessEventBatchRequest]) (*connect.Response[handlerpb.ProcessEventBatchResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("dev.reduction.handler.Handler.ProcessEventBatch is not implemented"))
}

func (UnimplementedHandlerHandler) KeyEventBatch(context.Context, *connect.Request[handlerpb.KeyEventBatchRequest]) (*connect.Response[handlerpb.KeyEventBatchResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("dev.reduction.handler.Handler.KeyEventBatch is not implemented"))
}
