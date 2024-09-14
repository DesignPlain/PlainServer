package datazone

import types "libds/aws/types"

type FormType struct {
	// Identifier of project that owns the form type. Must follow regex of ^[a-zA-Z0-9_-]{1,36}.
	OwningProjectIdentifier string `json:"owningProjectIdentifier,omitempty" yaml:"owningProjectIdentifier,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	//
	Timeouts types.Datazone_FormTypeTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Description of form type. Must have a length of between 1 and 2048 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Identifier of the domain.
	DomainIdentifier string `json:"domainIdentifier,omitempty" yaml:"domainIdentifier,omitempty"`

	// Object of the model of the form type that contains the following attributes.
	Model types.Datazone_FormTypeModel `json:"model,omitempty" yaml:"model,omitempty"`

	// Name of the form type. Must be the name of the structure in smithy document.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
