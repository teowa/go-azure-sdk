package deploymentoperations

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

type ListAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DeploymentOperation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAtSubscriptionScopeOperationResponse, error)
}

type ListAtSubscriptionScopeCompleteResult struct {
	Items []DeploymentOperation
}

func (r ListAtSubscriptionScopeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAtSubscriptionScopeOperationResponse) LoadMore(ctx context.Context) (resp ListAtSubscriptionScopeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListAtSubscriptionScopeOperationOptions struct {
	Top *int64
}

func DefaultListAtSubscriptionScopeOperationOptions() ListAtSubscriptionScopeOperationOptions {
	return ListAtSubscriptionScopeOperationOptions{}
}

func (o ListAtSubscriptionScopeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListAtSubscriptionScopeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListAtSubscriptionScope ...
func (c DeploymentOperationsClient) ListAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId, options ListAtSubscriptionScopeOperationOptions) (resp ListAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForListAtSubscriptionScope(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtSubscriptionScope", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAtSubscriptionScope(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtSubscriptionScope", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListAtSubscriptionScope prepares the ListAtSubscriptionScope request.
func (c DeploymentOperationsClient) preparerForListAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId, options ListAtSubscriptionScopeOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/operations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListAtSubscriptionScopeWithNextLink prepares the ListAtSubscriptionScope request with the given nextLink token.
func (c DeploymentOperationsClient) preparerForListAtSubscriptionScopeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListAtSubscriptionScope handles the response to the ListAtSubscriptionScope request. The method always
// closes the http.Response Body.
func (c DeploymentOperationsClient) responderForListAtSubscriptionScope(resp *http.Response) (result ListAtSubscriptionScopeOperationResponse, err error) {
	type page struct {
		Values   []DeploymentOperation `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAtSubscriptionScopeOperationResponse, err error) {
			req, err := c.preparerForListAtSubscriptionScopeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtSubscriptionScope", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtSubscriptionScope", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAtSubscriptionScope(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtSubscriptionScope", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListAtSubscriptionScopeComplete retrieves all of the results into a single object
func (c DeploymentOperationsClient) ListAtSubscriptionScopeComplete(ctx context.Context, id ProviderDeploymentId, options ListAtSubscriptionScopeOperationOptions) (ListAtSubscriptionScopeCompleteResult, error) {
	return c.ListAtSubscriptionScopeCompleteMatchingPredicate(ctx, id, options, DeploymentOperationOperationPredicate{})
}

// ListAtSubscriptionScopeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DeploymentOperationsClient) ListAtSubscriptionScopeCompleteMatchingPredicate(ctx context.Context, id ProviderDeploymentId, options ListAtSubscriptionScopeOperationOptions, predicate DeploymentOperationOperationPredicate) (resp ListAtSubscriptionScopeCompleteResult, err error) {
	items := make([]DeploymentOperation, 0)

	page, err := c.ListAtSubscriptionScope(ctx, id, options)
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

	out := ListAtSubscriptionScopeCompleteResult{
		Items: items,
	}
	return out, nil
}
