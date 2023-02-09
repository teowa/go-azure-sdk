package datasetmapping

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = ADLSGen2FolderDataSetMapping{}

type ADLSGen2FolderDataSetMapping struct {
	Properties ADLSGen2FolderDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = ADLSGen2FolderDataSetMapping{}

func (s ADLSGen2FolderDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper ADLSGen2FolderDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ADLSGen2FolderDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ADLSGen2FolderDataSetMapping: %+v", err)
	}
	decoded["kind"] = "AdlsGen2Folder"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ADLSGen2FolderDataSetMapping: %+v", err)
	}

	return encoded, nil
}
