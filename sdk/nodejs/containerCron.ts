// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Container Triggers. For the moment, the feature is limited to CRON Schedule (time-based).
 *
 * For more information consult
 * the [documentation](https://www.scaleway.com/en/docs/serverless/containers/)
 * .
 *
 * For more details about the limitation
 * check [containers-limitations](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/)
 * .
 *
 * You can check also
 * our [containers cron api documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#crons-942bf4).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const mainContainerNamespace = new scaleway.ContainerNamespace("mainContainerNamespace", {});
 * const mainContainer = new scaleway.Container("mainContainer", {namespaceId: mainContainerNamespace.id});
 * const mainContainerCron = new scaleway.ContainerCron("mainContainerCron", {
 *     containerId: mainContainer.id,
 *     schedule: "5 4 1 * *",
 *     args: JSON.stringify({
 *         address: {
 *             city: "Paris",
 *             country: "FR",
 *         },
 *         age: 23,
 *         firstName: "John",
 *         isAlive: true,
 *         lastName: "Smith",
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * Container Cron can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/containerCron:ContainerCron main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ContainerCron extends pulumi.CustomResource {
    /**
     * Get an existing ContainerCron resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerCronState, opts?: pulumi.CustomResourceOptions): ContainerCron {
        return new ContainerCron(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/containerCron:ContainerCron';

    /**
     * Returns true if the given object is an instance of ContainerCron.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerCron {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerCron.__pulumiType;
    }

    /**
     * The key-value mapping to define arguments that will be passed to your container’s event object
     * during
     */
    public readonly args!: pulumi.Output<string>;
    /**
     * The container ID to link with your cron.
     */
    public readonly containerId!: pulumi.Output<string>;
    /**
     * The name of the container cron. If not provided, the name is generated.
     * during
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * The cron status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ContainerCron resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerCronArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerCronArgs | ContainerCronState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerCronState | undefined;
            resourceInputs["args"] = state ? state.args : undefined;
            resourceInputs["containerId"] = state ? state.containerId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ContainerCronArgs | undefined;
            if ((!args || args.args === undefined) && !opts.urn) {
                throw new Error("Missing required property 'args'");
            }
            if ((!args || args.containerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerId'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            resourceInputs["args"] = args ? args.args : undefined;
            resourceInputs["containerId"] = args ? args.containerId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerCron.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerCron resources.
 */
export interface ContainerCronState {
    /**
     * The key-value mapping to define arguments that will be passed to your container’s event object
     * during
     */
    args?: pulumi.Input<string>;
    /**
     * The container ID to link with your cron.
     */
    containerId?: pulumi.Input<string>;
    /**
     * The name of the container cron. If not provided, the name is generated.
     * during
     */
    name?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     */
    region?: pulumi.Input<string>;
    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     */
    schedule?: pulumi.Input<string>;
    /**
     * The cron status.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerCron resource.
 */
export interface ContainerCronArgs {
    /**
     * The key-value mapping to define arguments that will be passed to your container’s event object
     * during
     */
    args: pulumi.Input<string>;
    /**
     * The container ID to link with your cron.
     */
    containerId: pulumi.Input<string>;
    /**
     * The name of the container cron. If not provided, the name is generated.
     * during
     */
    name?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     */
    region?: pulumi.Input<string>;
    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     */
    schedule: pulumi.Input<string>;
}
