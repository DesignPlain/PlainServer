package types

type Appautoscaling_PolicyStepScalingPolicyConfiguration struct {
	// Minimum number to adjust your scalable dimension as a result of a scaling activity. If the adjustment type is PercentChangeInCapacity, the scaling policy changes the scalable dimension of the scalable target by this amount.
	MinAdjustmentMagnitude int `json:"minAdjustmentMagnitude,omitempty" yaml:"minAdjustmentMagnitude,omitempty"`

	/*
	   Set of adjustments that manage scaling. These have the following structure:

	   <!--Start PulumiCodeChooser -->
	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const ecsPolicy = new aws.appautoscaling.Policy("ecs_policy", {stepScalingPolicyConfiguration: {
	       stepAdjustments: [
	           {
	               metricIntervalLowerBound: "1",
	               metricIntervalUpperBound: "2",
	               scalingAdjustment: -1,
	           },
	           {
	               metricIntervalLowerBound: "2",
	               metricIntervalUpperBound: "3",
	               scalingAdjustment: 1,
	           },
	       ],
	   }});
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   ecs_policy = aws.appautoscaling.Policy("ecs_policy", step_scaling_policy_configuration={
	       "step_adjustments": [
	           {
	               "metric_interval_lower_bound": "1",
	               "metric_interval_upper_bound": "2",
	               "scaling_adjustment": -1,
	           },
	           {
	               "metric_interval_lower_bound": "2",
	               "metric_interval_upper_bound": "3",
	               "scaling_adjustment": 1,
	           },
	       ],
	   })
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var ecsPolicy = new Aws.AppAutoScaling.Policy("ecs_policy", new()
	       {
	           StepScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyStepScalingPolicyConfigurationArgs
	           {
	               StepAdjustments = new[]
	               {
	                   new Aws.AppAutoScaling.Inputs.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs
	                   {
	                       MetricIntervalLowerBound = "1",
	                       MetricIntervalUpperBound = "2",
	                       ScalingAdjustment = -1,
	                   },
	                   new Aws.AppAutoScaling.Inputs.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs
	                   {
	                       MetricIntervalLowerBound = "2",
	                       MetricIntervalUpperBound = "3",
	                       ScalingAdjustment = 1,
	                   },
	               },
	           },
	       });

	   });
	   ```
	   ```go
	   package main

	   import (
	   	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	   	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	   )

	   func main() {
	   	pulumi.Run(func(ctx -pulumi.Context) error {
	   		_, err := appautoscaling.NewPolicy(ctx, "ecs_policy", &appautoscaling.PolicyArgs{
	   			StepScalingPolicyConfiguration: &appautoscaling.PolicyStepScalingPolicyConfigurationArgs{
	   				StepAdjustments: appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArray{
	   					&appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs{
	   						MetricIntervalLowerBound: pulumi.String("1"),
	   						MetricIntervalUpperBound: pulumi.String("2"),
	   						ScalingAdjustment:        int(-1),
	   					},
	   					&appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs{
	   						MetricIntervalLowerBound: pulumi.String("2"),
	   						MetricIntervalUpperBound: pulumi.String("3"),
	   						ScalingAdjustment:        pulumi.Int(1),
	   					},
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
	     ecsPolicy:
	       type: aws:appautoscaling:Policy
	       name: ecs_policy
	       properties:
	         stepScalingPolicyConfiguration:
	           stepAdjustments:
	             - metricIntervalLowerBound: 1
	               metricIntervalUpperBound: 2
	               scalingAdjustment: -1
	             - metricIntervalLowerBound: 2
	               metricIntervalUpperBound: 3
	               scalingAdjustment: 1
	   ```
	   <!--End PulumiCodeChooser -->
	*/
	StepAdjustments []Appautoscaling_PolicyStepScalingPolicyConfigurationStepAdjustment `json:"stepAdjustments,omitempty" yaml:"stepAdjustments,omitempty"`

	// Whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType string `json:"adjustmentType,omitempty" yaml:"adjustmentType,omitempty"`

	// Amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown int `json:"cooldown,omitempty" yaml:"cooldown,omitempty"`

	// Aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType string `json:"metricAggregationType,omitempty" yaml:"metricAggregationType,omitempty"`
}
