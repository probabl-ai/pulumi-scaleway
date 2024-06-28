# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MnqNatsCredentialsArgs', 'MnqNatsCredentials']

@pulumi.input_type
class MnqNatsCredentialsArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MnqNatsCredentials resource.
        :param pulumi.Input[str] account_id: The ID of the NATS account the credentials are generated from
        :param pulumi.Input[str] name: The unique name of the NATS credentials.
        :param pulumi.Input[str] region: `region`). The region
               in which the account exists.
        """
        pulumi.set(__self__, "account_id", account_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        """
        The ID of the NATS account the credentials are generated from
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the NATS credentials.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which the account exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _MnqNatsCredentialsState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MnqNatsCredentials resources.
        :param pulumi.Input[str] account_id: The ID of the NATS account the credentials are generated from
        :param pulumi.Input[str] file: The content of the credentials file.
        :param pulumi.Input[str] name: The unique name of the NATS credentials.
        :param pulumi.Input[str] region: `region`). The region
               in which the account exists.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if file is not None:
            pulumi.set(__self__, "file", file)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the NATS account the credentials are generated from
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the credentials file.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the NATS credentials.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which the account exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class MnqNatsCredentials(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Messaging and Queuing NATS credentials.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_mnq_nats_account = scaleway.MnqNatsAccount("mainMnqNatsAccount")
        main_mnq_nats_credentials = scaleway.MnqNatsCredentials("mainMnqNatsCredentials", account_id=main_mnq_nats_account.id)
        ```

        ## Import

        Namespaces can be imported using `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqNatsCredentials:MnqNatsCredentials main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the NATS account the credentials are generated from
        :param pulumi.Input[str] name: The unique name of the NATS credentials.
        :param pulumi.Input[str] region: `region`). The region
               in which the account exists.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MnqNatsCredentialsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Messaging and Queuing NATS credentials.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_mnq_nats_account = scaleway.MnqNatsAccount("mainMnqNatsAccount")
        main_mnq_nats_credentials = scaleway.MnqNatsCredentials("mainMnqNatsCredentials", account_id=main_mnq_nats_account.id)
        ```

        ## Import

        Namespaces can be imported using `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqNatsCredentials:MnqNatsCredentials main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param MnqNatsCredentialsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqNatsCredentialsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqNatsCredentialsArgs.__new__(MnqNatsCredentialsArgs)

            if account_id is None and not opts.urn:
                raise TypeError("Missing required property 'account_id'")
            __props__.__dict__["account_id"] = account_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["file"] = None
        super(MnqNatsCredentials, __self__).__init__(
            'scaleway:index/mnqNatsCredentials:MnqNatsCredentials',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            file: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'MnqNatsCredentials':
        """
        Get an existing MnqNatsCredentials resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of the NATS account the credentials are generated from
        :param pulumi.Input[str] file: The content of the credentials file.
        :param pulumi.Input[str] name: The unique name of the NATS credentials.
        :param pulumi.Input[str] region: `region`). The region
               in which the account exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqNatsCredentialsState.__new__(_MnqNatsCredentialsState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["file"] = file
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        return MnqNatsCredentials(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The ID of the NATS account the credentials are generated from
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output[str]:
        """
        The content of the credentials file.
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the NATS credentials.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region
        in which the account exists.
        """
        return pulumi.get(self, "region")

