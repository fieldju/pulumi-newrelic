// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.NewRelic
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("newrelic");

        public static string? ApiKey { get; set; } = __config.Get("apiKey") ?? Utilities.GetEnv("NEWRELIC_API_KEY");

        public static string? ApiUrl { get; set; } = __config.Get("apiUrl") ?? Utilities.GetEnv("NEWRELIC_API_URL") ?? "https://api.newrelic.com/v2";

        public static string? CacertFile { get; set; } = __config.Get("cacertFile") ?? Utilities.GetEnv("NEWRELIC_API_CACERT");

        public static string? InfraApiUrl { get; set; } = __config.Get("infraApiUrl") ?? Utilities.GetEnv("NEWRELIC_INFRA_API_URL") ?? "https://infra-api.newrelic.com/v2";

        public static string? InfrastructureApiUrl { get; set; } = __config.Get("infrastructureApiUrl");

        public static bool? InsecureSkipVerify { get; set; } = __config.GetBoolean("insecureSkipVerify") ?? Utilities.GetEnvBoolean("NEWRELIC_API_SKIP_VERIFY");

        public static string? InsightsAccountId { get; set; } = __config.Get("insightsAccountId") ?? Utilities.GetEnv("NEWRELIC_INSIGHTS_ACCOUNT_ID");

        public static string? InsightsInsertKey { get; set; } = __config.Get("insightsInsertKey") ?? Utilities.GetEnv("NEWRELIC_INSIGHTS_INSERT_KEY");

        public static string? InsightsInsertUrl { get; set; } = __config.Get("insightsInsertUrl") ?? Utilities.GetEnv("NEWRELIC_INSIGHTS_INSERT_URL") ?? "https://insights-collector.newrelic.com/v1/accounts";

        public static string? InsightsQueryKey { get; set; } = __config.Get("insightsQueryKey") ?? Utilities.GetEnv("NEWRELIC_INSIGHTS_QUERY_KEY");

        public static string? InsightsQueryUrl { get; set; } = __config.Get("insightsQueryUrl") ?? Utilities.GetEnv("NEWRELIC_INSIGHTS_QUERY_URL") ?? "https://insights-api.newrelic.com/v1/accounts";

        public static string? NerdgraphApiUrl { get; set; } = __config.Get("nerdgraphApiUrl");

        public static string? SyntheticsApiUrl { get; set; } = __config.Get("syntheticsApiUrl") ?? Utilities.GetEnv("NEWRELIC_SYNTHETICS_API_URL") ?? "https://synthetics.newrelic.com/synthetics/api/v3";

    }
    namespace ConfigTypes
    {
    }
}
