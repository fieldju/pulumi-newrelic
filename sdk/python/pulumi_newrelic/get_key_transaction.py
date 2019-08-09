# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetKeyTransactionResult:
    """
    A collection of values returned by getKeyTransaction.
    """
    def __init__(__self__, name=None, id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetKeyTransactionResult(GetKeyTransactionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyTransactionResult(
            name=self.name,
            id=self.id)

def get_key_transaction(name=None,opts=None):
    """
    Use this data source to get information about a specific key transaction in New Relic.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/d/key_transaction.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('newrelic:index/getKeyTransaction:getKeyTransaction', __args__, opts=opts).value

    return AwaitableGetKeyTransactionResult(
        name=__ret__.get('name'),
        id=__ret__.get('id'))
