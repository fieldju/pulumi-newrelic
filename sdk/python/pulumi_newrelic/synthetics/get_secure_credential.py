# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSecureCredentialResult:
    """
    A collection of values returned by getSecureCredential.
    """
    def __init__(__self__, created_at=None, description=None, key=None, last_updated=None, id=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        The time the secure credential was created.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The secure credential's description.
        """
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        __self__.key = key
        if last_updated and not isinstance(last_updated, str):
            raise TypeError("Expected argument 'last_updated' to be a str")
        __self__.last_updated = last_updated
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetSecureCredentialResult(GetSecureCredentialResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecureCredentialResult(
            created_at=self.created_at,
            description=self.description,
            key=self.key,
            last_updated=self.last_updated,
            id=self.id)

def get_secure_credential(key=None,opts=None):
    """
    Use this data source to get information about a specific Synthetics secure credential in New Relic that already exists.
    
    Note that the secure credential's value is not returned as an attribute for security reasons.
    
    :param str key: The secure credential's key name.  Regardless of the case used in the configuration, the provider will provide an upcased key to the underlying API.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-newrelic/blob/master/website/docs/d/synthetics_secure_credential.html.markdown.
    """
    __args__ = dict()

    __args__['key'] = key
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('newrelic:synthetics/getSecureCredential:getSecureCredential', __args__, opts=opts).value

    return AwaitableGetSecureCredentialResult(
        created_at=__ret__.get('createdAt'),
        description=__ret__.get('description'),
        key=__ret__.get('key'),
        last_updated=__ret__.get('lastUpdated'),
        id=__ret__.get('id'))