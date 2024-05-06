// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Serverless SQL Databases. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/serverless-databases/).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const database = new scaleway.SdbDatabase("database", {
 *     maxCpu: 8,
 *     minCpu: 0,
 * });
 * ```
 *
 * ## Import
 *
 * Serverless SQL Database can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/sdbDatabase:SdbDatabase database fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class SdbDatabase extends pulumi.CustomResource {
    /**
     * Get an existing SdbDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SdbDatabaseState, opts?: pulumi.CustomResourceOptions): SdbDatabase {
        return new SdbDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/sdbDatabase:SdbDatabase';

    /**
     * Returns true if the given object is an instance of SdbDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SdbDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SdbDatabase.__pulumiType;
    }

    /**
     * Endpoint of the database
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The maximum number of CPU units for your database. Defaults to 15.
     */
    public readonly maxCpu!: pulumi.Output<number | undefined>;
    /**
     * The minimum number of CPU units for your database. Defaults to 0.
     */
    public readonly minCpu!: pulumi.Output<number | undefined>;
    /**
     * Name of the database (e.g. `my-new-database`).
     *
     * > **Important:** Updates to `name` will recreate the database.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project_id you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a SdbDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SdbDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SdbDatabaseArgs | SdbDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SdbDatabaseState | undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["maxCpu"] = state ? state.maxCpu : undefined;
            resourceInputs["minCpu"] = state ? state.minCpu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as SdbDatabaseArgs | undefined;
            resourceInputs["maxCpu"] = args ? args.maxCpu : undefined;
            resourceInputs["minCpu"] = args ? args.minCpu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["endpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SdbDatabase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SdbDatabase resources.
 */
export interface SdbDatabaseState {
    /**
     * Endpoint of the database
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The maximum number of CPU units for your database. Defaults to 15.
     */
    maxCpu?: pulumi.Input<number>;
    /**
     * The minimum number of CPU units for your database. Defaults to 0.
     */
    minCpu?: pulumi.Input<number>;
    /**
     * Name of the database (e.g. `my-new-database`).
     *
     * > **Important:** Updates to `name` will recreate the database.
     */
    name?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SdbDatabase resource.
 */
export interface SdbDatabaseArgs {
    /**
     * The maximum number of CPU units for your database. Defaults to 15.
     */
    maxCpu?: pulumi.Input<number>;
    /**
     * The minimum number of CPU units for your database. Defaults to 0.
     */
    minCpu?: pulumi.Input<number>;
    /**
     * Name of the database (e.g. `my-new-database`).
     *
     * > **Important:** Updates to `name` will recreate the database.
     */
    name?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
}
