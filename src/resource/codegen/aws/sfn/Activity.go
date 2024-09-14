package sfn

import types "libds/aws/types"

type Activity struct {
	// The name of the activity to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Defines what encryption configuration is used to encrypt data in the Activity. For more information see the section [Data at rest encyption](https://docs.aws.amazon.com/step-functions/latest/dg/encryption-at-rest.html) in the AWS Step Functions User Guide.
	EncryptionConfiguration types.Sfn_ActivityEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`
}
