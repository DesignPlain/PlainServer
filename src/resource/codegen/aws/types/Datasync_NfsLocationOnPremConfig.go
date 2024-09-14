package types

type Datasync_NfsLocationOnPremConfig struct {
	// List of Amazon Resource Names (ARNs) of the DataSync Agents used to connect to the NFS server.
	AgentArns []string `json:"agentArns,omitempty" yaml:"agentArns,omitempty"`
}
