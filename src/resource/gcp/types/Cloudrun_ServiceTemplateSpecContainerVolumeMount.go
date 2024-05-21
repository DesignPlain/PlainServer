package types

type Cloudrun_ServiceTemplateSpecContainerVolumeMount struct {
	/*
	   Path within the container at which the volume should be mounted.  Must
	   not contain ':'.
	*/
	MountPath string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`

	// This must match the Name of a Volume.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
