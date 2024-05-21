package types

type Backupdisasterrecovery_getManagementServerNetwork struct {
	// Type of Network peeringMode Default value: "PRIVATE_SERVICE_ACCESS" Possible values: ["PRIVATE_SERVICE_ACCESS"]
	PeeringMode string `json:"peeringMode,omitempty" yaml:"peeringMode,omitempty"`

	// Network with format 'projects/{{project_id}}/global/networks/{{network_id}}'
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
