package types

type Apprunner_ServiceNetworkConfigurationIngressConfiguration struct {
	// Specifies whether your App Runner service is publicly accessible. To make the service publicly accessible set it to True. To make the service privately accessible, from only within an Amazon VPC set it to False.
	IsPubliclyAccessible bool `json:"isPubliclyAccessible,omitempty" yaml:"isPubliclyAccessible,omitempty"`
}
