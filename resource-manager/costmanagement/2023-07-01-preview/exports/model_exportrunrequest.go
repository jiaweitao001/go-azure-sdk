package exports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportRunRequest struct {
	TimePeriod *ExportTimePeriod `json:"timePeriod,omitempty"`
}
