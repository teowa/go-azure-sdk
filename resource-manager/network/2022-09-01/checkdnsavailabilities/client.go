package checkdnsavailabilities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckDnsAvailabilitiesClient struct {
	Client *resourcemanager.Client
}

func NewCheckDnsAvailabilitiesClientWithBaseURI(api environments.Api) (*CheckDnsAvailabilitiesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "checkdnsavailabilities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CheckDnsAvailabilitiesClient: %+v", err)
	}

	return &CheckDnsAvailabilitiesClient{
		Client: client,
	}, nil
}
