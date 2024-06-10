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
 * const main = new scaleway.IotHub("main", {productPlan: "plan_shared"});
 * ```
 *
 * ## Import
 *
 * IoT Hubs can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/iotHub:IotHub hub01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class IotHub extends pulumi.CustomResource {
    /**
     * Get an existing IotHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IotHubState, opts?: pulumi.CustomResourceOptions): IotHub {
        return new IotHub(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/iotHub:IotHub';

    /**
     * Returns true if the given object is an instance of IotHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotHub.__pulumiType;
    }

    /**
     * The current number of connected devices in the Hub.
     */
    public /*out*/ readonly connectedDeviceCount!: pulumi.Output<number>;
    /**
     * The date and time the Hub was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Wether to enable the device auto provisioning or not
     */
    public readonly deviceAutoProvisioning!: pulumi.Output<boolean | undefined>;
    /**
     * The number of registered devices in the Hub.
     */
    public /*out*/ readonly deviceCount!: pulumi.Output<number>;
    /**
     * Whether to enable the hub events or not
     */
    public readonly disableEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Wether the IoT Hub instance should be enabled or not.
     *
     * > **Important:** Updates to `enabled` will disconnect eventually connected devices.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The MQTT network endpoint to connect MQTT devices to.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Topic prefix for the hub events
     */
    public readonly eventsTopicPrefix!: pulumi.Output<string | undefined>;
    /**
     * Custom user provided certificate authority
     */
    public readonly hubCa!: pulumi.Output<string | undefined>;
    /**
     * Challenge certificate for the user provided hub CA
     */
    public readonly hubCaChallenge!: pulumi.Output<string | undefined>;
    /**
     * The MQTT certificat content
     */
    public /*out*/ readonly mqttCa!: pulumi.Output<string>;
    /**
     * The MQTT ca url
     */
    public /*out*/ readonly mqttCaUrl!: pulumi.Output<string>;
    /**
     * The name of the IoT Hub instance you want to create (e.g. `my-hub`).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organizationId you want to attach the resource to
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
     *
     * > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
     */
    public readonly productPlan!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the IoT Hub Instance is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The current status of the Hub.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The date and time the Hub resource was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a IotHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IotHubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IotHubArgs | IotHubState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IotHubState | undefined;
            resourceInputs["connectedDeviceCount"] = state ? state.connectedDeviceCount : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deviceAutoProvisioning"] = state ? state.deviceAutoProvisioning : undefined;
            resourceInputs["deviceCount"] = state ? state.deviceCount : undefined;
            resourceInputs["disableEvents"] = state ? state.disableEvents : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["eventsTopicPrefix"] = state ? state.eventsTopicPrefix : undefined;
            resourceInputs["hubCa"] = state ? state.hubCa : undefined;
            resourceInputs["hubCaChallenge"] = state ? state.hubCaChallenge : undefined;
            resourceInputs["mqttCa"] = state ? state.mqttCa : undefined;
            resourceInputs["mqttCaUrl"] = state ? state.mqttCaUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["productPlan"] = state ? state.productPlan : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IotHubArgs | undefined;
            if ((!args || args.productPlan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productPlan'");
            }
            resourceInputs["deviceAutoProvisioning"] = args ? args.deviceAutoProvisioning : undefined;
            resourceInputs["disableEvents"] = args ? args.disableEvents : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventsTopicPrefix"] = args ? args.eventsTopicPrefix : undefined;
            resourceInputs["hubCa"] = args ? args.hubCa : undefined;
            resourceInputs["hubCaChallenge"] = args ? args.hubCaChallenge : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["productPlan"] = args ? args.productPlan : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["connectedDeviceCount"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["deviceCount"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["mqttCa"] = undefined /*out*/;
            resourceInputs["mqttCaUrl"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IotHub.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IotHub resources.
 */
export interface IotHubState {
    /**
     * The current number of connected devices in the Hub.
     */
    connectedDeviceCount?: pulumi.Input<number>;
    /**
     * The date and time the Hub was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Wether to enable the device auto provisioning or not
     */
    deviceAutoProvisioning?: pulumi.Input<boolean>;
    /**
     * The number of registered devices in the Hub.
     */
    deviceCount?: pulumi.Input<number>;
    /**
     * Whether to enable the hub events or not
     */
    disableEvents?: pulumi.Input<boolean>;
    /**
     * Wether the IoT Hub instance should be enabled or not.
     *
     * > **Important:** Updates to `enabled` will disconnect eventually connected devices.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The MQTT network endpoint to connect MQTT devices to.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Topic prefix for the hub events
     */
    eventsTopicPrefix?: pulumi.Input<string>;
    /**
     * Custom user provided certificate authority
     */
    hubCa?: pulumi.Input<string>;
    /**
     * Challenge certificate for the user provided hub CA
     */
    hubCaChallenge?: pulumi.Input<string>;
    /**
     * The MQTT certificat content
     */
    mqttCa?: pulumi.Input<string>;
    /**
     * The MQTT ca url
     */
    mqttCaUrl?: pulumi.Input<string>;
    /**
     * The name of the IoT Hub instance you want to create (e.g. `my-hub`).
     */
    name?: pulumi.Input<string>;
    /**
     * The organizationId you want to attach the resource to
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
     *
     * > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
     */
    productPlan?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the IoT Hub Instance is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The current status of the Hub.
     */
    status?: pulumi.Input<string>;
    /**
     * The date and time the Hub resource was updated.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IotHub resource.
 */
export interface IotHubArgs {
    /**
     * Wether to enable the device auto provisioning or not
     */
    deviceAutoProvisioning?: pulumi.Input<boolean>;
    /**
     * Whether to enable the hub events or not
     */
    disableEvents?: pulumi.Input<boolean>;
    /**
     * Wether the IoT Hub instance should be enabled or not.
     *
     * > **Important:** Updates to `enabled` will disconnect eventually connected devices.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Topic prefix for the hub events
     */
    eventsTopicPrefix?: pulumi.Input<string>;
    /**
     * Custom user provided certificate authority
     */
    hubCa?: pulumi.Input<string>;
    /**
     * Challenge certificate for the user provided hub CA
     */
    hubCaChallenge?: pulumi.Input<string>;
    /**
     * The name of the IoT Hub instance you want to create (e.g. `my-hub`).
     */
    name?: pulumi.Input<string>;
    /**
     * Product plan to create the hub, see documentation for available product plans (e.g. `planShared`)
     *
     * > **Important:** Updates to `productPlan` will recreate the IoT Hub Instance.
     */
    productPlan: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the IoT Hub Instance is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}
