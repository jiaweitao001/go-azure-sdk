package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterFeaturesListByKubernetesClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]KubernetesClusterFeature
}

type KubernetesClusterFeaturesListByKubernetesClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []KubernetesClusterFeature
}

type KubernetesClusterFeaturesListByKubernetesClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *KubernetesClusterFeaturesListByKubernetesClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// KubernetesClusterFeaturesListByKubernetesCluster ...
func (c NetworkcloudsClient) KubernetesClusterFeaturesListByKubernetesCluster(ctx context.Context, id KubernetesClusterId) (result KubernetesClusterFeaturesListByKubernetesClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &KubernetesClusterFeaturesListByKubernetesClusterCustomPager{},
		Path:       fmt.Sprintf("%s/features", id.ID()),
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
		Values *[]KubernetesClusterFeature `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// KubernetesClusterFeaturesListByKubernetesClusterComplete retrieves all the results into a single object
func (c NetworkcloudsClient) KubernetesClusterFeaturesListByKubernetesClusterComplete(ctx context.Context, id KubernetesClusterId) (KubernetesClusterFeaturesListByKubernetesClusterCompleteResult, error) {
	return c.KubernetesClusterFeaturesListByKubernetesClusterCompleteMatchingPredicate(ctx, id, KubernetesClusterFeatureOperationPredicate{})
}

// KubernetesClusterFeaturesListByKubernetesClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) KubernetesClusterFeaturesListByKubernetesClusterCompleteMatchingPredicate(ctx context.Context, id KubernetesClusterId, predicate KubernetesClusterFeatureOperationPredicate) (result KubernetesClusterFeaturesListByKubernetesClusterCompleteResult, err error) {
	items := make([]KubernetesClusterFeature, 0)

	resp, err := c.KubernetesClusterFeaturesListByKubernetesCluster(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = KubernetesClusterFeaturesListByKubernetesClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
