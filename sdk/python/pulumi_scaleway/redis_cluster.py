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

__all__ = ['RedisClusterArgs', 'RedisCluster']

@pulumi.input_type
class RedisClusterArgs:
    def __init__(__self__, *,
                 node_type: pulumi.Input[str],
                 password: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 version: pulumi.Input[str],
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]]] = None,
                 cluster_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RedisCluster resource.
        :param pulumi.Input[str] node_type: Type of node to use for the cluster
        :param pulumi.Input[str] password: Password of the user
        :param pulumi.Input[str] user_name: Name of the user created when the cluster is created
        :param pulumi.Input[str] version: Redis version of the cluster
        :param pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]] acls: List of acl rules.
        :param pulumi.Input[int] cluster_size: Number of nodes for the cluster.
        :param pulumi.Input[str] name: Name of the redis cluster
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] settings: Map of settings to define for the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags ["tag1", "tag2", ...] attached to a redis cluster
        :param pulumi.Input[bool] tls_enabled: Whether or not TLS is enabled.
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        pulumi.set(__self__, "node_type", node_type)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "user_name", user_name)
        pulumi.set(__self__, "version", version)
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if cluster_size is not None:
            pulumi.set(__self__, "cluster_size", cluster_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Input[str]:
        """
        Type of node to use for the cluster
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password of the user
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        Name of the user created when the cluster is created
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Redis version of the cluster
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]]]:
        """
        List of acl rules.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="clusterSize")
    def cluster_size(self) -> Optional[pulumi.Input[int]]:
        """
        Number of nodes for the cluster.
        """
        return pulumi.get(self, "cluster_size")

    @cluster_size.setter
    def cluster_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the redis cluster
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
    def settings(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of settings to define for the cluster.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags ["tag1", "tag2", ...] attached to a redis cluster
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not TLS is enabled.
        """
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone you want to attach the resource to
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _RedisClusterState:
    def __init__(__self__, *,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]]] = None,
                 cluster_size: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RedisCluster resources.
        :param pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]] acls: List of acl rules.
        :param pulumi.Input[int] cluster_size: Number of nodes for the cluster.
        :param pulumi.Input[str] created_at: The date and time of the creation of the Redis cluster
        :param pulumi.Input[str] name: Name of the redis cluster
        :param pulumi.Input[str] node_type: Type of node to use for the cluster
        :param pulumi.Input[str] password: Password of the user
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] settings: Map of settings to define for the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags ["tag1", "tag2", ...] attached to a redis cluster
        :param pulumi.Input[bool] tls_enabled: Whether or not TLS is enabled.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the Redis cluster
        :param pulumi.Input[str] user_name: Name of the user created when the cluster is created
        :param pulumi.Input[str] version: Redis version of the cluster
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if cluster_size is not None:
            pulumi.set(__self__, "cluster_size", cluster_size)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_type is not None:
            pulumi.set(__self__, "node_type", node_type)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if settings is not None:
            pulumi.set(__self__, "settings", settings)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]]]:
        """
        List of acl rules.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RedisClusterAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="clusterSize")
    def cluster_size(self) -> Optional[pulumi.Input[int]]:
        """
        Number of nodes for the cluster.
        """
        return pulumi.get(self, "cluster_size")

    @cluster_size.setter
    def cluster_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_size", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the Redis cluster
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the redis cluster
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of node to use for the cluster
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password of the user
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

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
    def settings(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of settings to define for the cluster.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags ["tag1", "tag2", ...] attached to a redis cluster
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not TLS is enabled.
        """
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the Redis cluster
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the user created when the cluster is created
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Redis version of the cluster
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone you want to attach the resource to
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class RedisCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RedisClusterAclArgs']]]]] = None,
                 cluster_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a RedisCluster resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RedisClusterAclArgs']]]] acls: List of acl rules.
        :param pulumi.Input[int] cluster_size: Number of nodes for the cluster.
        :param pulumi.Input[str] name: Name of the redis cluster
        :param pulumi.Input[str] node_type: Type of node to use for the cluster
        :param pulumi.Input[str] password: Password of the user
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] settings: Map of settings to define for the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags ["tag1", "tag2", ...] attached to a redis cluster
        :param pulumi.Input[bool] tls_enabled: Whether or not TLS is enabled.
        :param pulumi.Input[str] user_name: Name of the user created when the cluster is created
        :param pulumi.Input[str] version: Redis version of the cluster
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RedisClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RedisCluster resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RedisClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RedisClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RedisClusterAclArgs']]]]] = None,
                 cluster_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RedisClusterArgs.__new__(RedisClusterArgs)

            __props__.__dict__["acls"] = acls
            __props__.__dict__["cluster_size"] = cluster_size
            __props__.__dict__["name"] = name
            if node_type is None and not opts.urn:
                raise TypeError("Missing required property 'node_type'")
            __props__.__dict__["node_type"] = node_type
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["settings"] = settings
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tls_enabled"] = tls_enabled
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(RedisCluster, __self__).__init__(
            'scaleway:index/redisCluster:RedisCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RedisClusterAclArgs']]]]] = None,
            cluster_size: Optional[pulumi.Input[int]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_type: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            settings: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tls_enabled: Optional[pulumi.Input[bool]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'RedisCluster':
        """
        Get an existing RedisCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RedisClusterAclArgs']]]] acls: List of acl rules.
        :param pulumi.Input[int] cluster_size: Number of nodes for the cluster.
        :param pulumi.Input[str] created_at: The date and time of the creation of the Redis cluster
        :param pulumi.Input[str] name: Name of the redis cluster
        :param pulumi.Input[str] node_type: Type of node to use for the cluster
        :param pulumi.Input[str] password: Password of the user
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] settings: Map of settings to define for the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags ["tag1", "tag2", ...] attached to a redis cluster
        :param pulumi.Input[bool] tls_enabled: Whether or not TLS is enabled.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the Redis cluster
        :param pulumi.Input[str] user_name: Name of the user created when the cluster is created
        :param pulumi.Input[str] version: Redis version of the cluster
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RedisClusterState.__new__(_RedisClusterState)

        __props__.__dict__["acls"] = acls
        __props__.__dict__["cluster_size"] = cluster_size
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["name"] = name
        __props__.__dict__["node_type"] = node_type
        __props__.__dict__["password"] = password
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["settings"] = settings
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tls_enabled"] = tls_enabled
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["user_name"] = user_name
        __props__.__dict__["version"] = version
        __props__.__dict__["zone"] = zone
        return RedisCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acls(self) -> pulumi.Output[Optional[Sequence['outputs.RedisClusterAcl']]]:
        """
        List of acl rules.
        """
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter(name="clusterSize")
    def cluster_size(self) -> pulumi.Output[int]:
        """
        Number of nodes for the cluster.
        """
        return pulumi.get(self, "cluster_size")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the Redis cluster
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the redis cluster
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Output[str]:
        """
        Type of node to use for the cluster
        """
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password of the user
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of settings to define for the cluster.
        """
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags ["tag1", "tag2", ...] attached to a redis cluster
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not TLS is enabled.
        """
        return pulumi.get(self, "tls_enabled")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the Redis cluster
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Name of the user created when the cluster is created
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Redis version of the cluster
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone you want to attach the resource to
        """
        return pulumi.get(self, "zone")

