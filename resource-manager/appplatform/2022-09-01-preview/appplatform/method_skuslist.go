package appplatform

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

type SkusListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceSku
}

type SkusListCompleteResult struct {
	Items []ResourceSku
}

// SkusList ...
func (c AppPlatformClient) SkusList(ctx context.Context, id commonids.SubscriptionId) (result SkusListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.AppPlatform/skus", id.ID()),
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
		Values *[]ResourceSku `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SkusListComplete retrieves all the results into a single object
func (c AppPlatformClient) SkusListComplete(ctx context.Context, id commonids.SubscriptionId) (SkusListCompleteResult, error) {
	return c.SkusListCompleteMatchingPredicate(ctx, id, ResourceSkuOperationPredicate{})
}

// SkusListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppPlatformClient) SkusListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ResourceSkuOperationPredicate) (result SkusListCompleteResult, err error) {
	items := make([]ResourceSku, 0)

	resp, err := c.SkusList(ctx, id)
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

	result = SkusListCompleteResult{
		Items: items,
	}
	return
}
