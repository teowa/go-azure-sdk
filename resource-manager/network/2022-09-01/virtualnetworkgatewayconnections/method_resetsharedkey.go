package virtualnetworkgatewayconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetSharedKeyOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ResetSharedKey ...
func (c VirtualNetworkGatewayConnectionsClient) ResetSharedKey(ctx context.Context, id ConnectionId, input ConnectionResetSharedKey) (result ResetSharedKeyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/sharedkey/reset", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// ResetSharedKeyThenPoll performs ResetSharedKey then polls until it's completed
func (c VirtualNetworkGatewayConnectionsClient) ResetSharedKeyThenPoll(ctx context.Context, id ConnectionId, input ConnectionResetSharedKey) error {
	result, err := c.ResetSharedKey(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ResetSharedKey: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ResetSharedKey: %+v", err)
	}

	return nil
}
