package filesystems

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileSystemsClient struct {
	Client *resourcemanager.Client
}

func NewFileSystemsClientWithBaseURI(api environments.Api) (*FileSystemsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "filesystems", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FileSystemsClient: %+v", err)
	}

	return &FileSystemsClient{
		Client: client,
	}, nil
}
