// Copyright (c) 2017 Uber Technologies, Inc.
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

package client

import (
	"time"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/.gen/go/admin/adminserviceclient"
	"github.com/uber/cadence/.gen/go/cadence/workflowserviceclient"
	"github.com/uber/cadence/.gen/go/history/historyserviceclient"
	"github.com/uber/cadence/.gen/go/matching/matchingserviceclient"
	"github.com/uber/cadence/client/admin"
	"github.com/uber/cadence/client/frontend"
	"github.com/uber/cadence/client/history"
	"github.com/uber/cadence/client/matching"
	"github.com/uber/cadence/client/public"
	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/membership"
	"github.com/uber/cadence/common/metrics"
	publicClientInterface "go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
)

const (
	publicCaller   = "cadence-public-client"
	frontendCaller = "cadence-frontend-client"
	historyCaller  = "history-service-client"
	matchingCaller = "matching-service-client"
	crossDCCaller  = "cadence-xdc-client"
)

const (
	clientKeyDispatcher = "client-key-dispatcher"
)

// Factory can be used to create RPC clients for cadence services
type Factory interface {
	NewHistoryClient() (history.Client, error)
	NewMatchingClient() (matching.Client, error)
	NewFrontendClient() (frontend.Client, error)
	NewPublicClient() (public.Client, error)

	NewHistoryClientWithTimeout(timeout time.Duration) (history.Client, error)
	NewMatchingClientWithTimeout(timeout time.Duration, longPollTimeout time.Duration) (matching.Client, error)
	NewFrontendClientWithTimeout(timeout time.Duration, longPollTimeout time.Duration) (frontend.Client, error)
	NewPublicClientWithTimeout(timeout time.Duration, longPollTimeout time.Duration) (public.Client, error)

	NewAdminClientWithTimeoutAndDispatcher(rpcName string, timeout time.Duration, dispatcher *yarpc.Dispatcher) (admin.Client, error)
	NewFrontendClientWithTimeoutAndDispatcher(rpcName string, timeout time.Duration, longPollTimeout time.Duration, dispatcher *yarpc.Dispatcher) (frontend.Client, error)
}

type rpcClientFactory struct {
	rpcFactory            common.RPCFactory
	monitor               membership.Monitor
	metricsClient         metrics.Client
	numberOfHistoryShards int
}

// NewRPCClientFactory creates an instance of client factory that knows how to dispatch RPC calls.
func NewRPCClientFactory(rpcFactory common.RPCFactory, monitor membership.Monitor,
	metricsClient metrics.Client, numberOfHistoryShards int) Factory {
	return &rpcClientFactory{
		rpcFactory:            rpcFactory,
		monitor:               monitor,
		metricsClient:         metricsClient,
		numberOfHistoryShards: numberOfHistoryShards,
	}
}

func (cf *rpcClientFactory) NewHistoryClient() (history.Client, error) {
	return cf.NewHistoryClientWithTimeout(history.DefaultTimeout)
}

func (cf *rpcClientFactory) NewMatchingClient() (matching.Client, error) {
	return cf.NewMatchingClientWithTimeout(matching.DefaultTimeout, matching.DefaultLongPollTimeout)
}

func (cf *rpcClientFactory) NewFrontendClient() (frontend.Client, error) {
	return cf.NewFrontendClientWithTimeout(frontend.DefaultTimeout, frontend.DefaultLongPollTimeout)
}

func (cf *rpcClientFactory) NewPublicClient() (public.Client, error) {
	return cf.NewPublicClientWithTimeout(public.DefaultTimeout, public.DefaultLongPollTimeout)
}

func (cf *rpcClientFactory) NewHistoryClientWithTimeout(timeout time.Duration) (history.Client, error) {
	resolver, err := cf.monitor.GetResolver(common.HistoryServiceName)
	if err != nil {
		return nil, err
	}

	keyResolver := func(key string) (string, error) {
		host, err := resolver.Lookup(key)
		if err != nil {
			return "", err
		}
		return host.GetAddress(), nil
	}

	clientProvider := func(clientKey string) (interface{}, error) {
		dispatcher := cf.rpcFactory.CreateDispatcherForOutbound(historyCaller, common.HistoryServiceName, clientKey)
		return historyserviceclient.New(dispatcher.ClientConfig(common.HistoryServiceName)), nil
	}

	client := history.NewClient(cf.numberOfHistoryShards, timeout, common.NewClientCache(keyResolver, clientProvider))
	if cf.metricsClient != nil {
		client = history.NewMetricClient(client, cf.metricsClient)
	}
	return client, nil
}

