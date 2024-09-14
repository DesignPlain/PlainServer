package customerprofiles

import types "libds/aws/types"

type Domain struct {
	// The default encryption key, which is an AWS managed key, is used when no specific type of encryption key is specified. It is used to encrypt all data before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey string `json:"defaultEncryptionKey,omitempty" yaml:"defaultEncryptionKey,omitempty"`

	/*
	   The default number of days until the data within the domain expires.

	   The following arguments are optional:
	*/
	DefaultExpirationDays int `json:"defaultExpirationDays,omitempty" yaml:"defaultExpirationDays,omitempty"`

	// The name for your Customer Profile domain. It must be unique for your AWS account.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// A block that specifies the process of matching duplicate profiles. Documented below.
	Matching types.Customerprofiles_DomainMatching `json:"matching,omitempty" yaml:"matching,omitempty"`

	// A block that specifies the process of matching duplicate profiles using the Rule-Based matching. Documented below.
	RuleBasedMatching types.Customerprofiles_DomainRuleBasedMatching `json:"ruleBasedMatching,omitempty" yaml:"ruleBasedMatching,omitempty"`

	// Tags to apply to the domain. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The URL of the SQS dead letter queue, which is used for reporting errors associated with ingesting data from third party applications.
	DeadLetterQueueUrl string `json:"deadLetterQueueUrl,omitempty" yaml:"deadLetterQueueUrl,omitempty"`
}
