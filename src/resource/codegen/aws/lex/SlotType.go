package lex

import types "libds/aws/types"

type SlotType struct {
	/*
	   Determines the slot resolution strategy that Amazon Lex
	   uses to return slot type values. `ORIGINAL_VALUE` returns the value entered by the user if the user
	   value is similar to the slot value. `TOP_RESOLUTION` returns the first value in the resolution list
	   if there is a resolution list for the slot, otherwise null is returned. Defaults to `ORIGINAL_VALUE`.
	*/
	ValueSelectionStrategy string `json:"valueSelectionStrategy,omitempty" yaml:"valueSelectionStrategy,omitempty"`

	/*
	   Determines if a new slot type version is created when the initial resource is created and on each
	   update. Defaults to `false`.
	*/
	CreateVersion bool `json:"createVersion,omitempty" yaml:"createVersion,omitempty"`

	// A description of the slot type. Must be less than or equal to 200 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A list of EnumerationValue objects that defines the values that
	   the slot type can take. Each value can have a list of synonyms, which are additional values that help
	   train the machine learning model about the values that it resolves for a slot. Attributes are
	   documented under enumeration_value.
	*/
	EnumerationValues []types.Lex_SlotTypeEnumerationValue `json:"enumerationValues,omitempty" yaml:"enumerationValues,omitempty"`

	// The name of the slot type. The name is not case sensitive. Must be less than or equal to 100 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
