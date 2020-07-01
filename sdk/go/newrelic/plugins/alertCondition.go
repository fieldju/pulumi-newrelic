// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package plugins

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this resource to create and manage plugins alert conditions in New Relic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-newrelic/sdk/v3/go/newrelic"
// 	"github.com/pulumi/pulumi-newrelic/sdk/v3/go/newrelic/plugins"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooPlugin, err := plugins.GetPlugin(ctx, &plugins.GetPluginArgs{
// 			Guid: "com.example.my-plugin",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooPluginComponent, err := plugins.GetPluginComponent(ctx, &plugins.GetPluginComponentArgs{
// 			PluginId: fooPlugin.Id,
// 			Name:     "MyPlugin",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooAlertPolicy, err := newrelic.NewAlertPolicy(ctx, "fooAlertPolicy", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = plugins.NewAlertCondition(ctx, "fooAlertCondition", &plugins.AlertConditionArgs{
// 			PolicyId: fooAlertPolicy.ID(),
// 			Entities: pulumi.IntArray{
// 				pulumi.String(fooPluginComponent.Id),
// 			},
// 			Metric:            pulumi.String("Component/Summary/Consumers[consumers]"),
// 			PluginId:          pulumi.String(fooPlugin.Id),
// 			PluginGuid:        pulumi.String(fooPlugin.Guid),
// 			ValueFunction:     pulumi.String("average"),
// 			MetricDescription: pulumi.String("Queue consumers"),
// 			Terms: plugins.AlertConditionTermArray{
// 				&plugins.AlertConditionTermArgs{
// 					Duration:     pulumi.Int(5),
// 					Operator:     pulumi.String("below"),
// 					Priority:     pulumi.String("critical"),
// 					Threshold:    pulumi.Float64(0.75),
// 					TimeFunction: pulumi.String("all"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Terms
//
// The `term` mapping supports the following arguments:
//
//   * `duration` - (Required) In minutes, must be in the range of `5` to `120`, inclusive.
//   * `operator` - (Optional) `above`, `below`, or `equal`.  Defaults to `equal`.
//   * `priority` - (Optional) `critical` or `warning`.  Defaults to `critical`.
//   * `threshold` - (Required) Must be 0 or greater.
//   * `timeFunction` - (Required) `all` or `any`.
type AlertCondition struct {
	pulumi.CustomResourceState

	// Whether or not this condition is enabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The plugin component IDs to target.
	Entities pulumi.IntArrayOutput `pulumi:"entities"`
	// The plugin metric to evaluate.
	Metric pulumi.StringOutput `pulumi:"metric"`
	// The metric description.
	MetricDescription pulumi.StringOutput `pulumi:"metricDescription"`
	// The title of the condition. Must be between 1 and 64 characters, inclusive.
	Name pulumi.StringOutput `pulumi:"name"`
	// The GUID of the plugin which produces the metric.
	PluginGuid pulumi.StringOutput `pulumi:"pluginGuid"`
	// The ID of the installed plugin instance which produces the metric.
	PluginId pulumi.StringOutput `pulumi:"pluginId"`
	// The ID of the policy where this condition should be used.
	PolicyId pulumi.IntOutput `pulumi:"policyId"`
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrOutput `pulumi:"runbookUrl"`
	// A list of terms for this condition. See Terms below for details.
	Terms AlertConditionTermArrayOutput `pulumi:"terms"`
	// The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
	ValueFunction pulumi.StringOutput `pulumi:"valueFunction"`
}

