package networkclouds

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

type TrunkedNetworksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TrunkedNetwork
}

type TrunkedNetworksListByResourceGroupCompleteResult struct {
	Items []TrunkedNetwork
}

// TrunkedNetworksListByResourceGroup ...
func (c NetworkcloudsClient) TrunkedNetworksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result TrunkedNetworksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/trunkedNetworks", id.ID()),
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
		Values *[]TrunkedNetwork `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// TrunkedNetworksListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) TrunkedNetworksListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (TrunkedNetworksListByResourceGroupCompleteResult, error) {
	return c.TrunkedNetworksListByResourceGroupCompleteMatchingPredicate(ctx, id, TrunkedNetworkOperationPredicate{})
}

// TrunkedNetworksListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) TrunkedNetworksListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate TrunkedNetworkOperationPredicate) (result TrunkedNetworksListByResourceGroupCompleteResult, err error) {
	items := make([]TrunkedNetwork, 0)

	resp, err := c.TrunkedNetworksListByResourceGroup(ctx, id)
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

	result = TrunkedNetworksListByResourceGroupCompleteResult{
		Items: items,
	}
	return
}
