// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package newrelic

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this resource to create and manage Infrastructure alert conditions in New Relic.
// 
// ## Thresholds
// 
// The `critical` and `warning` threshold mapping supports the following arguments:
// 
//   * `duration` - (Required) Identifies the number of minutes the threshold must be passed or met for the alert to trigger. Threshold durations must be between 1 and 60 minutes (inclusive).
//   * `value` - (Optional) Threshold value, computed against the `comparison` operator. Supported by `infraMetric` and `infraProcessRunning` alert condition types.
//   * `timeFunction` - (Optional) Indicates if the condition needs to be sustained or to just break the threshold once; `all` or `any`. Supported by the `infraMetric` alert condition type.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/r/infra_alert_condition.html.markdown.
type InfraAlertCondition struct {
	pulumi.CustomResourceState

	// The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infraMetric` and `infraProcessRunning` condition types.
	Comparison pulumi.StringPtrOutput `pulumi:"comparison"`
	CreatedAt pulumi.IntOutput `pulumi:"createdAt"`
	// Identifies the threshold parameters for opening a critial alert violation. See Thresholds below for details.
	Critical InfraAlertConditionCriticalPtrOutput `pulumi:"critical"`
	// Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infraMetric` condition type.
	Event pulumi.StringOutput `pulumi:"event"`
	// For alerts on integrations, use this instead of `event`.  Supported by the `infraMetric` condition type.
	IntegrationProvider pulumi.StringPtrOutput `pulumi:"integrationProvider"`
	// The Infrastructure alert condition's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the alert policy where this condition should be used.
	PolicyId pulumi.IntOutput `pulumi:"policyId"`
	// Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infraProcessRunning` condition type.
	ProcessWhere pulumi.StringPtrOutput `pulumi:"processWhere"`
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrOutput `pulumi:"runbookUrl"`
	// The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infraMetric` condition type.
	Select pulumi.StringPtrOutput `pulumi:"select"`
	// The type of Infrastructure alert condition.  Valid values are  `infraProcessRunning`, `infraMetric`, and `infraHostNotReporting`.
	Type pulumi.StringOutput `pulumi:"type"`
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
	// Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.
	Warning InfraAlertConditionWarningPtrOutput `pulumi:"warning"`
	// If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.
	Where pulumi.StringPtrOutput `pulumi:"where"`
}

// NewInfraAlertCondition registers a new resource with the given unique name, arguments, and options.
func NewInfraAlertCondition(ctx *pulumi.Context,
	name string, args *InfraAlertConditionArgs, opts ...pulumi.ResourceOption) (*InfraAlertCondition, error) {
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &InfraAlertConditionArgs{}
	}
	var resource InfraAlertCondition
	err := ctx.RegisterResource("newrelic:index/infraAlertCondition:InfraAlertCondition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInfraAlertCondition gets an existing InfraAlertCondition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInfraAlertCondition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InfraAlertConditionState, opts ...pulumi.ResourceOption) (*InfraAlertCondition, error) {
	var resource InfraAlertCondition
	err := ctx.ReadResource("newrelic:index/infraAlertCondition:InfraAlertCondition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InfraAlertCondition resources.
type infraAlertConditionState struct {
	// The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infraMetric` and `infraProcessRunning` condition types.
	Comparison *string `pulumi:"comparison"`
	CreatedAt *int `pulumi:"createdAt"`
	// Identifies the threshold parameters for opening a critial alert violation. See Thresholds below for details.
	Critical *InfraAlertConditionCritical `pulumi:"critical"`
	// Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infraMetric` condition type.
	Event *string `pulumi:"event"`
	// For alerts on integrations, use this instead of `event`.  Supported by the `infraMetric` condition type.
	IntegrationProvider *string `pulumi:"integrationProvider"`
	// The Infrastructure alert condition's name.
	Name *string `pulumi:"name"`
	// The ID of the alert policy where this condition should be used.
	PolicyId *int `pulumi:"policyId"`
	// Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infraProcessRunning` condition type.
	ProcessWhere *string `pulumi:"processWhere"`
	// Runbook URL to display in notifications.
	RunbookUrl *string `pulumi:"runbookUrl"`
	// The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infraMetric` condition type.
	Select *string `pulumi:"select"`
	// The type of Infrastructure alert condition.  Valid values are  `infraProcessRunning`, `infraMetric`, and `infraHostNotReporting`.
	Type *string `pulumi:"type"`
	UpdatedAt *int `pulumi:"updatedAt"`
	// Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.
	Warning *InfraAlertConditionWarning `pulumi:"warning"`
	// If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.
	Where *string `pulumi:"where"`
}

type InfraAlertConditionState struct {
	// The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infraMetric` and `infraProcessRunning` condition types.
	Comparison pulumi.StringPtrInput
	CreatedAt pulumi.IntPtrInput
	// Identifies the threshold parameters for opening a critial alert violation. See Thresholds below for details.
	Critical InfraAlertConditionCriticalPtrInput
	// Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infraMetric` condition type.
	Event pulumi.StringPtrInput
	// For alerts on integrations, use this instead of `event`.  Supported by the `infraMetric` condition type.
	IntegrationProvider pulumi.StringPtrInput
	// The Infrastructure alert condition's name.
	Name pulumi.StringPtrInput
	// The ID of the alert policy where this condition should be used.
	PolicyId pulumi.IntPtrInput
	// Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infraProcessRunning` condition type.
	ProcessWhere pulumi.StringPtrInput
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrInput
	// The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infraMetric` condition type.
	Select pulumi.StringPtrInput
	// The type of Infrastructure alert condition.  Valid values are  `infraProcessRunning`, `infraMetric`, and `infraHostNotReporting`.
	Type pulumi.StringPtrInput
	UpdatedAt pulumi.IntPtrInput
	// Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.
	Warning InfraAlertConditionWarningPtrInput
	// If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.
	Where pulumi.StringPtrInput
}

func (InfraAlertConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*infraAlertConditionState)(nil)).Elem()
}

