package types

type Ebs_getEbsVolumesFilter struct {
	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVolumes.html).
	   For example, if matching against the `size` filter, use:

	   <!--Start PulumiCodeChooser -->
	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const tenOrTwentyGbVolumes = aws.ebs.getEbsVolumes({
	       filters: [{
	           name: "size",
	           values: [
	               "10",
	               "20",
	           ],
	       }],
	   });
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   ten_or_twenty_gb_volumes = aws.ebs.get_ebs_volumes(filters=[{
	       "name": "size",
	       "values": [
	           "10",
	           "20",
	       ],
	   }])
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var tenOrTwentyGbVolumes = Aws.Ebs.GetEbsVolumes.Invoke(new()
	       {
	           Filters = new[]
	           {
	               new Aws.Ebs.Inputs.GetEbsVolumesFilterInputArgs
	               {
	                   Name = "size",
	                   Values = new[]
	                   {
	                       "10",
	                       "20",
	                   },
	               },
	           },
	       });

	   });
	   ```
	   ```go
	   package main

	   import (
	   	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	   	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	   )

	   func main() {
	   	pulumi.Run(func(ctx -pulumi.Context) error {
	   		_, err := ebs.GetEbsVolumes(ctx, &ebs.GetEbsVolumesArgs{
	   			Filters: []ebs.GetEbsVolumesFilter{
	   				{
	   					Name: "size",
	   					Values: []string{
	   						"10",
	   						"20",
	   					},
	   				},
	   			},
	   		}, nil)
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
	   import com.pulumi.aws.ebs.EbsFunctions;
	   import com.pulumi.aws.ebs.inputs.GetEbsVolumesArgs;
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
	           final var tenOrTwentyGbVolumes = EbsFunctions.getEbsVolumes(GetEbsVolumesArgs.builder()
	               .filters(GetEbsVolumesFilterArgs.builder()
	                   .name("size")
	                   .values(
	                       "10",
	                       "20")
	                   .build())
	               .build());

	       }
	   }
	   ```
	   ```yaml
	   variables:
	     tenOrTwentyGbVolumes:
	       fn::invoke:
	         Function: aws:ebs:getEbsVolumes
	         Arguments:
	           filters:
	             - name: size
	               values:
	                 - '10'
	                 - '20'
	   ```
	   <!--End PulumiCodeChooser -->
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Set of values that are accepted for the given field.
	   EBS Volume IDs will be selected if any one of the given values match.
	*/
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
