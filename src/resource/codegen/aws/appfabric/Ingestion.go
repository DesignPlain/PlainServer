package appfabric

type Ingestion struct {
	// Amazon Resource Name (ARN) of the app bundle to use for the request.
	AppBundleArn string `json:"appBundleArn,omitempty" yaml:"appBundleArn,omitempty"`

	// Ingestion type. Valid values are `auditLog`.
	IngestionType string `json:"ingestionType,omitempty" yaml:"ingestionType,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ID of the application tenant.
	TenantId string `json:"tenantId,omitempty" yaml:"tenantId,omitempty"`

	/*
	   Name of the application.
	   Refer to the AWS Documentation for the [list of valid values](https://docs.aws.amazon.com/appfabric/latest/api/API_CreateIngestion.html#appfabric-CreateIngestion-request-app)
	*/
	App string `json:"app,omitempty" yaml:"app,omitempty"`
}
