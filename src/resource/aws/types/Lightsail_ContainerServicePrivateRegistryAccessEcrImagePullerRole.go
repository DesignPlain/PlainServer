package types

type Lightsail_ContainerServicePrivateRegistryAccessEcrImagePullerRole struct {
	// A Boolean value that indicates whether to activate the role. The default is `false`.
	IsActive bool `json:"isActive,omitempty" yaml:"isActive,omitempty"`

	/*
	   The principal ARN of the container service. The principal ARN can be used to create a trust
	   relationship between your standard AWS account and your Lightsail container service. This allows you to give your
	   service permission to access resources in your standard AWS account.
	*/
	PrincipalArn string `json:"principalArn,omitempty" yaml:"principalArn,omitempty"`
}
