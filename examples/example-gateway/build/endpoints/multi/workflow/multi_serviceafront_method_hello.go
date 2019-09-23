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

	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi/module"
	"go.uber.org/zap"
)

// ServiceAFrontHelloWorkflow defines the interface for ServiceAFrontHello workflow
type ServiceAFrontHelloWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
	) (string, zanzibar.Header, error)
}

// NewServiceAFrontHelloWorkflow creates a workflow
func NewServiceAFrontHelloWorkflow(deps *module.Dependencies) ServiceAFrontHelloWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.multi.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.multi.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
		}
	}

	return &serviceAFrontHelloWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// serviceAFrontHelloWorkflow calls thrift client Multi.HelloA
type serviceAFrontHelloWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w serviceAFrontHelloWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (string, zanzibar.Header, error) {

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

	clientRespBody, _, err := w.Clients.Multi.HelloA(
		ctx, clientHeaders,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Multi"),
			)

			return "", nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertServiceABackHelloClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertServiceABackHelloClientResponse(in string) string {
	out := in

	return out
}
