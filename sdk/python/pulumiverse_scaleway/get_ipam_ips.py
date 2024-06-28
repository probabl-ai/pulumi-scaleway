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
    'GetIpamIpsResult',
    'AwaitableGetIpamIpsResult',
    'get_ipam_ips',
    'get_ipam_ips_output',
]

@pulumi.output_type
class GetIpamIpsResult:
    """
    A collection of values returned by getIpamIps.
    """
    def __init__(__self__, attached=None, id=None, ips=None, mac_address=None, organization_id=None, private_network_id=None, project_id=None, region=None, resource=None, tags=None, type=None, zonal=None):
        if attached and not isinstance(attached, bool):
            raise TypeError("Expected argument 'attached' to be a bool")
        pulumi.set(__self__, "attached", attached)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
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
    @pulumi.getter
    def ips(self) -> Sequence['outputs.GetIpamIpsIpResult']:
        """
        List of found IPs.
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        """
        The associated MAC address.
        """
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
        """
        The ID of the Project the resource is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region of the IP.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def resource(self) -> Optional['outputs.GetIpamIpsResourceResult']:
        """
        The list of public IPs attached to the resource.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The tags associated with the IP.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zonal(self) -> str:
        return pulumi.get(self, "zonal")


class AwaitableGetIpamIpsResult(GetIpamIpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpamIpsResult(
            attached=self.attached,
            id=self.id,
            ips=self.ips,
            mac_address=self.mac_address,
            organization_id=self.organization_id,
            private_network_id=self.private_network_id,
            project_id=self.project_id,
            region=self.region,
            resource=self.resource,
            tags=self.tags,
            type=self.type,
            zonal=self.zonal)


def get_ipam_ips(attached: Optional[bool] = None,
                 mac_address: Optional[str] = None,
                 private_network_id: Optional[str] = None,
                 project_id: Optional[str] = None,
                 region: Optional[str] = None,
                 resource: Optional[pulumi.InputType['GetIpamIpsResourceArgs']] = None,
                 tags: Optional[Sequence[str]] = None,
                 type: Optional[str] = None,
                 zonal: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpamIpsResult:
    """
    Gets information about multiple IP addresses managed by Scaleway's IP Address Management (IPAM) service.

    For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/network/vpc/concepts/#ipam).

    ## Examples

    ### By tag

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_tag = scaleway.get_ipam_ips(tags=["tag"])
    ```

    ### By type and resource

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    vpc01 = scaleway.Vpc("vpc01")
    pn01 = scaleway.VpcPrivateNetwork("pn01",
        vpc_id=vpc01.id,
        ipv4_subnet=scaleway.VpcPrivateNetworkIpv4SubnetArgs(
            subnet="172.16.32.0/22",
        ))
    redis01 = scaleway.RedisCluster("redis01",
        version="7.0.5",
        node_type="RED1-XS",
        user_name="my_initial_user",
        password="thiZ_is_v&ry_s3cret",
        cluster_size=3,
        private_networks=[scaleway.RedisClusterPrivateNetworkArgs(
            id=pn01.id,
        )])
    by_type_and_resource = scaleway.get_ipam_ips_output(type="ipv4",
        resource=scaleway.GetIpamIpsResourceArgs(
            id=redis01.id,
            type="redis_cluster",
        ))
    ```


    :param bool attached: Defines whether to filter only for IPs which are attached to a resource.
    :param str mac_address: The linked MAC address to filter for.
    :param str private_network_id: The ID of the Private Network to filter for.
    :param str project_id: The ID of the Project to filter for.
    :param str region: The region to filter for.
    :param pulumi.InputType['GetIpamIpsResourceArgs'] resource: Filter for a resource attached to the IP, using resource ID, type or name.
    :param Sequence[str] tags: The IP tags to filter for.
    :param str type: The type of IP to filter for (`ipv4` or `ipv6`).
    :param str zonal: Only IPs that are zonal, and in this zone, will be returned.
    """
    __args__ = dict()
    __args__['attached'] = attached
    __args__['macAddress'] = mac_address
    __args__['privateNetworkId'] = private_network_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['resource'] = resource
    __args__['tags'] = tags
    __args__['type'] = type
    __args__['zonal'] = zonal
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getIpamIps:getIpamIps', __args__, opts=opts, typ=GetIpamIpsResult).value

    return AwaitableGetIpamIpsResult(
        attached=pulumi.get(__ret__, 'attached'),
        id=pulumi.get(__ret__, 'id'),
        ips=pulumi.get(__ret__, 'ips'),
        mac_address=pulumi.get(__ret__, 'mac_address'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        private_network_id=pulumi.get(__ret__, 'private_network_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        resource=pulumi.get(__ret__, 'resource'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        zonal=pulumi.get(__ret__, 'zonal'))


@_utilities.lift_output_func(get_ipam_ips)
def get_ipam_ips_output(attached: Optional[pulumi.Input[Optional[bool]]] = None,
                        mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                        private_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                        project_id: Optional[pulumi.Input[Optional[str]]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        resource: Optional[pulumi.Input[Optional[pulumi.InputType['GetIpamIpsResourceArgs']]]] = None,
                        tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        type: Optional[pulumi.Input[Optional[str]]] = None,
                        zonal: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpamIpsResult]:
    """
    Gets information about multiple IP addresses managed by Scaleway's IP Address Management (IPAM) service.

    For more information about IPAM, see the main [documentation](https://www.scaleway.com/en/docs/network/vpc/concepts/#ipam).

    ## Examples

    ### By tag

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_tag = scaleway.get_ipam_ips(tags=["tag"])
    ```

    ### By type and resource

    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    vpc01 = scaleway.Vpc("vpc01")
    pn01 = scaleway.VpcPrivateNetwork("pn01",
        vpc_id=vpc01.id,
        ipv4_subnet=scaleway.VpcPrivateNetworkIpv4SubnetArgs(
            subnet="172.16.32.0/22",
        ))
    redis01 = scaleway.RedisCluster("redis01",
        version="7.0.5",
        node_type="RED1-XS",
        user_name="my_initial_user",
        password="thiZ_is_v&ry_s3cret",
        cluster_size=3,
        private_networks=[scaleway.RedisClusterPrivateNetworkArgs(
            id=pn01.id,
        )])
    by_type_and_resource = scaleway.get_ipam_ips_output(type="ipv4",
        resource=scaleway.GetIpamIpsResourceArgs(
            id=redis01.id,
            type="redis_cluster",
        ))
    ```


    :param bool attached: Defines whether to filter only for IPs which are attached to a resource.
    :param str mac_address: The linked MAC address to filter for.
    :param str private_network_id: The ID of the Private Network to filter for.
    :param str project_id: The ID of the Project to filter for.
    :param str region: The region to filter for.
    :param pulumi.InputType['GetIpamIpsResourceArgs'] resource: Filter for a resource attached to the IP, using resource ID, type or name.
    :param Sequence[str] tags: The IP tags to filter for.
    :param str type: The type of IP to filter for (`ipv4` or `ipv6`).
    :param str zonal: Only IPs that are zonal, and in this zone, will be returned.
    """
    ...
