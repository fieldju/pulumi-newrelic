// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.NewRelic.Synthetics
{
    /// <summary>
    /// Use this resource to create, update, and delete a Synthetics label in New Relic.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using NewRelic = Pulumi.NewRelic;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new NewRelic.Synthetics.Label("foo", new NewRelic.Synthetics.LabelArgs
    ///         {
    ///             MonitorId = newrelic_synthetics_monitor.Foo.Id,
    ///             Type = "MyCategory",
    ///             Value = "MyValue",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Label : Pulumi.CustomResource
    {
        /// <summary>
        /// The URL of the Synthetics label.
        /// </summary>
        [Output("href")]
        public Output<string> Href { get; private set; } = null!;

        /// <summary>
        /// The ID of the monitor that will be assigned the label.
        /// </summary>
        [Output("monitorId")]
        public Output<string> MonitorId { get; private set; } = null!;

        /// <summary>
        /// A string representing the label key/category.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A string representing the label value.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Label resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Label(string name, LabelArgs args, CustomResourceOptions? options = null)
            : base("newrelic:synthetics/label:Label", name, args ?? new LabelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Label(string name, Input<string> id, LabelState? state = null, CustomResourceOptions? options = null)
            : base("newrelic:synthetics/label:Label", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Label resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Label Get(string name, Input<string> id, LabelState? state = null, CustomResourceOptions? options = null)
        {
            return new Label(name, id, state, options);
        }
    }

    public sealed class LabelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the monitor that will be assigned the label.
        /// </summary>
        [Input("monitorId", required: true)]
        public Input<string> MonitorId { get; set; } = null!;

        /// <summary>
        /// A string representing the label key/category.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// A string representing the label value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public LabelArgs()
        {
        }
    }

    public sealed class LabelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the Synthetics label.
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        /// <summary>
        /// The ID of the monitor that will be assigned the label.
        /// </summary>
        [Input("monitorId")]
        public Input<string>? MonitorId { get; set; }

        /// <summary>
        /// A string representing the label key/category.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A string representing the label value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public LabelState()
        {
        }
    }
}
