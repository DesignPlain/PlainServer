package types

type Fms_PolicySecurityServicePolicyData struct {
	// The service that the policy is using to protect the resources. For the current list of supported types, please refer to the [AWS Firewall Manager SecurityServicePolicyData API Type Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html#fms-Type-SecurityServicePolicyData-Type).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Details about the service that are specific to the service type, in JSON format. For service type `SHIELD_ADVANCED`, this is an empty string. Examples depending on `type` can be found in the [AWS Firewall Manager SecurityServicePolicyData API Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_SecurityServicePolicyData.html).
	ManagedServiceData string `json:"managedServiceData,omitempty" yaml:"managedServiceData,omitempty"`

	// Contains the Network Firewall firewall policy options to configure a centralized deployment model. Documented below.
	PolicyOption Fms_PolicySecurityServicePolicyDataPolicyOption `json:"policyOption,omitempty" yaml:"policyOption,omitempty"`
}
