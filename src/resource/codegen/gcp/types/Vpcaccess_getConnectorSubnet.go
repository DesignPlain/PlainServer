package types

type Vpcaccess_getConnectorSubnet struct {
	/*
	   Name of the resource.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
