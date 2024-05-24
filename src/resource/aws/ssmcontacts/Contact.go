package ssmcontacts

type Contact struct {
	// A unique and identifiable alias for the contact or escalation plan. Must be between 1 and 255 characters, and may contain alphanumerics, underscores (`_`), and hyphens (`-`).
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// Full friendly name of the contact or escalation plan. If set, must be between 1 and 255 characters, and may contain alphanumerics, underscores (`_`), hyphens (`-`), periods (`.`), and spaces.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Map of tags to assign to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The type of contact engaged. A single contact is type PERSONAL and an escalation
	   plan is type ESCALATION.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
