package redshiftserverless

type Namespace struct {
	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying `default_iam_role_arn`, it also must be part of `iam_roles`.
	DefaultIamRoleArn string `json:"defaultIamRoleArn,omitempty" yaml:"defaultIamRoleArn,omitempty"`

	// A list of IAM roles to associate with the namespace.
	IamRoles []string `json:"iamRoles,omitempty" yaml:"iamRoles,omitempty"`

	// The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The name of the namespace.
	NamespaceName string `json:"namespaceName,omitempty" yaml:"namespaceName,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ID of the KMS key used to encrypt the namespace's admin credentials secret.
	AdminPasswordSecretKmsKeyId string `json:"adminPasswordSecretKmsKeyId,omitempty" yaml:"adminPasswordSecretKmsKeyId,omitempty"`

	/*
	   The password of the administrator for the first database created in the namespace.
	   Conflicts with `manage_admin_password`.
	*/
	AdminUserPassword string `json:"adminUserPassword,omitempty" yaml:"adminUserPassword,omitempty"`

	// The name of the first database created in the namespace.
	DbName string `json:"dbName,omitempty" yaml:"dbName,omitempty"`

	// The username of the administrator for the first database created in the namespace.
	AdminUsername string `json:"adminUsername,omitempty" yaml:"adminUsername,omitempty"`

	// The types of logs the namespace can export. Available export types are `userlog`, `connectionlog`, and `useractivitylog`.
	LogExports []string `json:"logExports,omitempty" yaml:"logExports,omitempty"`

	/*
	   Whether to use AWS SecretManager to manage namespace's admin credentials.
	   Conflicts with `admin_user_password`.
	*/
	ManageAdminPassword bool `json:"manageAdminPassword,omitempty" yaml:"manageAdminPassword,omitempty"`
}
