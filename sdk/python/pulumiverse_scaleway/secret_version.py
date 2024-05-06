# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SecretVersionArgs', 'SecretVersion']

@pulumi.input_type
class SecretVersionArgs:
    def __init__(__self__, *,
                 data: pulumi.Input[str],
                 secret_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretVersion resource.
        :param pulumi.Input[str] data: The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
        :param pulumi.Input[str] secret_id: The Secret ID associated wit the secret version.
        :param pulumi.Input[str] description: Description of the secret version (e.g. `my-new-description`).
        :param pulumi.Input[str] region: `region`) The region
               in which the resource exists.
        """
        pulumi.set(__self__, "data", data)
        pulumi.set(__self__, "secret_id", secret_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input[str]:
        """
        The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input[str]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Input[str]:
        """
        The Secret ID associated wit the secret version.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the secret version (e.g. `my-new-description`).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _SecretVersionState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 revision: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretVersion resources.
        :param pulumi.Input[str] created_at: Date and time of secret version's creation (RFC 3339 format).
        :param pulumi.Input[str] data: The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
        :param pulumi.Input[str] description: Description of the secret version (e.g. `my-new-description`).
        :param pulumi.Input[str] region: `region`) The region
               in which the resource exists.
        :param pulumi.Input[str] revision: The revision for this Secret Version.
        :param pulumi.Input[str] secret_id: The Secret ID associated wit the secret version.
        :param pulumi.Input[str] status: The status of the Secret Version.
        :param pulumi.Input[str] updated_at: Date and time of secret version's last update (RFC 3339 format).
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if revision is not None:
            pulumi.set(__self__, "revision", revision)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of secret version's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the secret version (e.g. `my-new-description`).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def revision(self) -> Optional[pulumi.Input[str]]:
        """
        The revision for this Secret Version.
        """
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "revision", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Secret ID associated wit the secret version.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the Secret Version.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of secret version's last update (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class SecretVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Secret Versions.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/#secret-versions-079501).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.Secret("main",
            description="barr",
            tags=[
                "foo",
                "terraform",
            ])
        v1 = scaleway.SecretVersion("v1",
            description="version1",
            secret_id=main.id,
            data="my_new_secret")
        ```

        ## Import

        The Secret Version can be imported using the `{region}/{id}/{revision}`, e.g.

        ~> **Important:** Be aware if you import with revision `latest` you will overwrite the version you used before.

        bash

        ```sh
        $ pulumi import scaleway:index/secretVersion:SecretVersion main fr-par/11111111-1111-1111-1111-111111111111/2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
        :param pulumi.Input[str] description: Description of the secret version (e.g. `my-new-description`).
        :param pulumi.Input[str] region: `region`) The region
               in which the resource exists.
        :param pulumi.Input[str] secret_id: The Secret ID associated wit the secret version.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Secret Versions.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/#secret-versions-079501).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.Secret("main",
            description="barr",
            tags=[
                "foo",
                "terraform",
            ])
        v1 = scaleway.SecretVersion("v1",
            description="version1",
            secret_id=main.id,
            data="my_new_secret")
        ```

        ## Import

        The Secret Version can be imported using the `{region}/{id}/{revision}`, e.g.

        ~> **Important:** Be aware if you import with revision `latest` you will overwrite the version you used before.

        bash

        ```sh
        $ pulumi import scaleway:index/secretVersion:SecretVersion main fr-par/11111111-1111-1111-1111-111111111111/2
        ```

        :param str resource_name: The name of the resource.
        :param SecretVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretVersionArgs.__new__(SecretVersionArgs)

            if data is None and not opts.urn:
                raise TypeError("Missing required property 'data'")
            __props__.__dict__["data"] = None if data is None else pulumi.Output.secret(data)
            __props__.__dict__["description"] = description
            __props__.__dict__["region"] = region
            if secret_id is None and not opts.urn:
                raise TypeError("Missing required property 'secret_id'")
            __props__.__dict__["secret_id"] = secret_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["revision"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["data"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SecretVersion, __self__).__init__(
            'scaleway:index/secretVersion:SecretVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            data: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            revision: Optional[pulumi.Input[str]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'SecretVersion':
        """
        Get an existing SecretVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Date and time of secret version's creation (RFC 3339 format).
        :param pulumi.Input[str] data: The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
        :param pulumi.Input[str] description: Description of the secret version (e.g. `my-new-description`).
        :param pulumi.Input[str] region: `region`) The region
               in which the resource exists.
        :param pulumi.Input[str] revision: The revision for this Secret Version.
        :param pulumi.Input[str] secret_id: The Secret ID associated wit the secret version.
        :param pulumi.Input[str] status: The status of the Secret Version.
        :param pulumi.Input[str] updated_at: Date and time of secret version's last update (RFC 3339 format).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretVersionState.__new__(_SecretVersionState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["data"] = data
        __props__.__dict__["description"] = description
        __props__.__dict__["region"] = region
        __props__.__dict__["revision"] = revision
        __props__.__dict__["secret_id"] = secret_id
        __props__.__dict__["status"] = status
        __props__.__dict__["updated_at"] = updated_at
        return SecretVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date and time of secret version's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[str]:
        """
        The data payload of the secret version. Must be no larger than 64KiB. (e.g. `my-secret-version-payload`). more on the data section
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the secret version (e.g. `my-new-description`).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region
        in which the resource exists.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Output[str]:
        """
        The revision for this Secret Version.
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        """
        The Secret ID associated wit the secret version.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the Secret Version.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date and time of secret version's last update (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

