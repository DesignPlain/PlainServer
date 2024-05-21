package cloudbuild

import types "DesignSphere_Server/src/resource/gcp/types"

type BitbucketServerConfig struct {
	/*
	   Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed.
	   Changing this field will result in deleting/ recreating the resource.
	*/
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// The ID to use for the BitbucketServerConfig, which will become the final component of the BitbucketServerConfig's resource name.
	ConfigId string `json:"configId,omitempty" yaml:"configId,omitempty"`

	/*
	   Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed.
	   If you need to change it, please create another BitbucketServerConfig.
	*/
	HostUri string `json:"hostUri,omitempty" yaml:"hostUri,omitempty"`

	// The location of this bitbucket server config.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Secret Manager secrets needed by the config.
	   Structure is documented below.
	*/
	Secrets types.Cloudbuild_BitbucketServerConfigSecrets `json:"secrets,omitempty" yaml:"secrets,omitempty"`

	// SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
	SslCa string `json:"sslCa,omitempty" yaml:"sslCa,omitempty"`

	/*
	   Connected Bitbucket Server repositories for this config.
	   Structure is documented below.
	*/
	ConnectedRepositories []types.Cloudbuild_BitbucketServerConfigConnectedRepository `json:"connectedRepositories,omitempty" yaml:"connectedRepositories,omitempty"`

	/*
	   The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection.
	   This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty,
	   no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format
	   projects/{project}/global/networks/{network}, where {project} is a project number or id and {network} is the name of a VPC network in the project.
	*/
	PeeredNetwork string `json:"peeredNetwork,omitempty" yaml:"peeredNetwork,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Username of the account Cloud Build will use on Bitbucket Server.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
