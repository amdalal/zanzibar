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

package mockpanictchannelworkflow

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/uber-go/tally"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	bazclientgeneratedmock "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz/mock-client"
	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/panic/module"
	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/panic/workflow"
	panictchannelendpointstatic "github.com/uber/zanzibar/examples/example-gateway/endpoints/tchannel/panic"
)

// NewSimpleServiceAnotherCallWorkflowMock creates a workflow with mock clients
func NewSimpleServiceAnotherCallWorkflowMock(t *testing.T) (workflow.SimpleServiceAnotherCallWorkflow, *MockClientNodes) {
	ctrl := gomock.NewController(t)

	initializedDefaultDependencies := &zanzibar.DefaultDependencies{
		Logger: zap.NewNop(),
		Scope:  tally.NewTestScope("", make(map[string]string)),
	}
	initializedDefaultDependencies.ContextLogger = zanzibar.NewContextLogger(initializedDefaultDependencies.Logger)
	contextExtractors := &zanzibar.ContextExtractors{}
	initializedDefaultDependencies.ContextExtractor = contextExtractors.MakeContextExtractor()

	initializedClientDependencies := &clientDependenciesNodes{}
	mockClientNodes := &MockClientNodes{
		Baz: bazclientgeneratedmock.NewMockClient(ctrl),
	}
	initializedClientDependencies.Baz = mockClientNodes.Baz

	w := panictchannelendpointstatic.NewSimpleServiceAnotherCallWorkflow(
		&module.Dependencies{
			Default: initializedDefaultDependencies,
			Client: &module.ClientDependencies{
				Baz: initializedClientDependencies.Baz,
			},
		},
	)

	return w, mockClientNodes
}
