package hosts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Host
}

type ListByResourceGroupCompleteResult struct {
	Items []Host
}

// ListByResourceGroup ...
func (c HostsClient) ListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.ConnectedVMwarevSphere/hosts", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Host `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByResourceGroupComplete retrieves all the results into a single object
func (c HostsClient) ListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ListByResourceGroupCompleteResult, error) {
	return c.ListByResourceGroupCompleteMatchingPredicate(ctx, id, HostOperationPredicate{})
}

// ListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HostsClient) ListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate HostOperationPredicate) (result ListByResourceGroupCompleteResult, err error) {
	items := make([]Host, 0)

	resp, err := c.ListByResourceGroup(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByResourceGroupCompleteResult{
		Items: items,
	}
	return
}
