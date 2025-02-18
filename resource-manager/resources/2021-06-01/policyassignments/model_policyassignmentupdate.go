package policyassignments

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentUpdate struct {
	Identity *identity.SystemOrUserAssignedMap `json:"identity,omitempty"`
	Location *string                           `json:"location,omitempty"`
}
