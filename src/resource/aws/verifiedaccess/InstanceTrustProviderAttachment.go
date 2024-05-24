package verifiedaccess

type InstanceTrustProviderAttachment struct {
	// The ID of the Verified Access instance to attach the Trust Provider to.
	VerifiedaccessInstanceId string `json:"verifiedaccessInstanceId,omitempty" yaml:"verifiedaccessInstanceId,omitempty"`

	// The ID of the Verified Access trust provider.
	VerifiedaccessTrustProviderId string `json:"verifiedaccessTrustProviderId,omitempty" yaml:"verifiedaccessTrustProviderId,omitempty"`
}
