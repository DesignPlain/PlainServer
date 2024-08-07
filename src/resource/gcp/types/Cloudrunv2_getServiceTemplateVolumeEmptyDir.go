package types

type Cloudrunv2_getServiceTemplateVolumeEmptyDir struct {
	// The different types of medium supported for EmptyDir. Default value: "MEMORY" Possible values: ["MEMORY"]
	Medium string `json:"medium,omitempty" yaml:"medium,omitempty"`

	// Limit on the storage usable by this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field's values are of the 'Quantity' k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir.
	SizeLimit string `json:"sizeLimit,omitempty" yaml:"sizeLimit,omitempty"`
}
