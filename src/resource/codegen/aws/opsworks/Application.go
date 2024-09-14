package opsworks

import types "libds/aws/types"

type Application struct {
	// SCM configuration of the app as described below.
	AppSources []types.Opsworks_ApplicationAppSource `json:"appSources,omitempty" yaml:"appSources,omitempty"`

	// A description of the app.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Opsworks application type. One of `aws-flow-ruby`, `java`, `rails`, `php`, `nodejs`, `static` or `other`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The database name.
	DataSourceDatabaseName string `json:"dataSourceDatabaseName,omitempty" yaml:"dataSourceDatabaseName,omitempty"`

	// A list of virtual host alias.
	Domains []string `json:"domains,omitempty" yaml:"domains,omitempty"`

	// Whether to enable SSL for the app. This must be set in order to let `ssl_configuration.private_key`, `ssl_configuration.certificate` and `ssl_configuration.chain` take effect.
	EnableSsl bool `json:"enableSsl,omitempty" yaml:"enableSsl,omitempty"`

	// Object to define environment variables.  Object is described below.
	Environments []types.Opsworks_ApplicationEnvironment `json:"environments,omitempty" yaml:"environments,omitempty"`

	// Run bundle install when deploying for application of type `rails`.
	AutoBundleOnDeploy string `json:"autoBundleOnDeploy,omitempty" yaml:"autoBundleOnDeploy,omitempty"`

	// The data source's type one of `AutoSelectOpsworksMysqlInstance`, `OpsworksMysqlInstance`, or `RdsDbInstance`.
	DataSourceType string `json:"dataSourceType,omitempty" yaml:"dataSourceType,omitempty"`

	// A human-readable name for the application.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The SSL configuration of the app. Object is described below.
	SslConfigurations []types.Opsworks_ApplicationSslConfiguration `json:"sslConfigurations,omitempty" yaml:"sslConfigurations,omitempty"`

	// ID of the stack the application will belong to.
	StackId string `json:"stackId,omitempty" yaml:"stackId,omitempty"`

	// Specify activity and workflow workers for your app using the aws-flow gem.
	AwsFlowRubySettings string `json:"awsFlowRubySettings,omitempty" yaml:"awsFlowRubySettings,omitempty"`

	// The data source's ARN.
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	// Subfolder for the document root for application of type `rails`.
	DocumentRoot string `json:"documentRoot,omitempty" yaml:"documentRoot,omitempty"`

	// The name of the Rails environment for application of type `rails`.
	RailsEnv string `json:"railsEnv,omitempty" yaml:"railsEnv,omitempty"`

	// A short, machine-readable name for the application. This can only be defined on resource creation and ignored on resource update.
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`
}
