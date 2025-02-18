package p2svpngateways

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetP2sVpnConnectionHealthDetailedOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetP2sVpnConnectionHealthDetailed ...
func (c P2sVpnGatewaysClient) GetP2sVpnConnectionHealthDetailed(ctx context.Context, id commonids.VirtualWANP2SVPNGatewayId, input P2SVpnConnectionHealthRequest) (result GetP2sVpnConnectionHealthDetailedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/getP2sVpnConnectionHealthDetailed", id.ID()),
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

// GetP2sVpnConnectionHealthDetailedThenPoll performs GetP2sVpnConnectionHealthDetailed then polls until it's completed
func (c P2sVpnGatewaysClient) GetP2sVpnConnectionHealthDetailedThenPoll(ctx context.Context, id commonids.VirtualWANP2SVPNGatewayId, input P2SVpnConnectionHealthRequest) error {
	result, err := c.GetP2sVpnConnectionHealthDetailed(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GetP2sVpnConnectionHealthDetailed: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GetP2sVpnConnectionHealthDetailed: %+v", err)
	}

	return nil
}
