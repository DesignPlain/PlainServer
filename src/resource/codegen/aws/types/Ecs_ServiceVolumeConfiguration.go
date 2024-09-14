package types

type Ecs_ServiceVolumeConfiguration struct {
	// Configuration for the Amazon EBS volume that Amazon ECS creates and manages on your behalf. See below.
	ManagedEbsVolume Ecs_ServiceVolumeConfigurationManagedEbsVolume `json:"managedEbsVolume,omitempty" yaml:"managedEbsVolume,omitempty"`

	// Name of the volume.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
