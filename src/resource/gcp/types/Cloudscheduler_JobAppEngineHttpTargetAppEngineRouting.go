package types

type Cloudscheduler_JobAppEngineHttpTargetAppEngineRouting struct {
	/*
	   App instance.
	   By default, the job is sent to an instance which is available when the job is attempted.
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   App service.
	   By default, the job is sent to the service which is the default service when the job is attempted.
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   App version.
	   By default, the job is sent to the version which is the default version when the job is attempted.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
