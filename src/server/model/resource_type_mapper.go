package model

import (
	"DesignSphere_Server/src/resource/aws/accessanalyzer"
	"DesignSphere_Server/src/resource/aws/account"
	"DesignSphere_Server/src/resource/aws/acm"
	"DesignSphere_Server/src/resource/aws/acmpca"
	"DesignSphere_Server/src/resource/aws/alb"
	"DesignSphere_Server/src/resource/aws/amp"
	"DesignSphere_Server/src/resource/aws/amplify"
	aws_apigateway "DesignSphere_Server/src/resource/aws/apigateway"
	"DesignSphere_Server/src/resource/aws/apigatewayv2"
	"DesignSphere_Server/src/resource/aws/appautoscaling"
	"DesignSphere_Server/src/resource/aws/appconfig"
	"DesignSphere_Server/src/resource/aws/appflow"
	"DesignSphere_Server/src/resource/aws/appintegrations"
	"DesignSphere_Server/src/resource/aws/applicationinsights"
	"DesignSphere_Server/src/resource/aws/appmesh"
	"DesignSphere_Server/src/resource/aws/apprunner"
	"DesignSphere_Server/src/resource/aws/appstream"
	"DesignSphere_Server/src/resource/aws/appsync"
	"DesignSphere_Server/src/resource/aws/athena"
	"DesignSphere_Server/src/resource/aws/auditmanager"
	"DesignSphere_Server/src/resource/aws/autoscaling"
	"DesignSphere_Server/src/resource/aws/autoscalingplans"
	"DesignSphere_Server/src/resource/aws/backup"
	"DesignSphere_Server/src/resource/aws/batch"
	"DesignSphere_Server/src/resource/aws/bedrock"
	"DesignSphere_Server/src/resource/aws/bedrockmodel"
	"DesignSphere_Server/src/resource/aws/budgets"
	"DesignSphere_Server/src/resource/aws/cfg"
	"DesignSphere_Server/src/resource/aws/chime"
	"DesignSphere_Server/src/resource/aws/chimesdkmediapipelines"
	"DesignSphere_Server/src/resource/aws/cleanrooms"
	"DesignSphere_Server/src/resource/aws/cloud9"
	"DesignSphere_Server/src/resource/aws/cloudcontrol"
	"DesignSphere_Server/src/resource/aws/cloudformation"
	"DesignSphere_Server/src/resource/aws/cloudfront"
	"DesignSphere_Server/src/resource/aws/cloudhsmv2"
	"DesignSphere_Server/src/resource/aws/cloudsearch"
	"DesignSphere_Server/src/resource/aws/cloudtrail"
	"DesignSphere_Server/src/resource/aws/cloudwatch"
	"DesignSphere_Server/src/resource/aws/codeartifact"
	"DesignSphere_Server/src/resource/aws/codebuild"
	"DesignSphere_Server/src/resource/aws/codecatalyst"
	"DesignSphere_Server/src/resource/aws/codecommit"
	"DesignSphere_Server/src/resource/aws/codedeploy"
	"DesignSphere_Server/src/resource/aws/codeguruprofiler"
	"DesignSphere_Server/src/resource/aws/codegurureviewer"
	"DesignSphere_Server/src/resource/aws/codepipeline"
	"DesignSphere_Server/src/resource/aws/codestarconnections"
	"DesignSphere_Server/src/resource/aws/codestarnotifications"
	"DesignSphere_Server/src/resource/aws/cognito"
	"DesignSphere_Server/src/resource/aws/comprehend"
	"DesignSphere_Server/src/resource/aws/connect"
	"DesignSphere_Server/src/resource/aws/controltower"
	"DesignSphere_Server/src/resource/aws/costexplorer"
	"DesignSphere_Server/src/resource/aws/cur"
	"DesignSphere_Server/src/resource/aws/customerprofiles"
	"DesignSphere_Server/src/resource/aws/dataexchange"
	"DesignSphere_Server/src/resource/aws/datapipeline"
	"DesignSphere_Server/src/resource/aws/datasync"
	"DesignSphere_Server/src/resource/aws/dax"
	"DesignSphere_Server/src/resource/aws/detective"
	"DesignSphere_Server/src/resource/aws/devicefarm"
	"DesignSphere_Server/src/resource/aws/directconnect"
	"DesignSphere_Server/src/resource/aws/directoryservice"
	"DesignSphere_Server/src/resource/aws/dlm"
	"DesignSphere_Server/src/resource/aws/dms"
	"DesignSphere_Server/src/resource/aws/docdb"
	"DesignSphere_Server/src/resource/aws/dynamodb"
	"DesignSphere_Server/src/resource/aws/ebs"
	"DesignSphere_Server/src/resource/aws/ec2"
	"DesignSphere_Server/src/resource/aws/ec2clientvpn"
	"DesignSphere_Server/src/resource/aws/ec2transitgateway"
	"DesignSphere_Server/src/resource/aws/ecr"
	"DesignSphere_Server/src/resource/aws/ecrpublic"
	"DesignSphere_Server/src/resource/aws/ecs"
	"DesignSphere_Server/src/resource/aws/efs"
	"DesignSphere_Server/src/resource/aws/eks"
	"DesignSphere_Server/src/resource/aws/elasticache"
	"DesignSphere_Server/src/resource/aws/elasticbeanstalk"
	"DesignSphere_Server/src/resource/aws/elasticsearch"
	"DesignSphere_Server/src/resource/aws/elastictranscoder"
	"DesignSphere_Server/src/resource/aws/elb"
	"DesignSphere_Server/src/resource/aws/emr"
	"DesignSphere_Server/src/resource/aws/emrcontainers"
	"DesignSphere_Server/src/resource/aws/emrserverless"
	"DesignSphere_Server/src/resource/aws/evidently"
	"DesignSphere_Server/src/resource/aws/finspace"
	"DesignSphere_Server/src/resource/aws/fis"
	"DesignSphere_Server/src/resource/aws/fms"
	"DesignSphere_Server/src/resource/aws/fsx"
	"DesignSphere_Server/src/resource/aws/gamelift"
	"DesignSphere_Server/src/resource/aws/glacier"
	"DesignSphere_Server/src/resource/aws/globalaccelerator"
	"DesignSphere_Server/src/resource/aws/glue"
	"DesignSphere_Server/src/resource/aws/grafana"
	"DesignSphere_Server/src/resource/aws/guardduty"
	aws_iam "DesignSphere_Server/src/resource/aws/iam"
	"DesignSphere_Server/src/resource/aws/identitystore"
	"DesignSphere_Server/src/resource/aws/imagebuilder"
	"DesignSphere_Server/src/resource/aws/inspector"
	"DesignSphere_Server/src/resource/aws/inspector2"
	"DesignSphere_Server/src/resource/aws/iot"
	"DesignSphere_Server/src/resource/aws/ivs"
	"DesignSphere_Server/src/resource/aws/ivschat"
	"DesignSphere_Server/src/resource/aws/kendra"
	"DesignSphere_Server/src/resource/aws/keyspaces"
	"DesignSphere_Server/src/resource/aws/kinesis"
	"DesignSphere_Server/src/resource/aws/kinesisanalyticsv2"
	aws_kms "DesignSphere_Server/src/resource/aws/kms"
	"DesignSphere_Server/src/resource/aws/lakeformation"
	"DesignSphere_Server/src/resource/aws/lambda"
	"DesignSphere_Server/src/resource/aws/lb"
	"DesignSphere_Server/src/resource/aws/lex"
	"DesignSphere_Server/src/resource/aws/licensemanager"
	"DesignSphere_Server/src/resource/aws/lightsail"
	"DesignSphere_Server/src/resource/aws/location"
	"DesignSphere_Server/src/resource/aws/macie"
	"DesignSphere_Server/src/resource/aws/macie2"
	"DesignSphere_Server/src/resource/aws/mediaconvert"
	"DesignSphere_Server/src/resource/aws/medialive"
	"DesignSphere_Server/src/resource/aws/mediapackage"
	"DesignSphere_Server/src/resource/aws/mediastore"
	"DesignSphere_Server/src/resource/aws/memorydb"
	"DesignSphere_Server/src/resource/aws/mq"
	"DesignSphere_Server/src/resource/aws/msk"
	"DesignSphere_Server/src/resource/aws/mskconnect"
	"DesignSphere_Server/src/resource/aws/mwaa"
	"DesignSphere_Server/src/resource/aws/neptune"
	"DesignSphere_Server/src/resource/aws/networkfirewall"
	"DesignSphere_Server/src/resource/aws/networkmanager"
	"DesignSphere_Server/src/resource/aws/oam"
	"DesignSphere_Server/src/resource/aws/opensearch"
	"DesignSphere_Server/src/resource/aws/opensearchingest"
	"DesignSphere_Server/src/resource/aws/opsworks"
	aws_organizations "DesignSphere_Server/src/resource/aws/organizations"
	"DesignSphere_Server/src/resource/aws/pinpoint"
	"DesignSphere_Server/src/resource/aws/pipes"
	"DesignSphere_Server/src/resource/aws/qldb"
	"DesignSphere_Server/src/resource/aws/quicksight"
	"DesignSphere_Server/src/resource/aws/ram"
	"DesignSphere_Server/src/resource/aws/rbin"
	"DesignSphere_Server/src/resource/aws/rds"
	"DesignSphere_Server/src/resource/aws/redshift"
	"DesignSphere_Server/src/resource/aws/redshiftdata"
	"DesignSphere_Server/src/resource/aws/redshiftserverless"
	"DesignSphere_Server/src/resource/aws/rekognition"
	"DesignSphere_Server/src/resource/aws/resourceexplorer"
	"DesignSphere_Server/src/resource/aws/resourcegroups"
	"DesignSphere_Server/src/resource/aws/rolesanywhere"
	"DesignSphere_Server/src/resource/aws/route53"
	"DesignSphere_Server/src/resource/aws/route53domains"
	"DesignSphere_Server/src/resource/aws/route53recoverycontrol"
	"DesignSphere_Server/src/resource/aws/route53recoveryreadiness"
	"DesignSphere_Server/src/resource/aws/rum"
	"DesignSphere_Server/src/resource/aws/s3"
	"DesignSphere_Server/src/resource/aws/s3control"
	"DesignSphere_Server/src/resource/aws/s3outposts"
	"DesignSphere_Server/src/resource/aws/sagemaker"
	"DesignSphere_Server/src/resource/aws/scheduler"
	"DesignSphere_Server/src/resource/aws/schemas"
	"DesignSphere_Server/src/resource/aws/secretsmanager"
	"DesignSphere_Server/src/resource/aws/securityhub"
	"DesignSphere_Server/src/resource/aws/securitylake"
	"DesignSphere_Server/src/resource/aws/serverlessrepository"
	"DesignSphere_Server/src/resource/aws/servicecatalog"
	"DesignSphere_Server/src/resource/aws/servicediscovery"
	"DesignSphere_Server/src/resource/aws/servicequotas"
	"DesignSphere_Server/src/resource/aws/ses"
	"DesignSphere_Server/src/resource/aws/sesv2"
	"DesignSphere_Server/src/resource/aws/sfn"
	"DesignSphere_Server/src/resource/aws/shield"
	"DesignSphere_Server/src/resource/aws/signer"
	"DesignSphere_Server/src/resource/aws/simpledb"
	"DesignSphere_Server/src/resource/aws/sns"
	"DesignSphere_Server/src/resource/aws/sqs"
	"DesignSphere_Server/src/resource/aws/ssm"
	"DesignSphere_Server/src/resource/aws/ssmcontacts"
	"DesignSphere_Server/src/resource/aws/ssmincidents"
	"DesignSphere_Server/src/resource/aws/ssoadmin"
	"DesignSphere_Server/src/resource/aws/storagegateway"
	"DesignSphere_Server/src/resource/aws/swf"
	"DesignSphere_Server/src/resource/aws/synthetics"
	"DesignSphere_Server/src/resource/aws/timestreamwrite"
	"DesignSphere_Server/src/resource/aws/transcribe"
	"DesignSphere_Server/src/resource/aws/transfer"
	"DesignSphere_Server/src/resource/aws/verifiedaccess"
	"DesignSphere_Server/src/resource/aws/verifiedpermissions"
	"DesignSphere_Server/src/resource/aws/vpc"
	"DesignSphere_Server/src/resource/aws/vpclattice"
	"DesignSphere_Server/src/resource/aws/waf"
	"DesignSphere_Server/src/resource/aws/wafregional"
	"DesignSphere_Server/src/resource/aws/wafv2"
	"DesignSphere_Server/src/resource/aws/worklink"
	"DesignSphere_Server/src/resource/aws/workspaces"
	"DesignSphere_Server/src/resource/aws/xray"
	"DesignSphere_Server/src/resource/gcp/accesscontextmanager"
	"DesignSphere_Server/src/resource/gcp/activedirectory"
	"DesignSphere_Server/src/resource/gcp/alloydb"
	"DesignSphere_Server/src/resource/gcp/apigateway"
	"DesignSphere_Server/src/resource/gcp/apigee"
	"DesignSphere_Server/src/resource/gcp/appengine"
	"DesignSphere_Server/src/resource/gcp/artifactregistry"
	"DesignSphere_Server/src/resource/gcp/assuredworkloads"
	"DesignSphere_Server/src/resource/gcp/backupdisasterrecovery"
	"DesignSphere_Server/src/resource/gcp/beyondcorp"
	"DesignSphere_Server/src/resource/gcp/biglake"
	"DesignSphere_Server/src/resource/gcp/bigquery"
	"DesignSphere_Server/src/resource/gcp/bigqueryanalyticshub"
	"DesignSphere_Server/src/resource/gcp/bigquerydatapolicy"
	"DesignSphere_Server/src/resource/gcp/bigtable"
	"DesignSphere_Server/src/resource/gcp/billing"
	"DesignSphere_Server/src/resource/gcp/binaryauthorization"
	"DesignSphere_Server/src/resource/gcp/blockchainnodeengine"
	"DesignSphere_Server/src/resource/gcp/certificateauthority"
	"DesignSphere_Server/src/resource/gcp/certificatemanager"
	"DesignSphere_Server/src/resource/gcp/cloudasset"
	"DesignSphere_Server/src/resource/gcp/cloudbuild"
	"DesignSphere_Server/src/resource/gcp/cloudbuildv2"
	"DesignSphere_Server/src/resource/gcp/clouddeploy"
	"DesignSphere_Server/src/resource/gcp/clouddomains"
	"DesignSphere_Server/src/resource/gcp/cloudfunctions"
	"DesignSphere_Server/src/resource/gcp/cloudfunctionsv2"
	"DesignSphere_Server/src/resource/gcp/cloudidentity"
	"DesignSphere_Server/src/resource/gcp/cloudids"
	"DesignSphere_Server/src/resource/gcp/cloudrun"
	"DesignSphere_Server/src/resource/gcp/cloudrunv2"
	"DesignSphere_Server/src/resource/gcp/cloudscheduler"
	"DesignSphere_Server/src/resource/gcp/cloudtasks"
	"DesignSphere_Server/src/resource/gcp/composer"
	"DesignSphere_Server/src/resource/gcp/compute"
	"DesignSphere_Server/src/resource/gcp/container"
	"DesignSphere_Server/src/resource/gcp/containeranalysis"
	"DesignSphere_Server/src/resource/gcp/databasemigrationservice"
	"DesignSphere_Server/src/resource/gcp/datacatalog"
	"DesignSphere_Server/src/resource/gcp/dataflow"
	"DesignSphere_Server/src/resource/gcp/dataform"
	"DesignSphere_Server/src/resource/gcp/datafusion"
	"DesignSphere_Server/src/resource/gcp/dataloss"
	"DesignSphere_Server/src/resource/gcp/dataplex"
	"DesignSphere_Server/src/resource/gcp/dataproc"
	"DesignSphere_Server/src/resource/gcp/datastore"
	"DesignSphere_Server/src/resource/gcp/datastream"
	"DesignSphere_Server/src/resource/gcp/deploymentmanager"
	"DesignSphere_Server/src/resource/gcp/diagflow"
	"DesignSphere_Server/src/resource/gcp/discoveryengine"
	"DesignSphere_Server/src/resource/gcp/dns"
	"DesignSphere_Server/src/resource/gcp/edgecontainer"
	"DesignSphere_Server/src/resource/gcp/edgenetwork"
	"DesignSphere_Server/src/resource/gcp/endpoints"
	"DesignSphere_Server/src/resource/gcp/essentialcontacts"
	"DesignSphere_Server/src/resource/gcp/eventarc"
	"DesignSphere_Server/src/resource/gcp/filestore"
	"DesignSphere_Server/src/resource/gcp/firebase"
	"DesignSphere_Server/src/resource/gcp/firebaserules"
	"DesignSphere_Server/src/resource/gcp/firestore"
	"DesignSphere_Server/src/resource/gcp/folder"
	"DesignSphere_Server/src/resource/gcp/gkebackup"
	"DesignSphere_Server/src/resource/gcp/gkehub"
	"DesignSphere_Server/src/resource/gcp/gkeonprem"
	"DesignSphere_Server/src/resource/gcp/healthcare"
	"DesignSphere_Server/src/resource/gcp/iam"
	"DesignSphere_Server/src/resource/gcp/iap"
	"DesignSphere_Server/src/resource/gcp/identityplatform"
	"DesignSphere_Server/src/resource/gcp/integrationconnectors"
	"DesignSphere_Server/src/resource/gcp/kms"
	"DesignSphere_Server/src/resource/gcp/logging"
	"DesignSphere_Server/src/resource/gcp/looker"
	"DesignSphere_Server/src/resource/gcp/memcache"
	"DesignSphere_Server/src/resource/gcp/migrationcenter"
	"DesignSphere_Server/src/resource/gcp/ml"
	"DesignSphere_Server/src/resource/gcp/monitoring"
	"DesignSphere_Server/src/resource/gcp/netapp"
	"DesignSphere_Server/src/resource/gcp/networkconnectivity"
	"DesignSphere_Server/src/resource/gcp/networkmanagement"
	"DesignSphere_Server/src/resource/gcp/networksecurity"
	"DesignSphere_Server/src/resource/gcp/networkservices"
	"DesignSphere_Server/src/resource/gcp/notebooks"
	"DesignSphere_Server/src/resource/gcp/organizations"
	"DesignSphere_Server/src/resource/gcp/orgpolicy"
	"DesignSphere_Server/src/resource/gcp/osconfig"
	"DesignSphere_Server/src/resource/gcp/oslogin"
	"DesignSphere_Server/src/resource/gcp/projects"
	"DesignSphere_Server/src/resource/gcp/pubsub"
	"DesignSphere_Server/src/resource/gcp/recaptcha"
	"DesignSphere_Server/src/resource/gcp/redis"
	"DesignSphere_Server/src/resource/gcp/resourcemanager"

	"DesignSphere_Server/src/resource/gcp/runtimeconfig"
	"DesignSphere_Server/src/resource/gcp/secretmanager"
	"DesignSphere_Server/src/resource/gcp/securesourcemanager"
	"DesignSphere_Server/src/resource/gcp/securitycenter"
	"DesignSphere_Server/src/resource/gcp/securityposture"
	"DesignSphere_Server/src/resource/gcp/serviceaccount"
	"DesignSphere_Server/src/resource/gcp/servicedirectory"
	"DesignSphere_Server/src/resource/gcp/servicenetworking"
	"DesignSphere_Server/src/resource/gcp/serviceusage"
	"DesignSphere_Server/src/resource/gcp/sourcerepo"
	"DesignSphere_Server/src/resource/gcp/spanner"
	"DesignSphere_Server/src/resource/gcp/sql"
	"DesignSphere_Server/src/resource/gcp/storage"
	"DesignSphere_Server/src/resource/gcp/tags"
	"DesignSphere_Server/src/resource/gcp/tpu"
	"DesignSphere_Server/src/resource/gcp/vertex"
	"DesignSphere_Server/src/resource/gcp/vmwareengine"
	"DesignSphere_Server/src/resource/gcp/vpcaccess"
	"DesignSphere_Server/src/resource/gcp/workbench"
	"DesignSphere_Server/src/resource/gcp/workflows"
	"DesignSphere_Server/src/resource/gcp/workstations"
)

