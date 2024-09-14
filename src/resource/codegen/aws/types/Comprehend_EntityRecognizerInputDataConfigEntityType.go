package types

type Comprehend_EntityRecognizerInputDataConfigEntityType struct {
	/*
	   An entity type to be matched by the Entity Recognizer.
	   Cannot contain a newline (`\n`), carriage return (`\r`), or tab (`\t`).
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
