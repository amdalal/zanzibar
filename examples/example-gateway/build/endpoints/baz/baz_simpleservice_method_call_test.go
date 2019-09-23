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

package bazendpoint

import (
	"bytes"
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/config"
	testbackend "github.com/uber/zanzibar/test/lib/test_backend"
	testGateway "github.com/uber/zanzibar/test/lib/test_gateway"
	"github.com/uber/zanzibar/test/lib/util"

	bazclient "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
)

func TestCallSuccessfulRequestOKResponse(t *testing.T) {

	confFiles := util.DefaultConfigFiles("example-gateway")
	staticConf := config.NewRuntimeConfigOrDie(confFiles, map[string]interface{}{})
	var alternateServiceDetail config.AlternateServiceDetail
	if staticConf.ContainsKey("clients.baz.alternates") {
		staticConf.MustGetStruct("clients.baz.alternates", &alternateServiceDetail)
	}
	var backends []*testbackend.TestTChannelBackend
	for serviceName := range alternateServiceDetail.ServicesDetailMap {
		if serviceName == "nomatch" {
			continue
		}
		backend, err := testbackend.CreateTChannelBackend(int32(0), serviceName)
		assert.NoError(t, err)
		err = backend.Bootstrap()
		assert.NoError(t, err)
		backends = append(backends, backend)
	}

	gateway, err := testGateway.CreateGateway(t, map[string]interface{}{
		"clients.baz.serviceName": "bazService",
	}, &testGateway.Options{
		KnownTChannelBackends: []string{"baz"},
		TestBinary:            util.DefaultMainFile("example-gateway"),
		ConfigFiles:           confFiles,
		Backends:              backends,
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	fakeCall := func(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Call_Args,
	) (map[string]string, error) {

		var resHeaders map[string]string

		return resHeaders, nil
	}

	headers := map[string]string{}
	err = gateway.TChannelBackends()["baz"].Register(
		"baz", "call", "SimpleService::call",
		bazclient.NewSimpleServiceCallHandler(fakeCall),
	)
	assert.NoError(t, err)
	makeRequestAndValidateCallSuccessfulRequest(t, gateway, headers)

	i := 1
	for serviceName := range alternateServiceDetail.ServicesDetailMap {
		headers := map[string]string{}

		if serviceName == "nomatch" {
			headers["x-container"] = "randomstr"
			headers["x-test-Env"] = "randomstr"
		} else {
			if i == 1 {
				headers["x-container"] = "sandbox"
			} else if i == 2 {
				headers["x-test-Env"] = "test1"
			}
			err = gateway.TChannelBackends()["baz:"+strconv.Itoa(i)].Register(
				"baz", "call", "SimpleService::call",
				bazclient.NewSimpleServiceCallHandler(fakeCall),
			)
		}

		makeRequestAndValidateCallSuccessfulRequest(t, gateway, headers)
		i++
	}

}

func makeRequestAndValidateCallSuccessfulRequest(t *testing.T, gateway testGateway.TestGateway, headers map[string]string) {
	headers["x-token"] = "token"
	headers["x-uuid"] = "uuid"

	endpointRequest := []byte(`{"arg":{"b1":true,"i3":42,"s2":"hello"}}`)

	res, err := gateway.MakeRequest(
		"POST",
		"/baz/call",
		headers,
		bytes.NewReader(endpointRequest),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	assert.Equal(t, 204, res.StatusCode)
}
