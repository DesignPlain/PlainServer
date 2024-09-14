package types

type Datacatalog_EntryBigqueryTableSpec struct {
	/*
	   (Output)
	   The table source type.
	*/
	TableSourceType string `json:"tableSourceType,omitempty" yaml:"tableSourceType,omitempty"`

	/*
	   (Output)
	   Spec of a BigQuery table. This field should only be populated if tableSourceType is BIGQUERY_TABLE.
	   Structure is documented below.
	*/
	TableSpecs []Datacatalog_EntryBigqueryTableSpecTableSpec `json:"tableSpecs,omitempty" yaml:"tableSpecs,omitempty"`

	/*
	   (Output)
	   Table view specification. This field should only be populated if tableSourceType is BIGQUERY_VIEW.
	   Structure is documented below.
	*/
	ViewSpecs []Datacatalog_EntryBigqueryTableSpecViewSpec `json:"viewSpecs,omitempty" yaml:"viewSpecs,omitempty"`
}
