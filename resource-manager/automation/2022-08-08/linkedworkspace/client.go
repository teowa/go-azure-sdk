package linkedworkspace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedWorkspaceClient struct {
	Client *resourcemanager.Client
}

func NewLinkedWorkspaceClientWithBaseURI(api environments.Api) (*LinkedWorkspaceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "linkedworkspace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LinkedWorkspaceClient: %+v", err)
	}

	return &LinkedWorkspaceClient{
		Client: client,
	}, nil
}
