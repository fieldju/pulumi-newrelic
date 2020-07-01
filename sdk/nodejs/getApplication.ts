// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * #### DEPRECATED! Use at your own risk. Use the `newrelic.getEntity` data source instead. This feature may be removed in the next major release.
 *
 * Use this data source to get information about a specific application in New Relic that already exists.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as newrelic from "@pulumi/newrelic";
 *
 * const app = newrelic.getApplication({
 *     name: "my-app",
 * });
 * const fooAlertPolicy = new newrelic.AlertPolicy("fooAlertPolicy", {});
 * const fooAlertCondition = new newrelic.AlertCondition("fooAlertCondition", {
 *     policyId: fooAlertPolicy.id,
 *     type: "apm_app_metric",
 *     entities: [app.then(app => app.id)],
 *     metric: "apdex",
 *     runbookUrl: "https://www.example.com",
 *     terms: [{
 *         duration: 5,
 *         operator: "below",
 *         priority: "critical",
 *         threshold: "0.75",
 *         timeFunction: "all",
 *     }],
 * });
 * ```
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("newrelic:index/getApplication:getApplication", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationArgs {
    /**
     * The name of the application in New Relic.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getApplication.
 */
export interface GetApplicationResult {
    /**
     * A list of host IDs associated with the application.
     */
    readonly hostIds: number[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of instance IDs associated with the application.
     */
    readonly instanceIds: number[];
    readonly name: string;
}
