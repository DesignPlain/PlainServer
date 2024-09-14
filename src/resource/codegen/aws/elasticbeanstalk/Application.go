package elasticbeanstalk

import types "libds/aws/types"

type Application struct {
	//
	AppversionLifecycle types.Elasticbeanstalk_ApplicationAppversionLifecycle `json:"appversionLifecycle,omitempty" yaml:"appversionLifecycle,omitempty"`

	// Short description of the application
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the application, must be unique within your account
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of tags for the Elastic Beanstalk Application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
