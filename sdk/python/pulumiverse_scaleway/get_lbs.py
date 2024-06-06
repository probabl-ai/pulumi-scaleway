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
    'GetLbsResult',
    'AwaitableGetLbsResult',
    'get_lbs',
    'get_lbs_output',
]

@pulumi.output_type
class GetLbsResult:
    """
    A collection of values returned by getLbs.
    """
    def __init__(__self__, id=None, lbs=None, name=None, organization_id=None, project_id=None, tags=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lbs and not isinstance(lbs, list):
            raise TypeError("Expected argument 'lbs' to be a list")
        pulumi.set(__self__, "lbs", lbs)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
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
    @pulumi.getter
    def lbs(self) -> Sequence['outputs.GetLbsLbResult']:
        """
        List of found LBs
        """
        return pulumi.get(self, "lbs")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the load-balancer.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The organization ID the load-balancer is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID of the project the load-balancer is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The tags associated with the load-balancer.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The zone in which the load-balancer is.
        """
        return pulumi.get(self, "zone")


class AwaitableGetLbsResult(GetLbsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbsResult(
            id=self.id,
            lbs=self.lbs,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            tags=self.tags,
            zone=self.zone)


def get_lbs(name: Optional[str] = None,
            project_id: Optional[str] = None,
            tags: Optional[Sequence[str]] = None,
            zone: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbsResult:
    """
    Gets information about multiple Load Balancers.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_lbs(name="foobar",
        zone="fr-par-2")
    lbs_by_tags = scaleway.get_lbs(tags=["a tag"])
    ```


    :param str name: The load balancer name used as a filter. LBs with a name like it are listed.
    :param str project_id: The ID of the project the load-balancer is associated with.
    :param Sequence[str] tags: List of tags used as filter. LBs with these exact tags are listed.
    :param str zone: `zone`) The zone in which LBs exist.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['tags'] = tags
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLbs:getLbs', __args__, opts=opts, typ=GetLbsResult).value

    return AwaitableGetLbsResult(
        id=pulumi.get(__ret__, 'id'),
        lbs=pulumi.get(__ret__, 'lbs'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        tags=pulumi.get(__ret__, 'tags'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_lbs)
def get_lbs_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                   project_id: Optional[pulumi.Input[Optional[str]]] = None,
                   tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                   zone: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLbsResult]:
    """
    Gets information about multiple Load Balancers.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_lbs(name="foobar",
        zone="fr-par-2")
    lbs_by_tags = scaleway.get_lbs(tags=["a tag"])
    ```


    :param str name: The load balancer name used as a filter. LBs with a name like it are listed.
    :param str project_id: The ID of the project the load-balancer is associated with.
    :param Sequence[str] tags: List of tags used as filter. LBs with these exact tags are listed.
    :param str zone: `zone`) The zone in which LBs exist.
    """
    ...
