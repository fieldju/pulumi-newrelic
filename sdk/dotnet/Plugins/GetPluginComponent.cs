// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.NewRelic.Plugins
{
    public static class GetPluginComponent
    {
        public static Task<GetPluginComponentResult> InvokeAsync(GetPluginComponentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPluginComponentResult>("newrelic:plugins/getPluginComponent:getPluginComponent", args ?? new GetPluginComponentArgs(), options.WithVersion());
    }


    public sealed class GetPluginComponentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the plugin component.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the plugin instance this component belongs to.
        /// </summary>
        [Input("pluginId", required: true)]
        public int PluginId { get; set; }

        public GetPluginComponentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPluginComponentResult
    {
        /// <summary>
        /// The health status of the plugin component.
        /// </summary>
        public readonly string HealthStatus;
        /// <summary>
        /// The ID of the plugin component.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly int PluginId;

        [OutputConstructor]
        private GetPluginComponentResult(
            string healthStatus,

            string id,

            string name,

            int pluginId)
        {
            HealthStatus = healthStatus;
            Id = id;
            Name = name;
            PluginId = pluginId;
        }
    }
}
