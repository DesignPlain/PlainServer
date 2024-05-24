package types

type Codebuild_ProjectFileSystemLocation struct {
	// The name used to access a file system created by Amazon EFS. CodeBuild creates an environment variable by appending the identifier in all capital letters to CODEBUILD\_. For example, if you specify my-efs for identifier, a new environment variable is create named CODEBUILD_MY-EFS.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// A string that specifies the location of the file system created by Amazon EFS. Its format is `efs-dns-name:/directory-path`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The mount options for a file system created by AWS EFS.
	MountOptions string `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`

	// The location in the container where you mount the file system.
	MountPoint string `json:"mountPoint,omitempty" yaml:"mountPoint,omitempty"`

	// The type of the file system. The one supported type is `EFS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
