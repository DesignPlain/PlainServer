package ram

type ResourceShare struct {
	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals bool `json:"allowExternalPrincipals,omitempty" yaml:"allowExternalPrincipals,omitempty"`

	// The name of the resource share.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the Amazon Resource Names (ARNs) of the RAM permission to associate with the resource share. If you do not specify an ARN for the permission, RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
	PermissionArns []string `json:"permissionArns,omitempty" yaml:"permissionArns,omitempty"`

	// A map of tags to assign to the resource share. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
