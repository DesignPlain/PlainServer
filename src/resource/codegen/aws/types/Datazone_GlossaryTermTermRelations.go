package types

type Datazone_GlossaryTermTermRelations struct {
	// String array that calssifies the term relations.
	Classifies []string `json:"classifies,omitempty" yaml:"classifies,omitempty"`

	//
	IsAs []string `json:"isAs,omitempty" yaml:"isAs,omitempty"`
}
