package elasticbeanstalk

type ApplicationVersion struct {
	// Name of the Beanstalk Application the version is associated with.
	Application string `json:"application,omitempty" yaml:"application,omitempty"`

	// S3 bucket that contains the Application Version source bundle.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Short description of the Application Version.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// On delete, force an Application Version to be deleted when it may be in use by multiple Elastic Beanstalk Environments.
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	// S3 object that is the Application Version source bundle.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   Unique name for the this Application Version.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of tags for the Elastic Beanstalk Application Version. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
