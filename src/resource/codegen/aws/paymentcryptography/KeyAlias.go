package paymentcryptography

type KeyAlias struct {
	/*
	   Name of the Key Alias.

	   The following arguments are optional:
	*/
	AliasName string `json:"aliasName,omitempty" yaml:"aliasName,omitempty"`

	// ARN of the key.
	KeyArn string `json:"keyArn,omitempty" yaml:"keyArn,omitempty"`
}
