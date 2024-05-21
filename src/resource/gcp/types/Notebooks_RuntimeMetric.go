package types

type Notebooks_RuntimeMetric struct {
	/*
	   (Output)
	   Contains runtime daemon metrics, such as OS and kernels and
	   sessions stats.
	*/
	SystemMetrics map[string]string `json:"systemMetrics,omitempty" yaml:"systemMetrics,omitempty"`
}
