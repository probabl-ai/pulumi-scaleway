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

__all__ = ['InstanceSecurityGroupArgs', 'InstanceSecurityGroup']

@pulumi.input_type
class InstanceSecurityGroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_default_security: Optional[pulumi.Input[bool]] = None,
                 external_rules: Optional[pulumi.Input[bool]] = None,
                 inbound_default_policy: Optional[pulumi.Input[str]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_default_policy: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 stateful: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceSecurityGroup resource.
        :param pulumi.Input[str] description: The description of the security group.
        :param pulumi.Input[bool] enable_default_security: Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        :param pulumi.Input[bool] external_rules: A boolean to specify whether to use instance_security_group_rules.
               If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        :param pulumi.Input[str] inbound_default_policy: The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] name: The name of the security group.
        :param pulumi.Input[str] outbound_default_policy: The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the security group is associated with.
        :param pulumi.Input[bool] stateful: A boolean to specify whether the security group should be stateful or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the security group
        :param pulumi.Input[str] zone: `zone`) The zone in which the security group should be created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_default_security is not None:
            pulumi.set(__self__, "enable_default_security", enable_default_security)
        if external_rules is not None:
            pulumi.set(__self__, "external_rules", external_rules)
        if inbound_default_policy is not None:
            pulumi.set(__self__, "inbound_default_policy", inbound_default_policy)
        if inbound_rules is not None:
            pulumi.set(__self__, "inbound_rules", inbound_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if outbound_default_policy is not None:
            pulumi.set(__self__, "outbound_default_policy", outbound_default_policy)
        if outbound_rules is not None:
            pulumi.set(__self__, "outbound_rules", outbound_rules)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if stateful is not None:
            pulumi.set(__self__, "stateful", stateful)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the security group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableDefaultSecurity")
    def enable_default_security(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        """
        return pulumi.get(self, "enable_default_security")

    @enable_default_security.setter
    def enable_default_security(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_default_security", value)

    @property
    @pulumi.getter(name="externalRules")
    def external_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether to use instance_security_group_rules.
        If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        """
        return pulumi.get(self, "external_rules")

    @external_rules.setter
    def external_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_rules", value)

    @property
    @pulumi.getter(name="inboundDefaultPolicy")
    def inbound_default_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "inbound_default_policy")

    @inbound_default_policy.setter
    def inbound_default_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inbound_default_policy", value)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]]]:
        """
        A list of inbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "inbound_rules")

    @inbound_rules.setter
    def inbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]]]):
        pulumi.set(self, "inbound_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outboundDefaultPolicy")
    def outbound_default_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "outbound_default_policy")

    @outbound_default_policy.setter
    def outbound_default_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outbound_default_policy", value)

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]]]:
        """
        A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "outbound_rules")

    @outbound_rules.setter
    def outbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]]]):
        pulumi.set(self, "outbound_rules", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the security group is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def stateful(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether the security group should be stateful or not.
        """
        return pulumi.get(self, "stateful")

    @stateful.setter
    def stateful(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "stateful", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the security group
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the security group should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstanceSecurityGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_default_security: Optional[pulumi.Input[bool]] = None,
                 external_rules: Optional[pulumi.Input[bool]] = None,
                 inbound_default_policy: Optional[pulumi.Input[str]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 outbound_default_policy: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 stateful: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceSecurityGroup resources.
        :param pulumi.Input[str] description: The description of the security group.
        :param pulumi.Input[bool] enable_default_security: Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        :param pulumi.Input[bool] external_rules: A boolean to specify whether to use instance_security_group_rules.
               If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        :param pulumi.Input[str] inbound_default_policy: The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] name: The name of the security group.
        :param pulumi.Input[str] organization_id: The organization ID the security group is associated with.
        :param pulumi.Input[str] outbound_default_policy: The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the security group is associated with.
        :param pulumi.Input[bool] stateful: A boolean to specify whether the security group should be stateful or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the security group
        :param pulumi.Input[str] zone: `zone`) The zone in which the security group should be created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_default_security is not None:
            pulumi.set(__self__, "enable_default_security", enable_default_security)
        if external_rules is not None:
            pulumi.set(__self__, "external_rules", external_rules)
        if inbound_default_policy is not None:
            pulumi.set(__self__, "inbound_default_policy", inbound_default_policy)
        if inbound_rules is not None:
            pulumi.set(__self__, "inbound_rules", inbound_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if outbound_default_policy is not None:
            pulumi.set(__self__, "outbound_default_policy", outbound_default_policy)
        if outbound_rules is not None:
            pulumi.set(__self__, "outbound_rules", outbound_rules)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if stateful is not None:
            pulumi.set(__self__, "stateful", stateful)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the security group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableDefaultSecurity")
    def enable_default_security(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        """
        return pulumi.get(self, "enable_default_security")

    @enable_default_security.setter
    def enable_default_security(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_default_security", value)

    @property
    @pulumi.getter(name="externalRules")
    def external_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether to use instance_security_group_rules.
        If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        """
        return pulumi.get(self, "external_rules")

    @external_rules.setter
    def external_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_rules", value)

    @property
    @pulumi.getter(name="inboundDefaultPolicy")
    def inbound_default_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "inbound_default_policy")

    @inbound_default_policy.setter
    def inbound_default_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inbound_default_policy", value)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]]]:
        """
        A list of inbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "inbound_rules")

    @inbound_rules.setter
    def inbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupInboundRuleArgs']]]]):
        pulumi.set(self, "inbound_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the security group is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="outboundDefaultPolicy")
    def outbound_default_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "outbound_default_policy")

    @outbound_default_policy.setter
    def outbound_default_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outbound_default_policy", value)

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]]]:
        """
        A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "outbound_rules")

    @outbound_rules.setter
    def outbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupOutboundRuleArgs']]]]):
        pulumi.set(self, "outbound_rules", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the security group is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def stateful(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether the security group should be stateful or not.
        """
        return pulumi.get(self, "stateful")

    @stateful.setter
    def stateful(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "stateful", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the security group
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the security group should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstanceSecurityGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_default_security: Optional[pulumi.Input[bool]] = None,
                 external_rules: Optional[pulumi.Input[bool]] = None,
                 inbound_default_policy: Optional[pulumi.Input[str]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupInboundRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_default_policy: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupOutboundRuleArgs']]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 stateful: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Instance security group can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceSecurityGroup:InstanceSecurityGroup web fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the security group.
        :param pulumi.Input[bool] enable_default_security: Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        :param pulumi.Input[bool] external_rules: A boolean to specify whether to use instance_security_group_rules.
               If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        :param pulumi.Input[str] inbound_default_policy: The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupInboundRuleArgs']]]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] name: The name of the security group.
        :param pulumi.Input[str] outbound_default_policy: The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupOutboundRuleArgs']]]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the security group is associated with.
        :param pulumi.Input[bool] stateful: A boolean to specify whether the security group should be stateful or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the security group
        :param pulumi.Input[str] zone: `zone`) The zone in which the security group should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InstanceSecurityGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Instance security group can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceSecurityGroup:InstanceSecurityGroup web fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceSecurityGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceSecurityGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_default_security: Optional[pulumi.Input[bool]] = None,
                 external_rules: Optional[pulumi.Input[bool]] = None,
                 inbound_default_policy: Optional[pulumi.Input[str]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupInboundRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_default_policy: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupOutboundRuleArgs']]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 stateful: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceSecurityGroupArgs.__new__(InstanceSecurityGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["enable_default_security"] = enable_default_security
            __props__.__dict__["external_rules"] = external_rules
            __props__.__dict__["inbound_default_policy"] = inbound_default_policy
            __props__.__dict__["inbound_rules"] = inbound_rules
            __props__.__dict__["name"] = name
            __props__.__dict__["outbound_default_policy"] = outbound_default_policy
            __props__.__dict__["outbound_rules"] = outbound_rules
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["stateful"] = stateful
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
            __props__.__dict__["organization_id"] = None
        super(InstanceSecurityGroup, __self__).__init__(
            'scaleway:index/instanceSecurityGroup:InstanceSecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable_default_security: Optional[pulumi.Input[bool]] = None,
            external_rules: Optional[pulumi.Input[bool]] = None,
            inbound_default_policy: Optional[pulumi.Input[str]] = None,
            inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupInboundRuleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            outbound_default_policy: Optional[pulumi.Input[str]] = None,
            outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupOutboundRuleArgs']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            stateful: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceSecurityGroup':
        """
        Get an existing InstanceSecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the security group.
        :param pulumi.Input[bool] enable_default_security: Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        :param pulumi.Input[bool] external_rules: A boolean to specify whether to use instance_security_group_rules.
               If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        :param pulumi.Input[str] inbound_default_policy: The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupInboundRuleArgs']]]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] name: The name of the security group.
        :param pulumi.Input[str] organization_id: The organization ID the security group is associated with.
        :param pulumi.Input[str] outbound_default_policy: The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupOutboundRuleArgs']]]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the security group is associated with.
        :param pulumi.Input[bool] stateful: A boolean to specify whether the security group should be stateful or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the security group
        :param pulumi.Input[str] zone: `zone`) The zone in which the security group should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceSecurityGroupState.__new__(_InstanceSecurityGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["enable_default_security"] = enable_default_security
        __props__.__dict__["external_rules"] = external_rules
        __props__.__dict__["inbound_default_policy"] = inbound_default_policy
        __props__.__dict__["inbound_rules"] = inbound_rules
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["outbound_default_policy"] = outbound_default_policy
        __props__.__dict__["outbound_rules"] = outbound_rules
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["stateful"] = stateful
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone"] = zone
        return InstanceSecurityGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the security group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableDefaultSecurity")
    def enable_default_security(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, [open a support ticket](https://console.scaleway.com/support/tickets).
        """
        return pulumi.get(self, "enable_default_security")

    @property
    @pulumi.getter(name="externalRules")
    def external_rules(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean to specify whether to use instance_security_group_rules.
        If `external_rules` is set to `true`, `inbound_rule` and `outbound_rule` can not be set directly in the security group.
        """
        return pulumi.get(self, "external_rules")

    @property
    @pulumi.getter(name="inboundDefaultPolicy")
    def inbound_default_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "inbound_default_policy")

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceSecurityGroupInboundRule']]]:
        """
        A list of inbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the security group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the security group is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="outboundDefaultPolicy")
    def outbound_default_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "outbound_default_policy")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceSecurityGroupOutboundRule']]]:
        """
        A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the security group is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def stateful(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean to specify whether the security group should be stateful or not.
        """
        return pulumi.get(self, "stateful")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The tags associated with the security group
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the security group should be created.
        """
        return pulumi.get(self, "zone")

