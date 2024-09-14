package types

type Efs_getAccessPointRootDirectory struct {
	// Path exposed as the root directory
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Single element list containing information on the creation permissions of the directory
	CreationInfos []Efs_getAccessPointRootDirectoryCreationInfo `json:"creationInfos,omitempty" yaml:"creationInfos,omitempty"`
}
