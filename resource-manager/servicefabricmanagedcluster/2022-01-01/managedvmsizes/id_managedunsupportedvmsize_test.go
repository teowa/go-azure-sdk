package managedvmsizes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ManagedUnsupportedVMSizeId{}

func TestNewManagedUnsupportedVMSizeID(t *testing.T) {
	id := NewManagedUnsupportedVMSizeID("12345678-1234-9876-4563-123456789012", "locationValue", "managedUnsupportedVMSizeValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.ManagedUnsupportedVMSizeName != "managedUnsupportedVMSizeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedUnsupportedVMSizeName'", id.ManagedUnsupportedVMSizeName, "managedUnsupportedVMSizeValue")
	}
}

func TestFormatManagedUnsupportedVMSizeID(t *testing.T) {
	actual := NewManagedUnsupportedVMSizeID("12345678-1234-9876-4563-123456789012", "locationValue", "managedUnsupportedVMSizeValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue/managedUnsupportedVMSizes/managedUnsupportedVMSizeValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseManagedUnsupportedVMSizeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedUnsupportedVMSizeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue/managedUnsupportedVMSizes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue/managedUnsupportedVMSizes/managedUnsupportedVMSizeValue",
			Expected: &ManagedUnsupportedVMSizeId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				LocationName:                 "locationValue",
				ManagedUnsupportedVMSizeName: "managedUnsupportedVMSizeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue/managedUnsupportedVMSizes/managedUnsupportedVMSizeValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedUnsupportedVMSizeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.ManagedUnsupportedVMSizeName != v.Expected.ManagedUnsupportedVMSizeName {
			t.Fatalf("Expected %q but got %q for ManagedUnsupportedVMSizeName", v.Expected.ManagedUnsupportedVMSizeName, actual.ManagedUnsupportedVMSizeName)
		}

	}
}

func TestParseManagedUnsupportedVMSizeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedUnsupportedVMSizeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sErViCeFaBrIc",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sErViCeFaBrIc/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sErViCeFaBrIc/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue/managedUnsupportedVMSizes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sErViCeFaBrIc/lOcAtIoNs/lOcAtIoNvAlUe/mAnAgEdUnSuPpOrTeDvMsIzEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue/managedUnsupportedVMSizes/managedUnsupportedVMSizeValue",
			Expected: &ManagedUnsupportedVMSizeId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				LocationName:                 "locationValue",
				ManagedUnsupportedVMSizeName: "managedUnsupportedVMSizeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.ServiceFabric/locations/locationValue/managedUnsupportedVMSizes/managedUnsupportedVMSizeValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sErViCeFaBrIc/lOcAtIoNs/lOcAtIoNvAlUe/mAnAgEdUnSuPpOrTeDvMsIzEs/mAnAgEdUnSuPpOrTeDvMsIzEvAlUe",
			Expected: &ManagedUnsupportedVMSizeId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				LocationName:                 "lOcAtIoNvAlUe",
				ManagedUnsupportedVMSizeName: "mAnAgEdUnSuPpOrTeDvMsIzEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sErViCeFaBrIc/lOcAtIoNs/lOcAtIoNvAlUe/mAnAgEdUnSuPpOrTeDvMsIzEs/mAnAgEdUnSuPpOrTeDvMsIzEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedUnsupportedVMSizeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.ManagedUnsupportedVMSizeName != v.Expected.ManagedUnsupportedVMSizeName {
			t.Fatalf("Expected %q but got %q for ManagedUnsupportedVMSizeName", v.Expected.ManagedUnsupportedVMSizeName, actual.ManagedUnsupportedVMSizeName)
		}

	}
}

func TestSegmentsForManagedUnsupportedVMSizeId(t *testing.T) {
	segments := ManagedUnsupportedVMSizeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ManagedUnsupportedVMSizeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
