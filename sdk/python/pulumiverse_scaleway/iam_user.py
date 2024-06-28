# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamUserArgs', 'IamUser']

@pulumi.input_type
class IamUserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 organization_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamUser resource.
        :param pulumi.Input[str] email: The email of the IAM user.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the user is associated with.
        """
        pulumi.set(__self__, "email", email)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        The email of the IAM user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        `organization_id`) The ID of the organization the user is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)


@pulumi.input_type
class _IamUserState:
    def __init__(__self__, *,
                 account_root_user_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 deletable: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 last_login_at: Optional[pulumi.Input[str]] = None,
                 mfa: Optional[pulumi.Input[bool]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamUser resources.
        :param pulumi.Input[str] account_root_user_id: The ID of the account root user associated with the user.
        :param pulumi.Input[str] created_at: The date and time of the creation of the IAM user.
        :param pulumi.Input[bool] deletable: Whether the IAM user is deletable.
        :param pulumi.Input[str] email: The email of the IAM user.
        :param pulumi.Input[str] last_login_at: The date of the last login.
        :param pulumi.Input[bool] mfa: Whether the MFA is enabled.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the user is associated with.
        :param pulumi.Input[str] status: The status of user invitation. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        :param pulumi.Input[str] type: The type of user. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        :param pulumi.Input[str] updated_at: The date and time of the last update of the IAM user.
        """
        if account_root_user_id is not None:
            pulumi.set(__self__, "account_root_user_id", account_root_user_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if deletable is not None:
            pulumi.set(__self__, "deletable", deletable)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if last_login_at is not None:
            pulumi.set(__self__, "last_login_at", last_login_at)
        if mfa is not None:
            pulumi.set(__self__, "mfa", mfa)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="accountRootUserId")
    def account_root_user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the account root user associated with the user.
        """
        return pulumi.get(self, "account_root_user_id")

    @account_root_user_id.setter
    def account_root_user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_root_user_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the IAM user.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def deletable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the IAM user is deletable.
        """
        return pulumi.get(self, "deletable")

    @deletable.setter
    def deletable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletable", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email of the IAM user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="lastLoginAt")
    def last_login_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date of the last login.
        """
        return pulumi.get(self, "last_login_at")

    @last_login_at.setter
    def last_login_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_login_at", value)

    @property
    @pulumi.getter
    def mfa(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the MFA is enabled.
        """
        return pulumi.get(self, "mfa")

    @mfa.setter
    def mfa(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mfa", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        `organization_id`) The ID of the organization the user is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of user invitation. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of user. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the IAM user.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class IamUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway IAM Users.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#path-users-list-users-of-an-organization).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        basic = scaleway.IamUser("basic", email="test@test.com")
        ```

        ## Import

        IAM users can be imported using the `{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iamUser:IamUser basic 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: The email of the IAM user.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the user is associated with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway IAM Users.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#path-users-list-users-of-an-organization).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        basic = scaleway.IamUser("basic", email="test@test.com")
        ```

        ## Import

        IAM users can be imported using the `{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iamUser:IamUser basic 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IamUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamUserArgs.__new__(IamUserArgs)

            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["account_root_user_id"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["deletable"] = None
            __props__.__dict__["last_login_at"] = None
            __props__.__dict__["mfa"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["updated_at"] = None
        super(IamUser, __self__).__init__(
            'scaleway:index/iamUser:IamUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_root_user_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            deletable: Optional[pulumi.Input[bool]] = None,
            email: Optional[pulumi.Input[str]] = None,
            last_login_at: Optional[pulumi.Input[str]] = None,
            mfa: Optional[pulumi.Input[bool]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'IamUser':
        """
        Get an existing IamUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_root_user_id: The ID of the account root user associated with the user.
        :param pulumi.Input[str] created_at: The date and time of the creation of the IAM user.
        :param pulumi.Input[bool] deletable: Whether the IAM user is deletable.
        :param pulumi.Input[str] email: The email of the IAM user.
        :param pulumi.Input[str] last_login_at: The date of the last login.
        :param pulumi.Input[bool] mfa: Whether the MFA is enabled.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the user is associated with.
        :param pulumi.Input[str] status: The status of user invitation. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        :param pulumi.Input[str] type: The type of user. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        :param pulumi.Input[str] updated_at: The date and time of the last update of the IAM user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamUserState.__new__(_IamUserState)

        __props__.__dict__["account_root_user_id"] = account_root_user_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["deletable"] = deletable
        __props__.__dict__["email"] = email
        __props__.__dict__["last_login_at"] = last_login_at
        __props__.__dict__["mfa"] = mfa
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        return IamUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountRootUserId")
    def account_root_user_id(self) -> pulumi.Output[str]:
        """
        The ID of the account root user associated with the user.
        """
        return pulumi.get(self, "account_root_user_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the IAM user.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def deletable(self) -> pulumi.Output[bool]:
        """
        Whether the IAM user is deletable.
        """
        return pulumi.get(self, "deletable")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The email of the IAM user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="lastLoginAt")
    def last_login_at(self) -> pulumi.Output[str]:
        """
        The date of the last login.
        """
        return pulumi.get(self, "last_login_at")

    @property
    @pulumi.getter
    def mfa(self) -> pulumi.Output[bool]:
        """
        Whether the MFA is enabled.
        """
        return pulumi.get(self, "mfa")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        `organization_id`) The ID of the organization the user is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of user invitation. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of user. Check the possible values in the [API doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the IAM user.
        """
        return pulumi.get(self, "updated_at")

