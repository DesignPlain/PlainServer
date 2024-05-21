package types

type Dataplex_DatascanExecutionSpec struct {
	// The unnested field (of type Date or Timestamp) that contains values which monotonically increase over time. If not specified, a data scan will run for all data in the table.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   Spec related to how often and when a scan should be triggered.
	   Structure is documented below.
	*/
	Trigger Dataplex_DatascanExecutionSpecTrigger `json:"trigger,omitempty" yaml:"trigger,omitempty"`
}
