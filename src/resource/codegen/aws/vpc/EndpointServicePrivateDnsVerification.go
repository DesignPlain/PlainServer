package vpc

import types "libds/aws/types"

type EndpointServicePrivateDnsVerification struct {
	/*
	   ID of the endpoint service.

	   The following arguments are optional:
	*/
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`

	//
	Timeouts types.Vpc_EndpointServicePrivateDnsVerificationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Whether to wait until the endpoint service returns a `Verified` status for the configured private DNS name.
	WaitForVerification bool `json:"waitForVerification,omitempty" yaml:"waitForVerification,omitempty"`
}
