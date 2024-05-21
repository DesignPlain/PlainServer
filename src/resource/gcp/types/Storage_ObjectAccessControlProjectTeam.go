package types

type Storage_ObjectAccessControlProjectTeam struct {
	// The project team associated with the entity
	ProjectNumber string `json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`

	/*
	   The team.
	   Possible values are: `editors`, `owners`, `viewers`.
	*/
	Team string `json:"team,omitempty" yaml:"team,omitempty"`
}
