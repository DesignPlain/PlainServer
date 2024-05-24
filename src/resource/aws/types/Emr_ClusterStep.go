package types

type Emr_ClusterStep struct {
	// JAR file used for the step. See below.
	HadoopJarStep Emr_ClusterStepHadoopJarStep `json:"hadoopJarStep,omitempty" yaml:"hadoopJarStep,omitempty"`

	// Name of the step.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Action to take if the step fails. Valid values: `TERMINATE_JOB_FLOW`, `TERMINATE_CLUSTER`, `CANCEL_AND_WAIT`, and `CONTINUE`
	ActionOnFailure string `json:"actionOnFailure,omitempty" yaml:"actionOnFailure,omitempty"`
}
