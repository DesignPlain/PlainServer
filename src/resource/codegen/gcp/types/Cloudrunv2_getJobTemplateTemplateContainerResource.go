package types

type Cloudrunv2_getJobTemplateTemplateContainerResource struct {
	// Only memory and CPU are supported. Use key 'cpu' for CPU limit and 'memory' for memory limit. Note: The only supported values for CPU are '1', '2', '4', and '8'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	Limits map[string]string `json:"limits,omitempty" yaml:"limits,omitempty"`
}
