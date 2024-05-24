package elasticbeanstalk

import types "DesignSphere_Server/src/resource/aws/types"

type Environment struct {
	// Short description of the Environment
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The time between polling the AWS API to
	   check if changes have been applied. Use this to adjust the rate of API calls
	   for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	   use the default behavior, which is an exponential backoff
	*/
	PollInterval string `json:"pollInterval,omitempty" yaml:"pollInterval,omitempty"`

	/*
	   Option settings to configure the new Environment. These
	   override specific values that are set as defaults. The format is detailed
	   below in Option Settings
	*/
	Settings []types.Elasticbeanstalk_EnvironmentSetting `json:"settings,omitempty" yaml:"settings,omitempty"`

	/*
	   A solution stack to base your environment
	   off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	*/
	SolutionStackName string `json:"solutionStackName,omitempty" yaml:"solutionStackName,omitempty"`

	/*
	   The name of the Elastic Beanstalk Configuration
	   template to use in deployment
	*/
	TemplateName string `json:"templateName,omitempty" yaml:"templateName,omitempty"`

	/*
	   Elastic Beanstalk Environment tier. Valid values are `Worker`
	   or `WebServer`. If tier is left blank `WebServer` will be used.
	*/
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`

	/*
	   Name of the application that contains the version
	   to be deployed
	*/
	Application string `json:"application,omitempty" yaml:"application,omitempty"`

	/*
	   Prefix to use for the fully qualified DNS name of
	   the Environment.
	*/
	CnamePrefix string `json:"cnamePrefix,omitempty" yaml:"cnamePrefix,omitempty"`

	/*
	   The maximum
	   [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	   wait for an Elastic Beanstalk Environment to be in a ready state before timing
	   out.
	*/
	WaitForReadyTimeout string `json:"waitForReadyTimeout,omitempty" yaml:"waitForReadyTimeout,omitempty"`

	// A set of tags to apply to the Environment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The name of the Elastic Beanstalk Application Version
	   to use in deployment.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   A unique name for this Environment. This name is used
	   in the application URL
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	   to use in deployment
	*/
	PlatformArn string `json:"platformArn,omitempty" yaml:"platformArn,omitempty"`
}
