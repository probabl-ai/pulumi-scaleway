# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamGroupMembershipArgs', 'IamGroupMembership']

@pulumi.input_type
class IamGroupMembershipArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[str],
                 application_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamGroupMembership resource.
        :param pulumi.Input[str] group_id: ID of the group to add members to.
        :param pulumi.Input[str] application_id: The ID of the application that will be added to the group.
        :param pulumi.Input[str] user_id: The ID of the user that will be added to the group
               
               - > Only one of `application_id` or `user_id` must be specified
        """
        pulumi.set(__self__, "group_id", group_id)
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        ID of the group to add members to.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the application that will be added to the group.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the user that will be added to the group

        - > Only one of `application_id` or `user_id` must be specified
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _IamGroupMembershipState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamGroupMembership resources.
        :param pulumi.Input[str] application_id: The ID of the application that will be added to the group.
        :param pulumi.Input[str] group_id: ID of the group to add members to.
        :param pulumi.Input[str] user_id: The ID of the user that will be added to the group
               
               - > Only one of `application_id` or `user_id` must be specified
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the application that will be added to the group.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the group to add members to.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the user that will be added to the group

        - > Only one of `application_id` or `user_id` must be specified
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class IamGroupMembership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Add members to an IAM group.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#groups-f592eb).

        ## Example Usage

        ### Application Membership

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        group = scaleway.IamGroup("group", external_membership=True)
        app = scaleway.IamApplication("app")
        member = scaleway.IamGroupMembership("member",
            group_id=group.id,
            application_id=app.id)
        ```

        ## Import

        IAM group memberships can be imported using two format:

        - For user: `{group_id}/user/{user_id}`

        - For application: `{group_id}/app/{application_id}`

        bash

        ```sh
        $ pulumi import scaleway:index/iamGroupMembership:IamGroupMembership app 11111111-1111-1111-1111-111111111111/app/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The ID of the application that will be added to the group.
        :param pulumi.Input[str] group_id: ID of the group to add members to.
        :param pulumi.Input[str] user_id: The ID of the user that will be added to the group
               
               - > Only one of `application_id` or `user_id` must be specified
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamGroupMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Add members to an IAM group.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#groups-f592eb).

        ## Example Usage

        ### Application Membership

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        group = scaleway.IamGroup("group", external_membership=True)
        app = scaleway.IamApplication("app")
        member = scaleway.IamGroupMembership("member",
            group_id=group.id,
            application_id=app.id)
        ```

        ## Import

        IAM group memberships can be imported using two format:

        - For user: `{group_id}/user/{user_id}`

        - For application: `{group_id}/app/{application_id}`

        bash

        ```sh
        $ pulumi import scaleway:index/iamGroupMembership:IamGroupMembership app 11111111-1111-1111-1111-111111111111/app/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IamGroupMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamGroupMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamGroupMembershipArgs.__new__(IamGroupMembershipArgs)

            __props__.__dict__["application_id"] = application_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["user_id"] = user_id
        super(IamGroupMembership, __self__).__init__(
            'scaleway:index/iamGroupMembership:IamGroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'IamGroupMembership':
        """
        Get an existing IamGroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The ID of the application that will be added to the group.
        :param pulumi.Input[str] group_id: ID of the group to add members to.
        :param pulumi.Input[str] user_id: The ID of the user that will be added to the group
               
               - > Only one of `application_id` or `user_id` must be specified
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamGroupMembershipState.__new__(_IamGroupMembershipState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["user_id"] = user_id
        return IamGroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the application that will be added to the group.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        ID of the group to add members to.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the user that will be added to the group

        - > Only one of `application_id` or `user_id` must be specified
        """
        return pulumi.get(self, "user_id")

