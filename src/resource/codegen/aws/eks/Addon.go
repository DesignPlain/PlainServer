package eks

type Addon struct {
	// custom configuration values for addons with single JSON string. This JSON string value must match the JSON schema derived from [describe-addon-configuration](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-configuration.html).
	ConfigurationValues string `json:"configurationValues,omitempty" yaml:"configurationValues,omitempty"`

	// Indicates if you want to preserve the created resources when deleting the EKS add-on.
	Preserve bool `json:"preserve,omitempty" yaml:"preserve,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// How to resolve field value conflicts for an Amazon EKS add-on if you've changed a value from the Amazon EKS default value. Valid values are `NONE`, `OVERWRITE`, and `PRESERVE`. For more details see the [UpdateAddon](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html) API Docs.
	ResolveConflictsOnUpdate string `json:"resolveConflictsOnUpdate,omitempty" yaml:"resolveConflictsOnUpdate,omitempty"`

	/*
	   The Amazon Resource Name (ARN) of an
	   existing IAM role to bind to the add-on's service account. The role must be
	   assigned the IAM permissions required by the add-on. If you don't specify
	   an existing IAM role, then the add-on uses the permissions assigned to the node
	   IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	   in the Amazon EKS User Guide.

	   > --Note:-- To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC)
	   provider created for your cluster. For more information, [see Enabling IAM roles
	   for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
	   in the Amazon EKS User Guide.
	*/
	ServiceAccountRoleArn string `json:"serviceAccountRoleArn,omitempty" yaml:"serviceAccountRoleArn,omitempty"`

	/*
	   Name of the EKS add-on. The name must match one of
	   the names returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
	*/
	AddonName string `json:"addonName,omitempty" yaml:"addonName,omitempty"`

	/*
	   The version of the EKS add-on. The version must
	   match one of the versions returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
	*/
	AddonVersion string `json:"addonVersion,omitempty" yaml:"addonVersion,omitempty"`

	/*
	   Name of the EKS Cluster.

	   The following arguments are optional:
	*/
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// Define how to resolve parameter value conflicts when migrating an existing add-on to an Amazon EKS add-on or when applying version updates to the add-on. Valid values are `NONE`, `OVERWRITE` and `PRESERVE`. Note that `PRESERVE` is only valid on addon update, not for initial addon creation. If you need to set this to `PRESERVE`, use the `resolve_conflicts_on_create` and `resolve_conflicts_on_update` attributes instead. For more details check [UpdateAddon](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html) API Docs.
	ResolveConflicts string `json:"resolveConflicts,omitempty" yaml:"resolveConflicts,omitempty"`

	// How to resolve field value conflicts when migrating a self-managed add-on to an Amazon EKS add-on. Valid values are `NONE` and `OVERWRITE`. For more details see the [CreateAddon](https://docs.aws.amazon.com/eks/latest/APIReference/API_CreateAddon.html) API Docs.
	ResolveConflictsOnCreate string `json:"resolveConflictsOnCreate,omitempty" yaml:"resolveConflictsOnCreate,omitempty"`
}
