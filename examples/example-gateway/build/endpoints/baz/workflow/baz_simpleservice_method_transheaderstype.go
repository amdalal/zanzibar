// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workflow

import (
	"context"
	"net/textproto"
	"strconv"

	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceTransHeadersTypeWorkflow defines the interface for SimpleServiceTransHeadersType workflow
type SimpleServiceTransHeadersTypeWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBazBaz.SimpleService_TransHeadersType_Args,
	) (*endpointsBazBaz.TransHeader, zanzibar.Header, error)
}

// NewSimpleServiceTransHeadersTypeWorkflow creates a workflow
func NewSimpleServiceTransHeadersTypeWorkflow(deps *module.Dependencies) SimpleServiceTransHeadersTypeWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.baz.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.baz.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
		}
	}

	return &simpleServiceTransHeadersTypeWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// simpleServiceTransHeadersTypeWorkflow calls thrift client Baz.TransHeadersType
type simpleServiceTransHeadersTypeWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w simpleServiceTransHeadersTypeWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_TransHeadersType_Args,
) (*endpointsBazBaz.TransHeader, zanzibar.Header, error) {
	clientRequest := convertToTransHeadersTypeClientRequest(r)
	clientRequest = propagateHeadersTransHeadersTypeClientRequests(clientRequest, reqHeaders)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	for _, whitelistedHeader := range w.whitelistedDynamicHeaders {
		headerVal, ok := reqHeaders.Get(whitelistedHeader)
		if ok {
			clientHeaders[whitelistedHeader] = headerVal
		}
	}

	clientRespBody, _, err := w.Clients.Baz.TransHeadersType(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertTransHeadersTypeAuthErr(
				errValue,
			)

			return nil, nil, serverErr

		case *clientsBazBaz.OtherAuthErr:
			serverErr := convertTransHeadersTypeOtherAuthErr(
				errValue,
			)

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceTransHeadersTypeClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToTransHeadersTypeClientRequest(in *endpointsBazBaz.SimpleService_TransHeadersType_Args) *clientsBazBaz.SimpleService_TransHeadersType_Args {
	out := &clientsBazBaz.SimpleService_TransHeadersType_Args{}

	if in.Req != nil {
		out.Req = &clientsBazBaz.TransHeaderType{}
	} else {
		out.Req = nil
	}

	return out
}

func convertTransHeadersTypeAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
func convertTransHeadersTypeOtherAuthErr(
	clientError *clientsBazBaz.OtherAuthErr,
) *endpointsBazBaz.OtherAuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.OtherAuthErr{}
	return serverError
}

func convertSimpleServiceTransHeadersTypeClientResponse(in *clientsBazBaz.TransHeaderType) *endpointsBazBaz.TransHeader {
	out := &endpointsBazBaz.TransHeader{}

	return out
}

func propagateHeadersTransHeadersTypeClientRequests(in *clientsBazBaz.SimpleService_TransHeadersType_Args, headers zanzibar.Header) *clientsBazBaz.SimpleService_TransHeadersType_Args {
	if in == nil {
		in = &clientsBazBaz.SimpleService_TransHeadersType_Args{}
	}
	if key, ok := headers.Get("x-boolean"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBaz.TransHeaderType{}
		}
		if v, err := strconv.ParseBool(key); err == nil {
			in.Req.B1 = v
		}

	}
	if key, ok := headers.Get("x-float"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBaz.TransHeaderType{}
		}
		if v, err := strconv.ParseFloat(key, 64); err == nil {
			in.Req.F3 = &v
		}

	}
	if key, ok := headers.Get("x-int"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBaz.TransHeaderType{}
		}
		if v, err := strconv.ParseInt(key, 10, 32); err == nil {
			val := int32(v)
			in.Req.I1 = &val
		}

	}
	if key, ok := headers.Get("x-int"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBaz.TransHeaderType{}
		}
		if v, err := strconv.ParseInt(key, 10, 64); err == nil {
			in.Req.I2 = v
		}

	}
	if key, ok := headers.Get("x-string"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBaz.TransHeaderType{}
		}
		in.Req.S6 = key

	}
	if key, ok := headers.Get("x-string"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBaz.TransHeaderType{}
		}
		val := clientsBazBaz.UUID(key)
		in.Req.U4 = val

	}
	if key, ok := headers.Get("x-string"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBaz.TransHeaderType{}
		}
		val := clientsBazBaz.UUID(key)
		in.Req.U5 = &val

	}
	return in
}
