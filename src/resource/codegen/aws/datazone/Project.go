package datazone

import types "libds/aws/types"

type Project struct {
	// Description of project.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Identifier of domain which the project is part of. Must follow the regex of `^dzd[-_][a-zA-Z0-9_-]{1,36}$`.
	DomainIdentifier string `json:"domainIdentifier,omitempty" yaml:"domainIdentifier,omitempty"`

	// List of glossary terms that can be used in the project. The list cannot be empty or include over 20 values. Each value must follow the regex of `[a-zA-Z0-9_-]{1,36}$`.
	GlossaryTerms []string `json:"glossaryTerms,omitempty" yaml:"glossaryTerms,omitempty"`

	/*
	   Name of the project. Must follow the regex of `^[\w -]+$`. and have a length of at most 64.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Optional flag to delete all child entities within the project.
	SkipDeletionCheck bool `json:"skipDeletionCheck,omitempty" yaml:"skipDeletionCheck,omitempty"`

	//
	Timeouts types.Datazone_ProjectTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
