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
    'GetTemDomainResult',
    'AwaitableGetTemDomainResult',
    'get_tem_domain',
    'get_tem_domain_output',
]

@pulumi.output_type
class GetTemDomainResult:
    """
    A collection of values returned by getTemDomain.
    """
    def __init__(__self__, accept_tos=None, created_at=None, dkim_config=None, domain_id=None, id=None, last_error=None, last_valid_at=None, mx_blackhole=None, name=None, next_check_at=None, project_id=None, region=None, reputations=None, revoked_at=None, smtp_host=None, smtp_port=None, smtp_port_alternative=None, smtp_port_unsecure=None, smtps_auth_user=None, smtps_port=None, smtps_port_alternative=None, spf_config=None, status=None):
        if accept_tos and not isinstance(accept_tos, bool):
            raise TypeError("Expected argument 'accept_tos' to be a bool")
        pulumi.set(__self__, "accept_tos", accept_tos)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if dkim_config and not isinstance(dkim_config, str):
            raise TypeError("Expected argument 'dkim_config' to be a str")
        pulumi.set(__self__, "dkim_config", dkim_config)
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        pulumi.set(__self__, "domain_id", domain_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_error and not isinstance(last_error, str):
            raise TypeError("Expected argument 'last_error' to be a str")
        pulumi.set(__self__, "last_error", last_error)
        if last_valid_at and not isinstance(last_valid_at, str):
            raise TypeError("Expected argument 'last_valid_at' to be a str")
        pulumi.set(__self__, "last_valid_at", last_valid_at)
        if mx_blackhole and not isinstance(mx_blackhole, str):
            raise TypeError("Expected argument 'mx_blackhole' to be a str")
        pulumi.set(__self__, "mx_blackhole", mx_blackhole)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if next_check_at and not isinstance(next_check_at, str):
            raise TypeError("Expected argument 'next_check_at' to be a str")
        pulumi.set(__self__, "next_check_at", next_check_at)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if reputations and not isinstance(reputations, list):
            raise TypeError("Expected argument 'reputations' to be a list")
        pulumi.set(__self__, "reputations", reputations)
        if revoked_at and not isinstance(revoked_at, str):
            raise TypeError("Expected argument 'revoked_at' to be a str")
        pulumi.set(__self__, "revoked_at", revoked_at)
        if smtp_host and not isinstance(smtp_host, str):
            raise TypeError("Expected argument 'smtp_host' to be a str")
        pulumi.set(__self__, "smtp_host", smtp_host)
        if smtp_port and not isinstance(smtp_port, int):
            raise TypeError("Expected argument 'smtp_port' to be a int")
        pulumi.set(__self__, "smtp_port", smtp_port)
        if smtp_port_alternative and not isinstance(smtp_port_alternative, int):
            raise TypeError("Expected argument 'smtp_port_alternative' to be a int")
        pulumi.set(__self__, "smtp_port_alternative", smtp_port_alternative)
        if smtp_port_unsecure and not isinstance(smtp_port_unsecure, int):
            raise TypeError("Expected argument 'smtp_port_unsecure' to be a int")
        pulumi.set(__self__, "smtp_port_unsecure", smtp_port_unsecure)
        if smtps_auth_user and not isinstance(smtps_auth_user, str):
            raise TypeError("Expected argument 'smtps_auth_user' to be a str")
        pulumi.set(__self__, "smtps_auth_user", smtps_auth_user)
        if smtps_port and not isinstance(smtps_port, int):
            raise TypeError("Expected argument 'smtps_port' to be a int")
        pulumi.set(__self__, "smtps_port", smtps_port)
        if smtps_port_alternative and not isinstance(smtps_port_alternative, int):
            raise TypeError("Expected argument 'smtps_port_alternative' to be a int")
        pulumi.set(__self__, "smtps_port_alternative", smtps_port_alternative)
        if spf_config and not isinstance(spf_config, str):
            raise TypeError("Expected argument 'spf_config' to be a str")
        pulumi.set(__self__, "spf_config", spf_config)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="acceptTos")
    def accept_tos(self) -> bool:
        return pulumi.get(self, "accept_tos")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dkimConfig")
    def dkim_config(self) -> str:
        return pulumi.get(self, "dkim_config")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[str]:
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastError")
    def last_error(self) -> str:
        return pulumi.get(self, "last_error")

    @property
    @pulumi.getter(name="lastValidAt")
    def last_valid_at(self) -> str:
        return pulumi.get(self, "last_valid_at")

    @property
    @pulumi.getter(name="mxBlackhole")
    def mx_blackhole(self) -> str:
        return pulumi.get(self, "mx_blackhole")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextCheckAt")
    def next_check_at(self) -> str:
        return pulumi.get(self, "next_check_at")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def reputations(self) -> Sequence['outputs.GetTemDomainReputationResult']:
        return pulumi.get(self, "reputations")

    @property
    @pulumi.getter(name="revokedAt")
    def revoked_at(self) -> str:
        return pulumi.get(self, "revoked_at")

    @property
    @pulumi.getter(name="smtpHost")
    def smtp_host(self) -> str:
        return pulumi.get(self, "smtp_host")

    @property
    @pulumi.getter(name="smtpPort")
    def smtp_port(self) -> int:
        return pulumi.get(self, "smtp_port")

    @property
    @pulumi.getter(name="smtpPortAlternative")
    def smtp_port_alternative(self) -> int:
        return pulumi.get(self, "smtp_port_alternative")

    @property
    @pulumi.getter(name="smtpPortUnsecure")
    def smtp_port_unsecure(self) -> int:
        return pulumi.get(self, "smtp_port_unsecure")

    @property
    @pulumi.getter(name="smtpsAuthUser")
    def smtps_auth_user(self) -> str:
        return pulumi.get(self, "smtps_auth_user")

    @property
    @pulumi.getter(name="smtpsPort")
    def smtps_port(self) -> int:
        return pulumi.get(self, "smtps_port")

    @property
    @pulumi.getter(name="smtpsPortAlternative")
    def smtps_port_alternative(self) -> int:
        return pulumi.get(self, "smtps_port_alternative")

    @property
    @pulumi.getter(name="spfConfig")
    def spf_config(self) -> str:
        return pulumi.get(self, "spf_config")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


