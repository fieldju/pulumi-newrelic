// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as newrelic from "@pulumi/newrelic";
 * 
 * const fooAlertPolicy = new newrelic.AlertPolicy("foo", {});
 * const fooNrqlAlertCondition = new newrelic.NrqlAlertCondition("foo", {
 *     enabled: true,
 *     nrql: {
 *         query: "SELECT count(*) FROM SyntheticCheck WHERE monitorId = '<monitorId>'",
 *         sinceValue: "3",
 *     },
 *     policyId: fooAlertPolicy.id,
 *     runbookUrl: "https://www.example.com",
 *     terms: [{
 *         duration: 5,
 *         operator: "below",
 *         priority: "critical",
 *         threshold: 1,
 *         timeFunction: "all",
 *     }],
 *     type: "static",
 *     valueFunction: "singleValue",
 * });
 * ```
 * 
 * ## Terms
 * 
 * The `term` mapping supports the following arguments:
 * 
 * - `duration` - (Required) In minutes, must be in the range of `1` to `120`, inclusive.
 * - `operator` - (Optional) `above`, `below`, or `equal`. Defaults to `equal`.
 * - `priority` - (Optional) `critical` or `warning`. Defaults to `critical`.
 * - `threshold` - (Required) Must be 0 or greater.
 * - `timeFunction` - (Required) `all` or `any`.
 * 
 * ## NRQL
 * 
 * The `nrql` attribute supports the following arguments:
 * 
 * - `query` - (Required) The NRQL query to execute for the condition.
 * - `sinceValue` - (Required) The value to be used in the `SINCE <X> MINUTES AGO` clause for the NRQL query. Must be between `1` and `20`.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/r/nrql_alert_condition.html.markdown.
 */
export class NrqlAlertCondition extends pulumi.CustomResource {
    /**
     * Get an existing NrqlAlertCondition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NrqlAlertConditionState, opts?: pulumi.CustomResourceOptions): NrqlAlertCondition {
        return new NrqlAlertCondition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'newrelic:index/nrqlAlertCondition:NrqlAlertCondition';

    /**
     * Returns true if the given object is an instance of NrqlAlertCondition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NrqlAlertCondition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NrqlAlertCondition.__pulumiType;
    }

    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly expectedGroups!: pulumi.Output<number | undefined>;
    public readonly ignoreOverlap!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly nrql!: pulumi.Output<outputs.NrqlAlertConditionNrql>;
    public readonly policyId!: pulumi.Output<number>;
    public readonly runbookUrl!: pulumi.Output<string | undefined>;
    public readonly terms!: pulumi.Output<outputs.NrqlAlertConditionTerm[]>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly valueFunction!: pulumi.Output<string | undefined>;

    /**
     * Create a NrqlAlertCondition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NrqlAlertConditionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NrqlAlertConditionArgs | NrqlAlertConditionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NrqlAlertConditionState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["expectedGroups"] = state ? state.expectedGroups : undefined;
            inputs["ignoreOverlap"] = state ? state.ignoreOverlap : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nrql"] = state ? state.nrql : undefined;
            inputs["policyId"] = state ? state.policyId : undefined;
            inputs["runbookUrl"] = state ? state.runbookUrl : undefined;
            inputs["terms"] = state ? state.terms : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["valueFunction"] = state ? state.valueFunction : undefined;
        } else {
            const args = argsOrState as NrqlAlertConditionArgs | undefined;
            if (!args || args.nrql === undefined) {
                throw new Error("Missing required property 'nrql'");
            }
            if (!args || args.policyId === undefined) {
                throw new Error("Missing required property 'policyId'");
            }
            if (!args || args.terms === undefined) {
                throw new Error("Missing required property 'terms'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["expectedGroups"] = args ? args.expectedGroups : undefined;
            inputs["ignoreOverlap"] = args ? args.ignoreOverlap : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nrql"] = args ? args.nrql : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["runbookUrl"] = args ? args.runbookUrl : undefined;
            inputs["terms"] = args ? args.terms : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["valueFunction"] = args ? args.valueFunction : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NrqlAlertCondition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NrqlAlertCondition resources.
 */
export interface NrqlAlertConditionState {
    readonly enabled?: pulumi.Input<boolean>;
    readonly expectedGroups?: pulumi.Input<number>;
    readonly ignoreOverlap?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly nrql?: pulumi.Input<inputs.NrqlAlertConditionNrql>;
    readonly policyId?: pulumi.Input<number>;
    readonly runbookUrl?: pulumi.Input<string>;
    readonly terms?: pulumi.Input<pulumi.Input<inputs.NrqlAlertConditionTerm>[]>;
    readonly type?: pulumi.Input<string>;
    readonly valueFunction?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NrqlAlertCondition resource.
 */
export interface NrqlAlertConditionArgs {
    readonly enabled?: pulumi.Input<boolean>;
    readonly expectedGroups?: pulumi.Input<number>;
    readonly ignoreOverlap?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly nrql: pulumi.Input<inputs.NrqlAlertConditionNrql>;
    readonly policyId: pulumi.Input<number>;
    readonly runbookUrl?: pulumi.Input<string>;
    readonly terms: pulumi.Input<pulumi.Input<inputs.NrqlAlertConditionTerm>[]>;
    readonly type?: pulumi.Input<string>;
    readonly valueFunction?: pulumi.Input<string>;
}
