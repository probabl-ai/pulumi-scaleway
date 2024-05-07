// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    public static class GetIpamIp
    {
        /// <summary>
        /// Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
        /// 
        /// ## Examples
        /// </summary>
        public static Task<GetIpamIpResult> InvokeAsync(GetIpamIpArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpamIpResult>("scaleway:index/getIpamIp:getIpamIp", args ?? new GetIpamIpArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
        /// 
        /// ## Examples
        /// </summary>
        public static Output<GetIpamIpResult> Invoke(GetIpamIpInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpamIpResult>("scaleway:index/getIpamIp:getIpamIp", args ?? new GetIpamIpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpamIpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("attached")]
        public bool? Attached { get; set; }

        /// <summary>
        /// The IPAM IP ID. Cannot be used with the rest of the arguments.
        /// </summary>
        [Input("ipamIpId")]
        public string? IpamIpId { get; set; }

        /// <summary>
        /// The Mac Address linked to the IP. Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("macAddress")]
        public string? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network the IP belong to. Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the IP exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Filter by resource ID, type or name. Cannot be used with `ipam_ip_id`.
        /// If specified, `type` is required, and at least one of `id` or `name` must be set.
        /// </summary>
        [Input("resource")]
        public Inputs.GetIpamIpResourceArgs? Resource { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The tags associated with the IP. Cannot be used with `ipam_ip_id`.
        /// As datasource only returns one IP, the search with given tags must return only one result.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of IP to search for (ipv4, ipv6). Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public string? Zonal { get; set; }

        public GetIpamIpArgs()
        {
        }
        public static new GetIpamIpArgs Empty => new GetIpamIpArgs();
    }

    public sealed class GetIpamIpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("attached")]
        public Input<bool>? Attached { get; set; }

        /// <summary>
        /// The IPAM IP ID. Cannot be used with the rest of the arguments.
        /// </summary>
        [Input("ipamIpId")]
        public Input<string>? IpamIpId { get; set; }

        /// <summary>
        /// The Mac Address linked to the IP. Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network the IP belong to. Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the IP exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Filter by resource ID, type or name. Cannot be used with `ipam_ip_id`.
        /// If specified, `type` is required, and at least one of `id` or `name` must be set.
        /// </summary>
        [Input("resource")]
        public Input<Inputs.GetIpamIpResourceInputArgs>? Resource { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the IP. Cannot be used with `ipam_ip_id`.
        /// As datasource only returns one IP, the search with given tags must return only one result.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of IP to search for (ipv4, ipv6). Cannot be used with `ipam_ip_id`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Only IPs that are zonal, and in this zone, will be returned.
        /// </summary>
        [Input("zonal")]
        public Input<string>? Zonal { get; set; }

        public GetIpamIpInvokeArgs()
        {
        }
        public static new GetIpamIpInvokeArgs Empty => new GetIpamIpInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpamIpResult
    {
        /// <summary>
        /// The IP address.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// the IP address with a CIDR notation.
        /// </summary>
        public readonly string AddressCidr;
        public readonly bool? Attached;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IpamIpId;
        public readonly string? MacAddress;
        public readonly string OrganizationId;
        public readonly string? PrivateNetworkId;
        public readonly string ProjectId;
        public readonly string Region;
        public readonly Outputs.GetIpamIpResourceResult? Resource;
        public readonly ImmutableArray<string> Tags;
        public readonly string? Type;
        public readonly string Zonal;

        [OutputConstructor]
        private GetIpamIpResult(
            string address,

            string addressCidr,

            bool? attached,

            string id,

            string? ipamIpId,

            string? macAddress,

            string organizationId,

            string? privateNetworkId,

            string projectId,

            string region,

            Outputs.GetIpamIpResourceResult? resource,

            ImmutableArray<string> tags,

            string? type,

            string zonal)
        {
            Address = address;
            AddressCidr = addressCidr;
            Attached = attached;
            Id = id;
            IpamIpId = ipamIpId;
            MacAddress = macAddress;
            OrganizationId = organizationId;
            PrivateNetworkId = privateNetworkId;
            ProjectId = projectId;
            Region = region;
            Resource = resource;
            Tags = tags;
            Type = type;
            Zonal = zonal;
        }
    }
}
