package types

type Organizations_getOrganizationalUnitsChild struct {
	// ARN of the organizational unit
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Parent identifier of the organizational units.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the organizational unit
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
