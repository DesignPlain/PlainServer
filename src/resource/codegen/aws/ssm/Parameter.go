package ssm

type Parameter struct {
	// Regular expression used to validate the parameter value.
	AllowedPattern string `json:"allowedPattern,omitempty" yaml:"allowedPattern,omitempty"`

	// ARN of the parameter.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	/*
	   Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
	DataType string `json:"dataType,omitempty" yaml:"dataType,omitempty"`

	// Description of the parameter.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Value of the parameter. --Use caution:-- This value is _never_ marked as sensitive in the pulumi preview output. This argument is not valid with a `type` of `SecureString`.
	InsecureValue string `json:"insecureValue,omitempty" yaml:"insecureValue,omitempty"`

	// KMS key ID or ARN for encrypting a SecureString.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
	Overwrite bool `json:"overwrite,omitempty" yaml:"overwrite,omitempty"`

	/*
	   Value of the parameter. This value is always marked as sensitive in the pulumi preview output, regardless of `type`.

	   > --NOTE:-- `aws:ssm:integration` data_type parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
