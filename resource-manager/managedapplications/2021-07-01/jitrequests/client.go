package jitrequests

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitRequestsClient struct {
	Client *resourcemanager.Client
}

func NewJitRequestsClientWithBaseURI(api environments.Api) (*JitRequestsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "jitrequests", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JitRequestsClient: %+v", err)
	}

	return &JitRequestsClient{
		Client: client,
	}, nil
}
