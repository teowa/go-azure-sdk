package publishedartifact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublishedArtifactClient struct {
	Client *resourcemanager.Client
}

func NewPublishedArtifactClientWithBaseURI(api environments.Api) (*PublishedArtifactClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "publishedartifact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PublishedArtifactClient: %+v", err)
	}

	return &PublishedArtifactClient{
		Client: client,
	}, nil
}
