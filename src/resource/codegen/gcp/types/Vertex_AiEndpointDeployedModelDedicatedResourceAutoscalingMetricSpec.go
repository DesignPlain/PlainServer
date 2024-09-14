package types

type Vertex_AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpec struct {
	/*
	   (Output)
	   The resource metric name. Supported metrics: - For Online Prediction: - `aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle` - `aiplatform.googleapis.com/prediction/online/cpu/utilization`
	*/
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	/*
	   (Output)
	   The target resource utilization in percentage (1%!)(MISSING) for the given metric; once the real usage deviates from the target by a certain percentage, the machine replicas change. The default value is 60 (representing 60%!!(MISSING))(MISSING) if not provided.
	*/
	Target int `json:"target,omitempty" yaml:"target,omitempty"`
}
