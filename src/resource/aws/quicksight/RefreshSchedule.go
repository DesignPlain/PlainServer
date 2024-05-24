package quicksight

import types "DesignSphere_Server/src/resource/aws/types"

type RefreshSchedule struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// The ID of the dataset.
	DataSetId string `json:"dataSetId,omitempty" yaml:"dataSetId,omitempty"`

	/*
	   The [refresh schedule](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_RefreshSchedule.html). See schedule

	   The following arguments are optional:
	*/
	Schedule types.Quicksight_RefreshScheduleSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// The ID of the refresh schedule.
	ScheduleId string `json:"scheduleId,omitempty" yaml:"scheduleId,omitempty"`
}
