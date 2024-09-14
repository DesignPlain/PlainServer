package cloudids

type Endpoint struct {
	// An optional description of the endpoint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The location for the endpoint.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name of the endpoint in the format projects/{project_id}/locations/{locationId}/endpoints/{endpointId}.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Name of the VPC network that is connected to the IDS endpoint. This can either contain the VPC network name itself (like "src-net") or the full URL to the network (like "projects/{project_id}/global/networks/src-net").
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The minimum alert severity level that is reported by the endpoint.
	   Possible values are: `INFORMATIONAL`, `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`.
	*/
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`

	// Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.
	ThreatExceptions []string `json:"threatExceptions,omitempty" yaml:"threatExceptions,omitempty"`
}
