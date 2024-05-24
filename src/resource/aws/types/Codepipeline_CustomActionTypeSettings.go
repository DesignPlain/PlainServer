package types

type Codepipeline_CustomActionTypeSettings struct {
	// The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system.
	EntityUrlTemplate string `json:"entityUrlTemplate,omitempty" yaml:"entityUrlTemplate,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system.
	ExecutionUrlTemplate string `json:"executionUrlTemplate,omitempty" yaml:"executionUrlTemplate,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	RevisionUrlTemplate string `json:"revisionUrlTemplate,omitempty" yaml:"revisionUrlTemplate,omitempty"`

	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	ThirdPartyConfigurationUrl string `json:"thirdPartyConfigurationUrl,omitempty" yaml:"thirdPartyConfigurationUrl,omitempty"`
}
