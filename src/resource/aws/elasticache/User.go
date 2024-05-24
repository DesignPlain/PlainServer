package elasticache

import types "DesignSphere_Server/src/resource/aws/types"

type User struct {
	// Denotes the user's authentication properties. Detailed below.
	AuthenticationMode types.Elasticache_UserAuthenticationMode `json:"authenticationMode,omitempty" yaml:"authenticationMode,omitempty"`

	// The current supported value is `REDIS`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Indicates a password is not required for this user.
	NoPasswordRequired bool `json:"noPasswordRequired,omitempty" yaml:"noPasswordRequired,omitempty"`

	// Passwords used for this user. You can create up to two passwords for each user.
	Passwords []string `json:"passwords,omitempty" yaml:"passwords,omitempty"`

	// A list of tags to be added to this resource. A tag is a key-value pair.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the user.
	UserId string `json:"userId,omitempty" yaml:"userId,omitempty"`

	/*
	   The username of the user.

	   The following arguments are optional:
	*/
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// Access permissions string used for this user. See [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html#Access-string) for more details.
	AccessString string `json:"accessString,omitempty" yaml:"accessString,omitempty"`
}
