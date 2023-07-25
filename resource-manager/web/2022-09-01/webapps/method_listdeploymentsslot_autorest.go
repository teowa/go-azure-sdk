package webapps

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeploymentsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Deployment

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListDeploymentsSlotOperationResponse, error)
}

type ListDeploymentsSlotCompleteResult struct {
	Items []Deployment
}

func (r ListDeploymentsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListDeploymentsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListDeploymentsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListDeploymentsSlot ...
func (c WebAppsClient) ListDeploymentsSlot(ctx context.Context, id SlotId) (resp ListDeploymentsSlotOperationResponse, err error) {
	req, err := c.preparerForListDeploymentsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListDeploymentsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListDeploymentsSlot prepares the ListDeploymentsSlot request.
func (c WebAppsClient) preparerForListDeploymentsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deployments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListDeploymentsSlotWithNextLink prepares the ListDeploymentsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListDeploymentsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListDeploymentsSlot handles the response to the ListDeploymentsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListDeploymentsSlot(resp *http.Response) (result ListDeploymentsSlotOperationResponse, err error) {
	type page struct {
		Values   []Deployment `json:"value"`
		NextLink *string      `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListDeploymentsSlotOperationResponse, err error) {
			req, err := c.preparerForListDeploymentsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListDeploymentsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListDeploymentsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListDeploymentsSlotComplete(ctx context.Context, id SlotId) (ListDeploymentsSlotCompleteResult, error) {
	return c.ListDeploymentsSlotCompleteMatchingPredicate(ctx, id, DeploymentOperationPredicate{})
}

// ListDeploymentsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListDeploymentsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate DeploymentOperationPredicate) (resp ListDeploymentsSlotCompleteResult, err error) {
	items := make([]Deployment, 0)

	page, err := c.ListDeploymentsSlot(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListDeploymentsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
