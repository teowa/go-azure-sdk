package expressroutecircuitauthorizations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCircuitAuthorizationsClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCircuitAuthorizationsClientWithBaseURI(api environments.Api) (*ExpressRouteCircuitAuthorizationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutecircuitauthorizations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCircuitAuthorizationsClient: %+v", err)
	}

	return &ExpressRouteCircuitAuthorizationsClient{
		Client: client,
	}, nil
}
