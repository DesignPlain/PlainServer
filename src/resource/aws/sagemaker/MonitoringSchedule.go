package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type MonitoringSchedule struct {
	// The configuration object that specifies the monitoring schedule and defines the monitoring job. Fields are documented below.
	MonitoringScheduleConfig types.Sagemaker_MonitoringScheduleMonitoringScheduleConfig `json:"monitoringScheduleConfig,omitempty" yaml:"monitoringScheduleConfig,omitempty"`

	// The name of the monitoring schedule. The name must be unique within an AWS Region within an AWS account. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
