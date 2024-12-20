package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigServerResourceRequests struct {
	Cpu           *string `json:"cpu,omitempty"`
	InstanceCount *int64  `json:"instanceCount,omitempty"`
	Memory        *string `json:"memory,omitempty"`
}
