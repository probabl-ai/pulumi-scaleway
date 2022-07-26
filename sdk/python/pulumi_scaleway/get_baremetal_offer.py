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

__all__ = [
    'GetBaremetalOfferResult',
    'AwaitableGetBaremetalOfferResult',
    'get_baremetal_offer',
    'get_baremetal_offer_output',
]

@pulumi.output_type
class GetBaremetalOfferResult:
    """
    A collection of values returned by getBaremetalOffer.
    """
    def __init__(__self__, bandwidth=None, commercial_range=None, cpu=None, disks=None, id=None, include_disabled=None, memories=None, name=None, offer_id=None, stock=None, zone=None):
        if bandwidth and not isinstance(bandwidth, int):
            raise TypeError("Expected argument 'bandwidth' to be a int")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if commercial_range and not isinstance(commercial_range, str):
            raise TypeError("Expected argument 'commercial_range' to be a str")
        pulumi.set(__self__, "commercial_range", commercial_range)
        if cpu and not isinstance(cpu, dict):
            raise TypeError("Expected argument 'cpu' to be a dict")
        pulumi.set(__self__, "cpu", cpu)
        if disks and not isinstance(disks, list):
            raise TypeError("Expected argument 'disks' to be a list")
        pulumi.set(__self__, "disks", disks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_disabled and not isinstance(include_disabled, bool):
            raise TypeError("Expected argument 'include_disabled' to be a bool")
        pulumi.set(__self__, "include_disabled", include_disabled)
        if memories and not isinstance(memories, list):
            raise TypeError("Expected argument 'memories' to be a list")
        pulumi.set(__self__, "memories", memories)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if offer_id and not isinstance(offer_id, str):
            raise TypeError("Expected argument 'offer_id' to be a str")
        pulumi.set(__self__, "offer_id", offer_id)
        if stock and not isinstance(stock, str):
            raise TypeError("Expected argument 'stock' to be a str")
        pulumi.set(__self__, "stock", stock)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="commercialRange")
    def commercial_range(self) -> str:
        return pulumi.get(self, "commercial_range")

    @property
    @pulumi.getter
    def cpu(self) -> 'outputs.GetBaremetalOfferCpuResult':
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter
    def disks(self) -> Sequence['outputs.GetBaremetalOfferDiskResult']:
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeDisabled")
    def include_disabled(self) -> Optional[bool]:
        return pulumi.get(self, "include_disabled")

    @property
    @pulumi.getter
    def memories(self) -> Sequence['outputs.GetBaremetalOfferMemoryResult']:
        return pulumi.get(self, "memories")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="offerId")
    def offer_id(self) -> Optional[str]:
        return pulumi.get(self, "offer_id")

    @property
    @pulumi.getter
    def stock(self) -> str:
        return pulumi.get(self, "stock")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetBaremetalOfferResult(GetBaremetalOfferResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBaremetalOfferResult(
            bandwidth=self.bandwidth,
            commercial_range=self.commercial_range,
            cpu=self.cpu,
            disks=self.disks,
            id=self.id,
            include_disabled=self.include_disabled,
            memories=self.memories,
            name=self.name,
            offer_id=self.offer_id,
            stock=self.stock,
            zone=self.zone)


def get_baremetal_offer(include_disabled: Optional[bool] = None,
                        name: Optional[str] = None,
                        offer_id: Optional[str] = None,
                        zone: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBaremetalOfferResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['includeDisabled'] = include_disabled
    __args__['name'] = name
    __args__['offerId'] = offer_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getBaremetalOffer:getBaremetalOffer', __args__, opts=opts, typ=GetBaremetalOfferResult).value

    return AwaitableGetBaremetalOfferResult(
        bandwidth=__ret__.bandwidth,
        commercial_range=__ret__.commercial_range,
        cpu=__ret__.cpu,
        disks=__ret__.disks,
        id=__ret__.id,
        include_disabled=__ret__.include_disabled,
        memories=__ret__.memories,
        name=__ret__.name,
        offer_id=__ret__.offer_id,
        stock=__ret__.stock,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_baremetal_offer)
def get_baremetal_offer_output(include_disabled: Optional[pulumi.Input[Optional[bool]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               offer_id: Optional[pulumi.Input[Optional[str]]] = None,
                               zone: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBaremetalOfferResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
