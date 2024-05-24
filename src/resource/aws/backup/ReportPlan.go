package backup

import types "DesignSphere_Server/src/resource/aws/types"

type ReportPlan struct {
	// The description of the report plan with a maximum of 1,024 characters
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// An object that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. Detailed below.
	ReportDeliveryChannel types.Backup_ReportPlanReportDeliveryChannel `json:"reportDeliveryChannel,omitempty" yaml:"reportDeliveryChannel,omitempty"`

	// An object that identifies the report template for the report. Reports are built using a report template. Detailed below.
	ReportSetting types.Backup_ReportPlanReportSetting `json:"reportSetting,omitempty" yaml:"reportSetting,omitempty"`

	// Metadata that you can assign to help organize the report plans you create. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
