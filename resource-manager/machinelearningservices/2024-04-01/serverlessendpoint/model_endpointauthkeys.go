package serverlessendpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointAuthKeys struct {
	PrimaryKey   *string `json:"primaryKey,omitempty"`
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}
