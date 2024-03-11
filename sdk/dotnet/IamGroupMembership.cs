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
    /// Add members to an IAM group.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).
    /// 
    /// ## Example Usage
    /// 
    /// ### Application Membership
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @group = new Scaleway.IamGroup("group", new()
    ///     {
    ///         ExternalMembership = true,
    ///     });
    /// 
    ///     var app = new Scaleway.IamApplication("app");
    /// 
    ///     var member = new Scaleway.IamGroupMembership("member", new()
    ///     {
    ///         GroupId = @group.Id,
    ///         ApplicationId = app.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// IAM group memberships can be imported using two format:
    /// 
    /// - For user: `{group_id}/user/{user_id}`
    /// 
    /// - For application: `{group_id}/app/{application_id}`
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/iamGroupMembership:IamGroupMembership app 11111111-1111-1111-1111-111111111111/app/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/iamGroupMembership:IamGroupMembership")]
    public partial class IamGroupMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the application that will be added to the group.
        /// </summary>
        [Output("applicationId")]
        public Output<string?> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// ID of the group to add members to.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the user that will be added to the group
        /// 
        /// - &gt; Only one of `application_id` or `user_id` must be specified
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a IamGroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamGroupMembership(string name, IamGroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/iamGroupMembership:IamGroupMembership", name, args ?? new IamGroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamGroupMembership(string name, Input<string> id, IamGroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iamGroupMembership:IamGroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamGroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamGroupMembership Get(string name, Input<string> id, IamGroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new IamGroupMembership(name, id, state, options);
        }
    }

    public sealed class IamGroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the application that will be added to the group.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// ID of the group to add members to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The ID of the user that will be added to the group
        /// 
        /// - &gt; Only one of `application_id` or `user_id` must be specified
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public IamGroupMembershipArgs()
        {
        }
        public static new IamGroupMembershipArgs Empty => new IamGroupMembershipArgs();
    }

    public sealed class IamGroupMembershipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the application that will be added to the group.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// ID of the group to add members to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The ID of the user that will be added to the group
        /// 
        /// - &gt; Only one of `application_id` or `user_id` must be specified
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public IamGroupMembershipState()
        {
        }
        public static new IamGroupMembershipState Empty => new IamGroupMembershipState();
    }
}