// TODO: This file was partially autogenerated and cleaned up manually,
//
//	will need to generate it properly
var ResourceTypeMap = map[ResourceType]func() Any{
	GLUE_SECURITYCONFIGURATION: func() Any {
		return &glue.SecurityConfiguration{}
	},
	GRAFANA_LICENSEASSOCIATION: func() Any {
		return &grafana.LicenseAssociation{}
	},
	KINESISANALYTICSV2_APPLICATIONSNAPSHOT: func() Any {
		return &kinesisanalyticsv2.ApplicationSnapshot{}
	},
	ATHENA_DATABASE: func() Any {
		return &athena.Database{}
	},
	QUICKSIGHT_TEMPLATEALIAS: func() Any {
		return &quicksight.TemplateAlias{}
	},
	LB_LOADBALANCER: func() Any {
		return &lb.LoadBalancer{}
	},
	OAM_SINK: func() Any {
		return &oam.Sink{}
	},
	OPSWORKS_PHPAPPLAYER: func() Any {
		return &opsworks.PhpAppLayer{}
	},
	RDS_INSTANCE: func() Any {
		return &rds.Instance{}
	},
	S3_BUCKETREQUESTPAYMENTCONFIGURATIONV2: func() Any {
		return &s3.BucketRequestPaymentConfigurationV2{}
	},
	APPSTREAM_STACK: func() Any {
		return &appstream.Stack{}
	},
	EC2_DEFAULTNETWORKACL: func() Any {
		return &ec2.DefaultNetworkAcl{}
	},
	EC2CLIENTVPN_ROUTE: func() Any {
		return &ec2clientvpn.Route{}
	},
	EMR_BLOCKPUBLICACCESSCONFIGURATION: func() Any {
		return &emr.BlockPublicAccessConfiguration{}
	},
	NETWORKMANAGER_CONNECTATTACHMENT: func() Any {
		return &networkmanager.ConnectAttachment{}
	},
	REDSHIFT_ENDPOINTACCESS: func() Any {
		return &redshift.EndpointAccess{}
	},
	APIGATEWAY_BASEPATHMAPPING: func() Any {
		return &aws_apigateway.BasePathMapping{}
	},
	EC2_SNAPSHOTCREATEVOLUMEPERMISSION: func() Any {
		return &ec2.SnapshotCreateVolumePermission{}
	},
	LIGHTSAIL_CERTIFICATE: func() Any {
		return &lightsail.Certificate{}
	},
	REDSHIFT_EVENTSUBSCRIPTION: func() Any {
		return &redshift.EventSubscription{}
	},
	S3_BUCKETOBJECTLOCKCONFIGURATIONV2: func() Any {
		return &s3.BucketObjectLockConfigurationV2{}
	},
	SCHEMAS_SCHEMA: func() Any {
		return &schemas.Schema{}
	},
	SECURITYHUB_MEMBER: func() Any {
		return &securityhub.Member{}
	},
	CLOUDHSMV2_HSM: func() Any {
		return &cloudhsmv2.Hsm{}
	},
	MEDIASTORE_CONTAINER: func() Any {
		return &mediastore.Container{}
	},
	AUDITMANAGER_ASSESSMENTREPORT: func() Any {
		return &auditmanager.AssessmentReport{}
	},
	EC2TRANSITGATEWAY_CONNECT: func() Any {
		return &ec2transitgateway.Connect{}
	},
	SECURITYHUB_ORGANIZATIONCONFIGURATION: func() Any {
		return &securityhub.OrganizationConfiguration{}
	},
	CLOUDWATCH_QUERYDEFINITION: func() Any {
		return &cloudwatch.QueryDefinition{}
	},
	EC2_AMIFROMINSTANCE: func() Any {
		return &ec2.AmiFromInstance{}
	},
	SIGNER_SIGNINGPROFILE: func() Any {
		return &signer.SigningProfile{}
	},
	CONNECT_USER: func() Any {
		return &connect.User{}
	},
	CONNECT_USERHIERARCHYGROUP: func() Any {
		return &connect.UserHierarchyGroup{}
	},
	DAX_SUBNETGROUP: func() Any {
		return &dax.SubnetGroup{}
	},
	APPCONFIG_APPLICATION: func() Any {
		return &appconfig.Application{}
	},
	CUSTOMERPROFILES_DOMAIN: func() Any {
		return &customerprofiles.Domain{}
	},
	DMS_REPLICATIONSUBNETGROUP: func() Any {
		return &dms.ReplicationSubnetGroup{}
	},
	EMRCONTAINERS_VIRTUALCLUSTER: func() Any {
		return &emrcontainers.VirtualCluster{}
	},
	RESOURCEEXPLORER_INDEX: func() Any {
		return &resourceexplorer.Index{}
	},
	SECRETSMANAGER_SECRET: func() Any {
		return &secretsmanager.Secret{}
	},
	SECURITYHUB_ACTIONTARGET: func() Any {
		return &securityhub.ActionTarget{}
	},
	SERVICECATALOG_TAGOPTIONRESOURCEASSOCIATION: func() Any {
		return &servicecatalog.TagOptionResourceAssociation{}
	},
	SES_IDENTITYNOTIFICATIONTOPIC: func() Any {
		return &ses.IdentityNotificationTopic{}
	},
	APIGATEWAYV2_DEPLOYMENT: func() Any {
		return &apigatewayv2.Deployment{}
	},
	CLOUDWATCH_EVENTBUS: func() Any {
		return &cloudwatch.EventBus{}
	},
	CONNECT_VOCABULARY: func() Any {
		return &connect.Vocabulary{}
	},
	KINESISANALYTICSV2_APPLICATION: func() Any {
		return &kinesisanalyticsv2.Application{}
	},
	ACMPCA_CERTIFICATEAUTHORITYCERTIFICATE: func() Any {
		return &acmpca.CertificateAuthorityCertificate{}
	},
	GUARDDUTY_ORGANIZATIONADMINACCOUNT: func() Any {
		return &guardduty.OrganizationAdminAccount{}
	},
	VERIFIEDPERMISSIONS_POLICYTEMPLATE: func() Any {
		return &verifiedpermissions.PolicyTemplate{}
	},
	CODEPIPELINE_CUSTOMACTIONTYPE: func() Any {
		return &codepipeline.CustomActionType{}
	},
	ELASTICACHE_USERGROUP: func() Any {
		return &elasticache.UserGroup{}
	},
	PIPES_PIPE: func() Any {
		return &pipes.Pipe{}
	},
	SAGEMAKER_SERVICECATALOGPORTFOLIOSTATUS: func() Any {
		return &sagemaker.ServicecatalogPortfolioStatus{}
	},
	SERVICEDISCOVERY_HTTPNAMESPACE: func() Any {
		return &servicediscovery.HttpNamespace{}
	},
	CONNECT_INSTANCESTORAGECONFIG: func() Any {
		return &connect.InstanceStorageConfig{}
	},
	INSPECTOR2_DELEGATEDADMINACCOUNT: func() Any {
		return &inspector2.DelegatedAdminAccount{}
	},
	OPSWORKS_CUSTOMLAYER: func() Any {
		return &opsworks.CustomLayer{}
	},
	PINPOINT_APNSSANDBOXCHANNEL: func() Any {
		return &pinpoint.ApnsSandboxChannel{}
	},
	RDS_CLUSTERPARAMETERGROUP: func() Any {
		return &rds.ClusterParameterGroup{}
	},
	CODEBUILD_REPORTGROUP: func() Any {
		return &codebuild.ReportGroup{}
	},
	QUICKSIGHT_REFRESHSCHEDULE: func() Any {
		return &quicksight.RefreshSchedule{}
	},
	S3_BUCKETV2: func() Any {
		return &s3.BucketV2{}
	},
	WAF_SQLINJECTIONMATCHSET: func() Any {
		return &waf.SqlInjectionMatchSet{}
	},
	COGNITO_RESOURCESERVER: func() Any {
		return &cognito.ResourceServer{}
	},
	FSX_ONTAPFILESYSTEM: func() Any {
		return &fsx.OntapFileSystem{}
	},
	SHIELD_PROTECTIONHEALTHCHECKASSOCIATION: func() Any {
		return &shield.ProtectionHealthCheckAssociation{}
	},
	ALB_TARGETGROUP: func() Any {
		return &alb.TargetGroup{}
	},
	APPSYNC_DATASOURCE: func() Any {
		return &appsync.DataSource{}
	},
	CLOUDFRONT_FIELDLEVELENCRYPTIONPROFILE: func() Any {
		return &cloudfront.FieldLevelEncryptionProfile{}
	},
	CLOUDWATCH_LOGGROUP: func() Any {
		return &cloudwatch.LogGroup{}
	},
	EC2_NETWORKINSIGHTSANALYSIS: func() Any {
		return &ec2.NetworkInsightsAnalysis{}
	},
	EC2_SECURITYGROUP: func() Any {
		return &ec2.SecurityGroup{}
	},
	SES_EMAILIDENTITY: func() Any {
		return &ses.EmailIdentity{}
	},
	ATHENA_NAMEDQUERY: func() Any {
		return &athena.NamedQuery{}
	},
	EC2_DEFAULTVPC: func() Any {
		return &ec2.DefaultVpc{}
	},
	ECS_CLUSTER: func() Any {
		return &ecs.Cluster{}
	},
	FINSPACE_KXDATABASE: func() Any {
		return &finspace.KxDatabase{}
	},
	MACIE2_MEMBER: func() Any {
		return &macie2.Member{}
	},
	REDSHIFT_USAGELIMIT: func() Any {
		return &redshift.UsageLimit{}
	},
	CONNECT_ROUTINGPROFILE: func() Any {
		return &connect.RoutingProfile{}
	},
	CONNECT_USERHIERARCHYSTRUCTURE: func() Any {
		return &connect.UserHierarchyStructure{}
	},
	IVS_RECORDINGCONFIGURATION: func() Any {
		return &ivs.RecordingConfiguration{}
	},
	OPENSEARCH_PACKAGE: func() Any {
		return &opensearch.Package{}
	},
	SECURITYHUB_PRODUCTSUBSCRIPTION: func() Any {
		return &securityhub.ProductSubscription{}
	},
	STORAGEGATEWAY_NFSFILESHARE: func() Any {
		return &storagegateway.NfsFileShare{}
	},
	APPRUNNER_VPCINGRESSCONNECTION: func() Any {
		return &apprunner.VpcIngressConnection{}
	},
	CLOUDFRONT_KEYVALUESTORE: func() Any {
		return &cloudfront.KeyValueStore{}
	},
	REDSHIFT_PARTNER: func() Any {
		return &redshift.Partner{}
	},
	SCHEDULER_SCHEDULEGROUP: func() Any {
		return &scheduler.ScheduleGroup{}
	},
	SESV2_EMAILIDENTITYFEEDBACKATTRIBUTES: func() Any {
		return &sesv2.EmailIdentityFeedbackAttributes{}
	},
	ALB_LISTENERRULE: func() Any {
		return &alb.ListenerRule{}
	},
	LIGHTSAIL_LBCERTIFICATEATTACHMENT: func() Any {
		return &lightsail.LbCertificateAttachment{}
	},
	APIGATEWAYV2_DOMAINNAME: func() Any {
		return &apigatewayv2.DomainName{}
	},
	EBS_SNAPSHOT: func() Any {
		return &ebs.Snapshot{}
	},
	EC2_INTERNETGATEWAYATTACHMENT: func() Any {
		return &ec2.InternetGatewayAttachment{}
	},
	EC2_SERIALCONSOLEACCESS: func() Any {
		return &ec2.SerialConsoleAccess{}
	},
	IAM_GROUP: func() Any {
		return &aws_iam.Group{}
	},
	IAM_GROUPMEMBERSHIP: func() Any {
		return &aws_iam.GroupMembership{}
	},
	S3_BUCKETCORSCONFIGURATIONV2: func() Any {
		return &s3.BucketCorsConfigurationV2{}
	},
	SSMCONTACTS_CONTACTCHANNEL: func() Any {
		return &ssmcontacts.ContactChannel{}
	},
	APIGATEWAYV2_ROUTERESPONSE: func() Any {
		return &apigatewayv2.RouteResponse{}
	},
	EC2_SUBNETCIDRRESERVATION: func() Any {
		return &ec2.SubnetCidrReservation{}
	},
	EC2_VPCIPAMRESOURCEDISCOVERYASSOCIATION: func() Any {
		return &ec2.VpcIpamResourceDiscoveryAssociation{}
	},
	GLUE_CRAWLER: func() Any {
		return &glue.Crawler{}
	},
	DMS_CERTIFICATE: func() Any {
		return &dms.Certificate{}
	},
	IOT_TOPICRULE: func() Any {
		return &iot.TopicRule{}
	},
	SECURITYHUB_FINDINGAGGREGATOR: func() Any {
		return &securityhub.FindingAggregator{}
	},
	SESV2_DEDICATEDIPPOOL: func() Any {
		return &sesv2.DedicatedIpPool{}
	},
	CLOUDWATCH_EVENTBUSPOLICY: func() Any {
		return &cloudwatch.EventBusPolicy{}
	},
	EC2_CUSTOMERGATEWAY: func() Any {
		return &ec2.CustomerGateway{}
	},
	EC2_NETWORKINTERFACEATTACHMENT: func() Any {
		return &ec2.NetworkInterfaceAttachment{}
	},
	RAM_RESOURCESHAREACCEPTER: func() Any {
		return &ram.ResourceShareAccepter{}
	},
	ROUTE53RECOVERYREADINESS_READINESSCHECK: func() Any {
		return &route53recoveryreadiness.ReadinessCheck{}
	},
	VPCLATTICE_RESOURCEPOLICY: func() Any {
		return &vpclattice.ResourcePolicy{}
	},
	EC2TRANSITGATEWAY_PEERINGATTACHMENTACCEPTER: func() Any {
		return &ec2transitgateway.PeeringAttachmentAccepter{}
	},
	EC2TRANSITGATEWAY_TRANSITGATEWAY: func() Any {
		return &ec2transitgateway.TransitGateway{}
	},
	IOT_POLICYATTACHMENT: func() Any {
		return &iot.PolicyAttachment{}
	},
	SAGEMAKER_WORKFORCE: func() Any {
		return &sagemaker.Workforce{}
	},
	APPMESH_GATEWAYROUTE: func() Any {
		return &appmesh.GatewayRoute{}
	},
	DIRECTORYSERVICE_LOGSERVICE: func() Any {
		return &directoryservice.LogService{}
	},
	FMS_ADMINACCOUNT: func() Any {
		return &fms.AdminAccount{}
	},
	MEDIALIVE_MULTIPLEXPROGRAM: func() Any {
		return &medialive.MultiplexProgram{}
	},
	RESOURCEGROUPS_RESOURCE: func() Any {
		return &resourcegroups.Resource{}
	},
	APPSTREAM_FLEET: func() Any {
		return &appstream.Fleet{}
	},
	CFG_ORGANIZATIONCUSTOMRULE: func() Any {
		return &cfg.OrganizationCustomRule{}
	},
	REDSHIFT_CLUSTER: func() Any {
		return &redshift.Cluster{}
	},
	SSM_MAINTENANCEWINDOWTASK: func() Any {
		return &ssm.MaintenanceWindowTask{}
	},
	SSM_PATCHBASELINE: func() Any {
		return &ssm.PatchBaseline{}
	},
	SYNTHETICS_GROUPASSOCIATION: func() Any {
		return &synthetics.GroupAssociation{}
	},
	AMP_ALERTMANAGERDEFINITION: func() Any {
		return &amp.AlertManagerDefinition{}
	},
	AUTOSCALING_ATTACHMENT: func() Any {
		return &autoscaling.Attachment{}
	},
	CLEANROOMS_CONFIGUREDTABLE: func() Any {
		return &cleanrooms.ConfiguredTable{}
	},
	CLOUDWATCH_COMPOSITEALARM: func() Any {
		return &cloudwatch.CompositeAlarm{}
	},
	EC2_VPCPEERINGCONNECTIONACCEPTER: func() Any {
		return &ec2.VpcPeeringConnectionAccepter{}
	},
	LAKEFORMATION_LFTAG: func() Any {
		return &lakeformation.LfTag{}
	},
	OPSWORKS_HAPROXYLAYER: func() Any {
		return &opsworks.HaproxyLayer{}
	},
	ORGANIZATIONS_DELEGATEDADMINISTRATOR: func() Any {
		return &aws_organizations.DelegatedAdministrator{}
	},
	SAGEMAKER_FEATUREGROUP: func() Any {
		return &sagemaker.FeatureGroup{}
	},
	SAGEMAKER_NOTEBOOKINSTANCELIFECYCLECONFIGURATION: func() Any {
		return &sagemaker.NotebookInstanceLifecycleConfiguration{}
	},
	CONTROLTOWER_LANDINGZONE: func() Any {
		return &controltower.LandingZone{}
	},
	DIRECTCONNECT_HOSTEDPUBLICVIRTUALINTERFACEACCEPTER: func() Any {
		return &directconnect.HostedPublicVirtualInterfaceAccepter{}
	},
	EC2_ROUTE: func() Any {
		return &ec2.Route{}
	},
	PINPOINT_SMSCHANNEL: func() Any {
		return &pinpoint.SmsChannel{}
	},
	RDS_INSTANCEAUTOMATEDBACKUPSREPLICATION: func() Any {
		return &rds.InstanceAutomatedBackupsReplication{}
	},
	SAGEMAKER_DEVICE: func() Any {
		return &sagemaker.Device{}
	},
	AUDITMANAGER_ACCOUNTREGISTRATION: func() Any {
		return &auditmanager.AccountRegistration{}
	},
	CHIME_VOICECONNECTORLOGGING: func() Any {
		return &chime.VoiceConnectorLogging{}
	},
	DAX_CLUSTER: func() Any {
		return &dax.Cluster{}
	},
	FSX_BACKUP: func() Any {
		return &fsx.Backup{}
	},
	MEDIAPACKAGE_CHANNEL: func() Any {
		return &mediapackage.Channel{}
	},
	RESOURCEGROUPS_GROUP: func() Any {
		return &resourcegroups.Group{}
	},
	ROUTE53_ZONE: func() Any {
		return &route53.Zone{}
	},
	SERVICECATALOG_SERVICEACTION: func() Any {
		return &servicecatalog.ServiceAction{}
	},
	STORAGEGATEWAY_UPLOADBUFFER: func() Any {
		return &storagegateway.UploadBuffer{}
	},
	APPRUNNER_DEFAULTAUTOSCALINGCONFIGURATIONVERSION: func() Any {
		return &apprunner.DefaultAutoScalingConfigurationVersion{}
	},
	DIRECTCONNECT_HOSTEDPRIVATEVIRTUALINTERFACE: func() Any {
		return &directconnect.HostedPrivateVirtualInterface{}
	},
	EC2TRANSITGATEWAY_POLICYTABLEASSOCIATION: func() Any {
		return &ec2transitgateway.PolicyTableAssociation{}
	},
	GAMELIFT_GAMESERVERGROUP: func() Any {
		return &gamelift.GameServerGroup{}
	},
	GLOBALACCELERATOR_LISTENER: func() Any {
		return &globalaccelerator.Listener{}
	},
	IAM_USERPOLICYATTACHMENT: func() Any {
		return &aws_iam.UserPolicyAttachment{}
	},
	SAGEMAKER_ENDPOINT: func() Any {
		return &sagemaker.Endpoint{}
	},
	SESV2_CONFIGURATIONSET: func() Any {
		return &sesv2.ConfigurationSet{}
	},
	SSM_PATCHGROUP: func() Any {
		return &ssm.PatchGroup{}
	},
	CODESTARCONNECTIONS_CONNECTION: func() Any {
		return &codestarconnections.Connection{}
	},
	GLUE_JOB: func() Any {
		return &glue.Job{}
	},
	NETWORKMANAGER_GLOBALNETWORK: func() Any {
		return &networkmanager.GlobalNetwork{}
	},
	SCHEMAS_REGISTRY: func() Any {
		return &schemas.Registry{}
	},
	TRANSCRIBE_VOCABULARY: func() Any {
		return &transcribe.Vocabulary{}
	},
	WORKLINK_WEBSITECERTIFICATEAUTHORITYASSOCIATION: func() Any {
		return &worklink.WebsiteCertificateAuthorityAssociation{}
	},
	APPFLOW_CONNECTORPROFILE: func() Any {
		return &appflow.ConnectorProfile{}
	},
	ATHENA_WORKGROUP: func() Any {
		return &athena.Workgroup{}
	},
	OPENSEARCH_SERVERLESSCOLLECTION: func() Any {
		return &opensearch.ServerlessCollection{}
	},
	VERIFIEDACCESS_GROUP: func() Any {
		return &verifiedaccess.Group{}
	},
	WAFREGIONAL_REGEXPATTERNSET: func() Any {
		return &wafregional.RegexPatternSet{}
	},
	CLOUDWATCH_LOGDATAPROTECTIONPOLICY: func() Any {
		return &cloudwatch.LogDataProtectionPolicy{}
	},
	DMS_REPLICATIONINSTANCE: func() Any {
		return &dms.ReplicationInstance{}
	},
	KMS_EXTERNALKEY: func() Any {
		return &aws_kms.ExternalKey{}
	},
	XRAY_ENCRYPTIONCONFIG: func() Any {
		return &xray.EncryptionConfig{}
	},
	BACKUP_VAULTNOTIFICATIONS: func() Any {
		return &backup.VaultNotifications{}
	},
	EMR_MANAGEDSCALINGPOLICY: func() Any {
		return &emr.ManagedScalingPolicy{}
	},
	SQS_REDRIVEPOLICY: func() Any {
		return &sqs.RedrivePolicy{}
	},
	ALB_LISTENERCERTIFICATE: func() Any {
		return &alb.ListenerCertificate{}
	},
	ECS_TASKDEFINITION: func() Any {
		return &ecs.TaskDefinition{}
	},
	NEPTUNE_SUBNETGROUP: func() Any {
		return &neptune.SubnetGroup{}
	},
	EC2_VPCENDPOINTPOLICY: func() Any {
		return &ec2.VpcEndpointPolicy{}
	},
	KEYSPACES_TABLE: func() Any {
		return &keyspaces.Table{}
	},
	MACIE2_ORGANIZATIONADMINACCOUNT: func() Any {
		return &macie2.OrganizationAdminAccount{}
	},
	MSK_CONFIGURATION: func() Any {
		return &msk.Configuration{}
	},
	QUICKSIGHT_USER: func() Any {
		return &quicksight.User{}
	},
	SERVICEDISCOVERY_PRIVATEDNSNAMESPACE: func() Any {
		return &servicediscovery.PrivateDnsNamespace{}
	},
	SSM_DOCUMENT: func() Any {
		return &ssm.Document{}
	},
	ACMPCA_PERMISSION: func() Any {
		return &acmpca.Permission{}
	},
	DATASYNC_FSXOPENZFSFILESYSTEM: func() Any {
		return &datasync.FsxOpenZfsFileSystem{}
	},
	EC2_VPCIPV6CIDRBLOCKASSOCIATION: func() Any {
		return &ec2.VpcIpv6CidrBlockAssociation{}
	},
	FSX_ONTAPSTORAGEVIRTUALMACHINE: func() Any {
		return &fsx.OntapStorageVirtualMachine{}
	},
	SERVICECATALOG_CONSTRAINT: func() Any {
		return &servicecatalog.Constraint{}
	},
	CFG_CONFORMANCEPACK: func() Any {
		return &cfg.ConformancePack{}
	},
	DATAEXCHANGE_REVISION: func() Any {
		return &dataexchange.Revision{}
	},
	LICENSEMANAGER_LICENSEGRANT: func() Any {
		return &licensemanager.LicenseGrant{}
	},
	S3_BUCKETNOTIFICATION: func() Any {
		return &s3.BucketNotification{}
	},
	SES_DOMAINIDENTITYVERIFICATION: func() Any {
		return &ses.DomainIdentityVerification{}
	},
	XRAY_GROUP: func() Any {
		return &xray.Group{}
	},
	CHIME_SDKVOICEVOICEPROFILEDOMAIN: func() Any {
		return &chime.SdkvoiceVoiceProfileDomain{}
	},
	DIRECTCONNECT_HOSTEDCONNECTION: func() Any {
		return &directconnect.HostedConnection{}
	},
	EC2_LAUNCHCONFIGURATION: func() Any {
		return &ec2.LaunchConfiguration{}
	},
	EC2_NETWORKACLRULE: func() Any {
		return &ec2.NetworkAclRule{}
	},
	EC2_VPCDHCPOPTIONS: func() Any {
		return &ec2.VpcDhcpOptions{}
	},
	MEDIALIVE_CHANNEL: func() Any {
		return &medialive.Channel{}
	},
	ROUTE53_HOSTEDZONEDNSSEC: func() Any {
		return &route53.HostedZoneDnsSec{}
	},
	ROUTE53_RESOLVERCONFIG: func() Any {
		return &route53.ResolverConfig{}
	},
	SECURITYHUB_STANDARDSCONTROL: func() Any {
		return &securityhub.StandardsControl{}
	},
	GLACIER_VAULT: func() Any {
		return &glacier.Vault{}
	},
	CLOUDCONTROL_RESOURCE: func() Any {
		return &cloudcontrol.Resource{}
	},
	EC2_ROUTETABLEASSOCIATION: func() Any {
		return &ec2.RouteTableAssociation{}
	},
	FINSPACE_KXVOLUME: func() Any {
		return &finspace.KxVolume{}
	},
	WAFREGIONAL_BYTEMATCHSET: func() Any {
		return &wafregional.ByteMatchSet{}
	},
	CLOUDFRONT_KEYGROUP: func() Any {
		return &cloudfront.KeyGroup{}
	},
	DOCDB_CLUSTERPARAMETERGROUP: func() Any {
		return &docdb.ClusterParameterGroup{}
	},
	EBS_SNAPSHOTCOPY: func() Any {
		return &ebs.SnapshotCopy{}
	},
	EC2TRANSITGATEWAY_MULTICASTDOMAINASSOCIATION: func() Any {
		return &ec2transitgateway.MulticastDomainAssociation{}
	},
	ELASTICBEANSTALK_APPLICATION: func() Any {
		return &elasticbeanstalk.Application{}
	},
	PINPOINT_ADMCHANNEL: func() Any {
		return &pinpoint.AdmChannel{}
	},
	SSM_MAINTENANCEWINDOWTARGET: func() Any {
		return &ssm.MaintenanceWindowTarget{}
	},
	DEVICEFARM_UPLOAD: func() Any {
		return &devicefarm.Upload{}
	},
	WORKSPACES_WORKSPACE: func() Any {
		return &workspaces.Workspace{}
	},
	EC2_NETWORKINSIGHTSPATH: func() Any {
		return &ec2.NetworkInsightsPath{}
	},
	SES_RECEIPTFILTER: func() Any {
		return &ses.ReceiptFilter{}
	},
	SHIELD_APPLICATIONLAYERAUTOMATICRESPONSE: func() Any {
		return &shield.ApplicationLayerAutomaticResponse{}
	},
	APIGATEWAYV2_INTEGRATIONRESPONSE: func() Any {
		return &apigatewayv2.IntegrationResponse{}
	},
	CONNECT_PHONENUMBER: func() Any {
		return &connect.PhoneNumber{}
	},
	DOCDB_CLUSTER: func() Any {
		return &docdb.Cluster{}
	},
	EC2_INTERNETGATEWAY: func() Any {
		return &ec2.InternetGateway{}
	},
	EC2_SUBNET: func() Any {
		return &ec2.Subnet{}
	},
	ECR_REPOSITORY: func() Any {
		return &ecr.Repository{}
	},
	IAM_SAMLPROVIDER: func() Any {
		return &aws_iam.SamlProvider{}
	},
	IAM_SECURITYTOKENSERVICEPREFERENCES: func() Any {
		return &aws_iam.SecurityTokenServicePreferences{}
	},
	CHIME_SDKVOICEGLOBALSETTINGS: func() Any {
		return &chime.SdkvoiceGlobalSettings{}
	},
	CODECATALYST_SOURCEREPOSITORY: func() Any {
		return &codecatalyst.SourceRepository{}
	},
	EC2_DEDICATEDHOST: func() Any {
		return &ec2.DedicatedHost{}
	},
	REDSHIFT_HSMCLIENTCERTIFICATE: func() Any {
		return &redshift.HsmClientCertificate{}
	},
	WAFV2_RULEGROUP: func() Any {
		return &wafv2.RuleGroup{}
	},
	BATCH_JOBQUEUE: func() Any {
		return &batch.JobQueue{}
	},
	GLACIER_VAULTLOCK: func() Any {
		return &glacier.VaultLock{}
	},
	GUARDDUTY_FILTER: func() Any {
		return &guardduty.Filter{}
	},
	LIGHTSAIL_CONTAINERSERVICEDEPLOYMENTVERSION: func() Any {
		return &lightsail.ContainerServiceDeploymentVersion{}
	},
	REDSHIFT_SNAPSHOTCOPYGRANT: func() Any {
		return &redshift.SnapshotCopyGrant{}
	},
	WAFREGIONAL_GEOMATCHSET: func() Any {
		return &wafregional.GeoMatchSet{}
	},
	APPCONFIG_DEPLOYMENT: func() Any {
		return &appconfig.Deployment{}
	},
	APPMESH_VIRTUALGATEWAY: func() Any {
		return &appmesh.VirtualGateway{}
	},
	CLOUDWATCH_EVENTENDPOINT: func() Any {
		return &cloudwatch.EventEndpoint{}
	},
	CLOUDWATCH_EVENTPERMISSION: func() Any {
		return &cloudwatch.EventPermission{}
	},
	CODESTARCONNECTIONS_HOST: func() Any {
		return &codestarconnections.Host{}
	},
	DATASYNC_AGENT: func() Any {
		return &datasync.Agent{}
	},
	SECRETSMANAGER_SECRETROTATION: func() Any {
		return &secretsmanager.SecretRotation{}
	},
	ELASTICACHE_SUBNETGROUP: func() Any {
		return &elasticache.SubnetGroup{}
	},
	LB_TRUSTSTORE: func() Any {
		return &lb.TrustStore{}
	},
	SECURITYHUB_INSIGHT: func() Any {
		return &securityhub.Insight{}
	},
	SERVICECATALOG_ORGANIZATIONSACCESS: func() Any {
		return &servicecatalog.OrganizationsAccess{}
	},
	APIGATEWAY_APIKEY: func() Any {
		return &aws_apigateway.ApiKey{}
	},
	CLOUDWATCH_METRICSTREAM: func() Any {
		return &cloudwatch.MetricStream{}
	},
	DETECTIVE_INVITATIONACCEPTER: func() Any {
		return &detective.InvitationAccepter{}
	},
	EC2_NATGATEWAY: func() Any {
		return &ec2.NatGateway{}
	},
	NETWORKMANAGER_CORENETWORK: func() Any {
		return &networkmanager.CoreNetwork{}
	},
	SSMINCIDENTS_REPLICATIONSET: func() Any {
		return &ssmincidents.ReplicationSet{}
	},
	IDENTITYSTORE_GROUPMEMBERSHIP: func() Any {
		return &identitystore.GroupMembership{}
	},
	PINPOINT_EVENTSTREAM: func() Any {
		return &pinpoint.EventStream{}
	},
	SNS_SMSPREFERENCES: func() Any {
		return &sns.SmsPreferences{}
	},
	WAFREGIONAL_SIZECONSTRAINTSET: func() Any {
		return &wafregional.SizeConstraintSet{}
	},
	LAMBDA_ALIAS: func() Any {
		return &lambda.Alias{}
	},
	ORGANIZATIONS_RESOURCEPOLICY: func() Any {
		return &aws_organizations.ResourcePolicy{}
	},
	PINPOINT_GCMCHANNEL: func() Any {
		return &pinpoint.GcmChannel{}
	},
	RDS_CLUSTERROLEASSOCIATION: func() Any {
		return &rds.ClusterRoleAssociation{}
	},
	ROUTE53_CIDRLOCATION: func() Any {
		return &route53.CidrLocation{}
	},
	SES_ACTIVERECEIPTRULESET: func() Any {
		return &ses.ActiveReceiptRuleSet{}
	},
	BUDGETS_BUDGET: func() Any {
		return &budgets.Budget{}
	},
	CFG_AGGREGATEAUTHORIZATION: func() Any {
		return &cfg.AggregateAuthorization{}
	},
	COGNITO_RISKCONFIGURATION: func() Any {
		return &cognito.RiskConfiguration{}
	},
	EC2TRANSITGATEWAY_PREFIXLISTREFERENCE: func() Any {
		return &ec2transitgateway.PrefixListReference{}
	},
	GLUE_DATACATALOGENCRYPTIONSETTINGS: func() Any {
		return &glue.DataCatalogEncryptionSettings{}
	},
	IOT_ROLEALIAS: func() Any {
		return &iot.RoleAlias{}
	},
	SHIELD_DRTACCESSLOGBUCKETASSOCIATION: func() Any {
		return &shield.DrtAccessLogBucketAssociation{}
	},
	WAF_REGEXPATTERNSET: func() Any {
		return &waf.RegexPatternSet{}
	},
	WAF_WEBACL: func() Any {
		return &waf.WebAcl{}
	},
	CLOUD9_ENVIRONMENTMEMBERSHIP: func() Any {
		return &cloud9.EnvironmentMembership{}
	},
	DATASYNC_NFSLOCATION: func() Any {
		return &datasync.NfsLocation{}
	},
	SSOADMIN_PERMISSIONSET: func() Any {
		return &ssoadmin.PermissionSet{}
	},
	APIGATEWAY_RESOURCE: func() Any {
		return &aws_apigateway.Resource{}
	},
	BEDROCKMODEL_INVOCATIONLOGGINGCONFIGURATION: func() Any {
		return &bedrockmodel.InvocationLoggingConfiguration{}
	},
	ELB_LISTENERPOLICY: func() Any {
		return &elb.ListenerPolicy{}
	},
	LOCATION_GEOFENCECOLLECTION: func() Any {
		return &location.GeofenceCollection{}
	},
	WAFV2_WEBACLASSOCIATION: func() Any {
		return &wafv2.WebAclAssociation{}
	},
	ATHENA_PREPAREDSTATEMENT: func() Any {
		return &athena.PreparedStatement{}
	},
	CLOUDFORMATION_STACKSETINSTANCE: func() Any {
		return &cloudformation.StackSetInstance{}
	},
	IAM_SERVICELINKEDROLE: func() Any {
		return &aws_iam.ServiceLinkedRole{}
	},
	QUICKSIGHT_ANALYSIS: func() Any {
		return &quicksight.Analysis{}
	},
	REDSHIFT_HSMCONFIGURATION: func() Any {
		return &redshift.HsmConfiguration{}
	},
	ROUTE53DOMAINS_DELEGATIONSIGNERRECORD: func() Any {
		return &route53domains.DelegationSignerRecord{}
	},
	SSOADMIN_INSTANCEACCESSCONTROLATTRIBUTES: func() Any {
		return &ssoadmin.InstanceAccessControlAttributes{}
	},
	CODEGURUREVIEWER_REPOSITORYASSOCIATION: func() Any {
		return &codegurureviewer.RepositoryAssociation{}
	},
	DIRECTORYSERVICE_SERVICEREGION: func() Any {
		return &directoryservice.ServiceRegion{}
	},
	EC2_AMI: func() Any {
		return &ec2.Ami{}
	},
	ECS_CLUSTERCAPACITYPROVIDERS: func() Any {
		return &ecs.ClusterCapacityProviders{}
	},
	KENDRA_FAQ: func() Any {
		return &kendra.Faq{}
	},
	S3_ACCOUNTPUBLICACCESSBLOCK: func() Any {
		return &s3.AccountPublicAccessBlock{}
	},
	EC2_CARRIERGATEWAY: func() Any {
		return &ec2.CarrierGateway{}
	},
	KENDRA_INDEX: func() Any {
		return &kendra.Index{}
	},
	LEX_V2MODELSBOTVERSION: func() Any {
		return &lex.V2modelsBotVersion{}
	},
	MEDIALIVE_INPUT: func() Any {
		return &medialive.Input{}
	},
	EFS_REPLICATIONCONFIGURATION: func() Any {
		return &efs.ReplicationConfiguration{}
	},
	OPSWORKS_ECSCLUSTERLAYER: func() Any {
		return &opsworks.EcsClusterLayer{}
	},
	QUICKSIGHT_FOLDER: func() Any {
		return &quicksight.Folder{}
	},
	VPCLATTICE_TARGETGROUP: func() Any {
		return &vpclattice.TargetGroup{}
	},
	APIGATEWAY_CLIENTCERTIFICATE: func() Any {
		return &aws_apigateway.ClientCertificate{}
	},
	DOCDB_SUBNETGROUP: func() Any {
		return &docdb.SubnetGroup{}
	},
	ROUTE53_KEYSIGNINGKEY: func() Any {
		return &route53.KeySigningKey{}
	},
	WAF_REGEXMATCHSET: func() Any {
		return &waf.RegexMatchSet{}
	},
	COGNITO_USERPOOLCLIENT: func() Any {
		return &cognito.UserPoolClient{}
	},
	ELASTICTRANSCODER_PRESET: func() Any {
		return &elastictranscoder.Preset{}
	},
	IAM_ROLE: func() Any {
		return &aws_iam.Role{}
	},
	SCHEMAS_REGISTRYPOLICY: func() Any {
		return &schemas.RegistryPolicy{}
	},
	IMAGEBUILDER_IMAGERECIPE: func() Any {
		return &imagebuilder.ImageRecipe{}
	},
	LICENSEMANAGER_LICENSEGRANTACCEPTER: func() Any {
		return &licensemanager.LicenseGrantAccepter{}
	},
	RDS_CLUSTER: func() Any {
		return &rds.Cluster{}
	},
	RDS_EXPORTTASK: func() Any {
		return &rds.ExportTask{}
	},
	DEVICEFARM_NETWORKPROFILE: func() Any {
		return &devicefarm.NetworkProfile{}
	},
	DOCDB_CLUSTERINSTANCE: func() Any {
		return &docdb.ClusterInstance{}
	},
	GLUE_TRIGGER: func() Any {
		return &glue.Trigger{}
	},
	OPSWORKS_MEMCACHEDLAYER: func() Any {
		return &opsworks.MemcachedLayer{}
	},
	APIGATEWAYV2_ROUTE: func() Any {
		return &apigatewayv2.Route{}
	},
	CHIME_SDKVOICESIPRULE: func() Any {
		return &chime.SdkvoiceSipRule{}
	},
	EC2_VPCENDPOINTCONNECTIONACCEPTER: func() Any {
		return &ec2.VpcEndpointConnectionAccepter{}
	},
	ECS_SERVICE: func() Any {
		return &ecs.Service{}
	},
	KINESIS_STREAM: func() Any {
		return &kinesis.Stream{}
	},
	SHIELD_DRTACCESSROLEARNASSOCIATION: func() Any {
		return &shield.DrtAccessRoleArnAssociation{}
	},
	TRANSCRIBE_MEDICALVOCABULARY: func() Any {
		return &transcribe.MedicalVocabulary{}
	},
	CLOUDWATCH_LOGDESTINATIONPOLICY: func() Any {
		return &cloudwatch.LogDestinationPolicy{}
	},
	CODEBUILD_RESOURCEPOLICY: func() Any {
		return &codebuild.ResourcePolicy{}
	},
	LIGHTSAIL_LBHTTPSREDIRECTIONPOLICY: func() Any {
		return &lightsail.LbHttpsRedirectionPolicy{}
	},
	SAGEMAKER_APP: func() Any {
		return &sagemaker.App{}
	},
	STORAGEGATEWAY_TAPEPOOL: func() Any {
		return &storagegateway.TapePool{}
	},
	VPCLATTICE_AUTHPOLICY: func() Any {
		return &vpclattice.AuthPolicy{}
	},
	CONNECT_SECURITYPROFILE: func() Any {
		return &connect.SecurityProfile{}
	},
	GLOBALACCELERATOR_ACCELERATOR: func() Any {
		return &globalaccelerator.Accelerator{}
	},
	OPSWORKS_STATICWEBLAYER: func() Any {
		return &opsworks.StaticWebLayer{}
	},
	RDS_SNAPSHOT: func() Any {
		return &rds.Snapshot{}
	},
	SAGEMAKER_SPACE: func() Any {
		return &sagemaker.Space{}
	},
	STORAGEGATEWAY_CACHE: func() Any {
		return &storagegateway.Cache{}
	},
	DIRECTCONNECT_LINKAGGREGATIONGROUP: func() Any {
		return &directconnect.LinkAggregationGroup{}
	},
	EKS_PODIDENTITYASSOCIATION: func() Any {
		return &eks.PodIdentityAssociation{}
	},
	GAMELIFT_ALIAS: func() Any {
		return &gamelift.Alias{}
	},
	RAM_SHARINGWITHORGANIZATION: func() Any {
		return &ram.SharingWithOrganization{}
	},
	RDS_RESERVEDINSTANCE: func() Any {
		return &rds.ReservedInstance{}
	},
	EBS_VOLUME: func() Any {
		return &ebs.Volume{}
	},
	LEX_V2MODELSSLOTTYPE: func() Any {
		return &lex.V2modelsSlotType{}
	},
	LIGHTSAIL_DOMAINENTRY: func() Any {
		return &lightsail.DomainEntry{}
	},
	NETWORKFIREWALL_RESOURCEPOLICY: func() Any {
		return &networkfirewall.ResourcePolicy{}
	},
	ROUTE53_DELEGATIONSET: func() Any {
		return &route53.DelegationSet{}
	},
	S3_BUCKETACCELERATECONFIGURATIONV2: func() Any {
		return &s3.BucketAccelerateConfigurationV2{}
	},
	SFN_STATEMACHINE: func() Any {
		return &sfn.StateMachine{}
	},
	DMS_EVENTSUBSCRIPTION: func() Any {
		return &dms.EventSubscription{}
	},
	APPSYNC_RESOLVER: func() Any {
		return &appsync.Resolver{}
	},
	CLOUDFRONT_ORIGINREQUESTPOLICY: func() Any {
		return &cloudfront.OriginRequestPolicy{}
	},
	CODECOMMIT_APPROVALRULETEMPLATE: func() Any {
		return &codecommit.ApprovalRuleTemplate{}
	},
	DIRECTCONNECT_TRANSITVIRTUALINTERFACE: func() Any {
		return &directconnect.TransitVirtualInterface{}
	},
	LEX_BOT: func() Any {
		return &lex.Bot{}
	},
	WAF_XSSMATCHSET: func() Any {
		return &waf.XssMatchSet{}
	},
	WORKLINK_FLEET: func() Any {
		return &worklink.Fleet{}
	},
	CLOUDTRAIL_EVENTDATASTORE: func() Any {
		return &cloudtrail.EventDataStore{}
	},
	EC2_MANAGEDPREFIXLIST: func() Any {
		return &ec2.ManagedPrefixList{}
	},
	PINPOINT_APNSCHANNEL: func() Any {
		return &pinpoint.ApnsChannel{}
	},
	ROUTE53_RESOLVERFIREWALLDOMAINLIST: func() Any {
		return &route53.ResolverFirewallDomainList{}
	},
	WAFV2_REGEXPATTERNSET: func() Any {
		return &wafv2.RegexPatternSet{}
	},
	ALB_TARGETGROUPATTACHMENT: func() Any {
		return &alb.TargetGroupAttachment{}
	},
	APIGATEWAY_AUTHORIZER: func() Any {
		return &aws_apigateway.Authorizer{}
	},
	APPSYNC_DOMAINNAMEAPIASSOCIATION: func() Any {
		return &appsync.DomainNameApiAssociation{}
	},
	AUTOSCALING_POLICY: func() Any {
		return &autoscaling.Policy{}
	},
	EC2_SPOTDATAFEEDSUBSCRIPTION: func() Any {
		return &ec2.SpotDatafeedSubscription{}
	},
	GLOBALACCELERATOR_ENDPOINTGROUP: func() Any {
		return &globalaccelerator.EndpointGroup{}
	},
	MEMORYDB_PARAMETERGROUP: func() Any {
		return &memorydb.ParameterGroup{}
	},
	OPSWORKS_APPLICATION: func() Any {
		return &opsworks.Application{}
	},
	S3_BUCKETSERVERSIDEENCRYPTIONCONFIGURATIONV2: func() Any {
		return &s3.BucketServerSideEncryptionConfigurationV2{}
	},
	SSMCONTACTS_CONTACT: func() Any {
		return &ssmcontacts.Contact{}
	},
	CLOUDFRONT_FIELDLEVELENCRYPTIONCONFIG: func() Any {
		return &cloudfront.FieldLevelEncryptionConfig{}
	},
	CONNECT_CONTACTFLOWMODULE: func() Any {
		return &connect.ContactFlowModule{}
	},
	DOCDB_GLOBALCLUSTER: func() Any {
		return &docdb.GlobalCluster{}
	},
	ECR_PULLTHROUGHCACHERULE: func() Any {
		return &ecr.PullThroughCacheRule{}
	},
	LIGHTSAIL_LBCERTIFICATE: func() Any {
		return &lightsail.LbCertificate{}
	},
	SECURITYLAKE_DATALAKE: func() Any {
		return &securitylake.DataLake{}
	},
	SYNTHETICS_CANARY: func() Any {
		return &synthetics.Canary{}
	},
	TRANSFER_PROFILE: func() Any {
		return &transfer.Profile{}
	},
	WORKSPACES_IPGROUP: func() Any {
		return &workspaces.IpGroup{}
	},
	EC2CLIENTVPN_AUTHORIZATIONRULE: func() Any {
		return &ec2clientvpn.AuthorizationRule{}
	},
	ECS_TASKSET: func() Any {
		return &ecs.TaskSet{}
	},
	ELB_APPCOOKIESTICKINESSPOLICY: func() Any {
		return &elb.AppCookieStickinessPolicy{}
	},
	LIGHTSAIL_DISTRIBUTION: func() Any {
		return &lightsail.Distribution{}
	},
	LOCATION_MAP: func() Any {
		return &location.Map{}
	},
	NEPTUNE_CLUSTERENDPOINT: func() Any {
		return &neptune.ClusterEndpoint{}
	},
	ORGANIZATIONS_ORGANIZATIONALUNIT: func() Any {
		return &aws_organizations.OrganizationalUnit{}
	},
	AMPLIFY_APP: func() Any {
		return &amplify.App{}
	},
	EC2_VPCIPAMORGANIZATIONADMINACCOUNT: func() Any {
		return &ec2.VpcIpamOrganizationAdminAccount{}
	},
	KINESIS_VIDEOSTREAM: func() Any {
		return &kinesis.VideoStream{}
	},
	LB_TARGETGROUPATTACHMENT: func() Any {
		return &lb.TargetGroupAttachment{}
	},
	CODEARTIFACT_DOMAINPERMISSIONS: func() Any {
		return &codeartifact.DomainPermissions{}
	},
	ELB_LOADBALANCER: func() Any {
		return &elb.LoadBalancer{}
	},
	GUARDDUTY_MEMBER: func() Any {
		return &guardduty.Member{}
	},
	NETWORKFIREWALL_LOGGINGCONFIGURATION: func() Any {
		return &networkfirewall.LoggingConfiguration{}
	},
	ROUTE53RECOVERYCONTROL_ROUTINGCONTROL: func() Any {
		return &route53recoverycontrol.RoutingControl{}
	},
	SERVICECATALOG_PROVISIONEDPRODUCT: func() Any {
		return &servicecatalog.ProvisionedProduct{}
	},
	CLOUDSEARCH_DOMAINSERVICEACCESSPOLICY: func() Any {
		return &cloudsearch.DomainServiceAccessPolicy{}
	},
	CODEBUILD_PROJECT: func() Any {
		return &codebuild.Project{}
	},
	EMR_STUDIO: func() Any {
		return &emr.Studio{}
	},
	EMR_STUDIOSESSIONMAPPING: func() Any {
		return &emr.StudioSessionMapping{}
	},
	GUARDDUTY_ORGANIZATIONCONFIGURATIONFEATURE: func() Any {
		return &guardduty.OrganizationConfigurationFeature{}
	},
	S3CONTROL_MULTIREGIONACCESSPOINT: func() Any {
		return &s3control.MultiRegionAccessPoint{}
	},
	VPC_SECURITYGROUPINGRESSRULE: func() Any {
		return &vpc.SecurityGroupIngressRule{}
	},
	XRAY_SAMPLINGRULE: func() Any {
		return &xray.SamplingRule{}
	},
	APPRUNNER_AUTOSCALINGCONFIGURATIONVERSION: func() Any {
		return &apprunner.AutoScalingConfigurationVersion{}
	},
	COGNITO_MANAGEDUSERPOOLCLIENT: func() Any {
		return &cognito.ManagedUserPoolClient{}
	},
	EC2_VPCENDPOINTROUTETABLEASSOCIATION: func() Any {
		return &ec2.VpcEndpointRouteTableAssociation{}
	},
	SES_RECEIPTRULESET: func() Any {
		return &ses.ReceiptRuleSet{}
	},
	ELASTICBEANSTALK_ENVIRONMENT: func() Any {
		return &elasticbeanstalk.Environment{}
	},
	KMS_KEY: func() Any {
		return &aws_kms.Key{}
	},
	VERIFIEDACCESS_INSTANCELOGGINGCONFIGURATION: func() Any {
		return &verifiedaccess.InstanceLoggingConfiguration{}
	},
	WAFREGIONAL_XSSMATCHSET: func() Any {
		return &wafregional.XssMatchSet{}
	},
	CONNECT_HOURSOFOPERATION: func() Any {
		return &connect.HoursOfOperation{}
	},
	FSX_OPENZFSSNAPSHOT: func() Any {
		return &fsx.OpenZfsSnapshot{}
	},
	KINESIS_STREAMCONSUMER: func() Any {
		return &kinesis.StreamConsumer{}
	},
	REDSHIFT_AUTHENTICATIONPROFILE: func() Any {
		return &redshift.AuthenticationProfile{}
	},
	IMAGEBUILDER_WORKFLOW: func() Any {
		return &imagebuilder.Workflow{}
	},
	KENDRA_DATASOURCE: func() Any {
		return &kendra.DataSource{}
	},
	OAM_LINK: func() Any {
		return &oam.Link{}
	},
	DEVICEFARM_TESTGRIDPROJECT: func() Any {
		return &devicefarm.TestGridProject{}
	},
	DIRECTCONNECT_GATEWAY: func() Any {
		return &directconnect.Gateway{}
	},
	FSX_WINDOWSFILESYSTEM: func() Any {
		return &fsx.WindowsFileSystem{}
	},
	NEPTUNE_EVENTSUBSCRIPTION: func() Any {
		return &neptune.EventSubscription{}
	},
	APIGATEWAY_RESPONSE: func() Any {
		return &aws_apigateway.Response{}
	},
	CLOUDWATCH_LOGSUBSCRIPTIONFILTER: func() Any {
		return &cloudwatch.LogSubscriptionFilter{}
	},
	EVIDENTLY_FEATURE: func() Any {
		return &evidently.Feature{}
	},
	OPSWORKS_MYSQLLAYER: func() Any {
		return &opsworks.MysqlLayer{}
	},
	REDSHIFT_PARAMETERGROUP: func() Any {
		return &redshift.ParameterGroup{}
	},
	TRANSFER_CONNECTOR: func() Any {
		return &transfer.Connector{}
	},
	CODEARTIFACT_REPOSITORYPERMISSIONSPOLICY: func() Any {
		return &codeartifact.RepositoryPermissionsPolicy{}
	},
	EC2_VPCPEERINGCONNECTION: func() Any {
		return &ec2.VpcPeeringConnection{}
	},
	EVIDENTLY_LAUNCH: func() Any {
		return &evidently.Launch{}
	},
	S3_BUCKETOWNERSHIPCONTROLS: func() Any {
		return &s3.BucketOwnershipControls{}
	},
	SECURITYHUB_ORGANIZATIONADMINACCOUNT: func() Any {
		return &securityhub.OrganizationAdminAccount{}
	},
	APPSTREAM_USER: func() Any {
		return &appstream.User{}
	},
	EC2_VPNGATEWAYATTACHMENT: func() Any {
		return &ec2.VpnGatewayAttachment{}
	},
	EMR_SECURITYCONFIGURATION: func() Any {
		return &emr.SecurityConfiguration{}
	},
	LOCATION_TRACKER: func() Any {
		return &location.Tracker{}
	},
	NETWORKMANAGER_DEVICE: func() Any {
		return &networkmanager.Device{}
	},
	NETWORKMANAGER_SITE: func() Any {
		return &networkmanager.Site{}
	},
	TRANSFER_SERVER: func() Any {
		return &transfer.Server{}
	},
	COGNITO_IDENTITYPOOLPROVIDERPRINCIPALTAG: func() Any {
		return &cognito.IdentityPoolProviderPrincipalTag{}
	},
	COGNITO_USERPOOLUICUSTOMIZATION: func() Any {
		return &cognito.UserPoolUICustomization{}
	},
	MSKCONNECT_CONNECTOR: func() Any {
		return &mskconnect.Connector{}
	},
	REDSHIFT_SUBNETGROUP: func() Any {
		return &redshift.SubnetGroup{}
	},
	S3CONTROL_OBJECTLAMBDAACCESSPOINTPOLICY: func() Any {
		return &s3control.ObjectLambdaAccessPointPolicy{}
	},
	WAFREGIONAL_RULEGROUP: func() Any {
		return &wafregional.RuleGroup{}
	},
	CLOUDWATCH_LOGDESTINATION: func() Any {
		return &cloudwatch.LogDestination{}
	},
	LEX_INTENT: func() Any {
		return &lex.Intent{}
	},
	RDS_CLUSTERSNAPSHOT: func() Any {
		return &rds.ClusterSnapshot{}
	},
	REDSHIFT_DATASHARECONSUMERASSOCIATION: func() Any {
		return &redshift.DataShareConsumerAssociation{}
	},
	CLOUDWATCH_EVENTTARGET: func() Any {
		return &cloudwatch.EventTarget{}
	},
	LAMBDA_LAYERVERSIONPERMISSION: func() Any {
		return &lambda.LayerVersionPermission{}
	},
	SES_MAILFROM: func() Any {
		return &ses.MailFrom{}
	},
	SSOADMIN_CUSTOMERMANAGEDPOLICYATTACHMENT: func() Any {
		return &ssoadmin.CustomerManagedPolicyAttachment{}
	},
	CLOUDFRONT_CONTINUOUSDEPLOYMENTPOLICY: func() Any {
		return &cloudfront.ContinuousDeploymentPolicy{}
	},
	IAM_OPENIDCONNECTPROVIDER: func() Any {
		return &aws_iam.OpenIdConnectProvider{}
	},
	KENDRA_EXPERIENCE: func() Any {
		return &kendra.Experience{}
	},
	OPENSEARCH_PACKAGEASSOCIATION: func() Any {
		return &opensearch.PackageAssociation{}
	},
	ROUTE53RECOVERYCONTROL_SAFETYRULE: func() Any {
		return &route53recoverycontrol.SafetyRule{}
	},
	NETWORKMANAGER_CUSTOMERGATEWAYASSOCIATION: func() Any {
		return &networkmanager.CustomerGatewayAssociation{}
	},
	EC2_MANAGEDPREFIXLISTENTRY: func() Any {
		return &ec2.ManagedPrefixListEntry{}
	},
	OPENSEARCH_SERVERLESSSECURITYPOLICY: func() Any {
		return &opensearch.ServerlessSecurityPolicy{}
	},
	AUTOSCALING_NOTIFICATION: func() Any {
		return &autoscaling.Notification{}
	},
	EC2_ROUTETABLE: func() Any {
		return &ec2.RouteTable{}
	},
	ECR_REGISTRYPOLICY: func() Any {
		return &ecr.RegistryPolicy{}
	},
	IOT_THINGGROUP: func() Any {
		return &iot.ThingGroup{}
	},
	LIGHTSAIL_INSTANCEPUBLICPORTS: func() Any {
		return &lightsail.InstancePublicPorts{}
	},
	NETWORKFIREWALL_FIREWALLPOLICY: func() Any {
		return &networkfirewall.FirewallPolicy{}
	},
	ROUTE53_RESOLVERQUERYLOGCONFIGASSOCIATION: func() Any {
		return &route53.ResolverQueryLogConfigAssociation{}
	},
	SERVICEDISCOVERY_PUBLICDNSNAMESPACE: func() Any {
		return &servicediscovery.PublicDnsNamespace{}
	},
	CODECOMMIT_REPOSITORY: func() Any {
		return &codecommit.Repository{}
	},
	EC2_LAUNCHTEMPLATE: func() Any {
		return &ec2.LaunchTemplate{}
	},
	EC2_SPOTINSTANCEREQUEST: func() Any {
		return &ec2.SpotInstanceRequest{}
	},
	GLUE_CONNECTION: func() Any {
		return &glue.Connection{}
	},
	GLUE_MLTRANSFORM: func() Any {
		return &glue.MLTransform{}
	},
	INSPECTOR_ASSESSMENTTEMPLATE: func() Any {
		return &inspector.AssessmentTemplate{}
	},
	IOT_POLICY: func() Any {
		return &iot.Policy{}
	},
	ORGANIZATIONS_ACCOUNT: func() Any {
		return &aws_organizations.Account{}
	},
	DEVICEFARM_INSTANCEPROFILE: func() Any {
		return &devicefarm.InstanceProfile{}
	},
	KMS_GRANT: func() Any {
		return &aws_kms.Grant{}
	},
	NEPTUNE_GLOBALCLUSTER: func() Any {
		return &neptune.GlobalCluster{}
	},
	NETWORKMANAGER_CORENETWORKPOLICYATTACHMENT: func() Any {
		return &networkmanager.CoreNetworkPolicyAttachment{}
	},
	SIGNER_SIGNINGPROFILEPERMISSION: func() Any {
		return &signer.SigningProfilePermission{}
	},
	TRANSFER_SSHKEY: func() Any {
		return &transfer.SshKey{}
	},
	TRANSFER_USER: func() Any {
		return &transfer.User{}
	},
	DATASYNC_TASK: func() Any {
		return &datasync.Task{}
	},
	EC2_PEERINGCONNECTIONOPTIONS: func() Any {
		return &ec2.PeeringConnectionOptions{}
	},
	KMS_REPLICAEXTERNALKEY: func() Any {
		return &aws_kms.ReplicaExternalKey{}
	},
	MEDIACONVERT_QUEUE: func() Any {
		return &mediaconvert.Queue{}
	},
	OPSWORKS_JAVAAPPLAYER: func() Any {
		return &opsworks.JavaAppLayer{}
	},
	GCP_ORGANIZATIONS_POLICY: func() Any {
		return &organizations.Policy{}
	},
	WAF_RULEGROUP: func() Any {
		return &waf.RuleGroup{}
	},
	CONNECT_CONTACTFLOW: func() Any {
		return &connect.ContactFlow{}
	},
	DETECTIVE_MEMBER: func() Any {
		return &detective.Member{}
	},
	DIRECTORYSERVICE_SHAREDDIRECTORY: func() Any {
		return &directoryservice.SharedDirectory{}
	},
	DYNAMODB_TABLEITEM: func() Any {
		return &dynamodb.TableItem{}
	},
	EC2_TRAFFICMIRRORFILTER: func() Any {
		return &ec2.TrafficMirrorFilter{}
	},
	IAM_SERVERCERTIFICATE: func() Any {
		return &aws_iam.ServerCertificate{}
	},
	IAM_VIRTUALMFADEVICE: func() Any {
		return &aws_iam.VirtualMfaDevice{}
	},
	SECRETSMANAGER_SECRETPOLICY: func() Any {
		return &secretsmanager.SecretPolicy{}
	},
	APIGATEWAY_METHOD: func() Any {
		return &aws_apigateway.Method{}
	},
	APPCONFIG_EXTENSION: func() Any {
		return &appconfig.Extension{}
	},
	BACKUP_VAULTLOCKCONFIGURATION: func() Any {
		return &backup.VaultLockConfiguration{}
	},
	GRAFANA_ROLEASSOCIATION: func() Any {
		return &grafana.RoleAssociation{}
	},
	ROUTE53_QUERYLOG: func() Any {
		return &route53.QueryLog{}
	},
	S3_ANALYTICSCONFIGURATION: func() Any {
		return &s3.AnalyticsConfiguration{}
	},
	SERVICEDISCOVERY_INSTANCE: func() Any {
		return &servicediscovery.Instance{}
	},
	IAM_SSHKEY: func() Any {
		return &aws_iam.SshKey{}
	},
	MACIE_CUSTOMDATAIDENTIFIER: func() Any {
		return &macie.CustomDataIdentifier{}
	},
	MSK_SERVERLESSCLUSTER: func() Any {
		return &msk.ServerlessCluster{}
	},
	QUICKSIGHT_GROUPMEMBERSHIP: func() Any {
		return &quicksight.GroupMembership{}
	},
	REDSHIFTSERVERLESS_WORKGROUP: func() Any {
		return &redshiftserverless.Workgroup{}
	},
	S3_BUCKETOBJECTV2: func() Any {
		return &s3.BucketObjectv2{}
	},
	SERVICECATALOG_PORTFOLIOSHARE: func() Any {
		return &servicecatalog.PortfolioShare{}
	},
	VPCLATTICE_LISTENERRULE: func() Any {
		return &vpclattice.ListenerRule{}
	},
	APIGATEWAY_DEPLOYMENT: func() Any {
		return &aws_apigateway.Deployment{}
	},
	APPSTREAM_DIRECTORYCONFIG: func() Any {
		return &appstream.DirectoryConfig{}
	},
	EC2_NETWORKACL: func() Any {
		return &ec2.NetworkAcl{}
	},
	LIGHTSAIL_BUCKETACCESSKEY: func() Any {
		return &lightsail.BucketAccessKey{}
	},
	NEPTUNE_PARAMETERGROUP: func() Any {
		return &neptune.ParameterGroup{}
	},
	REDSHIFT_SNAPSHOTSCHEDULEASSOCIATION: func() Any {
		return &redshift.SnapshotScheduleAssociation{}
	},
	SQS_REDRIVEALLOWPOLICY: func() Any {
		return &sqs.RedriveAllowPolicy{}
	},
	APIGATEWAYV2_INTEGRATION: func() Any {
		return &apigatewayv2.Integration{}
	},
	CHIME_VOICECONNECTORTERMINATIONCREDENTIALS: func() Any {
		return &chime.VoiceConnectorTerminationCredentials{}
	},
	DIRECTORYSERVICE_DIRECTORY: func() Any {
		return &directoryservice.Directory{}
	},
	EC2_EGRESSONLYINTERNETGATEWAY: func() Any {
		return &ec2.EgressOnlyInternetGateway{}
	},
	IDENTITYSTORE_GROUP: func() Any {
		return &identitystore.Group{}
	},
	OPSWORKS_GANGLIALAYER: func() Any {
		return &opsworks.GangliaLayer{}
	},
	QLDB_STREAM: func() Any {
		return &qldb.Stream{}
	},
	VERIFIEDPERMISSIONS_POLICYSTORE: func() Any {
		return &verifiedpermissions.PolicyStore{}
	},
	CLOUDWATCH_LOGSTREAM: func() Any {
		return &cloudwatch.LogStream{}
	},
	CLOUDWATCH_METRICALARM: func() Any {
		return &cloudwatch.MetricAlarm{}
	},
	GLUE_SCHEMA: func() Any {
		return &glue.Schema{}
	},
	LICENSEMANAGER_ASSOCIATION: func() Any {
		return &licensemanager.Association{}
	},
	MEMORYDB_USER: func() Any {
		return &memorydb.User{}
	},
	PINPOINT_APNSVOIPCHANNEL: func() Any {
		return &pinpoint.ApnsVoipChannel{}
	},
	RDS_CUSTOMDBENGINEVERSION: func() Any {
		return &rds.CustomDbEngineVersion{}
	},
	STORAGEGATEWAY_CACHESISCSIVOLUME: func() Any {
		return &storagegateway.CachesIscsiVolume{}
	},
	TRANSFER_CERTIFICATE: func() Any {
		return &transfer.Certificate{}
	},
	BACKUP_REPORTPLAN: func() Any {
		return &backup.ReportPlan{}
	},
	DATASYNC_EFSLOCATION: func() Any {
		return &datasync.EfsLocation{}
	},
	EC2_TRAFFICMIRRORTARGET: func() Any {
		return &ec2.TrafficMirrorTarget{}
	},
	KEYSPACES_KEYSPACE: func() Any {
		return &keyspaces.Keyspace{}
	},
	PINPOINT_EMAILCHANNEL: func() Any {
		return &pinpoint.EmailChannel{}
	},
	SESV2_CONTACTLIST: func() Any {
		return &sesv2.ContactList{}
	},
	TRANSFER_TAG: func() Any {
		return &transfer.Tag{}
	},
	EC2_NETWORKINTERFACESECURITYGROUPATTACHMENT: func() Any {
		return &ec2.NetworkInterfaceSecurityGroupAttachment{}
	},
	EC2_VPCIPAMPOOL: func() Any {
		return &ec2.VpcIpamPool{}
	},
	MSK_CLUSTER: func() Any {
		return &msk.Cluster{}
	},
	APPMESH_MESH: func() Any {
		return &appmesh.Mesh{}
	},
	CLOUDFRONT_RESPONSEHEADERSPOLICY: func() Any {
		return &cloudfront.ResponseHeadersPolicy{}
	},
	DATASYNC_LOCATIONFSXWINDOWS: func() Any {
		return &datasync.LocationFsxWindows{}
	},
	ECRPUBLIC_REPOSITORY: func() Any {
		return &ecrpublic.Repository{}
	},
	ELASTICSEARCH_DOMAIN: func() Any {
		return &elasticsearch.Domain{}
	},
	GUARDDUTY_DETECTORFEATURE: func() Any {
		return &guardduty.DetectorFeature{}
	},
	LAMBDA_INVOCATION: func() Any {
		return &lambda.Invocation{}
	},
	LOCATION_ROUTECALCULATION: func() Any {
		return &location.RouteCalculation{}
	},
	NEPTUNE_CLUSTERPARAMETERGROUP: func() Any {
		return &neptune.ClusterParameterGroup{}
	},
	SERVICEQUOTAS_TEMPLATE: func() Any {
		return &servicequotas.Template{}
	},
	APPMESH_VIRTUALSERVICE: func() Any {
		return &appmesh.VirtualService{}
	},
	CODEDEPLOY_APPLICATION: func() Any {
		return &codedeploy.Application{}
	},
	COGNITO_USERPOOL: func() Any {
		return &cognito.UserPool{}
	},
	EC2TRANSITGATEWAY_PEERINGATTACHMENT: func() Any {
		return &ec2transitgateway.PeeringAttachment{}
	},
	ECRPUBLIC_REPOSITORYPOLICY: func() Any {
		return &ecrpublic.RepositoryPolicy{}
	},
	IAM_ACCOUNTALIAS: func() Any {
		return &aws_iam.AccountAlias{}
	},
	LB_LISTENER: func() Any {
		return &lb.Listener{}
	},
	OPENSEARCH_SERVERLESSACCESSPOLICY: func() Any {
		return &opensearch.ServerlessAccessPolicy{}
	},
	WORKSPACES_DIRECTORY: func() Any {
		return &workspaces.Directory{}
	},
	DIRECTCONNECT_PUBLICVIRTUALINTERFACE: func() Any {
		return &directconnect.PublicVirtualInterface{}
	},
	RAM_PRINCIPALASSOCIATION: func() Any {
		return &ram.PrincipalAssociation{}
	},
	SERVICECATALOG_PRINCIPALPORTFOLIOASSOCIATION: func() Any {
		return &servicecatalog.PrincipalPortfolioAssociation{}
	},
	SSOADMIN_ACCOUNTASSIGNMENT: func() Any {
		return &ssoadmin.AccountAssignment{}
	},
	CLOUDWATCH_EVENTRULE: func() Any {
		return &cloudwatch.EventRule{}
	},
	EKS_ACCESSPOLICYASSOCIATION: func() Any {
		return &eks.AccessPolicyAssociation{}
	},
	ELASTICSEARCH_DOMAINPOLICY: func() Any {
		return &elasticsearch.DomainPolicy{}
	},
	FINSPACE_KXUSER: func() Any {
		return &finspace.KxUser{}
	},
	OPSWORKS_RAILSAPPLAYER: func() Any {
		return &opsworks.RailsAppLayer{}
	},
	RDS_OPTIONGROUP: func() Any {
		return &rds.OptionGroup{}
	},
	ROUTE53_TRAFFICPOLICYINSTANCE: func() Any {
		return &route53.TrafficPolicyInstance{}
	},
	SNS_TOPIC: func() Any {
		return &sns.Topic{}
	},
	ELB_SSLNEGOTIATIONPOLICY: func() Any {
		return &elb.SslNegotiationPolicy{}
	},
	KENDRA_THESAURUS: func() Any {
		return &kendra.Thesaurus{}
	},
	KINESIS_RESOURCEPOLICY: func() Any {
		return &kinesis.ResourcePolicy{}
	},
	ROUTE53RECOVERYREADINESS_RESOURCESET: func() Any {
		return &route53recoveryreadiness.ResourceSet{}
	},
	SSM_CONTACTSROTATION: func() Any {
		return &ssm.ContactsRotation{}
	},
	GLOBALACCELERATOR_CUSTOMROUTINGLISTENER: func() Any {
		return &globalaccelerator.CustomRoutingListener{}
	},
	RBIN_RULE: func() Any {
		return &rbin.Rule{}
	},
	APIGATEWAYV2_AUTHORIZER: func() Any {
		return &apigatewayv2.Authorizer{}
	},
	SES_IDENTITYPOLICY: func() Any {
		return &ses.IdentityPolicy{}
	},
	SESV2_DEDICATEDIPASSIGNMENT: func() Any {
		return &sesv2.DedicatedIpAssignment{}
	},
	SSM_DEFAULTPATCHBASELINE: func() Any {
		return &ssm.DefaultPatchBaseline{}
	},
	IAM_ACCESSKEY: func() Any {
		return &aws_iam.AccessKey{}
	},
	IAM_INSTANCEPROFILE: func() Any {
		return &aws_iam.InstanceProfile{}
	},
	LICENSEMANAGER_LICENSECONFIGURATION: func() Any {
		return &licensemanager.LicenseConfiguration{}
	},
	RDS_SUBNETGROUP: func() Any {
		return &rds.SubnetGroup{}
	},
	REDSHIFTSERVERLESS_USAGELIMIT: func() Any {
		return &redshiftserverless.UsageLimit{}
	},
	DIRECTCONNECT_HOSTEDTRANSITVIRTUALINTERFACE: func() Any {
		return &directconnect.HostedTransitVirtualInterface{}
	},
	DYNAMODB_CONTRIBUTORINSIGHTS: func() Any {
		return &dynamodb.ContributorInsights{}
	},
	EKS_FARGATEPROFILE: func() Any {
		return &eks.FargateProfile{}
	},
	CODECOMMIT_APPROVALRULETEMPLATEASSOCIATION: func() Any {
		return &codecommit.ApprovalRuleTemplateAssociation{}
	},
	EC2CLIENTVPN_ENDPOINT: func() Any {
		return &ec2clientvpn.Endpoint{}
	},
	NETWORKMANAGER_TRANSITGATEWAYCONNECTPEERASSOCIATION: func() Any {
		return &networkmanager.TransitGatewayConnectPeerAssociation{}
	},
	DYNAMODB_TABLE: func() Any {
		return &dynamodb.Table{}
	},
	GLUE_DEVENDPOINT: func() Any {
		return &glue.DevEndpoint{}
	},
	GUARDDUTY_ORGANIZATIONCONFIGURATION: func() Any {
		return &guardduty.OrganizationConfiguration{}
	},
	ROUTE53_RESOLVERFIREWALLRULEGROUP: func() Any {
		return &route53.ResolverFirewallRuleGroup{}
	},
	S3CONTROL_ACCESSGRANTSLOCATION: func() Any {
		return &s3control.AccessGrantsLocation{}
	},
	SNS_PLATFORMAPPLICATION: func() Any {
		return &sns.PlatformApplication{}
	},
	CODECATALYST_PROJECT: func() Any {
		return &codecatalyst.Project{}
	},
	EFS_ACCESSPOINT: func() Any {
		return &efs.AccessPoint{}
	},
	EKS_NODEGROUP: func() Any {
		return &eks.NodeGroup{}
	},
	OPENSEARCH_DOMAINPOLICY: func() Any {
		return &opensearch.DomainPolicy{}
	},
	RDS_CLUSTERACTIVITYSTREAM: func() Any {
		return &rds.ClusterActivityStream{}
	},
	CFG_RETENTIONCONFIGURATION: func() Any {
		return &cfg.RetentionConfiguration{}
	},
	EMRSERVERLESS_APPLICATION: func() Any {
		return &emrserverless.Application{}
	},
	BACKUP_SELECTION: func() Any {
		return &backup.Selection{}
	},
	CFG_CONFIGURATIONAGGREGATOR: func() Any {
		return &cfg.ConfigurationAggregator{}
	},
	EC2_DEFAULTSECURITYGROUP: func() Any {
		return &ec2.DefaultSecurityGroup{}
	},
	KMS_REPLICAKEY: func() Any {
		return &aws_kms.ReplicaKey{}
	},
	OPSWORKS_INSTANCE: func() Any {
		return &opsworks.Instance{}
	},
	CLOUDFRONT_ORIGINACCESSCONTROL: func() Any {
		return &cloudfront.OriginAccessControl{}
	},
	FINSPACE_KXCLUSTER: func() Any {
		return &finspace.KxCluster{}
	},
	LB_TARGETGROUP: func() Any {
		return &lb.TargetGroup{}
	},
	OPENSEARCH_VPCENDPOINT: func() Any {
		return &opensearch.VpcEndpoint{}
	},
	STORAGEGATEWAY_WORKINGSTORAGE: func() Any {
		return &storagegateway.WorkingStorage{}
	},
	WORKSPACES_CONNECTIONALIAS: func() Any {
		return &workspaces.ConnectionAlias{}
	},
	APIGATEWAY_RESTAPIPOLICY: func() Any {
		return &aws_apigateway.RestApiPolicy{}
	},
	ROUTE53_RECORD: func() Any {
		return &route53.Record{}
	},
	SSOADMIN_PERMISSIONSBOUNDARYATTACHMENT: func() Any {
		return &ssoadmin.PermissionsBoundaryAttachment{}
	},
	AUTOSCALING_GROUP: func() Any {
		return &autoscaling.Group{}
	},
	COSTEXPLORER_ANOMALYMONITOR: func() Any {
		return &costexplorer.AnomalyMonitor{}
	},
	GUARDDUTY_INVITEACCEPTER: func() Any {
		return &guardduty.InviteAccepter{}
	},
	IMAGEBUILDER_COMPONENT: func() Any {
		return &imagebuilder.Component{}
	},
	MSK_CLUSTERPOLICY: func() Any {
		return &msk.ClusterPolicy{}
	},
	CODEARTIFACT_DOMAIN: func() Any {
		return &codeartifact.Domain{}
	},
	DOCDB_CLUSTERSNAPSHOT: func() Any {
		return &docdb.ClusterSnapshot{}
	},
	EC2_VPCIPV4CIDRBLOCKASSOCIATION: func() Any {
		return &ec2.VpcIpv4CidrBlockAssociation{}
	},
	NETWORKMANAGER_CONNECTPEER: func() Any {
		return &networkmanager.ConnectPeer{}
	},
	SSMCONTACTS_PLAN: func() Any {
		return &ssmcontacts.Plan{}
	},
	APIGATEWAY_DOMAINNAME: func() Any {
		return &aws_apigateway.DomainName{}
	},
	DIRECTCONNECT_CONNECTION: func() Any {
		return &directconnect.Connection{}
	},
	EC2_VPCIPAMPREVIEWNEXTCIDR: func() Any {
		return &ec2.VpcIpamPreviewNextCidr{}
	},
	IOT_CERTIFICATE: func() Any {
		return &iot.Certificate{}
	},
	LIGHTSAIL_DATABASE: func() Any {
		return &lightsail.Database{}
	},
	ROUTE53_TRAFFICPOLICY: func() Any {
		return &route53.TrafficPolicy{}
	},
	EC2_AVAILABILITYZONEGROUP: func() Any {
		return &ec2.AvailabilityZoneGroup{}
	},
	QUICKSIGHT_INGESTION: func() Any {
		return &quicksight.Ingestion{}
	},
	CODEBUILD_WEBHOOK: func() Any {
		return &codebuild.Webhook{}
	},
	LEX_BOTALIAS: func() Any {
		return &lex.BotAlias{}
	},
	ROUTE53RECOVERYCONTROL_CONTROLPANEL: func() Any {
		return &route53recoverycontrol.ControlPanel{}
	},
	VPCLATTICE_SERVICENETWORK: func() Any {
		return &vpclattice.ServiceNetwork{}
	},
	WAFREGIONAL_WEBACLASSOCIATION: func() Any {
		return &wafregional.WebAclAssociation{}
	},
	APPMESH_VIRTUALNODE: func() Any {
		return &appmesh.VirtualNode{}
	},
	EC2_VPCIPAM: func() Any {
		return &ec2.VpcIpam{}
	},
	S3_BUCKET: func() Any {
		return &s3.Bucket{}
	},
	EFS_FILESYSTEM: func() Any {
		return &efs.FileSystem{}
	},
	MEDIALIVE_INPUTSECURITYGROUP: func() Any {
		return &medialive.InputSecurityGroup{}
	},
	MEMORYDB_CLUSTER: func() Any {
		return &memorydb.Cluster{}
	},
	QUICKSIGHT_DATASET: func() Any {
		return &quicksight.DataSet{}
	},
	CLOUDWATCH_LOGRESOURCEPOLICY: func() Any {
		return &cloudwatch.LogResourcePolicy{}
	},
	EC2_LOCALGATEWAYROUTETABLEVPCASSOCIATION: func() Any {
		return &ec2.LocalGatewayRouteTableVpcAssociation{}
	},
	INSPECTOR2_MEMBERASSOCIATION: func() Any {
		return &inspector2.MemberAssociation{}
	},
	ROUTE53_RESOLVERFIREWALLRULEGROUPASSOCIATION: func() Any {
		return &route53.ResolverFirewallRuleGroupAssociation{}
	},
	S3CONTROL_BUCKETPOLICY: func() Any {
		return &s3control.BucketPolicy{}
	},
	APIGATEWAY_USAGEPLANKEY: func() Any {
		return &aws_apigateway.UsagePlanKey{}
	},
	BACKUP_PLAN: func() Any {
		return &backup.Plan{}
	},
	BATCH_COMPUTEENVIRONMENT: func() Any {
		return &batch.ComputeEnvironment{}
	},
	CFG_REMEDIATIONCONFIGURATION: func() Any {
		return &cfg.RemediationConfiguration{}
	},
	DIRECTCONNECT_HOSTEDTRANSITVIRTUALINTERFACEACCEPTOR: func() Any {
		return &directconnect.HostedTransitVirtualInterfaceAcceptor{}
	},
	EC2_EIP: func() Any {
		return &ec2.Eip{}
	},
	GLUE_CATALOGTABLE: func() Any {
		return &glue.CatalogTable{}
	},
	IOT_DOMAINCONFIGURATION: func() Any {
		return &iot.DomainConfiguration{}
	},
	SERVICECATALOG_PRODUCTPORTFOLIOASSOCIATION: func() Any {
		return &servicecatalog.ProductPortfolioAssociation{}
	},
	VERIFIEDACCESS_INSTANCE: func() Any {
		return &verifiedaccess.Instance{}
	},
	VPCLATTICE_SERVICENETWORKVPCASSOCIATION: func() Any {
		return &vpclattice.ServiceNetworkVpcAssociation{}
	},
	COGNITO_USERINGROUP: func() Any {
		return &cognito.UserInGroup{}
	},
	DYNAMODB_GLOBALTABLE: func() Any {
		return &dynamodb.GlobalTable{}
	},
	EC2_AMILAUNCHPERMISSION: func() Any {
		return &ec2.AmiLaunchPermission{}
	},
	EC2_PROXYPROTOCOLPOLICY: func() Any {
		return &ec2.ProxyProtocolPolicy{}
	},
	EC2TRANSITGATEWAY_MULTICASTDOMAIN: func() Any {
		return &ec2transitgateway.MulticastDomain{}
	},
	GLUE_PARTITIONINDEX: func() Any {
		return &glue.PartitionIndex{}
	},
	WAF_SIZECONSTRAINTSET: func() Any {
		return &waf.SizeConstraintSet{}
	},
	CLOUDFRONT_PUBLICKEY: func() Any {
		return &cloudfront.PublicKey{}
	},
	CLOUDWATCH_LOGMETRICFILTER: func() Any {
		return &cloudwatch.LogMetricFilter{}
	},
	COSTEXPLORER_COSTALLOCATIONTAG: func() Any {
		return &costexplorer.CostAllocationTag{}
	},
	EMR_CLUSTER: func() Any {
		return &emr.Cluster{}
	},
	GLUE_USERDEFINEDFUNCTION: func() Any {
		return &glue.UserDefinedFunction{}
	},
	IAM_GROUPPOLICYATTACHMENT: func() Any {
		return &aws_iam.GroupPolicyAttachment{}
	},
	OAM_SINKPOLICY: func() Any {
		return &oam.SinkPolicy{}
	},
	ROUTE53_RESOLVERQUERYLOGCONFIG: func() Any {
		return &route53.ResolverQueryLogConfig{}
	},
	VERIFIEDACCESS_TRUSTPROVIDER: func() Any {
		return &verifiedaccess.TrustProvider{}
	},
	APIGATEWAY_ACCOUNT: func() Any {
		return &aws_apigateway.Account{}
	},
	APPSYNC_APICACHE: func() Any {
		return &appsync.ApiCache{}
	},
	CLOUDFRONT_DISTRIBUTION: func() Any {
		return &cloudfront.Distribution{}
	},
	SCHEMAS_DISCOVERER: func() Any {
		return &schemas.Discoverer{}
	},
	AMP_SCRAPER: func() Any {
		return &amp.Scraper{}
	},
	APPRUNNER_CUSTOMDOMAINASSOCIATION: func() Any {
		return &apprunner.CustomDomainAssociation{}
	},
	DATAPIPELINE_PIPELINEDEFINITION: func() Any {
		return &datapipeline.PipelineDefinition{}
	},
	EC2_VPCNETWORKPERFORMANCEMETRICSUBSCRIPTION: func() Any {
		return &ec2.VpcNetworkPerformanceMetricSubscription{}
	},
	ELASTICACHE_PARAMETERGROUP: func() Any {
		return &elasticache.ParameterGroup{}
	},
	IOT_EVENTCONFIGURATIONS: func() Any {
		return &iot.EventConfigurations{}
	},
	LIGHTSAIL_LB: func() Any {
		return &lightsail.Lb{}
	},
	LIGHTSAIL_LBSTICKINESSPOLICY: func() Any {
		return &lightsail.LbStickinessPolicy{}
	},
	REDSHIFTDATA_STATEMENT: func() Any {
		return &redshiftdata.Statement{}
	},
	REDSHIFTSERVERLESS_RESOURCEPOLICY: func() Any {
		return &redshiftserverless.ResourcePolicy{}
	},
	ROLESANYWHERE_TRUSTANCHOR: func() Any {
		return &rolesanywhere.TrustAnchor{}
	},
	SAGEMAKER_IMAGE: func() Any {
		return &sagemaker.Image{}
	},
	TIMESTREAMWRITE_TABLE: func() Any {
		return &timestreamwrite.Table{}
	},
	APIGATEWAY_INTEGRATION: func() Any {
		return &aws_apigateway.Integration{}
	},
	DEVICEFARM_PROJECT: func() Any {
		return &devicefarm.Project{}
	},
	EC2_MAINROUTETABLEASSOCIATION: func() Any {
		return &ec2.MainRouteTableAssociation{}
	},
	IOT_BILLINGGROUP: func() Any {
		return &iot.BillingGroup{}
	},
	OPENSEARCH_DOMAIN: func() Any {
		return &opensearch.Domain{}
	},
	RAM_RESOURCEASSOCIATION: func() Any {
		return &ram.ResourceAssociation{}
	},
	SNS_DATAPROTECTIONPOLICY: func() Any {
		return &sns.DataProtectionPolicy{}
	},
	SSOADMIN_APPLICATION: func() Any {
		return &ssoadmin.Application{}
	},
	TRANSFER_AGREEMENT: func() Any {
		return &transfer.Agreement{}
	},
	TRANSFER_WORKFLOW: func() Any {
		return &transfer.Workflow{}
	},
	DATASYNC_LOCATIONFSXLUSTRE: func() Any {
		return &datasync.LocationFsxLustre{}
	},
	EC2_DEFAULTROUTETABLE: func() Any {
		return &ec2.DefaultRouteTable{}
	},
	EC2_VPCIPAMPOOLCIDR: func() Any {
		return &ec2.VpcIpamPoolCidr{}
	},
	INSPECTOR2_ENABLER: func() Any {
		return &inspector2.Enabler{}
	},
	QUICKSIGHT_VPCCONNECTION: func() Any {
		return &quicksight.VpcConnection{}
	},
	ACCESSANALYZER_ANALYZER: func() Any {
		return &accessanalyzer.Analyzer{}
	},
	CLOUDSEARCH_DOMAIN: func() Any {
		return &cloudsearch.Domain{}
	},
	EC2TRANSITGATEWAY_POLICYTABLE: func() Any {
		return &ec2transitgateway.PolicyTable{}
	},
	ECS_CAPACITYPROVIDER: func() Any {
		return &ecs.CapacityProvider{}
	},
	ELB_LOADBALANCERPOLICY: func() Any {
		return &elb.LoadBalancerPolicy{}
	},
	KENDRA_QUERYSUGGESTIONSBLOCKLIST: func() Any {
		return &kendra.QuerySuggestionsBlockList{}
	},
	TIMESTREAMWRITE_DATABASE: func() Any {
		return &timestreamwrite.Database{}
	},
	EC2_KEYPAIR: func() Any {
		return &ec2.KeyPair{}
	},
	FSX_OPENZFSVOLUME: func() Any {
		return &fsx.OpenZfsVolume{}
	},
	GLUE_CLASSIFIER: func() Any {
		return &glue.Classifier{}
	},
	LIGHTSAIL_DOMAIN: func() Any {
		return &lightsail.Domain{}
	},
	NETWORKMANAGER_ATTACHMENTACCEPTER: func() Any {
		return &networkmanager.AttachmentAccepter{}
	},
	APPCONFIG_ENVIRONMENT: func() Any {
		return &appconfig.Environment{}
	},
	DETECTIVE_ORGANIZATIONADMINACCOUNT: func() Any {
		return &detective.OrganizationAdminAccount{}
	},
	SECURITYHUB_INVITEACCEPTER: func() Any {
		return &securityhub.InviteAccepter{}
	},
	VERIFIEDACCESS_INSTANCETRUSTPROVIDERATTACHMENT: func() Any {
		return &verifiedaccess.InstanceTrustProviderAttachment{}
	},
	APPINTEGRATIONS_DATAINTEGRATION: func() Any {
		return &appintegrations.DataIntegration{}
	},
	EC2_EIPASSOCIATION: func() Any {
		return &ec2.EipAssociation{}
	},
	LOCATION_TRACKERASSOCIATION: func() Any {
		return &location.TrackerAssociation{}
	},
	OPSWORKS_STACK: func() Any {
		return &opsworks.Stack{}
	},
	RESOURCEEXPLORER_VIEW: func() Any {
		return &resourceexplorer.View{}
	},
	CLOUDFRONT_REALTIMELOGCONFIG: func() Any {
		return &cloudfront.RealtimeLogConfig{}
	},
	ELASTICACHE_USER: func() Any {
		return &elasticache.User{}
	},
	FSX_FILECACHE: func() Any {
		return &fsx.FileCache{}
	},
	GUARDDUTY_IPSET: func() Any {
		return &guardduty.IPSet{}
	},
	IOT_THINGTYPE: func() Any {
		return &iot.ThingType{}
	},
	RDS_PROXYENDPOINT: func() Any {
		return &rds.ProxyEndpoint{}
	},
	S3_BUCKETOBJECT: func() Any {
		return &s3.BucketObject{}
	},
	SAGEMAKER_DOMAIN: func() Any {
		return &sagemaker.Domain{}
	},
	WAFV2_WEBACLLOGGINGCONFIGURATION: func() Any {
		return &wafv2.WebAclLoggingConfiguration{}
	},
	BACKUP_REGIONSETTINGS: func() Any {
		return &backup.RegionSettings{}
	},
	DMS_ENDPOINT: func() Any {
		return &dms.Endpoint{}
	},
	QUICKSIGHT_TEMPLATE: func() Any {
		return &quicksight.Template{}
	},
	ACMPCA_POLICY: func() Any {
		return &acmpca.Policy{}
	},
	APIGATEWAYV2_API: func() Any {
		return &apigatewayv2.Api{}
	},
	CLOUDTRAIL_TRAIL: func() Any {
		return &cloudtrail.Trail{}
	},
	OPSWORKS_PERMISSION: func() Any {
		return &opsworks.Permission{}
	},
	S3_OBJECTCOPY: func() Any {
		return &s3.ObjectCopy{}
	},
	SSM_SERVICESETTING: func() Any {
		return &ssm.ServiceSetting{}
	},
	VPCLATTICE_ACCESSLOGSUBSCRIPTION: func() Any {
		return &vpclattice.AccessLogSubscription{}
	},
	FSX_DATAREPOSITORYASSOCIATION: func() Any {
		return &fsx.DataRepositoryAssociation{}
	},
	LEX_SLOTTYPE: func() Any {
		return &lex.SlotType{}
	},
	CHIMESDKMEDIAPIPELINES_MEDIAINSIGHTSPIPELINECONFIGURATION: func() Any {
		return &chimesdkmediapipelines.MediaInsightsPipelineConfiguration{}
	},
	CLOUDHSMV2_CLUSTER: func() Any {
		return &cloudhsmv2.Cluster{}
	},
	COGNITO_USER: func() Any {
		return &cognito.User{}
	},
	DYNAMODB_KINESISSTREAMINGDESTINATION: func() Any {
		return &dynamodb.KinesisStreamingDestination{}
	},
	GAMELIFT_MATCHMAKINGRULESET: func() Any {
		return &gamelift.MatchmakingRuleSet{}
	},
	IAM_ACCOUNTPASSWORDPOLICY: func() Any {
		return &aws_iam.AccountPasswordPolicy{}
	},
	S3_BUCKETLOGGINGV2: func() Any {
		return &s3.BucketLoggingV2{}
	},
	CFG_ORGANIZATIONMANAGEDRULE: func() Any {
		return &cfg.OrganizationManagedRule{}
	},
	ROLESANYWHERE_PROFILE: func() Any {
		return &rolesanywhere.Profile{}
	},
	APIGATEWAY_DOCUMENTATIONVERSION: func() Any {
		return &aws_apigateway.DocumentationVersion{}
	},
	EC2TRANSITGATEWAY_ROUTETABLE: func() Any {
		return &ec2transitgateway.RouteTable{}
	},
	ECR_REGISTRYSCANNINGCONFIGURATION: func() Any {
		return &ecr.RegistryScanningConfiguration{}
	},
	NEPTUNE_CLUSTERINSTANCE: func() Any {
		return &neptune.ClusterInstance{}
	},
	ORGANIZATIONS_ORGANIZATION: func() Any {
		return &aws_organizations.Organization{}
	},
	AUTOSCALING_SCHEDULE: func() Any {
		return &autoscaling.Schedule{}
	},
	DATAPIPELINE_PIPELINE: func() Any {
		return &datapipeline.Pipeline{}
	},
	DIRECTCONNECT_HOSTEDPRIVATEVIRTUALINTERFACEACCEPTER: func() Any {
		return &directconnect.HostedPrivateVirtualInterfaceAccepter{}
	},
	SIMPLEDB_DOMAIN: func() Any {
		return &simpledb.Domain{}
	},
	KMS_CIPHERTEXT: func() Any {
		return &aws_kms.Ciphertext{}
	},
	S3CONTROL_BUCKET: func() Any {
		return &s3control.Bucket{}
	},
	AUDITMANAGER_FRAMEWORK: func() Any {
		return &auditmanager.Framework{}
	},
	CODEARTIFACT_REPOSITORY: func() Any {
		return &codeartifact.Repository{}
	},
	GAMELIFT_MATCHMAKINGCONFIGURATION: func() Any {
		return &gamelift.MatchmakingConfiguration{}
	},
	IDENTITYSTORE_USER: func() Any {
		return &identitystore.User{}
	},
	NETWORKMANAGER_TRANSITGATEWAYROUTETABLEATTACHMENT: func() Any {
		return &networkmanager.TransitGatewayRouteTableAttachment{}
	},
	VPCLATTICE_LISTENER: func() Any {
		return &vpclattice.Listener{}
	},
	CONNECT_LAMBDAFUNCTIONASSOCIATION: func() Any {
		return &connect.LambdaFunctionAssociation{}
	},
	IOT_AUTHORIZER: func() Any {
		return &iot.Authorizer{}
	},
	RDS_ROLEASSOCIATION: func() Any {
		return &rds.RoleAssociation{}
	},
	ROUTE53_ZONEASSOCIATION: func() Any {
		return &route53.ZoneAssociation{}
	},
	ROUTE53RECOVERYREADINESS_RECOVERYGROUP: func() Any {
		return &route53recoveryreadiness.RecoveryGroup{}
	},
	SESV2_EMAILIDENTITYPOLICY: func() Any {
		return &sesv2.EmailIdentityPolicy{}
	},
	APPCONFIG_DEPLOYMENTSTRATEGY: func() Any {
		return &appconfig.DeploymentStrategy{}
	},
	APPRUNNER_VPCCONNECTOR: func() Any {
		return &apprunner.VpcConnector{}
	},
	EC2_VPNCONNECTION: func() Any {
		return &ec2.VpnConnection{}
	},
	LEX_V2MODELSSLOT: func() Any {
		return &lex.V2modelsSlot{}
	},
	REKOGNITION_PROJECT: func() Any {
		return &rekognition.Project{}
	},
	APIGATEWAY_RESTAPI: func() Any {
		return &aws_apigateway.RestApi{}
	},
	CODECOMMIT_TRIGGER: func() Any {
		return &codecommit.Trigger{}
	},
	EVIDENTLY_PROJECT: func() Any {
		return &evidently.Project{}
	},
	GLUE_PARTITION: func() Any {
		return &glue.Partition{}
	},
	ROUTE53_RESOLVERRULE: func() Any {
		return &route53.ResolverRule{}
	},
	CODEDEPLOY_DEPLOYMENTGROUP: func() Any {
		return &codedeploy.DeploymentGroup{}
	},
	DMS_REPLICATIONCONFIG: func() Any {
		return &dms.ReplicationConfig{}
	},
	INSPECTOR_ASSESSMENTTARGET: func() Any {
		return &inspector.AssessmentTarget{}
	},
	KMS_CUSTOMKEYSTORE: func() Any {
		return &aws_kms.CustomKeyStore{}
	},
	CLOUDFRONT_ORIGINACCESSIDENTITY: func() Any {
		return &cloudfront.OriginAccessIdentity{}
	},
	FSX_OPENZFSFILESYSTEM: func() Any {
		return &fsx.OpenZfsFileSystem{}
	},
	S3CONTROL_ACCESSGRANTSINSTANCERESOURCEPOLICY: func() Any {
		return &s3control.AccessGrantsInstanceResourcePolicy{}
	},
	STORAGEGATEWAY_FILESYSTEMASSOCIATION: func() Any {
		return &storagegateway.FileSystemAssociation{}
	},
	CUR_REPORTDEFINITION: func() Any {
		return &cur.ReportDefinition{}
	},
	DMS_REPLICATIONTASK: func() Any {
		return &dms.ReplicationTask{}
	},
	EBS_ENCRYPTIONBYDEFAULT: func() Any {
		return &ebs.EncryptionByDefault{}
	},
	EC2_NETWORKACLASSOCIATION: func() Any {
		return &ec2.NetworkAclAssociation{}
	},
	LAMBDA_PROVISIONEDCONCURRENCYCONFIG: func() Any {
		return &lambda.ProvisionedConcurrencyConfig{}
	},
	SECURITYLAKE_AWSLOGSOURCE: func() Any {
		return &securitylake.AwsLogSource{}
	},
	SES_RECEIPTRULE: func() Any {
		return &ses.ReceiptRule{}
	},
	WAF_RATEBASEDRULE: func() Any {
		return &waf.RateBasedRule{}
	},
	CFG_RULE: func() Any {
		return &cfg.Rule{}
	},
	CLOUD9_ENVIRONMENTEC2: func() Any {
		return &cloud9.EnvironmentEC2{}
	},
	DATASYNC_LOCATIONOBJECTSTORAGE: func() Any {
		return &datasync.LocationObjectStorage{}
	},
	IOT_THINGPRINCIPALATTACHMENT: func() Any {
		return &iot.ThingPrincipalAttachment{}
	},
	LOCATION_PLACEINDEX: func() Any {
		return &location.PlaceIndex{}
	},
	NEPTUNE_CLUSTER: func() Any {
		return &neptune.Cluster{}
	},
	APIGATEWAYV2_STAGE: func() Any {
		return &apigatewayv2.Stage{}
	},
	EC2_TRAFFICMIRRORFILTERRULE: func() Any {
		return &ec2.TrafficMirrorFilterRule{}
	},
	EC2TRANSITGATEWAY_ROUTETABLEPROPAGATION: func() Any {
		return &ec2transitgateway.RouteTablePropagation{}
	},
	GLUE_REGISTRY: func() Any {
		return &glue.Registry{}
	},
	DIRECTCONNECT_HOSTEDPUBLICVIRTUALINTERFACE: func() Any {
		return &directconnect.HostedPublicVirtualInterface{}
	},
	ELASTICACHE_GLOBALREPLICATIONGROUP: func() Any {
		return &elasticache.GlobalReplicationGroup{}
	},
	LEX_V2MODELSINTENT: func() Any {
		return &lex.V2modelsIntent{}
	},
	S3_BUCKETMETRIC: func() Any {
		return &s3.BucketMetric{}
	},
	SAGEMAKER_MODELPACKAGEGROUP: func() Any {
		return &sagemaker.ModelPackageGroup{}
	},
	INSPECTOR_RESOURCEGROUP: func() Any {
		return &inspector.ResourceGroup{}
	},
	S3CONTROL_OBJECTLAMBDAACCESSPOINT: func() Any {
		return &s3control.ObjectLambdaAccessPoint{}
	},
	SFN_ALIAS: func() Any {
		return &sfn.Alias{}
	},
	APIGATEWAY_VPCLINK: func() Any {
		return &aws_apigateway.VpcLink{}
	},
	APPAUTOSCALING_SCHEDULEDACTION: func() Any {
		return &appautoscaling.ScheduledAction{}
	},
	COMPREHEND_ENTITYRECOGNIZER: func() Any {
		return &comprehend.EntityRecognizer{}
	},
	EC2_NETWORKINTERFACE: func() Any {
		return &ec2.NetworkInterface{}
	},
	ELASTICTRANSCODER_PIPELINE: func() Any {
		return &elastictranscoder.Pipeline{}
	},
	ROUTE53_CIDRCOLLECTION: func() Any {
		return &route53.CidrCollection{}
	},
	ROUTE53_RESOLVERFIREWALLRULE: func() Any {
		return &route53.ResolverFirewallRule{}
	},
	S3CONTROL_STORAGELENSCONFIGURATION: func() Any {
		return &s3control.StorageLensConfiguration{}
	},
	SAGEMAKER_CODEREPOSITORY: func() Any {
		return &sagemaker.CodeRepository{}
	},
	SSOADMIN_TRUSTEDTOKENISSUER: func() Any {
		return &ssoadmin.TrustedTokenIssuer{}
	},
	OPSWORKS_RDSDBINSTANCE: func() Any {
		return &opsworks.RdsDbInstance{}
	},
	S3_ACCESSPOINT: func() Any {
		return &s3.AccessPoint{}
	},
	SERVICECATALOG_PRODUCT: func() Any {
		return &servicecatalog.Product{}
	},
	SHIELD_PROACTIVEENGAGEMENT: func() Any {
		return &shield.ProactiveEngagement{}
	},
	CODEPIPELINE_PIPELINE: func() Any {
		return &codepipeline.Pipeline{}
	},
	INSPECTOR2_ORGANIZATIONCONFIGURATION: func() Any {
		return &inspector2.OrganizationConfiguration{}
	},
	IVS_PLAYBACKKEYPAIR: func() Any {
		return &ivs.PlaybackKeyPair{}
	},
	LAMBDA_FUNCTION: func() Any {
		return &lambda.Function{}
	},
	NETWORKMANAGER_TRANSITGATEWAYREGISTRATION: func() Any {
		return &networkmanager.TransitGatewayRegistration{}
	},
	SES_TEMPLATE: func() Any {
		return &ses.Template{}
	},
	DIRECTCONNECT_GATEWAYASSOCIATIONPROPOSAL: func() Any {
		return &directconnect.GatewayAssociationProposal{}
	},
	EC2_SPOTFLEETREQUEST: func() Any {
		return &ec2.SpotFleetRequest{}
	},
	ELASTICACHE_REPLICATIONGROUP: func() Any {
		return &elasticache.ReplicationGroup{}
	},
	IAM_POLICY: func() Any {
		return &aws_iam.Policy{}
	},
	KINESIS_FIREHOSEDELIVERYSTREAM: func() Any {
		return &kinesis.FirehoseDeliveryStream{}
	},
	ALB_LOADBALANCER: func() Any {
		return &alb.LoadBalancer{}
	},
	EC2TRANSITGATEWAY_MULTICASTGROUPSOURCE: func() Any {
		return &ec2transitgateway.MulticastGroupSource{}
	},
	ELASTICACHE_USERGROUPASSOCIATION: func() Any {
		return &elasticache.UserGroupAssociation{}
	},
	SAGEMAKER_DATAQUALITYJOBDEFINITION: func() Any {
		return &sagemaker.DataQualityJobDefinition{}
	},
	SAGEMAKER_PIPELINE: func() Any {
		return &sagemaker.Pipeline{}
	},
	SAGEMAKER_WORKTEAM: func() Any {
		return &sagemaker.Workteam{}
	},
	SECURITYLAKE_CUSTOMLOGSOURCE: func() Any {
		return &securitylake.CustomLogSource{}
	},
	SES_DOMAINDKIM: func() Any {
		return &ses.DomainDkim{}
	},
	VERIFIEDPERMISSIONS_SCHEMA: func() Any {
		return &verifiedpermissions.Schema{}
	},
	AUTOSCALING_LIFECYCLEHOOK: func() Any {
		return &autoscaling.LifecycleHook{}
	},
	CODESTARNOTIFICATIONS_NOTIFICATIONRULE: func() Any {
		return &codestarnotifications.NotificationRule{}
	},
	ROUTE53_VPCASSOCIATIONAUTHORIZATION: func() Any {
		return &route53.VpcAssociationAuthorization{}
	},
	SAGEMAKER_STUDIOLIFECYCLECONFIG: func() Any {
		return &sagemaker.StudioLifecycleConfig{}
	},
	SSOADMIN_APPLICATIONASSIGNMENT: func() Any {
		return &ssoadmin.ApplicationAssignment{}
	},
	CFG_ORGANIZATIONCUSTOMPOLICYRULE: func() Any {
		return &cfg.OrganizationCustomPolicyRule{}
	},
	CLOUDWATCH_EVENTCONNECTION: func() Any {
		return &cloudwatch.EventConnection{}
	},
	IMAGEBUILDER_DISTRIBUTIONCONFIGURATION: func() Any {
		return &imagebuilder.DistributionConfiguration{}
	},
	IOT_TOPICRULEDESTINATION: func() Any {
		return &iot.TopicRuleDestination{}
	},
	MWAA_ENVIRONMENT: func() Any {
		return &mwaa.Environment{}
	},
	QUICKSIGHT_IAMPOLICYASSIGNMENT: func() Any {
		return &quicksight.IamPolicyAssignment{}
	},
	ROUTE53_RESOLVERRULEASSOCIATION: func() Any {
		return &route53.ResolverRuleAssociation{}
	},
	ACCOUNT_PRIMARYCONTACT: func() Any {
		return &account.PrimaryContact{}
	},
	APIGATEWAY_METHODRESPONSE: func() Any {
		return &aws_apigateway.MethodResponse{}
	},
	APIGATEWAY_STAGE: func() Any {
		return &aws_apigateway.Stage{}
	},
	FSX_ONTAPVOLUME: func() Any {
		return &fsx.OntapVolume{}
	},
	GAMELIFT_SCRIPT: func() Any {
		return &gamelift.Script{}
	},
	LB_LISTENERRULE: func() Any {
		return &lb.ListenerRule{}
	},
	RDS_CLUSTERINSTANCE: func() Any {
		return &rds.ClusterInstance{}
	},
	REDSHIFT_ENDPOINTAUTHORIZATION: func() Any {
		return &redshift.EndpointAuthorization{}
	},
	RUM_METRICSDESTINATION: func() Any {
		return &rum.MetricsDestination{}
	},
	SCHEDULER_SCHEDULE: func() Any {
		return &scheduler.Schedule{}
	},
	APPCONFIG_CONFIGURATIONPROFILE: func() Any {
		return &appconfig.ConfigurationProfile{}
	},
	EC2_VPCENDPOINTSERVICEALLOWEDPRINCIPLE: func() Any {
		return &ec2.VpcEndpointServiceAllowedPrinciple{}
	},
	LIGHTSAIL_DISK: func() Any {
		return &lightsail.Disk{}
	},
	AUDITMANAGER_ASSESSMENTDELEGATION: func() Any {
		return &auditmanager.AssessmentDelegation{}
	},
	DYNAMODB_TABLEREPLICA: func() Any {
		return &dynamodb.TableReplica{}
	},
	EC2_VPCENDPOINTCONNECTIONNOTIFICATION: func() Any {
		return &ec2.VpcEndpointConnectionNotification{}
	},
	EC2_VPCENDPOINTSUBNETASSOCIATION: func() Any {
		return &ec2.VpcEndpointSubnetAssociation{}
	},
	MQ_BROKER: func() Any {
		return &mq.Broker{}
	},
	ALB_LISTENER: func() Any {
		return &alb.Listener{}
	},
	EC2_AMICOPY: func() Any {
		return &ec2.AmiCopy{}
	},
	EC2_CAPACITYRESERVATION: func() Any {
		return &ec2.CapacityReservation{}
	},
	EC2_VPNGATEWAY: func() Any {
		return &ec2.VpnGateway{}
	},
	IMAGEBUILDER_CONTAINERRECIPE: func() Any {
		return &imagebuilder.ContainerRecipe{}
	},
	LAMBDA_FUNCTIONURL: func() Any {
		return &lambda.FunctionUrl{}
	},
	QUICKSIGHT_THEME: func() Any {
		return &quicksight.Theme{}
	},
	SAGEMAKER_NOTEBOOKINSTANCE: func() Any {
		return &sagemaker.NotebookInstance{}
	},
	SERVICECATALOG_BUDGETRESOURCEASSOCIATION: func() Any {
		return &servicecatalog.BudgetResourceAssociation{}
	},
	SSM_PARAMETER: func() Any {
		return &ssm.Parameter{}
	},
	APIGATEWAYV2_VPCLINK: func() Any {
		return &apigatewayv2.VpcLink{}
	},
	ECR_REPLICATIONCONFIGURATION: func() Any {
		return &ecr.ReplicationConfiguration{}
	},
	ELASTICBEANSTALK_APPLICATIONVERSION: func() Any {
		return &elasticbeanstalk.ApplicationVersion{}
	},
	DYNAMODB_TAG: func() Any {
		return &dynamodb.Tag{}
	},
	EC2_VPCIPAMSCOPE: func() Any {
		return &ec2.VpcIpamScope{}
	},
	EC2TRANSITGATEWAY_CONNECTPEER: func() Any {
		return &ec2transitgateway.ConnectPeer{}
	},
	EC2TRANSITGATEWAY_INSTANCESTATE: func() Any {
		return &ec2transitgateway.InstanceState{}
	},
	EC2TRANSITGATEWAY_VPCATTACHMENT: func() Any {
		return &ec2transitgateway.VpcAttachment{}
	},
	LIGHTSAIL_BUCKETRESOURCEACCESS: func() Any {
		return &lightsail.BucketResourceAccess{}
	},
	OPENSEARCH_SERVERLESSSECURITYCONFIG: func() Any {
		return &opensearch.ServerlessSecurityConfig{}
	},
	QUICKSIGHT_DASHBOARD: func() Any {
		return &quicksight.Dashboard{}
	},
	GAMELIFT_GAMESESSIONQUEUE: func() Any {
		return &gamelift.GameSessionQueue{}
	},
	IAM_ROLEPOLICY: func() Any {
		return &aws_iam.RolePolicy{}
	},
	SAGEMAKER_HUMANTASKUI: func() Any {
		return &sagemaker.HumanTaskUI{}
	},
	WAF_RULE: func() Any {
		return &waf.Rule{}
	},
	CHIME_VOICECONNECTORORGANIZATION: func() Any {
		return &chime.VoiceConnectorOrganization{}
	},
	DIRECTORYSERVICE_RADIUSSETTINGS: func() Any {
		return &directoryservice.RadiusSettings{}
	},
	RDS_PROXYDEFAULTTARGETGROUP: func() Any {
		return &rds.ProxyDefaultTargetGroup{}
	},
	EMR_INSTANCEGROUP: func() Any {
		return &emr.InstanceGroup{}
	},
	NETWORKFIREWALL_RULEGROUP: func() Any {
		return &networkfirewall.RuleGroup{}
	},
	NETWORKMANAGER_TRANSITGATEWAYPEERING: func() Any {
		return &networkmanager.TransitGatewayPeering{}
	},
	SECURITYHUB_ACCOUNT: func() Any {
		return &securityhub.Account{}
	},
	SQS_QUEUEPOLICY: func() Any {
		return &sqs.QueuePolicy{}
	},
	SWF_DOMAIN: func() Any {
		return &swf.Domain{}
	},
	CHIME_VOICECONNECTORTERMINATION: func() Any {
		return &chime.VoiceConnectorTermination{}
	},
	DETECTIVE_GRAPH: func() Any {
		return &detective.Graph{}
	},
	KINESIS_ANALYTICSAPPLICATION: func() Any {
		return &kinesis.AnalyticsApplication{}
	},
	MSK_SCRAMSECRETASSOCIATION: func() Any {
		return &msk.ScramSecretAssociation{}
	},
	STORAGEGATEWAY_GATEWAY: func() Any {
		return &storagegateway.Gateway{}
	},
	APPRUNNER_CONNECTION: func() Any {
		return &apprunner.Connection{}
	},
	AUDITMANAGER_ASSESSMENT: func() Any {
		return &auditmanager.Assessment{}
	},
	IOT_PROVISIONINGTEMPLATE: func() Any {
		return &iot.ProvisioningTemplate{}
	},
	LAMBDA_FUNCTIONEVENTINVOKECONFIG: func() Any {
		return &lambda.FunctionEventInvokeConfig{}
	},
	DATASYNC_LOCATIONSMB: func() Any {
		return &datasync.LocationSmb{}
	},
	LIGHTSAIL_CONTAINERSERVICE: func() Any {
		return &lightsail.ContainerService{}
	},
	WAFREGIONAL_RATEBASEDRULE: func() Any {
		return &wafregional.RateBasedRule{}
	},
	IMAGEBUILDER_IMAGEPIPELINE: func() Any {
		return &imagebuilder.ImagePipeline{}
	},
	KMS_ALIAS: func() Any {
		return &aws_kms.Alias{}
	},
	IAM_USERPOLICY: func() Any {
		return &aws_iam.UserPolicy{}
	},
	EMR_INSTANCEFLEET: func() Any {
		return &emr.InstanceFleet{}
	},
	QUICKSIGHT_FOLDERMEMBERSHIP: func() Any {
		return &quicksight.FolderMembership{}
	},
	S3_BUCKETWEBSITECONFIGURATIONV2: func() Any {
		return &s3.BucketWebsiteConfigurationV2{}
	},
	ACM_CERTIFICATE: func() Any {
		return &acm.Certificate{}
	},
	APPSYNC_GRAPHQLAPI: func() Any {
		return &appsync.GraphQLApi{}
	},
	EBS_SNAPSHOTIMPORT: func() Any {
		return &ebs.SnapshotImport{}
	},
	GLUE_DATAQUALITYRULESET: func() Any {
		return &glue.DataQualityRuleset{}
	},
	OPENSEARCH_SERVERLESSLIFECYCLEPOLICY: func() Any {
		return &opensearch.ServerlessLifecyclePolicy{}
	},
	APPCONFIG_HOSTEDCONFIGURATIONVERSION: func() Any {
		return &appconfig.HostedConfigurationVersion{}
	},
	AUDITMANAGER_CONTROL: func() Any {
		return &auditmanager.Control{}
	},
	CLOUDFRONT_MONITORINGSUBSCRIPTION: func() Any {
		return &cloudfront.MonitoringSubscription{}
	},
	COMPREHEND_DOCUMENTCLASSIFIER: func() Any {
		return &comprehend.DocumentClassifier{}
	},
	DIRECTCONNECT_MACSECKEYASSOCIATION: func() Any {
		return &directconnect.MacsecKeyAssociation{}
	},
	APIGATEWAY_REQUESTVALIDATOR: func() Any {
		return &aws_apigateway.RequestValidator{}
	},
	BACKUP_GLOBALSETTINGS: func() Any {
		return &backup.GlobalSettings{}
	},
	CODECATALYST_DEVENVIRONMENT: func() Any {
		return &codecatalyst.DevEnvironment{}
	},
	EC2_VPNGATEWAYROUTEPROPAGATION: func() Any {
		return &ec2.VpnGatewayRoutePropagation{}
	},
	FMS_POLICY: func() Any {
		return &fms.Policy{}
	},
	REDSHIFT_SNAPSHOTSCHEDULE: func() Any {
		return &redshift.SnapshotSchedule{}
	},
	ROUTE53_HEALTHCHECK: func() Any {
		return &route53.HealthCheck{}
	},
	ROUTE53RECOVERYCONTROL_CLUSTER: func() Any {
		return &route53recoverycontrol.Cluster{}
	},
	S3CONTROL_ACCESSGRANT: func() Any {
		return &s3control.AccessGrant{}
	},
	S3CONTROL_BUCKETLIFECYCLECONFIGURATION: func() Any {
		return &s3control.BucketLifecycleConfiguration{}
	},
	VPCLATTICE_SERVICENETWORKSERVICEASSOCIATION: func() Any {
		return &vpclattice.ServiceNetworkServiceAssociation{}
	},
	ECS_TAG: func() Any {
		return &ecs.Tag{}
	},
	IOT_THINGGROUPMEMBERSHIP: func() Any {
		return &iot.ThingGroupMembership{}
	},
	LAMBDA_EVENTSOURCEMAPPING: func() Any {
		return &lambda.EventSourceMapping{}
	},
	CONNECT_BOTASSOCIATION: func() Any {
		return &connect.BotAssociation{}
	},
	DATASYNC_LOCATIONFSXONTAPFILESYSTEM: func() Any {
		return &datasync.LocationFsxOntapFileSystem{}
	},
	EC2TRANSITGATEWAY_MULTICASTGROUPMEMBER: func() Any {
		return &ec2transitgateway.MulticastGroupMember{}
	},
	IAM_ROLEPOLICYATTACHMENT: func() Any {
		return &aws_iam.RolePolicyAttachment{}
	},
	SSOADMIN_APPLICATIONACCESSSCOPE: func() Any {
		return &ssoadmin.ApplicationAccessScope{}
	},
	ACMPCA_CERTIFICATEAUTHORITY: func() Any {
		return &acmpca.CertificateAuthority{}
	},
	GAMELIFT_BUILD: func() Any {
		return &gamelift.Build{}
	},
	RAM_RESOURCESHARE: func() Any {
		return &ram.ResourceShare{}
	},
	DIRECTORYSERVICE_TRUST: func() Any {
		return &directoryservice.Trust{}
	},
	IAM_GROUPPOLICY: func() Any {
		return &aws_iam.GroupPolicy{}
	},
	ROUTE53_RESOLVERFIREWALLCONFIG: func() Any {
		return &route53.ResolverFirewallConfig{}
	},
	SAGEMAKER_DEVICEFLEET: func() Any {
		return &sagemaker.DeviceFleet{}
	},
	SSM_RESOURCEDATASYNC: func() Any {
		return &ssm.ResourceDataSync{}
	},
	EKS_CLUSTER: func() Any {
		return &eks.Cluster{}
	},
	FSX_LUSTREFILESYSTEM: func() Any {
		return &fsx.LustreFileSystem{}
	},
	IVS_CHANNEL: func() Any {
		return &ivs.Channel{}
	},
	S3_BUCKETPUBLICACCESSBLOCK: func() Any {
		return &s3.BucketPublicAccessBlock{}
	},
	WAF_GEOMATCHSET: func() Any {
		return &waf.GeoMatchSet{}
	},
	CLOUDFRONT_FUNCTION: func() Any {
		return &cloudfront.Function{}
	},
	GLUE_CATALOGDATABASE: func() Any {
		return &glue.CatalogDatabase{}
	},
	IAM_SERVICESPECIFICCREDENTIAL: func() Any {
		return &aws_iam.ServiceSpecificCredential{}
	},
	PINPOINT_APP: func() Any {
		return &pinpoint.App{}
	},
	DATAEXCHANGE_DATASET: func() Any {
		return &dataexchange.DataSet{}
	},
	DOCDB_ELASTICCLUSTER: func() Any {
		return &docdb.ElasticCluster{}
	},
	KMS_KEYPOLICY: func() Any {
		return &aws_kms.KeyPolicy{}
	},
	LAKEFORMATION_DATALAKESETTINGS: func() Any {
		return &lakeformation.DataLakeSettings{}
	},
	LB_TRUSTSTOREREVOCATION: func() Any {
		return &lb.TrustStoreRevocation{}
	},
	NETWORKMANAGER_LINKASSOCIATION: func() Any {
		return &networkmanager.LinkAssociation{}
	},
	PINPOINT_BAIDUCHANNEL: func() Any {
		return &pinpoint.BaiduChannel{}
	},
	REDSHIFTSERVERLESS_ENDPOINTACCESS: func() Any {
		return &redshiftserverless.EndpointAccess{}
	},
	S3_INVENTORY: func() Any {
		return &s3.Inventory{}
	},
	AMPLIFY_WEBHOOK: func() Any {
		return &amplify.Webhook{}
	},
	ECR_REPOSITORYPOLICY: func() Any {
		return &ecr.RepositoryPolicy{}
	},
	ELASTICBEANSTALK_CONFIGURATIONTEMPLATE: func() Any {
		return &elasticbeanstalk.ConfigurationTemplate{}
	},
	GRAFANA_WORKSPACESAMLCONFIGURATION: func() Any {
		return &grafana.WorkspaceSamlConfiguration{}
	},
	QUICKSIGHT_ACCOUNTSUBSCRIPTION: func() Any {
		return &quicksight.AccountSubscription{}
	},
	APPRUNNER_OBSERVABILITYCONFIGURATION: func() Any {
		return &apprunner.ObservabilityConfiguration{}
	},
	CFG_RECORDERSTATUS: func() Any {
		return &cfg.RecorderStatus{}
	},
	MACIE_FINDINGSFILTER: func() Any {
		return &macie.FindingsFilter{}
	},
	S3_BUCKETINTELLIGENTTIERINGCONFIGURATION: func() Any {
		return &s3.BucketIntelligentTieringConfiguration{}
	},
	COGNITO_USERGROUP: func() Any {
		return &cognito.UserGroup{}
	},
	DIRECTORYSERVICE_SHAREDDIRECTORYACCEPTER: func() Any {
		return &directoryservice.SharedDirectoryAccepter{}
	},
	EC2TRANSITGATEWAY_ROUTETABLEASSOCIATION: func() Any {
		return &ec2transitgateway.RouteTableAssociation{}
	},
	OPENSEARCH_SERVERLESSVPCENDPOINT: func() Any {
		return &opensearch.ServerlessVpcEndpoint{}
	},
	DIRECTCONNECT_GATEWAYASSOCIATION: func() Any {
		return &directconnect.GatewayAssociation{}
	},
	EC2_IMAGEBLOCKPUBLICACCESS: func() Any {
		return &ec2.ImageBlockPublicAccess{}
	},
	IAM_USER: func() Any {
		return &aws_iam.User{}
	},
	IVSCHAT_ROOM: func() Any {
		return &ivschat.Room{}
	},
	MEDIASTORE_CONTAINERPOLICY: func() Any {
		return &mediastore.ContainerPolicy{}
	},
	ORGANIZATIONS_POLICYATTACHMENT: func() Any {
		return &aws_organizations.PolicyAttachment{}
	},
	SNS_TOPICPOLICY: func() Any {
		return &sns.TopicPolicy{}
	},
	EBS_FASTSNAPSHOTRESTORE: func() Any {
		return &ebs.FastSnapshotRestore{}
	},
	EC2_DEFAULTVPCDHCPOPTIONS: func() Any {
		return &ec2.DefaultVpcDhcpOptions{}
	},
	LIGHTSAIL_STATICIP: func() Any {
		return &lightsail.StaticIp{}
	},
	NETWORKMANAGER_LINK: func() Any {
		return &networkmanager.Link{}
	},
	CHIME_VOICECONNECTOR: func() Any {
		return &chime.VoiceConnector{}
	},
	COSTEXPLORER_ANOMALYSUBSCRIPTION: func() Any {
		return &costexplorer.AnomalySubscription{}
	},
	ELASTICSEARCH_DOMAINSAMLOPTIONS: func() Any {
		return &elasticsearch.DomainSamlOptions{}
	},
	GLUE_RESOURCEPOLICY: func() Any {
		return &glue.ResourcePolicy{}
	},
	S3_BUCKETREPLICATIONCONFIG: func() Any {
		return &s3.BucketReplicationConfig{}
	},
	WAFREGIONAL_IPSET: func() Any {
		return &wafregional.IpSet{}
	},
	BATCH_JOBDEFINITION: func() Any {
		return &batch.JobDefinition{}
	},
	BUDGETS_BUDGETACTION: func() Any {
		return &budgets.BudgetAction{}
	},
	DIRECTCONNECT_PRIVATEVIRTUALINTERFACE: func() Any {
		return &directconnect.PrivateVirtualInterface{}
	},
	ROUTE53DOMAINS_REGISTEREDDOMAIN: func() Any {
		return &route53domains.RegisteredDomain{}
	},
	SERVICECATALOG_PORTFOLIO: func() Any {
		return &servicecatalog.Portfolio{}
	},
	APIGATEWAY_USAGEPLAN: func() Any {
		return &aws_apigateway.UsagePlan{}
	},
	ACCOUNT_ALTERNATIVECONTACT: func() Any {
		return &account.AlternativeContact{}
	},
	ECR_LIFECYCLEPOLICY: func() Any {
		return &ecr.LifecyclePolicy{}
	},
	ELASTICACHE_SERVERLESSCACHE: func() Any {
		return &elasticache.ServerlessCache{}
	},
	IAM_USERLOGINPROFILE: func() Any {
		return &aws_iam.UserLoginProfile{}
	},
	SECURITYHUB_AUTOMATIONRULE: func() Any {
		return &securityhub.AutomationRule{}
	},
	ECS_ACCOUNTSETTINGDEFAULT: func() Any {
		return &ecs.AccountSettingDefault{}
	},
	GLOBALACCELERATOR_CUSTOMROUTINGENDPOINTGROUP: func() Any {
		return &globalaccelerator.CustomRoutingEndpointGroup{}
	},
	TRANSCRIBE_VOCABULARYFILTER: func() Any {
		return &transcribe.VocabularyFilter{}
	},
	BEDROCK_PROVISIONEDMODELTHROUGHPUT: func() Any {
		return &bedrock.ProvisionedModelThroughput{}
	},
	CUSTOMERPROFILES_PROFILE: func() Any {
		return &customerprofiles.Profile{}
	},
	GAMELIFT_FLEET: func() Any {
		return &gamelift.Fleet{}
	},
	IOT_INDEXINGCONFIGURATION: func() Any {
		return &iot.IndexingConfiguration{}
	},
	MACIE2_CLASSIFICATIONJOB: func() Any {
		return &macie2.ClassificationJob{}
	},
	SECURITYHUB_STANDARDSSUBSCRIPTION: func() Any {
		return &securityhub.StandardsSubscription{}
	},
	SHIELD_PROTECTIONGROUP: func() Any {
		return &shield.ProtectionGroup{}
	},
	EC2_DEFAULTSUBNET: func() Any {
		return &ec2.DefaultSubnet{}
	},
	GRAFANA_WORKSPACEAPIKEY: func() Any {
		return &grafana.WorkspaceApiKey{}
	},
	REDSHIFTSERVERLESS_NAMESPACE: func() Any {
		return &redshiftserverless.Namespace{}
	},
	SERVERLESSREPOSITORY_CLOUDFORMATIONSTACK: func() Any {
		return &serverlessrepository.CloudFormationStack{}
	},
	RDS_CLUSTERENDPOINT: func() Any {
		return &rds.ClusterEndpoint{}
	},
	AMP_WORKSPACE: func() Any {
		return &amp.Workspace{}
	},
	APIGATEWAYV2_APIMAPPING: func() Any {
		return &apigatewayv2.ApiMapping{}
	},
	AUDITMANAGER_FRAMEWORKSHARE: func() Any {
		return &auditmanager.FrameworkShare{}
	},
	BEDROCK_CUSTOMMODEL: func() Any {
		return &bedrock.CustomModel{}
	},
	REDSHIFT_CLUSTERIAMROLES: func() Any {
		return &redshift.ClusterIamRoles{}
	},
	CONNECT_QUEUE: func() Any {
		return &connect.Queue{}
	},
	ELB_LOADBALANCERBACKENDSERVERPOLICY: func() Any {
		return &elb.LoadBalancerBackendServerPolicy{}
	},
	LB_LISTENERCERTIFICATE: func() Any {
		return &lb.ListenerCertificate{}
	},
	REKOGNITION_COLLECTION: func() Any {
		return &rekognition.Collection{}
	},
	SNS_TOPICSUBSCRIPTION: func() Any {
		return &sns.TopicSubscription{}
	},
	SYNTHETICS_GROUP: func() Any {
		return &synthetics.Group{}
	},
	CLOUDFORMATION_STACKSET: func() Any {
		return &cloudformation.StackSet{}
	},
	GUARDDUTY_DETECTOR: func() Any {
		return &guardduty.Detector{}
	},
	MSK_REPLICATOR: func() Any {
		return &msk.Replicator{}
	},
	OPENSEARCHINGEST_PIPELINE: func() Any {
		return &opensearchingest.Pipeline{}
	},
	ROUTE53_RESOLVERENDPOINT: func() Any {
		return &route53.ResolverEndpoint{}
	},
	SERVICECATALOG_PROVISIONINGARTIFACT: func() Any {
		return &servicecatalog.ProvisioningArtifact{}
	},
	SESV2_CONFIGURATIONSETEVENTDESTINATION: func() Any {
		return &sesv2.ConfigurationSetEventDestination{}
	},
	VPC_SECURITYGROUPEGRESSRULE: func() Any {
		return &vpc.SecurityGroupEgressRule{}
	},
	WAF_BYTEMATCHSET: func() Any {
		return &waf.ByteMatchSet{}
	},
	MSK_VPCCONNECTION: func() Any {
		return &msk.VpcConnection{}
	},
	ACCESSANALYZER_ARCHIVERULE: func() Any {
		return &accessanalyzer.ArchiveRule{}
	},
	APIGATEWAY_DOCUMENTATIONPART: func() Any {
		return &aws_apigateway.DocumentationPart{}
	},
	APPCONFIG_EXTENSIONASSOCIATION: func() Any {
		return &appconfig.ExtensionAssociation{}
	},
	IAM_SIGNINGCERTIFICATE: func() Any {
		return &aws_iam.SigningCertificate{}
	},
	LIGHTSAIL_STATICIPATTACHMENT: func() Any {
		return &lightsail.StaticIpAttachment{}
	},
	NETWORKMANAGER_CONNECTION: func() Any {
		return &networkmanager.Connection{}
	},
	OPSWORKS_NODEJSAPPLAYER: func() Any {
		return &opsworks.NodejsAppLayer{}
	},
	APPSYNC_TYPE: func() Any {
		return &appsync.Type{}
	},
	CFG_DELIVERYCHANNEL: func() Any {
		return &cfg.DeliveryChannel{}
	},
	EC2_VPCENDPOINTSERVICE: func() Any {
		return &ec2.VpcEndpointService{}
	},
	EVIDENTLY_SEGMENT: func() Any {
		return &evidently.Segment{}
	},
	APPCONFIG_EVENTINTEGRATION: func() Any {
		return &appconfig.EventIntegration{}
	},
	APPRUNNER_SERVICE: func() Any {
		return &apprunner.Service{}
	},
	COGNITO_IDENTITYPROVIDER: func() Any {
		return &cognito.IdentityProvider{}
	},
	ELB_LOADBALANCERCOOKIESTICKINESSPOLICY: func() Any {
		return &elb.LoadBalancerCookieStickinessPolicy{}
	},
	LEX_V2MODELSBOTLOCALE: func() Any {
		return &lex.V2modelsBotLocale{}
	},
	SSM_ASSOCIATION: func() Any {
		return &ssm.Association{}
	},
	APPMESH_ROUTE: func() Any {
		return &appmesh.Route{}
	},
	ATHENA_DATACATALOG: func() Any {
		return &athena.DataCatalog{}
	},
	AUTOSCALING_TRAFFICSOURCEATTACHMENT: func() Any {
		return &autoscaling.TrafficSourceAttachment{}
	},
	CLEANROOMS_COLLABORATION: func() Any {
		return &cleanrooms.Collaboration{}
	},
	DIRECTCONNECT_CONNECTIONCONFIRMATION: func() Any {
		return &directconnect.ConnectionConfirmation{}
	},
	IAM_USERGROUPMEMBERSHIP: func() Any {
		return &aws_iam.UserGroupMembership{}
	},
	SQS_QUEUE: func() Any {
		return &sqs.Queue{}
	},
	QUICKSIGHT_DATASOURCE: func() Any {
		return &quicksight.DataSource{}
	},
	S3_BUCKETPOLICY: func() Any {
		return &s3.BucketPolicy{}
	},
	SAGEMAKER_APPIMAGECONFIG: func() Any {
		return &sagemaker.AppImageConfig{}
	},
	SSOADMIN_APPLICATIONASSIGNMENTCONFIGURATION: func() Any {
		return &ssoadmin.ApplicationAssignmentConfiguration{}
	},
	BACKUP_FRAMEWORK: func() Any {
		return &backup.Framework{}
	},
	ELASTICSEARCH_VPCENDPOINT: func() Any {
		return &elasticsearch.VpcEndpoint{}
	},
	FINSPACE_KXDATAVIEW: func() Any {
		return &finspace.KxDataview{}
	},
	WAFREGIONAL_REGEXMATCHSET: func() Any {
		return &wafregional.RegexMatchSet{}
	},
	WAFREGIONAL_WEBACL: func() Any {
		return &wafregional.WebAcl{}
	},
	APIGATEWAY_MODEL: func() Any {
		return &aws_apigateway.Model{}
	},
	CONNECT_QUICKCONNECT: func() Any {
		return &connect.QuickConnect{}
	},
	GLOBALACCELERATOR_CUSTOMROUTINGACCELERATOR: func() Any {
		return &globalaccelerator.CustomRoutingAccelerator{}
	},
	LAKEFORMATION_RESOURCELFTAGS: func() Any {
		return &lakeformation.ResourceLfTags{}
	},
	SAGEMAKER_MODELPACKAGEGROUPPOLICY: func() Any {
		return &sagemaker.ModelPackageGroupPolicy{}
	},
	SESV2_EMAILIDENTITYMAILFROMATTRIBUTES: func() Any {
		return &sesv2.EmailIdentityMailFromAttributes{}
	},
	SSM_ACTIVATION: func() Any {
		return &ssm.Activation{}
	},
	AMPLIFY_DOMAINASSOCIATION: func() Any {
		return &amplify.DomainAssociation{}
	},
	APPAUTOSCALING_TARGET: func() Any {
		return &appautoscaling.Target{}
	},
	CODEPIPELINE_WEBHOOK: func() Any {
		return &codepipeline.Webhook{}
	},
	EC2_INSTANCE: func() Any {
		return &ec2.Instance{}
	},
	NEPTUNE_CLUSTERSNAPSHOT: func() Any {
		return &neptune.ClusterSnapshot{}
	},
	S3OUTPOSTS_ENDPOINT: func() Any {
		return &s3outposts.Endpoint{}
	},
	SERVICEQUOTAS_TEMPLATEASSOCIATION: func() Any {
		return &servicequotas.TemplateAssociation{}
	},
	AUTOSCALING_TAG: func() Any {
		return &autoscaling.Tag{}
	},
	CHIME_VOICECONNECTORGROUP: func() Any {
		return &chime.VoiceConnectorGroup{}
	},
	CLOUDWATCH_INTERNETMONITOR: func() Any {
		return &cloudwatch.InternetMonitor{}
	},
	DIRECTORYSERVICE_CONDITIONALFORWADER: func() Any {
		return &directoryservice.ConditionalForwader{}
	},
	DMS_S3ENDPOINT: func() Any {
		return &dms.S3Endpoint{}
	},
	MACIE2_INVITATIONACCEPTER: func() Any {
		return &macie2.InvitationAccepter{}
	},
	QUICKSIGHT_GROUP: func() Any {
		return &quicksight.Group{}
	},
	REDSHIFT_DATASHAREAUTHORIZATION: func() Any {
		return &redshift.DataShareAuthorization{}
	},
	DATASYNC_S3LOCATION: func() Any {
		return &datasync.S3Location{}
	},
	EFS_MOUNTTARGET: func() Any {
		return &efs.MountTarget{}
	},
	OPENSEARCH_INBOUNDCONNECTIONACCEPTER: func() Any {
		return &opensearch.InboundConnectionAccepter{}
	},
	RDS_EVENTSUBSCRIPTION: func() Any {
		return &rds.EventSubscription{}
	},
	TRANSFER_ACCESS: func() Any {
		return &transfer.Access{}
	},
	WAFV2_WEBACL: func() Any {
		return &wafv2.WebAcl{}
	},
	ACM_CERTIFICATEVALIDATION: func() Any {
		return &acm.CertificateValidation{}
	},
	APPSTREAM_FLEETSTACKASSOCIATION: func() Any {
		return &appstream.FleetStackAssociation{}
	},
	CONNECT_INSTANCE: func() Any {
		return &connect.Instance{}
	},
	EMRCONTAINERS_JOBTEMPLATE: func() Any {
		return &emrcontainers.JobTemplate{}
	},
	MEDIALIVE_MULTIPLEX: func() Any {
		return &medialive.Multiplex{}
	},
	TRANSCRIBE_LANGUAGEMODEL: func() Any {
		return &transcribe.LanguageModel{}
	},
	VPCLATTICE_TARGETGROUPATTACHMENT: func() Any {
		return &vpclattice.TargetGroupAttachment{}
	},
	AUDITMANAGER_ORGANIZATIONADMINACCOUNTREGISTRATION: func() Any {
		return &auditmanager.OrganizationAdminAccountRegistration{}
	},
	DIRECTCONNECT_CONNECTIONASSOCIATION: func() Any {
		return &directconnect.ConnectionAssociation{}
	},
	DOCDB_EVENTSUBSCRIPTION: func() Any {
		return &docdb.EventSubscription{}
	},
	REDSHIFT_RESOURCEPOLICY: func() Any {
		return &redshift.ResourcePolicy{}
	},
	ROUTE53RECOVERYREADINESS_CELL: func() Any {
		return &route53recoveryreadiness.Cell{}
	},
	RUM_APPMONITOR: func() Any {
		return &rum.AppMonitor{}
	},
	SES_CONFIGURATIONSET: func() Any {
		return &ses.ConfigurationSet{}
	},
	APPSTREAM_IMAGEBUILDER: func() Any {
		return &appstream.ImageBuilder{}
	},
	CONTROLTOWER_CONTROLTOWERCONTROL: func() Any {
		return &controltower.ControlTowerControl{}
	},
	LIGHTSAIL_DISK_ATTACHMENT: func() Any {
		return &lightsail.Disk_attachment{}
	},
	NETWORKMANAGER_SITETOSITEVPNATTACHMENT: func() Any {
		return &networkmanager.SiteToSiteVpnAttachment{}
	},
	SERVICECATALOG_TAGOPTION: func() Any {
		return &servicecatalog.TagOption{}
	},
	SESV2_EMAILIDENTITY: func() Any {
		return &sesv2.EmailIdentity{}
	},
	COSTEXPLORER_COSTCATEGORY: func() Any {
		return &costexplorer.CostCategory{}
	},
	EKS_ACCESSENTRY: func() Any {
		return &eks.AccessEntry{}
	},
	RDS_GLOBALCLUSTER: func() Any {
		return &rds.GlobalCluster{}
	},
	SERVICEDISCOVERY_SERVICE: func() Any {
		return &servicediscovery.Service{}
	},
	CHIME_VOICECONNECTORSTREAMING: func() Any {
		return &chime.VoiceConnectorStreaming{}
	},
	CLOUDFORMATION_CLOUDFORMATIONTYPE: func() Any {
		return &cloudformation.CloudFormationType{}
	},
	DATASYNC_LOCATIONAZUREBLOB: func() Any {
		return &datasync.LocationAzureBlob{}
	},
	FIS_EXPERIMENTTEMPLATE: func() Any {
		return &fis.ExperimentTemplate{}
	},
	GUARDDUTY_THREATINTELSET: func() Any {
		return &guardduty.ThreatIntelSet{}
	},
	LAKEFORMATION_PERMISSIONS: func() Any {
		return &lakeformation.Permissions{}
	},
	SSOADMIN_PERMISSIONSETINLINEPOLICY: func() Any {
		return &ssoadmin.PermissionSetInlinePolicy{}
	},
	VPCLATTICE_SERVICE: func() Any {
		return &vpclattice.Service{}
	},
	EC2_SECURITYGROUPRULE: func() Any {
		return &ec2.SecurityGroupRule{}
	},
	LIGHTSAIL_BUCKET: func() Any {
		return &lightsail.Bucket{}
	},
	S3_DIRECTORYBUCKET: func() Any {
		return &s3.DirectoryBucket{}
	},
	S3CONTROL_ACCESSGRANTSINSTANCE: func() Any {
		return &s3control.AccessGrantsInstance{}
	},
	SSM_MAINTENANCEWINDOW: func() Any {
		return &ssm.MaintenanceWindow{}
	},
	APPLICATIONINSIGHTS_APPLICATION: func() Any {
		return &applicationinsights.Application{}
	},
	APPMESH_VIRTUALROUTER: func() Any {
		return &appmesh.VirtualRouter{}
	},
	SIGNER_SIGNINGJOB: func() Any {
		return &signer.SigningJob{}
	},
	WAFREGIONAL_RULE: func() Any {
		return &wafregional.Rule{}
	},
	DETECTIVE_ORGANIZATIONCONFIGURATION: func() Any {
		return &detective.OrganizationConfiguration{}
	},
	EC2TRANSITGATEWAY_ROUTE: func() Any {
		return &ec2transitgateway.Route{}
	},
	EC2TRANSITGATEWAY_VPCATTACHMENTACCEPTER: func() Any {
		return &ec2transitgateway.VpcAttachmentAccepter{}
	},
	EKS_ADDON: func() Any {
		return &eks.Addon{}
	},
	LIGHTSAIL_INSTANCE: func() Any {
		return &lightsail.Instance{}
	},
	APPFLOW_FLOW: func() Any {
		return &appflow.Flow{}
	},
	EC2_LOCALGATEWAYROUTE: func() Any {
		return &ec2.LocalGatewayRoute{}
	},
	EC2_TAG: func() Any {
		return &ec2.Tag{}
	},
	EC2_VOLUMEATTACHMENT: func() Any {
		return &ec2.VolumeAttachment{}
	},
	REDSHIFT_SCHEDULEDACTION: func() Any {
		return &redshift.ScheduledAction{}
	},
	WAF_IPSET: func() Any {
		return &waf.IpSet{}
	},
	APIGATEWAY_METHODSETTINGS: func() Any {
		return &aws_apigateway.MethodSettings{}
	},
	CFG_RECORDER: func() Any {
		return &cfg.Recorder{}
	},
	MACIE2_ACCOUNT: func() Any {
		return &macie2.Account{}
	},
	MEMORYDB_ACL: func() Any {
		return &memorydb.Acl{}
	},
	MEMORYDB_SUBNETGROUP: func() Any {
		return &memorydb.SubnetGroup{}
	},
	NETWORKMANAGER_VPCATTACHMENT: func() Any {
		return &networkmanager.VpcAttachment{}
	},
	REDSHIFT_CLUSTERSNAPSHOT: func() Any {
		return &redshift.ClusterSnapshot{}
	},
	BACKUP_VAULTPOLICY: func() Any {
		return &backup.VaultPolicy{}
	},
	DIRECTCONNECT_BGPPEER: func() Any {
		return &directconnect.BgpPeer{}
	},
	REDSHIFTSERVERLESS_SNAPSHOT: func() Any {
		return &redshiftserverless.Snapshot{}
	},
	SAGEMAKER_ENDPOINTCONFIGURATION: func() Any {
		return &sagemaker.EndpointConfiguration{}
	},
	SAGEMAKER_PROJECT: func() Any {
		return &sagemaker.Project{}
	},
	APIGATEWAY_INTEGRATIONRESPONSE: func() Any {
		return &aws_apigateway.IntegrationResponse{}
	},
	DLM_LIFECYCLEPOLICY: func() Any {
		return &dlm.LifecyclePolicy{}
	},
	FINSPACE_KXENVIRONMENT: func() Any {
		return &finspace.KxEnvironment{}
	},
	LAMBDA_PERMISSION: func() Any {
		return &lambda.Permission{}
	},
	LIGHTSAIL_LBATTACHMENT: func() Any {
		return &lightsail.LbAttachment{}
	},
	SAGEMAKER_FLOWDEFINITION: func() Any {
		return &sagemaker.FlowDefinition{}
	},
	EBS_DEFAULTKMSKEY: func() Any {
		return &ebs.DefaultKmsKey{}
	},
	LAMBDA_LAYERVERSION: func() Any {
		return &lambda.LayerVersion{}
	},
	OPSWORKS_USERPROFILE: func() Any {
		return &opsworks.UserProfile{}
	},
	RDS_PARAMETERGROUP: func() Any {
		return &rds.ParameterGroup{}
	},
	SAGEMAKER_IMAGEVERSION: func() Any {
		return &sagemaker.ImageVersion{}
	},
	SAGEMAKER_MONITORINGSCHEDULE: func() Any {
		return &sagemaker.MonitoringSchedule{}
	},
	SECRETSMANAGER_SECRETVERSION: func() Any {
		return &secretsmanager.SecretVersion{}
	},
	SESV2_ACCOUNTVDMATTRIBUTES: func() Any {
		return &sesv2.AccountVdmAttributes{}
	},
	ACMPCA_CERTIFICATE: func() Any {
		return &acmpca.Certificate{}
	},
	APPRUNNER_DEPLOYMENT: func() Any {
		return &apprunner.Deployment{}
	},
	IMAGEBUILDER_IMAGE: func() Any {
		return &imagebuilder.Image{}
	},
	IOT_THING: func() Any {
		return &iot.Thing{}
	},
	MSKCONNECT_WORKERCONFIGURATION: func() Any {
		return &mskconnect.WorkerConfiguration{}
	},
	SSMINCIDENTS_RESPONSEPLAN: func() Any {
		return &ssmincidents.ResponsePlan{}
	},
	APPSTREAM_USERSTACKASSOCIATION: func() Any {
		return &appstream.UserStackAssociation{}
	},
	EC2_PLACEMENTGROUP: func() Any {
		return &ec2.PlacementGroup{}
	},
	SAGEMAKER_USERPROFILE: func() Any {
		return &sagemaker.UserProfile{}
	},
	WAFREGIONAL_SQLINJECTIONMATCHSET: func() Any {
		return &wafregional.SqlInjectionMatchSet{}
	},
	EFS_BACKUPPOLICY: func() Any {
		return &efs.BackupPolicy{}
	},
	S3CONTROL_ACCESSPOINTPOLICY: func() Any {
		return &s3control.AccessPointPolicy{}
	},
	APPSYNC_APIKEY: func() Any {
		return &appsync.ApiKey{}
	},
	EC2_TRAFFICMIRRORSESSION: func() Any {
		return &ec2.TrafficMirrorSession{}
	},
	IVSCHAT_LOGGINGCONFIGURATION: func() Any {
		return &ivschat.LoggingConfiguration{}
	},
	PINPOINT_APNSVOIPSANDBOXCHANNEL: func() Any {
		return &pinpoint.ApnsVoipSandboxChannel{}
	},
	RDS_PROXYTARGET: func() Any {
		return &rds.ProxyTarget{}
	},
	S3_BUCKETLIFECYCLECONFIGURATIONV2: func() Any {
		return &s3.BucketLifecycleConfigurationV2{}
	},
	SFN_ACTIVITY: func() Any {
		return &sfn.Activity{}
	},
	AMP_RULEGROUPNAMESPACE: func() Any {
		return &amp.RuleGroupNamespace{}
	},
	BACKUP_VAULT: func() Any {
		return &backup.Vault{}
	},
	CLOUDFORMATION_STACK: func() Any {
		return &cloudformation.Stack{}
	},
	CODEBUILD_SOURCECREDENTIAL: func() Any {
		return &codebuild.SourceCredential{}
	},
	LAMBDA_CODESIGNINGCONFIG: func() Any {
		return &lambda.CodeSigningConfig{}
	},
	RDS_PROXY: func() Any {
		return &rds.Proxy{}
	},
	RDS_SNAPSHOTCOPY: func() Any {
		return &rds.SnapshotCopy{}
	},
	AMPLIFY_BACKENDENVIRONMENT: func() Any {
		return &amplify.BackendEnvironment{}
	},
	APPSYNC_DOMAINNAME: func() Any {
		return &appsync.DomainName{}
	},
	OPENSEARCH_OUTBOUNDCONNECTION: func() Any {
		return &opensearch.OutboundConnection{}
	},
	SSOADMIN_MANAGEDPOLICYATTACHMENT: func() Any {
		return &ssoadmin.ManagedPolicyAttachment{}
	},
	WAFV2_IPSET: func() Any {
		return &wafv2.IpSet{}
	},
	CODEGURUPROFILER_PROFILINGGROUP: func() Any {
		return &codeguruprofiler.ProfilingGroup{}
	},
	IOT_LOGGINGOPTIONS: func() Any {
		return &iot.LoggingOptions{}
	},
	S3_BUCKETACLV2: func() Any {
		return &s3.BucketAclV2{}
	},
	S3_BUCKETVERSIONINGV2: func() Any {
		return &s3.BucketVersioningV2{}
	},
	DATASYNC_LOCATIONHDFS: func() Any {
		return &datasync.LocationHdfs{}
	},
	EC2_VPCIPAMPOOLCIDRALLOCATION: func() Any {
		return &ec2.VpcIpamPoolCidrAllocation{}
	},
	EC2_VPCIPAMRESOURCEDISCOVERY: func() Any {
		return &ec2.VpcIpamResourceDiscovery{}
	},
	ELB_ATTACHMENT: func() Any {
		return &elb.Attachment{}
	},
	GLUE_WORKFLOW: func() Any {
		return &glue.Workflow{}
	},
	MACIE2_CLASSIFICATIONEXPORTCONFIGURATION: func() Any {
		return &macie2.ClassificationExportConfiguration{}
	},
	STORAGEGATEWAY_STOREDISCSIVOLUME: func() Any {
		return &storagegateway.StoredIscsiVolume{}
	},
	CODEDEPLOY_DEPLOYMENTCONFIG: func() Any {
		return &codedeploy.DeploymentConfig{}
	},
	EC2_VPCENDPOINT: func() Any {
		return &ec2.VpcEndpoint{}
	},
	APIGATEWAYV2_MODEL: func() Any {
		return &apigatewayv2.Model{}
	},
	AUTOSCALINGPLANS_SCALINGPLAN: func() Any {
		return &autoscalingplans.ScalingPlan{}
	},
	CLOUDFRONT_CACHEPOLICY: func() Any {
		return &cloudfront.CachePolicy{}
	},
	GRAFANA_WORKSPACE: func() Any {
		return &grafana.Workspace{}
	},
	LAKEFORMATION_RESOURCE: func() Any {
		return &lakeformation.Resource{}
	},
	NETWORKFIREWALL_FIREWALL: func() Any {
		return &networkfirewall.Firewall{}
	},
	APPSYNC_FUNCTION: func() Any {
		return &appsync.Function{}
	},
	CFG_ORGANIZATIONCONFORMANCEPACK: func() Any {
		return &cfg.OrganizationConformancePack{}
	},
	COGNITO_USERPOOLDOMAIN: func() Any {
		return &cognito.UserPoolDomain{}
	},
	EC2_VPNCONNECTIONROUTE: func() Any {
		return &ec2.VpnConnectionRoute{}
	},
	FINSPACE_KXSCALINGGROUP: func() Any {
		return &finspace.KxScalingGroup{}
	},
	IOT_CACERTIFICATE: func() Any {
		return &iot.CaCertificate{}
	},
	MEMORYDB_SNAPSHOT: func() Any {
		return &memorydb.Snapshot{}
	},
	CLOUDWATCH_DASHBOARD: func() Any {
		return &cloudwatch.Dashboard{}
	},
	LEX_V2MODELSBOT: func() Any {
		return &lex.V2modelsBot{}
	},
	QUICKSIGHT_NAMESPACE: func() Any {
		return &quicksight.Namespace{}
	},
	SES_EVENTDESTINATION: func() Any {
		return &ses.EventDestination{}
	},
	VERIFIEDACCESS_ENDPOINT: func() Any {
		return &verifiedaccess.Endpoint{}
	},
	BATCH_SCHEDULINGPOLICY: func() Any {
		return &batch.SchedulingPolicy{}
	},
	CLOUDWATCH_EVENTARCHIVE: func() Any {
		return &cloudwatch.EventArchive{}
	},
	MSKCONNECT_CUSTOMPLUGIN: func() Any {
		return &mskconnect.CustomPlugin{}
	},
	SAGEMAKER_MODEL: func() Any {
		return &sagemaker.Model{}
	},
	CLOUDWATCH_EVENTAPIDESTINATION: func() Any {
		return &cloudwatch.EventApiDestination{}
	},
	COGNITO_IDENTITYPOOLROLEATTACHMENT: func() Any {
		return &cognito.IdentityPoolRoleAttachment{}
	},
	EC2CLIENTVPN_NETWORKASSOCIATION: func() Any {
		return &ec2clientvpn.NetworkAssociation{}
	},
	OPENSEARCH_DOMAINSAMLOPTIONS: func() Any {
		return &opensearch.DomainSamlOptions{}
	},
	AMPLIFY_BRANCH: func() Any {
		return &amplify.Branch{}
	},
	COGNITO_IDENTITYPOOL: func() Any {
		return &cognito.IdentityPool{}
	},
	EC2_VPCDHCPOPTIONSASSOCIATION: func() Any {
		return &ec2.VpcDhcpOptionsAssociation{}
	},
	ROUTE53_RESOLVERDNSSECCONFIG: func() Any {
		return &route53.ResolverDnsSecConfig{}
	},
	EC2_FLEET: func() Any {
		return &ec2.Fleet{}
	},
	GUARDDUTY_PUBLISHINGDESTINATION: func() Any {
		return &guardduty.PublishingDestination{}
	},
	SERVICEQUOTAS_SERVICEQUOTA: func() Any {
		return &servicequotas.ServiceQuota{}
	},
	STORAGEGATEWAY_SMBFILESHARE: func() Any {
		return &storagegateway.SmbFileShare{}
	},
	DEVICEFARM_DEVICEPOOL: func() Any {
		return &devicefarm.DevicePool{}
	},
	EC2TRANSITGATEWAY_INSTANCECONNECTENDPOINT: func() Any {
		return &ec2transitgateway.InstanceConnectEndpoint{}
	},
	MQ_CONFIGURATION: func() Any {
		return &mq.Configuration{}
	},
	QLDB_LEDGER: func() Any {
		return &qldb.Ledger{}
	},
	SES_DOMAINIDENTITY: func() Any {
		return &ses.DomainIdentity{}
	},
	APPAUTOSCALING_POLICY: func() Any {
		return &appautoscaling.Policy{}
	},
	CHIME_SDKVOICESIPMEDIAAPPLICATION: func() Any {
		return &chime.SdkvoiceSipMediaApplication{}
	},
	DAX_PARAMETERGROUP: func() Any {
		return &dax.ParameterGroup{}
	},
	IAM_POLICYATTACHMENT: func() Any {
		return &aws_iam.PolicyAttachment{}
	},
	S3CONTROL_MULTIREGIONACCESSPOINTPOLICY: func() Any {
		return &s3control.MultiRegionAccessPointPolicy{}
	},
	EFS_FILESYSTEMPOLICY: func() Any {
		return &efs.FileSystemPolicy{}
	},
	IMAGEBUILDER_INFRASTRUCTURECONFIGURATION: func() Any {
		return &imagebuilder.InfrastructureConfiguration{}
	},
	LIGHTSAIL_KEYPAIR: func() Any {
		return &lightsail.KeyPair{}
	},
	EC2_FLOWLOG: func() Any {
		return &ec2.FlowLog{}
	},
	EC2_SECURITYGROUPASSOCIATION: func() Any {
		return &ec2.SecurityGroupAssociation{}
	},
	EC2_VPC: func() Any {
		return &ec2.Vpc{}
	},
	EKS_IDENTITYPROVIDERCONFIG: func() Any {
		return &eks.IdentityProviderConfig{}
	},
	ELASTICACHE_CLUSTER: func() Any {
		return &elasticache.Cluster{}
	},
	SHIELD_PROTECTION: func() Any {
		return &shield.Protection{}
	},
	// GCP
	LOGGING_ORGANIZATIONBUCKETCONFIG: func() Any {
		return &logging.OrganizationBucketConfig{}
	},
	NETWORKSERVICES_EDGECACHEKEYSET: func() Any {
		return &networkservices.EdgeCacheKeyset{}
	},
	COMPUTE_HEALTHCHECK: func() Any {
		return &compute.HealthCheck{}
	},
	COMPUTE_INSTANCESETTINGS: func() Any {
		return &compute.InstanceSettings{}
	},
	DATACATALOG_TAXONOMYIAMBINDING: func() Any {
		return &datacatalog.TaxonomyIamBinding{}
	},
	IAP_WEBTYPEAPPENGINGIAMPOLICY: func() Any {
		return &iap.WebTypeAppEngingIamPolicy{}
	},
	CLOUDDEPLOY_CUSTOMTARGETTYPE: func() Any {
		return &clouddeploy.CustomTargetType{}
	},
	CLOUDRUNV2_JOBIAMBINDING: func() Any {
		return &cloudrunv2.JobIamBinding{}
	},
	PUBSUB_TOPICIAMMEMBER: func() Any {
		return &pubsub.TopicIAMMember{}
	},
	BIGQUERYDATAPOLICY_DATAPOLICYIAMBINDING: func() Any {
		return &bigquerydatapolicy.DataPolicyIamBinding{}
	},
	SOURCEREPO_REPOSITORY: func() Any {
		return &sourcerepo.Repository{}
	},
	BIGQUERYANALYTICSHUB_DATAEXCHANGEIAMPOLICY: func() Any {
		return &bigqueryanalyticshub.DataExchangeIamPolicy{}
	},
	COMPUTE_DISK: func() Any {
		return &compute.Disk{}
	},
	COMPUTE_FIREWALLPOLICYRULE: func() Any {
		return &compute.FirewallPolicyRule{}
	},
	LOGGING_BILLINGACCOUNTSINK: func() Any {
		return &logging.BillingAccountSink{}
	},
	COMPUTE_SUBNETWORKIAMPOLICY: func() Any {
		return &compute.SubnetworkIAMPolicy{}
	},
	FIREBASE_APPCHECKRECAPTCHAV3CONFIG: func() Any {
		return &firebase.AppCheckRecaptchaV3Config{}
	},
	FOLDER_ACCESSAPPROVALSETTINGS: func() Any {
		return &folder.AccessApprovalSettings{}
	},
	ESSENTIALCONTACTS_DOCUMENTAIPROCESSORDEFAULTVERSION: func() Any {
		return &essentialcontacts.DocumentAiProcessorDefaultVersion{}
	},
	FIREBASE_APPCHECKAPPATTESTCONFIG: func() Any {
		return &firebase.AppCheckAppAttestConfig{}
	},
	FIREBASE_HOSTINGSITE: func() Any {
		return &firebase.HostingSite{}
	},
	ORGANIZATIONS_IAMCUSTOMROLE: func() Any {
		return &organizations.IAMCustomRole{}
	},
	APIGEE_NATADDRESS: func() Any {
		return &apigee.NatAddress{}
	},
	CLOUDFUNCTIONSV2_FUNCTIONIAMBINDING: func() Any {
		return &cloudfunctionsv2.FunctionIamBinding{}
	},
	LOGGING_PROJECTBUCKETCONFIG: func() Any {
		return &logging.ProjectBucketConfig{}
	},
	IAP_TUNNELINSTANCEIAMMEMBER: func() Any {
		return &iap.TunnelInstanceIAMMember{}
	},
	KMS_KEYRINGIAMBINDING: func() Any {
		return &kms.KeyRingIAMBinding{}
	},
	STORAGE_BUCKETIAMMEMBER: func() Any {
		return &storage.BucketIAMMember{}
	},
	APIGEE_ENVKEYSTORE: func() Any {
		return &apigee.EnvKeystore{}
	},
	CLOUDDOMAINS_REGISTRATION: func() Any {
		return &clouddomains.Registration{}
	},
	COMPUTE_NETWORKPEERING: func() Any {
		return &compute.NetworkPeering{}
	},
	DATACATALOG_TAGTEMPLATEIAMBINDING: func() Any {
		return &datacatalog.TagTemplateIamBinding{}
	},
	COMPUTE_DISKIAMPOLICY: func() Any {
		return &compute.DiskIamPolicy{}
	},
	FILESTORE_INSTANCE: func() Any {
		return &filestore.Instance{}
	},
	NETWORKSECURITY_CLIENTTLSPOLICY: func() Any {
		return &networksecurity.ClientTlsPolicy{}
	},
	IAP_BRAND: func() Any {
		return &iap.Brand{}
	},
	COMPUTE_ROUTERNAT: func() Any {
		return &compute.RouterNat{}
	},
	SECURITYCENTER_MUTECONFIG: func() Any {
		return &securitycenter.MuteConfig{}
	},
	APIGATEWAY_GATEWAY: func() Any {
		return &apigateway.Gateway{}
	},
	COMPUTE_MACHINEIMAGE: func() Any {
		return &compute.MachineImage{}
	},
	BIGTABLE_TABLEIAMBINDING: func() Any {
		return &bigtable.TableIamBinding{}
	},
	LOGGING_ORGANIZATIONSINK: func() Any {
		return &logging.OrganizationSink{}
	},
	COMPUTE_VPNTUNNEL: func() Any {
		return &compute.VPNTunnel{}
	},
	OSCONFIG_OSPOLICYASSIGNMENT: func() Any {
		return &osconfig.OsPolicyAssignment{}
	},
	FIREBASE_HOSTINGCHANNEL: func() Any {
		return &firebase.HostingChannel{}
	},
	GKEBACKUP_RESTOREPLANIAMMEMBER: func() Any {
		return &gkebackup.RestorePlanIamMember{}
	},
	IDENTITYPLATFORM_TENANTDEFAULTSUPPORTEDIDPCONFIG: func() Any {
		return &identityplatform.TenantDefaultSupportedIdpConfig{}
	},
	COMPUTE_NODETEMPLATE: func() Any {
		return &compute.NodeTemplate{}
	},
	SERVICEUSAGE_CONSUMERQUOTAOVERRIDE: func() Any {
		return &serviceusage.ConsumerQuotaOverride{}
	},
	SQL_SOURCEREPRESENTATIONINSTANCE: func() Any {
		return &sql.SourceRepresentationInstance{}
	},
	DATALOSS_PREVENTIONINSPECTTEMPLATE: func() Any {
		return &dataloss.PreventionInspectTemplate{}
	},
	DATAPLEX_TASKIAMMEMBER: func() Any {
		return &dataplex.TaskIamMember{}
	},
	MONITORING_NOTIFICATIONCHANNEL: func() Any {
		return &monitoring.NotificationChannel{}
	},
	DATAPROC_AUTOSCALINGPOLICYIAMPOLICY: func() Any {
		return &dataproc.AutoscalingPolicyIamPolicy{}
	},
	TAGS_LOCATIONTAGBINDING: func() Any {
		return &tags.LocationTagBinding{}
	},
	ALLOYDB_USER: func() Any {
		return &alloydb.User{}
	},
	CLOUDFUNCTIONS_FUNCTIONIAMMEMBER: func() Any {
		return &cloudfunctions.FunctionIamMember{}
	},
	CLOUDRUNV2_JOBIAMMEMBER: func() Any {
		return &cloudrunv2.JobIamMember{}
	},
	COMPUTE_BACKENDBUCKET: func() Any {
		return &compute.BackendBucket{}
	},
	SECURITYCENTER_SOURCEIAMPOLICY: func() Any {
		return &securitycenter.SourceIamPolicy{}
	},
	CONTAINER_NODEPOOL: func() Any {
		return &container.NodePool{}
	},
	SECURITYCENTER_INSTANCEIAMPOLICY: func() Any {
		return &securitycenter.InstanceIamPolicy{}
	},
	STORAGE_TRANSFERJOB: func() Any {
		return &storage.TransferJob{}
	},
	LOGGING_ORGANIZATIONSETTINGS: func() Any {
		return &logging.OrganizationSettings{}
	},
	HEALTHCARE_DATASETIAMPOLICY: func() Any {
		return &healthcare.DatasetIamPolicy{}
	},
	IAP_TUNNELINSTANCEIAMBINDING: func() Any {
		return &iap.TunnelInstanceIAMBinding{}
	},
	KMS_CRYPTOKEYVERSION: func() Any {
		return &kms.CryptoKeyVersion{}
	},
	COMPUTE_INSTANCEFROMMACHINEIMAGE: func() Any {
		return &compute.InstanceFromMachineImage{}
	},
	ENDPOINTS_SERVICEIAMBINDING: func() Any {
		return &endpoints.ServiceIamBinding{}
	},
	ACCESSCONTEXTMANAGER_ACCESSPOLICYIAMPOLICY: func() Any {
		return &accesscontextmanager.AccessPolicyIamPolicy{}
	},
	APIGATEWAY_API: func() Any {
		return &apigateway.Api{}
	},
	CLOUDDEPLOY_TARGETIAMBINDING: func() Any {
		return &clouddeploy.TargetIamBinding{}
	},
	NETWORKSECURITY_ADDRESSGROUPIAMMEMBER: func() Any {
		return &networksecurity.AddressGroupIamMember{}
	},
	BIGQUERY_DATASET: func() Any {
		return &bigquery.Dataset{}
	},
	BILLING_SUBACCOUNT: func() Any {
		return &billing.SubAccount{}
	},
	IAP_TUNNELINSTANCEIAMPOLICY: func() Any {
		return &iap.TunnelInstanceIAMPolicy{}
	},
	IDENTITYPLATFORM_PROJECTDEFAULTCONFIG: func() Any {
		return &identityplatform.ProjectDefaultConfig{}
	},
	APPENGINE_DOMAINMAPPING: func() Any {
		return &appengine.DomainMapping{}
	},
	CLOUDRUNV2_SERVICEIAMMEMBER: func() Any {
		return &cloudrunv2.ServiceIamMember{}
	},
	COMPUTE_BACKENDSERVICEIAMPOLICY: func() Any {
		return &compute.BackendServiceIamPolicy{}
	},
	COMPUTE_PROJECTDEFAULTNETWORKTIER: func() Any {
		return &compute.ProjectDefaultNetworkTier{}
	},
	CERTIFICATEMANAGER_CERTIFICATEMAP: func() Any {
		return &certificatemanager.CertificateMap{}
	},
	SECURESOURCEMANAGER_INSTANCEIAMBINDING: func() Any {
		return &securesourcemanager.InstanceIamBinding{}
	},
	APIGEE_TARGETSERVER: func() Any {
		return &apigee.TargetServer{}
	},
	NETWORKCONNECTIVITY_HUB: func() Any {
		return &networkconnectivity.Hub{}
	},
	NETWORKSECURITY_SECURITYPROFILE: func() Any {
		return &networksecurity.SecurityProfile{}
	},
	VERTEX_AIMETADATASTORE: func() Any {
		return &vertex.AiMetadataStore{}
	},
	PROJECTS_IAMAUDITCONFIG: func() Any {
		return &projects.IAMAuditConfig{}
	},
	DNS_DNSMANAGEDZONEIAMPOLICY: func() Any {
		return &dns.DnsManagedZoneIamPolicy{}
	},
	COMPUTE_INSTANCE: func() Any {
		return &compute.Instance{}
	},
	COMPUTE_INSTANCEGROUPMANAGER: func() Any {
		return &compute.InstanceGroupManager{}
	},
	ACCESSCONTEXTMANAGER_SERVICEPERIMETERRESOURCE: func() Any {
		return &accesscontextmanager.ServicePerimeterResource{}
	},
	SECURITYCENTER_FOLDERCUSTOMMODULE: func() Any {
		return &securitycenter.FolderCustomModule{}
	},
	STORAGE_TRANSFERAGENTPOOL: func() Any {
		return &storage.TransferAgentPool{}
	},
	PROJECTS_USAGEEXPORTBUCKET: func() Any {
		return &projects.UsageExportBucket{}
	},
	CLOUDBUILD_BITBUCKETSERVERCONFIG: func() Any {
		return &cloudbuild.BitbucketServerConfig{}
	},
	GKEHUB_NAMESPACE: func() Any {
		return &gkehub.Namespace{}
	},
	IAP_WEBIAMPOLICY: func() Any {
		return &iap.WebIamPolicy{}
	},
	COMPUTE_FIREWALL: func() Any {
		return &compute.Firewall{}
	},
	COMPUTE_REGIONDISKRESOURCEPOLICYATTACHMENT: func() Any {
		return &compute.RegionDiskResourcePolicyAttachment{}
	},
	DNS_POLICY: func() Any {
		return &dns.Policy{}
	},
	DATACATALOG_TAGTEMPLATE: func() Any {
		return &datacatalog.TagTemplate{}
	},
	CLOUDIDENTITY_GROUP: func() Any {
		return &cloudidentity.Group{}
	},
	FIRESTORE_DATABASE: func() Any {
		return &firestore.Database{}
	},
	IAM_WORKFORCEPOOL: func() Any {
		return &iam.WorkforcePool{}
	},
	KMS_KEYRINGIAMMEMBER: func() Any {
		return &kms.KeyRingIAMMember{}
	},
	CLOUDRUNV2_SERVICEIAMBINDING: func() Any {
		return &cloudrunv2.ServiceIamBinding{}
	},
	CONTAINERANALYSIS_OCCURENCE: func() Any {
		return &containeranalysis.Occurence{}
	},
	PUBSUB_SUBSCRIPTIONIAMBINDING: func() Any {
		return &pubsub.SubscriptionIAMBinding{}
	},
	CERTIFICATEAUTHORITY_CAPOOL: func() Any {
		return &certificateauthority.CaPool{}
	},
	COMPUTE_TARGETTCPPROXY: func() Any {
		return &compute.TargetTCPProxy{}
	},
	NETWORKSECURITY_ADDRESSGROUPIAMPOLICY: func() Any {
		return &networksecurity.AddressGroupIamPolicy{}
	},
	DATACATALOG_TAGTEMPLATEIAMMEMBER: func() Any {
		return &datacatalog.TagTemplateIamMember{}
	},
	ML_ENGINEMODEL: func() Any {
		return &ml.EngineModel{}
	},
	COMPUTE_BACKENDBUCKETSIGNEDURLKEY: func() Any {
		return &compute.BackendBucketSignedUrlKey{}
	},
	DATACATALOG_POLICYTAG: func() Any {
		return &datacatalog.PolicyTag{}
	},
	IAM_ACCESSBOUNDARYPOLICY: func() Any {
		return &iam.AccessBoundaryPolicy{}
	},
	APIGEE_FLOWHOOK: func() Any {
		return &apigee.Flowhook{}
	},
	MONITORING_MONITOREDPROJECT: func() Any {
		return &monitoring.MonitoredProject{}
	},
	SECRETMANAGER_SECRETIAMMEMBER: func() Any {
		return &secretmanager.SecretIamMember{}
	},
	BIGQUERYANALYTICSHUB_DATAEXCHANGEIAMMEMBER: func() Any {
		return &bigqueryanalyticshub.DataExchangeIamMember{}
	},
	DNS_RESPONSEPOLICY: func() Any {
		return &dns.ResponsePolicy{}
	},
	OSLOGIN_SSHPUBLICKEY: func() Any {
		return &oslogin.SshPublicKey{}
	},
	APIGEE_ORGANIZATION: func() Any {
		return &apigee.Organization{}
	},
	COMPUTE_INSTANCETEMPLATE: func() Any {
		return &compute.InstanceTemplate{}
	},
	COMPUTE_NETWORKEDGESECURITYSERVICE: func() Any {
		return &compute.NetworkEdgeSecurityService{}
	},
	VERTEX_AIFEATUREGROUPFEATURE: func() Any {
		return &vertex.AiFeatureGroupFeature{}
	},
	VERTEX_AIFEATUREONLINESTOREFEATUREVIEW: func() Any {
		return &vertex.AiFeatureOnlineStoreFeatureview{}
	},
	CLOUDRUNV2_SERVICEIAMPOLICY: func() Any {
		return &cloudrunv2.ServiceIamPolicy{}
	},
	MONITORING_GROUP: func() Any {
		return &monitoring.Group{}
	},
	PUBSUB_LITETOPIC: func() Any {
		return &pubsub.LiteTopic{}
	},
	STORAGE_DEFAULTOBJECTACCESSCONTROL: func() Any {
		return &storage.DefaultObjectAccessControl{}
	},
	SECURITYCENTER_INSTANCEIAMBINDING: func() Any {
		return &securitycenter.InstanceIamBinding{}
	},
	COMPUTE_SNAPSHOTIAMBINDING: func() Any {
		return &compute.SnapshotIamBinding{}
	},
	FIRESTORE_BACKUPSCHEDULE: func() Any {
		return &firestore.BackupSchedule{}
	},
	STORAGE_DEFAULTOBJECTACL: func() Any {
		return &storage.DefaultObjectACL{}
	},
	BIGTABLE_TABLE: func() Any {
		return &bigtable.Table{}
	},
	CLOUDDEPLOY_DELIVERYPIPELINEIAMBINDING: func() Any {
		return &clouddeploy.DeliveryPipelineIamBinding{}
	},
	BLOCKCHAINNODEENGINE_BLOCKCHAINNODES: func() Any {
		return &blockchainnodeengine.BlockchainNodes{}
	},
	COMPUTE_PROJECTMETADATAITEM: func() Any {
		return &compute.ProjectMetadataItem{}
	},
	TAGS_TAGBINDING: func() Any {
		return &tags.TagBinding{}
	},
	WORKSTATIONS_WORKSTATIONCONFIGIAMMEMBER: func() Any {
		return &workstations.WorkstationConfigIamMember{}
	},
	HEALTHCARE_CONSENTSTOREIAMBINDING: func() Any {
		return &healthcare.ConsentStoreIamBinding{}
	},
	COMPUTE_MANAGEDSSLCERTIFICATE: func() Any {
		return &compute.ManagedSslCertificate{}
	},
	PROJECTS_ACCESSAPPROVALSETTINGS: func() Any {
		return &projects.AccessApprovalSettings{}
	},
	HEALTHCARE_DICOMSTOREIAMMEMBER: func() Any {
		return &healthcare.DicomStoreIamMember{}
	},
	STORAGE_BUCKETACL: func() Any {
		return &storage.BucketACL{}
	},
	COMPUTE_TARGETHTTPSPROXY: func() Any {
		return &compute.TargetHttpsProxy{}
	},
	DIAGFLOW_FULFILLMENT: func() Any {
		return &diagflow.Fulfillment{}
	},
	LOGGING_ORGANIZATIONEXCLUSION: func() Any {
		return &logging.OrganizationExclusion{}
	},
	PROJECTS_SERVICE: func() Any {
		return &projects.Service{}
	},
	DATAPLEX_LAKE: func() Any {
		return &dataplex.Lake{}
	},
	NETWORKSECURITY_FIREWALLENDPOINT: func() Any {
		return &networksecurity.FirewallEndpoint{}
	},
	ACCESSCONTEXTMANAGER_ACCESSLEVEL: func() Any {
		return &accesscontextmanager.AccessLevel{}
	},
	ACCESSCONTEXTMANAGER_SERVICEPERIMETER: func() Any {
		return &accesscontextmanager.ServicePerimeter{}
	},
	BEYONDCORP_APPCONNECTOR: func() Any {
		return &beyondcorp.AppConnector{}
	},
	RUNTIMECONFIG_CONFIGIAMMEMBER: func() Any {
		return &runtimeconfig.ConfigIamMember{}
	},
	COMPUTE_BACKENDSERVICEIAMBINDING: func() Any {
		return &compute.BackendServiceIamBinding{}
	},
	DATAFORM_REPOSITORYIAMBINDING: func() Any {
		return &dataform.RepositoryIamBinding{}
	},
	IAP_WEBREGIONBACKENDSERVICEIAMPOLICY: func() Any {
		return &iap.WebRegionBackendServiceIamPolicy{}
	},
	VMWAREENGINE_NETWORKPEERING: func() Any {
		return &vmwareengine.NetworkPeering{}
	},
	COMPUTE_FIREWALLPOLICYASSOCIATION: func() Any {
		return &compute.FirewallPolicyAssociation{}
	},
	IAP_WEBTYPECOMPUTEIAMBINDING: func() Any {
		return &iap.WebTypeComputeIamBinding{}
	},
	VERTEX_AIFEATURESTOREIAMBINDING: func() Any {
		return &vertex.AiFeatureStoreIamBinding{}
	},
	RESOURCEMANAGER_LIEN: func() Any {
		return &resourcemanager.Lien{}
	},
	BIGQUERYANALYTICSHUB_LISTINGIAMBINDING: func() Any {
		return &bigqueryanalyticshub.ListingIamBinding{}
	},
	COMPUTE_GLOBALNETWORKENDPOINTGROUP: func() Any {
		return &compute.GlobalNetworkEndpointGroup{}
	},
	BACKUPDISASTERRECOVERY_MANAGEMENTSERVER: func() Any {
		return &backupdisasterrecovery.ManagementServer{}
	},
	BIGTABLE_TABLEIAMPOLICY: func() Any {
		return &bigtable.TableIamPolicy{}
	},
	OSCONFIG_GUESTPOLICIES: func() Any {
		return &osconfig.GuestPolicies{}
	},
	BILLING_ACCOUNTIAMBINDING: func() Any {
		return &billing.AccountIamBinding{}
	},
	CERTIFICATEMANAGER_CERTIFICATEISSUANCECONFIG: func() Any {
		return &certificatemanager.CertificateIssuanceConfig{}
	},
	KMS_CRYPTOKEY: func() Any {
		return &kms.CryptoKey{}
	},
	NETWORKSECURITY_GATEWAYSECURITYPOLICYRULE: func() Any {
		return &networksecurity.GatewaySecurityPolicyRule{}
	},
	CLOUDRUN_SERVICE: func() Any {
		return &cloudrun.Service{}
	},
	COMPUTE_ORGANIZATIONSECURITYPOLICYASSOCIATION: func() Any {
		return &compute.OrganizationSecurityPolicyAssociation{}
	},
	COMPUTE_REGIONSSLCERTIFICATE: func() Any {
		return &compute.RegionSslCertificate{}
	},
	WORKFLOWS_WORKFLOW: func() Any {
		return &workflows.Workflow{}
	},
	BIGQUERY_IAMBINDING: func() Any {
		return &bigquery.IamBinding{}
	},
	COMPUTE_FIREWALLPOLICY: func() Any {
		return &compute.FirewallPolicy{}
	},
	IAP_WEBTYPECOMPUTEIAMMEMBER: func() Any {
		return &iap.WebTypeComputeIamMember{}
	},
	BIGQUERY_APPPROFILE: func() Any {
		return &bigquery.AppProfile{}
	},
	CONTAINER_AZURECLUSTER: func() Any {
		return &container.AzureCluster{}
	},
	DIAGFLOW_AGENT: func() Any {
		return &diagflow.Agent{}
	},
	SPANNER_INSTANCE: func() Any {
		return &spanner.Instance{}
	},
	GKEHUB_FEATURE: func() Any {
		return &gkehub.Feature{}
	},
	IAP_WEBREGIONBACKENDSERVICEIAMMEMBER: func() Any {
		return &iap.WebRegionBackendServiceIamMember{}
	},
	RUNTIMECONFIG_CONFIGIAMBINDING: func() Any {
		return &runtimeconfig.ConfigIamBinding{}
	},
	BIGQUERY_CAPACITYCOMMITMENT: func() Any {
		return &bigquery.CapacityCommitment{}
	},
	COMPUTE_INSTANCEGROUP: func() Any {
		return &compute.InstanceGroup{}
	},
	DNS_DNSMANAGEDZONEIAMBINDING: func() Any {
		return &dns.DnsManagedZoneIamBinding{}
	},
	IAP_CLIENT: func() Any {
		return &iap.Client{}
	},
	ACCESSCONTEXTMANAGER_AUTHORIZEDORGSDESC: func() Any {
		return &accesscontextmanager.AuthorizedOrgsDesc{}
	},
	APIGEE_KEYSTORESALIASESPKCS12: func() Any {
		return &apigee.KeystoresAliasesPkcs12{}
	},
	BIGQUERYANALYTICSHUB_LISTING: func() Any {
		return &bigqueryanalyticshub.Listing{}
	},
	COMPUTE_REGIONHEALTHCHECK: func() Any {
		return &compute.RegionHealthCheck{}
	},
	DATAPROC_JOBIAMMEMBER: func() Any {
		return &dataproc.JobIAMMember{}
	},
	FIRESTORE_FIELD: func() Any {
		return &firestore.Field{}
	},
	COMPUTE_REGIONTARGETHTTPPROXY: func() Any {
		return &compute.RegionTargetHttpProxy{}
	},
	IAP_TUNNELIAMMEMBER: func() Any {
		return &iap.TunnelIamMember{}
	},
	KMS_CRYPTOKEYIAMBINDING: func() Any {
		return &kms.CryptoKeyIAMBinding{}
	},
	PUBSUB_SCHEMAIAMBINDING: func() Any {
		return &pubsub.SchemaIamBinding{}
	},
	COMPUTE_REGIONINSTANCEGROUPMANAGER: func() Any {
		return &compute.RegionInstanceGroupManager{}
	},
	DATASTREAM_CONNECTIONPROFILE: func() Any {
		return &datastream.ConnectionProfile{}
	},
	FIRESTORE_DOCUMENT: func() Any {
		return &firestore.Document{}
	},
	CLOUDSCHEDULER_JOB: func() Any {
		return &cloudscheduler.Job{}
	},
	COMPUTE_BACKENDBUCKETIAMPOLICY: func() Any {
		return &compute.BackendBucketIamPolicy{}
	},
	FIREBASE_ANDROIDAPP: func() Any {
		return &firebase.AndroidApp{}
	},
	IAP_WEBTYPEAPPENGINGIAMBINDING: func() Any {
		return &iap.WebTypeAppEngingIamBinding{}
	},
	DATAFORM_REPOSITORY: func() Any {
		return &dataform.Repository{}
	},
	GKEONPREM_BAREMETALCLUSTER: func() Any {
		return &gkeonprem.BareMetalCluster{}
	},
	NETWORKSECURITY_GATEWAYSECURITYPOLICY: func() Any {
		return &networksecurity.GatewaySecurityPolicy{}
	},
	CONTAINER_CLUSTER: func() Any {
		return &container.Cluster{}
	},
	DATAFORM_REPOSITORYIAMPOLICY: func() Any {
		return &dataform.RepositoryIamPolicy{}
	},
	DATAPROC_AUTOSCALINGPOLICYIAMBINDING: func() Any {
		return &dataproc.AutoscalingPolicyIamBinding{}
	},
	ESSENTIALCONTACTS_DOCUMENTAIWAREHOUSELOCATION: func() Any {
		return &essentialcontacts.DocumentAiWarehouseLocation{}
	},
	CERTIFICATEAUTHORITY_CERTIFICATETEMPLATEIAMPOLICY: func() Any {
		return &certificateauthority.CertificateTemplateIamPolicy{}
	},
	SOURCEREPO_REPOSITORYIAMBINDING: func() Any {
		return &sourcerepo.RepositoryIamBinding{}
	},
	RUNTIMECONFIG_CONFIGIAMPOLICY: func() Any {
		return &runtimeconfig.ConfigIamPolicy{}
	},
	APIGEE_SHAREDFLOW: func() Any {
		return &apigee.Sharedflow{}
	},
	BIGQUERYANALYTICSHUB_DATAEXCHANGEIAMBINDING: func() Any {
		return &bigqueryanalyticshub.DataExchangeIamBinding{}
	},
	BINARYAUTHORIZATION_ATTESTORIAMMEMBER: func() Any {
		return &binaryauthorization.AttestorIamMember{}
	},
	VMWAREENGINE_NETWORK: func() Any {
		return &vmwareengine.Network{}
	},
	FIREBASE_APPCHECKDEBUGTOKEN: func() Any {
		return &firebase.AppCheckDebugToken{}
	},
	STORAGE_HMACKEY: func() Any {
		return &storage.HmacKey{}
	},
	ACCESSCONTEXTMANAGER_ACCESSPOLICYIAMMEMBER: func() Any {
		return &accesscontextmanager.AccessPolicyIamMember{}
	},
	DATACATALOG_TAGTEMPLATEIAMPOLICY: func() Any {
		return &datacatalog.TagTemplateIamPolicy{}
	},
	ENDPOINTS_CONSUMERSIAMBINDING: func() Any {
		return &endpoints.ConsumersIamBinding{}
	},
	TAGS_TAGVALUEIAMPOLICY: func() Any {
		return &tags.TagValueIamPolicy{}
	},
	HEALTHCARE_HL7STORE: func() Any {
		return &healthcare.Hl7Store{}
	},
	COMPUTE_INTERCONNECTATTACHMENT: func() Any {
		return &compute.InterconnectAttachment{}
	},
	COMPUTE_REGIONNETWORKFIREWALLPOLICY: func() Any {
		return &compute.RegionNetworkFirewallPolicy{}
	},
	COMPUTE_TARGETGRPCPROXY: func() Any {
		return &compute.TargetGrpcProxy{}
	},
	NETWORKSERVICES_TCPROUTE: func() Any {
		return &networkservices.TcpRoute{}
	},
	IDENTITYPLATFORM_CONFIG: func() Any {
		return &identityplatform.Config{}
	},
	SECURITYCENTER_SOURCEIAMBINDING: func() Any {
		return &securitycenter.SourceIamBinding{}
	},
	CONTAINER_AZURECLIENT: func() Any {
		return &container.AzureClient{}
	},
	FOLDER_IAMBINDING: func() Any {
		return &folder.IAMBinding{}
	},
	GKEBACKUP_BACKUPPLANIAMBINDING: func() Any {
		return &gkebackup.BackupPlanIamBinding{}
	},
	WORKSTATIONS_WORKSTATIONIAMBINDING: func() Any {
		return &workstations.WorkstationIamBinding{}
	},
	EDGENETWORK_NETWORK: func() Any {
		return &edgenetwork.Network{}
	},
	GKEHUB_SCOPEIAMBINDING: func() Any {
		return &gkehub.ScopeIamBinding{}
	},
	IAP_WEBBACKENDSERVICEIAMPOLICY: func() Any {
		return &iap.WebBackendServiceIamPolicy{}
	},
	WORKSTATIONS_WORKSTATIONCLUSTER: func() Any {
		return &workstations.WorkstationCluster{}
	},
	COMPUTE_ROUTERPEER: func() Any {
		return &compute.RouterPeer{}
	},
	COMPUTE_SSLPOLICY: func() Any {
		return &compute.SSLPolicy{}
	},
	NETWORKCONNECTIVITY_POLICYBASEDROUTE: func() Any {
		return &networkconnectivity.PolicyBasedRoute{}
	},
	ARTIFACTREGISTRY_REPOSITORYIAMPOLICY: func() Any {
		return &artifactregistry.RepositoryIamPolicy{}
	},
	COMPUTE_SNAPSHOTIAMMEMBER: func() Any {
		return &compute.SnapshotIamMember{}
	},
	CONTAINER_AZURENODEPOOL: func() Any {
		return &container.AzureNodePool{}
	},
	FIRESTORE_INDEX: func() Any {
		return &firestore.Index{}
	},
	CERTIFICATEAUTHORITY_CAPOOLIAMPOLICY: func() Any {
		return &certificateauthority.CaPoolIamPolicy{}
	},
	COMPUTE_TARGETSSLPROXY: func() Any {
		return &compute.TargetSSLProxy{}
	},
	GKEHUB_MEMBERSHIPRBACROLEBINDING: func() Any {
		return &gkehub.MembershipRbacRoleBinding{}
	},
	CONTAINER_REGISTRY: func() Any {
		return &container.Registry{}
	},
	DATACATALOG_POLICYTAGIAMMEMBER: func() Any {
		return &datacatalog.PolicyTagIamMember{}
	},
	SECRETMANAGER_SECRETVERSION: func() Any {
		return &secretmanager.SecretVersion{}
	},
	BILLING_PROJECTINFO: func() Any {
		return &billing.ProjectInfo{}
	},
	COMPUTE_DISKIAMBINDING: func() Any {
		return &compute.DiskIamBinding{}
	},
	DATAPROC_METASTOREFEDERATION: func() Any {
		return &dataproc.MetastoreFederation{}
	},
	NETWORKSERVICES_ENDPOINTPOLICY: func() Any {
		return &networkservices.EndpointPolicy{}
	},
	DATAPROC_METASTORESERVICEIAMPOLICY: func() Any {
		return &dataproc.MetastoreServiceIamPolicy{}
	},
	ORGPOLICY_POLICY: func() Any {
		return &orgpolicy.Policy{}
	},
	ACCESSCONTEXTMANAGER_GCPUSERACCESSBINDING: func() Any {
		return &accesscontextmanager.GcpUserAccessBinding{}
	},
	LOGGING_FOLDERSINK: func() Any {
		return &logging.FolderSink{}
	},
	NETWORKMANAGEMENT_CONNECTIVITYTEST: func() Any {
		return &networkmanagement.ConnectivityTest{}
	},
	BILLING_ACCOUNTIAMPOLICY: func() Any {
		return &billing.AccountIamPolicy{}
	},
	VERTEX_AIFEATURESTOREENTITYTYPEIAMPOLICY: func() Any {
		return &vertex.AiFeatureStoreEntityTypeIamPolicy{}
	},
	COMPUTE_REGIONNETWORKENDPOINT: func() Any {
		return &compute.RegionNetworkEndpoint{}
	},
	DATAPLEX_ASSETIAMMEMBER: func() Any {
		return &dataplex.AssetIamMember{}
	},
	NETWORKSECURITY_URLLIST: func() Any {
		return &networksecurity.UrlList{}
	},
	FOLDER_ORGANIZATIONPOLICY: func() Any {
		return &folder.OrganizationPolicy{}
	},
	WORKBENCH_INSTANCE: func() Any {
		return &workbench.Instance{}
	},
	GKEBACKUP_BACKUPPLANIAMMEMBER: func() Any {
		return &gkebackup.BackupPlanIamMember{}
	},
	HEALTHCARE_HL7STOREIAMMEMBER: func() Any {
		return &healthcare.Hl7StoreIamMember{}
	},
	APIGEE_ENVIRONMENTIAMBINDING: func() Any {
		return &apigee.EnvironmentIamBinding{}
	},
	CERTIFICATEAUTHORITY_AUTHORITY: func() Any {
		return &certificateauthority.Authority{}
	},
	DATALOSS_PREVENTIONJOBTRIGGER: func() Any {
		return &dataloss.PreventionJobTrigger{}
	},
	DATALOSS_PREVENTIONSTOREDINFOTYPE: func() Any {
		return &dataloss.PreventionStoredInfoType{}
	},
	CLOUDDEPLOY_AUTOMATION: func() Any {
		return &clouddeploy.Automation{}
	},
	FILESTORE_BACKUP: func() Any {
		return &filestore.Backup{}
	},
	IDENTITYPLATFORM_TENANTOAUTHIDPCONFIG: func() Any {
		return &identityplatform.TenantOauthIdpConfig{}
	},
	VERTEX_AIFEATURESTOREENTITYTYPEIAMBINDING: func() Any {
		return &vertex.AiFeatureStoreEntityTypeIamBinding{}
	},
	COMPUTE_MANGEDSSLCERTIFICATE: func() Any {
		return &compute.MangedSslCertificate{}
	},
	VERTEX_AIFEATURESTOREIAMMEMBER: func() Any {
		return &vertex.AiFeatureStoreIamMember{}
	},
	APIGATEWAY_APIIAMPOLICY: func() Any {
		return &apigateway.ApiIamPolicy{}
	},
	ESSENTIALCONTACTS_DOCUMENTAIPROCESSOR: func() Any {
		return &essentialcontacts.DocumentAiProcessor{}
	},
	BIGQUERYDATAPOLICY_DATAPOLICYIAMMEMBER: func() Any {
		return &bigquerydatapolicy.DataPolicyIamMember{}
	},
	CERTIFICATEAUTHORITY_CERTIFICATETEMPLATEIAMBINDING: func() Any {
		return &certificateauthority.CertificateTemplateIamBinding{}
	},
	GKEHUB_MEMBERSHIPIAMPOLICY: func() Any {
		return &gkehub.MembershipIamPolicy{}
	},
	MIGRATIONCENTER_GROUP: func() Any {
		return &migrationcenter.Group{}
	},
	TAGS_TAGVALUEIAMMEMBER: func() Any {
		return &tags.TagValueIamMember{}
	},
	ACCESSCONTEXTMANAGER_ACCESSPOLICY: func() Any {
		return &accesscontextmanager.AccessPolicy{}
	},
	APIGEE_ENDPOINTATTACHMENT: func() Any {
		return &apigee.EndpointAttachment{}
	},
	COMPUTE_IMAGEIAMMEMBER: func() Any {
		return &compute.ImageIamMember{}
	},
	SERVICEACCOUNT_ACCOUNT: func() Any {
		return &serviceaccount.Account{}
	},
	CLOUDRUN_IAMPOLICY: func() Any {
		return &cloudrun.IamPolicy{}
	},
	FIREBASE_WEBAPP: func() Any {
		return &firebase.WebApp{}
	},
	SECURITYCENTER_SOURCE: func() Any {
		return &securitycenter.Source{}
	},
	SPANNER_DATABASE: func() Any {
		return &spanner.Database{}
	},
	NETWORKSERVICES_SERVICEBINDING: func() Any {
		return &networkservices.ServiceBinding{}
	},
	BINARYAUTHORIZATION_ATTESTORIAMPOLICY: func() Any {
		return &binaryauthorization.AttestorIamPolicy{}
	},
	COMPUTE_CAEXTERNALACCOUNTKEY: func() Any {
		return &compute.CaExternalAccountKey{}
	},
	DATACATALOG_TAXONOMYIAMPOLICY: func() Any {
		return &datacatalog.TaxonomyIamPolicy{}
	},
	FIREBASE_PROJECT: func() Any {
		return &firebase.Project{}
	},
	APIGATEWAY_APIIAMMEMBER: func() Any {
		return &apigateway.ApiIamMember{}
	},
	BILLING_ACCOUNTIAMMEMBER: func() Any {
		return &billing.AccountIamMember{}
	},
	BILLING_BUDGET: func() Any {
		return &billing.Budget{}
	},
	PROJECTS_SERVICEIDENTITY: func() Any {
		return &projects.ServiceIdentity{}
	},
	IDENTITYPLATFORM_TENANTINBOUNDSAMLCONFIG: func() Any {
		return &identityplatform.TenantInboundSamlConfig{}
	},
	STORAGE_BUCKETACCESSCONTROL: func() Any {
		return &storage.BucketAccessControl{}
	},
	TAGS_TAGKEYIAMMEMBER: func() Any {
		return &tags.TagKeyIamMember{}
	},
	DATAPLEX_LAKEIAMMEMBER: func() Any {
		return &dataplex.LakeIamMember{}
	},
	WORKBENCH_INSTANCEIAMBINDING: func() Any {
		return &workbench.InstanceIamBinding{}
	},
	ACTIVEDIRECTORY_DOMAIN: func() Any {
		return &activedirectory.Domain{}
	},
	DNS_RECORDSET: func() Any {
		return &dns.RecordSet{}
	},
	COMPUTE_REGIONCOMMITMENT: func() Any {
		return &compute.RegionCommitment{}
	},
	BINARYAUTHORIZATION_ATTESTORIAMBINDING: func() Any {
		return &binaryauthorization.AttestorIamBinding{}
	},
	DATACATALOG_TAXONOMY: func() Any {
		return &datacatalog.Taxonomy{}
	},
	SECRETMANAGER_SECRETIAMBINDING: func() Any {
		return &secretmanager.SecretIamBinding{}
	},
	COMPUTE_NETWORKPEERINGROUTESCONFIG: func() Any {
		return &compute.NetworkPeeringRoutesConfig{}
	},
	SPANNER_DATABASEIAMBINDING: func() Any {
		return &spanner.DatabaseIAMBinding{}
	},
	APIGEE_SHAREDFLOWDEPLOYMENT: func() Any {
		return &apigee.SharedflowDeployment{}
	},
	IAM_WORKLOADIDENTITYPOOLPROVIDER: func() Any {
		return &iam.WorkloadIdentityPoolProvider{}
	},
	DATACATALOG_POLICYTAGIAMPOLICY: func() Any {
		return &datacatalog.PolicyTagIamPolicy{}
	},
	DATASTREAM_STREAM: func() Any {
		return &datastream.Stream{}
	},
	HEALTHCARE_CONSENTSTORE: func() Any {
		return &healthcare.ConsentStore{}
	},
	SECURESOURCEMANAGER_INSTANCEIAMPOLICY: func() Any {
		return &securesourcemanager.InstanceIamPolicy{}
	},
	COMPUTE_NETWORKFIREWALLPOLICYRULE: func() Any {
		return &compute.NetworkFirewallPolicyRule{}
	},
	PUBSUB_SUBSCRIPTION: func() Any {
		return &pubsub.Subscription{}
	},
	VERTEX_AIFEATUREGROUP: func() Any {
		return &vertex.AiFeatureGroup{}
	},
	DATAPLEX_ZONEIAMMEMBER: func() Any {
		return &dataplex.ZoneIamMember{}
	},
	PUBSUB_LITESUBSCRIPTION: func() Any {
		return &pubsub.LiteSubscription{}
	},
	PUBSUB_TOPICIAMPOLICY: func() Any {
		return &pubsub.TopicIAMPolicy{}
	},
	WORKSTATIONS_WORKSTATIONCONFIGIAMPOLICY: func() Any {
		return &workstations.WorkstationConfigIamPolicy{}
	},
	APIGEE_ENVREFERENCES: func() Any {
		return &apigee.EnvReferences{}
	},
	COMPUTE_ROUTER: func() Any {
		return &compute.Router{}
	},
	GKEONPREM_BAREMETALADMINCLUSTER: func() Any {
		return &gkeonprem.BareMetalAdminCluster{}
	},
	VERTEX_AIINDEXENDPOINT: func() Any {
		return &vertex.AiIndexEndpoint{}
	},
	NOTEBOOKS_LOCATION: func() Any {
		return &notebooks.Location{}
	},
	VERTEX_AIFEATURESTOREIAMPOLICY: func() Any {
		return &vertex.AiFeatureStoreIamPolicy{}
	},
	APIGATEWAY_GATEWAYIAMBINDING: func() Any {
		return &apigateway.GatewayIamBinding{}
	},
	BIGLAKE_TABLE: func() Any {
		return &biglake.Table{}
	},
	DATAPROC_JOBIAMBINDING: func() Any {
		return &dataproc.JobIAMBinding{}
	},
	FIREBASE_EXTENSIONSINSTANCE: func() Any {
		return &firebase.ExtensionsInstance{}
	},
	COMPUTE_HTTPHEALTHCHECK: func() Any {
		return &compute.HttpHealthCheck{}
	},
	DATAPROC_CLUSTERIAMMEMBER: func() Any {
		return &dataproc.ClusterIAMMember{}
	},
	GKEONPREM_BAREMETALNODEPOOL: func() Any {
		return &gkeonprem.BareMetalNodePool{}
	},
	NETAPP_VOLUMESNAPSHOT: func() Any {
		return &netapp.VolumeSnapshot{}
	},
	DATACATALOG_TAG: func() Any {
		return &datacatalog.Tag{}
	},
	KMS_KEYRINGIAMPOLICY: func() Any {
		return &kms.KeyRingIAMPolicy{}
	},
	CERTIFICATEMANAGER_DNSAUTHORIZATION: func() Any {
		return &certificatemanager.DnsAuthorization{}
	},
	CLOUDBUILDV2_CONNECTIONIAMMEMBER: func() Any {
		return &cloudbuildv2.ConnectionIAMMember{}
	},
	CLOUDDEPLOY_DELIVERYPIPELINE: func() Any {
		return &clouddeploy.DeliveryPipeline{}
	},
	COMPUTE_IMAGEIAMBINDING: func() Any {
		return &compute.ImageIamBinding{}
	},
	CLOUDTASKS_QUEUEIAMBINDING: func() Any {
		return &cloudtasks.QueueIamBinding{}
	},
	DATACATALOG_TAXONOMYIAMMEMBER: func() Any {
		return &datacatalog.TaxonomyIamMember{}
	},
	DATASTREAM_PRIVATECONNECTION: func() Any {
		return &datastream.PrivateConnection{}
	},
	DIAGFLOW_INTENT: func() Any {
		return &diagflow.Intent{}
	},
	GKEHUB_FEATUREIAMMEMBER: func() Any {
		return &gkehub.FeatureIamMember{}
	},
	HEALTHCARE_HL7STOREIAMPOLICY: func() Any {
		return &healthcare.Hl7StoreIamPolicy{}
	},
	PROJECTS_APIKEY: func() Any {
		return &projects.ApiKey{}
	},
	CLOUDRUN_IAMMEMBER: func() Any {
		return &cloudrun.IamMember{}
	},
	DATAPLEX_LAKEIAMBINDING: func() Any {
		return &dataplex.LakeIamBinding{}
	},
	FIREBASE_APPCHECKRECAPTCHAENTERPRISECONFIG: func() Any {
		return &firebase.AppCheckRecaptchaEnterpriseConfig{}
	},
	FIREBASE_DATABASEINSTANCE: func() Any {
		return &firebase.DatabaseInstance{}
	},
	SECURESOURCEMANAGER_INSTANCEIAMMEMBER: func() Any {
		return &securesourcemanager.InstanceIamMember{}
	},
	STORAGE_BUCKETIAMPOLICY: func() Any {
		return &storage.BucketIAMPolicy{}
	},
	VERTEX_AIFEATURESTOREENTITYTYPEFEATURE: func() Any {
		return &vertex.AiFeatureStoreEntityTypeFeature{}
	},
	SOURCEREPO_REPOSITORYIAMPOLICY: func() Any {
		return &sourcerepo.RepositoryIamPolicy{}
	},
	ARTIFACTREGISTRY_REPOSITORY: func() Any {
		return &artifactregistry.Repository{}
	},
	COMPUTE_ORGANIZATIONSECURITYPOLICYRULE: func() Any {
		return &compute.OrganizationSecurityPolicyRule{}
	},
	ENDPOINTS_CONSUMERSIAMPOLICY: func() Any {
		return &endpoints.ConsumersIamPolicy{}
	},
	ORGANIZATIONS_POLICY: func() Any {
		return &aws_organizations.Policy{}
	},
	SQL_DATABASE: func() Any {
		return &sql.Database{}
	},
	STORAGE_OBJECTACL: func() Any {
		return &storage.ObjectACL{}
	},
	CONTAINER_AWSCLUSTER: func() Any {
		return &container.AwsCluster{}
	},
	DATAFUSION_INSTANCE: func() Any {
		return &datafusion.Instance{}
	},
	DATAPROC_CLUSTERIAMPOLICY: func() Any {
		return &dataproc.ClusterIAMPolicy{}
	},
	LOGGING_PROJECTEXCLUSION: func() Any {
		return &logging.ProjectExclusion{}
	},
	COMPUTE_ADDRESS: func() Any {
		return &compute.Address{}
	},
	SPANNER_INSTANCEIAMPOLICY: func() Any {
		return &spanner.InstanceIAMPolicy{}
	},
	DATAPLEX_ZONEIAMBINDING: func() Any {
		return &dataplex.ZoneIamBinding{}
	},
	KMS_KEYRING: func() Any {
		return &kms.KeyRing{}
	},
	DATAPLEX_TASKIAMBINDING: func() Any {
		return &dataplex.TaskIamBinding{}
	},
	DATAPROC_METASTORESERVICEIAMBINDING: func() Any {
		return &dataproc.MetastoreServiceIamBinding{}
	},
	IAP_WEBIAMMEMBER: func() Any {
		return &iap.WebIamMember{}
	},
	COMPUTE_REGIONAUTOSCALER: func() Any {
		return &compute.RegionAutoscaler{}
	},
	DATAPROC_METASTOREFEDERATIONIAMPOLICY: func() Any {
		return &dataproc.MetastoreFederationIamPolicy{}
	},
	FIREBASERULES_RULESET: func() Any {
		return &firebaserules.Ruleset{}
	},
	CLOUDFUNCTIONS_FUNCTIONIAMBINDING: func() Any {
		return &cloudfunctions.FunctionIamBinding{}
	},
	COMPUTE_REGIONURLMAP: func() Any {
		return &compute.RegionUrlMap{}
	},
	SECURITYCENTER_ORGANIZATIONCUSTOMMODULE: func() Any {
		return &securitycenter.OrganizationCustomModule{}
	},
	SECURITYPOSTURE_POSTURE: func() Any {
		return &securityposture.Posture{}
	},
	ESSENTIALCONTACTS_DOCUMENTAIWAREHOUSEDOCUMENTSCHEMA: func() Any {
		return &essentialcontacts.DocumentAiWarehouseDocumentSchema{}
	},
	MIGRATIONCENTER_PREFERENCESET: func() Any {
		return &migrationcenter.PreferenceSet{}
	},
	SERVICEACCOUNT_KEY: func() Any {
		return &serviceaccount.Key{}
	},
	APIGEE_ENVIRONMENT: func() Any {
		return &apigee.Environment{}
	},
	CLOUDIDENTITY_GROUPMEMBERSHIP: func() Any {
		return &cloudidentity.GroupMembership{}
	},
	DATACATALOG_ENTRYGROUPIAMMEMBER: func() Any {
		return &datacatalog.EntryGroupIamMember{}
	},
	DATAPROC_JOBIAMPOLICY: func() Any {
		return &dataproc.JobIAMPolicy{}
	},
	BIGQUERY_DATASETIAMBINDING: func() Any {
		return &bigquery.DatasetIamBinding{}
	},
	COMPUTE_MACHINEIMAGEIAMPOLICY: func() Any {
		return &compute.MachineImageIamPolicy{}
	},
	COMPUTE_TARGETINSTANCE: func() Any {
		return &compute.TargetInstance{}
	},
	BIGQUERYANALYTICSHUB_LISTINGIAMMEMBER: func() Any {
		return &bigqueryanalyticshub.ListingIamMember{}
	},
	COMPUTE_REGIONINSTANCETEMPLATE: func() Any {
		return &compute.RegionInstanceTemplate{}
	},
	EVENTARC_TRIGGER: func() Any {
		return &eventarc.Trigger{}
	},
	PUBSUB_TOPICIAMBINDING: func() Any {
		return &pubsub.TopicIAMBinding{}
	},
	SECRETMANAGER_SECRETIAMPOLICY: func() Any {
		return &secretmanager.SecretIamPolicy{}
	},
	STORAGE_NOTIFICATION: func() Any {
		return &storage.Notification{}
	},
	COMPUTE_GLOBALNETWORKENDPOINT: func() Any {
		return &compute.GlobalNetworkEndpoint{}
	},
	COMPUTE_ROUTERINTERFACE: func() Any {
		return &compute.RouterInterface{}
	},
	MONITORING_SLO: func() Any {
		return &monitoring.Slo{}
	},
	NETWORKSECURITY_ADDRESSGROUP: func() Any {
		return &networksecurity.AddressGroup{}
	},
	PUBSUB_SCHEMAIAMMEMBER: func() Any {
		return &pubsub.SchemaIamMember{}
	},
	SECURITYCENTER_NOTIFICATIONCONFIG: func() Any {
		return &securitycenter.NotificationConfig{}
	},
	TAGS_TAGKEYIAMPOLICY: func() Any {
		return &tags.TagKeyIamPolicy{}
	},
	BIGQUERY_CONNECTIONIAMBINDING: func() Any {
		return &bigquery.ConnectionIamBinding{}
	},
	BIGQUERY_CONNECTIONIAMMEMBER: func() Any {
		return &bigquery.ConnectionIamMember{}
	},
	DIAGFLOW_CXSECURITYSETTINGS: func() Any {
		return &diagflow.CxSecuritySettings{}
	},
	EDGECONTAINER_CLUSTER: func() Any {
		return &edgecontainer.Cluster{}
	},
	PUBSUB_SUBSCRIPTIONIAMPOLICY: func() Any {
		return &pubsub.SubscriptionIAMPolicy{}
	},
	SQL_SSLCERT: func() Any {
		return &sql.SslCert{}
	},
	DIAGFLOW_CXINTENT: func() Any {
		return &diagflow.CxIntent{}
	},
	GKEHUB_MEMBERSHIPIAMMEMBER: func() Any {
		return &gkehub.MembershipIamMember{}
	},
	IAP_TUNNELIAMPOLICY: func() Any {
		return &iap.TunnelIamPolicy{}
	},
	CLOUDBUILDV2_REPOSITORY: func() Any {
		return &cloudbuildv2.Repository{}
	},
	COMPUTE_SUBNETWORKIAMMEMBER: func() Any {
		return &compute.SubnetworkIAMMember{}
	},
	NETAPP_STORAGEPOOL: func() Any {
		return &netapp.StoragePool{}
	},
	VMWAREENGINE_NETWORKPOLICY: func() Any {
		return &vmwareengine.NetworkPolicy{}
	},
	ACCESSCONTEXTMANAGER_SERVICEPERIMETERINGRESSPOLICY: func() Any {
		return &accesscontextmanager.ServicePerimeterIngressPolicy{}
	},
	BIGQUERY_RESERVATIONASSIGNMENT: func() Any {
		return &bigquery.ReservationAssignment{}
	},
	DATAPLEX_ASSETIAMPOLICY: func() Any {
		return &dataplex.AssetIamPolicy{}
	},
	DISCOVERYENGINE_SEARCHENGINE: func() Any {
		return &discoveryengine.SearchEngine{}
	},
	VMWAREENGINE_EXTERNALACCESSRULE: func() Any {
		return &vmwareengine.ExternalAccessRule{}
	},
	COMPUTE_REGIONSECURITYPOLICYRULE: func() Any {
		return &compute.RegionSecurityPolicyRule{}
	},
	COMPUTE_REGIONTARGETTCPPROXY: func() Any {
		return &compute.RegionTargetTcpProxy{}
	},
	GKEBACKUP_RESTOREPLANIAMBINDING: func() Any {
		return &gkebackup.RestorePlanIamBinding{}
	},
	KMS_CRYPTOKEYIAMMEMBER: func() Any {
		return &kms.CryptoKeyIAMMember{}
	},
	COMPUTE_MACHINEIMAGEIAMBINDING: func() Any {
		return &compute.MachineImageIamBinding{}
	},
	NOTEBOOKS_ENVIRONMENT: func() Any {
		return &notebooks.Environment{}
	},
	COMPUTE_REGIONBACKENDSERVICEIAMPOLICY: func() Any {
		return &compute.RegionBackendServiceIamPolicy{}
	},
	DATAFORM_REPOSITORYRELEASECONFIG: func() Any {
		return &dataform.RepositoryReleaseConfig{}
	},
	DISCOVERYENGINE_CHATENGINE: func() Any {
		return &discoveryengine.ChatEngine{}
	},
	LOGGING_LINKEDDATASET: func() Any {
		return &logging.LinkedDataset{}
	},
	NOTEBOOKS_INSTANCE: func() Any {
		return &notebooks.Instance{}
	},
	SERVICENETWORKING_PEEREDDNSDOMAIN: func() Any {
		return &servicenetworking.PeeredDnsDomain{}
	},
	PROJECTS_IAMMEMBER: func() Any {
		return &projects.IAMMember{}
	},
	VERTEX_AIDATASET: func() Any {
		return &vertex.AiDataset{}
	},
	DATAFLOW_JOB: func() Any {
		return &dataflow.Job{}
	},
	EVENTARC_GOOGLECHANNELCONFIG: func() Any {
		return &eventarc.GoogleChannelConfig{}
	},
	IAM_WORKLOADIDENTITYPOOL: func() Any {
		return &iam.WorkloadIdentityPool{}
	},
	ACCESSCONTEXTMANAGER_SERVICEPERIMETERS: func() Any {
		return &accesscontextmanager.ServicePerimeters{}
	},
	HEALTHCARE_FHIRSTOREIAMBINDING: func() Any {
		return &healthcare.FhirStoreIamBinding{}
	},
	HEALTHCARE_DATASETIAMMEMBER: func() Any {
		return &healthcare.DatasetIamMember{}
	},
	SERVICEACCOUNT_IAMMEMBER: func() Any {
		return &serviceaccount.IAMMember{}
	},
	SERVICEACCOUNT_IAMPOLICY: func() Any {
		return &serviceaccount.IAMPolicy{}
	},
	SOURCEREPO_REPOSITORYIAMMEMBER: func() Any {
		return &sourcerepo.RepositoryIamMember{}
	},
	APPENGINE_SERVICENETWORKSETTINGS: func() Any {
		return &appengine.ServiceNetworkSettings{}
	},
	COMPUTE_NETWORKFIREWALLPOLICY: func() Any {
		return &compute.NetworkFirewallPolicy{}
	},
	COMPUTE_SSLCERTIFICATE: func() Any {
		return &compute.SSLCertificate{}
	},
	ENDPOINTS_SERVICE: func() Any {
		return &endpoints.Service{}
	},
	CLOUDFUNCTIONSV2_FUNCTIONIAMMEMBER: func() Any {
		return &cloudfunctionsv2.FunctionIamMember{}
	},
	GKEHUB_SCOPEIAMMEMBER: func() Any {
		return &gkehub.ScopeIamMember{}
	},
	STORAGE_INSIGHTSREPORTCONFIG: func() Any {
		return &storage.InsightsReportConfig{}
	},
	LOGGING_BILLINGACCOUNTBUCKETCONFIG: func() Any {
		return &logging.BillingAccountBucketConfig{}
	},
	PUBSUB_TOPIC: func() Any {
		return &pubsub.Topic{}
	},
	SPANNER_INSTANCEIAMMEMBER: func() Any {
		return &spanner.InstanceIAMMember{}
	},
	ACCESSCONTEXTMANAGER_ACCESSPOLICYIAMBINDING: func() Any {
		return &accesscontextmanager.AccessPolicyIamBinding{}
	},
	APPENGINE_APPLICATIONURLDISPATCHRULES: func() Any {
		return &appengine.ApplicationUrlDispatchRules{}
	},
	DATAPLEX_ZONEIAMPOLICY: func() Any {
		return &dataplex.ZoneIamPolicy{}
	},
	ESSENTIALCONTACTS_CONTACT: func() Any {
		return &essentialcontacts.Contact{}
	},
	IAP_WEBREGIONBACKENDSERVICEIAMBINDING: func() Any {
		return &iap.WebRegionBackendServiceIamBinding{}
	},
	IDENTITYPLATFORM_TENANT: func() Any {
		return &identityplatform.Tenant{}
	},
	COMPUTE_ATTACHEDDISK: func() Any {
		return &compute.AttachedDisk{}
	},
	DATACATALOG_ENTRYGROUP: func() Any {
		return &datacatalog.EntryGroup{}
	},
	HEALTHCARE_DATASET: func() Any {
		return &healthcare.Dataset{}
	},
	COMPUTE_PACKETMIRRORING: func() Any {
		return &compute.PacketMirroring{}
	},
	EDGENETWORK_SUBNET: func() Any {
		return &edgenetwork.Subnet{}
	},
	BIGTABLE_GCPOLICY: func() Any {
		return &bigtable.GCPolicy{}
	},
	COMPOSER_ENVIRONMENT: func() Any {
		return &composer.Environment{}
	},
	CERTIFICATEAUTHORITY_CERTIFICATE: func() Any {
		return &certificateauthority.Certificate{}
	},
	COMPUTE_REGIONDISKIAMMEMBER: func() Any {
		return &compute.RegionDiskIamMember{}
	},
	SERVICENETWORKING_CONNECTION: func() Any {
		return &servicenetworking.Connection{}
	},
	APPENGINE_APPLICATION: func() Any {
		return &appengine.Application{}
	},
	COMPUTE_INSTANCEIAMMEMBER: func() Any {
		return &compute.InstanceIAMMember{}
	},
	VERTEX_AIINDEX: func() Any {
		return &vertex.AiIndex{}
	},
	FIREBASERULES_RELEASE: func() Any {
		return &firebaserules.Release{}
	},
	ACCESSCONTEXTMANAGER_ACCESSLEVELS: func() Any {
		return &accesscontextmanager.AccessLevels{}
	},
	CONTAINER_ATTACHEDCLUSTER: func() Any {
		return &container.AttachedCluster{}
	},
	DATAFLOW_FLEXTEMPLATEJOB: func() Any {
		return &dataflow.FlexTemplateJob{}
	},
	MONITORING_METRICDESCRIPTOR: func() Any {
		return &monitoring.MetricDescriptor{}
	},
	NETWORKSERVICES_GATEWAY: func() Any {
		return &networkservices.Gateway{}
	},
	ORGANIZATIONS_IAMPOLICY: func() Any {
		return &organizations.IAMPolicy{}
	},
	NOTEBOOKS_RUNTIMEIAMMEMBER: func() Any {
		return &notebooks.RuntimeIamMember{}
	},
	COMPUTE_REGIONNETWORKFIREWALLPOLICYASSOCIATION: func() Any {
		return &compute.RegionNetworkFirewallPolicyAssociation{}
	},
	GKEBACKUP_BACKUPPLANIAMPOLICY: func() Any {
		return &gkebackup.BackupPlanIamPolicy{}
	},
	GKEONPREM_VMWARECLUSTER: func() Any {
		return &gkeonprem.VMwareCluster{}
	},
	LOGGING_FOLDEREXCLUSION: func() Any {
		return &logging.FolderExclusion{}
	},
	NETAPP_ACTIVEDIRECTORY: func() Any {
		return &netapp.ActiveDirectory{}
	},
	CERTIFICATEAUTHORITY_CAPOOLIAMBINDING: func() Any {
		return &certificateauthority.CaPoolIamBinding{}
	},
	DATAPROC_CLUSTER: func() Any {
		return &dataproc.Cluster{}
	},
	ENDPOINTS_SERVICEIAMPOLICY: func() Any {
		return &endpoints.ServiceIamPolicy{}
	},
	DATAPROC_JOB: func() Any {
		return &dataproc.Job{}
	},
	DIAGFLOW_CXENVIRONMENT: func() Any {
		return &diagflow.CxEnvironment{}
	},
	ORGANIZATIONS_IAMBINDING: func() Any {
		return &organizations.IAMBinding{}
	},
	COMPUTE_TARGETHTTPPROXY: func() Any {
		return &compute.TargetHttpProxy{}
	},
	CLOUDFUNCTIONS_FUNCTIONIAMPOLICY: func() Any {
		return &cloudfunctions.FunctionIamPolicy{}
	},
	HEALTHCARE_DICOMSTOREIAMBINDING: func() Any {
		return &healthcare.DicomStoreIamBinding{}
	},
	KMS_SECRETCIPHERTEXT: func() Any {
		return &kms.SecretCiphertext{}
	},
	NETWORKSECURITY_TLSINSPECTIONPOLICY: func() Any {
		return &networksecurity.TlsInspectionPolicy{}
	},
	IAM_WORKFORCEPOOLPROVIDER: func() Any {
		return &iam.WorkforcePoolProvider{}
	},
	OSCONFIG_PATCHDEPLOYMENT: func() Any {
		return &osconfig.PatchDeployment{}
	},
	CLOUDDEPLOY_DELIVERYPIPELINEIAMMEMBER: func() Any {
		return &clouddeploy.DeliveryPipelineIamMember{}
	},
	COMPUTE_BACKENDSERVICEIAMMEMBER: func() Any {
		return &compute.BackendServiceIamMember{}
	},
	COMPUTE_GLOBALADDRESS: func() Any {
		return &compute.GlobalAddress{}
	},
	ACCESSCONTEXTMANAGER_EGRESSPOLICY: func() Any {
		return &accesscontextmanager.EgressPolicy{}
	},
	APPENGINE_FIREWALLRULE: func() Any {
		return &appengine.FirewallRule{}
	},
	FIREBASE_APPCHECKSERVICECONFIG: func() Any {
		return &firebase.AppCheckServiceConfig{}
	},
	HEALTHCARE_CONSENTSTOREIAMPOLICY: func() Any {
		return &healthcare.ConsentStoreIamPolicy{}
	},
	ORGANIZATIONS_ACCESSAPPROVALSETTINGS: func() Any {
		return &organizations.AccessApprovalSettings{}
	},
	COMPUTE_TARGETPOOL: func() Any {
		return &compute.TargetPool{}
	},
	COMPUTE_NETWORKATTACHMENT: func() Any {
		return &compute.NetworkAttachment{}
	},
	EDGECONTAINER_VPNCONNECTION: func() Any {
		return &edgecontainer.VpnConnection{}
	},
	BINARYAUTHORIZATION_ATTESTOR: func() Any {
		return &binaryauthorization.Attestor{}
	},
	IAP_WEBBACKENDSERVICEIAMBINDING: func() Any {
		return &iap.WebBackendServiceIamBinding{}
	},
	PUBSUB_LITERESERVATION: func() Any {
		return &pubsub.LiteReservation{}
	},
	STORAGE_OBJECTACCESSCONTROL: func() Any {
		return &storage.ObjectAccessControl{}
	},
	NOTEBOOKS_INSTANCEIAMBINDING: func() Any {
		return &notebooks.InstanceIamBinding{}
	},
	WORKSTATIONS_WORKSTATIONIAMMEMBER: func() Any {
		return &workstations.WorkstationIamMember{}
	},
	ACTIVEDIRECTORY_PEERING: func() Any {
		return &activedirectory.Peering{}
	},
	DATAPLEX_DATASCANIAMBINDING: func() Any {
		return &dataplex.DatascanIamBinding{}
	},
	DATAPROC_AUTOSCALINGPOLICY: func() Any {
		return &dataproc.AutoscalingPolicy{}
	},
	DNS_RESPONSEPOLICYRULE: func() Any {
		return &dns.ResponsePolicyRule{}
	},
	CERTIFICATEAUTHORITY_CERTIFICATETEMPLATE: func() Any {
		return &certificateauthority.CertificateTemplate{}
	},
	COMPUTE_NETWORK: func() Any {
		return &compute.Network{}
	},
	VERTEX_AIENDPOINT: func() Any {
		return &vertex.AiEndpoint{}
	},
	SECURITYPOSTURE_POSTUREDEPLOYMENT: func() Any {
		return &securityposture.PostureDeployment{}
	},
	COMPUTE_PERINSTANCECONFIG: func() Any {
		return &compute.PerInstanceConfig{}
	},
	IAP_WEBBACKENDSERVICEIAMMEMBER: func() Any {
		return &iap.WebBackendServiceIamMember{}
	},
	BIGQUERYANALYTICSHUB_DATAEXCHANGE: func() Any {
		return &bigqueryanalyticshub.DataExchange{}
	},
	COMPUTE_PROJECTMETADATA: func() Any {
		return &compute.ProjectMetadata{}
	},
	COMPUTE_SNAPSHOTIAMPOLICY: func() Any {
		return &compute.SnapshotIamPolicy{}
	},
	PROJECTS_IAMPOLICY: func() Any {
		return &projects.IAMPolicy{}
	},
	PUBSUB_SUBSCRIPTIONIAMMEMBER: func() Any {
		return &pubsub.SubscriptionIAMMember{}
	},
	DATABASEMIGRATIONSERVICE_PRIVATECONNECTION: func() Any {
		return &databasemigrationservice.PrivateConnection{}
	},
	GKEHUB_SCOPE: func() Any {
		return &gkehub.Scope{}
	},
	ORGANIZATIONS_PROJECT: func() Any {
		return &organizations.Project{}
	},
	CONTAINERANALYSIS_NOTEIAMBINDING: func() Any {
		return &containeranalysis.NoteIamBinding{}
	},
	HEALTHCARE_DATASETIAMBINDING: func() Any {
		return &healthcare.DatasetIamBinding{}
	},
	ACTIVEDIRECTORY_DOMAINTRUST: func() Any {
		return &activedirectory.DomainTrust{}
	},
	ARTIFACTREGISTRY_REPOSITORYIAMMEMBER: func() Any {
		return &artifactregistry.RepositoryIamMember{}
	},
	BIGTABLE_TABLEIAMMEMBER: func() Any {
		return &bigtable.TableIamMember{}
	},
	PUBSUB_SCHEMAIAMPOLICY: func() Any {
		return &pubsub.SchemaIamPolicy{}
	},
	APIGEE_KEYSTORESALIASESKEYCERTFILE: func() Any {
		return &apigee.KeystoresAliasesKeyCertFile{}
	},
	ENDPOINTS_CONSUMERSIAMMEMBER: func() Any {
		return &endpoints.ConsumersIamMember{}
	},
	GKEHUB_FEATUREIAMBINDING: func() Any {
		return &gkehub.FeatureIamBinding{}
	},
	GKEHUB_SCOPEIAMPOLICY: func() Any {
		return &gkehub.ScopeIamPolicy{}
	},
	RUNTIMECONFIG_VARIABLE: func() Any {
		return &runtimeconfig.Variable{}
	},
	CLOUDFUNCTIONSV2_FUNCTIONIAMPOLICY: func() Any {
		return &cloudfunctionsv2.FunctionIamPolicy{}
	},
	CERTIFICATEAUTHORITY_CAPOOLIAMMEMBER: func() Any {
		return &certificateauthority.CaPoolIamMember{}
	},
	CLOUDTASKS_QUEUEIAMMEMBER: func() Any {
		return &cloudtasks.QueueIamMember{}
	},
	COMPUTE_GLOBALFORWARDINGRULE: func() Any {
		return &compute.GlobalForwardingRule{}
	},
	DIAGFLOW_CXVERSION: func() Any {
		return &diagflow.CxVersion{}
	},
	COMPUTE_BACKENDSERVICESIGNEDURLKEY: func() Any {
		return &compute.BackendServiceSignedUrlKey{}
	},
	COMPUTE_INSTANCEGROUPNAMEDPORT: func() Any {
		return &compute.InstanceGroupNamedPort{}
	},
	MONITORING_DASHBOARD: func() Any {
		return &monitoring.Dashboard{}
	},
	ALLOYDB_CLUSTER: func() Any {
		return &alloydb.Cluster{}
	},
	MONITORING_ALERTPOLICY: func() Any {
		return &monitoring.AlertPolicy{}
	},
	ACCESSCONTEXTMANAGER_ACCESSLEVELCONDITION: func() Any {
		return &accesscontextmanager.AccessLevelCondition{}
	},
	ARTIFACTREGISTRY_REPOSITORYIAMBINDING: func() Any {
		return &artifactregistry.RepositoryIamBinding{}
	},
	BIGQUERY_IAMPOLICY: func() Any {
		return &bigquery.IamPolicy{}
	},
	IAP_APPENGINESERVICEIAMBINDING: func() Any {
		return &iap.AppEngineServiceIamBinding{}
	},
	APIGEE_ENVGROUP: func() Any {
		return &apigee.EnvGroup{}
	},
	COMPUTE_INSTANCEIAMPOLICY: func() Any {
		return &compute.InstanceIAMPolicy{}
	},
	HEALTHCARE_DICOMSTORE: func() Any {
		return &healthcare.DicomStore{}
	},
	INTEGRATIONCONNECTORS_ENDPOINTATTACHMENT: func() Any {
		return &integrationconnectors.EndpointAttachment{}
	},
	BIGQUERY_CONNECTION: func() Any {
		return &bigquery.Connection{}
	},
	COMPUTE_SUBNETWORK: func() Any {
		return &compute.Subnetwork{}
	},
	DATACATALOG_POLICYTAGIAMBINDING: func() Any {
		return &datacatalog.PolicyTagIamBinding{}
	},
	GKEHUB_FEATUREIAMPOLICY: func() Any {
		return &gkehub.FeatureIamPolicy{}
	},
	ORGANIZATIONS_IAMMEMBER: func() Any {
		return &organizations.IAMMember{}
	},
	STORAGE_BUCKETOBJECT: func() Any {
		return &storage.BucketObject{}
	},
	NETAPP_VOLUME: func() Any {
		return &netapp.Volume{}
	},
	ORGANIZATIONS_FOLDER: func() Any {
		return &organizations.Folder{}
	},
	BEYONDCORP_APPGATEWAY: func() Any {
		return &beyondcorp.AppGateway{}
	},
	COMPUTE_HAVPNGATEWAY: func() Any {
		return &compute.HaVpnGateway{}
	},
	APIGEE_INSTANCEATTACHMENT: func() Any {
		return &apigee.InstanceAttachment{}
	},
	TAGS_TAGVALUE: func() Any {
		return &tags.TagValue{}
	},
	COMPUTE_PUBLICDELEGATEDPREFIX: func() Any {
		return &compute.PublicDelegatedPrefix{}
	},
	NETAPP_BACKUPPOLICY: func() Any {
		return &netapp.BackupPolicy{}
	},
	NOTEBOOKS_RUNTIMEIAMPOLICY: func() Any {
		return &notebooks.RuntimeIamPolicy{}
	},
	BIGQUERY_BIRESERVATION: func() Any {
		return &bigquery.BiReservation{}
	},
	COMPUTE_DISKRESOURCEPOLICYATTACHMENT: func() Any {
		return &compute.DiskResourcePolicyAttachment{}
	},
	COMPUTE_IMAGEIAMPOLICY: func() Any {
		return &compute.ImageIamPolicy{}
	},
	GKEONPREM_VMWARENODEPOOL: func() Any {
		return &gkeonprem.VMwareNodePool{}
	},
	APIGATEWAY_GATEWAYIAMMEMBER: func() Any {
		return &apigateway.GatewayIamMember{}
	},
	BIGQUERYDATAPOLICY_DATAPOLICY: func() Any {
		return &bigquerydatapolicy.DataPolicy{}
	},
	BIGQUERYDATAPOLICY_DATAPOLICYIAMPOLICY: func() Any {
		return &bigquerydatapolicy.DataPolicyIamPolicy{}
	},
	BIGTABLE_INSTANCEIAMPOLICY: func() Any {
		return &bigtable.InstanceIamPolicy{}
	},
	NETWORKSERVICES_MESH: func() Any {
		return &networkservices.Mesh{}
	},
	CLOUDRUNV2_JOBIAMPOLICY: func() Any {
		return &cloudrunv2.JobIamPolicy{}
	},
	GKEHUB_FLEET: func() Any {
		return &gkehub.Fleet{}
	},
	IAP_TUNNELIAMBINDING: func() Any {
		return &iap.TunnelIamBinding{}
	},
	NOTEBOOKS_INSTANCEIAMMEMBER: func() Any {
		return &notebooks.InstanceIamMember{}
	},
	VPCACCESS_CONNECTOR: func() Any {
		return &vpcaccess.Connector{}
	},
	CLOUDASSET_ORGANIZATIONFEED: func() Any {
		return &cloudasset.OrganizationFeed{}
	},
	CLOUDFUNCTIONS_FUNCTION: func() Any {
		return &cloudfunctions.Function{}
	},
	COMPUTE_SHAREDVPCHOSTPROJECT: func() Any {
		return &compute.SharedVPCHostProject{}
	},
	DATAPROC_AUTOSCALINGPOLICYIAMMEMBER: func() Any {
		return &dataproc.AutoscalingPolicyIamMember{}
	},
	CLOUDFUNCTIONSV2_FUNCTION: func() Any {
		return &cloudfunctionsv2.Function{}
	},
	COMPUTE_REGIONDISKIAMBINDING: func() Any {
		return &compute.RegionDiskIamBinding{}
	},
	DATASTORE_DATASTOREINDEX: func() Any {
		return &datastore.DataStoreIndex{}
	},
	FILESTORE_SNAPSHOT: func() Any {
		return &filestore.Snapshot{}
	},
	CLOUDDEPLOY_TARGETIAMMEMBER: func() Any {
		return &clouddeploy.TargetIamMember{}
	},
	COMPUTE_NETWORKENDPOINT: func() Any {
		return &compute.NetworkEndpoint{}
	},
	CONTAINERANALYSIS_NOTEIAMMEMBER: func() Any {
		return &containeranalysis.NoteIamMember{}
	},
	KMS_CRYPTOKEYIAMPOLICY: func() Any {
		return &kms.CryptoKeyIAMPolicy{}
	},
	CLOUDRUNV2_SERVICE: func() Any {
		return &cloudrunv2.Service{}
	},
	DATAPROC_METASTORESERVICE: func() Any {
		return &dataproc.MetastoreService{}
	},
	ALLOYDB_BACKUP: func() Any {
		return &alloydb.Backup{}
	},
	DIAGFLOW_ENTITYTYPE: func() Any {
		return &diagflow.EntityType{}
	},
	VERTEX_AIENDPOINTIAMPOLICY: func() Any {
		return &vertex.AiEndpointIamPolicy{}
	},
	DATAFORM_REPOSITORYWORKFLOWCONFIG: func() Any {
		return &dataform.RepositoryWorkflowConfig{}
	},
	NETWORKCONNECTIVITY_SERVICECONNECTIONPOLICY: func() Any {
		return &networkconnectivity.ServiceConnectionPolicy{}
	},
	PUBSUB_SCHEMA: func() Any {
		return &pubsub.Schema{}
	},
	TAGS_TAGKEY: func() Any {
		return &tags.TagKey{}
	},
	ALLOYDB_INSTANCE: func() Any {
		return &alloydb.Instance{}
	},
	APIGEE_SYNCAUTHORIZATION: func() Any {
		return &apigee.SyncAuthorization{}
	},
	COMPUTE_MACHINEIMAGEIAMMEMBER: func() Any {
		return &compute.MachineImageIamMember{}
	},
	COMPUTE_REGIONSECURITYPOLICY: func() Any {
		return &compute.RegionSecurityPolicy{}
	},
	VMWAREENGINE_SUBNET: func() Any {
		return &vmwareengine.Subnet{}
	},
	BIGQUERY_DATASETIAMPOLICY: func() Any {
		return &bigquery.DatasetIamPolicy{}
	},
	DATAPLEX_LAKEIAMPOLICY: func() Any {
		return &dataplex.LakeIamPolicy{}
	},
	ORGPOLICY_CUSTOMCONSTRAINT: func() Any {
		return &orgpolicy.CustomConstraint{}
	},
	SECURESOURCEMANAGER_INSTANCE: func() Any {
		return &securesourcemanager.Instance{}
	},
	COMPUTE_NETWORKENDPOINTLIST: func() Any {
		return &compute.NetworkEndpointList{}
	},
	COMPUTE_REGIONDISK: func() Any {
		return &compute.RegionDisk{}
	},
	DATALOSS_PREVENTIONDEIDENTIFYTEMPLATE: func() Any {
		return &dataloss.PreventionDeidentifyTemplate{}
	},
	FIREBASE_APPCHECKPLAYINTEGRITYCONFIG: func() Any {
		return &firebase.AppCheckPlayIntegrityConfig{}
	},
	APIGATEWAY_APICONFIGIAMMEMBER: func() Any {
		return &apigateway.ApiConfigIamMember{}
	},
	BIGQUERYANALYTICSHUB_LISTINGIAMPOLICY: func() Any {
		return &bigqueryanalyticshub.ListingIamPolicy{}
	},
	CLOUDTASKS_QUEUEIAMPOLICY: func() Any {
		return &cloudtasks.QueueIamPolicy{}
	},
	COMPUTE_FORWARDINGRULE: func() Any {
		return &compute.ForwardingRule{}
	},
	WORKSTATIONS_WORKSTATION: func() Any {
		return &workstations.Workstation{}
	},
	IAP_APPENGINEVERSIONIAMBINDING: func() Any {
		return &iap.AppEngineVersionIamBinding{}
	},
	SECURITYCENTER_PROJECTCUSTOMMODULE: func() Any {
		return &securitycenter.ProjectCustomModule{}
	},
	SERVICEDIRECTORY_NAMESPACEIAMMEMBER: func() Any {
		return &servicedirectory.NamespaceIamMember{}
	},
	GKEHUB_MEMBERSHIP: func() Any {
		return &gkehub.Membership{}
	},
	FIREBASE_HOSTINGVERSION: func() Any {
		return &firebase.HostingVersion{}
	},
	DATAPLEX_ASSET: func() Any {
		return &dataplex.Asset{}
	},
	COMPUTE_URLMAP: func() Any {
		return &compute.URLMap{}
	},
	DIAGFLOW_CXPAGE: func() Any {
		return &diagflow.CxPage{}
	},
	FOLDER_IAMPOLICY: func() Any {
		return &folder.IAMPolicy{}
	},
	SECRETMANAGER_SECRET: func() Any {
		return &secretmanager.Secret{}
	},
	REDIS_CLUSTER: func() Any {
		return &redis.Cluster{}
	},
	APIGEE_KEYSTORESALIASESSELFSIGNEDCERT: func() Any {
		return &apigee.KeystoresAliasesSelfSignedCert{}
	},
	ASSUREDWORKLOADS_WORKLOAD: func() Any {
		return &assuredworkloads.Workload{}
	},
	CLOUDDEPLOY_TARGETIAMPOLICY: func() Any {
		return &clouddeploy.TargetIamPolicy{}
	},
	COMPUTE_SERVICEATTACHMENT: func() Any {
		return &compute.ServiceAttachment{}
	},
	VERTEX_AIFEATURESTOREENTITYTYPE: func() Any {
		return &vertex.AiFeatureStoreEntityType{}
	},
	WORKBENCH_INSTANCEIAMPOLICY: func() Any {
		return &workbench.InstanceIamPolicy{}
	},
	APIGEE_ENVIRONMENTIAMPOLICY: func() Any {
		return &apigee.EnvironmentIamPolicy{}
	},
	BIGQUERY_TABLE: func() Any {
		return &bigquery.Table{}
	},
	CLOUDBUILDV2_CONNECTIONIAMPOLICY: func() Any {
		return &cloudbuildv2.ConnectionIAMPolicy{}
	},
	GKEHUB_MEMBERSHIPBINDING: func() Any {
		return &gkehub.MembershipBinding{}
	},
	ACCESSCONTEXTMANAGER_INGRESSPOLICY: func() Any {
		return &accesscontextmanager.IngressPolicy{}
	},
	COMPUTE_REGIONDISKIAMPOLICY: func() Any {
		return &compute.RegionDiskIamPolicy{}
	},
	DATAPLEX_DATASCANIAMMEMBER: func() Any {
		return &dataplex.DatascanIamMember{}
	},
	BIGQUERY_IAMMEMBER: func() Any {
		return &bigquery.IamMember{}
	},
	COMPUTE_BACKENDBUCKETIAMMEMBER: func() Any {
		return &compute.BackendBucketIamMember{}
	},
	COMPUTE_BACKENDSERVICE: func() Any {
		return &compute.BackendService{}
	},
	SECURITYCENTER_EVENTTHREATDETECTIONCUSTOMMODULE: func() Any {
		return &securitycenter.EventThreatDetectionCustomModule{}
	},
	HEALTHCARE_DICOMSTOREIAMPOLICY: func() Any {
		return &healthcare.DicomStoreIamPolicy{}
	},
	HEALTHCARE_FHIRSTOREIAMMEMBER: func() Any {
		return &healthcare.FhirStoreIamMember{}
	},
	MONITORING_UPTIMECHECKCONFIG: func() Any {
		return &monitoring.UptimeCheckConfig{}
	},
	APIGATEWAY_APICONFIGIAMPOLICY: func() Any {
		return &apigateway.ApiConfigIamPolicy{}
	},
	COMPUTE_SHAREDVPCSERVICEPROJECT: func() Any {
		return &compute.SharedVPCServiceProject{}
	},
	DATAPLEX_TASKIAMPOLICY: func() Any {
		return &dataplex.TaskIamPolicy{}
	},
	VERTEX_AIFEATURESTOREENTITYTYPEIAMMEMBER: func() Any {
		return &vertex.AiFeatureStoreEntityTypeIamMember{}
	},
	CLOUDTASKS_QUEUE: func() Any {
		return &cloudtasks.Queue{}
	},
	COMPUTE_INSTANCEIAMBINDING: func() Any {
		return &compute.InstanceIAMBinding{}
	},
	DATACATALOG_ENTRYGROUPIAMPOLICY: func() Any {
		return &datacatalog.EntryGroupIamPolicy{}
	},
	APIGEE_ADDONSCONFIG: func() Any {
		return &apigee.AddonsConfig{}
	},
	COMPUTE_REGIONPERINSTANCECONFIG: func() Any {
		return &compute.RegionPerInstanceConfig{}
	},
	DNS_MANAGEDZONE: func() Any {
		return &dns.ManagedZone{}
	},
	VERTEX_AIFEATUREONLINESTORE: func() Any {
		return &vertex.AiFeatureOnlineStore{}
	},
	VERTEX_AITENSORBOARD: func() Any {
		return &vertex.AiTensorboard{}
	},
	APIGATEWAY_APIIAMBINDING: func() Any {
		return &apigateway.ApiIamBinding{}
	},
	BIGQUERY_DATASETIAMMEMBER: func() Any {
		return &bigquery.DatasetIamMember{}
	},
	COMPUTE_ROUTE: func() Any {
		return &compute.Route{}
	},
	DATAPLEX_DATASCAN: func() Any {
		return &dataplex.Datascan{}
	},
	DATACATALOG_ENTRY: func() Any {
		return &datacatalog.Entry{}
	},
	HEALTHCARE_FHIRSTOREIAMPOLICY: func() Any {
		return &healthcare.FhirStoreIamPolicy{}
	},
	IDENTITYPLATFORM_DEFAULTSUPPORTEDIDPCONFIG: func() Any {
		return &identityplatform.DefaultSupportedIdpConfig{}
	},
	CLOUDASSET_FOLDERFEED: func() Any {
		return &cloudasset.FolderFeed{}
	},
	EDGECONTAINER_NODEPOOL: func() Any {
		return &edgecontainer.NodePool{}
	},
	APIGEE_ENVGROUPATTACHMENT: func() Any {
		return &apigee.EnvGroupAttachment{}
	},
	DATABASEMIGRATIONSERVICE_CONNECTIONPROFILE: func() Any {
		return &databasemigrationservice.ConnectionProfile{}
	},
	FIREBASE_APPLEAPP: func() Any {
		return &firebase.AppleApp{}
	},
	SERVICEDIRECTORY_ENDPOINT: func() Any {
		return &servicedirectory.Endpoint{}
	},
	BEYONDCORP_APPCONNECTION: func() Any {
		return &beyondcorp.AppConnection{}
	},
	COMPUTE_SUBNETWORKIAMBINDING: func() Any {
		return &compute.SubnetworkIAMBinding{}
	},
	VERTEX_AIENDPOINTIAMMEMBER: func() Any {
		return &vertex.AiEndpointIamMember{}
	},
	GKEHUB_FEATUREMEMBERSHIP: func() Any {
		return &gkehub.FeatureMembership{}
	},
	LOOKER_INSTANCE: func() Any {
		return &looker.Instance{}
	},
	MEMCACHE_INSTANCE: func() Any {
		return &memcache.Instance{}
	},
	PROJECTS_IAMCUSTOMROLE: func() Any {
		return &projects.IAMCustomRole{}
	},
	ACCESSCONTEXTMANAGER_SERVICEPERIMETEREGRESSPOLICY: func() Any {
		return &accesscontextmanager.ServicePerimeterEgressPolicy{}
	},
	COMPUTE_IMAGE: func() Any {
		return &compute.Image{}
	},
	COMPUTE_PUBLICADVERTISEDPREFIX: func() Any {
		return &compute.PublicAdvertisedPrefix{}
	},
	SERVICEDIRECTORY_SERVICEIAMPOLICY: func() Any {
		return &servicedirectory.ServiceIamPolicy{}
	},
	CLOUDDEPLOY_DELIVERYPIPELINEIAMPOLICY: func() Any {
		return &clouddeploy.DeliveryPipelineIamPolicy{}
	},
	COMPUTE_REGIONBACKENDSERVICEIAMBINDING: func() Any {
		return &compute.RegionBackendServiceIamBinding{}
	},
	IAP_WEBTYPECOMPUTEIAMPOLICY: func() Any {
		return &iap.WebTypeComputeIamPolicy{}
	},
	SQL_USER: func() Any {
		return &sql.User{}
	},
	CERTIFICATEMANAGER_TRUSTCONFIG: func() Any {
		return &certificatemanager.TrustConfig{}
	},
	WORKBENCH_INSTANCEIAMMEMBER: func() Any {
		return &workbench.InstanceIamMember{}
	},
	COMPUTE_INSTANCEGROUPMEMBERSHIP: func() Any {
		return &compute.InstanceGroupMembership{}
	},
	IAP_WEBTYPEAPPENGINGIAMMEMBER: func() Any {
		return &iap.WebTypeAppEngingIamMember{}
	},
	BIGTABLE_INSTANCEIAMBINDING: func() Any {
		return &bigtable.InstanceIamBinding{}
	},
	CLOUDBUILD_TRIGGER: func() Any {
		return &cloudbuild.Trigger{}
	},
	BIGQUERY_DATATRANSFERCONFIG: func() Any {
		return &bigquery.DataTransferConfig{}
	},
	COMPUTE_REGIONSSLPOLICY: func() Any {
		return &compute.RegionSslPolicy{}
	},
	DATAPROC_METASTORESERVICEIAMMEMBER: func() Any {
		return &dataproc.MetastoreServiceIamMember{}
	},
	DIAGFLOW_CXWEBHOOK: func() Any {
		return &diagflow.CxWebhook{}
	},
	VMWAREENGINE_CLUSTER: func() Any {
		return &vmwareengine.Cluster{}
	},
	CLOUDRUNV2_JOB: func() Any {
		return &cloudrunv2.Job{}
	},
	COMPUTE_HTTPSHEALTHCHECK: func() Any {
		return &compute.HttpsHealthCheck{}
	},
	COMPUTE_SECURITYPOLICY: func() Any {
		return &compute.SecurityPolicy{}
	},
	DATAPROC_METASTOREFEDERATIONIAMBINDING: func() Any {
		return &dataproc.MetastoreFederationIamBinding{}
	},
	DIAGFLOW_CXFLOW: func() Any {
		return &diagflow.CxFlow{}
	},
	SECURITYCENTER_SOURCEIAMMEMBER: func() Any {
		return &securitycenter.SourceIamMember{}
	},
	SERVICEDIRECTORY_NAMESPACEIAMPOLICY: func() Any {
		return &servicedirectory.NamespaceIamPolicy{}
	},
	TPU_NODE: func() Any {
		return &tpu.Node{}
	},
	NETAPP_BACKUPVAULT: func() Any {
		return &netapp.BackupVault{}
	},
	NETWORKSERVICES_EDGECACHEORIGIN: func() Any {
		return &networkservices.EdgeCacheOrigin{}
	},
	WORKSTATIONS_WORKSTATIONCONFIG: func() Any {
		return &workstations.WorkstationConfig{}
	},
	BIGLAKE_CATALOG: func() Any {
		return &biglake.Catalog{}
	},
	CLOUDBUILD_WORKERPOOL: func() Any {
		return &cloudbuild.WorkerPool{}
	},
	CLOUDBUILDV2_CONNECTION: func() Any {
		return &cloudbuildv2.Connection{}
	},
	IAP_APPENGINEVERSIONIAMPOLICY: func() Any {
		return &iap.AppEngineVersionIamPolicy{}
	},
	APIGEE_ENVIRONMENTIAMMEMBER: func() Any {
		return &apigee.EnvironmentIamMember{}
	},
	BIGTABLE_INSTANCE: func() Any {
		return &bigtable.Instance{}
	},
	DEPLOYMENTMANAGER_DEPLOYMENT: func() Any {
		return &deploymentmanager.Deployment{}
	},
	DATAPLEX_DATASCANIAMPOLICY: func() Any {
		return &dataplex.DatascanIamPolicy{}
	},
	LOGGING_LOGVIEW: func() Any {
		return &logging.LogView{}
	},
	SERVICEDIRECTORY_NAMESPACEIAMBINDING: func() Any {
		return &servicedirectory.NamespaceIamBinding{}
	},
	DATAFORM_REPOSITORYIAMMEMBER: func() Any {
		return &dataform.RepositoryIamMember{}
	},
	IDENTITYPLATFORM_INBOUNDSAMLCONFIG: func() Any {
		return &identityplatform.InboundSamlConfig{}
	},
	TAGS_TAGKEYIAMBINDING: func() Any {
		return &tags.TagKeyIamBinding{}
	},
	DIAGFLOW_CXENTITYTYPE: func() Any {
		return &diagflow.CxEntityType{}
	},
	CONTAINERANALYSIS_NOTE: func() Any {
		return &containeranalysis.Note{}
	},
	GKEHUB_SCOPERBACROLEBINDING: func() Any {
		return &gkehub.ScopeRbacRoleBinding{}
	},
	LOGGING_FOLDERSETTINGS: func() Any {
		return &logging.FolderSettings{}
	},
	CLOUDIDS_ENDPOINT: func() Any {
		return &cloudids.Endpoint{}
	},
	COMPUTE_BACKENDBUCKETIAMBINDING: func() Any {
		return &compute.BackendBucketIamBinding{}
	},
	TAGS_TAGVALUEIAMBINDING: func() Any {
		return &tags.TagValueIamBinding{}
	},
	NETWORKSERVICES_TLSROUTE: func() Any {
		return &networkservices.TlsRoute{}
	},
	CLOUDDEPLOY_TARGET: func() Any {
		return &clouddeploy.Target{}
	},
	COMPUTE_INSTANCEFROMTEMPLATE: func() Any {
		return &compute.InstanceFromTemplate{}
	},
	IAP_APPENGINESERVICEIAMMEMBER: func() Any {
		return &iap.AppEngineServiceIamMember{}
	},
	INTEGRATIONCONNECTORS_CONNECTION: func() Any {
		return &integrationconnectors.Connection{}
	},
	COMPUTE_SNAPSHOT: func() Any {
		return &compute.Snapshot{}
	},
	NETWORKSECURITY_AUTHORIZATIONPOLICY: func() Any {
		return &networksecurity.AuthorizationPolicy{}
	},
	NOTEBOOKS_RUNTIMEIAMBINDING: func() Any {
		return &notebooks.RuntimeIamBinding{}
	},
	PROJECTS_DEFAULTSERVICEACCOUNTS: func() Any {
		return &projects.DefaultServiceAccounts{}
	},
	VMWAREENGINE_EXTERNALADDRESS: func() Any {
		return &vmwareengine.ExternalAddress{}
	},
	ORGANIZATIONS_IAMAUDITCONFIG: func() Any {
		return &organizations.IamAuditConfig{}
	},
	BIGQUERY_JOB: func() Any {
		return &bigquery.Job{}
	},
	CLOUDRUN_DOMAINMAPPING: func() Any {
		return &cloudrun.DomainMapping{}
	},
	DNS_DNSMANAGEDZONEIAMMEMBER: func() Any {
		return &dns.DnsManagedZoneIamMember{}
	},
	GKEBACKUP_RESTOREPLANIAMPOLICY: func() Any {
		return &gkebackup.RestorePlanIamPolicy{}
	},
	DATAFLOW_PIPELINE: func() Any {
		return &dataflow.Pipeline{}
	},
	REDIS_INSTANCE: func() Any {
		return &redis.Instance{}
	},
	DATAPROC_METASTOREFEDERATIONIAMMEMBER: func() Any {
		return &dataproc.MetastoreFederationIamMember{}
	},
	FIREBASE_HOSTINGCUSTOMDOMAIN: func() Any {
		return &firebase.HostingCustomDomain{}
	},
	IAM_DENYPOLICY: func() Any {
		return &iam.DenyPolicy{}
	},
	CLOUDRUN_IAMBINDING: func() Any {
		return &cloudrun.IamBinding{}
	},
	HEALTHCARE_HL7STOREIAMBINDING: func() Any {
		return &healthcare.Hl7StoreIamBinding{}
	},
	NETWORKSECURITY_SERVERTLSPOLICY: func() Any {
		return &networksecurity.ServerTlsPolicy{}
	},
	VERTEX_AIENDPOINTIAMBINDING: func() Any {
		return &vertex.AiEndpointIamBinding{}
	},
	IDENTITYPLATFORM_OAUTHIDPCONFIG: func() Any {
		return &identityplatform.OauthIdpConfig{}
	},
	NOTEBOOKS_INSTANCEIAMPOLICY: func() Any {
		return &notebooks.InstanceIamPolicy{}
	},
	RECAPTCHA_ENTERPRISEKEY: func() Any {
		return &recaptcha.EnterpriseKey{}
	},
	APPENGINE_STANDARDAPPVERSION: func() Any {
		return &appengine.StandardAppVersion{}
	},
	FIREBASE_STORAGEBUCKET: func() Any {
		return &firebase.StorageBucket{}
	},
	LOGGING_PROJECTSINK: func() Any {
		return &logging.ProjectSink{}
	},
	NOTEBOOKS_RUNTIME: func() Any {
		return &notebooks.Runtime{}
	},
	SERVICEDIRECTORY_SERVICE: func() Any {
		return &servicedirectory.Service{}
	},
	COMPUTE_SECURITYSCANCONFIG: func() Any {
		return &compute.SecurityScanConfig{}
	},
	DATAPROC_WORKFLOWTEMPLATE: func() Any {
		return &dataproc.WorkflowTemplate{}
	},
	DIAGFLOW_CXTESTCASE: func() Any {
		return &diagflow.CxTestCase{}
	},
	SERVICEDIRECTORY_NAMESPACE: func() Any {
		return &servicedirectory.Namespace{}
	},
	COMPUTE_DISKASYNCREPLICATION: func() Any {
		return &compute.DiskAsyncReplication{}
	},
	COMPUTE_NETWORKENDPOINTGROUP: func() Any {
		return &compute.NetworkEndpointGroup{}
	},
	COMPUTE_RESOURCEPOLICY: func() Any {
		return &compute.ResourcePolicy{}
	},
	EVENTARC_CHANNEL: func() Any {
		return &eventarc.Channel{}
	},
	HEALTHCARE_CONSENTSTOREIAMMEMBER: func() Any {
		return &healthcare.ConsentStoreIamMember{}
	},
	KMS_KEYRINGIMPORTJOB: func() Any {
		return &kms.KeyRingImportJob{}
	},
	IAP_APPENGINESERVICEIAMPOLICY: func() Any {
		return &iap.AppEngineServiceIamPolicy{}
	},
	VERTEX_AIFEATURESTORE: func() Any {
		return &vertex.AiFeatureStore{}
	},
	SERVICEACCOUNT_IAMBINDING: func() Any {
		return &serviceaccount.IAMBinding{}
	},
	PROJECTS_ORGANIZATIONPOLICY: func() Any {
		return &projects.OrganizationPolicy{}
	},
	VMWAREENGINE_PRIVATECLOUD: func() Any {
		return &vmwareengine.PrivateCloud{}
	},
	APIGATEWAY_APICONFIGIAMBINDING: func() Any {
		return &apigateway.ApiConfigIamBinding{}
	},
	BIGTABLE_INSTANCEIAMMEMBER: func() Any {
		return &bigtable.InstanceIamMember{}
	},
	DISCOVERYENGINE_DATASTORE: func() Any {
		return &discoveryengine.DataStore{}
	},
	ENDPOINTS_SERVICEIAMMEMBER: func() Any {
		return &endpoints.ServiceIamMember{}
	},
	BIGQUERY_CONNECTIONIAMPOLICY: func() Any {
		return &bigquery.ConnectionIamPolicy{}
	},
	COMPUTE_NODEGROUP: func() Any {
		return &compute.NodeGroup{}
	},
	COMPUTE_REGIONNETWORKENDPOINTGROUP: func() Any {
		return &compute.RegionNetworkEndpointGroup{}
	},
	NETWORKSERVICES_EDGECACHESERVICE: func() Any {
		return &networkservices.EdgeCacheService{}
	},
	BINARYAUTHORIZATION_POLICY: func() Any {
		return &binaryauthorization.Policy{}
	},
	NETWORKSERVICES_HTTPROUTE: func() Any {
		return &networkservices.HttpRoute{}
	},
	STORAGE_BUCKET: func() Any {
		return &storage.Bucket{}
	},
	BIGLAKE_DATABASE: func() Any {
		return &biglake.Database{}
	},
	NETAPP_KMSCONFIG: func() Any {
		return &netapp.Kmsconfig{}
	},
	SPANNER_DATABASEIAMMEMBER: func() Any {
		return &spanner.DatabaseIAMMember{}
	},
	FOLDER_IAMAUDITCONFIG: func() Any {
		return &folder.IamAuditConfig{}
	},
	SPANNER_INSTANCEIAMBINDING: func() Any {
		return &spanner.InstanceIAMBinding{}
	},
	APIGATEWAY_GATEWAYIAMPOLICY: func() Any {
		return &apigateway.GatewayIamPolicy{}
	},
	APIGEE_INSTANCE: func() Any {
		return &apigee.Instance{}
	},
	BIGQUERY_DATASETACCESS: func() Any {
		return &bigquery.DatasetAccess{}
	},
	FOLDER_IAMMEMBER: func() Any {
		return &folder.IAMMember{}
	},
	DIAGFLOW_CXAGENT: func() Any {
		return &diagflow.CxAgent{}
	},
	COMPUTE_AUTOSCALER: func() Any {
		return &compute.Autoscaler{}
	},
	CONTAINERANALYSIS_NOTEIAMPOLICY: func() Any {
		return &containeranalysis.NoteIamPolicy{}
	},
	DATAPLEX_TASK: func() Any {
		return &dataplex.Task{}
	},
	DATAPROC_CLUSTERIAMBINDING: func() Any {
		return &dataproc.ClusterIAMBinding{}
	},
	CERTIFICATEMANAGER_CERTIFICATE: func() Any {
		return &certificatemanager.Certificate{}
	},
	LOGGING_BILLINGACCOUNTEXCLUSION: func() Any {
		return &logging.BillingAccountExclusion{}
	},
	SQL_DATABASEINSTANCE: func() Any {
		return &sql.DatabaseInstance{}
	},
	LOGGING_METRIC: func() Any {
		return &logging.Metric{}
	},
	WORKSTATIONS_WORKSTATIONIAMPOLICY: func() Any {
		return &workstations.WorkstationIamPolicy{}
	},
	NETAPP_VOLUMEREPLICATION: func() Any {
		return &netapp.VolumeReplication{}
	},
	SECURITYCENTER_INSTANCEIAMMEMBER: func() Any {
		return &securitycenter.InstanceIamMember{}
	},
	COMPUTE_EXTERNALVPNGATEWAY: func() Any {
		return &compute.ExternalVpnGateway{}
	},
	COMPUTE_RESERVATION: func() Any {
		return &compute.Reservation{}
	},
	DATAPLEX_ZONE: func() Any {
		return &dataplex.Zone{}
	},
	GKEBACKUP_BACKUPPLAN: func() Any {
		return &gkebackup.BackupPlan{}
	},
	COMPUTE_NETWORKFIREWALLPOLICYASSOCIATION: func() Any {
		return &compute.NetworkFirewallPolicyAssociation{}
	},
	COMPUTE_VPNGATEWAY: func() Any {
		return &compute.VPNGateway{}
	},
	MONITORING_GENERICSERVICE: func() Any {
		return &monitoring.GenericService{}
	},
	RUNTIMECONFIG_CONFIG: func() Any {
		return &runtimeconfig.Config{}
	},
	COMPUTE_REGIONTARGETHTTPSPROXY: func() Any {
		return &compute.RegionTargetHttpsProxy{}
	},
	STORAGE_BUCKETIAMBINDING: func() Any {
		return &storage.BucketIAMBinding{}
	},
	SERVICEDIRECTORY_SERVICEIAMMEMBER: func() Any {
		return &servicedirectory.ServiceIamMember{}
	},
	CERTIFICATEAUTHORITY_CERTIFICATETEMPLATEIAMMEMBER: func() Any {
		return &certificateauthority.CertificateTemplateIamMember{}
	},
	DATAPLEX_ASSETIAMBINDING: func() Any {
		return &dataplex.AssetIamBinding{}
	},
	GKEBACKUP_RESTOREPLAN: func() Any {
		return &gkebackup.RestorePlan{}
	},
	NETWORKSECURITY_ADDRESSGROUPIAMBINDING: func() Any {
		return &networksecurity.AddressGroupIamBinding{}
	},
	BIGQUERY_RESERVATION: func() Any {
		return &bigquery.Reservation{}
	},
	CLOUDASSET_PROJECTFEED: func() Any {
		return &cloudasset.ProjectFeed{}
	},
	COMPUTE_DISKIAMMEMBER: func() Any {
		return &compute.DiskIamMember{}
	},
	MONITORING_CUSTOMSERVICE: func() Any {
		return &monitoring.CustomService{}
	},
	NETWORKSECURITY_SECURITYPROFILEGROUP: func() Any {
		return &networksecurity.SecurityProfileGroup{}
	},
	WORKSTATIONS_WORKSTATIONCONFIGIAMBINDING: func() Any {
		return &workstations.WorkstationConfigIamBinding{}
	},
	IAP_APPENGINEVERSIONIAMMEMBER: func() Any {
		return &iap.AppEngineVersionIamMember{}
	},
	PROJECTS_IAMBINDING: func() Any {
		return &projects.IAMBinding{}
	},
	NETWORKCONNECTIVITY_SPOKE: func() Any {
		return &networkconnectivity.Spoke{}
	},
	APIGATEWAY_APICONFIG: func() Any {
		return &apigateway.ApiConfig{}
	},
	GKEHUB_MEMBERSHIPIAMBINDING: func() Any {
		return &gkehub.MembershipIamBinding{}
	},
	SERVICEDIRECTORY_SERVICEIAMBINDING: func() Any {
		return &servicedirectory.ServiceIamBinding{}
	},
	ARTIFACTREGISTRY_VPCSCCONFIG: func() Any {
		return &artifactregistry.VpcscConfig{}
	},
	TPU_V2VM: func() Any {
		return &tpu.V2Vm{}
	},
	CLOUDBUILDV2_CONNECTIONIAMBINDING: func() Any {
		return &cloudbuildv2.ConnectionIAMBinding{}
	},
	NETWORKSERVICES_GRPCROUTE: func() Any {
		return &networkservices.GrpcRoute{}
	},
	COMPUTE_REGIONBACKENDSERVICE: func() Any {
		return &compute.RegionBackendService{}
	},
	COMPUTE_REGIONBACKENDSERVICEIAMMEMBER: func() Any {
		return &compute.RegionBackendServiceIamMember{}
	},
	CONTAINER_AWSNODEPOOL: func() Any {
		return &container.AwsNodePool{}
	},
	DATACATALOG_ENTRYGROUPIAMBINDING: func() Any {
		return &datacatalog.EntryGroupIamBinding{}
	},
	APPENGINE_ENGINESPLITTRAFFIC: func() Any {
		return &appengine.EngineSplitTraffic{}
	},
	APPENGINE_FLEXIBLEAPPVERSION: func() Any {
		return &appengine.FlexibleAppVersion{}
	},
	BIGQUERY_ROUTINE: func() Any {
		return &bigquery.Routine{}
	},
	COMPUTE_ORGANIZATIONSECURITYPOLICY: func() Any {
		return &compute.OrganizationSecurityPolicy{}
	},
	FIREBASE_HOSTINGRELEASE: func() Any {
		return &firebase.HostingRelease{}
	},
	HEALTHCARE_FHIRSTORE: func() Any {
		return &healthcare.FhirStore{}
	},
	IAP_WEBIAMBINDING: func() Any {
		return &iap.WebIamBinding{}
	},
	COMPUTE_REGIONNETWORKFIREWALLPOLICYRULE: func() Any {
		return &compute.RegionNetworkFirewallPolicyRule{}
	},
	SPANNER_DATABASEIAMPOLICY: func() Any {
		return &spanner.DatabaseIAMPolicy{}
	},
	CERTIFICATEMANAGER_CERTIFICATEMAPENTRY: func() Any {
		return &certificatemanager.CertificateMapEntry{}
	},
	LOGGING_FOLDERBUCKETCONFIG: func() Any {
		return &logging.FolderBucketConfig{}
	},
}
