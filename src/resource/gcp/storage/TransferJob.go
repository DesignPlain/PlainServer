package storage

import types "DesignSphere_Server/src/resource/gcp/types"

type TransferJob struct {
	/*
	   Transfer specification. Structure documented below.

	   - - -
	*/
	TransferSpec types.Storage_TransferJobTransferSpec `json:"transferSpec,omitempty" yaml:"transferSpec,omitempty"`

	// Unique description to identify the Transfer Job.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the Event-driven transfer options. Event-driven transfers listen to an event stream to transfer updated files. Structure documented below Either `event_stream` or `schedule` must be set.
	EventStream types.Storage_TransferJobEventStream `json:"eventStream,omitempty" yaml:"eventStream,omitempty"`

	// The name of the Transfer Job. This name must start with "transferJobs/" prefix and end with a letter or a number, and should be no more than 128 characters ( `transferJobs/^(?!OPI)[A-Za-z0-9-._~]-[A-Za-z0-9]$` ). For transfers involving PosixFilesystem, this name must start with transferJobs/OPI specifically ( `transferJobs/OPI^[A-Za-z0-9-._~]-[A-Za-z0-9]$` ). For all other transfer types, this name must not start with transferJobs/OPI. Default the provider will assign a random unique name with `transferJobs/{{name}}` format, where `name` is a numeric value.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Notification configuration. This is not supported for transfers involving PosixFilesystem. Structure documented below.
	NotificationConfig types.Storage_TransferJobNotificationConfig `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`

	/*
	   The project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below. Either `schedule` or `event_stream` must be set.
	Schedule types.Storage_TransferJobSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Status of the job. Default: `ENABLED`. --NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.--
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
