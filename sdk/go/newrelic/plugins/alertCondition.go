// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package plugins

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this resource to create and manage plugins alert conditions in New Relic.
// 
// ## Terms
// 
// The `term` mapping supports the following arguments:
// 
//   * `duration` - (Required) In minutes, must be in the range of `5` to `120`, inclusive.
//   * `operator` - (Optional) `above`, `below`, or `equal`.  Defaults to `equal`.
//   * `priority` - (Optional) `critical` or `warning`.  Defaults to `critical`.
//   * `threshold` - (Required) Must be 0 or greater.
//   * `timeFunction` - (Required) `all` or `any`.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/r/plugins_alert_condition.html.markdown.
type AlertCondition struct {
	pulumi.CustomResourceState

	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	Entities pulumi.IntArrayOutput `pulumi:"entities"`
	// The metric field accepts parameters based on the `type` set.
	Metric pulumi.StringOutput `pulumi:"metric"`
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
	Enabled *bool `pulumi:"enabled"`
	Entities []int `pulumi:"entities"`
	// The metric field accepts parameters based on the `type` set.
	Metric *string `pulumi:"metric"`
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
	ValueFunction *string `pulumi:"valueFunction"`
}

type AlertConditionState struct {
	Enabled pulumi.BoolPtrInput
	Entities pulumi.IntArrayInput
	// The metric field accepts parameters based on the `type` set.
	Metric pulumi.StringPtrInput
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
	ValueFunction pulumi.StringPtrInput
}

func (AlertConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertConditionState)(nil)).Elem()
}

type alertConditionArgs struct {
	Enabled *bool `pulumi:"enabled"`
	Entities []int `pulumi:"entities"`
	// The metric field accepts parameters based on the `type` set.
	Metric string `pulumi:"metric"`
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
	ValueFunction string `pulumi:"valueFunction"`
}

// The set of arguments for constructing a AlertCondition resource.
type AlertConditionArgs struct {
	Enabled pulumi.BoolPtrInput
	Entities pulumi.IntArrayInput
	// The metric field accepts parameters based on the `type` set.
	Metric pulumi.StringInput
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
	ValueFunction pulumi.StringInput
}

func (AlertConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertConditionArgs)(nil)).Elem()
}

