package types

type Cloudtasks_QueueAppEngineRoutingOverride struct {
	/*
	   App service.
	   By default, the task is sent to the service which is the default service when the task is attempted.
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   App version.
	   By default, the task is sent to the version which is the default version when the task is attempted.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   (Output)
	   The host that the task is sent to.
	*/
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   App instance.
	   By default, the task is sent to an instance which is available when the task is attempted.
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`
}
