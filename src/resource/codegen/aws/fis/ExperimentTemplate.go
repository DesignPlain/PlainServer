package fis

import types "libds/aws/types"

type ExperimentTemplate struct {
	// The experiment options for the experiment template. See experiment_options below for more details!
	ExperimentOptions types.Fis_ExperimentTemplateExperimentOptions `json:"experimentOptions,omitempty" yaml:"experimentOptions,omitempty"`

	// The configuration for experiment logging. See below.
	LogConfiguration types.Fis_ExperimentTemplateLogConfiguration `json:"logConfiguration,omitempty" yaml:"logConfiguration,omitempty"`

	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	/*
	   When an ongoing experiment should be stopped. See below.

	   The following arguments are optional:
	*/
	StopConditions []types.Fis_ExperimentTemplateStopCondition `json:"stopConditions,omitempty" yaml:"stopConditions,omitempty"`

	// Key-value mapping of tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Target of an action. See below.
	Targets []types.Fis_ExperimentTemplateTarget `json:"targets,omitempty" yaml:"targets,omitempty"`

	// Action to be performed during an experiment. See below.
	Actions []types.Fis_ExperimentTemplateAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	// Description for the experiment template.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
