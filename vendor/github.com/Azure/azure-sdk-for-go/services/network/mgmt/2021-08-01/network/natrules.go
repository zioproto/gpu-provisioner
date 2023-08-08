package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NatRulesClient is the network Client
type NatRulesClient struct {
	BaseClient
}

// NewNatRulesClient creates an instance of the NatRulesClient client.
func NewNatRulesClient(subscriptionID string) NatRulesClient {
	return NewNatRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNatRulesClientWithBaseURI creates an instance of the NatRulesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNatRulesClientWithBaseURI(baseURI string, subscriptionID string) NatRulesClient {
	return NatRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a nat rule to a scalable vpn gateway if it doesn't exist else updates the existing nat rules.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// natRuleName - the name of the nat rule.
// natRuleParameters - parameters supplied to create or Update a Nat Rule.
func (client NatRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VpnGatewayNatRule) (result NatRulesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, gatewayName, natRuleName, natRuleParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client NatRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters VpnGatewayNatRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"natRuleName":       autorest.Encode("path", natRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	natRuleParameters.Etag = nil
	natRuleParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules/{natRuleName}", pathParameters),
		autorest.WithJSON(natRuleParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client NatRulesClient) CreateOrUpdateSender(req *http.Request) (future NatRulesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client NatRulesClient) CreateOrUpdateResponder(resp *http.Response) (result VpnGatewayNatRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a nat rule.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// natRuleName - the name of the nat rule.
func (client NatRulesClient) Delete(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string) (result NatRulesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatRulesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, gatewayName, natRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NatRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"natRuleName":       autorest.Encode("path", natRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules/{natRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NatRulesClient) DeleteSender(req *http.Request) (future NatRulesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NatRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a nat ruleGet.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// natRuleName - the name of the nat rule.
func (client NatRulesClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string) (result VpnGatewayNatRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, gatewayName, natRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client NatRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"natRuleName":       autorest.Encode("path", natRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules/{natRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NatRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NatRulesClient) GetResponder(resp *http.Response) (result VpnGatewayNatRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVpnGateway retrieves all nat rules for a particular virtual wan vpn gateway.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
func (client NatRulesClient) ListByVpnGateway(ctx context.Context, resourceGroupName string, gatewayName string) (result ListVpnGatewayNatRulesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatRulesClient.ListByVpnGateway")
		defer func() {
			sc := -1
			if result.lvgnrr.Response.Response != nil {
				sc = result.lvgnrr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVpnGatewayNextResults
	req, err := client.ListByVpnGatewayPreparer(ctx, resourceGroupName, gatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "ListByVpnGateway", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVpnGatewaySender(req)
	if err != nil {
		result.lvgnrr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "ListByVpnGateway", resp, "Failure sending request")
		return
	}

	result.lvgnrr, err = client.ListByVpnGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "ListByVpnGateway", resp, "Failure responding to request")
		return
	}
	if result.lvgnrr.hasNextLink() && result.lvgnrr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVpnGatewayPreparer prepares the ListByVpnGateway request.
func (client NatRulesClient) ListByVpnGatewayPreparer(ctx context.Context, resourceGroupName string, gatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/natRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVpnGatewaySender sends the ListByVpnGateway request. The method will close the
// http.Response Body if it receives an error.
func (client NatRulesClient) ListByVpnGatewaySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVpnGatewayResponder handles the response to the ListByVpnGateway request. The method always
// closes the http.Response Body.
func (client NatRulesClient) ListByVpnGatewayResponder(resp *http.Response) (result ListVpnGatewayNatRulesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVpnGatewayNextResults retrieves the next set of results, if any.
func (client NatRulesClient) listByVpnGatewayNextResults(ctx context.Context, lastResults ListVpnGatewayNatRulesResult) (result ListVpnGatewayNatRulesResult, err error) {
	req, err := lastResults.listVpnGatewayNatRulesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NatRulesClient", "listByVpnGatewayNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVpnGatewaySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.NatRulesClient", "listByVpnGatewayNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVpnGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatRulesClient", "listByVpnGatewayNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVpnGatewayComplete enumerates all values, automatically crossing page boundaries as required.
func (client NatRulesClient) ListByVpnGatewayComplete(ctx context.Context, resourceGroupName string, gatewayName string) (result ListVpnGatewayNatRulesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatRulesClient.ListByVpnGateway")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVpnGateway(ctx, resourceGroupName, gatewayName)
	return
}
