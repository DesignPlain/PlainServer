package types

type Integrationconnectors_ConnectionSslConfigAdditionalVariable struct {
	// Key for the configVariable
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   Secret value of configVariable
	   Structure is documented below.
	*/
	SecretValue Integrationconnectors_ConnectionSslConfigAdditionalVariableSecretValue `json:"secretValue,omitempty" yaml:"secretValue,omitempty"`

	// String Value of configVariabley.
	StringValue string `json:"stringValue,omitempty" yaml:"stringValue,omitempty"`

	// Boolean Value of configVariable.
	BooleanValue bool `json:"booleanValue,omitempty" yaml:"booleanValue,omitempty"`

	/*
	   Encription key value of configVariable.
	   Structure is documented below.
	*/
	EncryptionKeyValue Integrationconnectors_ConnectionSslConfigAdditionalVariableEncryptionKeyValue `json:"encryptionKeyValue,omitempty" yaml:"encryptionKeyValue,omitempty"`

	// Integer Value of configVariable.
	IntegerValue int `json:"integerValue,omitempty" yaml:"integerValue,omitempty"`
}
