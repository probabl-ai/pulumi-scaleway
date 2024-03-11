// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Compute Instance User Data values.
 *
 * User data is a key value store API you can use to provide data from and to your server without authentication. It is the mechanism by which a user can pass information contained in a local file to an Instance at launch time.
 *
 * The typical use case is to pass something like a shell script or a configuration file as user data.
 *
 * For more information about [userData](https://developers.scaleway.com/en/products/instance/api/#patch-9ef3ec)  check our documentation guide [here](https://www.scaleway.com/en/docs/compute/instances/how-to/use-boot-modes/#how-to-use-cloud-init).
 *
 * About cloud-init documentation please check this [link](https://cloudinit.readthedocs.io/en/latest/).
 *
 * ## Examples
 *
 * ### Basic
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const config = new pulumi.Config();
 * const userData = config.getObject("userData") || {
 *     "cloud-init": `#cloud-config
 * apt-update: true
 * apt-upgrade: true
 * `,
 *     foo: "bar",
 * };
 * const mainInstanceServer = new scaleway.InstanceServer("mainInstanceServer", {
 *     image: "ubuntu_focal",
 *     type: "DEV1-S",
 * });
 * // User data with a single value
 * const mainInstanceUserData = new scaleway.InstanceUserData("mainInstanceUserData", {
 *     serverId: mainInstanceServer.id,
 *     key: "foo",
 *     value: "bar",
 * });
 * // User Data with many keys.
 * const data: scaleway.InstanceUserData[] = [];
 * for (const range = {value: 0}; range.value < userData; range.value++) {
 *     data.push(new scaleway.InstanceUserData(`data-${range.value}`, {
 *         serverId: mainInstanceServer.id,
 *         key: range.key,
 *         value: range.value,
 *     }));
 * }
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * User data can be imported using the `{zone}/{key}/{server_id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/instanceUserData:InstanceUserData main fr-par-1/cloud-init/11111111-1111-1111-1111-111111111111
 * ```
 */
export class InstanceUserData extends pulumi.CustomResource {
    /**
     * Get an existing InstanceUserData resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceUserDataState, opts?: pulumi.CustomResourceOptions): InstanceUserData {
        return new InstanceUserData(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/instanceUserData:InstanceUserData';

    /**
     * Returns true if the given object is an instance of InstanceUserData.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceUserData {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceUserData.__pulumiType;
    }

    /**
     * Key of the user data.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the server associated with.
     */
    public readonly serverId!: pulumi.Output<string>;
    /**
     * Value associated with your key
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the server should be created.
     *
     * > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
     * You can define values using:
     * - string
     * - UTF-8 encoded file content using file
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceUserData resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceUserDataArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceUserDataArgs | InstanceUserDataState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceUserDataState | undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceUserDataArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.serverId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverId'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["serverId"] = args ? args.serverId : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceUserData.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceUserData resources.
 */
export interface InstanceUserDataState {
    /**
     * Key of the user data.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the server associated with.
     */
    serverId?: pulumi.Input<string>;
    /**
     * Value associated with your key
     */
    value?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server should be created.
     *
     * > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
     * You can define values using:
     * - string
     * - UTF-8 encoded file content using file
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceUserData resource.
 */
export interface InstanceUserDataArgs {
    /**
     * Key of the user data.
     */
    key: pulumi.Input<string>;
    /**
     * The ID of the server associated with.
     */
    serverId: pulumi.Input<string>;
    /**
     * Value associated with your key
     */
    value: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server should be created.
     *
     * > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
     * You can define values using:
     * - string
     * - UTF-8 encoded file content using file
     */
    zone?: pulumi.Input<string>;
}
