package azurestack

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RegistrationsClient is the azure Stack
type RegistrationsClient struct {
	BaseClient
}

// NewRegistrationsClient creates an instance of the RegistrationsClient client.
func NewRegistrationsClient(subscriptionID string) RegistrationsClient {
	return NewRegistrationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegistrationsClientWithBaseURI creates an instance of the RegistrationsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRegistrationsClientWithBaseURI(baseURI string, subscriptionID string) RegistrationsClient {
	return RegistrationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update an Azure Stack registration.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// tokenParameter - registration token
func (client RegistrationsClient) CreateOrUpdate(ctx context.Context, resourceGroup string, registrationName string, tokenParameter RegistrationParameter) (result Registration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: tokenParameter,
			Constraints: []validation.Constraint{{Target: "tokenParameter.RegistrationParameterProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "tokenParameter.RegistrationParameterProperties.RegistrationToken", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("azurestack.RegistrationsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroup, registrationName, tokenParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RegistrationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroup string, registrationName string, tokenParameter RegistrationParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", pathParameters),
		autorest.WithJSON(tokenParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RegistrationsClient) CreateOrUpdateResponder(resp *http.Response) (result Registration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the requested Azure Stack registration.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
func (client RegistrationsClient) Delete(ctx context.Context, resourceGroup string, registrationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroup, registrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegistrationsClient) DeletePreparer(ctx context.Context, resourceGroup string, registrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegistrationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the properties of an Azure Stack registration.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
func (client RegistrationsClient) Get(ctx context.Context, resourceGroup string, registrationName string) (result Registration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroup, registrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegistrationsClient) GetPreparer(ctx context.Context, resourceGroup string, registrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegistrationsClient) GetResponder(resp *http.Response) (result Registration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetActivationKey returns Azure Stack Activation Key.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
func (client RegistrationsClient) GetActivationKey(ctx context.Context, resourceGroup string, registrationName string) (result ActivationKeyResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationsClient.GetActivationKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetActivationKeyPreparer(ctx, resourceGroup, registrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "GetActivationKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetActivationKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "GetActivationKey", resp, "Failure sending request")
		return
	}

	result, err = client.GetActivationKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "GetActivationKey", resp, "Failure responding to request")
	}

	return
}

// GetActivationKeyPreparer prepares the GetActivationKey request.
func (client RegistrationsClient) GetActivationKeyPreparer(ctx context.Context, resourceGroup string, registrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/getactivationkey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetActivationKeySender sends the GetActivationKey request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationsClient) GetActivationKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetActivationKeyResponder handles the response to the GetActivationKey request. The method always
// closes the http.Response Body.
func (client RegistrationsClient) GetActivationKeyResponder(resp *http.Response) (result ActivationKeyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns a list of all registrations.
// Parameters:
// resourceGroup - name of the resource group.
func (client RegistrationsClient) List(ctx context.Context, resourceGroup string) (result RegistrationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationsClient.List")
		defer func() {
			sc := -1
			if result.rl.Response.Response != nil {
				sc = result.rl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "List", resp, "Failure sending request")
		return
	}

	result.rl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RegistrationsClient) ListPreparer(ctx context.Context, resourceGroup string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroup":  autorest.Encode("path", resourceGroup),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegistrationsClient) ListResponder(resp *http.Response) (result RegistrationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RegistrationsClient) listNextResults(ctx context.Context, lastResults RegistrationList) (result RegistrationList, err error) {
	req, err := lastResults.registrationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegistrationsClient) ListComplete(ctx context.Context, resourceGroup string) (result RegistrationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroup)
	return
}

// Update patch an Azure Stack registration.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// tokenParameter - registration token
func (client RegistrationsClient) Update(ctx context.Context, resourceGroup string, registrationName string, tokenParameter RegistrationParameter) (result Registration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroup, registrationName, tokenParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.RegistrationsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RegistrationsClient) UpdatePreparer(ctx context.Context, resourceGroup string, registrationName string, tokenParameter RegistrationParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", pathParameters),
		autorest.WithJSON(tokenParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RegistrationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RegistrationsClient) UpdateResponder(resp *http.Response) (result Registration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
