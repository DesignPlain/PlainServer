package types

type Iam_WorkforcePoolProviderOidcClientSecretValue struct {
	/*
	   The plain text of the client secret value.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	PlainText string `json:"plainText,omitempty" yaml:"plainText,omitempty"`

	/*
	   (Output)
	   A thumbprint to represent the current client secret value.
	*/
	Thumbprint string `json:"thumbprint,omitempty" yaml:"thumbprint,omitempty"`
}
