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
    'GetInstancePrivateNicResult',
    'AwaitableGetInstancePrivateNicResult',
    'get_instance_private_nic',
    'get_instance_private_nic_output',
]

@pulumi.output_type
class GetInstancePrivateNicResult:
    """
    A collection of values returned by getInstancePrivateNic.
    """
    def __init__(__self__, id=None, ip_ids=None, mac_address=None, private_network_id=None, private_nic_id=None, server_id=None, tags=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_ids and not isinstance(ip_ids, list):
            raise TypeError("Expected argument 'ip_ids' to be a list")
        pulumi.set(__self__, "ip_ids", ip_ids)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if private_network_id and not isinstance(private_network_id, str):
            raise TypeError("Expected argument 'private_network_id' to be a str")
        pulumi.set(__self__, "private_network_id", private_network_id)
        if private_nic_id and not isinstance(private_nic_id, str):
            raise TypeError("Expected argument 'private_nic_id' to be a str")
        pulumi.set(__self__, "private_nic_id", private_nic_id)
        if server_id and not isinstance(server_id, str):
            raise TypeError("Expected argument 'server_id' to be a str")
        pulumi.set(__self__, "server_id", server_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipIds")
    def ip_ids(self) -> Sequence[str]:
        return pulumi.get(self, "ip_ids")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[str]:
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter(name="privateNicId")
    def private_nic_id(self) -> Optional[str]:
        return pulumi.get(self, "private_nic_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> str:
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetInstancePrivateNicResult(GetInstancePrivateNicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancePrivateNicResult(
            id=self.id,
            ip_ids=self.ip_ids,
            mac_address=self.mac_address,
            private_network_id=self.private_network_id,
            private_nic_id=self.private_nic_id,
            server_id=self.server_id,
            tags=self.tags,
            zone=self.zone)


def get_instance_private_nic(private_network_id: Optional[str] = None,
                             private_nic_id: Optional[str] = None,
                             server_id: Optional[str] = None,
                             tags: Optional[Sequence[str]] = None,
                             zone: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancePrivateNicResult:
    """
    Gets information about an instance private NIC.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_nic_id = scaleway.get_instance_private_nic(private_nic_id="11111111-1111-1111-1111-111111111111",
        server_id="11111111-1111-1111-1111-111111111111")
    by_pn_id = scaleway.get_instance_private_nic(private_network_id="11111111-1111-1111-1111-111111111111",
        server_id="11111111-1111-1111-1111-111111111111")
    by_tags = scaleway.get_instance_private_nic(server_id="11111111-1111-1111-1111-111111111111",
        tags=["mytag"])
    ```
    <!--End PulumiCodeChooser -->


    :param str private_network_id: The ID of the private network
           Only one of `private_nic_id` and `private_network_id` should be specified.
    :param str private_nic_id: The ID of the instance server private nic
           Only one of `private_nic_id` and `private_network_id` should be specified.
    :param str server_id: The server's id
    :param Sequence[str] tags: The tags associated with the private NIC.
           As datasource only returns one private NIC, the search with given tags must return only one result
    :param str zone: `zone`) The zone in which the private nic exists.
    """
    __args__ = dict()
    __args__['privateNetworkId'] = private_network_id
    __args__['privateNicId'] = private_nic_id
    __args__['serverId'] = server_id
    __args__['tags'] = tags
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getInstancePrivateNic:getInstancePrivateNic', __args__, opts=opts, typ=GetInstancePrivateNicResult).value

    return AwaitableGetInstancePrivateNicResult(
        id=pulumi.get(__ret__, 'id'),
        ip_ids=pulumi.get(__ret__, 'ip_ids'),
        mac_address=pulumi.get(__ret__, 'mac_address'),
        private_network_id=pulumi.get(__ret__, 'private_network_id'),
        private_nic_id=pulumi.get(__ret__, 'private_nic_id'),
        server_id=pulumi.get(__ret__, 'server_id'),
        tags=pulumi.get(__ret__, 'tags'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_instance_private_nic)
def get_instance_private_nic_output(private_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    private_nic_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    server_id: Optional[pulumi.Input[str]] = None,
                                    tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    zone: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancePrivateNicResult]:
    """
    Gets information about an instance private NIC.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_nic_id = scaleway.get_instance_private_nic(private_nic_id="11111111-1111-1111-1111-111111111111",
        server_id="11111111-1111-1111-1111-111111111111")
    by_pn_id = scaleway.get_instance_private_nic(private_network_id="11111111-1111-1111-1111-111111111111",
        server_id="11111111-1111-1111-1111-111111111111")
    by_tags = scaleway.get_instance_private_nic(server_id="11111111-1111-1111-1111-111111111111",
        tags=["mytag"])
    ```
    <!--End PulumiCodeChooser -->


    :param str private_network_id: The ID of the private network
           Only one of `private_nic_id` and `private_network_id` should be specified.
    :param str private_nic_id: The ID of the instance server private nic
           Only one of `private_nic_id` and `private_network_id` should be specified.
    :param str server_id: The server's id
    :param Sequence[str] tags: The tags associated with the private NIC.
           As datasource only returns one private NIC, the search with given tags must return only one result
    :param str zone: `zone`) The zone in which the private nic exists.
    """
    ...
