package sessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceInfoPatchProperties struct {
	ExactVersion *string `json:"exactVersion,omitempty"`
	Offer        *string `json:"offer,omitempty"`
	Publisher    *string `json:"publisher,omitempty"`
	Sku          *string `json:"sku,omitempty"`
}
