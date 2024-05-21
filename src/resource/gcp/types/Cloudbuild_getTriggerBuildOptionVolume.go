package types

type Cloudbuild_getTriggerBuildOptionVolume struct {
	/*
	   Name of the volume to mount.

	   Volume names must be unique per build step and must be valid names for Docker volumes.
	   Each named volume must be used by at least two build steps.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Path at which to mount the volume.

	   Paths must be absolute and cannot conflict with other volume paths on the same
	   build step or with certain reserved volume paths.
	*/
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
