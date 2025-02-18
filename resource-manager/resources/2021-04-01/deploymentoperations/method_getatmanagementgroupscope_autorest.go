package deploymentoperations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAtManagementGroupScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentOperation
}

// GetAtManagementGroupScope ...
func (c DeploymentOperationsClient) GetAtManagementGroupScope(ctx context.Context, id Providers2DeploymentOperationId) (result GetAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForGetAtManagementGroupScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAtManagementGroupScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "GetAtManagementGroupScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAtManagementGroupScope prepares the GetAtManagementGroupScope request.
func (c DeploymentOperationsClient) preparerForGetAtManagementGroupScope(ctx context.Context, id Providers2DeploymentOperationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAtManagementGroupScope handles the response to the GetAtManagementGroupScope request. The method always
// closes the http.Response Body.
func (c DeploymentOperationsClient) responderForGetAtManagementGroupScope(resp *http.Response) (result GetAtManagementGroupScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
