package apigee

type Instance struct {
	// Display name of the instance.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   IP range represents the customer-provided CIDR block of length 22 that will be used for
	   the Apigee instance creation. This optional range, if provided, should be freely
	   available as part of larger named range the customer has allocated to the Service
	   Networking peering. If this is not provided, Apigee will automatically request for any
	   available /22 CIDR block from Service Networking. The customer should use this CIDR block
	   for configuring their firewall needs to allow traffic from Apigee.
	   Input format: "a.b.c.d/22"
	*/
	IpRange string `json:"ipRange,omitempty" yaml:"ipRange,omitempty"`

	// Resource ID of the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The Apigee Organization associated with the Apigee instance,
	   in the format `organizations/{{org_name}}`.


	   - - -
	*/
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	/*
	   Optional. Customer accept list represents the list of projects (id/number) on customer
	   side that can privately connect to the service attachment. It is an optional field
	   which the customers can provide during the instance creation. By default, the customer
	   project associated with the Apigee organization will be included to the list.
	*/
	ConsumerAcceptLists []string `json:"consumerAcceptLists,omitempty" yaml:"consumerAcceptLists,omitempty"`

	// Description of the instance.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	   Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	*/
	DiskEncryptionKeyName string `json:"diskEncryptionKeyName,omitempty" yaml:"diskEncryptionKeyName,omitempty"`

	// Required. Compute Engine location where the instance resides.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The size of the CIDR block range that will be reserved by the instance. For valid values,
	   see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	*/
	PeeringCidrRange string `json:"peeringCidrRange,omitempty" yaml:"peeringCidrRange,omitempty"`
}
