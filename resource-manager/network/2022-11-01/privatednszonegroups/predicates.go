package privatednszonegroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateDnsZoneGroupOperationPredicate struct {
	Etag *string
	Id   *string
	Name *string
}

func (p PrivateDnsZoneGroupOperationPredicate) Matches(input PrivateDnsZoneGroup) bool {

	if p.Etag != nil && (input.Etag == nil && *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
