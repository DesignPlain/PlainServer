package transfer

type SshKey struct {
	// The public key portion of an SSH key pair.
	Body string `json:"body,omitempty" yaml:"body,omitempty"`

	// The Server ID of the Transfer Server (e.g., `s-12345678`)
	ServerId string `json:"serverId,omitempty" yaml:"serverId,omitempty"`

	// The name of the user account that is assigned to one or more servers.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`
}
