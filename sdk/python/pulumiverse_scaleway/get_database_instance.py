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

__all__ = [
    'GetDatabaseInstanceResult',
    'AwaitableGetDatabaseInstanceResult',
    'get_database_instance',
    'get_database_instance_output',
]

@pulumi.output_type
class GetDatabaseInstanceResult:
    """
    A collection of values returned by getDatabaseInstance.
    """
    def __init__(__self__, backup_same_region=None, backup_schedule_frequency=None, backup_schedule_retention=None, certificate=None, disable_backup=None, endpoint_ip=None, endpoint_port=None, engine=None, id=None, init_settings=None, instance_id=None, is_ha_cluster=None, load_balancers=None, name=None, node_type=None, organization_id=None, password=None, private_networks=None, project_id=None, read_replicas=None, region=None, settings=None, tags=None, user_name=None, volume_size_in_gb=None, volume_type=None):
        if backup_same_region and not isinstance(backup_same_region, bool):
            raise TypeError("Expected argument 'backup_same_region' to be a bool")
        pulumi.set(__self__, "backup_same_region", backup_same_region)
        if backup_schedule_frequency and not isinstance(backup_schedule_frequency, int):
            raise TypeError("Expected argument 'backup_schedule_frequency' to be a int")
        pulumi.set(__self__, "backup_schedule_frequency", backup_schedule_frequency)
        if backup_schedule_retention and not isinstance(backup_schedule_retention, int):
            raise TypeError("Expected argument 'backup_schedule_retention' to be a int")
        pulumi.set(__self__, "backup_schedule_retention", backup_schedule_retention)
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        pulumi.set(__self__, "certificate", certificate)
        if disable_backup and not isinstance(disable_backup, bool):
            raise TypeError("Expected argument 'disable_backup' to be a bool")
        pulumi.set(__self__, "disable_backup", disable_backup)
        if endpoint_ip and not isinstance(endpoint_ip, str):
            raise TypeError("Expected argument 'endpoint_ip' to be a str")
        pulumi.set(__self__, "endpoint_ip", endpoint_ip)
        if endpoint_port and not isinstance(endpoint_port, int):
            raise TypeError("Expected argument 'endpoint_port' to be a int")
        pulumi.set(__self__, "endpoint_port", endpoint_port)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if init_settings and not isinstance(init_settings, dict):
            raise TypeError("Expected argument 'init_settings' to be a dict")
        pulumi.set(__self__, "init_settings", init_settings)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if is_ha_cluster and not isinstance(is_ha_cluster, bool):
            raise TypeError("Expected argument 'is_ha_cluster' to be a bool")
        pulumi.set(__self__, "is_ha_cluster", is_ha_cluster)
        if load_balancers and not isinstance(load_balancers, list):
            raise TypeError("Expected argument 'load_balancers' to be a list")
        pulumi.set(__self__, "load_balancers", load_balancers)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_type and not isinstance(node_type, str):
            raise TypeError("Expected argument 'node_type' to be a str")
        pulumi.set(__self__, "node_type", node_type)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if private_networks and not isinstance(private_networks, list):
            raise TypeError("Expected argument 'private_networks' to be a list")
        pulumi.set(__self__, "private_networks", private_networks)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if read_replicas and not isinstance(read_replicas, list):
            raise TypeError("Expected argument 'read_replicas' to be a list")
        pulumi.set(__self__, "read_replicas", read_replicas)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if settings and not isinstance(settings, dict):
            raise TypeError("Expected argument 'settings' to be a dict")
        pulumi.set(__self__, "settings", settings)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)
        if volume_size_in_gb and not isinstance(volume_size_in_gb, int):
            raise TypeError("Expected argument 'volume_size_in_gb' to be a int")
        pulumi.set(__self__, "volume_size_in_gb", volume_size_in_gb)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="backupSameRegion")
    def backup_same_region(self) -> bool:
        return pulumi.get(self, "backup_same_region")

    @property
    @pulumi.getter(name="backupScheduleFrequency")
    def backup_schedule_frequency(self) -> int:
        return pulumi.get(self, "backup_schedule_frequency")

    @property
    @pulumi.getter(name="backupScheduleRetention")
    def backup_schedule_retention(self) -> int:
        return pulumi.get(self, "backup_schedule_retention")

    @property
    @pulumi.getter
    def certificate(self) -> str:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="disableBackup")
    def disable_backup(self) -> bool:
        return pulumi.get(self, "disable_backup")

    @property
    @pulumi.getter(name="endpointIp")
    def endpoint_ip(self) -> str:
        return pulumi.get(self, "endpoint_ip")

    @property
    @pulumi.getter(name="endpointPort")
    def endpoint_port(self) -> int:
        return pulumi.get(self, "endpoint_port")

    @property
    @pulumi.getter
    def engine(self) -> str:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="initSettings")
    def init_settings(self) -> Mapping[str, str]:
        return pulumi.get(self, "init_settings")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isHaCluster")
    def is_ha_cluster(self) -> bool:
        return pulumi.get(self, "is_ha_cluster")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> Sequence['outputs.GetDatabaseInstanceLoadBalancerResult']:
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> str:
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Sequence['outputs.GetDatabaseInstancePrivateNetworkResult']:
        return pulumi.get(self, "private_networks")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="readReplicas")
    def read_replicas(self) -> Sequence['outputs.GetDatabaseInstanceReadReplicaResult']:
        return pulumi.get(self, "read_replicas")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def settings(self) -> Mapping[str, str]:
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="volumeSizeInGb")
    def volume_size_in_gb(self) -> int:
        return pulumi.get(self, "volume_size_in_gb")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> str:
        return pulumi.get(self, "volume_type")


