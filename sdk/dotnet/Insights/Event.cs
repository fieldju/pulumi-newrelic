// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.NewRelic.Insights
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/r/insights_event.html.markdown.
    /// </summary>
    public partial class Event : Pulumi.CustomResource
    {
        /// <summary>
        /// An event to insert into Insights. Multiple event blocks can be defined. See Events below for details.
        /// </summary>
        [Output("events")]
        public Output<ImmutableArray<Outputs.EventEvents>> Events { get; private set; } = null!;


        /// <summary>
        /// Create a Event resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Event(string name, EventArgs args, CustomResourceOptions? options = null)
            : base("newrelic:insights/event:Event", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Event(string name, Input<string> id, EventState? state = null, CustomResourceOptions? options = null)
            : base("newrelic:insights/event:Event", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Event resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Event Get(string name, Input<string> id, EventState? state = null, CustomResourceOptions? options = null)
        {
            return new Event(name, id, state, options);
        }
    }

    public sealed class EventArgs : Pulumi.ResourceArgs
    {
        [Input("events", required: true)]
        private InputList<Inputs.EventEventsArgs>? _events;

        /// <summary>
        /// An event to insert into Insights. Multiple event blocks can be defined. See Events below for details.
        /// </summary>
        public InputList<Inputs.EventEventsArgs> Events
        {
            get => _events ?? (_events = new InputList<Inputs.EventEventsArgs>());
            set => _events = value;
        }

        public EventArgs()
        {
        }
    }

    public sealed class EventState : Pulumi.ResourceArgs
    {
        [Input("events")]
        private InputList<Inputs.EventEventsGetArgs>? _events;

        /// <summary>
        /// An event to insert into Insights. Multiple event blocks can be defined. See Events below for details.
        /// </summary>
        public InputList<Inputs.EventEventsGetArgs> Events
        {
            get => _events ?? (_events = new InputList<Inputs.EventEventsGetArgs>());
            set => _events = value;
        }

        public EventState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class EventEventsArgs : Pulumi.ResourceArgs
    {
        [Input("attributes", required: true)]
        private InputList<EventEventsAttributesArgs>? _attributes;
        public InputList<EventEventsAttributesArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<EventEventsAttributesArgs>());
            set => _attributes = value;
        }

        [Input("timestamp")]
        public Input<int>? Timestamp { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public EventEventsArgs()
        {
        }
    }

    public sealed class EventEventsAttributesArgs : Pulumi.ResourceArgs
    {
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public EventEventsAttributesArgs()
        {
        }
    }

    public sealed class EventEventsAttributesGetArgs : Pulumi.ResourceArgs
    {
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public EventEventsAttributesGetArgs()
        {
        }
    }

    public sealed class EventEventsGetArgs : Pulumi.ResourceArgs
    {
        [Input("attributes", required: true)]
        private InputList<EventEventsAttributesGetArgs>? _attributes;
        public InputList<EventEventsAttributesGetArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<EventEventsAttributesGetArgs>());
            set => _attributes = value;
        }

        [Input("timestamp")]
        public Input<int>? Timestamp { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public EventEventsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class EventEvents
    {
        public readonly ImmutableArray<EventEventsAttributes> Attributes;
        public readonly int? Timestamp;
        public readonly string Type;

        [OutputConstructor]
        private EventEvents(
            ImmutableArray<EventEventsAttributes> attributes,
            int? timestamp,
            string type)
        {
            Attributes = attributes;
            Timestamp = timestamp;
            Type = type;
        }
    }

    [OutputType]
    public sealed class EventEventsAttributes
    {
        public readonly string Key;
        public readonly string? Type;
        public readonly string Value;

        [OutputConstructor]
        private EventEventsAttributes(
            string key,
            string? type,
            string value)
        {
            Key = key;
            Type = type;
            Value = value;
        }
    }
    }
}
