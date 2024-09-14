package types

type Healthcare_DicomStoreStreamConfigBigqueryDestination struct {
	// a fully qualified BigQuery table URI where DICOM instance metadata will be streamed.
	TableUri string `json:"tableUri,omitempty" yaml:"tableUri,omitempty"`
}
