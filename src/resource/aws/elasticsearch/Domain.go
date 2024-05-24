package elasticsearch

import types "DesignSphere_Server/src/resource/aws/types"

type Domain struct {
	// Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)). Detailed below.
	VpcOptions types.Elasticsearch_DomainVpcOptions `json:"vpcOptions,omitempty" yaml:"vpcOptions,omitempty"`

	// IAM policy document specifying the access policies for the domain.
	AccessPolicies string `json:"accessPolicies,omitempty" yaml:"accessPolicies,omitempty"`

	// Configuration block for [fine-grained access control](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/fgac.html). Detailed below.
	AdvancedSecurityOptions types.Elasticsearch_DomainAdvancedSecurityOptions `json:"advancedSecurityOptions,omitempty" yaml:"advancedSecurityOptions,omitempty"`

	// Configuration block for the Auto-Tune options of the domain. Detailed below.
	AutoTuneOptions types.Elasticsearch_DomainAutoTuneOptions `json:"autoTuneOptions,omitempty" yaml:"autoTuneOptions,omitempty"`

	// Configuration block for domain endpoint HTTP(S) related options. Detailed below.
	DomainEndpointOptions types.Elasticsearch_DomainDomainEndpointOptions `json:"domainEndpointOptions,omitempty" yaml:"domainEndpointOptions,omitempty"`

	// Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running Elasticsearch 5.3 and later, Amazon ES takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions of Elasticsearch, Amazon ES takes daily automated snapshots.
	SnapshotOptions types.Elasticsearch_DomainSnapshotOptions `json:"snapshotOptions,omitempty" yaml:"snapshotOptions,omitempty"`

	// Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). Detailed below.
	EbsOptions types.Elasticsearch_DomainEbsOptions `json:"ebsOptions,omitempty" yaml:"ebsOptions,omitempty"`

	// Configuration block for encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). Detailed below.
	EncryptAtRest types.Elasticsearch_DomainEncryptAtRest `json:"encryptAtRest,omitempty" yaml:"encryptAtRest,omitempty"`

	// Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
	LogPublishingOptions []types.Elasticsearch_DomainLogPublishingOption `json:"logPublishingOptions,omitempty" yaml:"logPublishingOptions,omitempty"`

	// Configuration block for the cluster of the domain. Detailed below.
	ClusterConfig types.Elasticsearch_DomainClusterConfig `json:"clusterConfig,omitempty" yaml:"clusterConfig,omitempty"`

	// Configuration block for authenticating Kibana with Cognito. Detailed below.
	CognitoOptions types.Elasticsearch_DomainCognitoOptions `json:"cognitoOptions,omitempty" yaml:"cognitoOptions,omitempty"`

	// Configuration block for node-to-node encryption options. Detailed below.
	NodeToNodeEncryption types.Elasticsearch_DomainNodeToNodeEncryption `json:"nodeToNodeEncryption,omitempty" yaml:"nodeToNodeEncryption,omitempty"`

	// Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your Elasticsearch domain on every apply.
	AdvancedOptions map[string]string `json:"advancedOptions,omitempty" yaml:"advancedOptions,omitempty"`

	/*
	   Name of the domain.

	   The following arguments are optional:
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Version of Elasticsearch to deploy. Defaults to `1.5`.
	ElasticsearchVersion string `json:"elasticsearchVersion,omitempty" yaml:"elasticsearchVersion,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
