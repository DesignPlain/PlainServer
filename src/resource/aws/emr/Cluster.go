package emr

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	// Custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
	CustomAmiId string `json:"customAmiId,omitempty" yaml:"customAmiId,omitempty"`

	// Release label for the Amazon EMR release.
	ReleaseLabel string `json:"releaseLabel,omitempty" yaml:"releaseLabel,omitempty"`

	// Ordered list of bootstrap actions that will be run before Hadoop is started on the cluster nodes. See below.
	BootstrapActions []types.Emr_ClusterBootstrapAction `json:"bootstrapActions,omitempty" yaml:"bootstrapActions,omitempty"`

	// AWS KMS customer master key (CMK) key ID or arn used for encrypting log files. This attribute is only available with EMR version 5.30.0 and later, excluding EMR 6.0.0.
	LogEncryptionKmsKeyId string `json:"logEncryptionKmsKeyId,omitempty" yaml:"logEncryptionKmsKeyId,omitempty"`

	// Configuration block to use an [Instance Fleet](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-fleet.html) for the master node type. Cannot be specified if any `master_instance_group` configuration blocks are set. Detailed below.
	MasterInstanceFleet types.Emr_ClusterMasterInstanceFleet `json:"masterInstanceFleet,omitempty" yaml:"masterInstanceFleet,omitempty"`

	// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [master node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-master).
	MasterInstanceGroup types.Emr_ClusterMasterInstanceGroup `json:"masterInstanceGroup,omitempty" yaml:"masterInstanceGroup,omitempty"`

	/*
	   IAM role that will be assumed by the Amazon EMR service to access AWS resources.

	   The following arguments are optional:
	*/
	ServiceRole string `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`

	// Configuration block to use an [Instance Fleet](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-fleet.html) for the core node type. Cannot be specified if any `core_instance_group` configuration blocks are set. Detailed below.
	CoreInstanceFleet types.Emr_ClusterCoreInstanceFleet `json:"coreInstanceFleet,omitempty" yaml:"coreInstanceFleet,omitempty"`

	// Kerberos configuration for the cluster. See below.
	KerberosAttributes types.Emr_ClusterKerberosAttributes `json:"kerberosAttributes,omitempty" yaml:"kerberosAttributes,omitempty"`

	// List of [step states](https://docs.aws.amazon.com/emr/latest/APIReference/API_StepStatus.html) used to filter returned steps
	ListStepsStates []string `json:"listStepsStates,omitempty" yaml:"listStepsStates,omitempty"`

	// Attributes for the EC2 instances running the job flow. See below.
	Ec2Attributes types.Emr_ClusterEc2Attributes `json:"ec2Attributes,omitempty" yaml:"ec2Attributes,omitempty"`

	// Way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized.
	ScaleDownBehavior string `json:"scaleDownBehavior,omitempty" yaml:"scaleDownBehavior,omitempty"`

	// Switch on/off termination protection (default is `false`, except when using multiple master nodes). Before attempting to destroy the resource when termination protection is enabled, this configuration must be applied with its value set to `false`.
	TerminationProtection bool `json:"terminationProtection,omitempty" yaml:"terminationProtection,omitempty"`

	// List of configurations supplied for the EMR cluster you are creating. Supply a configuration object for applications to override their default configuration. See [AWS Documentation](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) for more information.
	Configurations string `json:"configurations,omitempty" yaml:"configurations,omitempty"`

	/*
	   JSON string for supplying list of configurations for the EMR cluster.

	   > --NOTE on `configurations_json`:-- If the `Configurations` value is empty then you should skip the `Configurations` field instead of providing an empty list as a value, `"Configurations": []`.

	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const cluster = new aws.emr.Cluster("cluster", {configurationsJson: `[
	   {
	   "Classification": "hadoop-env",
	   "Configurations": [
	   {
	   "Classification": "export",
	   "Properties": {
	   "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
	   }
	   }
	   ],
	   "Properties": {}
	   }
	   ]
	   `});
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   cluster = aws.emr.Cluster("cluster", configurations_json="""[
	   {
	   "Classification": "hadoop-env",
	   "Configurations": [
	   {
	   "Classification": "export",
	   "Properties": {
	   "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
	   }
	   }
	   ],
	   "Properties": {}
	   }
	   ]
	   """)
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var cluster = new Aws.Emr.Cluster("cluster", new()
	       {
	           ConfigurationsJson = @"[
	   {
	   ""Classification"": ""hadoop-env"",
	   ""Configurations"": [
	   {
	   ""Classification"": ""export"",
	   ""Properties"": {
	   ""JAVA_HOME"": ""/usr/lib/jvm/java-1.8.0""
	   }
	   }
	   ],
	   ""Properties"": {}
	   }
	   ]
	   ",
	       });

	   });
	   ```
	   ```go
	   package main

	   import (
	   	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	   	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	   )

	   func main() {
	   	pulumi.Run(func(ctx -pulumi.Context) error {
	   		_, err := emr.NewCluster(ctx, "cluster", &emr.ClusterArgs{
	   			ConfigurationsJson: pulumi.String(`[
	   {
	   "Classification": "hadoop-env",
	   "Configurations": [
	   {
	   "Classification": "export",
	   "Properties": {
	   "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
	   }
	   }
	   ],
	   "Properties": {}
	   }
	   ]
	   `),
	   		})
	   		if err != nil {
	   			return err
	   		}
	   		return nil
	   	})
	   }
	   ```
	   ```java
	   package generated_program;

	   import com.pulumi.Context;
	   import com.pulumi.Pulumi;
	   import com.pulumi.core.Output;
	   import com.pulumi.aws.emr.Cluster;
	   import com.pulumi.aws.emr.ClusterArgs;
	   import java.util.List;
	   import java.util.ArrayList;
	   import java.util.Map;
	   import java.io.File;
	   import java.nio.file.Files;
	   import java.nio.file.Paths;

	   public class App {
	       public static void main(String[] args) {
	           Pulumi.run(App::stack);
	       }

	       public static void stack(Context ctx) {
	           var cluster = new Cluster("cluster", ClusterArgs.builder()
	               .configurationsJson("""
	   [
	   {
	   "Classification": "hadoop-env",
	   "Configurations": [
	   {
	   "Classification": "export",
	   "Properties": {
	   "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
	   }
	   }
	   ],
	   "Properties": {}
	   }
	   ]
	               """)
	               .build());

	       }
	   }
	   ```
	   ```yaml
	   resources:
	     cluster:
	       type: aws:emr:Cluster
	       properties:
	         configurationsJson: |
	           [
	           {
	           "Classification": "hadoop-env",
	           "Configurations": [
	           {
	           "Classification": "export",
	           "Properties": {
	           "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
	           }
	           }
	           ],
	           "Properties": {}
	           }
	           ]
	   ```
	*/
	ConfigurationsJson string `json:"configurationsJson,omitempty" yaml:"configurationsJson,omitempty"`

	// Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
	EbsRootVolumeSize int `json:"ebsRootVolumeSize,omitempty" yaml:"ebsRootVolumeSize,omitempty"`

	// Name of the job flow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// JSON string for selecting additional features such as adding proxy information. Note: Currently there is no API to retrieve the value of this argument after EMR cluster creation from provider, therefore the provider cannot detect drift from the actual EMR cluster if its value is changed outside the provider.
	AdditionalInfo string `json:"additionalInfo,omitempty" yaml:"additionalInfo,omitempty"`

	// An auto-termination policy for an Amazon EMR cluster. An auto-termination policy defines the amount of idle time in seconds after which a cluster automatically terminates. See Auto Termination Policy Below.
	AutoTerminationPolicy types.Emr_ClusterAutoTerminationPolicy `json:"autoTerminationPolicy,omitempty" yaml:"autoTerminationPolicy,omitempty"`

	// Switch on/off run cluster with no steps or when all steps are complete (default is on)
	KeepJobFlowAliveWhenNoSteps bool `json:"keepJobFlowAliveWhenNoSteps,omitempty" yaml:"keepJobFlowAliveWhenNoSteps,omitempty"`

	// list of tags to apply to the EMR Cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default value is `true`.
	VisibleToAllUsers bool `json:"visibleToAllUsers,omitempty" yaml:"visibleToAllUsers,omitempty"`

	// The specified placement group configuration for an Amazon EMR cluster.
	PlacementGroupConfigs []types.Emr_ClusterPlacementGroupConfig `json:"placementGroupConfigs,omitempty" yaml:"placementGroupConfigs,omitempty"`

	// Security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater.
	SecurityConfiguration string `json:"securityConfiguration,omitempty" yaml:"securityConfiguration,omitempty"`

	// List of steps to run when creating the cluster. See below. It is highly recommended to utilize the lifecycle resource options block with `ignoreChanges` if other steps are being managed outside of this provider.
	Steps []types.Emr_ClusterStep `json:"steps,omitempty" yaml:"steps,omitempty"`

	// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [core node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-core).
	CoreInstanceGroup types.Emr_ClusterCoreInstanceGroup `json:"coreInstanceGroup,omitempty" yaml:"coreInstanceGroup,omitempty"`

	// Number of steps that can be executed concurrently. You can specify a maximum of 256 steps. Only valid for EMR clusters with `release_label` 5.28.0 or greater (default is 1).
	StepConcurrencyLevel int `json:"stepConcurrencyLevel,omitempty" yaml:"stepConcurrencyLevel,omitempty"`

	// A case-insensitive list of applications for Amazon EMR to install and configure when launching the cluster. For a list of applications available for each Amazon EMR release version, see the [Amazon EMR Release Guide](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-release-components.html).
	Applications []string `json:"applications,omitempty" yaml:"applications,omitempty"`

	// IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
	AutoscalingRole string `json:"autoscalingRole,omitempty" yaml:"autoscalingRole,omitempty"`

	// S3 bucket to write the log files of the job flow. If a value is not provided, logs are not created.
	LogUri string `json:"logUri,omitempty" yaml:"logUri,omitempty"`
}
