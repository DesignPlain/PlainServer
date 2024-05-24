package transfer

import types "DesignSphere_Server/src/resource/aws/types"

type Server struct {
	// Specifies the workflow details. See Workflow Details below.
	WorkflowDetails types.Transfer_ServerWorkflowDetails `json:"workflowDetails,omitempty" yaml:"workflowDetails,omitempty"`

	// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`. This option only applies to servers configured with a `SERVICE_MANAGED` `identity_provider_type`.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
	LoggingRole string `json:"loggingRole,omitempty" yaml:"loggingRole,omitempty"`

	// Specify a string to display when users connect to a server. This string is displayed after the user authenticates. The SFTP protocol does not support post-authentication display banners.
	PostAuthenticationLoginBanner string `json:"postAuthenticationLoginBanner,omitempty" yaml:"postAuthenticationLoginBanner,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
	InvocationRole string `json:"invocationRole,omitempty" yaml:"invocationRole,omitempty"`

	// Specify a string to display when users connect to a server. This string is displayed before the user authenticates.
	PreAuthenticationLoginBanner string `json:"preAuthenticationLoginBanner,omitempty" yaml:"preAuthenticationLoginBanner,omitempty"`

	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to `SFTP` . The available protocols are:
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	// A set of ARNs of destinations that will receive structured logs from the transfer server such as CloudWatch Log Group ARNs. If provided this enables the transfer server to emit structured logs to the specified locations.
	StructuredLogDestinations []string `json:"structuredLogDestinations,omitempty" yaml:"structuredLogDestinations,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when `protocols` is set to `FTPS`
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// The domain of the storage system that is used for file transfers. Valid values are: `S3` and `EFS`. The default value is `S3`.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// The ARN for a lambda function to use for the Identity provider.
	Function string `json:"function,omitempty" yaml:"function,omitempty"`

	/*
	   Specifies the name of the security policy that is attached to the server. Default value is: `TransferSecurityPolicy-2018-11`. The available values are:
	   - `TransferSecurityPolicy-2024-01`
	   - `TransferSecurityPolicy-2023-05`
	   - `TransferSecurityPolicy-2022-03`
	   - `TransferSecurityPolicy-2020-06`
	   - `TransferSecurityPolicy-2018-11`
	   - `TransferSecurityPolicy-FIPS-2024-01`
	   - `TransferSecurityPolicy-FIPS-2023-05`
	   - `TransferSecurityPolicy-FIPS-2020-06`
	   - `TransferSecurityPolicy-PQ-SSH-Experimental-2023-04`
	   - `TransferSecurityPolicy-PQ-SSH-FIPS-Experimental-2023-04`
	*/
	SecurityPolicyName string `json:"securityPolicyName,omitempty" yaml:"securityPolicyName,omitempty"`

	// The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
	EndpointDetails types.Transfer_ServerEndpointDetails `json:"endpointDetails,omitempty" yaml:"endpointDetails,omitempty"`

	// The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using `AWS_DIRECTORY_SERVICE` will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors. Use the `AWS_LAMBDA` value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the `function` argument.
	IdentityProviderType string `json:"identityProviderType,omitempty" yaml:"identityProviderType,omitempty"`

	// The protocol settings that are configured for your server.
	ProtocolDetails types.Transfer_ServerProtocolDetails `json:"protocolDetails,omitempty" yaml:"protocolDetails,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The directory service ID of the directory service you want to connect to with an `identity_provider_type` of `AWS_DIRECTORY_SERVICE`.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC` (or `VPC_ENDPOINT`), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`

	// RSA, ECDSA, or ED25519 private key (e.g., as generated by the `ssh-keygen -t rsa -b 2048 -N "" -m PEM -f my-new-server-key`, `ssh-keygen -t ecdsa -b 256 -N "" -m PEM -f my-new-server-key` or `ssh-keygen -t ed25519 -N "" -f my-new-server-key` commands).
	HostKey string `json:"hostKey,omitempty" yaml:"hostKey,omitempty"`
}
