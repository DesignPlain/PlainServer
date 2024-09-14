package codestarconnections

type Connection struct {
	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `provider_type`
	HostArn string `json:"hostArn,omitempty" yaml:"hostArn,omitempty"`

	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub`, `GitHubEnterpriseServer`, `GitLab` or `GitLabSelfManaged`. Changing `provider_type` will create a new resource. Conflicts with `host_arn`
	ProviderType string `json:"providerType,omitempty" yaml:"providerType,omitempty"`

	// Map of key-value resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
