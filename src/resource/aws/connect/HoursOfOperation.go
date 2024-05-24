package connect

import types "DesignSphere_Server/src/resource/aws/types"

type HoursOfOperation struct {
	// Specifies the name of the Hours of Operation.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Tags to apply to the Hours of Operation. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the time zone of the Hours of Operation.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	// One or more config blocks which define the configuration information for the hours of operation: day, start time, and end time . Config blocks are documented below.
	Configs []types.Connect_HoursOfOperationConfig `json:"configs,omitempty" yaml:"configs,omitempty"`

	// Specifies the description of the Hours of Operation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
}
