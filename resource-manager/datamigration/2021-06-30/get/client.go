package get

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GETClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGETClientWithBaseURI(endpoint string) GETClient {
	return GETClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
