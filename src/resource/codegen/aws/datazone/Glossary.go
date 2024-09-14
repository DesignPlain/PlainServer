package datazone

type Glossary struct {
	//
	DomainIdentifier string `json:"domainIdentifier,omitempty" yaml:"domainIdentifier,omitempty"`

	// Name of the glossary. Must have length between 1 and 256.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   ID of the project that owns business glossary. Must follow regex of ^[a-zA-Z0-9_-]{1,36}$.

	   The following arguments are optional:
	*/
	OwningProjectIdentifier string `json:"owningProjectIdentifier,omitempty" yaml:"owningProjectIdentifier,omitempty"`

	// Status of business glossary. Valid values are DISABLED and ENABLED.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Description of the glossary. Must have a length between 0 and 4096.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
