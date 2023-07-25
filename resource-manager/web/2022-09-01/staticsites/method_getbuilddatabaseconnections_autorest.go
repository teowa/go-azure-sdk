package staticsites

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

type GetBuildDatabaseConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DatabaseConnection

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetBuildDatabaseConnectionsOperationResponse, error)
}

type GetBuildDatabaseConnectionsCompleteResult struct {
	Items []DatabaseConnection
}

func (r GetBuildDatabaseConnectionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetBuildDatabaseConnectionsOperationResponse) LoadMore(ctx context.Context) (resp GetBuildDatabaseConnectionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetBuildDatabaseConnections ...
func (c StaticSitesClient) GetBuildDatabaseConnections(ctx context.Context, id BuildId) (resp GetBuildDatabaseConnectionsOperationResponse, err error) {
	req, err := c.preparerForGetBuildDatabaseConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnections", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnections", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetBuildDatabaseConnections(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnections", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetBuildDatabaseConnections prepares the GetBuildDatabaseConnections request.
func (c StaticSitesClient) preparerForGetBuildDatabaseConnections(ctx context.Context, id BuildId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/databaseConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetBuildDatabaseConnectionsWithNextLink prepares the GetBuildDatabaseConnections request with the given nextLink token.
func (c StaticSitesClient) preparerForGetBuildDatabaseConnectionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetBuildDatabaseConnections handles the response to the GetBuildDatabaseConnections request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetBuildDatabaseConnections(resp *http.Response) (result GetBuildDatabaseConnectionsOperationResponse, err error) {
	type page struct {
		Values   []DatabaseConnection `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetBuildDatabaseConnectionsOperationResponse, err error) {
			req, err := c.preparerForGetBuildDatabaseConnectionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnections", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnections", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetBuildDatabaseConnections(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetBuildDatabaseConnections", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetBuildDatabaseConnectionsComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetBuildDatabaseConnectionsComplete(ctx context.Context, id BuildId) (GetBuildDatabaseConnectionsCompleteResult, error) {
	return c.GetBuildDatabaseConnectionsCompleteMatchingPredicate(ctx, id, DatabaseConnectionOperationPredicate{})
}

// GetBuildDatabaseConnectionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetBuildDatabaseConnectionsCompleteMatchingPredicate(ctx context.Context, id BuildId, predicate DatabaseConnectionOperationPredicate) (resp GetBuildDatabaseConnectionsCompleteResult, err error) {
	items := make([]DatabaseConnection, 0)

	page, err := c.GetBuildDatabaseConnections(ctx, id)
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

	out := GetBuildDatabaseConnectionsCompleteResult{
		Items: items,
	}
	return out, nil
}