// NewAlertCondition registers a new resource with the given unique name, arguments, and options.
func NewAlertCondition(ctx *pulumi.Context,
	name string, args *AlertConditionArgs, opts ...pulumi.ResourceOption) (*AlertCondition, error) {
	if args == nil || args.Entities == nil {
		return nil, errors.New("missing required argument 'Entities'")
	}
	if args == nil || args.Metric == nil {
		return nil, errors.New("missing required argument 'Metric'")
	}
	if args == nil || args.MetricDescription == nil {
		return nil, errors.New("missing required argument 'MetricDescription'")
	}
	if args == nil || args.PluginGuid == nil {
		return nil, errors.New("missing required argument 'PluginGuid'")
	}
	if args == nil || args.PluginId == nil {
		return nil, errors.New("missing required argument 'PluginId'")
	}
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil || args.Terms == nil {
		return nil, errors.New("missing required argument 'Terms'")
	}
	if args == nil || args.ValueFunction == nil {
		return nil, errors.New("missing required argument 'ValueFunction'")
	}
	if args == nil {
		args = &AlertConditionArgs{}
	}
	var resource AlertCondition
	err := ctx.RegisterResource("newrelic:plugins/alertCondition:AlertCondition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertCondition gets an existing AlertCondition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertCondition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertConditionState, opts ...pulumi.ResourceOption) (*AlertCondition, error) {
	var resource AlertCondition
	err := ctx.ReadResource("newrelic:plugins/alertCondition:AlertCondition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertCondition resources.
type alertConditionState struct {
	// Whether or not this condition is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The plugin component IDs to target.
	Entities []int `pulumi:"entities"`
	// The plugin metric to evaluate.
	Metric *string `pulumi:"metric"`
	// The metric description.
	MetricDescription *string `pulumi:"metricDescription"`
	// The title of the condition. Must be between 1 and 64 characters, inclusive.
	Name *string `pulumi:"name"`
	// The GUID of the plugin which produces the metric.
	PluginGuid *string `pulumi:"pluginGuid"`
	// The ID of the installed plugin instance which produces the metric.
	PluginId *string `pulumi:"pluginId"`
	// The ID of the policy where this condition should be used.
	PolicyId *int `pulumi:"policyId"`
	// Runbook URL to display in notifications.
	RunbookUrl *string `pulumi:"runbookUrl"`
	// A list of terms for this condition. See Terms below for details.
	Terms []AlertConditionTerm `pulumi:"terms"`
	// The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
	ValueFunction *string `pulumi:"valueFunction"`
}

type AlertConditionState struct {
	// Whether or not this condition is enabled.
	Enabled pulumi.BoolPtrInput
	// The plugin component IDs to target.
	Entities pulumi.IntArrayInput
	// The plugin metric to evaluate.
	Metric pulumi.StringPtrInput
	// The metric description.
	MetricDescription pulumi.StringPtrInput
	// The title of the condition. Must be between 1 and 64 characters, inclusive.
	Name pulumi.StringPtrInput
	// The GUID of the plugin which produces the metric.
	PluginGuid pulumi.StringPtrInput
	// The ID of the installed plugin instance which produces the metric.
	PluginId pulumi.StringPtrInput
	// The ID of the policy where this condition should be used.
	PolicyId pulumi.IntPtrInput
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrInput
	// A list of terms for this condition. See Terms below for details.
	Terms AlertConditionTermArrayInput
	// The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
	ValueFunction pulumi.StringPtrInput
}

func (AlertConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertConditionState)(nil)).Elem()
}

type alertConditionArgs struct {
	// Whether or not this condition is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The plugin component IDs to target.
	Entities []int `pulumi:"entities"`
	// The plugin metric to evaluate.
	Metric string `pulumi:"metric"`
	// The metric description.
	MetricDescription string `pulumi:"metricDescription"`
	// The title of the condition. Must be between 1 and 64 characters, inclusive.
	Name *string `pulumi:"name"`
	// The GUID of the plugin which produces the metric.
	PluginGuid string `pulumi:"pluginGuid"`
	// The ID of the installed plugin instance which produces the metric.
	PluginId string `pulumi:"pluginId"`
	// The ID of the policy where this condition should be used.
	PolicyId int `pulumi:"policyId"`
	// Runbook URL to display in notifications.
	RunbookUrl *string `pulumi:"runbookUrl"`
	// A list of terms for this condition. See Terms below for details.
	Terms []AlertConditionTerm `pulumi:"terms"`
	// The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
	ValueFunction string `pulumi:"valueFunction"`
}

// The set of arguments for constructing a AlertCondition resource.
type AlertConditionArgs struct {
	// Whether or not this condition is enabled.
	Enabled pulumi.BoolPtrInput
	// The plugin component IDs to target.
	Entities pulumi.IntArrayInput
	// The plugin metric to evaluate.
	Metric pulumi.StringInput
	// The metric description.
	MetricDescription pulumi.StringInput
	// The title of the condition. Must be between 1 and 64 characters, inclusive.
	Name pulumi.StringPtrInput
	// The GUID of the plugin which produces the metric.
	PluginGuid pulumi.StringInput
	// The ID of the installed plugin instance which produces the metric.
	PluginId pulumi.StringInput
	// The ID of the policy where this condition should be used.
	PolicyId pulumi.IntInput
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrInput
	// A list of terms for this condition. See Terms below for details.
	Terms AlertConditionTermArrayInput
	// The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
	ValueFunction pulumi.StringInput
}

func (AlertConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertConditionArgs)(nil)).Elem()
}
