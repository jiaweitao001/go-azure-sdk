package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type L2NetworkAttachmentConfiguration struct {
	NetworkId  string                `json:"networkId"`
	PluginType *KubernetesPluginType `json:"pluginType,omitempty"`
}
