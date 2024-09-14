package types

type Datazone_FormTypeModel struct {
	/*
	   Smithy document that indicates the model of the API. Must be between the lengths 1 and 100,000 and be encoded as a smithy document.

	   The following arguments are optional:
	*/
	Smithy string `json:"smithy,omitempty" yaml:"smithy,omitempty"`
}
