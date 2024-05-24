package finspace

import types "DesignSphere_Server/src/resource/aws/types"

type KxVolume struct {
	// Description of the volume.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A unique identifier for the kdb environment, whose clusters can attach to the volume.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	// Unique name for the volumr that you want to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the configuration for the Network attached storage (`NAS_1`) file system volume. This parameter is required when `volume_type` is `NAS_1`. See `nas1_configuration` Argument Reference below.
	Nas1Configurations []types.Finspace_KxVolumeNas1Configuration `json:"nas1Configurations,omitempty" yaml:"nas1Configurations,omitempty"`

	// A list of key-value pairs to label the volume. You can add up to 50 tags to a volume
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of file system volume. Currently, FinSpace only supports the `NAS_1` volume type. When you select the `NAS_1` volume type, you must also provide `nas1_configuration`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The identifier of the AWS Availability Zone IDs.

	   The following arguments are optional:
	*/
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// The number of availability zones you want to assign per volume. Currently, Finspace only support SINGLE for volumes.
	AzMode string `json:"azMode,omitempty" yaml:"azMode,omitempty"`
}
