# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import _utilities, _tables


class InfraAlertCondition(pulumi.CustomResource):
    comparison: pulumi.Output[str]
    """
    The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infra_metric` and `infra_process_running` condition types.
    """
    created_at: pulumi.Output[float]
    """
    The timestamp the alert condition was created.
    """
    critical: pulumi.Output[dict]
    """
    Identifies the threshold parameters for opening a critical alert violation. See Thresholds below for details.

      * `duration` (`float`)
      * `timeFunction` (`str`)
      * `value` (`float`)
    """
    enabled: pulumi.Output[bool]
    """
    Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
    """
    event: pulumi.Output[str]
    """
    The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infra_metric` condition type.
    """
    integration_provider: pulumi.Output[str]
    """
    For alerts on integrations, use this instead of `event`.  Supported by the `infra_metric` condition type.
    """
    name: pulumi.Output[str]
    """
    The Infrastructure alert condition's name.
    """
    policy_id: pulumi.Output[float]
    """
    The ID of the alert policy where this condition should be used.
    """
    process_where: pulumi.Output[str]
    """
    Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infra_process_running` condition type.
    """
    runbook_url: pulumi.Output[str]
    """
    Runbook URL to display in notifications.
    """
    select: pulumi.Output[str]
    """
    The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infra_metric` condition type.
    """
    type: pulumi.Output[str]
    """
    The type of Infrastructure alert condition.  Valid values are  `infra_process_running`, `infra_metric`, and `infra_host_not_reporting`.
    """
    updated_at: pulumi.Output[float]
    """
    The timestamp the alert condition was last updated.
    """
    violation_close_timer: pulumi.Output[float]
    """
    Determines how much time will pass before a violation is automatically closed. Setting the time limit to 0 prevents a violation from being force-closed.
    """
    warning: pulumi.Output[dict]
    """
    Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.

      * `duration` (`float`)
      * `timeFunction` (`str`)
      * `value` (`float`)
    """
    where: pulumi.Output[str]
    """
    If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.
    """
    def __init__(__self__, resource_name, opts=None, comparison=None, critical=None, enabled=None, event=None, integration_provider=None, name=None, policy_id=None, process_where=None, runbook_url=None, select=None, type=None, violation_close_timer=None, warning=None, where=None, __props__=None, __name__=None, __opts__=None):
        """
        Use this resource to create and manage Infrastructure alert conditions in New Relic.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_newrelic as newrelic

        foo = newrelic.AlertPolicy("foo")
        high_disk_usage = newrelic.InfraAlertCondition("highDiskUsage",
            policy_id=foo.id,
            type="infra_metric",
            event="StorageSample",
            select="diskUsedPercent",
            comparison="above",
            where="(`hostname` LIKE '%frontend%')",
            critical={
                "duration": 25,
                "value": 90,
                "timeFunction": "all",
            },
            warning={
                "duration": 10,
                "value": 80,
                "timeFunction": "all",
            })
        high_db_conn_count = newrelic.InfraAlertCondition("highDbConnCount",
            policy_id=foo.id,
            type="infra_metric",
            event="DatastoreSample",
            select="provider.databaseConnections.Average",
            comparison="above",
            where="(`hostname` LIKE '%db%')",
            integration_provider="RdsDbInstance",
            critical={
                "duration": 25,
                "value": 90,
                "timeFunction": "all",
            })
        process_not_running = newrelic.InfraAlertCondition("processNotRunning",
            policy_id=foo.id,
            type="infra_process_running",
            comparison="equal",
            process_where="`commandName` = '/usr/bin/ruby'",
            critical={
                "duration": 5,
                "value": 0,
            })
        host_not_reporting = newrelic.InfraAlertCondition("hostNotReporting",
            policy_id=foo.id,
            type="infra_host_not_reporting",
            where="(`hostname` LIKE '%frontend%')",
            critical={
                "duration": 5,
            })
        ```
        ## Thresholds

        The `critical` and `warning` threshold mapping supports the following arguments:

          * `duration` - (Required) Identifies the number of minutes the threshold must be passed or met for the alert to trigger. Threshold durations must be between 1 and 60 minutes (inclusive).
          * `value` - (Optional) Threshold value, computed against the `comparison` operator. Supported by `infra_metric` and `infra_process_running` alert condition types.
          * `time_function` - (Optional) Indicates if the condition needs to be sustained or to just break the threshold once; `all` or `any`. Supported by the `infra_metric` alert condition type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comparison: The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infra_metric` and `infra_process_running` condition types.
        :param pulumi.Input[dict] critical: Identifies the threshold parameters for opening a critical alert violation. See Thresholds below for details.
        :param pulumi.Input[bool] enabled: Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
        :param pulumi.Input[str] event: The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infra_metric` condition type.
        :param pulumi.Input[str] integration_provider: For alerts on integrations, use this instead of `event`.  Supported by the `infra_metric` condition type.
        :param pulumi.Input[str] name: The Infrastructure alert condition's name.
        :param pulumi.Input[float] policy_id: The ID of the alert policy where this condition should be used.
        :param pulumi.Input[str] process_where: Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infra_process_running` condition type.
        :param pulumi.Input[str] runbook_url: Runbook URL to display in notifications.
        :param pulumi.Input[str] select: The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infra_metric` condition type.
        :param pulumi.Input[str] type: The type of Infrastructure alert condition.  Valid values are  `infra_process_running`, `infra_metric`, and `infra_host_not_reporting`.
        :param pulumi.Input[float] violation_close_timer: Determines how much time will pass before a violation is automatically closed. Setting the time limit to 0 prevents a violation from being force-closed.
        :param pulumi.Input[dict] warning: Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.
        :param pulumi.Input[str] where: If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.

        The **critical** object supports the following:

          * `duration` (`pulumi.Input[float]`)
          * `timeFunction` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[float]`)

        The **warning** object supports the following:

          * `duration` (`pulumi.Input[float]`)
          * `timeFunction` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[float]`)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['comparison'] = comparison
            __props__['critical'] = critical
            __props__['enabled'] = enabled
            __props__['event'] = event
            __props__['integration_provider'] = integration_provider
            __props__['name'] = name
            if policy_id is None:
                raise TypeError("Missing required property 'policy_id'")
            __props__['policy_id'] = policy_id
            __props__['process_where'] = process_where
            __props__['runbook_url'] = runbook_url
            __props__['select'] = select
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['violation_close_timer'] = violation_close_timer
            __props__['warning'] = warning
            __props__['where'] = where
            __props__['created_at'] = None
            __props__['updated_at'] = None
        super(InfraAlertCondition, __self__).__init__(
            'newrelic:index/infraAlertCondition:InfraAlertCondition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, comparison=None, created_at=None, critical=None, enabled=None, event=None, integration_provider=None, name=None, policy_id=None, process_where=None, runbook_url=None, select=None, type=None, updated_at=None, violation_close_timer=None, warning=None, where=None):
        """
        Get an existing InfraAlertCondition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comparison: The operator used to evaluate the threshold value.  Valid values are `above`, `below`, and `equal`.  Supported by the `infra_metric` and `infra_process_running` condition types.
        :param pulumi.Input[float] created_at: The timestamp the alert condition was created.
        :param pulumi.Input[dict] critical: Identifies the threshold parameters for opening a critical alert violation. See Thresholds below for details.
        :param pulumi.Input[bool] enabled: Whether the condition is turned on or off.  Valid values are `true` and `false`.  Defaults to `true`.
        :param pulumi.Input[str] event: The metric event; for example, `SystemSample` or `StorageSample`.  Supported by the `infra_metric` condition type.
        :param pulumi.Input[str] integration_provider: For alerts on integrations, use this instead of `event`.  Supported by the `infra_metric` condition type.
        :param pulumi.Input[str] name: The Infrastructure alert condition's name.
        :param pulumi.Input[float] policy_id: The ID of the alert policy where this condition should be used.
        :param pulumi.Input[str] process_where: Any filters applied to processes; for example: `commandName = 'java'`.  Supported by the `infra_process_running` condition type.
        :param pulumi.Input[str] runbook_url: Runbook URL to display in notifications.
        :param pulumi.Input[str] select: The attribute name to identify the metric being targeted; for example, `cpuPercent`, `diskFreePercent`, or `memoryResidentSizeBytes`.  The underlying API will automatically populate this value for Infrastructure integrations (for example `diskFreePercent`), so make sure to explicitly include this value to avoid diff issues.  Supported by the `infra_metric` condition type.
        :param pulumi.Input[str] type: The type of Infrastructure alert condition.  Valid values are  `infra_process_running`, `infra_metric`, and `infra_host_not_reporting`.
        :param pulumi.Input[float] updated_at: The timestamp the alert condition was last updated.
        :param pulumi.Input[float] violation_close_timer: Determines how much time will pass before a violation is automatically closed. Setting the time limit to 0 prevents a violation from being force-closed.
        :param pulumi.Input[dict] warning: Identifies the threshold parameters for opening a warning alert violation. See Thresholds below for details.
        :param pulumi.Input[str] where: If applicable, this identifies any Infrastructure host filters used; for example: `hostname LIKE '%cassandra%'`.

        The **critical** object supports the following:

          * `duration` (`pulumi.Input[float]`)
          * `timeFunction` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[float]`)

        The **warning** object supports the following:

          * `duration` (`pulumi.Input[float]`)
          * `timeFunction` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["comparison"] = comparison
        __props__["created_at"] = created_at
        __props__["critical"] = critical
        __props__["enabled"] = enabled
        __props__["event"] = event
        __props__["integration_provider"] = integration_provider
        __props__["name"] = name
        __props__["policy_id"] = policy_id
        __props__["process_where"] = process_where
        __props__["runbook_url"] = runbook_url
        __props__["select"] = select
        __props__["type"] = type
        __props__["updated_at"] = updated_at
        __props__["violation_close_timer"] = violation_close_timer
        __props__["warning"] = warning
        __props__["where"] = where
        return InfraAlertCondition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
