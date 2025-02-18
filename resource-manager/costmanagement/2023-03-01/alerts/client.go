package alerts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertsClient struct {
	Client *resourcemanager.Client
}

func NewAlertsClientWithBaseURI(api environments.Api) (*AlertsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "alerts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AlertsClient: %+v", err)
	}

	return &AlertsClient{
		Client: client,
	}, nil
}
