package appservicecertificateorders

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateDetails struct {
	Issuer             *string `json:"issuer,omitempty"`
	NotAfter           *string `json:"notAfter,omitempty"`
	NotBefore          *string `json:"notBefore,omitempty"`
	RawData            *string `json:"rawData,omitempty"`
	SerialNumber       *string `json:"serialNumber,omitempty"`
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	Subject            *string `json:"subject,omitempty"`
	Thumbprint         *string `json:"thumbprint,omitempty"`
	Version            *int64  `json:"version,omitempty"`
}

func (o *CertificateDetails) GetNotAfterAsTime() (*time.Time, error) {
	if o.NotAfter == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NotAfter, "2006-01-02T15:04:05Z07:00")
}

func (o *CertificateDetails) SetNotAfterAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NotAfter = &formatted
}

func (o *CertificateDetails) GetNotBeforeAsTime() (*time.Time, error) {
	if o.NotBefore == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NotBefore, "2006-01-02T15:04:05Z07:00")
}

func (o *CertificateDetails) SetNotBeforeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NotBefore = &formatted
}
