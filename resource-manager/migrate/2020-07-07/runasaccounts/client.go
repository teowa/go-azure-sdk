package runasaccounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunAsAccountsClient struct {
	Client *resourcemanager.Client
}

func NewRunAsAccountsClientWithBaseURI(api environments.Api) (*RunAsAccountsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "runasaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RunAsAccountsClient: %+v", err)
	}

	return &RunAsAccountsClient{
		Client: client,
	}, nil
}
