// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package newrelic

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this resource to map alert policies to alert channels in New Relic.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/r/alert_policy_channel.html.markdown.
type AlertPolicyChannel struct {
	pulumi.CustomResourceState

	// The ID of the channel.
	ChannelId pulumi.IntOutput `pulumi:"channelId"`
	// The ID of the policy.
	PolicyId pulumi.IntOutput `pulumi:"policyId"`
}

// NewAlertPolicyChannel registers a new resource with the given unique name, arguments, and options.
func NewAlertPolicyChannel(ctx *pulumi.Context,
	name string, args *AlertPolicyChannelArgs, opts ...pulumi.ResourceOption) (*AlertPolicyChannel, error) {
	if args == nil || args.ChannelId == nil {
		return nil, errors.New("missing required argument 'ChannelId'")
	}
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil {
		args = &AlertPolicyChannelArgs{}
	}
	var resource AlertPolicyChannel
	err := ctx.RegisterResource("newrelic:index/alertPolicyChannel:AlertPolicyChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertPolicyChannel gets an existing AlertPolicyChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertPolicyChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertPolicyChannelState, opts ...pulumi.ResourceOption) (*AlertPolicyChannel, error) {
	var resource AlertPolicyChannel
	err := ctx.ReadResource("newrelic:index/alertPolicyChannel:AlertPolicyChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertPolicyChannel resources.
type alertPolicyChannelState struct {
	// The ID of the channel.
	ChannelId *int `pulumi:"channelId"`
	// The ID of the policy.
	PolicyId *int `pulumi:"policyId"`
}

type AlertPolicyChannelState struct {
	// The ID of the channel.
	ChannelId pulumi.IntPtrInput
	// The ID of the policy.
	PolicyId pulumi.IntPtrInput
}

func (AlertPolicyChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPolicyChannelState)(nil)).Elem()
}

type alertPolicyChannelArgs struct {
	// The ID of the channel.
	ChannelId int `pulumi:"channelId"`
	// The ID of the policy.
	PolicyId int `pulumi:"policyId"`
}

// The set of arguments for constructing a AlertPolicyChannel resource.
type AlertPolicyChannelArgs struct {
	// The ID of the channel.
	ChannelId pulumi.IntInput
	// The ID of the policy.
	PolicyId pulumi.IntInput
}

func (AlertPolicyChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertPolicyChannelArgs)(nil)).Elem()
}

