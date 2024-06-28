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
    /// <summary>
    /// Creates and manages Scaleway Load Balancer IP addresses.
    /// 
    /// For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ip = new Scaleway.LoadbalancerIp("ip", new()
    ///     {
    ///         Reverse = "my-reverse.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With IPv6
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ipv6 = new Scaleway.LoadbalancerIp("ipv6", new()
    ///     {
    ///         IsIpv6 = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IPs can be imported using `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/loadbalancerIp:LoadbalancerIp ip01 fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/loadbalancerIp:LoadbalancerIp")]
    public partial class LoadbalancerIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// If true, creates a flexible IP with an IPv6 address.
        /// </summary>
        [Output("isIpv6")]
        public Output<bool?> IsIpv6 { get; private set; } = null!;

        /// <summary>
        /// The associated Load Balancer ID if any
        /// </summary>
        [Output("lbId")]
        public Output<string> LbId { get; private set; } = null!;

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the Project the IP is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region of the resource
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The reverse domain associated with this IP.
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a LoadbalancerIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadbalancerIp(string name, LoadbalancerIpArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/loadbalancerIp:LoadbalancerIp", name, args ?? new LoadbalancerIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadbalancerIp(string name, Input<string> id, LoadbalancerIpState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/loadbalancerIp:LoadbalancerIp", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadbalancerIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadbalancerIp Get(string name, Input<string> id, LoadbalancerIpState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadbalancerIp(name, id, state, options);
        }
    }

    public sealed class LoadbalancerIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, creates a flexible IP with an IPv6 address.
        /// </summary>
        [Input("isIpv6")]
        public Input<bool>? IsIpv6 { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The reverse domain associated with this IP.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LoadbalancerIpArgs()
        {
        }
        public static new LoadbalancerIpArgs Empty => new LoadbalancerIpArgs();
    }

    public sealed class LoadbalancerIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// If true, creates a flexible IP with an IPv6 address.
        /// </summary>
        [Input("isIpv6")]
        public Input<bool>? IsIpv6 { get; set; }

        /// <summary>
        /// The associated Load Balancer ID if any
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the Project the IP is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region of the resource
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The reverse domain associated with this IP.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public LoadbalancerIpState()
        {
        }
        public static new LoadbalancerIpState Empty => new LoadbalancerIpState();
    }
}
