# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    """
    A collection of values returned by getDatabase.
    """
    def __init__(__self__, id=None, instance_id=None, managed=None, name=None, owner=None, region=None, size=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if managed and not isinstance(managed, bool):
            raise TypeError("Expected argument 'managed' to be a bool")
        pulumi.set(__self__, "managed", managed)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if size and not isinstance(size, str):
            raise TypeError("Expected argument 'size' to be a str")
        pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def managed(self) -> bool:
        """
        Whether the database is managed or not.
        """
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The name of the owner of the database.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> str:
        """
        Size of the database (in bytes).
        """
        return pulumi.get(self, "size")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            id=self.id,
            instance_id=self.instance_id,
            managed=self.managed,
            name=self.name,
            owner=self.owner,
            region=self.region,
            size=self.size)


def get_database(instance_id: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Gets information about a database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_db = scaleway.get_database(instance_id="11111111-1111-1111-1111-111111111111",
        name="foobar")
    ```


    :param str instance_id: The RDB instance ID.
    :param str name: The name of the RDB instance.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getDatabase:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        managed=pulumi.get(__ret__, 'managed'),
        name=pulumi.get(__ret__, 'name'),
        owner=pulumi.get(__ret__, 'owner'),
        region=pulumi.get(__ret__, 'region'),
        size=pulumi.get(__ret__, 'size'))


@_utilities.lift_output_func(get_database)
def get_database_output(instance_id: Optional[pulumi.Input[str]] = None,
                        name: Optional[pulumi.Input[str]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Gets information about a database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_db = scaleway.get_database(instance_id="11111111-1111-1111-1111-111111111111",
        name="foobar")
    ```


    :param str instance_id: The RDB instance ID.
    :param str name: The name of the RDB instance.
    """
    ...
