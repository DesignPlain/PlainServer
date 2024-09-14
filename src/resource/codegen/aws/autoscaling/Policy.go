package autoscaling

import types "libds/aws/types"

type Policy struct {
	// Estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup int `json:"estimatedInstanceWarmup,omitempty" yaml:"estimatedInstanceWarmup,omitempty"`

	// Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType string `json:"metricAggregationType,omitempty" yaml:"metricAggregationType,omitempty"`

	// Name of the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Policy type, either "SimpleScaling", "StepScaling", "TargetTrackingScaling", or "PredictiveScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType string `json:"policyType,omitempty" yaml:"policyType,omitempty"`

	/*
	   Number of members by which to
	   scale, when the adjustment bounds are breached. A positive value scales
	   up. A negative value scales down.
	*/
	ScalingAdjustment int `json:"scalingAdjustment,omitempty" yaml:"scalingAdjustment,omitempty"`

	// Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown int `json:"cooldown,omitempty" yaml:"cooldown,omitempty"`

	/*
	   Whether the scaling policy is enabled or disabled. Default: `true`.

	   The following argument is only available to "SimpleScaling" and "StepScaling" type policies:
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   Minimum value to scale by when `adjustment_type` is set to `PercentChangeInCapacity`.

	   The following arguments are only available to "SimpleScaling" type policies:
	*/
	MinAdjustmentMagnitude int `json:"minAdjustmentMagnitude,omitempty" yaml:"minAdjustmentMagnitude,omitempty"`

	// Predictive scaling policy configuration to use with Amazon EC2 Auto Scaling.
	PredictiveScalingConfiguration types.Autoscaling_PolicyPredictiveScalingConfiguration `json:"predictiveScalingConfiguration,omitempty" yaml:"predictiveScalingConfiguration,omitempty"`

	/*
	   Set of adjustments that manage
	   group scaling. These have the following structure:

	   <!--Start PulumiCodeChooser -->
	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const example = new aws.autoscaling.Policy("example", {stepAdjustments: [
	       {
	           scalingAdjustment: -1,
	           metricIntervalLowerBound: "1",
	           metricIntervalUpperBound: "2",
	       },
	       {
	           scalingAdjustment: 1,
	           metricIntervalLowerBound: "2",
	           metricIntervalUpperBound: "3",
	       },
	   ]});
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   example = aws.autoscaling.Policy("example", step_adjustments=[
	       {
	           "scaling_adjustment": -1,
	           "metric_interval_lower_bound": "1",
	           "metric_interval_upper_bound": "2",
	       },
	       {
	           "scaling_adjustment": 1,
	           "metric_interval_lower_bound": "2",
	           "metric_interval_upper_bound": "3",
	       },
	   ])
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var example = new Aws.AutoScaling.Policy("example", new()
	       {
	           StepAdjustments = new[]
	           {
	               new Aws.AutoScaling.Inputs.PolicyStepAdjustmentArgs
	               {
	                   ScalingAdjustment = -1,
	                   MetricIntervalLowerBound = "1",
	                   MetricIntervalUpperBound = "2",
	               },
	               new Aws.AutoScaling.Inputs.PolicyStepAdjustmentArgs
	               {
	                   ScalingAdjustment = 1,
	                   MetricIntervalLowerBound = "2",
	                   MetricIntervalUpperBound = "3",
	               },
	           },
	       });

	   });
	   ```
	   ```go
	   package main

	   import (
	   	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	   	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	   )

	   func main() {
	   	pulumi.Run(func(ctx -pulumi.Context) error {
	   		_, err := autoscaling.NewPolicy(ctx, "example", &autoscaling.PolicyArgs{
	   			StepAdjustments: autoscaling.PolicyStepAdjustmentArray{
	   				&autoscaling.PolicyStepAdjustmentArgs{
	   					ScalingAdjustment:        int(-1),
	   					MetricIntervalLowerBound: pulumi.String("1"),
	   					MetricIntervalUpperBound: pulumi.String("2"),
	   				},
	   				&autoscaling.PolicyStepAdjustmentArgs{
	   					ScalingAdjustment:        pulumi.Int(1),
	   					MetricIntervalLowerBound: pulumi.String("2"),
	   					MetricIntervalUpperBound: pulumi.String("3"),
	   				},
	   			},
	   		})
	   		if err != nil {
	   			return err
	   		}
	   		return nil
	   	})
	   }
	   ```
	   ```yaml
	   resources:
	     example:
	       type: aws:autoscaling:Policy
	       properties:
	         stepAdjustments:
	           - scalingAdjustment: -1
	             metricIntervalLowerBound: 1
	             metricIntervalUpperBound: 2
	           - scalingAdjustment: 1
	             metricIntervalLowerBound: 2
	             metricIntervalUpperBound: 3
	   ```
	   <!--End PulumiCodeChooser -->

	   The following fields are available in step adjustments:
	*/
	StepAdjustments []types.Autoscaling_PolicyStepAdjustment `json:"stepAdjustments,omitempty" yaml:"stepAdjustments,omitempty"`

	/*
	   Target tracking policy. These have the following structure:

	   <!--Start PulumiCodeChooser -->
	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const example = new aws.autoscaling.Policy("example", {targetTrackingConfiguration: {
	       predefinedMetricSpecification: {
	           predefinedMetricType: "ASGAverageCPUUtilization",
	       },
	       targetValue: 40,
	   }});
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   example = aws.autoscaling.Policy("example", target_tracking_configuration={
	       "predefined_metric_specification": {
	           "predefined_metric_type": "ASGAverageCPUUtilization",
	       },
	       "target_value": 40,
	   })
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var example = new Aws.AutoScaling.Policy("example", new()
	       {
	           TargetTrackingConfiguration = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationArgs
	           {
	               PredefinedMetricSpecification = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs
	               {
	                   PredefinedMetricType = "ASGAverageCPUUtilization",
	               },
	               TargetValue = 40,
	           },
	       });

	   });
	   ```
	   ```go
	   package main

	   import (
	   	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	   	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	   )

	   func main() {
	   	pulumi.Run(func(ctx -pulumi.Context) error {
	   		_, err := autoscaling.NewPolicy(ctx, "example", &autoscaling.PolicyArgs{
	   			TargetTrackingConfiguration: &autoscaling.PolicyTargetTrackingConfigurationArgs{
	   				PredefinedMetricSpecification: &autoscaling.PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs{
	   					PredefinedMetricType: pulumi.String("ASGAverageCPUUtilization"),
	   				},
	   				TargetValue: pulumi.Float64(40),
	   			},
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
	   import com.pulumi.aws.autoscaling.Policy;
	   import com.pulumi.aws.autoscaling.PolicyArgs;
	   import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationArgs;
	   import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs;
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
	           var example = new Policy("example", PolicyArgs.builder()
	               .targetTrackingConfiguration(PolicyTargetTrackingConfigurationArgs.builder()
	                   .predefinedMetricSpecification(PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs.builder()
	                       .predefinedMetricType("ASGAverageCPUUtilization")
	                       .build())
	                   .targetValue(40)
	                   .build())
	               .build());

	       }
	   }
	   ```
	   ```yaml
	   resources:
	     example:
	       type: aws:autoscaling:Policy
	       properties:
	         targetTrackingConfiguration:
	           predefinedMetricSpecification:
	             predefinedMetricType: ASGAverageCPUUtilization
	           targetValue: 40
	   ```
	   <!--End PulumiCodeChooser -->

	   The following fields are available in target tracking configuration:
	*/
	TargetTrackingConfiguration types.Autoscaling_PolicyTargetTrackingConfiguration `json:"targetTrackingConfiguration,omitempty" yaml:"targetTrackingConfiguration,omitempty"`

	// Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType string `json:"adjustmentType,omitempty" yaml:"adjustmentType,omitempty"`

	// Name of the autoscaling group.
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty" yaml:"autoscalingGroupName,omitempty"`
}
