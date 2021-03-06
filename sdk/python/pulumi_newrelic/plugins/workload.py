# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Workload(pulumi.CustomResource):
    account_id: pulumi.Output[float]
    """
    The New Relic account ID where you want to create the workload.
    """
    composite_entity_search_query: pulumi.Output[str]
    """
    The composite query used to compose a dynamic workload.
    """
    entity_guids: pulumi.Output[list]
    """
    A list of entity GUIDs manually assigned to this workload.
    """
    entity_search_queries: pulumi.Output[list]
    """
    A list of search queries that define a dynamic workload.  See Nested entity_search_query blocks below for details.

      * `query` (`str`) - The query.
    """
    guid: pulumi.Output[str]
    """
    The unique entity identifier of the workload in New Relic.
    """
    name: pulumi.Output[str]
    """
    The workload's name.
    """
    permalink: pulumi.Output[str]
    """
    The URL of the workload.
    """
    scope_account_ids: pulumi.Output[list]
    """
    A list of account IDs that will be used to get entities from.
    """
    workload_id: pulumi.Output[float]
    """
    The unique entity identifier of the workload.
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, entity_guids=None, entity_search_queries=None, name=None, scope_account_ids=None, __props__=None, __name__=None, __opts__=None):
        """
        Use this resource to create, update, and delete a New Relic One workload.

        A New Relic Personal API key is required to provision this resource.  Set the `api_key`
        attribute in the `provider` block or the `NEW_RELIC_API_KEY` environment
        variable with your Personal API key.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_newrelic as newrelic

        foo = newrelic.plugins.Workload("foo",
            account_id=12345678,
            entity_guids=["MjUyMDUyOHxBUE18QVBQTElDQVRJT058MjE1MDM3Nzk1"],
            entity_search_queries=[{
                "query": "name like 'Example application'",
            }],
            scope_account_ids=[12345678])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] account_id: The New Relic account ID where you want to create the workload.
        :param pulumi.Input[list] entity_guids: A list of entity GUIDs manually assigned to this workload.
        :param pulumi.Input[list] entity_search_queries: A list of search queries that define a dynamic workload.  See Nested entity_search_query blocks below for details.
        :param pulumi.Input[str] name: The workload's name.
        :param pulumi.Input[list] scope_account_ids: A list of account IDs that will be used to get entities from.

        The **entity_search_queries** object supports the following:

          * `query` (`pulumi.Input[str]`) - The query.
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

            if account_id is None:
                raise TypeError("Missing required property 'account_id'")
            __props__['account_id'] = account_id
            __props__['entity_guids'] = entity_guids
            __props__['entity_search_queries'] = entity_search_queries
            __props__['name'] = name
            __props__['scope_account_ids'] = scope_account_ids
            __props__['composite_entity_search_query'] = None
            __props__['guid'] = None
            __props__['permalink'] = None
            __props__['workload_id'] = None
        super(Workload, __self__).__init__(
            'newrelic:plugins/workload:Workload',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_id=None, composite_entity_search_query=None, entity_guids=None, entity_search_queries=None, guid=None, name=None, permalink=None, scope_account_ids=None, workload_id=None):
        """
        Get an existing Workload resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] account_id: The New Relic account ID where you want to create the workload.
        :param pulumi.Input[str] composite_entity_search_query: The composite query used to compose a dynamic workload.
        :param pulumi.Input[list] entity_guids: A list of entity GUIDs manually assigned to this workload.
        :param pulumi.Input[list] entity_search_queries: A list of search queries that define a dynamic workload.  See Nested entity_search_query blocks below for details.
        :param pulumi.Input[str] guid: The unique entity identifier of the workload in New Relic.
        :param pulumi.Input[str] name: The workload's name.
        :param pulumi.Input[str] permalink: The URL of the workload.
        :param pulumi.Input[list] scope_account_ids: A list of account IDs that will be used to get entities from.
        :param pulumi.Input[float] workload_id: The unique entity identifier of the workload.

        The **entity_search_queries** object supports the following:

          * `query` (`pulumi.Input[str]`) - The query.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["composite_entity_search_query"] = composite_entity_search_query
        __props__["entity_guids"] = entity_guids
        __props__["entity_search_queries"] = entity_search_queries
        __props__["guid"] = guid
        __props__["name"] = name
        __props__["permalink"] = permalink
        __props__["scope_account_ids"] = scope_account_ids
        __props__["workload_id"] = workload_id
        return Workload(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
