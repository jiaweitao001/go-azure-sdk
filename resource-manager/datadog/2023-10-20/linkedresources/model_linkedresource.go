package linkedresources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedResource struct {
	Id       *string `json:"id,omitempty"`
	Location *string `json:"location,omitempty"`
}
