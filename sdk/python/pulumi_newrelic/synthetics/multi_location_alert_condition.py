# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class MultiLocationAlertCondition(pulumi.CustomResource):
    critical: pulumi.Output[dict]
    """
    A condition term with the priority set to critical.

      * `threshold` (`float`)
    """
    enabled: pulumi.Output[bool]
    """
    Set whether to enable the alert condition.  Defaults to true.
    """
    entities: pulumi.Output[list]
    """
    The GUIDs of the Synthetics monitors to alert on.
    """
    name: pulumi.Output[str]
    """
    The title of the condition.
    """
    policy_id: pulumi.Output[float]
    """
    The ID of the policy where this condition will be used.
    """
    runbook_url: pulumi.Output[str]
    """
    Runbook URL to display in notifications.
    """
    violation_time_limit_seconds: pulumi.Output[float]
    """
    The maximum number of seconds a violation can remain open before being closed by the system. Must be one of: 0, 3600,
    7200, 14400, 28800, 43200, 86400
    """
    warning: pulumi.Output[dict]
    """
    A condition term with the priority set to warning.

      * `threshold` (`float`)
    """
    def __init__(__self__, resource_name, opts=None, critical=None, enabled=None, entities=None, name=None, policy_id=None, runbook_url=None, violation_time_limit_seconds=None, warning=None, __props__=None, __name__=None, __opts__=None):
        """
        Use this resource to create, update, and delete a New Relic Synthetics Location Alerts.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] critical: A condition term with the priority set to critical.
        :param pulumi.Input[bool] enabled: Set whether to enable the alert condition.  Defaults to true.
        :param pulumi.Input[list] entities: The GUIDs of the Synthetics monitors to alert on.
        :param pulumi.Input[str] name: The title of the condition.
        :param pulumi.Input[float] policy_id: The ID of the policy where this condition will be used.
        :param pulumi.Input[str] runbook_url: Runbook URL to display in notifications.
        :param pulumi.Input[float] violation_time_limit_seconds: The maximum number of seconds a violation can remain open before being closed by the system. Must be one of: 0, 3600,
               7200, 14400, 28800, 43200, 86400
        :param pulumi.Input[dict] warning: A condition term with the priority set to warning.

        The **critical** object supports the following:

          * `threshold` (`pulumi.Input[float]`)

        The **warning** object supports the following:

          * `threshold` (`pulumi.Input[float]`)
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if critical is None:
                raise TypeError("Missing required property 'critical'")
            __props__['critical'] = critical
            __props__['enabled'] = enabled
            if entities is None:
                raise TypeError("Missing required property 'entities'")
            __props__['entities'] = entities
            __props__['name'] = name
            if policy_id is None:
                raise TypeError("Missing required property 'policy_id'")
            __props__['policy_id'] = policy_id
            __props__['runbook_url'] = runbook_url
            if violation_time_limit_seconds is None:
                raise TypeError("Missing required property 'violation_time_limit_seconds'")
            __props__['violation_time_limit_seconds'] = violation_time_limit_seconds
            __props__['warning'] = warning
        super(MultiLocationAlertCondition, __self__).__init__(
            'newrelic:synthetics/multiLocationAlertCondition:MultiLocationAlertCondition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, critical=None, enabled=None, entities=None, name=None, policy_id=None, runbook_url=None, violation_time_limit_seconds=None, warning=None):
        """
        Get an existing MultiLocationAlertCondition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] critical: A condition term with the priority set to critical.
        :param pulumi.Input[bool] enabled: Set whether to enable the alert condition.  Defaults to true.
        :param pulumi.Input[list] entities: The GUIDs of the Synthetics monitors to alert on.
        :param pulumi.Input[str] name: The title of the condition.
        :param pulumi.Input[float] policy_id: The ID of the policy where this condition will be used.
        :param pulumi.Input[str] runbook_url: Runbook URL to display in notifications.
        :param pulumi.Input[float] violation_time_limit_seconds: The maximum number of seconds a violation can remain open before being closed by the system. Must be one of: 0, 3600,
               7200, 14400, 28800, 43200, 86400
        :param pulumi.Input[dict] warning: A condition term with the priority set to warning.

        The **critical** object supports the following:

          * `threshold` (`pulumi.Input[float]`)

        The **warning** object supports the following:

          * `threshold` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["critical"] = critical
        __props__["enabled"] = enabled
        __props__["entities"] = entities
        __props__["name"] = name
        __props__["policy_id"] = policy_id
        __props__["runbook_url"] = runbook_url
        __props__["violation_time_limit_seconds"] = violation_time_limit_seconds
        __props__["warning"] = warning
        return MultiLocationAlertCondition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
