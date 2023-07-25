package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteHybridConnectionOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteHybridConnection ...
func (c WebAppsClient) DeleteHybridConnection(ctx context.Context, id RelayId) (result DeleteHybridConnectionOperationResponse, err error) {
	req, err := c.preparerForDeleteHybridConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHybridConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHybridConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteHybridConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHybridConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteHybridConnection prepares the DeleteHybridConnection request.
func (c WebAppsClient) preparerForDeleteHybridConnection(ctx context.Context, id RelayId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteHybridConnection handles the response to the DeleteHybridConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteHybridConnection(resp *http.Response) (result DeleteHybridConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
