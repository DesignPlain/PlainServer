package types

type Dataplex_DatascanDataQualitySpecPostScanActions struct {
	/*
	   If set, results will be exported to the provided BigQuery table.
	   Structure is documented below.
	*/
	BigqueryExport Dataplex_DatascanDataQualitySpecPostScanActionsBigqueryExport `json:"bigqueryExport,omitempty" yaml:"bigqueryExport,omitempty"`
}
