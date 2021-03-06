// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.NewRelic
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("newrelic");
        public static int? AccountId { get; set; } = __config.GetInt32("accountId") ?? Utilities.GetEnvInt32("NEW_RELIC_ACCOUNT_ID");

        public static string? AdminApiKey { get; set; } = __config.Get("adminApiKey") ?? Utilities.GetEnv("NEW_RELIC_ADMIN_API_KEY");

        public static string? ApiKey { get; set; } = __config.Get("apiKey") ?? Utilities.GetEnv("NEW_RELIC_API_KEY");

        public static string? ApiUrl { get; set; } = __config.Get("apiUrl");

        public static string? CacertFile { get; set; } = __config.Get("cacertFile") ?? Utilities.GetEnv("NEW_RELIC_API_CACERT");

        public static string? InfrastructureApiUrl { get; set; } = __config.Get("infrastructureApiUrl");

        public static bool? InsecureSkipVerify { get; set; } = __config.GetBoolean("insecureSkipVerify") ?? Utilities.GetEnvBoolean("NEW_RELIC_API_SKIP_VERIFY");

        public static string? InsightsInsertKey { get; set; } = __config.Get("insightsInsertKey") ?? Utilities.GetEnv("NEW_RELIC_INSIGHTS_INSERT_KEY");

        public static string? InsightsInsertUrl { get; set; } = __config.Get("insightsInsertUrl") ?? Utilities.GetEnv("NEW_RELIC_INSIGHTS_INSERT_URL") ?? "https://insights-collector.newrelic.com/v1/accounts";

        public static string? InsightsQueryUrl { get; set; } = __config.Get("insightsQueryUrl") ?? Utilities.GetEnv("NEW_RELIC_INSIGHTS_QUERY_URL") ?? "https://insights-api.newrelic.com/v1/accounts";

        public static string? NerdgraphApiUrl { get; set; } = __config.Get("nerdgraphApiUrl");

        /// <summary>
        /// The data center for which your New Relic account is configured. Only one region per provider block is permitted.
        /// </summary>
        public static string? Region { get; set; } = __config.Get("region") ?? Utilities.GetEnv("NEW_RELIC_REGION") ?? "US";

        public static string? SyntheticsApiUrl { get; set; } = __config.Get("syntheticsApiUrl");

    }
}
