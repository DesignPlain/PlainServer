package types

type Appengine_StandardAppVersionManualScaling struct {
	/*
	   Number of instances to assign to the service at the start.
	   --Note:-- When managing the number of instances at runtime through the App Engine Admin API or the (now deprecated) Python 2
	   Modules API set_num_instances() you must use `lifecycle.ignore_changes = ["manual_scaling"[0].instances]` to prevent drift detection.
	*/
	Instances int `json:"instances,omitempty" yaml:"instances,omitempty"`
}
