package pipelines

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelinesClient struct {
	Client *resourcemanager.Client
}

func NewPipelinesClientWithBaseURI(api environments.Api) (*PipelinesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "pipelines", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PipelinesClient: %+v", err)
	}

	return &PipelinesClient{
		Client: client,
	}, nil
}
