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
    'GetKubernetesNodePoolResult',
    'AwaitableGetKubernetesNodePoolResult',
    'get_kubernetes_node_pool',
    'get_kubernetes_node_pool_output',
]

@pulumi.output_type
class GetKubernetesNodePoolResult:
    """
    A collection of values returned by getKubernetesNodePool.
    """
    def __init__(__self__, autohealing=None, autoscaling=None, cluster_id=None, container_runtime=None, created_at=None, current_size=None, id=None, kubelet_args=None, max_size=None, min_size=None, name=None, node_type=None, nodes=None, placement_group_id=None, pool_id=None, public_ip_disabled=None, region=None, root_volume_size_in_gb=None, root_volume_type=None, size=None, status=None, tags=None, updated_at=None, upgrade_policies=None, version=None, wait_for_pool_ready=None, zone=None):
        if autohealing and not isinstance(autohealing, bool):
            raise TypeError("Expected argument 'autohealing' to be a bool")
        pulumi.set(__self__, "autohealing", autohealing)
        if autoscaling and not isinstance(autoscaling, bool):
            raise TypeError("Expected argument 'autoscaling' to be a bool")
        pulumi.set(__self__, "autoscaling", autoscaling)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if container_runtime and not isinstance(container_runtime, str):
            raise TypeError("Expected argument 'container_runtime' to be a str")
        pulumi.set(__self__, "container_runtime", container_runtime)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if current_size and not isinstance(current_size, int):
            raise TypeError("Expected argument 'current_size' to be a int")
        pulumi.set(__self__, "current_size", current_size)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kubelet_args and not isinstance(kubelet_args, dict):
            raise TypeError("Expected argument 'kubelet_args' to be a dict")
        pulumi.set(__self__, "kubelet_args", kubelet_args)
        if max_size and not isinstance(max_size, int):
            raise TypeError("Expected argument 'max_size' to be a int")
        pulumi.set(__self__, "max_size", max_size)
        if min_size and not isinstance(min_size, int):
            raise TypeError("Expected argument 'min_size' to be a int")
        pulumi.set(__self__, "min_size", min_size)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_type and not isinstance(node_type, str):
            raise TypeError("Expected argument 'node_type' to be a str")
        pulumi.set(__self__, "node_type", node_type)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)
        if placement_group_id and not isinstance(placement_group_id, str):
            raise TypeError("Expected argument 'placement_group_id' to be a str")
        pulumi.set(__self__, "placement_group_id", placement_group_id)
        if pool_id and not isinstance(pool_id, str):
            raise TypeError("Expected argument 'pool_id' to be a str")
        pulumi.set(__self__, "pool_id", pool_id)
        if public_ip_disabled and not isinstance(public_ip_disabled, bool):
            raise TypeError("Expected argument 'public_ip_disabled' to be a bool")
        pulumi.set(__self__, "public_ip_disabled", public_ip_disabled)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if root_volume_size_in_gb and not isinstance(root_volume_size_in_gb, int):
            raise TypeError("Expected argument 'root_volume_size_in_gb' to be a int")
        pulumi.set(__self__, "root_volume_size_in_gb", root_volume_size_in_gb)
        if root_volume_type and not isinstance(root_volume_type, str):
            raise TypeError("Expected argument 'root_volume_type' to be a str")
        pulumi.set(__self__, "root_volume_type", root_volume_type)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if upgrade_policies and not isinstance(upgrade_policies, list):
            raise TypeError("Expected argument 'upgrade_policies' to be a list")
        pulumi.set(__self__, "upgrade_policies", upgrade_policies)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if wait_for_pool_ready and not isinstance(wait_for_pool_ready, bool):
            raise TypeError("Expected argument 'wait_for_pool_ready' to be a bool")
        pulumi.set(__self__, "wait_for_pool_ready", wait_for_pool_ready)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def autohealing(self) -> bool:
        """
        True if the autohealing feature is enabled for this pool.
        """
        return pulumi.get(self, "autohealing")

    @property
    @pulumi.getter
    def autoscaling(self) -> bool:
        """
        True if the autoscaling feature is enabled for this pool.
        """
        return pulumi.get(self, "autoscaling")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="containerRuntime")
    def container_runtime(self) -> str:
        """
        The container runtime of the pool.
        """
        return pulumi.get(self, "container_runtime")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The creation date of the pool.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentSize")
    def current_size(self) -> int:
        return pulumi.get(self, "current_size")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubeletArgs")
    def kubelet_args(self) -> Mapping[str, str]:
        return pulumi.get(self, "kubelet_args")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> int:
        """
        The maximum size of the pool, used by the autoscaling feature.
        """
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> int:
        """
        The minimum size of the pool, used by the autoscaling feature.
        """
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the node.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> str:
        """
        The commercial type of the pool instances.
        """
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.GetKubernetesNodePoolNodeResult']:
        """
        (List of) The nodes in the default pool.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> str:
        """
        [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool are attached to.
        """
        return pulumi.get(self, "placement_group_id")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[str]:
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter(name="publicIpDisabled")
    def public_ip_disabled(self) -> bool:
        return pulumi.get(self, "public_ip_disabled")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="rootVolumeSizeInGb")
    def root_volume_size_in_gb(self) -> int:
        return pulumi.get(self, "root_volume_size_in_gb")

    @property
    @pulumi.getter(name="rootVolumeType")
    def root_volume_type(self) -> str:
        return pulumi.get(self, "root_volume_type")

    @property
    @pulumi.getter
    def size(self) -> Optional[int]:
        """
        The size of the pool.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the node.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        The tags associated with the pool.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The last update date of the pool.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="upgradePolicies")
    def upgrade_policies(self) -> Sequence['outputs.GetKubernetesNodePoolUpgradePolicyResult']:
        return pulumi.get(self, "upgrade_policies")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of the pool.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="waitForPoolReady")
    def wait_for_pool_ready(self) -> bool:
        return pulumi.get(self, "wait_for_pool_ready")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetKubernetesNodePoolResult(GetKubernetesNodePoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesNodePoolResult(
            autohealing=self.autohealing,
            autoscaling=self.autoscaling,
            cluster_id=self.cluster_id,
            container_runtime=self.container_runtime,
            created_at=self.created_at,
            current_size=self.current_size,
            id=self.id,
            kubelet_args=self.kubelet_args,
            max_size=self.max_size,
            min_size=self.min_size,
            name=self.name,
            node_type=self.node_type,
            nodes=self.nodes,
            placement_group_id=self.placement_group_id,
            pool_id=self.pool_id,
            public_ip_disabled=self.public_ip_disabled,
            region=self.region,
            root_volume_size_in_gb=self.root_volume_size_in_gb,
            root_volume_type=self.root_volume_type,
            size=self.size,
            status=self.status,
            tags=self.tags,
            updated_at=self.updated_at,
            upgrade_policies=self.upgrade_policies,
            version=self.version,
            wait_for_pool_ready=self.wait_for_pool_ready,
            zone=self.zone)


def get_kubernetes_node_pool(cluster_id: Optional[str] = None,
                             name: Optional[str] = None,
                             pool_id: Optional[str] = None,
                             region: Optional[str] = None,
                             size: Optional[int] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubernetesNodePoolResult:
    """
    Gets information about a Kubernetes Cluster's Pool.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_kubernetes_node_pool(pool_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: The cluster ID. Required when `name` is set.
    :param str name: The pool name. Only one of `name` and `pool_id` should be specified. `cluster_id` should be specified with `name`.
    :param str pool_id: The pool's ID. Only one of `name` and `pool_id` should be specified.
    :param str region: `region`) The region in which the pool exists.
    :param int size: The size of the pool.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    __args__['poolId'] = pool_id
    __args__['region'] = region
    __args__['size'] = size
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getKubernetesNodePool:getKubernetesNodePool', __args__, opts=opts, typ=GetKubernetesNodePoolResult).value

    return AwaitableGetKubernetesNodePoolResult(
        autohealing=pulumi.get(__ret__, 'autohealing'),
        autoscaling=pulumi.get(__ret__, 'autoscaling'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        container_runtime=pulumi.get(__ret__, 'container_runtime'),
        created_at=pulumi.get(__ret__, 'created_at'),
        current_size=pulumi.get(__ret__, 'current_size'),
        id=pulumi.get(__ret__, 'id'),
        kubelet_args=pulumi.get(__ret__, 'kubelet_args'),
        max_size=pulumi.get(__ret__, 'max_size'),
        min_size=pulumi.get(__ret__, 'min_size'),
        name=pulumi.get(__ret__, 'name'),
        node_type=pulumi.get(__ret__, 'node_type'),
        nodes=pulumi.get(__ret__, 'nodes'),
        placement_group_id=pulumi.get(__ret__, 'placement_group_id'),
        pool_id=pulumi.get(__ret__, 'pool_id'),
        public_ip_disabled=pulumi.get(__ret__, 'public_ip_disabled'),
        region=pulumi.get(__ret__, 'region'),
        root_volume_size_in_gb=pulumi.get(__ret__, 'root_volume_size_in_gb'),
        root_volume_type=pulumi.get(__ret__, 'root_volume_type'),
        size=pulumi.get(__ret__, 'size'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        upgrade_policies=pulumi.get(__ret__, 'upgrade_policies'),
        version=pulumi.get(__ret__, 'version'),
        wait_for_pool_ready=pulumi.get(__ret__, 'wait_for_pool_ready'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_kubernetes_node_pool)
def get_kubernetes_node_pool_output(cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    name: Optional[pulumi.Input[Optional[str]]] = None,
                                    pool_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    region: Optional[pulumi.Input[Optional[str]]] = None,
                                    size: Optional[pulumi.Input[Optional[int]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubernetesNodePoolResult]:
    """
    Gets information about a Kubernetes Cluster's Pool.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_kubernetes_node_pool(pool_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: The cluster ID. Required when `name` is set.
    :param str name: The pool name. Only one of `name` and `pool_id` should be specified. `cluster_id` should be specified with `name`.
    :param str pool_id: The pool's ID. Only one of `name` and `pool_id` should be specified.
    :param str region: `region`) The region in which the pool exists.
    :param int size: The size of the pool.
    """
    ...
