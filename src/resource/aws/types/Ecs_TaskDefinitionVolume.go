package types

type Ecs_TaskDefinitionVolume struct {
	// Configuration block to configure a docker volume. Detailed below.
	DockerVolumeConfiguration Ecs_TaskDefinitionVolumeDockerVolumeConfiguration `json:"dockerVolumeConfiguration,omitempty" yaml:"dockerVolumeConfiguration,omitempty"`

	// Configuration block for an EFS volume. Detailed below.
	EfsVolumeConfiguration Ecs_TaskDefinitionVolumeEfsVolumeConfiguration `json:"efsVolumeConfiguration,omitempty" yaml:"efsVolumeConfiguration,omitempty"`

	// Configuration block for an FSX Windows File Server volume. Detailed below.
	FsxWindowsFileServerVolumeConfiguration Ecs_TaskDefinitionVolumeFsxWindowsFileServerVolumeConfiguration `json:"fsxWindowsFileServerVolumeConfiguration,omitempty" yaml:"fsxWindowsFileServerVolumeConfiguration,omitempty"`

	// Path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
	HostPath string `json:"hostPath,omitempty" yaml:"hostPath,omitempty"`

	/*
	   Name of the volume. This name is referenced in the `sourceVolume`
	   parameter of container definition in the `mountPoints` section.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
