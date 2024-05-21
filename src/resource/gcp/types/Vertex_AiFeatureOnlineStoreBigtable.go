package types

type Vertex_AiFeatureOnlineStoreBigtable struct {
	/*
	   Autoscaling config applied to Bigtable Instance.
	   Structure is documented below.
	*/
	AutoScaling Vertex_AiFeatureOnlineStoreBigtableAutoScaling `json:"autoScaling,omitempty" yaml:"autoScaling,omitempty"`
}
