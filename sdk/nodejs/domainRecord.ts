// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Domain record.\
 * For more information, see [the documentation](https://www.scaleway.com/en/docs/scaleway-dns/).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const www = new scaleway.DomainRecord("www", {
 *     data: "1.2.3.4",
 *     dnsZone: "domain.tld",
 *     ttl: 3600,
 *     type: "A",
 * });
 * const www2 = new scaleway.DomainRecord("www2", {
 *     data: "1.2.3.5",
 *     dnsZone: "domain.tld",
 *     ttl: 3600,
 *     type: "A",
 * });
 * const mx = new scaleway.DomainRecord("mx", {
 *     data: "mx.online.net.",
 *     dnsZone: "domain.tld",
 *     priority: 10,
 *     ttl: 3600,
 *     type: "MX",
 * });
 * const mx2 = new scaleway.DomainRecord("mx2", {
 *     data: "mx-cache.online.net.",
 *     dnsZone: "domain.tld",
 *     priority: 20,
 *     ttl: 3600,
 *     type: "MX",
 * });
 * ```
 *
 * ### With dynamic records
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const geoIp = new scaleway.DomainRecord("geo_ip", {
 *     data: "1.2.3.4",
 *     dnsZone: "domain.tld",
 *     geoIp: {
 *         matches: [
 *             {
 *                 continents: ["EU"],
 *                 countries: ["FR"],
 *                 data: "1.2.3.5",
 *             },
 *             {
 *                 continents: ["NA"],
 *                 data: "4.3.2.1",
 *             },
 *         ],
 *     },
 *     ttl: 3600,
 *     type: "A",
 * });
 * const httpService = new scaleway.DomainRecord("http_service", {
 *     data: "1.2.3.4",
 *     dnsZone: "domain.tld",
 *     httpService: {
 *         ips: [
 *             "1.2.3.5",
 *             "1.2.3.6",
 *         ],
 *         mustContain: "up",
 *         strategy: "hashed",
 *         url: "http://mywebsite.com/health",
 *         userAgent: "scw_service_up",
 *     },
 *     ttl: 3600,
 *     type: "A",
 * });
 * const view = new scaleway.DomainRecord("view", {
 *     data: "1.2.3.4",
 *     dnsZone: "domain.tld",
 *     ttl: 3600,
 *     type: "A",
 *     views: [
 *         {
 *             data: "1.2.3.5",
 *             subnet: "100.0.0.0/16",
 *         },
 *         {
 *             data: "1.2.3.6",
 *             subnet: "100.1.0.0/16",
 *         },
 *     ],
 * });
 * const weighted = new scaleway.DomainRecord("weighted", {
 *     data: "1.2.3.4",
 *     dnsZone: "domain.tld",
 *     ttl: 3600,
 *     type: "A",
 *     weighteds: [
 *         {
 *             ip: "1.2.3.5",
 *             weight: 1,
 *         },
 *         {
 *             ip: "1.2.3.6",
 *             weight: 2,
 *         },
 *     ],
 * });
 * ```
 *
 * ### Create an instance and add records with the new instance IP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@lbrlabs/pulumi-scaleway";
 *
 * const config = new pulumi.Config();
 * const projectId = config.require("projectId");
 * const dnsZone = config.require("dnsZone");
 * const publicIp = new scaleway.InstanceIp("publicIp", {projectId: projectId});
 * const web = new scaleway.InstanceServer("web", {
 *     projectId: projectId,
 *     type: "DEV1-S",
 *     image: "ubuntu_focal",
 *     tags: [
 *         "front",
 *         "web",
 *     ],
 *     ipId: publicIp.id,
 *     rootVolume: {
 *         sizeInGb: 20,
 *     },
 * });
 * const webA = new scaleway.DomainRecord("webA", {
 *     dnsZone: dnsZone,
 *     type: "A",
 *     data: web.publicIp,
 *     ttl: 3600,
 * });
 * const webCname = new scaleway.DomainRecord("webCname", {
 *     dnsZone: dnsZone,
 *     type: "CNAME",
 *     data: `web.${dnsZone}.`,
 *     ttl: 3600,
 * });
 * const webAlias = new scaleway.DomainRecord("webAlias", {
 *     dnsZone: dnsZone,
 *     type: "ALIAS",
 *     data: `web.${dnsZone}.`,
 *     ttl: 3600,
 * });
 * ```
 *
 * ## Multiple records
 *
 * Some record types can have multiple `data` with the same `name` (eg: `A`, `AAAA`, `MX`, `NS`...).\
 * You can duplicate a resource `scaleway.DomainRecord` with the same `name`, the records will be added.
 *
 * Please note, some record (eg: `CNAME`, Multiple dynamic records of different types...) has to be unique.
 *
 * ## Import
 *
 * Record can be imported using the `{dns_zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/domainRecord:DomainRecord www subdomain.domain.tld/11111111-1111-1111-1111-111111111111
 * ```
 */
export class DomainRecord extends pulumi.CustomResource {
    /**
     * Get an existing DomainRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainRecordState, opts?: pulumi.CustomResourceOptions): DomainRecord {
        return new DomainRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/domainRecord:DomainRecord';

    /**
     * Returns true if the given object is an instance of DomainRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainRecord.__pulumiType;
    }

    /**
     * The data of the view record
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
     */
    public readonly dnsZone!: pulumi.Output<string>;
    /**
     * The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Geo-IP-Records)
     */
    public readonly geoIp!: pulumi.Output<outputs.DomainRecordGeoIp | undefined>;
    /**
     * The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Healthcheck-records)
     */
    public readonly httpService!: pulumi.Output<outputs.DomainRecordHttpService | undefined>;
    /**
     * When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
     */
    public readonly keepEmptyZone!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the record (can be an empty string for a root record).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The priority of the record (mostly used with an `MX` record)
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The project_id you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Does the DNS zone is the root zone or not
     */
    public /*out*/ readonly rootZone!: pulumi.Output<boolean>;
    /**
     * Time To Live of the record in seconds.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Views-records)
     */
    public readonly views!: pulumi.Output<outputs.DomainRecordView[] | undefined>;
    /**
     * You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Weight-Records)
     */
    public readonly weighteds!: pulumi.Output<outputs.DomainRecordWeighted[] | undefined>;

    /**
     * Create a DomainRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainRecordArgs | DomainRecordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainRecordState | undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["dnsZone"] = state ? state.dnsZone : undefined;
            resourceInputs["geoIp"] = state ? state.geoIp : undefined;
            resourceInputs["httpService"] = state ? state.httpService : undefined;
            resourceInputs["keepEmptyZone"] = state ? state.keepEmptyZone : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["rootZone"] = state ? state.rootZone : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["views"] = state ? state.views : undefined;
            resourceInputs["weighteds"] = state ? state.weighteds : undefined;
        } else {
            const args = argsOrState as DomainRecordArgs | undefined;
            if ((!args || args.data === undefined) && !opts.urn) {
                throw new Error("Missing required property 'data'");
            }
            if ((!args || args.dnsZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dnsZone'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["dnsZone"] = args ? args.dnsZone : undefined;
            resourceInputs["geoIp"] = args ? args.geoIp : undefined;
            resourceInputs["httpService"] = args ? args.httpService : undefined;
            resourceInputs["keepEmptyZone"] = args ? args.keepEmptyZone : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["views"] = args ? args.views : undefined;
            resourceInputs["weighteds"] = args ? args.weighteds : undefined;
            resourceInputs["rootZone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainRecord.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainRecord resources.
 */
export interface DomainRecordState {
    /**
     * The data of the view record
     */
    data?: pulumi.Input<string>;
    /**
     * The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
     */
    dnsZone?: pulumi.Input<string>;
    /**
     * The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Geo-IP-Records)
     */
    geoIp?: pulumi.Input<inputs.DomainRecordGeoIp>;
    /**
     * The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Healthcheck-records)
     */
    httpService?: pulumi.Input<inputs.DomainRecordHttpService>;
    /**
     * When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
     */
    keepEmptyZone?: pulumi.Input<boolean>;
    /**
     * The name of the record (can be an empty string for a root record).
     */
    name?: pulumi.Input<string>;
    /**
     * The priority of the record (mostly used with an `MX` record)
     */
    priority?: pulumi.Input<number>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * Does the DNS zone is the root zone or not
     */
    rootZone?: pulumi.Input<boolean>;
    /**
     * Time To Live of the record in seconds.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
     */
    type?: pulumi.Input<string>;
    /**
     * The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Views-records)
     */
    views?: pulumi.Input<pulumi.Input<inputs.DomainRecordView>[]>;
    /**
     * You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Weight-Records)
     */
    weighteds?: pulumi.Input<pulumi.Input<inputs.DomainRecordWeighted>[]>;
}

/**
 * The set of arguments for constructing a DomainRecord resource.
 */
export interface DomainRecordArgs {
    /**
     * The data of the view record
     */
    data: pulumi.Input<string>;
    /**
     * The DNS Zone of the domain. If the DNS zone doesn't exist, it will be automatically created.
     */
    dnsZone: pulumi.Input<string>;
    /**
     * The Geo IP feature provides DNS resolution, based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Geo-IP-Records)
     */
    geoIp?: pulumi.Input<inputs.DomainRecordGeoIp>;
    /**
     * The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs by excluding the ones not responding to the given string to check. [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Healthcheck-records)
     */
    httpService?: pulumi.Input<inputs.DomainRecordHttpService>;
    /**
     * When destroying a resource, if only NS records remain and this is set to `false`, the zone will be deleted. Please note, each zone not deleted will [cost you money](https://www.scaleway.com/en/dns/)
     */
    keepEmptyZone?: pulumi.Input<boolean>;
    /**
     * The name of the record (can be an empty string for a root record).
     */
    name?: pulumi.Input<string>;
    /**
     * The priority of the record (mostly used with an `MX` record)
     */
    priority?: pulumi.Input<number>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * Time To Live of the record in seconds.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
     */
    type: pulumi.Input<string>;
    /**
     * The answer to a DNS request is based on the client’s (resolver) subnet. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Views-records)
     */
    views?: pulumi.Input<pulumi.Input<inputs.DomainRecordView>[]>;
    /**
     * You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with its related IP compared to the others in the list. *(Can be more than 1)* [Documentation and usage example](https://www.scaleway.com/en/docs/scaleway-dns/#-Weight-Records)
     */
    weighteds?: pulumi.Input<pulumi.Input<inputs.DomainRecordWeighted>[]>;
}
