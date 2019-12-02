// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Newrelic
{
    /// <summary>
    /// ## Thresholds
    /// 
    /// The `critical` and `warning` threshold mapping supports the following arguments:
    /// 
    ///   * `duration` - (Required) Identifies the number of minutes the threshold must be passed or met for the alert to trigger. Threshold durations must be between 1 and 60 minutes (inclusive).
    ///   * `value` - (Optional) Threshold value, computed against the `comparison` operator. Supported by "infra_metric" and "infra_process_running" alert condition types.
    ///   * `time_function` - (Optional) Indicates if the condition needs to be sustained or to just break the threshold once; `all` or `any`. Supported by the "infra_metric" alert condition type.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/r/infra_alert_condition.html.markdown.
    /// </summary>
    public partial class InfraAlertCondition : Pulumi.CustomResource
    {
        /// <summary>
        /// The operator used to evaluate the threshold value; "above", "below", "equal".
        /// </summary>
        [Output("comparison")]
        public Output<string?> Comparison { get; private set; } = null!;

        [Output("createdAt")]
        public Output<int> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Identifies the critical threshold parameters for triggering an alert notification. See Thresholds below for details.
        /// </summary>
        [Output("critical")]
        public Output<Outputs.InfraAlertConditionCritical?> Critical { get; private set; } = null!;

        /// <summary>
        /// Set whether to enable the alert condition. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The metric event; for example, system metrics, process metrics, storage metrics, or network metrics.
        /// </summary>
        [Output("event")]
        public Output<string?> Event { get; private set; } = null!;

        /// <summary>
        /// For alerts on integrations, use this instead of `event`.
        /// </summary>
        [Output("integrationProvider")]
        public Output<string?> IntegrationProvider { get; private set; } = null!;

        /// <summary>
        /// The Infrastructure alert condition's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the alert policy where this condition should be used.
        /// </summary>
        [Output("policyId")]
        public Output<int> PolicyId { get; private set; } = null!;

        /// <summary>
        /// Any filters applied to processes; for example: `"commandName = 'java'"`.
        /// </summary>
        [Output("processWhere")]
        public Output<string?> ProcessWhere { get; private set; } = null!;

        /// <summary>
        /// Runbook URL to display in notifications.
        /// </summary>
        [Output("runbookUrl")]
        public Output<string?> RunbookUrl { get; private set; } = null!;

        /// <summary>
        /// The attribute name to identify the type of metric condition; for example, "network", "process", "system", or "storage".
        /// </summary>
        [Output("select")]
        public Output<string?> Select { get; private set; } = null!;

        /// <summary>
        /// The type of Infrastructure alert condition: "infra_process_running", "infra_metric", or "infra_host_not_reporting".
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<int> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Identifies the warning threshold parameters. See Thresholds below for details.
        /// </summary>
        [Output("warning")]
        public Output<Outputs.InfraAlertConditionWarning?> Warning { get; private set; } = null!;

        /// <summary>
        /// Infrastructure host filter for the alert condition.
        /// </summary>
        [Output("where")]
        public Output<string?> Where { get; private set; } = null!;


        /// <summary>
        /// Create a InfraAlertCondition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InfraAlertCondition(string name, InfraAlertConditionArgs args, CustomResourceOptions? options = null)
            : base("newrelic:index/infraAlertCondition:InfraAlertCondition", name, args, MakeResourceOptions(options, ""))
        {
        }

        private InfraAlertCondition(string name, Input<string> id, InfraAlertConditionState? state = null, CustomResourceOptions? options = null)
            : base("newrelic:index/infraAlertCondition:InfraAlertCondition", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InfraAlertCondition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InfraAlertCondition Get(string name, Input<string> id, InfraAlertConditionState? state = null, CustomResourceOptions? options = null)
        {
            return new InfraAlertCondition(name, id, state, options);
        }
    }

    public sealed class InfraAlertConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operator used to evaluate the threshold value; "above", "below", "equal".
        /// </summary>
        [Input("comparison")]
        public Input<string>? Comparison { get; set; }

        /// <summary>
        /// Identifies the critical threshold parameters for triggering an alert notification. See Thresholds below for details.
        /// </summary>
        [Input("critical")]
        public Input<Inputs.InfraAlertConditionCriticalArgs>? Critical { get; set; }

        /// <summary>
        /// Set whether to enable the alert condition. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The metric event; for example, system metrics, process metrics, storage metrics, or network metrics.
        /// </summary>
        [Input("event")]
        public Input<string>? Event { get; set; }

        /// <summary>
        /// For alerts on integrations, use this instead of `event`.
        /// </summary>
        [Input("integrationProvider")]
        public Input<string>? IntegrationProvider { get; set; }

        /// <summary>
        /// The Infrastructure alert condition's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the alert policy where this condition should be used.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<int> PolicyId { get; set; } = null!;

        /// <summary>
        /// Any filters applied to processes; for example: `"commandName = 'java'"`.
        /// </summary>
        [Input("processWhere")]
        public Input<string>? ProcessWhere { get; set; }

        /// <summary>
        /// Runbook URL to display in notifications.
        /// </summary>
        [Input("runbookUrl")]
        public Input<string>? RunbookUrl { get; set; }

        /// <summary>
        /// The attribute name to identify the type of metric condition; for example, "network", "process", "system", or "storage".
        /// </summary>
        [Input("select")]
        public Input<string>? Select { get; set; }

        /// <summary>
        /// The type of Infrastructure alert condition: "infra_process_running", "infra_metric", or "infra_host_not_reporting".
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Identifies the warning threshold parameters. See Thresholds below for details.
        /// </summary>
        [Input("warning")]
        public Input<Inputs.InfraAlertConditionWarningArgs>? Warning { get; set; }

        /// <summary>
        /// Infrastructure host filter for the alert condition.
        /// </summary>
        [Input("where")]
        public Input<string>? Where { get; set; }

        public InfraAlertConditionArgs()
        {
        }
    }

    public sealed class InfraAlertConditionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operator used to evaluate the threshold value; "above", "below", "equal".
        /// </summary>
        [Input("comparison")]
        public Input<string>? Comparison { get; set; }

        [Input("createdAt")]
        public Input<int>? CreatedAt { get; set; }

        /// <summary>
        /// Identifies the critical threshold parameters for triggering an alert notification. See Thresholds below for details.
        /// </summary>
        [Input("critical")]
        public Input<Inputs.InfraAlertConditionCriticalGetArgs>? Critical { get; set; }

        /// <summary>
        /// Set whether to enable the alert condition. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The metric event; for example, system metrics, process metrics, storage metrics, or network metrics.
        /// </summary>
        [Input("event")]
        public Input<string>? Event { get; set; }

        /// <summary>
        /// For alerts on integrations, use this instead of `event`.
        /// </summary>
        [Input("integrationProvider")]
        public Input<string>? IntegrationProvider { get; set; }

        /// <summary>
        /// The Infrastructure alert condition's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the alert policy where this condition should be used.
        /// </summary>
        [Input("policyId")]
        public Input<int>? PolicyId { get; set; }

        /// <summary>
        /// Any filters applied to processes; for example: `"commandName = 'java'"`.
        /// </summary>
        [Input("processWhere")]
        public Input<string>? ProcessWhere { get; set; }

        /// <summary>
        /// Runbook URL to display in notifications.
        /// </summary>
        [Input("runbookUrl")]
        public Input<string>? RunbookUrl { get; set; }

        /// <summary>
        /// The attribute name to identify the type of metric condition; for example, "network", "process", "system", or "storage".
        /// </summary>
        [Input("select")]
        public Input<string>? Select { get; set; }

        /// <summary>
        /// The type of Infrastructure alert condition: "infra_process_running", "infra_metric", or "infra_host_not_reporting".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("updatedAt")]
        public Input<int>? UpdatedAt { get; set; }

        /// <summary>
        /// Identifies the warning threshold parameters. See Thresholds below for details.
        /// </summary>
        [Input("warning")]
        public Input<Inputs.InfraAlertConditionWarningGetArgs>? Warning { get; set; }

        /// <summary>
        /// Infrastructure host filter for the alert condition.
        /// </summary>
        [Input("where")]
        public Input<string>? Where { get; set; }

        public InfraAlertConditionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class InfraAlertConditionCriticalArgs : Pulumi.ResourceArgs
    {
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        [Input("timeFunction")]
        public Input<string>? TimeFunction { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public InfraAlertConditionCriticalArgs()
        {
        }
    }

    public sealed class InfraAlertConditionCriticalGetArgs : Pulumi.ResourceArgs
    {
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        [Input("timeFunction")]
        public Input<string>? TimeFunction { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public InfraAlertConditionCriticalGetArgs()
        {
        }
    }

    public sealed class InfraAlertConditionWarningArgs : Pulumi.ResourceArgs
    {
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        [Input("timeFunction")]
        public Input<string>? TimeFunction { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public InfraAlertConditionWarningArgs()
        {
        }
    }

    public sealed class InfraAlertConditionWarningGetArgs : Pulumi.ResourceArgs
    {
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        [Input("timeFunction")]
        public Input<string>? TimeFunction { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public InfraAlertConditionWarningGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class InfraAlertConditionCritical
    {
        public readonly int Duration;
        public readonly string? TimeFunction;
        public readonly int? Value;

        [OutputConstructor]
        private InfraAlertConditionCritical(
            int duration,
            string? timeFunction,
            int? value)
        {
            Duration = duration;
            TimeFunction = timeFunction;
            Value = value;
        }
    }

    [OutputType]
    public sealed class InfraAlertConditionWarning
    {
        public readonly int Duration;
        public readonly string? TimeFunction;
        public readonly int? Value;

        [OutputConstructor]
        private InfraAlertConditionWarning(
            int duration,
            string? timeFunction,
            int? value)
        {
            Duration = duration;
            TimeFunction = timeFunction;
            Value = value;
        }
    }
    }
}
