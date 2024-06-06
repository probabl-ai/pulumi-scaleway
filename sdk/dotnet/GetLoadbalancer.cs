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
    public static class GetLoadbalancer
    {
        /// <summary>
        /// Gets information about a Load Balancer.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetLoadbalancer.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.GetLoadbalancer.Invoke(new()
        ///     {
        ///         LbId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLoadbalancerResult> InvokeAsync(GetLoadbalancerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoadbalancerResult>("scaleway:index/getLoadbalancer:getLoadbalancer", args ?? new GetLoadbalancerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Load Balancer.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var byName = Scaleway.GetLoadbalancer.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.GetLoadbalancer.Invoke(new()
        ///     {
        ///         LbId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLoadbalancerResult> Invoke(GetLoadbalancerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadbalancerResult>("scaleway:index/getLoadbalancer:getLoadbalancer", args ?? new GetLoadbalancerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadbalancerArgs : global::Pulumi.InvokeArgs
    {
        [Input("lbId")]
        public string? LbId { get; set; }

        /// <summary>
        /// The load balancer name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the LB is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("releaseIp")]
        public bool? ReleaseIp { get; set; }

        /// <summary>
        /// (Defaults to provider `zone`) The zone in which the LB exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetLoadbalancerArgs()
        {
        }
        public static new GetLoadbalancerArgs Empty => new GetLoadbalancerArgs();
    }

    public sealed class GetLoadbalancerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The load balancer name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the LB is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("releaseIp")]
        public Input<bool>? ReleaseIp { get; set; }

        /// <summary>
        /// (Defaults to provider `zone`) The zone in which the LB exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetLoadbalancerInvokeArgs()
        {
        }
        public static new GetLoadbalancerInvokeArgs Empty => new GetLoadbalancerInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoadbalancerResult
    {
        public readonly bool AssignFlexibleIp;
        public readonly bool AssignFlexibleIpv6;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The load-balancer public IP Address.
        /// </summary>
        public readonly string IpAddress;
        public readonly string IpId;
        public readonly ImmutableArray<string> IpIds;
        public readonly string Ipv6Address;
        public readonly string? LbId;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly ImmutableArray<Outputs.GetLoadbalancerPrivateNetworkResult> PrivateNetworks;
        public readonly string? ProjectId;
        public readonly string Region;
        public readonly bool? ReleaseIp;
        public readonly string SslCompatibilityLevel;
        /// <summary>
        /// The tags associated with the load-balancer.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of the load-balancer.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// (Defaults to provider `zone`) The zone in which the LB exists.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private GetLoadbalancerResult(
            bool assignFlexibleIp,

            bool assignFlexibleIpv6,

            string description,

            string id,

            string ipAddress,

            string ipId,

            ImmutableArray<string> ipIds,

            string ipv6Address,

            string? lbId,

            string? name,

            string organizationId,

            ImmutableArray<Outputs.GetLoadbalancerPrivateNetworkResult> privateNetworks,

            string? projectId,

            string region,

            bool? releaseIp,

            string sslCompatibilityLevel,

            ImmutableArray<string> tags,

            string type,

            string? zone)
        {
            AssignFlexibleIp = assignFlexibleIp;
            AssignFlexibleIpv6 = assignFlexibleIpv6;
            Description = description;
            Id = id;
            IpAddress = ipAddress;
            IpId = ipId;
            IpIds = ipIds;
            Ipv6Address = ipv6Address;
            LbId = lbId;
            Name = name;
            OrganizationId = organizationId;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            Region = region;
            ReleaseIp = releaseIp;
            SslCompatibilityLevel = sslCompatibilityLevel;
            Tags = tags;
            Type = type;
            Zone = zone;
        }
    }
}
