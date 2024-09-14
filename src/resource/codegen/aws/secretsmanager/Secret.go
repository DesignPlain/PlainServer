package secretsmanager

import types "libds/aws/types"

type Secret struct {
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `aws.secretsmanager.SecretPolicy`. To delete the `policy`, set it to `"{}"` (an empty JSON document).
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
	RecoveryWindowInDays int `json:"recoveryWindowInDays,omitempty" yaml:"recoveryWindowInDays,omitempty"`

	// Configuration block to support secret replication. See details below.
	Replicas []types.Secretsmanager_SecretReplica `json:"replicas,omitempty" yaml:"replicas,omitempty"`

	// Description of the secret.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ARN or Id of the AWS KMS key to be used to encrypt the secret values in the versions stored in this secret. If you need to reference a CMK in a different account, you can use only the key ARN. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default KMS key (the one named `aws/secretsmanager`). If the default KMS key with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of user-defined tags that are attached to the secret. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Accepts boolean value to specify whether to overwrite a secret with the same name in the destination Region.
	ForceOverwriteReplicaSecret bool `json:"forceOverwriteReplicaSecret,omitempty" yaml:"forceOverwriteReplicaSecret,omitempty"`
}
