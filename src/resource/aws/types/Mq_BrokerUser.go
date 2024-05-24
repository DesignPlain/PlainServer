package types

type Mq_BrokerUser struct {
	// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user. Applies to `engine_type` of `ActiveMQ` only.
	ConsoleAccess bool `json:"consoleAccess,omitempty" yaml:"consoleAccess,omitempty"`

	// List of groups (20 maximum) to which the ActiveMQ user belongs. Applies to `engine_type` of `ActiveMQ` only.
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	// Password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Whether to set set replication user. Defaults to `false`.
	ReplicationUser bool `json:"replicationUser,omitempty" yaml:"replicationUser,omitempty"`

	/*
	   Username of the user.

	   > --NOTE:-- AWS currently does not support updating RabbitMQ users. Updates to users can only be in the RabbitMQ UI.
	*/
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
