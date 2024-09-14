package redshift

type Partner struct {
	// The Amazon Web Services account ID that owns the cluster.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The cluster identifier of the cluster that receives data from the partner.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The name of the database that receives data from the partner.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The name of the partner that is authorized to send data.
	PartnerName string `json:"partnerName,omitempty" yaml:"partnerName,omitempty"`
}
