// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this resource to create and manage New Relic dashboards.
 *
 * ## Example Usage
 * ### Create A New Relic Dashboard
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as newrelic from "@pulumi/newrelic";
 *
 * const myApplication = newrelic.getEntity({
 *     name: "My Application",
 *     type: "APPLICATION",
 *     domain: "APM",
 * });
 * const exampledash = new newrelic.Dashboard("exampledash", {
 *     title: "New Relic Terraform Example",
 *     filter: {
 *         eventTypes: ["Transaction"],
 *         attributes: [
 *             "appName",
 *             "name",
 *         ],
 *     },
 *     widgets: [
 *         {
 *             title: "Requests per minute",
 *             visualization: "billboard",
 *             nrql: "SELECT rate(count(*), 1 minute) FROM Transaction",
 *             row: 1,
 *             column: 1,
 *         },
 *         {
 *             title: "Error rate",
 *             visualization: "gauge",
 *             nrql: "SELECT percentage(count(*), WHERE error IS True) FROM Transaction",
 *             thresholdRed: 2.5,
 *             row: 1,
 *             column: 2,
 *         },
 *         {
 *             title: "Average transaction duration, by application",
 *             visualization: "facet_bar_chart",
 *             nrql: "SELECT average(duration) FROM Transaction FACET appName",
 *             row: 1,
 *             column: 3,
 *         },
 *         {
 *             title: "Apdex, top 5 by host",
 *             duration: 1800000,
 *             visualization: "metric_line_chart",
 *             entityIds: [data.newrelic_application.my_application.application_id],
 *             metrics: [{
 *                 name: "Apdex",
 *                 values: ["score"],
 *             }],
 *             facet: "host",
 *             limit: 5,
 *             row: 2,
 *             column: 1,
 *         },
 *         {
 *             title: "Requests per minute, by transaction",
 *             visualization: "facet_table",
 *             nrql: "SELECT rate(count(*), 1 minute) FROM Transaction FACET name",
 *             row: 2,
 *             column: 2,
 *         },
 *         {
 *             title: "Dashboard Note",
 *             visualization: "markdown",
 *             source: `### Helpful Links
 *
 * * [New Relic One](https://one.newrelic.com)
 * * [Developer Portal](https://developer.newrelic.com)`,
 *             row: 2,
 *             column: 3,
 *         },
 *     ],
 * });
 * ```
 * ## Attribute Refence
 *
 * In addition to all arguments above, the following attributes are exported:
 *
 *   * `dashboardUrl` - The URL for viewing the dashboard.
 *
 * ### Nested `widget` blocks
 *
 * All nested `widget` blocks support the following common arguments:
 *
 *   * `title` - (Required) A title for the widget.
 *   * `visualization` - (Required) How the widget visualizes data.  Valid values are `billboard`, `gauge`, `billboardComparison`, `facetBarChart`, `facetedLineChart`, `facetPieChart`, `facetTable`, `facetedAreaChart`, `heatmap`, `attributeSheet`, `singleEvent`, `histogram`, `funnel`, `rawJson`, `eventFeed`, `eventTable`, `uniquesList`, `lineChart`, `comparisonLineChart`, `markdown`, and `metricLineChart`.
 *   * `row` - (Required) Row position of widget from top left, starting at `1`.
 *   * `column` - (Required) Column position of widget from top left, starting at `1`.
 *   * `width` - (Optional) Width of the widget.  Valid values are `1` to `3` inclusive.  Defaults to `1`.
 *   * `height` - (Optional) Height of the widget.  Valid values are `1` to `3` inclusive.  Defaults to `1`.
 *   * `notes` - (Optional) Description of the widget.
 *
 * Each `visualization` type supports an additional set of arguments:
 *
 *   * `billboard`, `billboardComparison`:
 *     * `nrql` - (Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.
 *     * `thresholdRed` - (Optional) Threshold above which the displayed value will be styled with a red color.
 *     * `thresholdYellow` - (Optional) Threshold above which the displayed value will be styled with a yellow color.
 *   * `gauge`:
 *     * `nrql` - (Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.
 *     * `thresholdRed` - (Required) Threshold above which the displayed value will be styled with a red color.
 *     * `thresholdYellow` - (Optional) Threshold above which the displayed value will be styled with a yellow color.
 *   * `facetBarChart`, `facetPieChart`, `facetTable`, `facetedAreaChart`, `facetedLineChart`, or `heatmap`:
 *     * `nrql` - (Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.
 *     * `drilldownDashboardId` - (Optional) The ID of a dashboard to link to from the widget's facets.
 *   * `attributeSheet`, `comparisonLineChart`, `eventFeed`, `eventTable`, `funnel`, `histogram`, `lineChart`, `rawJson`, `singleEvent`, or `uniquesList`:
 *     * `nrql` - (Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.
 *   * `markdown`:
 *     * `source` - (Required) The markdown source to be rendered in the widget.
 *   * `metricLineChart`:
 *     * `entityIds` - (Required) A collection of entity ids to display data for.  These are typically application IDs.
 *     * `metric` - (Required) A nested block that describes a metric.  Nested `metric` blocks support the following arguments:
 *       * `name` - (Required) The metric name to display.
 *       * `values` - (Required) The metric values to display.
 *     * `duration` - (Required) The duration, in ms, of the time window represented in the chart.
 *     * `endTime` - (Optional) The end time of the time window represented in the chart in epoch time.  When not set, the time window will end at the current time.
 *     * `facet` - (Optional) Can be set to "host" to facet the metric data by host.
 *     * `limit` - (Optional) The limit of distinct data series to display.
 *   * `applicationBreakdown`:
 *     * `entityIds` - (Required) A collection of entity IDs to display data. These are typically application IDs.
 *
 * ### Nested `filter` block
 *
 * The optional filter block supports the following arguments:
 *   * `eventTypes` - (Optional) A list of event types to enable filtering for.
 *   * `attributes` - (Optional) A list of attributes belonging to the specified event types to enable filtering for.
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'newrelic:index/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    /**
     * The URL for viewing the dashboard.
     */
    public /*out*/ readonly dashboardUrl!: pulumi.Output<string>;
    /**
     * Determines who can edit the dashboard in an account. Valid values are `all`,  `editableByAll`, `editableByOwner`, or `readOnly`.  Defaults to `editableByAll`.
     */
    public readonly editable!: pulumi.Output<string | undefined>;
    /**
     * A nested block that describes a dashboard filter.  Exactly one nested `filter` block is allowed. See Nested filter block below for details.
     */
    public readonly filter!: pulumi.Output<outputs.DashboardFilter | undefined>;
    /**
     * The number of columns to use when organizing and displaying widgets. New Relic One supports a 3 column grid and a 12 column grid. New Relic Insights supports a 3 column grid.
     */
    public readonly gridColumnCount!: pulumi.Output<number | undefined>;
    /**
     * The icon for the dashboard.  Valid values are `adjust`, `archive`, `bar-chart`, `bell`, `bolt`, `bug`, `bullhorn`, `bullseye`, `clock-o`, `cloud`, `cog`, `comments-o`, `crosshairs`, `dashboard`, `envelope`, `fire`, `flag`, `flask`, `globe`, `heart`, `leaf`, `legal`, `life-ring`, `line-chart`, `magic`, `mobile`, `money`, `none`, `paper-plane`, `pie-chart`, `puzzle-piece`, `road`, `rocket`, `shopping-cart`, `sitemap`, `sliders`, `tablet`, `thumbs-down`, `thumbs-up`, `trophy`, `usd`, `user`, and `users`.  Defaults to `bar-chart`.
     */
    public readonly icon!: pulumi.Output<string | undefined>;
    /**
     * The title of the dashboard.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Determines who can see the dashboard in an account. Valid values are `all` or `owner`.  Defaults to `all`.
     */
    public readonly visibility!: pulumi.Output<string | undefined>;
    /**
     * A nested block that describes a visualization.  Up to 300 `widget` blocks are allowed in a dashboard definition.  See Nested widget blocks below for details.
     */
    public readonly widgets!: pulumi.Output<outputs.DashboardWidget[] | undefined>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DashboardState | undefined;
            inputs["dashboardUrl"] = state ? state.dashboardUrl : undefined;
            inputs["editable"] = state ? state.editable : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["gridColumnCount"] = state ? state.gridColumnCount : undefined;
            inputs["icon"] = state ? state.icon : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["visibility"] = state ? state.visibility : undefined;
            inputs["widgets"] = state ? state.widgets : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["editable"] = args ? args.editable : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["gridColumnCount"] = args ? args.gridColumnCount : undefined;
            inputs["icon"] = args ? args.icon : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
            inputs["widgets"] = args ? args.widgets : undefined;
            inputs["dashboardUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Dashboard.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    /**
     * The URL for viewing the dashboard.
     */
    readonly dashboardUrl?: pulumi.Input<string>;
    /**
     * Determines who can edit the dashboard in an account. Valid values are `all`,  `editableByAll`, `editableByOwner`, or `readOnly`.  Defaults to `editableByAll`.
     */
    readonly editable?: pulumi.Input<string>;
    /**
     * A nested block that describes a dashboard filter.  Exactly one nested `filter` block is allowed. See Nested filter block below for details.
     */
    readonly filter?: pulumi.Input<inputs.DashboardFilter>;
    /**
     * The number of columns to use when organizing and displaying widgets. New Relic One supports a 3 column grid and a 12 column grid. New Relic Insights supports a 3 column grid.
     */
    readonly gridColumnCount?: pulumi.Input<number>;
    /**
     * The icon for the dashboard.  Valid values are `adjust`, `archive`, `bar-chart`, `bell`, `bolt`, `bug`, `bullhorn`, `bullseye`, `clock-o`, `cloud`, `cog`, `comments-o`, `crosshairs`, `dashboard`, `envelope`, `fire`, `flag`, `flask`, `globe`, `heart`, `leaf`, `legal`, `life-ring`, `line-chart`, `magic`, `mobile`, `money`, `none`, `paper-plane`, `pie-chart`, `puzzle-piece`, `road`, `rocket`, `shopping-cart`, `sitemap`, `sliders`, `tablet`, `thumbs-down`, `thumbs-up`, `trophy`, `usd`, `user`, and `users`.  Defaults to `bar-chart`.
     */
    readonly icon?: pulumi.Input<string>;
    /**
     * The title of the dashboard.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * Determines who can see the dashboard in an account. Valid values are `all` or `owner`.  Defaults to `all`.
     */
    readonly visibility?: pulumi.Input<string>;
    /**
     * A nested block that describes a visualization.  Up to 300 `widget` blocks are allowed in a dashboard definition.  See Nested widget blocks below for details.
     */
    readonly widgets?: pulumi.Input<pulumi.Input<inputs.DashboardWidget>[]>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * Determines who can edit the dashboard in an account. Valid values are `all`,  `editableByAll`, `editableByOwner`, or `readOnly`.  Defaults to `editableByAll`.
     */
    readonly editable?: pulumi.Input<string>;
    /**
     * A nested block that describes a dashboard filter.  Exactly one nested `filter` block is allowed. See Nested filter block below for details.
     */
    readonly filter?: pulumi.Input<inputs.DashboardFilter>;
    /**
     * The number of columns to use when organizing and displaying widgets. New Relic One supports a 3 column grid and a 12 column grid. New Relic Insights supports a 3 column grid.
     */
    readonly gridColumnCount?: pulumi.Input<number>;
    /**
     * The icon for the dashboard.  Valid values are `adjust`, `archive`, `bar-chart`, `bell`, `bolt`, `bug`, `bullhorn`, `bullseye`, `clock-o`, `cloud`, `cog`, `comments-o`, `crosshairs`, `dashboard`, `envelope`, `fire`, `flag`, `flask`, `globe`, `heart`, `leaf`, `legal`, `life-ring`, `line-chart`, `magic`, `mobile`, `money`, `none`, `paper-plane`, `pie-chart`, `puzzle-piece`, `road`, `rocket`, `shopping-cart`, `sitemap`, `sliders`, `tablet`, `thumbs-down`, `thumbs-up`, `trophy`, `usd`, `user`, and `users`.  Defaults to `bar-chart`.
     */
    readonly icon?: pulumi.Input<string>;
    /**
     * The title of the dashboard.
     */
    readonly title: pulumi.Input<string>;
    /**
     * Determines who can see the dashboard in an account. Valid values are `all` or `owner`.  Defaults to `all`.
     */
    readonly visibility?: pulumi.Input<string>;
    /**
     * A nested block that describes a visualization.  Up to 300 `widget` blocks are allowed in a dashboard definition.  See Nested widget blocks below for details.
     */
    readonly widgets?: pulumi.Input<pulumi.Input<inputs.DashboardWidget>[]>;
}
