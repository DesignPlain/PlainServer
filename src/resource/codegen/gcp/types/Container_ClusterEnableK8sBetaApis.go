package types

type Container_ClusterEnableK8sBetaApis struct {
	// Enabled Kubernetes Beta APIs. To list a Beta API resource, use the representation {group}/{version}/{resource}. The version must be a Beta version. Note that you cannot disable beta APIs that are already enabled on a cluster without recreating it. See the [Configure beta APIs](https://cloud.google.com/kubernetes-engine/docs/how-to/use-beta-apis#configure-beta-apis) for more information.
	EnabledApis []string `json:"enabledApis,omitempty" yaml:"enabledApis,omitempty"`
}
