package worklink

import types "libds/aws/types"

type Fleet struct {
	// A region-unique name for the AMI.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network types.Worklink_FleetNetwork `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.

	   --network-- requires the following:

	   > --NOTE:-- `network` is cannot removed without force recreating.
	*/
	OptimizeForEndUserLocation bool `json:"optimizeForEndUserLocation,omitempty" yaml:"optimizeForEndUserLocation,omitempty"`

	// The ARN of the Amazon Kinesis data stream that receives the audit events. Kinesis data stream name must begin with `"AmazonWorkLink-"`.
	AuditStreamArn string `json:"auditStreamArn,omitempty" yaml:"auditStreamArn,omitempty"`

	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate string `json:"deviceCaCertificate,omitempty" yaml:"deviceCaCertificate,omitempty"`

	// The name of the fleet.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider types.Worklink_FleetIdentityProvider `json:"identityProvider,omitempty" yaml:"identityProvider,omitempty"`
}
