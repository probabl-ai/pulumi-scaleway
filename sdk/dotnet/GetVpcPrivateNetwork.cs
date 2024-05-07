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
    public static class GetVpcPrivateNetwork
    {
        /// <summary>
        /// Gets information about a private network.
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
        ///     var myName = Scaleway.GetVpcPrivateNetwork.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var myNameAndVpcId = Scaleway.GetVpcPrivateNetwork.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///         VpcId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var myId = Scaleway.GetVpcPrivateNetwork.Invoke(new()
        ///     {
        ///         PrivateNetworkId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVpcPrivateNetworkResult> InvokeAsync(GetVpcPrivateNetworkArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcPrivateNetworkResult>("scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork", args ?? new GetVpcPrivateNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a private network.
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
        ///     var myName = Scaleway.GetVpcPrivateNetwork.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var myNameAndVpcId = Scaleway.GetVpcPrivateNetwork.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///         VpcId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var myId = Scaleway.GetVpcPrivateNetwork.Invoke(new()
        ///     {
        ///         PrivateNetworkId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVpcPrivateNetworkResult> Invoke(GetVpcPrivateNetworkInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcPrivateNetworkResult>("scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork", args ?? new GetVpcPrivateNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcPrivateNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the private network. Cannot be used with `private_network_id`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// ID of the private network. Cannot be used with `name` and `vpc_id`.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the project the private network is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// ID of the VPC in which the private network is. Cannot be used with `private_network_id`.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVpcPrivateNetworkArgs()
        {
        }
        public static new GetVpcPrivateNetworkArgs Empty => new GetVpcPrivateNetworkArgs();
    }

    public sealed class GetVpcPrivateNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the private network. Cannot be used with `private_network_id`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the private network. Cannot be used with `name` and `vpc_id`.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the project the private network is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// ID of the VPC in which the private network is. Cannot be used with `private_network_id`.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetVpcPrivateNetworkInvokeArgs()
        {
        }
        public static new GetVpcPrivateNetworkInvokeArgs Empty => new GetVpcPrivateNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcPrivateNetworkResult
    {
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IPv4 subnet associated with the private network.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcPrivateNetworkIpv4SubnetResult> Ipv4Subnets;
        /// <summary>
        /// The IPv6 subnets associated with the private network.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcPrivateNetworkIpv6SubnetResult> Ipv6Subnets;
        public readonly bool IsRegional;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string? PrivateNetworkId;
        public readonly string? ProjectId;
        public readonly string Region;
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdatedAt;
        public readonly string? VpcId;
        public readonly string Zone;

        [OutputConstructor]
        private GetVpcPrivateNetworkResult(
            string createdAt,

            string id,

            ImmutableArray<Outputs.GetVpcPrivateNetworkIpv4SubnetResult> ipv4Subnets,

            ImmutableArray<Outputs.GetVpcPrivateNetworkIpv6SubnetResult> ipv6Subnets,

            bool isRegional,

            string? name,

            string organizationId,

            string? privateNetworkId,

            string? projectId,

            string region,

            ImmutableArray<string> tags,

            string updatedAt,

            string? vpcId,

            string zone)
        {
            CreatedAt = createdAt;
            Id = id;
            Ipv4Subnets = ipv4Subnets;
            Ipv6Subnets = ipv6Subnets;
            IsRegional = isRegional;
            Name = name;
            OrganizationId = organizationId;
            PrivateNetworkId = privateNetworkId;
            ProjectId = projectId;
            Region = region;
            Tags = tags;
            UpdatedAt = updatedAt;
            VpcId = vpcId;
            Zone = zone;
        }
    }
}
