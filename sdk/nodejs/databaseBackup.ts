// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway RDB database backup.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
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
 * const main = new scaleway.DatabaseBackup("main", {
 *     instanceId: data.scaleway_rdb_instance.main.id,
 *     databaseName: data.scaleway_rdb_database.main.name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### With expiration
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.DatabaseBackup("main", {
 *     instanceId: data.scaleway_rdb_instance.main.id,
 *     databaseName: data.scaleway_rdb_database.main.name,
 *     expiresAt: "2022-06-16T07:48:44Z",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * RDB Database can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/databaseBackup:DatabaseBackup mybackup fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class DatabaseBackup extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseBackup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseBackupState, opts?: pulumi.CustomResourceOptions): DatabaseBackup {
        return new DatabaseBackup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/databaseBackup:DatabaseBackup';

    /**
     * Returns true if the given object is an instance of DatabaseBackup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseBackup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseBackup.__pulumiType;
    }

    /**
     * Creation date (Format ISO 8601).
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the database of this backup.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * Expiration date (Format ISO 8601).
     *
     * > **Important:** `expiresAt` cannot be removed after being set.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Backup.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Name of the instance of the backup.
     */
    public /*out*/ readonly instanceName!: pulumi.Output<string>;
    /**
     * Name of the database (e.g. `my-database`).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Size of the backup (in bytes).
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Updated date (Format ISO 8601).
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a DatabaseBackup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseBackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseBackupArgs | DatabaseBackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseBackupState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as DatabaseBackupArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["instanceName"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseBackup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseBackup resources.
 */
export interface DatabaseBackupState {
    /**
     * Creation date (Format ISO 8601).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the database of this backup.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * Expiration date (Format ISO 8601).
     *
     * > **Important:** `expiresAt` cannot be removed after being set.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Backup.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Name of the instance of the backup.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Name of the database (e.g. `my-database`).
     */
    name?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
    /**
     * Size of the backup (in bytes).
     */
    size?: pulumi.Input<number>;
    /**
     * Updated date (Format ISO 8601).
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseBackup resource.
 */
export interface DatabaseBackupArgs {
    /**
     * Name of the database of this backup.
     */
    databaseName: pulumi.Input<string>;
    /**
     * Expiration date (Format ISO 8601).
     *
     * > **Important:** `expiresAt` cannot be removed after being set.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Backup.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Name of the database (e.g. `my-database`).
     */
    name?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
}
