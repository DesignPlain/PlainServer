package types

type Appflow_FlowTriggerConfigTriggerPropertiesScheduled struct {
	// Scheduling expression that determines the rate at which the schedule will run, for example `rate(5minutes)`.
	ScheduleExpression string `json:"scheduleExpression,omitempty" yaml:"scheduleExpression,omitempty"`

	// Optional offset that is added to the time interval for a schedule-triggered flow. Maximum value of 36000.
	ScheduleOffset int `json:"scheduleOffset,omitempty" yaml:"scheduleOffset,omitempty"`

	// Scheduled start time for a schedule-triggered flow. Must be a valid RFC3339 timestamp.
	ScheduleStartTime string `json:"scheduleStartTime,omitempty" yaml:"scheduleStartTime,omitempty"`

	/*
	   Time zone used when referring to the date and time of a scheduled-triggered flow, such as `America/New_York`.

	   <!--Start PulumiCodeChooser -->
	   ```java
	   package generated_program;

	   import com.pulumi.Context;
	   import com.pulumi.Pulumi;
	   import com.pulumi.core.Output;
	   import com.pulumi.aws.appflow.Flow;
	   import com.pulumi.aws.appflow.FlowArgs;
	   import com.pulumi.aws.appflow.inputs.FlowTriggerConfigArgs;
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
	           var example = new Flow("example", FlowArgs.builder()
	               .triggerConfig(FlowTriggerConfigArgs.builder()
	                   .scheduled(%!!(MISSING)v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
	                   .build())
	               .build());

	       }
	   }
	   ```
	   ```yaml
	   resources:
	     example:
	       type: aws:appflow:Flow
	       properties:
	         triggerConfig:
	           scheduled:
	             - scheduleExpression: rate(1minutes)
	   ```
	   <!--End PulumiCodeChooser -->
	*/
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`

	// Whether a scheduled flow has an incremental data transfer or a complete data transfer for each flow run. Valid values are `Incremental` and `Complete`.
	DataPullMode string `json:"dataPullMode,omitempty" yaml:"dataPullMode,omitempty"`

	// Date range for the records to import from the connector in the first flow run. Must be a valid RFC3339 timestamp.
	FirstExecutionFrom string `json:"firstExecutionFrom,omitempty" yaml:"firstExecutionFrom,omitempty"`

	// Scheduled end time for a schedule-triggered flow. Must be a valid RFC3339 timestamp.
	ScheduleEndTime string `json:"scheduleEndTime,omitempty" yaml:"scheduleEndTime,omitempty"`
}
