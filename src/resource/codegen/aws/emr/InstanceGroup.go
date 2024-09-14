package emr

import types "libds/aws/types"

type InstanceGroup struct {
	// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Human friendly name given to the instance group. Changing this forces a new resource to be created.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice string `json:"bidPrice,omitempty" yaml:"bidPrice,omitempty"`

	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
	EbsOptimized bool `json:"ebsOptimized,omitempty" yaml:"ebsOptimized,omitempty"`

	// target number of instances for the instance group. defaults to 0.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
	AutoscalingPolicy string `json:"autoscalingPolicy,omitempty" yaml:"autoscalingPolicy,omitempty"`

	/*
	   A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.

	   <!--Start PulumiCodeChooser -->
	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const task = new aws.emr.InstanceGroup("task", {configurationsJson: `[
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

	   task = aws.emr.InstanceGroup("task", configurations_json="""[
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
	       var task = new Aws.Emr.InstanceGroup("task", new()
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
	   		_, err := emr.NewInstanceGroup(ctx, "task", &emr.InstanceGroupArgs{
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
	   import com.pulumi.aws.emr.InstanceGroup;
	   import com.pulumi.aws.emr.InstanceGroupArgs;
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
	           var task = new InstanceGroup("task", InstanceGroupArgs.builder()
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
	     task:
	       type: aws:emr:InstanceGroup
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
	   <!--End PulumiCodeChooser -->
	*/
	ConfigurationsJson string `json:"configurationsJson,omitempty" yaml:"configurationsJson,omitempty"`

	// One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
	EbsConfigs []types.Emr_InstanceGroupEbsConfig `json:"ebsConfigs,omitempty" yaml:"ebsConfigs,omitempty"`
}
