package dataplex

import types "libds/gcp/types"

type Task struct {
	// User-provided description of the task.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Configuration for the cluster
	   Structure is documented below.
	*/
	ExecutionSpec types.Dataplex_TaskExecutionSpec `json:"executionSpec,omitempty" yaml:"executionSpec,omitempty"`

	// The location in which the task will be created in.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	   Structure is documented below.
	*/
	Notebook types.Dataplex_TaskNotebook `json:"notebook,omitempty" yaml:"notebook,omitempty"`

	// The task Id of the task.
	TaskId string `json:"taskId,omitempty" yaml:"taskId,omitempty"`

	/*
	   Configuration for the cluster
	   Structure is documented below.
	*/
	TriggerSpec types.Dataplex_TaskTriggerSpec `json:"triggerSpec,omitempty" yaml:"triggerSpec,omitempty"`

	// User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   User-defined labels for the task.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The lake in which the task will be created in.
	Lake string `json:"lake,omitempty" yaml:"lake,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	   Structure is documented below.
	*/
	Spark types.Dataplex_TaskSpark `json:"spark,omitempty" yaml:"spark,omitempty"`
}
