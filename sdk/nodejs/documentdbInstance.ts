// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Database Instances.
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
 *
 * ## Examples
 *
 * ### Example Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.DocumentdbInstance("main", {
 *     engine: "FerretDB-1",
 *     nodeType: "docdb-play2-pico",
 *     password: "thiZ_is_v&ry_s3cret",
 *     tags: [
 *         "terraform-test",
 *         "scaleway_documentdb_instance",
 *         "minimal",
 *     ],
 *     userName: "my_initial_user",
 *     volumeSizeInGb: 20,
 * });
 * ```
 *
 * ## Import
 *
 * Database Instance can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/documentdbInstance:DocumentdbInstance db fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class DocumentdbInstance extends pulumi.CustomResource {
    /**
     * Get an existing DocumentdbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentdbInstanceState, opts?: pulumi.CustomResourceOptions): DocumentdbInstance {
        return new DocumentdbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/documentdbInstance:DocumentdbInstance';

    /**
     * Returns true if the given object is an instance of DocumentdbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DocumentdbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DocumentdbInstance.__pulumiType;
    }

    /**
     * Database Instance's engine version (e.g. `FerretDB-1`).
     *
     * > **Important:** Updates to `engine` will recreate the Database Instance.
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * Enable or disable high availability for the database instance.
     */
    public readonly isHaCluster!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Database Instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of database instance you want to create (e.g. `docdb-play2-pico`).
     *
     * > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
     * interruption. Keep in mind that you cannot downgrade a Database Instance.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * Password for the first user of the database instance.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * `projectId`) The ID of the project the Database
     * Instance is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region
     * in which the Database Instance should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The tags associated with the Database Instance.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
     *
     * > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
     */
    public readonly telemetryEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Identifier for the first user of the database instance.
     *
     * > **Important:** Updates to `userName` will recreate the Database Instance.
     */
    public readonly userName!: pulumi.Output<string | undefined>;
    /**
     * Volume size (in GB) when `volumeType` is set to `bssd`.
     */
    public readonly volumeSizeInGb!: pulumi.Output<number>;
    /**
     * Type of volume where data are stored (`bssd` or `lssd`).
     */
    public readonly volumeType!: pulumi.Output<string | undefined>;

    /**
     * Create a DocumentdbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentdbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentdbInstanceArgs | DocumentdbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentdbInstanceState | undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["isHaCluster"] = state ? state.isHaCluster : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["telemetryEnabled"] = state ? state.telemetryEnabled : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["volumeSizeInGb"] = state ? state.volumeSizeInGb : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as DocumentdbInstanceArgs | undefined;
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["isHaCluster"] = args ? args.isHaCluster : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["telemetryEnabled"] = args ? args.telemetryEnabled : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["volumeSizeInGb"] = args ? args.volumeSizeInGb : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DocumentdbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DocumentdbInstance resources.
 */
export interface DocumentdbInstanceState {
    /**
     * Database Instance's engine version (e.g. `FerretDB-1`).
     *
     * > **Important:** Updates to `engine` will recreate the Database Instance.
     */
    engine?: pulumi.Input<string>;
    /**
     * Enable or disable high availability for the database instance.
     */
    isHaCluster?: pulumi.Input<boolean>;
    /**
     * The name of the Database Instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of database instance you want to create (e.g. `docdb-play2-pico`).
     *
     * > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
     * interruption. Keep in mind that you cannot downgrade a Database Instance.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Password for the first user of the database instance.
     */
    password?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the Database
     * Instance is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region
     * in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the Database Instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
     *
     * > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
     */
    telemetryEnabled?: pulumi.Input<boolean>;
    /**
     * Identifier for the first user of the database instance.
     *
     * > **Important:** Updates to `userName` will recreate the Database Instance.
     */
    userName?: pulumi.Input<string>;
    /**
     * Volume size (in GB) when `volumeType` is set to `bssd`.
     */
    volumeSizeInGb?: pulumi.Input<number>;
    /**
     * Type of volume where data are stored (`bssd` or `lssd`).
     */
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DocumentdbInstance resource.
 */
export interface DocumentdbInstanceArgs {
    /**
     * Database Instance's engine version (e.g. `FerretDB-1`).
     *
     * > **Important:** Updates to `engine` will recreate the Database Instance.
     */
    engine: pulumi.Input<string>;
    /**
     * Enable or disable high availability for the database instance.
     */
    isHaCluster?: pulumi.Input<boolean>;
    /**
     * The name of the Database Instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of database instance you want to create (e.g. `docdb-play2-pico`).
     *
     * > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
     * interruption. Keep in mind that you cannot downgrade a Database Instance.
     */
    nodeType: pulumi.Input<string>;
    /**
     * Password for the first user of the database instance.
     */
    password?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the Database
     * Instance is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region
     * in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The tags associated with the Database Instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
     *
     * > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
     */
    telemetryEnabled?: pulumi.Input<boolean>;
    /**
     * Identifier for the first user of the database instance.
     *
     * > **Important:** Updates to `userName` will recreate the Database Instance.
     */
    userName?: pulumi.Input<string>;
    /**
     * Volume size (in GB) when `volumeType` is set to `bssd`.
     */
    volumeSizeInGb?: pulumi.Input<number>;
    /**
     * Type of volume where data are stored (`bssd` or `lssd`).
     */
    volumeType?: pulumi.Input<string>;
}
