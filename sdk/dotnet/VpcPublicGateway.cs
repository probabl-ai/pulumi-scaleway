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
    /// Creates and manages Scaleway VPC Public Gateway.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/public-gateway).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.VpcPublicGateway("main", new()
    ///     {
    ///         Tags = new[]
    ///         {
    ///             "demo",
    ///             "terraform",
    ///         },
    ///         Type = "VPC-GW-S",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Public gateway can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/vpcPublicGateway:VpcPublicGateway main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/vpcPublicGateway:VpcPublicGateway")]
    public partial class VpcPublicGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable SSH bastion on the gateway
        /// </summary>
        [Output("bastionEnabled")]
        public Output<bool?> BastionEnabled { get; private set; } = null!;

        /// <summary>
        /// The port on which the SSH bastion will listen.
        /// </summary>
        [Output("bastionPort")]
        public Output<int> BastionPort { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the public gateway.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Enable SMTP on the gateway
        /// </summary>
        [Output("enableSmtp")]
        public Output<bool> EnableSmtp { get; private set; } = null!;

        /// <summary>
        /// attach an existing flexible IP to the gateway
        /// </summary>
        [Output("ipId")]
        public Output<string> IpId { get; private set; } = null!;

        /// <summary>
        /// The name of the public gateway. If not provided it will be randomly generated.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the public gateway is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the public gateway is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The status of the public gateway.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the public gateway.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The gateway type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the public gateway.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// override the gateway's default recursive DNS servers, if DNS features are enabled.
        /// </summary>
        [Output("upstreamDnsServers")]
        public Output<ImmutableArray<string>> UpstreamDnsServers { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the public gateway should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPublicGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPublicGateway(string name, VpcPublicGatewayArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, args ?? new VpcPublicGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPublicGateway(string name, Input<string> id, VpcPublicGatewayState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGateway:VpcPublicGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcPublicGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPublicGateway Get(string name, Input<string> id, VpcPublicGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPublicGateway(name, id, state, options);
        }
    }

    public sealed class VpcPublicGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable SSH bastion on the gateway
        /// </summary>
        [Input("bastionEnabled")]
        public Input<bool>? BastionEnabled { get; set; }

        /// <summary>
        /// The port on which the SSH bastion will listen.
        /// </summary>
        [Input("bastionPort")]
        public Input<int>? BastionPort { get; set; }

        /// <summary>
        /// Enable SMTP on the gateway
        /// </summary>
        [Input("enableSmtp")]
        public Input<bool>? EnableSmtp { get; set; }

        /// <summary>
        /// attach an existing flexible IP to the gateway
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// The name of the public gateway. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the public gateway is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the public gateway.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The gateway type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("upstreamDnsServers")]
        private InputList<string>? _upstreamDnsServers;

        /// <summary>
        /// override the gateway's default recursive DNS servers, if DNS features are enabled.
        /// </summary>
        public InputList<string> UpstreamDnsServers
        {
            get => _upstreamDnsServers ?? (_upstreamDnsServers = new InputList<string>());
            set => _upstreamDnsServers = value;
        }

        /// <summary>
        /// `zone`) The zone in which the public gateway should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayArgs()
        {
        }
        public static new VpcPublicGatewayArgs Empty => new VpcPublicGatewayArgs();
    }

    public sealed class VpcPublicGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable SSH bastion on the gateway
        /// </summary>
        [Input("bastionEnabled")]
        public Input<bool>? BastionEnabled { get; set; }

        /// <summary>
        /// The port on which the SSH bastion will listen.
        /// </summary>
        [Input("bastionPort")]
        public Input<int>? BastionPort { get; set; }

        /// <summary>
        /// The date and time of the creation of the public gateway.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Enable SMTP on the gateway
        /// </summary>
        [Input("enableSmtp")]
        public Input<bool>? EnableSmtp { get; set; }

        /// <summary>
        /// attach an existing flexible IP to the gateway
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        /// <summary>
        /// The name of the public gateway. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the public gateway is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the public gateway is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The status of the public gateway.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the public gateway.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The gateway type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The date and time of the last update of the public gateway.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("upstreamDnsServers")]
        private InputList<string>? _upstreamDnsServers;

        /// <summary>
        /// override the gateway's default recursive DNS servers, if DNS features are enabled.
        /// </summary>
        public InputList<string> UpstreamDnsServers
        {
            get => _upstreamDnsServers ?? (_upstreamDnsServers = new InputList<string>());
            set => _upstreamDnsServers = value;
        }

        /// <summary>
        /// `zone`) The zone in which the public gateway should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayState()
        {
        }
        public static new VpcPublicGatewayState Empty => new VpcPublicGatewayState();
    }
}
