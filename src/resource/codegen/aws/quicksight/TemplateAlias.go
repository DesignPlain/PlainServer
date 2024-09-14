package quicksight

type TemplateAlias struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// ID of the template.
	TemplateId string `json:"templateId,omitempty" yaml:"templateId,omitempty"`

	/*
	   Version number of the template.

	   The following arguments are optional:
	*/
	TemplateVersionNumber int `json:"templateVersionNumber,omitempty" yaml:"templateVersionNumber,omitempty"`

	// Display name of the template alias.
	AliasName string `json:"aliasName,omitempty" yaml:"aliasName,omitempty"`
}
