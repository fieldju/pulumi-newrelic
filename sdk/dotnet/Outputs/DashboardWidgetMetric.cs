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
    public sealed class DashboardWidgetMetric
    {
        public readonly string Name;
        public readonly string? Scope;
        public readonly string? Units;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private DashboardWidgetMetric(
            string name,

            string? scope,

            string? units,

            ImmutableArray<string> values)
        {
            Name = name;
            Scope = scope;
            Units = units;
            Values = values;
        }
    }
}
