package datazone

import types "libds/aws/types"

type Domain struct {
	// Description of the Domain.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   ARN of the role used by DataZone to configure the Domain.

	   The following arguments are optional:
	*/
	DomainExecutionRole string `json:"domainExecutionRole,omitempty" yaml:"domainExecutionRole,omitempty"`

	// ARN of the KMS key used to encrypt the Amazon DataZone domain, metadata and reporting data.
	KmsKeyIdentifier string `json:"kmsKeyIdentifier,omitempty" yaml:"kmsKeyIdentifier,omitempty"`

	// Name of the Domain.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Single sign on options, used to [enable AWS IAM Identity Center](https://docs.aws.amazon.com/datazone/latest/userguide/enable-IAM-identity-center-for-datazone.html) for DataZone.
	SingleSignOn types.Datazone_DomainSingleSignOn `json:"singleSignOn,omitempty" yaml:"singleSignOn,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Datazone_DomainTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
