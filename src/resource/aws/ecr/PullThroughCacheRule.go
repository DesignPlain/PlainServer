package ecr

type PullThroughCacheRule struct {
	// ARN of the Secret which will be used to authenticate against the registry.
	CredentialArn string `json:"credentialArn,omitempty" yaml:"credentialArn,omitempty"`

	// The repository name prefix to use when caching images from the source registry.
	EcrRepositoryPrefix string `json:"ecrRepositoryPrefix,omitempty" yaml:"ecrRepositoryPrefix,omitempty"`

	// The registry URL of the upstream public registry to use as the source.
	UpstreamRegistryUrl string `json:"upstreamRegistryUrl,omitempty" yaml:"upstreamRegistryUrl,omitempty"`
}
