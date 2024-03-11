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
    'GetVpcPublicGatewayDhcpReservationResult',
    'AwaitableGetVpcPublicGatewayDhcpReservationResult',
    'get_vpc_public_gateway_dhcp_reservation',
    'get_vpc_public_gateway_dhcp_reservation_output',
]

@pulumi.output_type
class GetVpcPublicGatewayDhcpReservationResult:
    """
    A collection of values returned by getVpcPublicGatewayDhcpReservation.
    """
    def __init__(__self__, created_at=None, gateway_network_id=None, hostname=None, id=None, ip_address=None, mac_address=None, reservation_id=None, type=None, updated_at=None, wait_for_dhcp=None, zone=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if gateway_network_id and not isinstance(gateway_network_id, str):
            raise TypeError("Expected argument 'gateway_network_id' to be a str")
        pulumi.set(__self__, "gateway_network_id", gateway_network_id)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if reservation_id and not isinstance(reservation_id, str):
            raise TypeError("Expected argument 'reservation_id' to be a str")
        pulumi.set(__self__, "reservation_id", reservation_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if wait_for_dhcp and not isinstance(wait_for_dhcp, bool):
            raise TypeError("Expected argument 'wait_for_dhcp' to be a bool")
        pulumi.set(__self__, "wait_for_dhcp", wait_for_dhcp)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time of the creation of the public gateway DHCP config.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="gatewayNetworkId")
    def gateway_network_id(self) -> Optional[str]:
        """
        The ID of the owning GatewayNetwork.
        """
        return pulumi.get(self, "gateway_network_id")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        The Hostname of the client machine.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The IP address to give to the machine (IP address).
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="reservationId")
    def reservation_id(self) -> Optional[str]:
        return pulumi.get(self, "reservation_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date and time of the last update of the public gateway DHCP config.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="waitForDhcp")
    def wait_for_dhcp(self) -> Optional[bool]:
        return pulumi.get(self, "wait_for_dhcp")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetVpcPublicGatewayDhcpReservationResult(GetVpcPublicGatewayDhcpReservationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcPublicGatewayDhcpReservationResult(
            created_at=self.created_at,
            gateway_network_id=self.gateway_network_id,
            hostname=self.hostname,
            id=self.id,
            ip_address=self.ip_address,
            mac_address=self.mac_address,
            reservation_id=self.reservation_id,
            type=self.type,
            updated_at=self.updated_at,
            wait_for_dhcp=self.wait_for_dhcp,
            zone=self.zone)


def get_vpc_public_gateway_dhcp_reservation(gateway_network_id: Optional[str] = None,
                                            mac_address: Optional[str] = None,
                                            reservation_id: Optional[str] = None,
                                            wait_for_dhcp: Optional[bool] = None,
                                            zone: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcPublicGatewayDhcpReservationResult:
    """
    Gets information about a dhcp entries. For further information please check the
    API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6)

    ## Example Dynamic

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    main_vpc_private_network = scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork")
    main_instance_server = scaleway.InstanceServer("mainInstanceServer",
        image="ubuntu_jammy",
        type="DEV1-S",
        zone="fr-par-1",
        private_networks=[scaleway.InstanceServerPrivateNetworkArgs(
            pn_id=main_vpc_private_network.id,
        )])
    main_vpc_public_gateway_ip = scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp")
    main_vpc_public_gateway_dhcp = scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", subnet="192.168.1.0/24")
    main_vpc_public_gateway = scaleway.VpcPublicGateway("mainVpcPublicGateway",
        type="VPC-GW-S",
        ip_id=main_vpc_public_gateway_ip.id)
    main_vpc_gateway_network = scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork",
        gateway_id=main_vpc_public_gateway.id,
        private_network_id=main_vpc_private_network.id,
        dhcp_id=main_vpc_public_gateway_dhcp.id,
        cleanup_dhcp=True,
        enable_masquerade=True)
    ## Retrieve the dynamic entries generated by mac address & gateway network
    by_mac_address_and_gw_network = pulumi.Output.all(main_instance_server.private_networks, main_vpc_gateway_network.id).apply(lambda private_networks, id: scaleway.get_vpc_public_gateway_dhcp_reservation_output(mac_address=private_networks[0].mac_address,
        gateway_network_id=id))
    ```
    <!--End PulumiCodeChooser -->

    ## Example Static and PAT rule

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    main_vpc_private_network = scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork")
    main_instance_security_group = scaleway.InstanceSecurityGroup("mainInstanceSecurityGroup",
        inbound_default_policy="drop",
        outbound_default_policy="accept",
        inbound_rules=[scaleway.InstanceSecurityGroupInboundRuleArgs(
            action="accept",
            port=22,
        )])
    main_instance_server = scaleway.InstanceServer("mainInstanceServer",
        image="ubuntu_jammy",
        type="DEV1-S",
        zone="fr-par-1",
        security_group_id=main_instance_security_group.id,
        private_networks=[scaleway.InstanceServerPrivateNetworkArgs(
            pn_id=main_vpc_private_network.id,
        )])
    main_vpc_public_gateway_ip = scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp")
    main_vpc_public_gateway_dhcp = scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", subnet="192.168.1.0/24")
    main_vpc_public_gateway = scaleway.VpcPublicGateway("mainVpcPublicGateway",
        type="VPC-GW-S",
        ip_id=main_vpc_public_gateway_ip.id)
    main_vpc_gateway_network = scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork",
        gateway_id=main_vpc_public_gateway.id,
        private_network_id=main_vpc_private_network.id,
        dhcp_id=main_vpc_public_gateway_dhcp.id,
        cleanup_dhcp=True,
        enable_masquerade=True)
    main_vpc_public_gateway_dhcp_reservation = scaleway.VpcPublicGatewayDhcpReservation("mainVpcPublicGatewayDhcpReservation",
        gateway_network_id=main_vpc_gateway_network.id,
        mac_address=main_instance_server.private_networks[0].mac_address,
        ip_address="192.168.1.4")
    ### VPC PAT RULE
    main_vpc_public_gateway_pat_rule = scaleway.VpcPublicGatewayPatRule("mainVpcPublicGatewayPatRule",
        gateway_id=main_vpc_public_gateway.id,
        private_ip=main_vpc_public_gateway_dhcp_reservation.ip_address,
        private_port=22,
        public_port=2222,
        protocol="tcp")
    by_id = scaleway.get_vpc_public_gateway_dhcp_reservation_output(reservation_id=main_vpc_public_gateway_dhcp_reservation.id)
    ```
    <!--End PulumiCodeChooser -->


    :param str gateway_network_id: The ID of the owning GatewayNetwork
    :param str mac_address: The MAC address of the reservation to retrieve
    :param str reservation_id: The ID of the Reservation to retrieve
    :param bool wait_for_dhcp: Boolean to wait for mac_address to exist in dhcp
    :param str zone: `zone`) The zone in which
           the image exists.
    """
    __args__ = dict()
    __args__['gatewayNetworkId'] = gateway_network_id
    __args__['macAddress'] = mac_address
    __args__['reservationId'] = reservation_id
    __args__['waitForDhcp'] = wait_for_dhcp
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcPublicGatewayDhcpReservation:getVpcPublicGatewayDhcpReservation', __args__, opts=opts, typ=GetVpcPublicGatewayDhcpReservationResult).value

    return AwaitableGetVpcPublicGatewayDhcpReservationResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        gateway_network_id=pulumi.get(__ret__, 'gateway_network_id'),
        hostname=pulumi.get(__ret__, 'hostname'),
        id=pulumi.get(__ret__, 'id'),
        ip_address=pulumi.get(__ret__, 'ip_address'),
        mac_address=pulumi.get(__ret__, 'mac_address'),
        reservation_id=pulumi.get(__ret__, 'reservation_id'),
        type=pulumi.get(__ret__, 'type'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        wait_for_dhcp=pulumi.get(__ret__, 'wait_for_dhcp'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_vpc_public_gateway_dhcp_reservation)
def get_vpc_public_gateway_dhcp_reservation_output(gateway_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                   mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                                                   reservation_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                   wait_for_dhcp: Optional[pulumi.Input[Optional[bool]]] = None,
                                                   zone: Optional[pulumi.Input[Optional[str]]] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcPublicGatewayDhcpReservationResult]:
    """
    Gets information about a dhcp entries. For further information please check the
    API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6)

    ## Example Dynamic

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    main_vpc_private_network = scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork")
    main_instance_server = scaleway.InstanceServer("mainInstanceServer",
        image="ubuntu_jammy",
        type="DEV1-S",
        zone="fr-par-1",
        private_networks=[scaleway.InstanceServerPrivateNetworkArgs(
            pn_id=main_vpc_private_network.id,
        )])
    main_vpc_public_gateway_ip = scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp")
    main_vpc_public_gateway_dhcp = scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", subnet="192.168.1.0/24")
    main_vpc_public_gateway = scaleway.VpcPublicGateway("mainVpcPublicGateway",
        type="VPC-GW-S",
        ip_id=main_vpc_public_gateway_ip.id)
    main_vpc_gateway_network = scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork",
        gateway_id=main_vpc_public_gateway.id,
        private_network_id=main_vpc_private_network.id,
        dhcp_id=main_vpc_public_gateway_dhcp.id,
        cleanup_dhcp=True,
        enable_masquerade=True)
    ## Retrieve the dynamic entries generated by mac address & gateway network
    by_mac_address_and_gw_network = pulumi.Output.all(main_instance_server.private_networks, main_vpc_gateway_network.id).apply(lambda private_networks, id: scaleway.get_vpc_public_gateway_dhcp_reservation_output(mac_address=private_networks[0].mac_address,
        gateway_network_id=id))
    ```
    <!--End PulumiCodeChooser -->

    ## Example Static and PAT rule

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway
    import pulumiverse_scaleway as scaleway

    main_vpc_private_network = scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork")
    main_instance_security_group = scaleway.InstanceSecurityGroup("mainInstanceSecurityGroup",
        inbound_default_policy="drop",
        outbound_default_policy="accept",
        inbound_rules=[scaleway.InstanceSecurityGroupInboundRuleArgs(
            action="accept",
            port=22,
        )])
    main_instance_server = scaleway.InstanceServer("mainInstanceServer",
        image="ubuntu_jammy",
        type="DEV1-S",
        zone="fr-par-1",
        security_group_id=main_instance_security_group.id,
        private_networks=[scaleway.InstanceServerPrivateNetworkArgs(
            pn_id=main_vpc_private_network.id,
        )])
    main_vpc_public_gateway_ip = scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp")
    main_vpc_public_gateway_dhcp = scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", subnet="192.168.1.0/24")
    main_vpc_public_gateway = scaleway.VpcPublicGateway("mainVpcPublicGateway",
        type="VPC-GW-S",
        ip_id=main_vpc_public_gateway_ip.id)
    main_vpc_gateway_network = scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork",
        gateway_id=main_vpc_public_gateway.id,
        private_network_id=main_vpc_private_network.id,
        dhcp_id=main_vpc_public_gateway_dhcp.id,
        cleanup_dhcp=True,
        enable_masquerade=True)
    main_vpc_public_gateway_dhcp_reservation = scaleway.VpcPublicGatewayDhcpReservation("mainVpcPublicGatewayDhcpReservation",
        gateway_network_id=main_vpc_gateway_network.id,
        mac_address=main_instance_server.private_networks[0].mac_address,
        ip_address="192.168.1.4")
    ### VPC PAT RULE
    main_vpc_public_gateway_pat_rule = scaleway.VpcPublicGatewayPatRule("mainVpcPublicGatewayPatRule",
        gateway_id=main_vpc_public_gateway.id,
        private_ip=main_vpc_public_gateway_dhcp_reservation.ip_address,
        private_port=22,
        public_port=2222,
        protocol="tcp")
    by_id = scaleway.get_vpc_public_gateway_dhcp_reservation_output(reservation_id=main_vpc_public_gateway_dhcp_reservation.id)
    ```
    <!--End PulumiCodeChooser -->


    :param str gateway_network_id: The ID of the owning GatewayNetwork
    :param str mac_address: The MAC address of the reservation to retrieve
    :param str reservation_id: The ID of the Reservation to retrieve
    :param bool wait_for_dhcp: Boolean to wait for mac_address to exist in dhcp
    :param str zone: `zone`) The zone in which
           the image exists.
    """
    ...
