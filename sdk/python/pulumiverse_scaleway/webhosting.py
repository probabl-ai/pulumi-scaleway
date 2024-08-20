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

__all__ = ['WebhostingArgs', 'Webhosting']

@pulumi.input_type
class WebhostingArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 email: pulumi.Input[str],
                 offer_id: pulumi.Input[str],
                 option_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Webhosting resource.
        :param pulumi.Input[str] domain: The domain name of the hosting.
        :param pulumi.Input[str] email: The contact email of the client for the hosting.
        :param pulumi.Input[str] offer_id: The ID of the selected offer for the hosting.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] option_ids: The IDs of the selected options for the hosting.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the VPC is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Hosting.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the hosting.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "offer_id", offer_id)
        if option_ids is not None:
            pulumi.set(__self__, "option_ids", option_ids)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain name of the hosting.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        The contact email of the client for the hosting.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="offerId")
    def offer_id(self) -> pulumi.Input[str]:
        """
        The ID of the selected offer for the hosting.
        """
        return pulumi.get(self, "offer_id")

    @offer_id.setter
    def offer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "offer_id", value)

    @property
    @pulumi.getter(name="optionIds")
    def option_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of the selected options for the hosting.
        """
        return pulumi.get(self, "option_ids")

    @option_ids.setter
    def option_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "option_ids", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the VPC is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the Hosting.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the hosting.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _WebhostingState:
    def __init__(__self__, *,
                 cpanel_urls: Optional[pulumi.Input[Sequence[pulumi.Input['WebhostingCpanelUrlArgs']]]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 dns_status: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 offer_id: Optional[pulumi.Input[str]] = None,
                 offer_name: Optional[pulumi.Input[str]] = None,
                 option_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input['WebhostingOptionArgs']]]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 platform_hostname: Optional[pulumi.Input[str]] = None,
                 platform_number: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Webhosting resources.
        :param pulumi.Input[Sequence[pulumi.Input['WebhostingCpanelUrlArgs']]] cpanel_urls: The URL to connect to cPanel Dashboard and to Webmail interface.
        :param pulumi.Input[str] created_at: Date and time of hosting's creation (RFC 3339 format).
        :param pulumi.Input[str] dns_status: The DNS status of the hosting.
        :param pulumi.Input[str] domain: The domain name of the hosting.
        :param pulumi.Input[str] email: The contact email of the client for the hosting.
        :param pulumi.Input[str] offer_id: The ID of the selected offer for the hosting.
        :param pulumi.Input[str] offer_name: The name of the active offer.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] option_ids: The IDs of the selected options for the hosting.
        :param pulumi.Input[Sequence[pulumi.Input['WebhostingOptionArgs']]] options: The active options of the hosting.
        :param pulumi.Input[str] organization_id: The organization ID the hosting is associated with.
        :param pulumi.Input[str] platform_hostname: The hostname of the host platform.
        :param pulumi.Input[int] platform_number: The number of the host platform.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the VPC is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Hosting.
        :param pulumi.Input[str] status: The hosting status.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the hosting.
        :param pulumi.Input[str] updated_at: Date and time of hosting's last update (RFC 3339 format).
        :param pulumi.Input[str] username: The main hosting cPanel username.
        """
        if cpanel_urls is not None:
            pulumi.set(__self__, "cpanel_urls", cpanel_urls)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if dns_status is not None:
            pulumi.set(__self__, "dns_status", dns_status)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if offer_id is not None:
            pulumi.set(__self__, "offer_id", offer_id)
        if offer_name is not None:
            pulumi.set(__self__, "offer_name", offer_name)
        if option_ids is not None:
            pulumi.set(__self__, "option_ids", option_ids)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if platform_hostname is not None:
            pulumi.set(__self__, "platform_hostname", platform_hostname)
        if platform_number is not None:
            pulumi.set(__self__, "platform_number", platform_number)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="cpanelUrls")
    def cpanel_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebhostingCpanelUrlArgs']]]]:
        """
        The URL to connect to cPanel Dashboard and to Webmail interface.
        """
        return pulumi.get(self, "cpanel_urls")

    @cpanel_urls.setter
    def cpanel_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebhostingCpanelUrlArgs']]]]):
        pulumi.set(self, "cpanel_urls", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of hosting's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dnsStatus")
    def dns_status(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS status of the hosting.
        """
        return pulumi.get(self, "dns_status")

    @dns_status.setter
    def dns_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_status", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name of the hosting.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The contact email of the client for the hosting.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="offerId")
    def offer_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the selected offer for the hosting.
        """
        return pulumi.get(self, "offer_id")

    @offer_id.setter
    def offer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "offer_id", value)

    @property
    @pulumi.getter(name="offerName")
    def offer_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the active offer.
        """
        return pulumi.get(self, "offer_name")

    @offer_name.setter
    def offer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "offer_name", value)

    @property
    @pulumi.getter(name="optionIds")
    def option_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of the selected options for the hosting.
        """
        return pulumi.get(self, "option_ids")

    @option_ids.setter
    def option_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "option_ids", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebhostingOptionArgs']]]]:
        """
        The active options of the hosting.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebhostingOptionArgs']]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the hosting is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="platformHostname")
    def platform_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The hostname of the host platform.
        """
        return pulumi.get(self, "platform_hostname")

    @platform_hostname.setter
    def platform_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_hostname", value)

    @property
    @pulumi.getter(name="platformNumber")
    def platform_number(self) -> Optional[pulumi.Input[int]]:
        """
        The number of the host platform.
        """
        return pulumi.get(self, "platform_number")

    @platform_number.setter
    def platform_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "platform_number", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the VPC is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the Hosting.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The hosting status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the hosting.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of hosting's last update (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The main hosting cPanel username.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Webhosting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 offer_id: Optional[pulumi.Input[str]] = None,
                 option_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Web Hostings.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/webhosting/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        by_name = scaleway.get_web_host_offer(name="lite")
        main = scaleway.Webhosting("main",
            offer_id=by_name.offer_id,
            email="your@email.com",
            domain="yourdomain.com",
            tags=[
                "webhosting",
                "provider",
                "terraform",
            ])
        ```

        ## Import

        Hostings can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/webhosting:Webhosting hosting01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The domain name of the hosting.
        :param pulumi.Input[str] email: The contact email of the client for the hosting.
        :param pulumi.Input[str] offer_id: The ID of the selected offer for the hosting.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] option_ids: The IDs of the selected options for the hosting.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the VPC is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Hosting.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the hosting.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebhostingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Web Hostings.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/webhosting/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_scaleway as scaleway
        import pulumiverse_scaleway as scaleway

        by_name = scaleway.get_web_host_offer(name="lite")
        main = scaleway.Webhosting("main",
            offer_id=by_name.offer_id,
            email="your@email.com",
            domain="yourdomain.com",
            tags=[
                "webhosting",
                "provider",
                "terraform",
            ])
        ```

        ## Import

        Hostings can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/webhosting:Webhosting hosting01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param WebhostingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebhostingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 offer_id: Optional[pulumi.Input[str]] = None,
                 option_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebhostingArgs.__new__(WebhostingArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            if offer_id is None and not opts.urn:
                raise TypeError("Missing required property 'offer_id'")
            __props__.__dict__["offer_id"] = offer_id
            __props__.__dict__["option_ids"] = option_ids
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
            __props__.__dict__["cpanel_urls"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["dns_status"] = None
            __props__.__dict__["offer_name"] = None
            __props__.__dict__["options"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["platform_hostname"] = None
            __props__.__dict__["platform_number"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["username"] = None
        super(Webhosting, __self__).__init__(
            'scaleway:index/webhosting:Webhosting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cpanel_urls: Optional[pulumi.Input[Sequence[pulumi.Input[Union['WebhostingCpanelUrlArgs', 'WebhostingCpanelUrlArgsDict']]]]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            dns_status: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            offer_id: Optional[pulumi.Input[str]] = None,
            offer_name: Optional[pulumi.Input[str]] = None,
            option_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['WebhostingOptionArgs', 'WebhostingOptionArgsDict']]]]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            platform_hostname: Optional[pulumi.Input[str]] = None,
            platform_number: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'Webhosting':
        """
        Get an existing Webhosting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['WebhostingCpanelUrlArgs', 'WebhostingCpanelUrlArgsDict']]]] cpanel_urls: The URL to connect to cPanel Dashboard and to Webmail interface.
        :param pulumi.Input[str] created_at: Date and time of hosting's creation (RFC 3339 format).
        :param pulumi.Input[str] dns_status: The DNS status of the hosting.
        :param pulumi.Input[str] domain: The domain name of the hosting.
        :param pulumi.Input[str] email: The contact email of the client for the hosting.
        :param pulumi.Input[str] offer_id: The ID of the selected offer for the hosting.
        :param pulumi.Input[str] offer_name: The name of the active offer.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] option_ids: The IDs of the selected options for the hosting.
        :param pulumi.Input[Sequence[pulumi.Input[Union['WebhostingOptionArgs', 'WebhostingOptionArgsDict']]]] options: The active options of the hosting.
        :param pulumi.Input[str] organization_id: The organization ID the hosting is associated with.
        :param pulumi.Input[str] platform_hostname: The hostname of the host platform.
        :param pulumi.Input[int] platform_number: The number of the host platform.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the VPC is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Hosting.
        :param pulumi.Input[str] status: The hosting status.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the hosting.
        :param pulumi.Input[str] updated_at: Date and time of hosting's last update (RFC 3339 format).
        :param pulumi.Input[str] username: The main hosting cPanel username.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebhostingState.__new__(_WebhostingState)

        __props__.__dict__["cpanel_urls"] = cpanel_urls
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["dns_status"] = dns_status
        __props__.__dict__["domain"] = domain
        __props__.__dict__["email"] = email
        __props__.__dict__["offer_id"] = offer_id
        __props__.__dict__["offer_name"] = offer_name
        __props__.__dict__["option_ids"] = option_ids
        __props__.__dict__["options"] = options
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["platform_hostname"] = platform_hostname
        __props__.__dict__["platform_number"] = platform_number
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["username"] = username
        return Webhosting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cpanelUrls")
    def cpanel_urls(self) -> pulumi.Output[Sequence['outputs.WebhostingCpanelUrl']]:
        """
        The URL to connect to cPanel Dashboard and to Webmail interface.
        """
        return pulumi.get(self, "cpanel_urls")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date and time of hosting's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dnsStatus")
    def dns_status(self) -> pulumi.Output[str]:
        """
        The DNS status of the hosting.
        """
        return pulumi.get(self, "dns_status")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain name of the hosting.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The contact email of the client for the hosting.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="offerId")
    def offer_id(self) -> pulumi.Output[str]:
        """
        The ID of the selected offer for the hosting.
        """
        return pulumi.get(self, "offer_id")

    @property
    @pulumi.getter(name="offerName")
    def offer_name(self) -> pulumi.Output[str]:
        """
        The name of the active offer.
        """
        return pulumi.get(self, "offer_name")

    @property
    @pulumi.getter(name="optionIds")
    def option_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of the selected options for the hosting.
        """
        return pulumi.get(self, "option_ids")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Sequence['outputs.WebhostingOption']]:
        """
        The active options of the hosting.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the hosting is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="platformHostname")
    def platform_hostname(self) -> pulumi.Output[str]:
        """
        The hostname of the host platform.
        """
        return pulumi.get(self, "platform_hostname")

    @property
    @pulumi.getter(name="platformNumber")
    def platform_number(self) -> pulumi.Output[int]:
        """
        The number of the host platform.
        """
        return pulumi.get(self, "platform_number")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the VPC is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region of the Hosting.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The hosting status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        The tags associated with the hosting.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date and time of hosting's last update (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The main hosting cPanel username.
        """
        return pulumi.get(self, "username")

