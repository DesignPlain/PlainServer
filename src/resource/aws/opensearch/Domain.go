package opensearch

import types "DesignSphere_Server/src/resource/aws/types"

type Domain struct {
	/*
	   Name of the domain.

	   The following arguments are optional:
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/opensearch-service/pricing/). Detailed below.
	EbsOptions types.Opensearch_DomainEbsOptions `json:"ebsOptions,omitempty" yaml:"ebsOptions,omitempty"`

	// Configuration to add Off Peak update options. ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html)). Detailed below.
	OffPeakWindowOptions types.Opensearch_DomainOffPeakWindowOptions `json:"offPeakWindowOptions,omitempty" yaml:"offPeakWindowOptions,omitempty"`

	// Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running OpenSearch 5.3 and later, Amazon OpenSearch takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions, OpenSearch takes daily automated snapshots.
	SnapshotOptions types.Opensearch_DomainSnapshotOptions `json:"snapshotOptions,omitempty" yaml:"snapshotOptions,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for the cluster of the domain. Detailed below.
	ClusterConfig types.Opensearch_DomainClusterConfig `json:"clusterConfig,omitempty" yaml:"clusterConfig,omitempty"`

	// Configuration block for authenticating dashboard with Cognito. Detailed below.
	CognitoOptions types.Opensearch_DomainCognitoOptions `json:"cognitoOptions,omitempty" yaml:"cognitoOptions,omitempty"`

	// Software update options for the domain. Detailed below.
	SoftwareUpdateOptions types.Opensearch_DomainSoftwareUpdateOptions `json:"softwareUpdateOptions,omitempty" yaml:"softwareUpdateOptions,omitempty"`

	// IAM policy document specifying the access policies for the domain.
	AccessPolicies string `json:"accessPolicies,omitempty" yaml:"accessPolicies,omitempty"`

	/*
	   Either `Elasticsearch_X.Y` or `OpenSearch_X.Y` to specify the engine version for the Amazon OpenSearch Service domain. For example, `OpenSearch_1.0` or `Elasticsearch_7.9`.
	   See [Creating and managing Amazon OpenSearch Service domains](http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).
	   Defaults to the lastest version of OpenSearch.
	*/
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	// Configuration block for encrypt at rest options. Only available for [certain instance types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html). Detailed below.
	EncryptAtRest types.Opensearch_DomainEncryptAtRest `json:"encryptAtRest,omitempty" yaml:"encryptAtRest,omitempty"`

	// Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
	LogPublishingOptions []types.Opensearch_DomainLogPublishingOption `json:"logPublishingOptions,omitempty" yaml:"logPublishingOptions,omitempty"`

	// Configuration block for node-to-node encryption options. Detailed below.
	NodeToNodeEncryption types.Opensearch_DomainNodeToNodeEncryption `json:"nodeToNodeEncryption,omitempty" yaml:"nodeToNodeEncryption,omitempty"`

	// Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your OpenSearch domain on every apply.
	AdvancedOptions map[string]string `json:"advancedOptions,omitempty" yaml:"advancedOptions,omitempty"`

	// Configuration block for the Auto-Tune options of the domain. Detailed below.
	AutoTuneOptions types.Opensearch_DomainAutoTuneOptions `json:"autoTuneOptions,omitempty" yaml:"autoTuneOptions,omitempty"`

	// Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html)). Detailed below.
	VpcOptions types.Opensearch_DomainVpcOptions `json:"vpcOptions,omitempty" yaml:"vpcOptions,omitempty"`

	// Configuration block for [fine-grained access control](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html). Detailed below.
	AdvancedSecurityOptions types.Opensearch_DomainAdvancedSecurityOptions `json:"advancedSecurityOptions,omitempty" yaml:"advancedSecurityOptions,omitempty"`

	// Configuration block for domain endpoint HTTP(S) related options. Detailed below.
	DomainEndpointOptions types.Opensearch_DomainDomainEndpointOptions `json:"domainEndpointOptions,omitempty" yaml:"domainEndpointOptions,omitempty"`
}
