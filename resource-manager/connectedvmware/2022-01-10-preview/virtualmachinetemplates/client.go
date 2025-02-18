package virtualmachinetemplates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineTemplatesClient struct {
	Client *resourcemanager.Client
}

func NewVirtualMachineTemplatesClientWithBaseURI(api environments.Api) (*VirtualMachineTemplatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "virtualmachinetemplates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualMachineTemplatesClient: %+v", err)
	}

	return &VirtualMachineTemplatesClient{
		Client: client,
	}, nil
}
