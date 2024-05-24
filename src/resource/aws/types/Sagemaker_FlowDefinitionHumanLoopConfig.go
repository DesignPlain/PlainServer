package types

type Sagemaker_FlowDefinitionHumanLoopConfig struct {
	// The Amazon Resource Name (ARN) of the human task user interface.
	HumanTaskUiArn string `json:"humanTaskUiArn,omitempty" yaml:"humanTaskUiArn,omitempty"`

	// The number of distinct workers who will perform the same task on each object. Valid value range between `1` and `3`.
	TaskCount int `json:"taskCount,omitempty" yaml:"taskCount,omitempty"`

	// A description for the human worker task.
	TaskDescription string `json:"taskDescription,omitempty" yaml:"taskDescription,omitempty"`

	// The Amazon Resource Name (ARN) of the human task user interface. Amazon Resource Name (ARN) of a team of workers. For Public workforces see [AWS Docs](https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-management-public.html).
	WorkteamArn string `json:"workteamArn,omitempty" yaml:"workteamArn,omitempty"`

	// A title for the human worker task.
	TaskTitle string `json:"taskTitle,omitempty" yaml:"taskTitle,omitempty"`

	// Defines the amount of money paid to an Amazon Mechanical Turk worker for each task performed. See Public Workforce Task Price details below.
	PublicWorkforceTaskPrice Sagemaker_FlowDefinitionHumanLoopConfigPublicWorkforceTaskPrice `json:"publicWorkforceTaskPrice,omitempty" yaml:"publicWorkforceTaskPrice,omitempty"`

	// The length of time that a task remains available for review by human workers. Valid value range between `1` and `864000`.
	TaskAvailabilityLifetimeInSeconds int `json:"taskAvailabilityLifetimeInSeconds,omitempty" yaml:"taskAvailabilityLifetimeInSeconds,omitempty"`

	// An array of keywords used to describe the task so that workers can discover the task.
	TaskKeywords []string `json:"taskKeywords,omitempty" yaml:"taskKeywords,omitempty"`

	// The amount of time that a worker has to complete a task. The default value is `3600` seconds.
	TaskTimeLimitInSeconds int `json:"taskTimeLimitInSeconds,omitempty" yaml:"taskTimeLimitInSeconds,omitempty"`
}
