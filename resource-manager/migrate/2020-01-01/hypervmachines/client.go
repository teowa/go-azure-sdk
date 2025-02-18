package hypervmachines

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVMachinesClient struct {
	Client *resourcemanager.Client
}

func NewHyperVMachinesClientWithBaseURI(api environments.Api) (*HyperVMachinesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "hypervmachines", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HyperVMachinesClient: %+v", err)
	}

	return &HyperVMachinesClient{
		Client: client,
	}, nil
}
