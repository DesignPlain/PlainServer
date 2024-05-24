package types

type Sagemaker_DomainDomainSettingsRStudioServerProDomainSettings struct {
	// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see `default_resource_spec` Block above.
	DefaultResourceSpec Sagemaker_DomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec `json:"defaultResourceSpec,omitempty" yaml:"defaultResourceSpec,omitempty"`

	// The ARN of the execution role for the RStudioServerPro Domain-level app.
	DomainExecutionRoleArn string `json:"domainExecutionRoleArn,omitempty" yaml:"domainExecutionRoleArn,omitempty"`

	// A URL pointing to an RStudio Connect server.
	RStudioConnectUrl string `json:"rStudioConnectUrl,omitempty" yaml:"rStudioConnectUrl,omitempty"`

	// A URL pointing to an RStudio Package Manager server.
	RStudioPackageManagerUrl string `json:"rStudioPackageManagerUrl,omitempty" yaml:"rStudioPackageManagerUrl,omitempty"`
}
