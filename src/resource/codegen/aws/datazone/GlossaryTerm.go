package datazone

import types "libds/aws/types"

type GlossaryTerm struct {
	// If glossary term is ENABLED or DISABLED.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Object classifying the term relations through the following attributes:
	TermRelations types.Datazone_GlossaryTermTermRelations `json:"termRelations,omitempty" yaml:"termRelations,omitempty"`

	//
	Timeouts types.Datazone_GlossaryTermTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Identifier of domain.
	DomainIdentifier string `json:"domainIdentifier,omitempty" yaml:"domainIdentifier,omitempty"`

	// Identifier of glossary.
	GlossaryIdentifier string `json:"glossaryIdentifier,omitempty" yaml:"glossaryIdentifier,omitempty"`

	// Long description of entry.
	LongDescription string `json:"longDescription,omitempty" yaml:"longDescription,omitempty"`

	/*
	   Name of glossary term.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Short description of entry.
	ShortDescription string `json:"shortDescription,omitempty" yaml:"shortDescription,omitempty"`
}
