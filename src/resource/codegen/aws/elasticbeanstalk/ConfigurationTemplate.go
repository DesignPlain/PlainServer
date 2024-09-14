package elasticbeanstalk

import types "libds/aws/types"

type ConfigurationTemplate struct {
	// name of the application to associate with this configuration template
	Application string `json:"application,omitempty" yaml:"application,omitempty"`

	// Short description of the Template
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the environment used with this configuration template
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	// A unique name for this Template.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Option settings to configure the new Environment. These
	   override specific values that are set as defaults. The format is detailed
	   below in Option Settings
	*/
	Settings []types.Elasticbeanstalk_ConfigurationTemplateSetting `json:"settings,omitempty" yaml:"settings,omitempty"`

	/*
	   A solution stack to base your Template
	   off of. Example stacks can be found in the [Amazon API documentation][1]
	*/
	SolutionStackName string `json:"solutionStackName,omitempty" yaml:"solutionStackName,omitempty"`
}
