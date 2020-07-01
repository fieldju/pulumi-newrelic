# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetKeyTransactionResult:
    """
    A collection of values returned by getKeyTransaction.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetKeyTransactionResult(GetKeyTransactionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyTransactionResult(
            id=self.id,
            name=self.name)

def get_key_transaction(name=None,opts=None):
    """
    Use this data source to get information about a specific key transaction in New Relic that already exists.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_newrelic as newrelic

    txn = newrelic.get_key_transaction(name="txn")
    foo_alert_policy = newrelic.AlertPolicy("fooAlertPolicy")
    foo_alert_condition = newrelic.AlertCondition("fooAlertCondition",
        policy_id=foo_alert_policy.id,
        type="apm_kt_metric",
        entities=[txn.id],
        metric="error_percentage",
        runbook_url="https://www.example.com",
        terms=[{
            "duration": 5,
            "operator": "below",
            "priority": "critical",
            "threshold": "0.75",
            "timeFunction": "all",
        }])
    ```


    :param str name: The name of the key transaction in New Relic.
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('newrelic:index/getKeyTransaction:getKeyTransaction', __args__, opts=opts).value

    return AwaitableGetKeyTransactionResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'))
