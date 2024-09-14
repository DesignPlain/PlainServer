package types

type Dataplex_DatascanDataProfileSpecPostScanActions struct {
	/*
	   If set, results will be exported to the provided BigQuery table.
	   Structure is documented below.
	*/
	BigqueryExport Dataplex_DatascanDataProfileSpecPostScanActionsBigqueryExport `json:"bigqueryExport,omitempty" yaml:"bigqueryExport,omitempty"`
}