type infraAlertConditionArgs struct {
	// The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infraMetric` and `infraProcessRunning` condition types.
	Comparison *string `pulumi:"comparison"`
	// Identifies the threshold parameters for opening a critial alert violation. See Thresholds below for details.
	Critical *InfraAlertConditionCritical `pulumi:"critical"`
	// Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infraMetric` condition type.
	Event *string `pulumi:"event"`
	// For alerts on integrations, use this instead of `event`.  Supported by the `infraMetric` condition type.
	IntegrationProvider *string `pulumi:"integrationProvider"`
	// The Infrastructure alert condition's name.
	Name *string `pulumi:"name"`
	// The ID of the alert policy where this condition should be used.
	PolicyId int `pulumi:"policyId"`
	// Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infraProcessRunning` condition type.
	ProcessWhere *string `pulumi:"processWhere"`
	// Runbook URL to display in notifications.
	RunbookUrl *string `pulumi:"runbookUrl"`
	// The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infraMetric` condition type.
	Select *string `pulumi:"select"`
	// The type of Infrastructure alert condition.  Valid values are  `infraProcessRunning`, `infraMetric`, and `infraHostNotReporting`.
	Type string `pulumi:"type"`
	// Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.
	Warning *InfraAlertConditionWarning `pulumi:"warning"`
	// If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.
	Where *string `pulumi:"where"`
}

// The set of arguments for constructing a InfraAlertCondition resource.
type InfraAlertConditionArgs struct {
	// The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infraMetric` and `infraProcessRunning` condition types.
	Comparison pulumi.StringPtrInput
	// Identifies the threshold parameters for opening a critial alert violation. See Thresholds below for details.
	Critical InfraAlertConditionCriticalPtrInput
	// Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infraMetric` condition type.
	Event pulumi.StringPtrInput
	// For alerts on integrations, use this instead of `event`.  Supported by the `infraMetric` condition type.
	IntegrationProvider pulumi.StringPtrInput
	// The Infrastructure alert condition's name.
	Name pulumi.StringPtrInput
	// The ID of the alert policy where this condition should be used.
	PolicyId pulumi.IntInput
	// Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infraProcessRunning` condition type.
	ProcessWhere pulumi.StringPtrInput
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrInput
	// The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infraMetric` condition type.
	Select pulumi.StringPtrInput
	// The type of Infrastructure alert condition.  Valid values are  `infraProcessRunning`, `infraMetric`, and `infraHostNotReporting`.
	Type pulumi.StringInput
	// Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.
	Warning InfraAlertConditionWarningPtrInput
	// If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.
	Where pulumi.StringPtrInput
}

func (InfraAlertConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*infraAlertConditionArgs)(nil)).Elem()
}

