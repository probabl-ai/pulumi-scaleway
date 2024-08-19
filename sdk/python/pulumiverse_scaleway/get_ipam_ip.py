# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetIpamIpResult',
    'AwaitableGetIpamIpResult',
    'get_ipam_ip',
    'get_ipam_ip_output',
]

@pulumi.output_type
class GetIpamIpResult:
    """
    A collection of values returned by getIpamIp.
    """
    def __init__(__self__, address=None, address_cidr=None, attached=None, id=None, ipam_ip_id=None, mac_address=None, organization_id=None, private_network_id=None, project_id=None, region=None, resource=None, tags=None, type=None, zonal=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if address_cidr and not isinstance(address_cidr, str):
            raise TypeError("Expected argument 'address_cidr' to be a str")
        pulumi.set(__self__, "address_cidr", address_cidr)
        if attached and not isinstance(attached, bool):
            raise TypeError("Expected argument 'attached' to be a bool")
        pulumi.set(__self__, "attached", attached)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipam_ip_id and not isinstance(ipam_ip_id, str):
            raise TypeError("Expected argument 'ipam_ip_id' to be a str")
        pulumi.set(__self__, "ipam_ip_id", ipam_ip_id)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if private_network_id and not isinstance(private_network_id, str):
            raise TypeError("Expected argument 'private_network_id' to be a str")
        pulumi.set(__self__, "private_network_id", private_network_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if resource and not isinstance(resource, dict):
            raise TypeError("Expected argument 'resource' to be a dict")
        pulumi.set(__self__, "resource", resource)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zonal and not isinstance(zonal, str):
            raise TypeError("Expected argument 'zonal' to be a str")
        pulumi.set(__self__, "zonal", zonal)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The IP address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="addressCidr")
    def address_cidr(self) -> str:
        """
        the IP address in CIDR notation.
        """
        return pulumi.get(self, "address_cidr")

    @property
    @pulumi.getter
    def attached(self) -> Optional[bool]:
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipamIpId")
    def ipam_ip_id(self) -> Optional[str]:
        return pulumi.get(self, "ipam_ip_id")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[str]:
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def resource(self) -> Optional['outputs.GetIpamIpResourceResult']:
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zonal(self) -> str:
        return pulumi.get(self, "zonal")


class AwaitableGetIpamIpResult(GetIpamIpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpamIpResult(
            address=self.address,
            address_cidr=self.address_cidr,
            attached=self.attached,
            id=self.id,
            ipam_ip_id=self.ipam_ip_id,
            mac_address=self.mac_address,
            organization_id=self.organization_id,
            private_network_id=self.private_network_id,
            project_id=self.project_id,
            region=self.region,
            resource=self.resource,
            tags=self.tags,
            type=self.type,
            zonal=self.zonal)


def get_ipam_ip(attached: Optional[bool] = None,
                ipam_ip_id: Optional[str] = None,
                mac_address: Optional[str] = None,
                private_network_id: Optional[str] = None,
                project_id: Optional[str] = None,
                region: Optional[str] = None,
                resource: Optional[Union['GetIpamIpResourceArgs', 'GetIpamIpResourceArgsDict']] = None,
                tags: Optional[Sequence[str]] = None,
                type: Optional[str] = None,
                zonal: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpamIpResult:
    """
    Gets information about IP addresses managed by Scaleway's IP Address Management (IPAM) service. IPAM is used for the DHCP bundled with VPC Private Networks.

    For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/network/vpc/concepts/#ipam).

    ## Examples

    ### IPAM IP ID

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_id = scaleway.get_ipam_ip(ipam_ip_id="11111111-1111-1111-1111-111111111111")
    ```

    ### Instance Private Network IP

    Get an Instance's IP on a Private Network.

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    # Connect your instance to a private network using a private nic.
    nic = scaleway.InstancePrivateNic("nic",
        server_id=scaleway_instance_server["server"]["id"],
        private_network_id=scaleway_vpc_private_network["pn"]["id"])
    by_mac = scaleway.get_ipam_ip_output(mac_address=nic.mac_address,
        type="ipv4")
    by_id = scaleway.get_ipam_ip_output(resource={
            "id": nic.id,
            "type": "instance_private_nic",
        },
        type="ipv4")
    ```

    ### RDB instance

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    # Find the private IPv4 using resource name
    pn = scaleway.VpcPrivateNetwork("pn")
    main = scaleway.DatabaseInstance("main",
        node_type="DB-DEV-S",
        engine="PostgreSQL-15",
        is_ha_cluster=True,
        disable_backup=True,
        user_name="my_initial_user",
        password="thiZ_is_v&ry_s3cret",
        private_network={
            "pn_id": pn.id,
        })
    by_name = scaleway.get_ipam_ip_output(resource={
            "name": main.name,
            "type": "rdb_instance",
        },
        type="ipv4")
    ```


    :param bool attached: Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipam_ip_id`.
    :param str ipam_ip_id: The IPAM IP ID. Cannot be used with any other arguments.
    :param str mac_address: The MAC address linked to the IP. Cannot be used with `ipam_ip_id`.
    :param str private_network_id: The ID of the Private Network the IP belongs to. Cannot be used with `ipam_ip_id`.
    :param str project_id: `project_id`) The ID of the Project the IP is associated with.
    :param str region: `region`) The region in which the IP exists.
    :param Union['GetIpamIpResourceArgs', 'GetIpamIpResourceArgsDict'] resource: Filter by resource ID, type or name. Cannot be used with `ipam_ip_id`.
           If specified, `type` is required, and at least one of `id` or `name` must be set.
    :param Sequence[str] tags: The tags associated with the IP. Cannot be used with `ipam_ip_id`.
           As datasource only returns one IP, the search with given tags must return only one result.
    :param str type: The type of IP to search for (`ipv4` or `ipv6`). Cannot be used with `ipam_ip_id`.
    :param str zonal: Only IPs that are zonal, and in this zone, will be returned.
    """
    __args__ = dict()
    __args__['attached'] = attached
    __args__['ipamIpId'] = ipam_ip_id
    __args__['macAddress'] = mac_address
    __args__['privateNetworkId'] = private_network_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['resource'] = resource
    __args__['tags'] = tags
    __args__['type'] = type
    __args__['zonal'] = zonal
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getIpamIp:getIpamIp', __args__, opts=opts, typ=GetIpamIpResult).value

    return AwaitableGetIpamIpResult(
        address=pulumi.get(__ret__, 'address'),
        address_cidr=pulumi.get(__ret__, 'address_cidr'),
        attached=pulumi.get(__ret__, 'attached'),
        id=pulumi.get(__ret__, 'id'),
        ipam_ip_id=pulumi.get(__ret__, 'ipam_ip_id'),
        mac_address=pulumi.get(__ret__, 'mac_address'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        private_network_id=pulumi.get(__ret__, 'private_network_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        resource=pulumi.get(__ret__, 'resource'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        zonal=pulumi.get(__ret__, 'zonal'))


@_utilities.lift_output_func(get_ipam_ip)
def get_ipam_ip_output(attached: Optional[pulumi.Input[Optional[bool]]] = None,
                       ipam_ip_id: Optional[pulumi.Input[Optional[str]]] = None,
                       mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                       private_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                       project_id: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       resource: Optional[pulumi.Input[Optional[Union['GetIpamIpResourceArgs', 'GetIpamIpResourceArgsDict']]]] = None,
                       tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       type: Optional[pulumi.Input[Optional[str]]] = None,
                       zonal: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpamIpResult]:
    """
    Gets information about IP addresses managed by Scaleway's IP Address Management (IPAM) service. IPAM is used for the DHCP bundled with VPC Private Networks.

    For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/network/vpc/concepts/#ipam).

    ## Examples

    ### IPAM IP ID

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_id = scaleway.get_ipam_ip(ipam_ip_id="11111111-1111-1111-1111-111111111111")
    ```

    ### Instance Private Network IP

    Get an Instance's IP on a Private Network.

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    # Connect your instance to a private network using a private nic.
    nic = scaleway.InstancePrivateNic("nic",
        server_id=scaleway_instance_server["server"]["id"],
        private_network_id=scaleway_vpc_private_network["pn"]["id"])
    by_mac = scaleway.get_ipam_ip_output(mac_address=nic.mac_address,
        type="ipv4")
    by_id = scaleway.get_ipam_ip_output(resource={
            "id": nic.id,
            "type": "instance_private_nic",
        },
        type="ipv4")
    ```

    ### RDB instance

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    # Find the private IPv4 using resource name
    pn = scaleway.VpcPrivateNetwork("pn")
    main = scaleway.DatabaseInstance("main",
        node_type="DB-DEV-S",
        engine="PostgreSQL-15",
        is_ha_cluster=True,
        disable_backup=True,
        user_name="my_initial_user",
        password="thiZ_is_v&ry_s3cret",
        private_network={
            "pn_id": pn.id,
        })
    by_name = scaleway.get_ipam_ip_output(resource={
            "name": main.name,
            "type": "rdb_instance",
        },
        type="ipv4")
    ```


    :param bool attached: Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipam_ip_id`.
    :param str ipam_ip_id: The IPAM IP ID. Cannot be used with any other arguments.
    :param str mac_address: The MAC address linked to the IP. Cannot be used with `ipam_ip_id`.
    :param str private_network_id: The ID of the Private Network the IP belongs to. Cannot be used with `ipam_ip_id`.
    :param str project_id: `project_id`) The ID of the Project the IP is associated with.
    :param str region: `region`) The region in which the IP exists.
    :param Union['GetIpamIpResourceArgs', 'GetIpamIpResourceArgsDict'] resource: Filter by resource ID, type or name. Cannot be used with `ipam_ip_id`.
           If specified, `type` is required, and at least one of `id` or `name` must be set.
    :param Sequence[str] tags: The tags associated with the IP. Cannot be used with `ipam_ip_id`.
           As datasource only returns one IP, the search with given tags must return only one result.
    :param str type: The type of IP to search for (`ipv4` or `ipv6`). Cannot be used with `ipam_ip_id`.
    :param str zonal: Only IPs that are zonal, and in this zone, will be returned.
    """
    ...
