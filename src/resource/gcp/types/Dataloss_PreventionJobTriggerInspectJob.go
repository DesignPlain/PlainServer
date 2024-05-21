package types

type Dataloss_PreventionJobTriggerInspectJob struct {
	/*
	   Information on where to inspect
	   Structure is documented below.
	*/
	StorageConfig Dataloss_PreventionJobTriggerInspectJobStorageConfig `json:"storageConfig,omitempty" yaml:"storageConfig,omitempty"`

	/*
	   Configuration block for the actions to execute on the completion of a job. Can be specified multiple times, but only one for each type. Each action block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	   Structure is documented below.
	*/
	Actions []Dataloss_PreventionJobTriggerInspectJobAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	/*
	   The core content of the template.
	   Structure is documented below.
	*/
	InspectConfig Dataloss_PreventionJobTriggerInspectJobInspectConfig `json:"inspectConfig,omitempty" yaml:"inspectConfig,omitempty"`

	// The name of the template to run when this job is triggered.
	InspectTemplateName string `json:"inspectTemplateName,omitempty" yaml:"inspectTemplateName,omitempty"`
}