func (cf *rpcClientFactory) NewMatchingClientWithTimeout(
	timeout time.Duration,
	longPollTimeout time.Duration,
) (matching.Client, error) {
	resolver, err := cf.monitor.GetResolver(common.MatchingServiceName)
	if err != nil {
		return nil, err
	}

	keyResolver := func(key string) (string, error) {
		host, err := resolver.Lookup(key)
		if err != nil {
			return "", err
		}
		return host.GetAddress(), nil
	}

	clientProvider := func(clientKey string) (interface{}, error) {
		dispatcher := cf.rpcFactory.CreateDispatcherForOutbound(matchingCaller, common.MatchingServiceName, clientKey)
		return matchingserviceclient.New(dispatcher.ClientConfig(common.MatchingServiceName)), nil
	}

	client := matching.NewClient(timeout, longPollTimeout, common.NewClientCache(keyResolver, clientProvider))
	if cf.metricsClient != nil {
		client = matching.NewMetricClient(client, cf.metricsClient)
	}
	return client, nil

}

func (cf *rpcClientFactory) NewFrontendClientWithTimeout(
	timeout time.Duration,
	longPollTimeout time.Duration,
) (frontend.Client, error) {

	resolver, err := cf.monitor.GetResolver(common.FrontendServiceName)
	if err != nil {
		return nil, err
	}

	keyResolver := func(key string) (string, error) {
		host, err := resolver.Lookup(key)
		if err != nil {
			return "", err
		}
		return host.GetAddress(), nil
	}

	clientProvider := func(clientKey string) (interface{}, error) {
		dispatcher := cf.rpcFactory.CreateDispatcherForOutbound(frontendCaller, common.FrontendServiceName, clientKey)
		return workflowserviceclient.New(dispatcher.ClientConfig(common.FrontendServiceName)), nil
	}

	client := frontend.NewClient(timeout, longPollTimeout, common.NewClientCache(keyResolver, clientProvider))
	if cf.metricsClient != nil {
		client = frontend.NewMetricClient(client, cf.metricsClient)
	}
	return client, nil
}

func (cf *rpcClientFactory) NewPublicClientWithTimeout(
	timeout time.Duration,
	longPollTimeout time.Duration,
) (public.Client, error) {

	// public client and frontend client are essentially the same,
	// except the interface definition
	resolver, err := cf.monitor.GetResolver(common.FrontendServiceName)
	if err != nil {
		return nil, err
	}

	keyResolver := func(key string) (string, error) {
		host, err := resolver.Lookup(key)
		if err != nil {
			return "", err
		}
		return host.GetAddress(), nil
	}

	clientProvider := func(clientKey string) (interface{}, error) {
		dispatcher := cf.rpcFactory.CreateDispatcherForOutbound(publicCaller, common.FrontendServiceName, clientKey)
		return publicClientInterface.New(dispatcher.ClientConfig(common.FrontendServiceName)), nil
	}

	client := public.NewClient(timeout, longPollTimeout, common.NewClientCache(keyResolver, clientProvider))
	if cf.metricsClient != nil {
		client = public.NewMetricClient(client, cf.metricsClient)
	}
	return client, nil
}

func (cf *rpcClientFactory) NewAdminClientWithTimeoutAndDispatcher(
	rpcName string,
	timeout time.Duration,
	dispatcher *yarpc.Dispatcher,
) (admin.Client, error) {
	keyResolver := func(key string) (string, error) {
		return clientKeyDispatcher, nil
	}

	clientProvider := func(clientKey string) (interface{}, error) {
		return adminserviceclient.New(dispatcher.ClientConfig(rpcName)), nil
	}

	client := admin.NewClient(timeout, common.NewClientCache(keyResolver, clientProvider))
	if cf.metricsClient != nil {
		client = admin.NewMetricClient(client, cf.metricsClient)
	}
	return client, nil
}

func (cf *rpcClientFactory) NewFrontendClientWithTimeoutAndDispatcher(
	rpcName string,
	timeout time.Duration,
	longPollTimeout time.Duration,
	dispatcher *yarpc.Dispatcher,
) (frontend.Client, error) {
	keyResolver := func(key string) (string, error) {
		return clientKeyDispatcher, nil
	}

	clientProvider := func(clientKey string) (interface{}, error) {
		return workflowserviceclient.New(dispatcher.ClientConfig(rpcName)), nil
	}

	client := frontend.NewClient(timeout, longPollTimeout, common.NewClientCache(keyResolver, clientProvider))
	if cf.metricsClient != nil {
		client = frontend.NewMetricClient(client, cf.metricsClient)
	}
	return client, nil
}
