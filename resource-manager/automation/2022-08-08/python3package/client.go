package python3package

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Python3PackageClient struct {
	Client *resourcemanager.Client
}

func NewPython3PackageClientWithBaseURI(api environments.Api) (*Python3PackageClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "python3package", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating Python3PackageClient: %+v", err)
	}

	return &Python3PackageClient{
		Client: client,
	}, nil
}
