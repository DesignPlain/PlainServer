package types

type Gkebackup_RestorePlanRestoreConfigTransformationRuleFieldAction struct {
	/*
	   A string containing a JSON Pointer value that references the
	   location in the target document to move the value from.
	*/
	FromPath string `json:"fromPath,omitempty" yaml:"fromPath,omitempty"`

	/*
	   Specifies the operation to perform.
	   Possible values are: `REMOVE`, `MOVE`, `COPY`, `ADD`, `TEST`, `REPLACE`.
	*/
	Op string `json:"op,omitempty" yaml:"op,omitempty"`

	/*
	   A string containing a JSON-Pointer value that references a
	   location within the target document where the operation is performed.
	*/
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   A string that specifies the desired value in string format
	   to use for transformation.

	   - - -
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
