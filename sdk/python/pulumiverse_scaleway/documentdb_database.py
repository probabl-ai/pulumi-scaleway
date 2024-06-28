# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DocumentdbDatabaseArgs', 'DocumentdbDatabase']

@pulumi.input_type
class DocumentdbDatabaseArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DocumentdbDatabase resource.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the documentdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database (e.g. `my-new-database`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _DocumentdbDatabaseState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DocumentdbDatabase resources.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database.
        :param pulumi.Input[bool] managed: Whether the database is managed or not.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
        :param pulumi.Input[str] owner: The name of the owner of the database.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[str] size: Size in gigabytes of the database.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if managed is not None:
            pulumi.set(__self__, "managed", managed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if size is not None:
            pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the documentdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def managed(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the database is managed or not.
        """
        return pulumi.get(self, "managed")

    @managed.setter
    def managed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "managed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database (e.g. `my-new-database`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the owner of the database.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[str]]:
        """
        Size in gigabytes of the database.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size", value)


class DocumentdbDatabase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway DocumentDB database.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        instance = scaleway.DocumentdbInstance("instance",
            node_type="docdb-play2-pico",
            engine="FerretDB-1",
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret",
            volume_size_in_gb=20)
        main = scaleway.DocumentdbDatabase("main", instance_id=instance.id)
        ```

        ## Import

        DocumentDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/documentdbDatabase:DocumentdbDatabase mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DocumentdbDatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway DocumentDB database.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        instance = scaleway.DocumentdbInstance("instance",
            node_type="docdb-play2-pico",
            engine="FerretDB-1",
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret",
            volume_size_in_gb=20)
        main = scaleway.DocumentdbDatabase("main", instance_id=instance.id)
        ```

        ## Import

        DocumentDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/documentdbDatabase:DocumentdbDatabase mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
        ```

        :param str resource_name: The name of the resource.
        :param DocumentdbDatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DocumentdbDatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DocumentdbDatabaseArgs.__new__(DocumentdbDatabaseArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["managed"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["size"] = None
        super(DocumentdbDatabase, __self__).__init__(
            'scaleway:index/documentdbDatabase:DocumentdbDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            managed: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[str]] = None) -> 'DocumentdbDatabase':
        """
        Get an existing DocumentdbDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database.
        :param pulumi.Input[bool] managed: Whether the database is managed or not.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
        :param pulumi.Input[str] owner: The name of the owner of the database.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[str] size: Size in gigabytes of the database.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DocumentdbDatabaseState.__new__(_DocumentdbDatabaseState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["managed"] = managed
        __props__.__dict__["name"] = name
        __props__.__dict__["owner"] = owner
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["size"] = size
        return DocumentdbDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the documentdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def managed(self) -> pulumi.Output[bool]:
        """
        Whether the database is managed or not.
        """
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the database (e.g. `my-new-database`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        The name of the owner of the database.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        """
        Size in gigabytes of the database.
        """
        return pulumi.get(self, "size")

