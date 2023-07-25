package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSyncStatusSlotOperationResponse struct {
	HttpResponse *http.Response
}

// ListSyncStatusSlot ...
func (c WebAppsClient) ListSyncStatusSlot(ctx context.Context, id SlotId) (result ListSyncStatusSlotOperationResponse, err error) {
	req, err := c.preparerForListSyncStatusSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncStatusSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncStatusSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListSyncStatusSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncStatusSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListSyncStatusSlot prepares the ListSyncStatusSlot request.
func (c WebAppsClient) preparerForListSyncStatusSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/host/default/listsyncstatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSyncStatusSlot handles the response to the ListSyncStatusSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSyncStatusSlot(resp *http.Response) (result ListSyncStatusSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
