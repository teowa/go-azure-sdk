package deploymentoperations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = OperationId{}

func TestNewOperationID(t *testing.T) {
	id := NewOperationID("deploymentValue", "operationIdValue")

	if id.DeploymentName != "deploymentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeploymentName'", id.DeploymentName, "deploymentValue")
	}

	if id.OperationId != "operationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationId'", id.OperationId, "operationIdValue")
	}
}

func TestFormatOperationID(t *testing.T) {
	actual := NewOperationID("deploymentValue", "operationIdValue").ID()
	expected := "/providers/Microsoft.Resources/deployments/deploymentValue/operations/operationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OperationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/deployments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue/operations/operationIdValue",
			Expected: &OperationId{
				DeploymentName: "deploymentValue",
				OperationId:    "operationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue/operations/operationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeploymentName != v.Expected.DeploymentName {
			t.Fatalf("Expected %q but got %q for DeploymentName", v.Expected.DeploymentName, actual.DeploymentName)
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestParseOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *OperationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/deployments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue/operations/operationIdValue",
			Expected: &OperationId{
				DeploymentName: "deploymentValue",
				OperationId:    "operationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Resources/deployments/deploymentValue/operations/operationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe/oPeRaTiOnS/oPeRaTiOnIdVaLuE",
			Expected: &OperationId{
				DeploymentName: "dEpLoYmEnTvAlUe",
				OperationId:    "oPeRaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.rEsOuRcEs/dEpLoYmEnTs/dEpLoYmEnTvAlUe/oPeRaTiOnS/oPeRaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeploymentName != v.Expected.DeploymentName {
			t.Fatalf("Expected %q but got %q for DeploymentName", v.Expected.DeploymentName, actual.DeploymentName)
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestSegmentsForOperationId(t *testing.T) {
	segments := OperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("OperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
