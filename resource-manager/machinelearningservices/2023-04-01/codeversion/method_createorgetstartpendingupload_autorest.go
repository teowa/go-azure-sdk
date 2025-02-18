package codeversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrGetStartPendingUploadOperationResponse struct {
	HttpResponse *http.Response
	Model        *PendingUploadResponseDto
}

// CreateOrGetStartPendingUpload ...
func (c CodeVersionClient) CreateOrGetStartPendingUpload(ctx context.Context, id CodeVersionId, input PendingUploadRequestDto) (result CreateOrGetStartPendingUploadOperationResponse, err error) {
	req, err := c.preparerForCreateOrGetStartPendingUpload(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "CreateOrGetStartPendingUpload", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "CreateOrGetStartPendingUpload", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrGetStartPendingUpload(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "CreateOrGetStartPendingUpload", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrGetStartPendingUpload prepares the CreateOrGetStartPendingUpload request.
func (c CodeVersionClient) preparerForCreateOrGetStartPendingUpload(ctx context.Context, id CodeVersionId, input PendingUploadRequestDto) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/startPendingUpload", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrGetStartPendingUpload handles the response to the CreateOrGetStartPendingUpload request. The method always
// closes the http.Response Body.
func (c CodeVersionClient) responderForCreateOrGetStartPendingUpload(resp *http.Response) (result CreateOrGetStartPendingUploadOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
