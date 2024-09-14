package iam

type ServiceSpecificCredential struct {
	// The name of the AWS service that is to be associated with the credentials. The service you specify here is the only service that can be accessed using these credentials.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	// The status to be assigned to the service-specific credential. Valid values are `Active` and `Inactive`. Default value is `Active`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The name of the IAM user that is to be associated with the credentials. The new service-specific credentials have the same permissions as the associated user except that they can be used only to access the specified service.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`
}
