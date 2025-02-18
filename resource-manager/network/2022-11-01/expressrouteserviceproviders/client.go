package expressrouteserviceproviders

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteServiceProvidersClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteServiceProvidersClientWithBaseURI(api environments.Api) (*ExpressRouteServiceProvidersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressrouteserviceproviders", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteServiceProvidersClient: %+v", err)
	}

	return &ExpressRouteServiceProvidersClient{
		Client: client,
	}, nil
}
