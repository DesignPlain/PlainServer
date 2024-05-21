package types

type Dataplex_DatascanDataProfileSpecPostScanActionsBigqueryExport struct {
	/*
	   The BigQuery table to export DataProfileScan results to.
	   Format://bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	*/
	ResultsTable string `json:"resultsTable,omitempty" yaml:"resultsTable,omitempty"`
}
