package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointDeploymentResourcePropertiesBasicResource struct {
	Id         *string                              `json:"id,omitempty"`
	Name       *string                              `json:"name,omitempty"`
	Properties EndpointDeploymentResourceProperties `json:"properties"`
	SystemData *systemdata.SystemData               `json:"systemData,omitempty"`
	Type       *string                              `json:"type,omitempty"`
}

var _ json.Unmarshaler = &EndpointDeploymentResourcePropertiesBasicResource{}

func (s *EndpointDeploymentResourcePropertiesBasicResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id         *string                `json:"id,omitempty"`
		Name       *string                `json:"name,omitempty"`
		SystemData *systemdata.SystemData `json:"systemData,omitempty"`
		Type       *string                `json:"type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.SystemData = decoded.SystemData
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EndpointDeploymentResourcePropertiesBasicResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalEndpointDeploymentResourcePropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'EndpointDeploymentResourcePropertiesBasicResource': %+v", err)
		}
		s.Properties = impl
	}

	return nil
}
