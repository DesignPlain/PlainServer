package types

type Eks_NodeGroupResourceAutoscalingGroup struct {
	// Name of the EC2 Launch Template. Conflicts with `id`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
