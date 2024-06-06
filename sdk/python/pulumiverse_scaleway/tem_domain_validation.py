# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TemDomainValidationArgs', 'TemDomainValidation']

@pulumi.input_type
class TemDomainValidationArgs:
    def __init__(__self__, *,
                 domain_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a TemDomainValidation resource.
        :param pulumi.Input[str] domain_id: The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
        :param pulumi.Input[str] region: `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
        :param pulumi.Input[int] timeout: The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
        """
        pulumi.set(__self__, "domain_id", domain_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Input[str]:
        """
        The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _TemDomainValidationState:
    def __init__(__self__, *,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 validated: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering TemDomainValidation resources.
        :param pulumi.Input[str] domain_id: The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
        :param pulumi.Input[str] region: `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
        :param pulumi.Input[int] timeout: The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
        :param pulumi.Input[bool] validated: Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
        """
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if validated is not None:
            pulumi.set(__self__, "validated", validated)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def validated(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
        """
        return pulumi.get(self, "validated")

    @validated.setter
    def validated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validated", value)


class TemDomainValidation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        example = scaleway.TemDomainValidation("example",
            domain_id="your-domain-id",
            region="fr-par",
            timeout=300)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_id: The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
        :param pulumi.Input[str] region: `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
        :param pulumi.Input[int] timeout: The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemDomainValidationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        example = scaleway.TemDomainValidation("example",
            domain_id="your-domain-id",
            region="fr-par",
            timeout=300)
        ```

        :param str resource_name: The name of the resource.
        :param TemDomainValidationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemDomainValidationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemDomainValidationArgs.__new__(TemDomainValidationArgs)

            if domain_id is None and not opts.urn:
                raise TypeError("Missing required property 'domain_id'")
            __props__.__dict__["domain_id"] = domain_id
            __props__.__dict__["region"] = region
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["validated"] = None
        super(TemDomainValidation, __self__).__init__(
            'scaleway:index/temDomainValidation:TemDomainValidation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            validated: Optional[pulumi.Input[bool]] = None) -> 'TemDomainValidation':
        """
        Get an existing TemDomainValidation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_id: The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
        :param pulumi.Input[str] region: `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
        :param pulumi.Input[int] timeout: The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
        :param pulumi.Input[bool] validated: Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TemDomainValidationState.__new__(_TemDomainValidationState)

        __props__.__dict__["domain_id"] = domain_id
        __props__.__dict__["region"] = region
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["validated"] = validated
        return TemDomainValidation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[str]:
        """
        The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def validated(self) -> pulumi.Output[bool]:
        """
        Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
        """
        return pulumi.get(self, "validated")

