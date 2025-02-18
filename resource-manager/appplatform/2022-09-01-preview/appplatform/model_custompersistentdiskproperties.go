package appplatform

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomPersistentDiskProperties interface {
}

func unmarshalCustomPersistentDiskPropertiesImplementation(input []byte) (CustomPersistentDiskProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomPersistentDiskProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureFileVolume") {
		var out AzureFileVolume
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureFileVolume: %+v", err)
		}
		return out, nil
	}

	type RawCustomPersistentDiskPropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawCustomPersistentDiskPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
