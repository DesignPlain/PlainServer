package directoryservice

import types "DesignSphere_Server/src/resource/aws/types"

type ServiceRegion struct {
	// The number of domain controllers desired in the replicated directory. Minimum value of `2`.
	DesiredNumberOfDomainControllers int `json:"desiredNumberOfDomainControllers,omitempty" yaml:"desiredNumberOfDomainControllers,omitempty"`

	// The identifier of the directory to which you want to add Region replication.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// The name of the Region where you want to add domain controllers for replication.
	RegionName string `json:"regionName,omitempty" yaml:"regionName,omitempty"`

	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// VPC information in the replicated Region. Detailed below.
	VpcSettings types.Directoryservice_ServiceRegionVpcSettings `json:"vpcSettings,omitempty" yaml:"vpcSettings,omitempty"`
}