class AwaitableGetTemDomainResult(GetTemDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTemDomainResult(
            accept_tos=self.accept_tos,
            created_at=self.created_at,
            dkim_config=self.dkim_config,
            domain_id=self.domain_id,
            id=self.id,
            last_error=self.last_error,
            last_valid_at=self.last_valid_at,
            mx_blackhole=self.mx_blackhole,
            name=self.name,
            next_check_at=self.next_check_at,
            project_id=self.project_id,
            region=self.region,
            reputations=self.reputations,
            revoked_at=self.revoked_at,
            smtp_host=self.smtp_host,
            smtp_port=self.smtp_port,
            smtp_port_alternative=self.smtp_port_alternative,
            smtp_port_unsecure=self.smtp_port_unsecure,
            smtps_auth_user=self.smtps_auth_user,
            smtps_port=self.smtps_port,
            smtps_port_alternative=self.smtps_port_alternative,
            spf_config=self.spf_config,
            status=self.status)


def get_tem_domain(domain_id: Optional[str] = None,
                   name: Optional[str] = None,
                   project_id: Optional[str] = None,
                   region: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTemDomainResult:
    """
    Gets information about a transactional email domain.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_domain = scaleway.get_tem_domain(domain_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str domain_id: The domain id.
           Only one of `name` and `domain_id` should be specified.
    :param str name: The domain name.
           Only one of `name` and `domain_id` should be specified.
    :param str project_id: `project_id`) The ID of the project the domain is associated with.
    :param str region: `region`) The region in which the domain exists.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getTemDomain:getTemDomain', __args__, opts=opts, typ=GetTemDomainResult).value

    return AwaitableGetTemDomainResult(
        accept_tos=pulumi.get(__ret__, 'accept_tos'),
        created_at=pulumi.get(__ret__, 'created_at'),
        dkim_config=pulumi.get(__ret__, 'dkim_config'),
        domain_id=pulumi.get(__ret__, 'domain_id'),
        id=pulumi.get(__ret__, 'id'),
        last_error=pulumi.get(__ret__, 'last_error'),
        last_valid_at=pulumi.get(__ret__, 'last_valid_at'),
        mx_blackhole=pulumi.get(__ret__, 'mx_blackhole'),
        name=pulumi.get(__ret__, 'name'),
        next_check_at=pulumi.get(__ret__, 'next_check_at'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        reputations=pulumi.get(__ret__, 'reputations'),
        revoked_at=pulumi.get(__ret__, 'revoked_at'),
        smtp_host=pulumi.get(__ret__, 'smtp_host'),
        smtp_port=pulumi.get(__ret__, 'smtp_port'),
        smtp_port_alternative=pulumi.get(__ret__, 'smtp_port_alternative'),
        smtp_port_unsecure=pulumi.get(__ret__, 'smtp_port_unsecure'),
        smtps_auth_user=pulumi.get(__ret__, 'smtps_auth_user'),
        smtps_port=pulumi.get(__ret__, 'smtps_port'),
        smtps_port_alternative=pulumi.get(__ret__, 'smtps_port_alternative'),
        spf_config=pulumi.get(__ret__, 'spf_config'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_tem_domain)
def get_tem_domain_output(domain_id: Optional[pulumi.Input[Optional[str]]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          project_id: Optional[pulumi.Input[Optional[str]]] = None,
                          region: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTemDomainResult]:
    """
    Gets information about a transactional email domain.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_domain = scaleway.get_tem_domain(domain_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str domain_id: The domain id.
           Only one of `name` and `domain_id` should be specified.
    :param str name: The domain name.
           Only one of `name` and `domain_id` should be specified.
    :param str project_id: `project_id`) The ID of the project the domain is associated with.
    :param str region: `region`) The region in which the domain exists.
    """
    ...
