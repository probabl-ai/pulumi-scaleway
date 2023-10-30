// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway
{
    /// <summary>
    /// Creates and manages Scaleway Web Hostings.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/webhosting/).
    /// 
    /// ## Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Lbrlabs.PulumiPackage.Scaleway;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var byName = Scaleway.GetWebHostOffer.Invoke(new()
    ///     {
    ///         Name = "lite",
    ///     });
    /// 
    ///     var main = new Scaleway.Webhosting("main", new()
    ///     {
    ///         OfferId = byName.Apply(getWebHostOfferResult =&gt; getWebHostOfferResult.OfferId),
    ///         Email = "your@email.com",
    ///         Domain = "yourdomain.com",
    ///         Tags = new[]
    ///         {
    ///             "webhosting",
    ///             "provider",
    ///             "terraform",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Hostings can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/webhosting:Webhosting hosting01 fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/webhosting:Webhosting")]
    public partial class Webhosting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The URL to connect to cPanel Dashboard and to Webmail interface.
        /// </summary>
        [Output("cpanelUrls")]
        public Output<ImmutableArray<Outputs.WebhostingCpanelUrl>> CpanelUrls { get; private set; } = null!;

        /// <summary>
        /// Date and time of hosting's creation (RFC 3339 format).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The DNS status of the hosting.
        /// </summary>
        [Output("dnsStatus")]
        public Output<string> DnsStatus { get; private set; } = null!;

        /// <summary>
        /// The domain name of the hosting.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The contact email of the client for the hosting.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The ID of the selected offer for the hosting.
        /// </summary>
        [Output("offerId")]
        public Output<string> OfferId { get; private set; } = null!;

        /// <summary>
        /// The name of the active offer.
        /// </summary>
        [Output("offerName")]
        public Output<string> OfferName { get; private set; } = null!;

        /// <summary>
        /// The IDs of the selected options for the hosting.
        /// </summary>
        [Output("optionIds")]
        public Output<ImmutableArray<string>> OptionIds { get; private set; } = null!;

        /// <summary>
        /// The active options of the hosting.
        /// </summary>
        [Output("options")]
        public Output<ImmutableArray<Outputs.WebhostingOption>> Options { get; private set; } = null!;

        /// <summary>
        /// The organization ID the hosting is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The hostname of the host platform.
        /// </summary>
        [Output("platformHostname")]
        public Output<string> PlatformHostname { get; private set; } = null!;

        /// <summary>
        /// The number of the host platform.
        /// </summary>
        [Output("platformNumber")]
        public Output<int> PlatformNumber { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region of the Hosting.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The hosting status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the hosting.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date and time of hosting's last update (RFC 3339 format).
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The main hosting cPanel username.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Webhosting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Webhosting(string name, WebhostingArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/webhosting:Webhosting", name, args ?? new WebhostingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Webhosting(string name, Input<string> id, WebhostingState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/webhosting:Webhosting", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Webhosting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Webhosting Get(string name, Input<string> id, WebhostingState? state = null, CustomResourceOptions? options = null)
        {
            return new Webhosting(name, id, state, options);
        }
    }

    public sealed class WebhostingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name of the hosting.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The contact email of the client for the hosting.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// The ID of the selected offer for the hosting.
        /// </summary>
        [Input("offerId", required: true)]
        public Input<string> OfferId { get; set; } = null!;

        [Input("optionIds")]
        private InputList<string>? _optionIds;

        /// <summary>
        /// The IDs of the selected options for the hosting.
        /// </summary>
        public InputList<string> OptionIds
        {
            get => _optionIds ?? (_optionIds = new InputList<string>());
            set => _optionIds = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region of the Hosting.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the hosting.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public WebhostingArgs()
        {
        }
        public static new WebhostingArgs Empty => new WebhostingArgs();
    }

    public sealed class WebhostingState : global::Pulumi.ResourceArgs
    {
        [Input("cpanelUrls")]
        private InputList<Inputs.WebhostingCpanelUrlGetArgs>? _cpanelUrls;

        /// <summary>
        /// The URL to connect to cPanel Dashboard and to Webmail interface.
        /// </summary>
        public InputList<Inputs.WebhostingCpanelUrlGetArgs> CpanelUrls
        {
            get => _cpanelUrls ?? (_cpanelUrls = new InputList<Inputs.WebhostingCpanelUrlGetArgs>());
            set => _cpanelUrls = value;
        }

        /// <summary>
        /// Date and time of hosting's creation (RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The DNS status of the hosting.
        /// </summary>
        [Input("dnsStatus")]
        public Input<string>? DnsStatus { get; set; }

        /// <summary>
        /// The domain name of the hosting.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The contact email of the client for the hosting.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The ID of the selected offer for the hosting.
        /// </summary>
        [Input("offerId")]
        public Input<string>? OfferId { get; set; }

        /// <summary>
        /// The name of the active offer.
        /// </summary>
        [Input("offerName")]
        public Input<string>? OfferName { get; set; }

        [Input("optionIds")]
        private InputList<string>? _optionIds;

        /// <summary>
        /// The IDs of the selected options for the hosting.
        /// </summary>
        public InputList<string> OptionIds
        {
            get => _optionIds ?? (_optionIds = new InputList<string>());
            set => _optionIds = value;
        }

        [Input("options")]
        private InputList<Inputs.WebhostingOptionGetArgs>? _options;

        /// <summary>
        /// The active options of the hosting.
        /// </summary>
        public InputList<Inputs.WebhostingOptionGetArgs> Options
        {
            get => _options ?? (_options = new InputList<Inputs.WebhostingOptionGetArgs>());
            set => _options = value;
        }

        /// <summary>
        /// The organization ID the hosting is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The hostname of the host platform.
        /// </summary>
        [Input("platformHostname")]
        public Input<string>? PlatformHostname { get; set; }

        /// <summary>
        /// The number of the host platform.
        /// </summary>
        [Input("platformNumber")]
        public Input<int>? PlatformNumber { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region of the Hosting.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The hosting status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the hosting.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Date and time of hosting's last update (RFC 3339 format).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The main hosting cPanel username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public WebhostingState()
        {
        }
        public static new WebhostingState Empty => new WebhostingState();
    }
}
