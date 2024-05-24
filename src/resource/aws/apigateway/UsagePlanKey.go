package apigateway

type UsagePlanKey struct {
	// Type of the API key resource. Currently, the valid key type is API_KEY.
	KeyType string `json:"keyType,omitempty" yaml:"keyType,omitempty"`

	// Id of the usage plan resource representing to associate the key to.
	UsagePlanId string `json:"usagePlanId,omitempty" yaml:"usagePlanId,omitempty"`

	// Identifier of the API key resource.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`
}
