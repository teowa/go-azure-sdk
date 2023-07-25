package diagnostics

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

type ListHostingEnvironmentDetectorResponsesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DetectorResponse

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListHostingEnvironmentDetectorResponsesOperationResponse, error)
}

type ListHostingEnvironmentDetectorResponsesCompleteResult struct {
	Items []DetectorResponse
}

func (r ListHostingEnvironmentDetectorResponsesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListHostingEnvironmentDetectorResponsesOperationResponse) LoadMore(ctx context.Context) (resp ListHostingEnvironmentDetectorResponsesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListHostingEnvironmentDetectorResponses ...
func (c DiagnosticsClient) ListHostingEnvironmentDetectorResponses(ctx context.Context, id HostingEnvironmentId) (resp ListHostingEnvironmentDetectorResponsesOperationResponse, err error) {
	req, err := c.preparerForListHostingEnvironmentDetectorResponses(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListHostingEnvironmentDetectorResponses", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListHostingEnvironmentDetectorResponses", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListHostingEnvironmentDetectorResponses(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListHostingEnvironmentDetectorResponses", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListHostingEnvironmentDetectorResponses prepares the ListHostingEnvironmentDetectorResponses request.
func (c DiagnosticsClient) preparerForListHostingEnvironmentDetectorResponses(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detectors", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListHostingEnvironmentDetectorResponsesWithNextLink prepares the ListHostingEnvironmentDetectorResponses request with the given nextLink token.
func (c DiagnosticsClient) preparerForListHostingEnvironmentDetectorResponsesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListHostingEnvironmentDetectorResponses handles the response to the ListHostingEnvironmentDetectorResponses request. The method always
// closes the http.Response Body.
func (c DiagnosticsClient) responderForListHostingEnvironmentDetectorResponses(resp *http.Response) (result ListHostingEnvironmentDetectorResponsesOperationResponse, err error) {
	type page struct {
		Values   []DetectorResponse `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListHostingEnvironmentDetectorResponsesOperationResponse, err error) {
			req, err := c.preparerForListHostingEnvironmentDetectorResponsesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListHostingEnvironmentDetectorResponses", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListHostingEnvironmentDetectorResponses", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListHostingEnvironmentDetectorResponses(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "diagnostics.DiagnosticsClient", "ListHostingEnvironmentDetectorResponses", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListHostingEnvironmentDetectorResponsesComplete retrieves all of the results into a single object
func (c DiagnosticsClient) ListHostingEnvironmentDetectorResponsesComplete(ctx context.Context, id HostingEnvironmentId) (ListHostingEnvironmentDetectorResponsesCompleteResult, error) {
	return c.ListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// ListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DiagnosticsClient) ListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate DetectorResponseOperationPredicate) (resp ListHostingEnvironmentDetectorResponsesCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	page, err := c.ListHostingEnvironmentDetectorResponses(ctx, id)
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

	out := ListHostingEnvironmentDetectorResponsesCompleteResult{
		Items: items,
	}
	return out, nil
}
