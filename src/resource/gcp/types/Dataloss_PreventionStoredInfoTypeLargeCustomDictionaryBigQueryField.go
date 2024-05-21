package types

type Dataloss_PreventionStoredInfoTypeLargeCustomDictionaryBigQueryField struct {
	/*
	   Designated field in the BigQuery table.
	   Structure is documented below.
	*/
	Field Dataloss_PreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   Field in a BigQuery table where each cell represents a dictionary phrase.
	   Structure is documented below.
	*/
	Table Dataloss_PreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable `json:"table,omitempty" yaml:"table,omitempty"`
}
