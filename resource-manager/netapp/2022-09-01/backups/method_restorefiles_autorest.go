package backups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreFilesOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestoreFiles ...
func (c BackupsClient) RestoreFiles(ctx context.Context, id BackupId, input BackupRestoreFiles) (result RestoreFilesOperationResponse, err error) {
	req, err := c.preparerForRestoreFiles(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "RestoreFiles", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestoreFiles(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "RestoreFiles", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestoreFilesThenPoll performs RestoreFiles then polls until it's completed
func (c BackupsClient) RestoreFilesThenPoll(ctx context.Context, id BackupId, input BackupRestoreFiles) error {
	result, err := c.RestoreFiles(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreFiles: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestoreFiles: %+v", err)
	}

	return nil
}

// preparerForRestoreFiles prepares the RestoreFiles request.
func (c BackupsClient) preparerForRestoreFiles(ctx context.Context, id BackupId, input BackupRestoreFiles) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restoreFiles", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestoreFiles sends the RestoreFiles request. The method will close the
// http.Response Body if it receives an error.
func (c BackupsClient) senderForRestoreFiles(ctx context.Context, req *http.Request) (future RestoreFilesOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
