package route53

type ResolverConfig struct {
	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: `ENABLE`, `DISABLE`.
	AutodefinedReverseFlag string `json:"autodefinedReverseFlag,omitempty" yaml:"autodefinedReverseFlag,omitempty"`

	// The ID of the VPC that the configuration is for.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`
}
