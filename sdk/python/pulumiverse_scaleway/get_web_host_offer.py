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
    'GetWebHostOfferResult',
    'AwaitableGetWebHostOfferResult',
    'get_web_host_offer',
    'get_web_host_offer_output',
]

@pulumi.output_type
class GetWebHostOfferResult:
    """
    A collection of values returned by getWebHostOffer.
    """
    def __init__(__self__, billing_operation_path=None, id=None, name=None, offer_id=None, price=None, products=None, region=None):
        if billing_operation_path and not isinstance(billing_operation_path, str):
            raise TypeError("Expected argument 'billing_operation_path' to be a str")
        pulumi.set(__self__, "billing_operation_path", billing_operation_path)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if offer_id and not isinstance(offer_id, str):
            raise TypeError("Expected argument 'offer_id' to be a str")
        pulumi.set(__self__, "offer_id", offer_id)
        if price and not isinstance(price, str):
            raise TypeError("Expected argument 'price' to be a str")
        pulumi.set(__self__, "price", price)
        if products and not isinstance(products, list):
            raise TypeError("Expected argument 'products' to be a list")
        pulumi.set(__self__, "products", products)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="billingOperationPath")
    def billing_operation_path(self) -> str:
        """
        The unique identifier used for billing.
        """
        return pulumi.get(self, "billing_operation_path")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

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
    def price(self) -> str:
        """
        The offer price.
        """
        return pulumi.get(self, "price")

    @property
    @pulumi.getter
    def products(self) -> Sequence['outputs.GetWebHostOfferProductResult']:
        """
        The offer product.
        """
        return pulumi.get(self, "products")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetWebHostOfferResult(GetWebHostOfferResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebHostOfferResult(
            billing_operation_path=self.billing_operation_path,
            id=self.id,
            name=self.name,
            offer_id=self.offer_id,
            price=self.price,
            products=self.products,
            region=self.region)


def get_web_host_offer(name: Optional[str] = None,
                       offer_id: Optional[str] = None,
                       region: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebHostOfferResult:
    """
    Gets information about a webhosting offer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.get_web_host_offer(name="performance")
    by_id = scaleway.get_web_host_offer(offer_id="de2426b4-a9e9-11ec-b909-0242ac120002")
    ```


    :param str name: The offer name. Only one of `name` and `offer_id` should be specified.
    :param str offer_id: The offer id. Only one of `name` and `offer_id` should be specified.
    :param str region: `region`) The region in which offer exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['offerId'] = offer_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getWebHostOffer:getWebHostOffer', __args__, opts=opts, typ=GetWebHostOfferResult).value

    return AwaitableGetWebHostOfferResult(
        billing_operation_path=pulumi.get(__ret__, 'billing_operation_path'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        offer_id=pulumi.get(__ret__, 'offer_id'),
        price=pulumi.get(__ret__, 'price'),
        products=pulumi.get(__ret__, 'products'),
        region=pulumi.get(__ret__, 'region'))


@_utilities.lift_output_func(get_web_host_offer)
def get_web_host_offer_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                              offer_id: Optional[pulumi.Input[Optional[str]]] = None,
                              region: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWebHostOfferResult]:
    """
    Gets information about a webhosting offer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.get_web_host_offer(name="performance")
    by_id = scaleway.get_web_host_offer(offer_id="de2426b4-a9e9-11ec-b909-0242ac120002")
    ```


    :param str name: The offer name. Only one of `name` and `offer_id` should be specified.
    :param str offer_id: The offer id. Only one of `name` and `offer_id` should be specified.
    :param str region: `region`) The region in which offer exists.
    """
    ...
