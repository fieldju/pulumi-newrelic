// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > **NOTE:** Applications are not created by this resource, but are created by
 * a reporting agent.
 * 
 * Use this resource to manage configuration for an application that already
 * exists in New Relic.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as newrelic from "@pulumi/newrelic";
 * 
 * const app = new newrelic.plugins.ApplicationSettings("app", {
 *     appApdexThreshold: 0.7,
 *     enableRealUserMonitoring: false,
 *     endUserApdexThreshold: 0.8,
 * });
 * ```
 * 
 * ## Notes
 * 
 * > **NOTE:** Applications that have reported data in the last twelve hours
 * cannot be deleted.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/r/application_settings.html.markdown.
 */
export class ApplicationSettings extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationSettingsState, opts?: pulumi.CustomResourceOptions): ApplicationSettings {
        return new ApplicationSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'newrelic:plugins/applicationSettings:ApplicationSettings';

    /**
     * Returns true if the given object is an instance of ApplicationSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationSettings.__pulumiType;
    }

    /**
     * The appex threshold for the New Relic application.
     */
    public readonly appApdexThreshold!: pulumi.Output<number>;
    /**
     * Enable or disable real user monitoring for the New Relic application.
     */
    public readonly enableRealUserMonitoring!: pulumi.Output<boolean>;
    /**
     * The user's apdex threshold for the New Relic application.
     */
    public readonly endUserApdexThreshold!: pulumi.Output<number>;
    /**
     * The name of the application in New Relic APM.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ApplicationSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationSettingsArgs | ApplicationSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApplicationSettingsState | undefined;
            inputs["appApdexThreshold"] = state ? state.appApdexThreshold : undefined;
            inputs["enableRealUserMonitoring"] = state ? state.enableRealUserMonitoring : undefined;
            inputs["endUserApdexThreshold"] = state ? state.endUserApdexThreshold : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ApplicationSettingsArgs | undefined;
            if (!args || args.appApdexThreshold === undefined) {
                throw new Error("Missing required property 'appApdexThreshold'");
            }
            if (!args || args.enableRealUserMonitoring === undefined) {
                throw new Error("Missing required property 'enableRealUserMonitoring'");
            }
            if (!args || args.endUserApdexThreshold === undefined) {
                throw new Error("Missing required property 'endUserApdexThreshold'");
            }
            inputs["appApdexThreshold"] = args ? args.appApdexThreshold : undefined;
            inputs["enableRealUserMonitoring"] = args ? args.enableRealUserMonitoring : undefined;
            inputs["endUserApdexThreshold"] = args ? args.endUserApdexThreshold : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ApplicationSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationSettings resources.
 */
export interface ApplicationSettingsState {
    /**
     * The appex threshold for the New Relic application.
     */
    readonly appApdexThreshold?: pulumi.Input<number>;
    /**
     * Enable or disable real user monitoring for the New Relic application.
     */
    readonly enableRealUserMonitoring?: pulumi.Input<boolean>;
    /**
     * The user's apdex threshold for the New Relic application.
     */
    readonly endUserApdexThreshold?: pulumi.Input<number>;
    /**
     * The name of the application in New Relic APM.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationSettings resource.
 */
export interface ApplicationSettingsArgs {
    /**
     * The appex threshold for the New Relic application.
     */
    readonly appApdexThreshold: pulumi.Input<number>;
    /**
     * Enable or disable real user monitoring for the New Relic application.
     */
    readonly enableRealUserMonitoring: pulumi.Input<boolean>;
    /**
     * The user's apdex threshold for the New Relic application.
     */
    readonly endUserApdexThreshold: pulumi.Input<number>;
    /**
     * The name of the application in New Relic APM.
     */
    readonly name?: pulumi.Input<string>;
}
