// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Database DocumentDB Users.
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as random from "@pulumi/random";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const dbPassword = new random.RandomPassword("dbPassword", {
 *     length: 16,
 *     special: true,
 * });
 * const dbAdmin = new scaleway.DocumentdbUser("dbAdmin", {
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 *     password: dbPassword.result,
 *     isAdmin: true,
 * });
 * ```
 *
 * ## Import
 *
 * Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/documentdbUser:DocumentdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
 * ```
 */
export class DocumentdbUser extends pulumi.CustomResource {
    /**
     * Get an existing DocumentdbUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentdbUserState, opts?: pulumi.CustomResourceOptions): DocumentdbUser {
        return new DocumentdbUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/documentdbUser:DocumentdbUser';

    /**
     * Returns true if the given object is an instance of DocumentdbUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DocumentdbUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DocumentdbUser.__pulumiType;
    }

    /**
     * UUID of the documentDB instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database User.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Grant admin permissions to the Database User.
     */
    public readonly isAdmin!: pulumi.Output<boolean | undefined>;
    /**
     * Database Username.
     *
     * > **Important:** Updates to `name` will recreate the Database User.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Database User password.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The Scaleway region this resource resides in.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a DocumentdbUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentdbUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentdbUserArgs | DocumentdbUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentdbUserState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isAdmin"] = state ? state.isAdmin : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as DocumentdbUserArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isAdmin"] = args ? args.isAdmin : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DocumentdbUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DocumentdbUser resources.
 */
export interface DocumentdbUserState {
    /**
     * UUID of the documentDB instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database User.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Grant admin permissions to the Database User.
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * Database Username.
     *
     * > **Important:** Updates to `name` will recreate the Database User.
     */
    name?: pulumi.Input<string>;
    /**
     * Database User password.
     */
    password?: pulumi.Input<string>;
    /**
     * The Scaleway region this resource resides in.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DocumentdbUser resource.
 */
export interface DocumentdbUserArgs {
    /**
     * UUID of the documentDB instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database User.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Grant admin permissions to the Database User.
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * Database Username.
     *
     * > **Important:** Updates to `name` will recreate the Database User.
     */
    name?: pulumi.Input<string>;
    /**
     * Database User password.
     */
    password: pulumi.Input<string>;
    /**
     * The Scaleway region this resource resides in.
     */
    region?: pulumi.Input<string>;
}
