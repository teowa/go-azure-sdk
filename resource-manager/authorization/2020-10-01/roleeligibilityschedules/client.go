package roleeligibilityschedules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilitySchedulesClient struct {
	Client *resourcemanager.Client
}

func NewRoleEligibilitySchedulesClientWithBaseURI(api environments.Api) (*RoleEligibilitySchedulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "roleeligibilityschedules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleEligibilitySchedulesClient: %+v", err)
	}

	return &RoleEligibilitySchedulesClient{
		Client: client,
	}, nil
}
