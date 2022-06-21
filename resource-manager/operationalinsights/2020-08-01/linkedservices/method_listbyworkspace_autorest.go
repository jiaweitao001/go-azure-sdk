package linkedservices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type ListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	Model        *LinkedServiceListResult
}

// ListByWorkspace ...
func (c LinkedServicesClient) ListByWorkspace(ctx context.Context, id WorkspaceId) (result ListByWorkspaceOperationResponse, err error) {
	req, err := c.preparerForListByWorkspace(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByWorkspace", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByWorkspace(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "linkedservices.LinkedServicesClient", "ListByWorkspace", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByWorkspace prepares the ListByWorkspace request.
func (c LinkedServicesClient) preparerForListByWorkspace(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/linkedServices", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByWorkspace handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (c LinkedServicesClient) responderForListByWorkspace(resp *http.Response) (result ListByWorkspaceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
