package dataplex

import types "libds/gcp/types"

type Datascan struct {
	// User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   DataScan execution settings.
	   Structure is documented below.
	*/
	ExecutionSpec types.Dataplex_DatascanExecutionSpec `json:"executionSpec,omitempty" yaml:"executionSpec,omitempty"`

	// The location where the data scan should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Description of the rule.
	   The maximum length is 1,024 characters.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   DataProfileScan related setting.
	   Structure is documented below.
	*/
	DataProfileSpec types.Dataplex_DatascanDataProfileSpec `json:"dataProfileSpec,omitempty" yaml:"dataProfileSpec,omitempty"`

	/*
	   DataQualityScan related setting.
	   Structure is documented below.
	*/
	DataQualitySpec types.Dataplex_DatascanDataQualitySpec `json:"dataQualitySpec,omitempty" yaml:"dataQualitySpec,omitempty"`

	// DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
	DataScanId string `json:"dataScanId,omitempty" yaml:"dataScanId,omitempty"`

	/*
	   User-defined labels for the scan. A list of key->value pairs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The data source for DataScan.
	   Structure is documented below.
	*/
	Data types.Dataplex_DatascanData `json:"data,omitempty" yaml:"data,omitempty"`
}
