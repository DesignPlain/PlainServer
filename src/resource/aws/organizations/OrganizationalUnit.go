package organizations

type OrganizationalUnit struct {
	// The name for the organizational unit
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ID of the parent organizational unit, which may be the root
	ParentId string `json:"parentId,omitempty" yaml:"parentId,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
