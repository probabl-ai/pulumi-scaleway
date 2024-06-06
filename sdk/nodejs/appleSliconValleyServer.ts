// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Apple silicon. For more information,
 * see [the documentation](https://www.scaleway.com/en/developers/api/apple-silicon/).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const server = new scaleway.AppleSliconValleyServer("server", {type: "M1-M"});
 * ```
 *
 * ## Import
 *
 * Instance servers can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/appleSliconValleyServer:AppleSliconValleyServer main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class AppleSliconValleyServer extends pulumi.CustomResource {
    /**
     * Get an existing AppleSliconValleyServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppleSliconValleyServerState, opts?: pulumi.CustomResourceOptions): AppleSliconValleyServer {
        return new AppleSliconValleyServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/appleSliconValleyServer:AppleSliconValleyServer';

    /**
     * Returns true if the given object is an instance of AppleSliconValleyServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppleSliconValleyServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppleSliconValleyServer.__pulumiType;
    }

    /**
     * The date and time of the creation of the Apple Silicon server.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The minimal date and time on which you can delete this server due to Apple licence
     */
    public /*out*/ readonly deletableAt!: pulumi.Output<string>;
    /**
     * IPv4 address of the server (IPv4 address).
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * The name of the server.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the server is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the server is
     * associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The state of the server.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The commercial type of the server. You find all the available types on
     * the [pricing page](https://www.scaleway.com/en/pricing/apple-silicon/). Updates to this field will recreate a new
     * resource.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the Apple Silicon server.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * URL of the VNC.
     */
    public /*out*/ readonly vncUrl!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which
     * the server should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a AppleSliconValleyServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppleSliconValleyServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppleSliconValleyServerArgs | AppleSliconValleyServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppleSliconValleyServerState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deletableAt"] = state ? state.deletableAt : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["vncUrl"] = state ? state.vncUrl : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as AppleSliconValleyServerArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["deletableAt"] = undefined /*out*/;
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["vncUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppleSliconValleyServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppleSliconValleyServer resources.
 */
export interface AppleSliconValleyServerState {
    /**
     * The date and time of the creation of the Apple Silicon server.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The minimal date and time on which you can delete this server due to Apple licence
     */
    deletableAt?: pulumi.Input<string>;
    /**
     * IPv4 address of the server (IPv4 address).
     */
    ip?: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the server is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the server is
     * associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The state of the server.
     */
    state?: pulumi.Input<string>;
    /**
     * The commercial type of the server. You find all the available types on
     * the [pricing page](https://www.scaleway.com/en/pricing/apple-silicon/). Updates to this field will recreate a new
     * resource.
     */
    type?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the Apple Silicon server.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * URL of the VNC.
     */
    vncUrl?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which
     * the server should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppleSliconValleyServer resource.
 */
export interface AppleSliconValleyServerArgs {
    /**
     * The name of the server.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the server is
     * associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The commercial type of the server. You find all the available types on
     * the [pricing page](https://www.scaleway.com/en/pricing/apple-silicon/). Updates to this field will recreate a new
     * resource.
     */
    type: pulumi.Input<string>;
    /**
     * `zone`) The zone in which
     * the server should be created.
     */
    zone?: pulumi.Input<string>;
}
