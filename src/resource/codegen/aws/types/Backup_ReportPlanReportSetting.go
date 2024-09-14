package types

type Backup_ReportPlanReportSetting struct {
	// Specifies the list of Organizational Units a report covers.
	OrganizationUnits []string `json:"organizationUnits,omitempty" yaml:"organizationUnits,omitempty"`

	// Specifies the list of regions a report covers.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`

	// Identifies the report template for the report. Reports are built using a report template. The report templates are: `RESOURCE_COMPLIANCE_REPORT` | `CONTROL_COMPLIANCE_REPORT` | `BACKUP_JOB_REPORT` | `COPY_JOB_REPORT` | `RESTORE_JOB_REPORT`.
	ReportTemplate string `json:"reportTemplate,omitempty" yaml:"reportTemplate,omitempty"`

	// Specifies the list of accounts a report covers.
	Accounts []string `json:"accounts,omitempty" yaml:"accounts,omitempty"`

	// Specifies the Amazon Resource Names (ARNs) of the frameworks a report covers.
	FrameworkArns []string `json:"frameworkArns,omitempty" yaml:"frameworkArns,omitempty"`

	// Specifies the number of frameworks a report covers.
	NumberOfFrameworks int `json:"numberOfFrameworks,omitempty" yaml:"numberOfFrameworks,omitempty"`
}
