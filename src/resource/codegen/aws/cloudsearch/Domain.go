package cloudsearch

import types "libds/aws/types"

type Domain struct {
	// Domain endpoint options. Documented below.
	EndpointOptions types.Cloudsearch_DomainEndpointOptions `json:"endpointOptions,omitempty" yaml:"endpointOptions,omitempty"`

	// The index fields for documents added to the domain. Documented below.
	IndexFields []types.Cloudsearch_DomainIndexField `json:"indexFields,omitempty" yaml:"indexFields,omitempty"`

	// Whether or not to maintain extra instances for the domain in a second Availability Zone to ensure high availability.
	MultiAz bool `json:"multiAz,omitempty" yaml:"multiAz,omitempty"`

	// The name of the CloudSearch domain.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Domain scaling parameters. Documented below.
	ScalingParameters types.Cloudsearch_DomainScalingParameters `json:"scalingParameters,omitempty" yaml:"scalingParameters,omitempty"`
}
