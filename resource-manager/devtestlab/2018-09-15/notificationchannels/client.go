package notificationchannels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationChannelsClient struct {
	Client *resourcemanager.Client
}

func NewNotificationChannelsClientWithBaseURI(sdkApi sdkEnv.Api) (*NotificationChannelsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "notificationchannels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NotificationChannelsClient: %+v", err)
	}

	return &NotificationChannelsClient{
		Client: client,
	}, nil
}
