// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.NewRelic.Outputs
{

    [OutputType]
    public sealed class DashboardFilter
    {
        public readonly ImmutableArray<string> Attributes;
        public readonly ImmutableArray<string> EventTypes;

        [OutputConstructor]
        private DashboardFilter(
            ImmutableArray<string> attributes,

            ImmutableArray<string> eventTypes)
        {
            Attributes = attributes;
            EventTypes = eventTypes;
        }
    }
}
