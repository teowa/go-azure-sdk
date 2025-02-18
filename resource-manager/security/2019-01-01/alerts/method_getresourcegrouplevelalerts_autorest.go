package alerts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetResourceGroupLevelAlertsOperationResponse struct {
	HttpResponse *http.Response
	Model        *Alert
}

// GetResourceGroupLevelAlerts ...
func (c AlertsClient) GetResourceGroupLevelAlerts(ctx context.Context, id LocationAlertId) (result GetResourceGroupLevelAlertsOperationResponse, err error) {
	req, err := c.preparerForGetResourceGroupLevelAlerts(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "GetResourceGroupLevelAlerts", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "GetResourceGroupLevelAlerts", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetResourceGroupLevelAlerts(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "GetResourceGroupLevelAlerts", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetResourceGroupLevelAlerts prepares the GetResourceGroupLevelAlerts request.
func (c AlertsClient) preparerForGetResourceGroupLevelAlerts(ctx context.Context, id LocationAlertId) (*http.Request, error) {
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

// responderForGetResourceGroupLevelAlerts handles the response to the GetResourceGroupLevelAlerts request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForGetResourceGroupLevelAlerts(resp *http.Response) (result GetResourceGroupLevelAlertsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
