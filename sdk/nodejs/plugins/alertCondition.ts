// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this resource to create and manage plugins alert conditions in New Relic.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as newrelic from "@pulumi/newrelic";
 *
 * const fooPlugin = newrelic.plugins.getPlugin({
 *     guid: "com.example.my-plugin",
 * });
 * const fooPluginComponent = fooPlugin.then(fooPlugin => newrelic.plugins.getPluginComponent({
 *     pluginId: fooPlugin.id,
 *     name: "MyPlugin",
 * }));
 * const fooAlertPolicy = new newrelic.AlertPolicy("fooAlertPolicy", {});
 * const fooAlertCondition = new newrelic.plugins.AlertCondition("fooAlertCondition", {
 *     policyId: fooAlertPolicy.id,
 *     entities: [fooPluginComponent.then(fooPluginComponent => fooPluginComponent.id)],
 *     metric: "Component/Summary/Consumers[consumers]",
 *     pluginId: fooPlugin.then(fooPlugin => fooPlugin.id),
 *     pluginGuid: fooPlugin.then(fooPlugin => fooPlugin.guid),
 *     valueFunction: "average",
 *     metricDescription: "Queue consumers",
 *     terms: [{
 *         duration: 5,
 *         operator: "below",
 *         priority: "critical",
 *         threshold: "0.75",
 *         timeFunction: "all",
 *     }],
 * });
 * ```
 * ## Terms
 *
 * The `term` mapping supports the following arguments:
 *
 *   * `duration` - (Required) In minutes, must be in the range of `5` to `120`, inclusive.
 *   * `operator` - (Optional) `above`, `below`, or `equal`.  Defaults to `equal`.
 *   * `priority` - (Optional) `critical` or `warning`.  Defaults to `critical`.
 *   * `threshold` - (Required) Must be 0 or greater.
 *   * `timeFunction` - (Required) `all` or `any`.
 */
export class AlertCondition extends pulumi.CustomResource {
    /**
     * Get an existing AlertCondition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertConditionState, opts?: pulumi.CustomResourceOptions): AlertCondition {
        return new AlertCondition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'newrelic:plugins/alertCondition:AlertCondition';

    /**
     * Returns true if the given object is an instance of AlertCondition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertCondition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertCondition.__pulumiType;
    }

    /**
     * Whether or not this condition is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The plugin component IDs to target.
     */
    public readonly entities!: pulumi.Output<number[]>;
    /**
     * The plugin metric to evaluate.
     */
    public readonly metric!: pulumi.Output<string>;
    /**
     * The metric description.
     */
    public readonly metricDescription!: pulumi.Output<string>;
    /**
     * The title of the condition. Must be between 1 and 64 characters, inclusive.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The GUID of the plugin which produces the metric.
     */
    public readonly pluginGuid!: pulumi.Output<string>;
    /**
     * The ID of the installed plugin instance which produces the metric.
     */
    public readonly pluginId!: pulumi.Output<string>;
    /**
     * The ID of the policy where this condition should be used.
     */
    public readonly policyId!: pulumi.Output<number>;
    /**
     * Runbook URL to display in notifications.
     */
    public readonly runbookUrl!: pulumi.Output<string | undefined>;
    /**
     * A list of terms for this condition. See Terms below for details.
     */
    public readonly terms!: pulumi.Output<outputs.plugins.AlertConditionTerm[]>;
    /**
     * The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
     */
    public readonly valueFunction!: pulumi.Output<string>;

    /**
     * Create a AlertCondition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertConditionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertConditionArgs | AlertConditionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AlertConditionState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["entities"] = state ? state.entities : undefined;
            inputs["metric"] = state ? state.metric : undefined;
            inputs["metricDescription"] = state ? state.metricDescription : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pluginGuid"] = state ? state.pluginGuid : undefined;
            inputs["pluginId"] = state ? state.pluginId : undefined;
            inputs["policyId"] = state ? state.policyId : undefined;
            inputs["runbookUrl"] = state ? state.runbookUrl : undefined;
            inputs["terms"] = state ? state.terms : undefined;
            inputs["valueFunction"] = state ? state.valueFunction : undefined;
        } else {
            const args = argsOrState as AlertConditionArgs | undefined;
            if (!args || args.entities === undefined) {
                throw new Error("Missing required property 'entities'");
            }
            if (!args || args.metric === undefined) {
                throw new Error("Missing required property 'metric'");
            }
            if (!args || args.metricDescription === undefined) {
                throw new Error("Missing required property 'metricDescription'");
            }
            if (!args || args.pluginGuid === undefined) {
                throw new Error("Missing required property 'pluginGuid'");
            }
            if (!args || args.pluginId === undefined) {
                throw new Error("Missing required property 'pluginId'");
            }
            if (!args || args.policyId === undefined) {
                throw new Error("Missing required property 'policyId'");
            }
            if (!args || args.terms === undefined) {
                throw new Error("Missing required property 'terms'");
            }
            if (!args || args.valueFunction === undefined) {
                throw new Error("Missing required property 'valueFunction'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["entities"] = args ? args.entities : undefined;
            inputs["metric"] = args ? args.metric : undefined;
            inputs["metricDescription"] = args ? args.metricDescription : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pluginGuid"] = args ? args.pluginGuid : undefined;
            inputs["pluginId"] = args ? args.pluginId : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["runbookUrl"] = args ? args.runbookUrl : undefined;
            inputs["terms"] = args ? args.terms : undefined;
            inputs["valueFunction"] = args ? args.valueFunction : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AlertCondition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertCondition resources.
 */
export interface AlertConditionState {
    /**
     * Whether or not this condition is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The plugin component IDs to target.
     */
    readonly entities?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The plugin metric to evaluate.
     */
    readonly metric?: pulumi.Input<string>;
    /**
     * The metric description.
     */
    readonly metricDescription?: pulumi.Input<string>;
    /**
     * The title of the condition. Must be between 1 and 64 characters, inclusive.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The GUID of the plugin which produces the metric.
     */
    readonly pluginGuid?: pulumi.Input<string>;
    /**
     * The ID of the installed plugin instance which produces the metric.
     */
    readonly pluginId?: pulumi.Input<string>;
    /**
     * The ID of the policy where this condition should be used.
     */
    readonly policyId?: pulumi.Input<number>;
    /**
     * Runbook URL to display in notifications.
     */
    readonly runbookUrl?: pulumi.Input<string>;
    /**
     * A list of terms for this condition. See Terms below for details.
     */
    readonly terms?: pulumi.Input<pulumi.Input<inputs.plugins.AlertConditionTerm>[]>;
    /**
     * The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
     */
    readonly valueFunction?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlertCondition resource.
 */
export interface AlertConditionArgs {
    /**
     * Whether or not this condition is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The plugin component IDs to target.
     */
    readonly entities: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The plugin metric to evaluate.
     */
    readonly metric: pulumi.Input<string>;
    /**
     * The metric description.
     */
    readonly metricDescription: pulumi.Input<string>;
    /**
     * The title of the condition. Must be between 1 and 64 characters, inclusive.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The GUID of the plugin which produces the metric.
     */
    readonly pluginGuid: pulumi.Input<string>;
    /**
     * The ID of the installed plugin instance which produces the metric.
     */
    readonly pluginId: pulumi.Input<string>;
    /**
     * The ID of the policy where this condition should be used.
     */
    readonly policyId: pulumi.Input<number>;
    /**
     * Runbook URL to display in notifications.
     */
    readonly runbookUrl?: pulumi.Input<string>;
    /**
     * A list of terms for this condition. See Terms below for details.
     */
    readonly terms: pulumi.Input<pulumi.Input<inputs.plugins.AlertConditionTerm>[]>;
    /**
     * The value function to apply to the metric data.  One of `min`, `max`, `average`, `sampleSize`, `total`, or `percent`.
     */
    readonly valueFunction: pulumi.Input<string>;
}
