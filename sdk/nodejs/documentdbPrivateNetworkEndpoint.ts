// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Database Private Network Endpoint.
 *
 * ## Example Usage
 *
 * ### Example Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const pn = new scaleway.VpcPrivateNetwork("pn", {});
 * const instance = new scaleway.DocumentdbInstance("instance", {
 *     nodeType: "docdb-play2-pico",
 *     engine: "FerretDB-1",
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     volumeSizeInGb: 20,
 * });
 * const main = new scaleway.DocumentdbPrivateNetworkEndpoint("main", {
 *     instanceId: instance.id,
 *     privateNetwork: {
 *         ipNet: "172.16.32.3/22",
 *         id: pn.id,
 *     },
 * }, {
 *     dependsOn: [pn],
 * });
 * ```
 *
 * ## Import
 *
 * Database Instance Endpoint can be imported using the `{region}/{endpoint_id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint end fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class DocumentdbPrivateNetworkEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing DocumentdbPrivateNetworkEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentdbPrivateNetworkEndpointState, opts?: pulumi.CustomResourceOptions): DocumentdbPrivateNetworkEndpoint {
        return new DocumentdbPrivateNetworkEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint';

    /**
     * Returns true if the given object is an instance of DocumentdbPrivateNetworkEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DocumentdbPrivateNetworkEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DocumentdbPrivateNetworkEndpoint.__pulumiType;
    }

    /**
     * UUID of the documentdb instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The private network specs details. This is a list with maximum one element and supports the following attributes:
     */
    public readonly privateNetwork!: pulumi.Output<outputs.DocumentdbPrivateNetworkEndpointPrivateNetwork | undefined>;
    /**
     * The region of the endpoint.
     *
     *
     * > **NOTE:** Please calculate your host IP.
     * using cirhost. Otherwise, lets IPAM service
     * handle the host IP on the network.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a DocumentdbPrivateNetworkEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentdbPrivateNetworkEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentdbPrivateNetworkEndpointArgs | DocumentdbPrivateNetworkEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentdbPrivateNetworkEndpointState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["privateNetwork"] = state ? state.privateNetwork : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as DocumentdbPrivateNetworkEndpointArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["privateNetwork"] = args ? args.privateNetwork : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DocumentdbPrivateNetworkEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DocumentdbPrivateNetworkEndpoint resources.
 */
export interface DocumentdbPrivateNetworkEndpointState {
    /**
     * UUID of the documentdb instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The private network specs details. This is a list with maximum one element and supports the following attributes:
     */
    privateNetwork?: pulumi.Input<inputs.DocumentdbPrivateNetworkEndpointPrivateNetwork>;
    /**
     * The region of the endpoint.
     *
     *
     * > **NOTE:** Please calculate your host IP.
     * using cirhost. Otherwise, lets IPAM service
     * handle the host IP on the network.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DocumentdbPrivateNetworkEndpoint resource.
 */
export interface DocumentdbPrivateNetworkEndpointArgs {
    /**
     * UUID of the documentdb instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The private network specs details. This is a list with maximum one element and supports the following attributes:
     */
    privateNetwork?: pulumi.Input<inputs.DocumentdbPrivateNetworkEndpointPrivateNetwork>;
    /**
     * The region of the endpoint.
     *
     *
     * > **NOTE:** Please calculate your host IP.
     * using cirhost. Otherwise, lets IPAM service
     * handle the host IP on the network.
     */
    region?: pulumi.Input<string>;
}
