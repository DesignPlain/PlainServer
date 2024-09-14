package types

type Cloudrun_getServiceTemplateSpecContainerVolumeMount struct {
	/*
	   Path within the container at which the volume should be mounted.  Must
	   not contain ':'.
	*/
	MountPath string `json:"mountPath,omitempty" yaml:"mountPath,omitempty"`

	// The name of the Cloud Run Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
