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
    public static class GetBaremetalServer
    {
        /// <summary>
        /// Gets information about a baremetal server.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
        ///     var byName = Scaleway.GetBaremetalServer.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        ///     var byId = Scaleway.GetBaremetalServer.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBaremetalServerResult> InvokeAsync(GetBaremetalServerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBaremetalServerResult>("scaleway:index/getBaremetalServer:getBaremetalServer", args ?? new GetBaremetalServerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a baremetal server.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
        ///     var byName = Scaleway.GetBaremetalServer.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        ///     var byId = Scaleway.GetBaremetalServer.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBaremetalServerResult> Invoke(GetBaremetalServerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBaremetalServerResult>("scaleway:index/getBaremetalServer:getBaremetalServer", args ?? new GetBaremetalServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBaremetalServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The server name. Only one of `name` and `server_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the baremetal server is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("serverId")]
        public string? ServerId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetBaremetalServerArgs()
        {
        }
        public static new GetBaremetalServerArgs Empty => new GetBaremetalServerArgs();
    }

    public sealed class GetBaremetalServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The server name. Only one of `name` and `server_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the baremetal server is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetBaremetalServerInvokeArgs()
        {
        }
        public static new GetBaremetalServerInvokeArgs Empty => new GetBaremetalServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetBaremetalServerResult
    {
        public readonly string Description;
        public readonly string Domain;
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool InstallConfigAfterward;
        public readonly ImmutableArray<Outputs.GetBaremetalServerIpResult> Ips;
        public readonly ImmutableArray<Outputs.GetBaremetalServerIpv4Result> Ipv4s;
        public readonly ImmutableArray<Outputs.GetBaremetalServerIpv6Result> Ipv6s;
        public readonly string? Name;
        public readonly string Offer;
        public readonly string OfferId;
        public readonly string OfferName;
        public readonly ImmutableArray<Outputs.GetBaremetalServerOptionResult> Options;
        public readonly string OrganizationId;
        public readonly string Os;
        public readonly string OsName;
        public readonly string Password;
        public readonly ImmutableArray<Outputs.GetBaremetalServerPrivateNetworkResult> PrivateNetworks;
        public readonly string? ProjectId;
        public readonly bool ReinstallOnConfigChanges;
        public readonly string? ServerId;
        public readonly string ServicePassword;
        public readonly string ServiceUser;
        public readonly ImmutableArray<string> SshKeyIds;
        public readonly ImmutableArray<string> Tags;
        public readonly string User;
        public readonly string? Zone;

        [OutputConstructor]
        private GetBaremetalServerResult(
            string description,

            string domain,

            string hostname,

            string id,

            bool installConfigAfterward,

            ImmutableArray<Outputs.GetBaremetalServerIpResult> ips,

            ImmutableArray<Outputs.GetBaremetalServerIpv4Result> ipv4s,

            ImmutableArray<Outputs.GetBaremetalServerIpv6Result> ipv6s,

            string? name,

            string offer,

            string offerId,

            string offerName,

            ImmutableArray<Outputs.GetBaremetalServerOptionResult> options,

            string organizationId,

            string os,

            string osName,

            string password,

            ImmutableArray<Outputs.GetBaremetalServerPrivateNetworkResult> privateNetworks,

            string? projectId,

            bool reinstallOnConfigChanges,

            string? serverId,

            string servicePassword,

            string serviceUser,

            ImmutableArray<string> sshKeyIds,

            ImmutableArray<string> tags,

            string user,

            string? zone)
        {
            Description = description;
            Domain = domain;
            Hostname = hostname;
            Id = id;
            InstallConfigAfterward = installConfigAfterward;
            Ips = ips;
            Ipv4s = ipv4s;
            Ipv6s = ipv6s;
            Name = name;
            Offer = offer;
            OfferId = offerId;
            OfferName = offerName;
            Options = options;
            OrganizationId = organizationId;
            Os = os;
            OsName = osName;
            Password = password;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            ReinstallOnConfigChanges = reinstallOnConfigChanges;
            ServerId = serverId;
            ServicePassword = servicePassword;
            ServiceUser = serviceUser;
            SshKeyIds = sshKeyIds;
            Tags = tags;
            User = user;
            Zone = zone;
        }
    }
}
