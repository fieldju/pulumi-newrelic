// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package newrelic

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/d/alert_policy.html.markdown.
func LookupAlertPolicy(ctx *pulumi.Context, args *LookupAlertPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAlertPolicyResult, error) {
	var rv LookupAlertPolicyResult
	err := ctx.Invoke("newrelic:index/getAlertPolicy:getAlertPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlertPolicy.
type LookupAlertPolicyArgs struct {
	IncidentPreference *string `pulumi:"incidentPreference"`
	// The name of the alert policy in New Relic.
	Name string `pulumi:"name"`
}


// A collection of values returned by getAlertPolicy.
type LookupAlertPolicyResult struct {
	// The time the policy was created.
	CreatedAt string `pulumi:"createdAt"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The rollup strategy for the policy. Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET. The default is PER_POLICY.
	IncidentPreference *string `pulumi:"incidentPreference"`
	Name string `pulumi:"name"`
	// The time the policy was last updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

