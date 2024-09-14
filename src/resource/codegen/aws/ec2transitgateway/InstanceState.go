package ec2transitgateway

type InstanceState struct {
	// Whether to request a forced stop when `state` is `stopped`. Otherwise (_i.e._, `state` is `running`), ignored. When an instance is forced to stop, it does not flush file system caches or file system metadata, and you must subsequently perform file system check and repair. Not recommended for Windows instances. Defaults to `false`.
	Force bool `json:"force,omitempty" yaml:"force,omitempty"`

	// ID of the instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   State of the instance. Valid values are `stopped`, `running`.

	   The following arguments are optional:
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
