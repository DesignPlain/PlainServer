package elasticbeanstalk

type ApplicationVersion struct {
	// On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	// S3 object that is the Application Version source bundle.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   Unique name for the this Application Version.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Pre-processes and validates the environment manifest (env.yaml ) and configuration files (-.config files in the .ebextensions folder) in the source bundle. Validating configuration files can identify issues prior to deploying the application version to an environment. You must turn processing on for application versions that you create using AWS CodeBuild or AWS CodeCommit. For application versions built from a source bundle in Amazon S3, processing is optional. It validates Elastic Beanstalk configuration files. It doesn’t validate your application’s configuration files, like proxy server or Docker configuration.
	Process bool `json:"process,omitempty" yaml:"process,omitempty"`

	// Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of the Beanstalk Application the version is associated with.
	Application string `json:"application,omitempty" yaml:"application,omitempty"`

	// S3 bucket that contains the Application Version source bundle.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Short description of the Application Version.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
