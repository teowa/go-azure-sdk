package virtualwans

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

type VirtualHubRouteTableV2sCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// VirtualHubRouteTableV2sCreateOrUpdate ...
func (c VirtualWANsClient) VirtualHubRouteTableV2sCreateOrUpdate(ctx context.Context, id RouteTableId, input VirtualHubRouteTableV2) (result VirtualHubRouteTableV2sCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

// VirtualHubRouteTableV2sCreateOrUpdateThenPoll performs VirtualHubRouteTableV2sCreateOrUpdate then polls until it's completed
func (c VirtualWANsClient) VirtualHubRouteTableV2sCreateOrUpdateThenPoll(ctx context.Context, id RouteTableId, input VirtualHubRouteTableV2) error {
	result, err := c.VirtualHubRouteTableV2sCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing VirtualHubRouteTableV2sCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after VirtualHubRouteTableV2sCreateOrUpdate: %+v", err)
	}

	return nil
}
