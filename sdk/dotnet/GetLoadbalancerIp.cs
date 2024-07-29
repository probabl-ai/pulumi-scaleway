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
    public static class GetLoadbalancerIp
    {
        /// <summary>
        /// Gets information about a Load Balancer IP address.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
        /// 
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
        ///     var myIp = Scaleway.GetLoadbalancerIp.Invoke(new()
        ///     {
        ///         IpId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLoadbalancerIpResult> InvokeAsync(GetLoadbalancerIpArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoadbalancerIpResult>("scaleway:index/getLoadbalancerIp:getLoadbalancerIp", args ?? new GetLoadbalancerIpArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Load Balancer IP address.
        /// 
        /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
        /// 
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
        ///     var myIp = Scaleway.GetLoadbalancerIp.Invoke(new()
        ///     {
        ///         IpId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLoadbalancerIpResult> Invoke(GetLoadbalancerIpInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadbalancerIpResult>("scaleway:index/getLoadbalancerIp:getLoadbalancerIp", args ?? new GetLoadbalancerIpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadbalancerIpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        /// <summary>
        /// The IP ID.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipId")]
        public string? IpId { get; set; }

        /// <summary>
        /// The ID of the Project the Load Balancer IP is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetLoadbalancerIpArgs()
        {
        }
        public static new GetLoadbalancerIpArgs Empty => new GetLoadbalancerIpArgs();
    }

    public sealed class GetLoadbalancerIpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The IP ID.
        /// Only one of `ip_address` and `ip_id` should be specified.
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// The ID of the Project the Load Balancer IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetLoadbalancerIpInvokeArgs()
        {
        }
        public static new GetLoadbalancerIpInvokeArgs Empty => new GetLoadbalancerIpInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoadbalancerIpResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IpAddress;
        public readonly string? IpId;
        public readonly bool IsIpv6;
        /// <summary>
        /// The ID of the associated Load Balancer, if any
        /// </summary>
        public readonly string LbId;
        /// <summary>
        /// (Defaults to provider `organization_id`) The ID of the Organization the Load Balancer IP is associated with.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string Region;
        /// <summary>
        /// The reverse domain associated with this IP.
        /// </summary>
        public readonly string Reverse;
        /// <summary>
        /// The tags associated with this IP.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly string Zone;

        [OutputConstructor]
        private GetLoadbalancerIpResult(
            string id,

            string? ipAddress,

            string? ipId,

            bool isIpv6,

            string lbId,

            string organizationId,

            string projectId,

            string region,

            string reverse,

            ImmutableArray<string> tags,

            string zone)
        {
            Id = id;
            IpAddress = ipAddress;
            IpId = ipId;
            IsIpv6 = isIpv6;
            LbId = lbId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Reverse = reverse;
            Tags = tags;
            Zone = zone;
        }
    }
}
