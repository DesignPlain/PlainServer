package types

type Healthcare_DicomStoreStreamConfig struct {
	/*
	   BigQueryDestination to include a fully qualified BigQuery table URI where DICOM instance metadata will be streamed.
	   Structure is documented below.
	*/
	BigqueryDestination Healthcare_DicomStoreStreamConfigBigqueryDestination `json:"bigqueryDestination,omitempty" yaml:"bigqueryDestination,omitempty"`
}
