package datastore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Datastore = HdfsDatastore{}

type HdfsDatastore struct {
	HdfsServerCertificate *string `json:"hdfsServerCertificate,omitempty"`
	NameNodeAddress       string  `json:"nameNodeAddress"`
	Protocol              *string `json:"protocol,omitempty"`

	// Fields inherited from Datastore
	Credentials          DatastoreCredentials  `json:"credentials"`
	Description          *string               `json:"description,omitempty"`
	IntellectualProperty *IntellectualProperty `json:"intellectualProperty,omitempty"`
	IsDefault            *bool                 `json:"isDefault,omitempty"`
	Properties           *map[string]string    `json:"properties,omitempty"`
	Tags                 *map[string]string    `json:"tags,omitempty"`
}

var _ json.Marshaler = HdfsDatastore{}

func (s HdfsDatastore) MarshalJSON() ([]byte, error) {
	type wrapper HdfsDatastore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HdfsDatastore: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HdfsDatastore: %+v", err)
	}
	decoded["datastoreType"] = "Hdfs"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HdfsDatastore: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &HdfsDatastore{}

func (s *HdfsDatastore) UnmarshalJSON(bytes []byte) error {
	type alias HdfsDatastore
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into HdfsDatastore: %+v", err)
	}

	s.Description = decoded.Description
	s.HdfsServerCertificate = decoded.HdfsServerCertificate
	s.IntellectualProperty = decoded.IntellectualProperty
	s.IsDefault = decoded.IsDefault
	s.NameNodeAddress = decoded.NameNodeAddress
	s.Properties = decoded.Properties
	s.Protocol = decoded.Protocol
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling HdfsDatastore into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["credentials"]; ok {
		impl, err := unmarshalDatastoreCredentialsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Credentials' for 'HdfsDatastore': %+v", err)
		}
		s.Credentials = impl
	}
	return nil
}