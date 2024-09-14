package types

type Elasticbeanstalk_EnvironmentSetting struct {
	//
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	/*
	   A unique name for this Environment. This name is used
	   in the application URL
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	//
	Resource string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
