package storage

import types "DesignSphere_Server/src/resource/gcp/types"

type InsightsReportConfig struct {
	/*
	   Options for configuring the format of the inventory report CSV file.
	   Structure is documented below.
	*/
	CsvOptions types.Storage_InsightsReportConfigCsvOptions `json:"csvOptions,omitempty" yaml:"csvOptions,omitempty"`

	// The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Options for configuring how inventory reports are generated.
	   Structure is documented below.
	*/
	FrequencyOptions types.Storage_InsightsReportConfigFrequencyOptions `json:"frequencyOptions,omitempty" yaml:"frequencyOptions,omitempty"`

	/*
	   The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
	   must be in the same location.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Options for including metadata in an inventory report.
	   Structure is documented below.
	*/
	ObjectMetadataReportOptions types.Storage_InsightsReportConfigObjectMetadataReportOptions `json:"objectMetadataReportOptions,omitempty" yaml:"objectMetadataReportOptions,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
