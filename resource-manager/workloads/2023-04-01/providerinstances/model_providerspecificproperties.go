package providerinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderSpecificProperties interface {
}

func unmarshalProviderSpecificPropertiesImplementation(input []byte) (ProviderSpecificProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProviderSpecificProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["providerType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Db2") {
		var out DB2ProviderInstanceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DB2ProviderInstanceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SapHana") {
		var out HanaDbProviderInstanceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HanaDbProviderInstanceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MsSqlServer") {
		var out MsSqlServerProviderInstanceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MsSqlServerProviderInstanceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PrometheusHaCluster") {
		var out PrometheusHaClusterProviderInstanceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrometheusHaClusterProviderInstanceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PrometheusOS") {
		var out PrometheusOSProviderInstanceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrometheusOSProviderInstanceProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SapNetWeaver") {
		var out SapNetWeaverProviderInstanceProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SapNetWeaverProviderInstanceProperties: %+v", err)
		}
		return out, nil
	}

	type RawProviderSpecificPropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawProviderSpecificPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
