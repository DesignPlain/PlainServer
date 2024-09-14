package types

type Backupdisasterrecovery_ManagementServerNetwork struct {
	// Network with format `projects/{{project_id}}/global/networks/{{network_id}}`
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Type of Network peeringMode
	   Default value is `PRIVATE_SERVICE_ACCESS`.
	   Possible values are: `PRIVATE_SERVICE_ACCESS`.

	   - - -
	*/
	PeeringMode string `json:"peeringMode,omitempty" yaml:"peeringMode,omitempty"`
}
