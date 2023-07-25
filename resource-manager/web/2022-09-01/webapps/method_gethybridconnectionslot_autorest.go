package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHybridConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnection
}

// GetHybridConnectionSlot ...
func (c WebAppsClient) GetHybridConnectionSlot(ctx context.Context, id SlotHybridConnectionNamespaceRelayId) (result GetHybridConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForGetHybridConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHybridConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHybridConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetHybridConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHybridConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetHybridConnectionSlot prepares the GetHybridConnectionSlot request.
func (c WebAppsClient) preparerForGetHybridConnectionSlot(ctx context.Context, id SlotHybridConnectionNamespaceRelayId) (*http.Request, error) {
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

// responderForGetHybridConnectionSlot handles the response to the GetHybridConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetHybridConnectionSlot(resp *http.Response) (result GetHybridConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
