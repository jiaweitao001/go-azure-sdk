package assets

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Asset struct {
	ExtendedLocation ExtendedLocation       `json:"extendedLocation"`
	Id               *string                `json:"id,omitempty"`
	Location         string                 `json:"location"`
	Name             *string                `json:"name,omitempty"`
	Properties       *AssetProperties       `json:"properties,omitempty"`
	SystemData       *systemdata.SystemData `json:"systemData,omitempty"`
	Tags             *map[string]string     `json:"tags,omitempty"`
	Type             *string                `json:"type,omitempty"`
}
