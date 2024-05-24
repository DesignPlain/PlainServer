package codebuild

import types "DesignSphere_Server/src/resource/aws/types"

type ReportGroup struct {
	// If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
	DeleteReports bool `json:"deleteReports,omitempty" yaml:"deleteReports,omitempty"`

	// Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
	ExportConfig types.Codebuild_ReportGroupExportConfig `json:"exportConfig,omitempty" yaml:"exportConfig,omitempty"`

	// The name of a Report Group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of the Report Group. Valid value are `TEST` and `CODE_COVERAGE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
