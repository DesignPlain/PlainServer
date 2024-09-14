package compute

import types "libds/gcp/types"

type RegionNetworkEndpointGroup struct {
	/*
	   This field is only used for SERVERLESS NEGs.
	   Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
	   Structure is documented below.
	*/
	ServerlessDeployment types.Compute_RegionNetworkEndpointGroupServerlessDeployment `json:"serverlessDeployment,omitempty" yaml:"serverlessDeployment,omitempty"`

	/*
	   This field is only used for SERVERLESS NEGs.
	   Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	   Structure is documented below.
	*/
	AppEngine types.Compute_RegionNetworkEndpointGroupAppEngine `json:"appEngine,omitempty" yaml:"appEngine,omitempty"`

	/*
	   This field is only used for SERVERLESS NEGs.
	   Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	   Structure is documented below.
	*/
	CloudRun types.Compute_RegionNetworkEndpointGroupCloudRun `json:"cloudRun,omitempty" yaml:"cloudRun,omitempty"`

	/*
	   This field is only used for PSC and INTERNET NEGs.
	   The URL of the network to which all network endpoints in the NEG belong. Uses
	   "default" project network if unspecified.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A reference to the region where the regional NEGs reside.


	   - - -
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   This field is only used for PSC NEGs.
	   Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   This field is only used for SERVERLESS NEGs.
	   Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
	   Structure is documented below.
	*/
	CloudFunction types.Compute_RegionNetworkEndpointGroupCloudFunction `json:"cloudFunction,omitempty" yaml:"cloudFunction,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Type of network endpoints in this network endpoint group. Defaults to SERVERLESS.
	   Default value is `SERVERLESS`.
	   Possible values are: `SERVERLESS`, `PRIVATE_SERVICE_CONNECT`, `INTERNET_IP_PORT`, `INTERNET_FQDN_PORT`.
	*/
	NetworkEndpointType string `json:"networkEndpointType,omitempty" yaml:"networkEndpointType,omitempty"`

	/*
	   This field is only used for PSC and INTERNET NEGs.
	   The target service url used to set up private service connection to
	   a Google API or a PSC Producer Service Attachment.
	*/
	PscTargetService string `json:"pscTargetService,omitempty" yaml:"pscTargetService,omitempty"`
}
