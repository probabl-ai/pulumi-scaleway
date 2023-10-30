// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.scaleway.inputs.KubernetesNodePoolNodeArgs;
import com.pulumi.scaleway.inputs.KubernetesNodePoolUpgradePolicyArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubernetesNodePoolState extends com.pulumi.resources.ResourceArgs {

    public static final KubernetesNodePoolState Empty = new KubernetesNodePoolState();

    /**
     * Enables the autohealing feature for this pool.
     * 
     */
    @Import(name="autohealing")
    private @Nullable Output<Boolean> autohealing;

    /**
     * @return Enables the autohealing feature for this pool.
     * 
     */
    public Optional<Output<Boolean>> autohealing() {
        return Optional.ofNullable(this.autohealing);
    }

    /**
     * Enables the autoscaling feature for this pool.
     * &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
     * 
     */
    @Import(name="autoscaling")
    private @Nullable Output<Boolean> autoscaling;

    /**
     * @return Enables the autoscaling feature for this pool.
     * &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
     * 
     */
    public Optional<Output<Boolean>> autoscaling() {
        return Optional.ofNullable(this.autoscaling);
    }

    /**
     * The ID of the Kubernetes cluster on which this pool will be created.
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return The ID of the Kubernetes cluster on which this pool will be created.
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * The container runtime of the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="containerRuntime")
    private @Nullable Output<String> containerRuntime;

    /**
     * @return The container runtime of the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<String>> containerRuntime() {
        return Optional.ofNullable(this.containerRuntime);
    }

    /**
     * The creation date of the pool.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The creation date of the pool.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The actual size of the pool
     * 
     */
    @Import(name="currentSize")
    private @Nullable Output<Integer> currentSize;

    /**
     * @return The actual size of the pool
     * 
     */
    public Optional<Output<Integer>> currentSize() {
        return Optional.ofNullable(this.currentSize);
    }

    /**
     * The Kubelet arguments to be used by this pool
     * 
     */
    @Import(name="kubeletArgs")
    private @Nullable Output<Map<String,String>> kubeletArgs;

    /**
     * @return The Kubelet arguments to be used by this pool
     * 
     */
    public Optional<Output<Map<String,String>>> kubeletArgs() {
        return Optional.ofNullable(this.kubeletArgs);
    }

    /**
     * The maximum size of the pool, used by the autoscaling feature.
     * 
     */
    @Import(name="maxSize")
    private @Nullable Output<Integer> maxSize;

    /**
     * @return The maximum size of the pool, used by the autoscaling feature.
     * 
     */
    public Optional<Output<Integer>> maxSize() {
        return Optional.ofNullable(this.maxSize);
    }

    /**
     * The minimum size of the pool, used by the autoscaling feature.
     * 
     */
    @Import(name="minSize")
    private @Nullable Output<Integer> minSize;

    /**
     * @return The minimum size of the pool, used by the autoscaling feature.
     * 
     */
    public Optional<Output<Integer>> minSize() {
        return Optional.ofNullable(this.minSize);
    }

    /**
     * The name for the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name for the pool.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
     * 
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="nodeType")
    private @Nullable Output<String> nodeType;

    /**
     * @return The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
     * 
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<String>> nodeType() {
        return Optional.ofNullable(this.nodeType);
    }

    /**
     * (List of) The nodes in the default pool.
     * 
     */
    @Import(name="nodes")
    private @Nullable Output<List<KubernetesNodePoolNodeArgs>> nodes;

    /**
     * @return (List of) The nodes in the default pool.
     * 
     */
    public Optional<Output<List<KubernetesNodePoolNodeArgs>>> nodes() {
        return Optional.ofNullable(this.nodes);
    }

    /**
     * The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool will be attached to.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="placementGroupId")
    private @Nullable Output<String> placementGroupId;

    /**
     * @return The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool will be attached to.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<String>> placementGroupId() {
        return Optional.ofNullable(this.placementGroupId);
    }

    /**
     * Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="publicIpDisabled")
    private @Nullable Output<Boolean> publicIpDisabled;

    /**
     * @return Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<Boolean>> publicIpDisabled() {
        return Optional.ofNullable(this.publicIpDisabled);
    }

    /**
     * `region`) The region in which the pool should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the pool should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The size of the system volume of the nodes in gigabyte
     * 
     */
    @Import(name="rootVolumeSizeInGb")
    private @Nullable Output<Integer> rootVolumeSizeInGb;

    /**
     * @return The size of the system volume of the nodes in gigabyte
     * 
     */
    public Optional<Output<Integer>> rootVolumeSizeInGb() {
        return Optional.ofNullable(this.rootVolumeSizeInGb);
    }

    /**
     * System volume type of the nodes composing the pool
     * 
     */
    @Import(name="rootVolumeType")
    private @Nullable Output<String> rootVolumeType;

    /**
     * @return System volume type of the nodes composing the pool
     * 
     */
    public Optional<Output<String>> rootVolumeType() {
        return Optional.ofNullable(this.rootVolumeType);
    }

    /**
     * The size of the pool.
     * &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return The size of the pool.
     * &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * The status of the node.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the node.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags associated with the pool.
     * &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool&#39;s nodes are applied using tags. (Example: &#34;taint=taintName=taineValue:Effect&#34;)
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags associated with the pool.
     * &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool&#39;s nodes are applied using tags. (Example: &#34;taint=taintName=taineValue:Effect&#34;)
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The last update date of the pool.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return The last update date of the pool.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    /**
     * The Pool upgrade policy
     * 
     */
    @Import(name="upgradePolicy")
    private @Nullable Output<KubernetesNodePoolUpgradePolicyArgs> upgradePolicy;

    /**
     * @return The Pool upgrade policy
     * 
     */
    public Optional<Output<KubernetesNodePoolUpgradePolicyArgs>> upgradePolicy() {
        return Optional.ofNullable(this.upgradePolicy);
    }

    /**
     * The version of the pool.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The version of the pool.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * Whether to wait for the pool to be ready.
     * 
     */
    @Import(name="waitForPoolReady")
    private @Nullable Output<Boolean> waitForPoolReady;

    /**
     * @return Whether to wait for the pool to be ready.
     * 
     */
    public Optional<Output<Boolean>> waitForPoolReady() {
        return Optional.ofNullable(this.waitForPoolReady);
    }

    /**
     * `zone`) The zone in which the pool should be created.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone in which the pool should be created.
     * &gt; **Important:** Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private KubernetesNodePoolState() {}

    private KubernetesNodePoolState(KubernetesNodePoolState $) {
        this.autohealing = $.autohealing;
        this.autoscaling = $.autoscaling;
        this.clusterId = $.clusterId;
        this.containerRuntime = $.containerRuntime;
        this.createdAt = $.createdAt;
        this.currentSize = $.currentSize;
        this.kubeletArgs = $.kubeletArgs;
        this.maxSize = $.maxSize;
        this.minSize = $.minSize;
        this.name = $.name;
        this.nodeType = $.nodeType;
        this.nodes = $.nodes;
        this.placementGroupId = $.placementGroupId;
        this.publicIpDisabled = $.publicIpDisabled;
        this.region = $.region;
        this.rootVolumeSizeInGb = $.rootVolumeSizeInGb;
        this.rootVolumeType = $.rootVolumeType;
        this.size = $.size;
        this.status = $.status;
        this.tags = $.tags;
        this.updatedAt = $.updatedAt;
        this.upgradePolicy = $.upgradePolicy;
        this.version = $.version;
        this.waitForPoolReady = $.waitForPoolReady;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubernetesNodePoolState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubernetesNodePoolState $;

        public Builder() {
            $ = new KubernetesNodePoolState();
        }

        public Builder(KubernetesNodePoolState defaults) {
            $ = new KubernetesNodePoolState(Objects.requireNonNull(defaults));
        }

        /**
         * @param autohealing Enables the autohealing feature for this pool.
         * 
         * @return builder
         * 
         */
        public Builder autohealing(@Nullable Output<Boolean> autohealing) {
            $.autohealing = autohealing;
            return this;
        }

        /**
         * @param autohealing Enables the autohealing feature for this pool.
         * 
         * @return builder
         * 
         */
        public Builder autohealing(Boolean autohealing) {
            return autohealing(Output.of(autohealing));
        }

        /**
         * @param autoscaling Enables the autoscaling feature for this pool.
         * &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
         * 
         * @return builder
         * 
         */
        public Builder autoscaling(@Nullable Output<Boolean> autoscaling) {
            $.autoscaling = autoscaling;
            return this;
        }

        /**
         * @param autoscaling Enables the autoscaling feature for this pool.
         * &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
         * 
         * @return builder
         * 
         */
        public Builder autoscaling(Boolean autoscaling) {
            return autoscaling(Output.of(autoscaling));
        }

        /**
         * @param clusterId The ID of the Kubernetes cluster on which this pool will be created.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId The ID of the Kubernetes cluster on which this pool will be created.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param containerRuntime The container runtime of the pool.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder containerRuntime(@Nullable Output<String> containerRuntime) {
            $.containerRuntime = containerRuntime;
            return this;
        }

        /**
         * @param containerRuntime The container runtime of the pool.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder containerRuntime(String containerRuntime) {
            return containerRuntime(Output.of(containerRuntime));
        }

        /**
         * @param createdAt The creation date of the pool.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The creation date of the pool.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param currentSize The actual size of the pool
         * 
         * @return builder
         * 
         */
        public Builder currentSize(@Nullable Output<Integer> currentSize) {
            $.currentSize = currentSize;
            return this;
        }

        /**
         * @param currentSize The actual size of the pool
         * 
         * @return builder
         * 
         */
        public Builder currentSize(Integer currentSize) {
            return currentSize(Output.of(currentSize));
        }

        /**
         * @param kubeletArgs The Kubelet arguments to be used by this pool
         * 
         * @return builder
         * 
         */
        public Builder kubeletArgs(@Nullable Output<Map<String,String>> kubeletArgs) {
            $.kubeletArgs = kubeletArgs;
            return this;
        }

        /**
         * @param kubeletArgs The Kubelet arguments to be used by this pool
         * 
         * @return builder
         * 
         */
        public Builder kubeletArgs(Map<String,String> kubeletArgs) {
            return kubeletArgs(Output.of(kubeletArgs));
        }

        /**
         * @param maxSize The maximum size of the pool, used by the autoscaling feature.
         * 
         * @return builder
         * 
         */
        public Builder maxSize(@Nullable Output<Integer> maxSize) {
            $.maxSize = maxSize;
            return this;
        }

        /**
         * @param maxSize The maximum size of the pool, used by the autoscaling feature.
         * 
         * @return builder
         * 
         */
        public Builder maxSize(Integer maxSize) {
            return maxSize(Output.of(maxSize));
        }

        /**
         * @param minSize The minimum size of the pool, used by the autoscaling feature.
         * 
         * @return builder
         * 
         */
        public Builder minSize(@Nullable Output<Integer> minSize) {
            $.minSize = minSize;
            return this;
        }

        /**
         * @param minSize The minimum size of the pool, used by the autoscaling feature.
         * 
         * @return builder
         * 
         */
        public Builder minSize(Integer minSize) {
            return minSize(Output.of(minSize));
        }

        /**
         * @param name The name for the pool.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name for the pool.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nodeType The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
         * 
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder nodeType(@Nullable Output<String> nodeType) {
            $.nodeType = nodeType;
            return this;
        }

        /**
         * @param nodeType The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
         * 
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder nodeType(String nodeType) {
            return nodeType(Output.of(nodeType));
        }

        /**
         * @param nodes (List of) The nodes in the default pool.
         * 
         * @return builder
         * 
         */
        public Builder nodes(@Nullable Output<List<KubernetesNodePoolNodeArgs>> nodes) {
            $.nodes = nodes;
            return this;
        }

        /**
         * @param nodes (List of) The nodes in the default pool.
         * 
         * @return builder
         * 
         */
        public Builder nodes(List<KubernetesNodePoolNodeArgs> nodes) {
            return nodes(Output.of(nodes));
        }

        /**
         * @param nodes (List of) The nodes in the default pool.
         * 
         * @return builder
         * 
         */
        public Builder nodes(KubernetesNodePoolNodeArgs... nodes) {
            return nodes(List.of(nodes));
        }

        /**
         * @param placementGroupId The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool will be attached to.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder placementGroupId(@Nullable Output<String> placementGroupId) {
            $.placementGroupId = placementGroupId;
            return this;
        }

        /**
         * @param placementGroupId The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool will be attached to.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder placementGroupId(String placementGroupId) {
            return placementGroupId(Output.of(placementGroupId));
        }

        /**
         * @param publicIpDisabled Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder publicIpDisabled(@Nullable Output<Boolean> publicIpDisabled) {
            $.publicIpDisabled = publicIpDisabled;
            return this;
        }

        /**
         * @param publicIpDisabled Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder publicIpDisabled(Boolean publicIpDisabled) {
            return publicIpDisabled(Output.of(publicIpDisabled));
        }

        /**
         * @param region `region`) The region in which the pool should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the pool should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param rootVolumeSizeInGb The size of the system volume of the nodes in gigabyte
         * 
         * @return builder
         * 
         */
        public Builder rootVolumeSizeInGb(@Nullable Output<Integer> rootVolumeSizeInGb) {
            $.rootVolumeSizeInGb = rootVolumeSizeInGb;
            return this;
        }

        /**
         * @param rootVolumeSizeInGb The size of the system volume of the nodes in gigabyte
         * 
         * @return builder
         * 
         */
        public Builder rootVolumeSizeInGb(Integer rootVolumeSizeInGb) {
            return rootVolumeSizeInGb(Output.of(rootVolumeSizeInGb));
        }

        /**
         * @param rootVolumeType System volume type of the nodes composing the pool
         * 
         * @return builder
         * 
         */
        public Builder rootVolumeType(@Nullable Output<String> rootVolumeType) {
            $.rootVolumeType = rootVolumeType;
            return this;
        }

        /**
         * @param rootVolumeType System volume type of the nodes composing the pool
         * 
         * @return builder
         * 
         */
        public Builder rootVolumeType(String rootVolumeType) {
            return rootVolumeType(Output.of(rootVolumeType));
        }

        /**
         * @param size The size of the pool.
         * &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The size of the pool.
         * &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param status The status of the node.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the node.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tags associated with the pool.
         * &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool&#39;s nodes are applied using tags. (Example: &#34;taint=taintName=taineValue:Effect&#34;)
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags associated with the pool.
         * &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool&#39;s nodes are applied using tags. (Example: &#34;taint=taintName=taineValue:Effect&#34;)
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags associated with the pool.
         * &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool&#39;s nodes are applied using tags. (Example: &#34;taint=taintName=taineValue:Effect&#34;)
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param updatedAt The last update date of the pool.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt The last update date of the pool.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        /**
         * @param upgradePolicy The Pool upgrade policy
         * 
         * @return builder
         * 
         */
        public Builder upgradePolicy(@Nullable Output<KubernetesNodePoolUpgradePolicyArgs> upgradePolicy) {
            $.upgradePolicy = upgradePolicy;
            return this;
        }

        /**
         * @param upgradePolicy The Pool upgrade policy
         * 
         * @return builder
         * 
         */
        public Builder upgradePolicy(KubernetesNodePoolUpgradePolicyArgs upgradePolicy) {
            return upgradePolicy(Output.of(upgradePolicy));
        }

        /**
         * @param version The version of the pool.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of the pool.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param waitForPoolReady Whether to wait for the pool to be ready.
         * 
         * @return builder
         * 
         */
        public Builder waitForPoolReady(@Nullable Output<Boolean> waitForPoolReady) {
            $.waitForPoolReady = waitForPoolReady;
            return this;
        }

        /**
         * @param waitForPoolReady Whether to wait for the pool to be ready.
         * 
         * @return builder
         * 
         */
        public Builder waitForPoolReady(Boolean waitForPoolReady) {
            return waitForPoolReady(Output.of(waitForPoolReady));
        }

        /**
         * @param zone `zone`) The zone in which the pool should be created.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which the pool should be created.
         * &gt; **Important:** Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public KubernetesNodePoolState build() {
            return $;
        }
    }

}
