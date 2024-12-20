package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = MicrosoftPurviewInformationProtectionCheckRequirements{}

type MicrosoftPurviewInformationProtectionCheckRequirements struct {
	Properties *DataConnectorTenantId `json:"properties,omitempty"`

	// Fields inherited from DataConnectorsCheckRequirements

	Kind DataConnectorKind `json:"kind"`
}

func (s MicrosoftPurviewInformationProtectionCheckRequirements) DataConnectorsCheckRequirements() BaseDataConnectorsCheckRequirementsImpl {
	return BaseDataConnectorsCheckRequirementsImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = MicrosoftPurviewInformationProtectionCheckRequirements{}

func (s MicrosoftPurviewInformationProtectionCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftPurviewInformationProtectionCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftPurviewInformationProtectionCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftPurviewInformationProtectionCheckRequirements: %+v", err)
	}

	decoded["kind"] = "MicrosoftPurviewInformationProtection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftPurviewInformationProtectionCheckRequirements: %+v", err)
	}

	return encoded, nil
}
