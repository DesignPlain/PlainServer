package types

type Appengine_StandardAppVersionVpcAccessConnector struct {
	// The egress setting for the connector, controlling what traffic is diverted through it.
	EgressSetting string `json:"egressSetting,omitempty" yaml:"egressSetting,omitempty"`

	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
