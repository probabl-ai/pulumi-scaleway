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
    public static class GetInstanceIp
    {
        /// <summary>
        /// Gets information about an instance IP.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myIp = Scaleway.GetInstanceIp.Invoke(new()
        ///     {
        ///         Id = "fr-par-1/11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetInstanceIpResult> InvokeAsync(GetInstanceIpArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceIpResult>("scaleway:index/getInstanceIp:getInstanceIp", args ?? new GetInstanceIpArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an instance IP.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myIp = Scaleway.GetInstanceIp.Invoke(new()
        ///     {
        ///         Id = "fr-par-1/11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetInstanceIpResult> Invoke(GetInstanceIpInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceIpResult>("scaleway:index/getInstanceIp:getInstanceIp", args ?? new GetInstanceIpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceIpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IPv4 address to retrieve
        /// Only one of `address` and `id` should be specified.
        /// </summary>
        [Input("address")]
        public string? Address { get; set; }

        /// <summary>
        /// The ID of the IP address to retrieve
        /// Only one of `address` and `id` should be specified.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        public GetInstanceIpArgs()
        {
        }
        public static new GetInstanceIpArgs Empty => new GetInstanceIpArgs();
    }

    public sealed class GetInstanceIpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IPv4 address to retrieve
        /// Only one of `address` and `id` should be specified.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The ID of the IP address to retrieve
        /// Only one of `address` and `id` should be specified.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetInstanceIpInvokeArgs()
        {
        }
        public static new GetInstanceIpInvokeArgs Empty => new GetInstanceIpInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceIpResult
    {
        /// <summary>
        /// The IP address.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The ID of the IP.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The organization ID the IP is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The IP Prefix.
        /// </summary>
        public readonly string Prefix;
        public readonly string ProjectId;
        /// <summary>
        /// The reverse dns attached to this IP
        /// </summary>
        public readonly string Reverse;
        public readonly string ServerId;
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of the IP
        /// </summary>
        public readonly string Type;
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceIpResult(
            string? address,

            string? id,

            string organizationId,

            string prefix,

            string projectId,

            string reverse,

            string serverId,

            ImmutableArray<string> tags,

            string type,

            string zone)
        {
            Address = address;
            Id = id;
            OrganizationId = organizationId;
            Prefix = prefix;
            ProjectId = projectId;
            Reverse = reverse;
            ServerId = serverId;
            Tags = tags;
            Type = type;
            Zone = zone;
        }
    }
}
