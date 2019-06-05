// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as newrelic from "@pulumi/newrelic";
 * 
 * const fooAlertPolicy = new newrelic.AlertPolicy("foo", {});
 * const fooInfraAlertCondition = new newrelic.InfraAlertCondition("foo", {
 *     comparison: "above",
 *     critical: {
 *         duration: 25,
 *         timeFunction: "all",
 *         value: 90,
 *     },
 *     event: "StorageSample",
 *     policyId: fooAlertPolicy.id,
 *     select: "diskUsedPercent",
 *     type: "infra_metric",
 *     where: "(`hostname` LIKE '%frontend%')",
 * });
 * ```
 * 
 * ## Thresholds
 * 
 * The `critical` and `warning` threshold mapping supports the following arguments:
 * 
 *   * `duration` - (Required) Identifies the number of minutes the threshold must be passed or met for the alert to trigger. Threshold durations must be between 1 and 60 minutes (inclusive).
 *   * `value` - (Optional) Threshold value, computed against the `comparison` operator. Supported by "infra_metric" and "infra_process_running" alert condition types.
 *   * `time_function` - (Optional) Indicates if the condition needs to be sustained or to just break the threshold once; `all` or `any`. Supported by the "infra_metric" alert condition type.
 */
export class InfraAlertCondition extends pulumi.CustomResource {
    /**
     * Get an existing InfraAlertCondition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InfraAlertConditionState, opts?: pulumi.CustomResourceOptions): InfraAlertCondition {
        return new InfraAlertCondition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'newrelic:index/infraAlertCondition:InfraAlertCondition';

    /**
     * Returns true if the given object is an instance of InfraAlertCondition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InfraAlertCondition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InfraAlertCondition.__pulumiType;
    }

    /**
     * The operator used to evaluate the threshold value; "above", "below", "equal".
     */
    public readonly comparison!: pulumi.Output<string | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<number>;
    /**
     * Identifies the critical threshold parameters for triggering an alert notification. See Thresholds below for details.
     */
    public readonly critical!: pulumi.Output<{ duration: number, timeFunction?: string, value?: number } | undefined>;
    /**
     * Set whether to enable the alert condition. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The metric event; for example, system metrics, process metrics, storage metrics, or network metrics.
     */
    public readonly event!: pulumi.Output<string | undefined>;
    /**
     * For alerts on integrations, use this instead of `event`. 
     */
    public readonly integrationProvider!: pulumi.Output<string | undefined>;
    /**
     * The Infrastructure alert condition's name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the alert policy where this condition should be used.
     */
    public readonly policyId!: pulumi.Output<number>;
    /**
     * Any filters applied to processes; for example: `"commandName = 'java'"`.
     */
    public readonly processWhere!: pulumi.Output<string | undefined>;
    /**
     * The attribute name to identify the type of metric condition; for example, "network", "process", "system", or "storage".
     */
    public readonly select!: pulumi.Output<string | undefined>;
    /**
     * The type of Infrastructure alert condition: "infra_process_running", "infra_metric", or "infra_host_not_reporting".
     */
    public readonly type!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<number>;
    /**
     * Identifies the warning threshold parameters. See Thresholds below for details.
     */
    public readonly warning!: pulumi.Output<{ duration: number, timeFunction?: string, value?: number } | undefined>;
    /**
     * Infrastructure host filter for the alert condition.
     */
    public readonly where!: pulumi.Output<string | undefined>;

    /**
     * Create a InfraAlertCondition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InfraAlertConditionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InfraAlertConditionArgs | InfraAlertConditionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InfraAlertConditionState | undefined;
            inputs["comparison"] = state ? state.comparison : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["critical"] = state ? state.critical : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["event"] = state ? state.event : undefined;
            inputs["integrationProvider"] = state ? state.integrationProvider : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyId"] = state ? state.policyId : undefined;
            inputs["processWhere"] = state ? state.processWhere : undefined;
            inputs["select"] = state ? state.select : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["warning"] = state ? state.warning : undefined;
            inputs["where"] = state ? state.where : undefined;
        } else {
            const args = argsOrState as InfraAlertConditionArgs | undefined;
            if (!args || args.policyId === undefined) {
                throw new Error("Missing required property 'policyId'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["comparison"] = args ? args.comparison : undefined;
            inputs["critical"] = args ? args.critical : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["event"] = args ? args.event : undefined;
            inputs["integrationProvider"] = args ? args.integrationProvider : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["processWhere"] = args ? args.processWhere : undefined;
            inputs["select"] = args ? args.select : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["warning"] = args ? args.warning : undefined;
            inputs["where"] = args ? args.where : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        super(InfraAlertCondition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InfraAlertCondition resources.
 */
export interface InfraAlertConditionState {
    /**
     * The operator used to evaluate the threshold value; "above", "below", "equal".
     */
    readonly comparison?: pulumi.Input<string>;
    readonly createdAt?: pulumi.Input<number>;
    /**
     * Identifies the critical threshold parameters for triggering an alert notification. See Thresholds below for details.
     */
    readonly critical?: pulumi.Input<{ duration: pulumi.Input<number>, timeFunction?: pulumi.Input<string>, value?: pulumi.Input<number> }>;
    /**
     * Set whether to enable the alert condition. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The metric event; for example, system metrics, process metrics, storage metrics, or network metrics.
     */
    readonly event?: pulumi.Input<string>;
    /**
     * For alerts on integrations, use this instead of `event`. 
     */
    readonly integrationProvider?: pulumi.Input<string>;
    /**
     * The Infrastructure alert condition's name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the alert policy where this condition should be used.
     */
    readonly policyId?: pulumi.Input<number>;
    /**
     * Any filters applied to processes; for example: `"commandName = 'java'"`.
     */
    readonly processWhere?: pulumi.Input<string>;
    /**
     * The attribute name to identify the type of metric condition; for example, "network", "process", "system", or "storage".
     */
    readonly select?: pulumi.Input<string>;
    /**
     * The type of Infrastructure alert condition: "infra_process_running", "infra_metric", or "infra_host_not_reporting".
     */
    readonly type?: pulumi.Input<string>;
    readonly updatedAt?: pulumi.Input<number>;
    /**
     * Identifies the warning threshold parameters. See Thresholds below for details.
     */
    readonly warning?: pulumi.Input<{ duration: pulumi.Input<number>, timeFunction?: pulumi.Input<string>, value?: pulumi.Input<number> }>;
    /**
     * Infrastructure host filter for the alert condition.
     */
    readonly where?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InfraAlertCondition resource.
 */
export interface InfraAlertConditionArgs {
    /**
     * The operator used to evaluate the threshold value; "above", "below", "equal".
     */
    readonly comparison?: pulumi.Input<string>;
    /**
     * Identifies the critical threshold parameters for triggering an alert notification. See Thresholds below for details.
     */
    readonly critical?: pulumi.Input<{ duration: pulumi.Input<number>, timeFunction?: pulumi.Input<string>, value?: pulumi.Input<number> }>;
    /**
     * Set whether to enable the alert condition. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The metric event; for example, system metrics, process metrics, storage metrics, or network metrics.
     */
    readonly event?: pulumi.Input<string>;
    /**
     * For alerts on integrations, use this instead of `event`. 
     */
    readonly integrationProvider?: pulumi.Input<string>;
    /**
     * The Infrastructure alert condition's name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the alert policy where this condition should be used.
     */
    readonly policyId: pulumi.Input<number>;
    /**
     * Any filters applied to processes; for example: `"commandName = 'java'"`.
     */
    readonly processWhere?: pulumi.Input<string>;
    /**
     * The attribute name to identify the type of metric condition; for example, "network", "process", "system", or "storage".
     */
    readonly select?: pulumi.Input<string>;
    /**
     * The type of Infrastructure alert condition: "infra_process_running", "infra_metric", or "infra_host_not_reporting".
     */
    readonly type: pulumi.Input<string>;
    /**
     * Identifies the warning threshold parameters. See Thresholds below for details.
     */
    readonly warning?: pulumi.Input<{ duration: pulumi.Input<number>, timeFunction?: pulumi.Input<string>, value?: pulumi.Input<number> }>;
    /**
     * Infrastructure host filter for the alert condition.
     */
    readonly where?: pulumi.Input<string>;
}
