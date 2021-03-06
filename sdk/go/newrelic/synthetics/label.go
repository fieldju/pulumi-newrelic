// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synthetics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this resource to create, update, and delete a Synthetics label in New Relic.
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
// 		_, err := synthetics.NewLabel(ctx, "foo", &synthetics.LabelArgs{
// 			MonitorId: pulumi.Any(newrelic_synthetics_monitor.Foo.Id),
// 			Type:      pulumi.String("MyCategory"),
// 			Value:     pulumi.String("MyValue"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Label struct {
	pulumi.CustomResourceState

	// The URL of the Synthetics label.
	Href pulumi.StringOutput `pulumi:"href"`
	// The ID of the monitor that will be assigned the label.
	MonitorId pulumi.StringOutput `pulumi:"monitorId"`
	// A string representing the label key/category.
	Type pulumi.StringOutput `pulumi:"type"`
	// A string representing the label value.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewLabel registers a new resource with the given unique name, arguments, and options.
func NewLabel(ctx *pulumi.Context,
	name string, args *LabelArgs, opts ...pulumi.ResourceOption) (*Label, error) {
	if args == nil || args.MonitorId == nil {
		return nil, errors.New("missing required argument 'MonitorId'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &LabelArgs{}
	}
	var resource Label
	err := ctx.RegisterResource("newrelic:synthetics/label:Label", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLabel gets an existing Label resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelState, opts ...pulumi.ResourceOption) (*Label, error) {
	var resource Label
	err := ctx.ReadResource("newrelic:synthetics/label:Label", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Label resources.
type labelState struct {
	// The URL of the Synthetics label.
	Href *string `pulumi:"href"`
	// The ID of the monitor that will be assigned the label.
	MonitorId *string `pulumi:"monitorId"`
	// A string representing the label key/category.
	Type *string `pulumi:"type"`
	// A string representing the label value.
	Value *string `pulumi:"value"`
}

type LabelState struct {
	// The URL of the Synthetics label.
	Href pulumi.StringPtrInput
	// The ID of the monitor that will be assigned the label.
	MonitorId pulumi.StringPtrInput
	// A string representing the label key/category.
	Type pulumi.StringPtrInput
	// A string representing the label value.
	Value pulumi.StringPtrInput
}

func (LabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*labelState)(nil)).Elem()
}

type labelArgs struct {
	// The ID of the monitor that will be assigned the label.
	MonitorId string `pulumi:"monitorId"`
	// A string representing the label key/category.
	Type string `pulumi:"type"`
	// A string representing the label value.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Label resource.
type LabelArgs struct {
	// The ID of the monitor that will be assigned the label.
	MonitorId pulumi.StringInput
	// A string representing the label key/category.
	Type pulumi.StringInput
	// A string representing the label value.
	Value pulumi.StringInput
}

func (LabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labelArgs)(nil)).Elem()
}
