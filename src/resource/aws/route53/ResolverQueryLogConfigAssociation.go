package route53

type ResolverQueryLogConfigAssociation struct {
	// The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
	ResolverQueryLogConfigId string `json:"resolverQueryLogConfigId,omitempty" yaml:"resolverQueryLogConfigId,omitempty"`

	// The ID of a VPC that you want this query logging configuration to log queries for.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`
}
