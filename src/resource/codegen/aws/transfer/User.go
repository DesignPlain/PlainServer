package transfer

import types "libds/aws/types"

type User struct {
	// Amazon Resource Name (ARN) of an IAM role that allows the service to control your user’s access to your Amazon S3 bucket.
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// The name used for log in to your SFTP server.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// Logical directory mappings that specify what S3 paths and keys should be visible to your user and how you want to make them visible. See Home Directory Mappings below.
	HomeDirectoryMappings []types.Transfer_UserHomeDirectoryMapping `json:"homeDirectoryMappings,omitempty" yaml:"homeDirectoryMappings,omitempty"`

	// The type of landing directory (folder) you mapped for your users' home directory. Valid values are `PATH` and `LOGICAL`.
	HomeDirectoryType string `json:"homeDirectoryType,omitempty" yaml:"homeDirectoryType,omitempty"`

	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Specifies the full POSIX identity, including user ID (Uid), group ID (Gid), and any secondary groups IDs (SecondaryGids), that controls your users' access to your Amazon EFS file systems. See Posix Profile below.
	PosixProfile types.Transfer_UserPosixProfile `json:"posixProfile,omitempty" yaml:"posixProfile,omitempty"`

	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory string `json:"homeDirectory,omitempty" yaml:"homeDirectory,omitempty"`

	// The Server ID of the Transfer Server (e.g., `s-12345678`)
	ServerId string `json:"serverId,omitempty" yaml:"serverId,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
