package vpngateways

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

type ResetOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetOperationOptions struct {
	IPConfigurationId *string
}

func DefaultResetOperationOptions() ResetOperationOptions {
	return ResetOperationOptions{}
}

func (o ResetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ResetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IPConfigurationId != nil {
		out.Append("ipConfigurationId", fmt.Sprintf("%v", *o.IPConfigurationId))
	}
	return &out
}

// Reset ...
func (c VpnGatewaysClient) Reset(ctx context.Context, id VpnGatewayId, options ResetOperationOptions) (result ResetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/reset", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

// ResetThenPoll performs Reset then polls until it's completed
func (c VpnGatewaysClient) ResetThenPoll(ctx context.Context, id VpnGatewayId, options ResetOperationOptions) error {
	result, err := c.Reset(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing Reset: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after Reset: %+v", err)
	}

	return nil
}
