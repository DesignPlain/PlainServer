package types

type Route53_RecordCidrRoutingPolicy struct {
	// The CIDR collection ID. See the `aws.route53.CidrCollection` resource for more details.
	CollectionId string `json:"collectionId,omitempty" yaml:"collectionId,omitempty"`

	// The CIDR collection location name. See the `aws.route53.CidrLocation` resource for more details. A `location_name` with an asterisk `"-"` can be used to create a default CIDR record. `collection_id` is still required for default record.
	LocationName string `json:"locationName,omitempty" yaml:"locationName,omitempty"`
}