class AwaitableGetDatabaseInstanceResult(GetDatabaseInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseInstanceResult(
            backup_same_region=self.backup_same_region,
            backup_schedule_frequency=self.backup_schedule_frequency,
            backup_schedule_retention=self.backup_schedule_retention,
            certificate=self.certificate,
            disable_backup=self.disable_backup,
            endpoint_ip=self.endpoint_ip,
            endpoint_port=self.endpoint_port,
            engine=self.engine,
            id=self.id,
            init_settings=self.init_settings,
            instance_id=self.instance_id,
            is_ha_cluster=self.is_ha_cluster,
            load_balancers=self.load_balancers,
            name=self.name,
            node_type=self.node_type,
            organization_id=self.organization_id,
            password=self.password,
            private_networks=self.private_networks,
            project_id=self.project_id,
            read_replicas=self.read_replicas,
            region=self.region,
            settings=self.settings,
            tags=self.tags,
            user_name=self.user_name,
            volume_size_in_gb=self.volume_size_in_gb,
            volume_type=self.volume_type)


def get_database_instance(instance_id: Optional[str] = None,
                          name: Optional[str] = None,
                          project_id: Optional[str] = None,
                          region: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseInstanceResult:
    """
    Gets information about an RDB instance. For further information see our [developers website](https://developers.scaleway.com/en/products/rdb/api/#database-instance)


    :param str instance_id: The RDB instance ID.
           Only one of `name` and `instance_id` should be specified.
    :param str name: The name of the RDB instance.
           Only one of `name` and `instance_id` should be specified.
    :param str project_id: The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
    :param str region: `region`) The region in which the RDB instance exists.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getDatabaseInstance:getDatabaseInstance', __args__, opts=opts, typ=GetDatabaseInstanceResult).value

    return AwaitableGetDatabaseInstanceResult(
        backup_same_region=pulumi.get(__ret__, 'backup_same_region'),
        backup_schedule_frequency=pulumi.get(__ret__, 'backup_schedule_frequency'),
        backup_schedule_retention=pulumi.get(__ret__, 'backup_schedule_retention'),
        certificate=pulumi.get(__ret__, 'certificate'),
        disable_backup=pulumi.get(__ret__, 'disable_backup'),
        endpoint_ip=pulumi.get(__ret__, 'endpoint_ip'),
        endpoint_port=pulumi.get(__ret__, 'endpoint_port'),
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        init_settings=pulumi.get(__ret__, 'init_settings'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        is_ha_cluster=pulumi.get(__ret__, 'is_ha_cluster'),
        load_balancers=pulumi.get(__ret__, 'load_balancers'),
        name=pulumi.get(__ret__, 'name'),
        node_type=pulumi.get(__ret__, 'node_type'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        password=pulumi.get(__ret__, 'password'),
        private_networks=pulumi.get(__ret__, 'private_networks'),
        project_id=pulumi.get(__ret__, 'project_id'),
        read_replicas=pulumi.get(__ret__, 'read_replicas'),
        region=pulumi.get(__ret__, 'region'),
        settings=pulumi.get(__ret__, 'settings'),
        tags=pulumi.get(__ret__, 'tags'),
        user_name=pulumi.get(__ret__, 'user_name'),
        volume_size_in_gb=pulumi.get(__ret__, 'volume_size_in_gb'),
        volume_type=pulumi.get(__ret__, 'volume_type'))


@_utilities.lift_output_func(get_database_instance)
def get_database_instance_output(instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 name: Optional[pulumi.Input[Optional[str]]] = None,
                                 project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 region: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseInstanceResult]:
    """
    Gets information about an RDB instance. For further information see our [developers website](https://developers.scaleway.com/en/products/rdb/api/#database-instance)


    :param str instance_id: The RDB instance ID.
           Only one of `name` and `instance_id` should be specified.
    :param str name: The name of the RDB instance.
           Only one of `name` and `instance_id` should be specified.
    :param str project_id: The ID of the project the RDB instance is in. Can be used to filter instances when using `name`.
    :param str region: `region`) The region in which the RDB instance exists.
    """
    ...
