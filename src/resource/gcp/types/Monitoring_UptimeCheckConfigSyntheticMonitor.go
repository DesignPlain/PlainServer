package types

type Monitoring_UptimeCheckConfigSyntheticMonitor struct {
	/*
	   Target a Synthetic Monitor GCFv2 Instance
	   Structure is documented below.


	   <a name="nested_cloud_function_v2"></a>The `cloud_function_v2` block supports:
	*/
	CloudFunctionV2 Monitoring_UptimeCheckConfigSyntheticMonitorCloudFunctionV2 `json:"cloudFunctionV2,omitempty" yaml:"cloudFunctionV2,omitempty"`
}
