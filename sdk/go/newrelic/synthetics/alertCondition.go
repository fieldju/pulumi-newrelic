// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synthetics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this resource to create and manage synthetics alert conditions in New Relic.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-newrelic/sdk/v3/go/newrelic/synthetics"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooMonitor, err := synthetics.LookupMonitor(ctx, &synthetics.LookupMonitorArgs{
// 			Name: "foo",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synthetics.NewAlertCondition(ctx, "fooAlertCondition", &synthetics.AlertConditionArgs{
// 			PolicyId:   pulumi.Any(newrelic_alert_policy.Foo.Id),
// 			MonitorId:  pulumi.String(fooMonitor.Id),
// 			RunbookUrl: pulumi.String("https://www.example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AlertCondition struct {
	pulumi.CustomResourceState

	// Set whether to enable the alert condition. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The ID of the Synthetics monitor to be referenced in the alert condition.
	MonitorId pulumi.StringOutput `pulumi:"monitorId"`
	// The title of this condition.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the policy where this condition should be used.
	PolicyId pulumi.IntOutput `pulumi:"policyId"`
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrOutput `pulumi:"runbookUrl"`
}

// NewAlertCondition registers a new resource with the given unique name, arguments, and options.
func NewAlertCondition(ctx *pulumi.Context,
	name string, args *AlertConditionArgs, opts ...pulumi.ResourceOption) (*AlertCondition, error) {
	if args == nil || args.MonitorId == nil {
		return nil, errors.New("missing required argument 'MonitorId'")
	}
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil {
		args = &AlertConditionArgs{}
	}
	var resource AlertCondition
	err := ctx.RegisterResource("newrelic:synthetics/alertCondition:AlertCondition", name, args, &resource, opts...)
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
	err := ctx.ReadResource("newrelic:synthetics/alertCondition:AlertCondition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertCondition resources.
type alertConditionState struct {
	// Set whether to enable the alert condition. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Synthetics monitor to be referenced in the alert condition.
	MonitorId *string `pulumi:"monitorId"`
	// The title of this condition.
	Name *string `pulumi:"name"`
	// The ID of the policy where this condition should be used.
	PolicyId *int `pulumi:"policyId"`
	// Runbook URL to display in notifications.
	RunbookUrl *string `pulumi:"runbookUrl"`
}

type AlertConditionState struct {
	// Set whether to enable the alert condition. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the Synthetics monitor to be referenced in the alert condition.
	MonitorId pulumi.StringPtrInput
	// The title of this condition.
	Name pulumi.StringPtrInput
	// The ID of the policy where this condition should be used.
	PolicyId pulumi.IntPtrInput
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrInput
}

func (AlertConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertConditionState)(nil)).Elem()
}

type alertConditionArgs struct {
	// Set whether to enable the alert condition. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Synthetics monitor to be referenced in the alert condition.
	MonitorId string `pulumi:"monitorId"`
	// The title of this condition.
	Name *string `pulumi:"name"`
	// The ID of the policy where this condition should be used.
	PolicyId int `pulumi:"policyId"`
	// Runbook URL to display in notifications.
	RunbookUrl *string `pulumi:"runbookUrl"`
}

// The set of arguments for constructing a AlertCondition resource.
type AlertConditionArgs struct {
	// Set whether to enable the alert condition. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The ID of the Synthetics monitor to be referenced in the alert condition.
	MonitorId pulumi.StringInput
	// The title of this condition.
	Name pulumi.StringPtrInput
	// The ID of the policy where this condition should be used.
	PolicyId pulumi.IntInput
	// Runbook URL to display in notifications.
	RunbookUrl pulumi.StringPtrInput
}

func (AlertConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertConditionArgs)(nil)).Elem()
}
