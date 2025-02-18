package apipolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByApiOperationResponse struct {
	HttpResponse *http.Response
	Model        *PolicyCollection
}

// ListByApi ...
func (c ApiPolicyClient) ListByApi(ctx context.Context, id ApiId) (result ListByApiOperationResponse, err error) {
	req, err := c.preparerForListByApi(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apipolicy.ApiPolicyClient", "ListByApi", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apipolicy.ApiPolicyClient", "ListByApi", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByApi(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apipolicy.ApiPolicyClient", "ListByApi", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByApi prepares the ListByApi request.
func (c ApiPolicyClient) preparerForListByApi(ctx context.Context, id ApiId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/policies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByApi handles the response to the ListByApi request. The method always
// closes the http.Response Body.
func (c ApiPolicyClient) responderForListByApi(resp *http.Response) (result ListByApiOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
