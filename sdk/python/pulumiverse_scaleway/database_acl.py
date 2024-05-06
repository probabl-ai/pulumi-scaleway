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

__all__ = ['DatabaseAclArgs', 'DatabaseAcl']

@pulumi.input_type
class DatabaseAclArgs:
    def __init__(__self__, *,
                 acl_rules: pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]],
                 instance_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabaseAcl resource.
        :param pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]] acl_rules: A list of ACLs (structure is described below)
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database ACL.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        """
        pulumi.set(__self__, "acl_rules", acl_rules)
        pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="aclRules")
    def acl_rules(self) -> pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]]:
        """
        A list of ACLs (structure is described below)
        """
        return pulumi.get(self, "acl_rules")

    @acl_rules.setter
    def acl_rules(self, value: pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]]):
        pulumi.set(self, "acl_rules", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database ACL.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _DatabaseAclState:
    def __init__(__self__, *,
                 acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabaseAcl resources.
        :param pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]] acl_rules: A list of ACLs (structure is described below)
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database ACL.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        """
        if acl_rules is not None:
            pulumi.set(__self__, "acl_rules", acl_rules)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="aclRules")
    def acl_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]]]:
        """
        A list of ACLs (structure is described below)
        """
        return pulumi.get(self, "acl_rules")

    @acl_rules.setter
    def acl_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseAclAclRuleArgs']]]]):
        pulumi.set(self, "acl_rules", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database ACL.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class DatabaseAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseAclAclRuleArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Database instance authorized IPs.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#acl-rules-allowed-ips).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.DatabaseAcl("main",
            instance_id=scaleway_rdb_instance["main"]["id"],
            acl_rules=[scaleway.DatabaseAclAclRuleArgs(
                ip="1.2.3.4/32",
                description="foo",
            )])
        ```

        ## Import

        Database Instance can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/databaseAcl:DatabaseAcl acl01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseAclAclRuleArgs']]]] acl_rules: A list of ACLs (structure is described below)
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database ACL.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Database instance authorized IPs.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#acl-rules-allowed-ips).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.DatabaseAcl("main",
            instance_id=scaleway_rdb_instance["main"]["id"],
            acl_rules=[scaleway.DatabaseAclAclRuleArgs(
                ip="1.2.3.4/32",
                description="foo",
            )])
        ```

        ## Import

        Database Instance can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/databaseAcl:DatabaseAcl acl01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param DatabaseAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseAclAclRuleArgs']]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseAclArgs.__new__(DatabaseAclArgs)

            if acl_rules is None and not opts.urn:
                raise TypeError("Missing required property 'acl_rules'")
            __props__.__dict__["acl_rules"] = acl_rules
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region"] = region
        super(DatabaseAcl, __self__).__init__(
            'scaleway:index/databaseAcl:DatabaseAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseAclAclRuleArgs']]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'DatabaseAcl':
        """
        Get an existing DatabaseAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatabaseAclAclRuleArgs']]]] acl_rules: A list of ACLs (structure is described below)
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database ACL.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseAclState.__new__(_DatabaseAclState)

        __props__.__dict__["acl_rules"] = acl_rules
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["region"] = region
        return DatabaseAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aclRules")
    def acl_rules(self) -> pulumi.Output[Sequence['outputs.DatabaseAclAclRule']]:
        """
        A list of ACLs (structure is described below)
        """
        return pulumi.get(self, "acl_rules")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database ACL.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

