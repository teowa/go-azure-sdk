package backupinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupDatasourceParameters interface {
}

func unmarshalBackupDatasourceParametersImplementation(input []byte) (BackupDatasourceParameters, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BackupDatasourceParameters into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "BlobBackupDatasourceParameters") {
		var out BlobBackupDatasourceParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlobBackupDatasourceParameters: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "KubernetesClusterBackupDatasourceParameters") {
		var out KubernetesClusterBackupDatasourceParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KubernetesClusterBackupDatasourceParameters: %+v", err)
		}
		return out, nil
	}

	type RawBackupDatasourceParametersImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawBackupDatasourceParametersImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
