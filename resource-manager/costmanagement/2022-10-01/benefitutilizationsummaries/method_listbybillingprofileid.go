package benefitutilizationsummaries

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingProfileIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BenefitUtilizationSummary
}

type ListByBillingProfileIdCompleteResult struct {
	Items []BenefitUtilizationSummary
}

type ListByBillingProfileIdOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultListByBillingProfileIdOperationOptions() ListByBillingProfileIdOperationOptions {
	return ListByBillingProfileIdOperationOptions{}
}

func (o ListByBillingProfileIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingProfileIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBillingProfileIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.GrainParameter != nil {
		out.Append("grainParameter", fmt.Sprintf("%v", *o.GrainParameter))
	}
	return &out
}

// ListByBillingProfileId ...
func (c BenefitUtilizationSummariesClient) ListByBillingProfileId(ctx context.Context, id BillingProfileId, options ListByBillingProfileIdOperationOptions) (result ListByBillingProfileIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.CostManagement/benefitUtilizationSummaries", id.ID()),
		OptionsObject: options,
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
		Values *[]BenefitUtilizationSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingProfileIdComplete retrieves all the results into a single object
func (c BenefitUtilizationSummariesClient) ListByBillingProfileIdComplete(ctx context.Context, id BillingProfileId, options ListByBillingProfileIdOperationOptions) (ListByBillingProfileIdCompleteResult, error) {
	return c.ListByBillingProfileIdCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// ListByBillingProfileIdCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BenefitUtilizationSummariesClient) ListByBillingProfileIdCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options ListByBillingProfileIdOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (result ListByBillingProfileIdCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	resp, err := c.ListByBillingProfileId(ctx, id, options)
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

	result = ListByBillingProfileIdCompleteResult{
		Items: items,
	}
	return
}
