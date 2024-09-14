package types

type Sagemaker_DataQualityJobDefinitionDataQualityBaselineConfig struct {
	// The constraints resource for a monitoring job. Fields are documented below.
	ConstraintsResource Sagemaker_DataQualityJobDefinitionDataQualityBaselineConfigConstraintsResource `json:"constraintsResource,omitempty" yaml:"constraintsResource,omitempty"`

	// The statistics resource for a monitoring job. Fields are documented below.
	StatisticsResource Sagemaker_DataQualityJobDefinitionDataQualityBaselineConfigStatisticsResource `json:"statisticsResource,omitempty" yaml:"statisticsResource,omitempty"`
}
