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

type DeleteSourceControlSlotOperationResponse struct {
	HttpResponse *http.Response
}

type DeleteSourceControlSlotOperationOptions struct {
	AdditionalFlags *string
}

func DefaultDeleteSourceControlSlotOperationOptions() DeleteSourceControlSlotOperationOptions {
	return DeleteSourceControlSlotOperationOptions{}
}

func (o DeleteSourceControlSlotOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DeleteSourceControlSlotOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.AdditionalFlags != nil {
		out["additionalFlags"] = *o.AdditionalFlags
	}

	return out
}

// DeleteSourceControlSlot ...
func (c WebAppsClient) DeleteSourceControlSlot(ctx context.Context, id SlotId, options DeleteSourceControlSlotOperationOptions) (result DeleteSourceControlSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteSourceControlSlot(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSourceControlSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSourceControlSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSourceControlSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSourceControlSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSourceControlSlot prepares the DeleteSourceControlSlot request.
func (c WebAppsClient) preparerForDeleteSourceControlSlot(ctx context.Context, id SlotId, options DeleteSourceControlSlotOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/sourceControls/web", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteSourceControlSlot handles the response to the DeleteSourceControlSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteSourceControlSlot(resp *http.Response) (result DeleteSourceControlSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
