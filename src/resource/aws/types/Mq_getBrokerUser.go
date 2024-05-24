package types

type Mq_getBrokerUser struct {
	//
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	//
	ConsoleAccess bool `json:"consoleAccess,omitempty" yaml:"consoleAccess,omitempty"`

	//
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	//
	ReplicationUser bool `json:"replicationUser,omitempty" yaml:"replicationUser,omitempty"`
}
