package dataloss

import types "DesignSphere_Server/src/resource/gcp/types"

type PreventionJobTrigger struct {
	// A short description of where the data is coming from. Will be stored once in the job. 256 max length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// User set display name of the job trigger.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Controls what and how to inspect for findings.
	   Structure is documented below.
	*/
	InspectJob types.Dataloss_PreventionJobTriggerInspectJob `json:"inspectJob,omitempty" yaml:"inspectJob,omitempty"`

	/*
	   The parent of the trigger, either in the format `projects/{{project}}`
	   or `projects/{{project}}/locations/{{location}}`
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   Whether the trigger is currently active.
	   Default value is `HEALTHY`.
	   Possible values are: `PAUSED`, `HEALTHY`, `CANCELLED`.
	*/
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	/*
	   The trigger id can contain uppercase and lowercase letters, numbers, and hyphens;
	   that is, it must match the regular expression: [a-zA-Z\d-_]+.
	   The maximum length is 100 characters. Can be empty to allow the system to generate one.
	*/
	TriggerId string `json:"triggerId,omitempty" yaml:"triggerId,omitempty"`

	/*
	   What event needs to occur for a new job to be started.
	   Structure is documented below.
	*/
	Triggers []types.Dataloss_PreventionJobTriggerTrigger `json:"triggers,omitempty" yaml:"triggers,omitempty"`
}
