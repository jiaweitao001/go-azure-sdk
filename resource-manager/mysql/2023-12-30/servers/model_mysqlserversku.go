package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MySQLServerSku struct {
	Name string        `json:"name"`
	Tier ServerSkuTier `json:"tier"`
}
