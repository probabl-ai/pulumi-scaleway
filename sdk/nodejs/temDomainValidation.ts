// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumiverse/scaleway";
 *
 * const main = new scaleway.TemDomain("main", {acceptTos: true});
 * const example = new scaleway.TemDomainValidation("example", {
 *     domainId: main.id,
 *     region: "fr-par",
 *     timeout: 300,
 * });
 * ```
 */
export class TemDomainValidation extends pulumi.CustomResource {
    /**
     * Get an existing TemDomainValidation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemDomainValidationState, opts?: pulumi.CustomResourceOptions): TemDomainValidation {
        return new TemDomainValidation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/temDomainValidation:TemDomainValidation';

    /**
     * Returns true if the given object is an instance of TemDomainValidation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TemDomainValidation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TemDomainValidation.__pulumiType;
    }

    /**
     * The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
     */
    public readonly domainId!: pulumi.Output<string>;
    /**
     * `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
     */
    public /*out*/ readonly validated!: pulumi.Output<boolean>;

    /**
     * Create a TemDomainValidation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemDomainValidationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemDomainValidationArgs | TemDomainValidationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemDomainValidationState | undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["validated"] = state ? state.validated : undefined;
        } else {
            const args = argsOrState as TemDomainValidationArgs | undefined;
            if ((!args || args.domainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainId'");
            }
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["validated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TemDomainValidation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TemDomainValidation resources.
 */
export interface TemDomainValidationState {
    /**
     * The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
     */
    domainId?: pulumi.Input<string>;
    /**
     * `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
     */
    region?: pulumi.Input<string>;
    /**
     * The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Indicates if the domain has been verified for email sending. This is computed after the creation or update of the domain validation resource.
     */
    validated?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a TemDomainValidation resource.
 */
export interface TemDomainValidationArgs {
    /**
     * The ID of the domain name used when sending emails. This ID must correspond to a domain already registered with Scaleway's Transactional Email service.
     */
    domainId: pulumi.Input<string>;
    /**
     * `region`). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
     */
    region?: pulumi.Input<string>;
    /**
     * The maximum wait time in seconds before returning an error if the domain validation does not complete. The default is 300 seconds.
     */
    timeout?: pulumi.Input<number>;
}
