// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package synthetics

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about a specific synthetics monitor in New Relic that already exists. This can be used to set up a Synthetics alert condition.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/d/synthetics_monitor.html.markdown.
func LookupMonitor(ctx *pulumi.Context, args *LookupMonitorArgs, opts ...pulumi.InvokeOption) (*LookupMonitorResult, error) {
	var rv LookupMonitorResult
	err := ctx.Invoke("newrelic:synthetics/getMonitor:getMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMonitor.
type LookupMonitorArgs struct {
	// The name of the synthetics monitor in New Relic.
	Name string `pulumi:"name"`
}


// A collection of values returned by getMonitor.
type LookupMonitorResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the synthetics monitor.
	MonitorId string `pulumi:"monitorId"`
	Name string `pulumi:"name"`
}

