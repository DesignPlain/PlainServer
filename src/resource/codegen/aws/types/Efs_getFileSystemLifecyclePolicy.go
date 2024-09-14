package types

type Efs_getFileSystemLifecyclePolicy struct {
	//
	TransitionToArchive string `json:"transitionToArchive,omitempty" yaml:"transitionToArchive,omitempty"`

	//
	TransitionToIa string `json:"transitionToIa,omitempty" yaml:"transitionToIa,omitempty"`

	//
	TransitionToPrimaryStorageClass string `json:"transitionToPrimaryStorageClass,omitempty" yaml:"transitionToPrimaryStorageClass,omitempty"`
}
