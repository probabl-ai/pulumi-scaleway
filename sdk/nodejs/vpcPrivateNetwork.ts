// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway VPC Private Networks.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc/api/#private-networks-ac2df4).
 *
 * ## Example
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const pnPriv = new scaleway.VpcPrivateNetwork("pnPriv", {tags: [
 *     "demo",
 *     "terraform",
 * ]});
 * ```
 *
 * ### With subnets
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const pnPriv = new scaleway.VpcPrivateNetwork("pnPriv", {
 *     ipv4Subnet: {
 *         subnet: "192.168.0.0/24",
 *     },
 *     ipv6Subnets: [
 *         {
 *             subnet: "fd46:78ab:30b8:177c::/64",
 *         },
 *         {
 *             subnet: "fd46:78ab:30b8:c7df::/64",
 *         },
 *     ],
 *     tags: [
 *         "demo",
 *         "terraform",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Private networks can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork vpc_demo fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class VpcPrivateNetwork extends pulumi.CustomResource {
    /**
     * Get an existing VpcPrivateNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcPrivateNetworkState, opts?: pulumi.CustomResourceOptions): VpcPrivateNetwork {
        return new VpcPrivateNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork';

    /**
     * Returns true if the given object is an instance of VpcPrivateNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcPrivateNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcPrivateNetwork.__pulumiType;
    }

    /**
     * The date and time of the creation of the subnet.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The IPv4 subnet to associate with the private network.
     */
    public readonly ipv4Subnet!: pulumi.Output<outputs.VpcPrivateNetworkIpv4Subnet>;
    /**
     * The IPv6 subnets to associate with the private network.
     */
    public readonly ipv6Subnets!: pulumi.Output<outputs.VpcPrivateNetworkIpv6Subnet[]>;
    /**
     * The private networks are necessarily regional now.
     *
     * @deprecated This field is deprecated and will be removed in the next major version
     */
    public readonly isRegional!: pulumi.Output<boolean>;
    /**
     * The name of the private network. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the private network is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the private network is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region of the private network.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The tags associated with the private network.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The date and time of the last update of the subnet.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The VPC in which to create the private network.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
     *
     * @deprecated This field is deprecated and will be removed in the next major version, please use `region` instead
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a VpcPrivateNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpcPrivateNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcPrivateNetworkArgs | VpcPrivateNetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcPrivateNetworkState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["ipv4Subnet"] = state ? state.ipv4Subnet : undefined;
            resourceInputs["ipv6Subnets"] = state ? state.ipv6Subnets : undefined;
            resourceInputs["isRegional"] = state ? state.isRegional : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as VpcPrivateNetworkArgs | undefined;
            resourceInputs["ipv4Subnet"] = args ? args.ipv4Subnet : undefined;
            resourceInputs["ipv6Subnets"] = args ? args.ipv6Subnets : undefined;
            resourceInputs["isRegional"] = args ? args.isRegional : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcPrivateNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcPrivateNetwork resources.
 */
export interface VpcPrivateNetworkState {
    /**
     * The date and time of the creation of the subnet.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The IPv4 subnet to associate with the private network.
     */
    ipv4Subnet?: pulumi.Input<inputs.VpcPrivateNetworkIpv4Subnet>;
    /**
     * The IPv6 subnets to associate with the private network.
     */
    ipv6Subnets?: pulumi.Input<pulumi.Input<inputs.VpcPrivateNetworkIpv6Subnet>[]>;
    /**
     * The private networks are necessarily regional now.
     *
     * @deprecated This field is deprecated and will be removed in the next major version
     */
    isRegional?: pulumi.Input<boolean>;
    /**
     * The name of the private network. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the private network is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the private network is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the private network.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the private network.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time of the last update of the subnet.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The VPC in which to create the private network.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
     *
     * @deprecated This field is deprecated and will be removed in the next major version, please use `region` instead
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcPrivateNetwork resource.
 */
export interface VpcPrivateNetworkArgs {
    /**
     * The IPv4 subnet to associate with the private network.
     */
    ipv4Subnet?: pulumi.Input<inputs.VpcPrivateNetworkIpv4Subnet>;
    /**
     * The IPv6 subnets to associate with the private network.
     */
    ipv6Subnets?: pulumi.Input<pulumi.Input<inputs.VpcPrivateNetworkIpv6Subnet>[]>;
    /**
     * The private networks are necessarily regional now.
     *
     * @deprecated This field is deprecated and will be removed in the next major version
     */
    isRegional?: pulumi.Input<boolean>;
    /**
     * The name of the private network. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the private network is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region of the private network.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the private network.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The VPC in which to create the private network.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
     *
     * @deprecated This field is deprecated and will be removed in the next major version, please use `region` instead
     */
    zone?: pulumi.Input<string>;
}
