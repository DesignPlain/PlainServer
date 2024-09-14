package compute

import types "libds/gcp/types"

type ServiceAttachment struct {
	/*
	   The connection preference to use for this service attachment. Valid
	   values include "ACCEPT_AUTOMATIC", "ACCEPT_MANUAL".
	*/
	ConnectionPreference string `json:"connectionPreference,omitempty" yaml:"connectionPreference,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. The name must be 1-63 characters long, and
	   comply with RFC1035. Specifically, the name must be 1-63 characters
	   long and match the regular expression `a-z?`
	   which means the first character must be a lowercase letter, and all
	   following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints.
	   If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified .
	   If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list.
	*/
	ReconcileConnections bool `json:"reconcileConnections,omitempty" yaml:"reconcileConnections,omitempty"`

	// URL of the region where the resource resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The URL of a forwarding rule that represents the service identified by
	   this service attachment.
	*/
	TargetService string `json:"targetService,omitempty" yaml:"targetService,omitempty"`

	/*
	   An array of projects that are allowed to connect to this service
	   attachment.
	   Structure is documented below.
	*/
	ConsumerAcceptLists []types.Compute_ServiceAttachmentConsumerAcceptList `json:"consumerAcceptLists,omitempty" yaml:"consumerAcceptLists,omitempty"`

	/*
	   An array of projects that are not allowed to connect to this service
	   attachment.
	*/
	ConsumerRejectLists []string `json:"consumerRejectLists,omitempty" yaml:"consumerRejectLists,omitempty"`

	/*
	   If specified, the domain name will be used during the integration between
	   the PSC connected endpoints and the Cloud DNS. For example, this is a
	   valid domain name: "p.mycompany.com.". Current max number of domain names
	   supported is 1.
	*/
	DomainNames []string `json:"domainNames,omitempty" yaml:"domainNames,omitempty"`

	/*
	   If true, enable the proxy protocol which is for supplying client TCP/IP
	   address data in TCP connections that traverse proxies on their way to
	   destination servers.


	   - - -
	*/
	EnableProxyProtocol bool `json:"enableProxyProtocol,omitempty" yaml:"enableProxyProtocol,omitempty"`

	// An array of subnets that is provided for NAT in this service attachment.
	NatSubnets []string `json:"natSubnets,omitempty" yaml:"natSubnets,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
