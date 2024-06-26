package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionPolicies struct {
	LocationPlacementId *string `json:"locationPlacementId,omitempty"`
	QuotaId             *string `json:"quotaId,omitempty"`
}
