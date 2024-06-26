package managedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler struct {
	AddonAutoscaling *AddonAutoscaling `json:"addonAutoscaling,omitempty"`
	Enabled          bool              `json:"enabled"`
}
