package types

type Firestore_IndexField struct {
	/*
	   Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	   Only one of `order` and `arrayConfig` can be specified.
	   Possible values are: `ASCENDING`, `DESCENDING`.
	*/
	Order string `json:"order,omitempty" yaml:"order,omitempty"`

	/*
	   Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
	   be specified.
	   Possible values are: `CONTAINS`.

	   - - -
	*/
	ArrayConfig string `json:"arrayConfig,omitempty" yaml:"arrayConfig,omitempty"`

	// Name of the field.
	FieldPath string `json:"fieldPath,omitempty" yaml:"fieldPath,omitempty"`
}
