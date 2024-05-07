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
    /// Provides an Object bucket lock configuration resource.
    /// For more information, see [Setting up object lock](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/).
    /// 
    /// ## Example Usage
    /// 
    /// ### Configure an Object Lock for a new bucket
    /// 
    /// Please note that `object_lock_enabled` must be set to `true` before configuring the lock.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainObjectBucket = new Scaleway.ObjectBucket("mainObjectBucket", new()
    ///     {
    ///         Acl = "public-read",
    ///         ObjectLockEnabled = true,
    ///     });
    /// 
    ///     var mainObjectBucketLockConfiguration = new Scaleway.ObjectBucketLockConfiguration("mainObjectBucketLockConfiguration", new()
    ///     {
    ///         Bucket = mainObjectBucket.Name,
    ///         Rule = new Scaleway.Inputs.ObjectBucketLockConfigurationRuleArgs
    ///         {
    ///             DefaultRetention = new Scaleway.Inputs.ObjectBucketLockConfigurationRuleDefaultRetentionArgs
    ///             {
    ///                 Mode = "GOVERNANCE",
    ///                 Days = 1,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Configure an Object Lock for an existing bucket
    /// 
    /// You should [contact Scaleway support](https://console.scaleway.com/support/tickets/create) to enable object lock on an existing bucket.
    /// 
    /// ## Import
    /// 
    /// Bucket lock configurations can be imported using the `{region}/{bucketName}` identifier, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration some_bucket fr-par/some-bucket
    /// ```
    /// 
    /// ~&gt; **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
    /// 
    /// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration")]
    public partial class ObjectBucketLockConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Specifies the Object Lock rule for the specified object.
        /// </summary>
        [Output("rule")]
        public Output<Outputs.ObjectBucketLockConfigurationRule> Rule { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectBucketLockConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectBucketLockConfiguration(string name, ObjectBucketLockConfigurationArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration", name, args ?? new ObjectBucketLockConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectBucketLockConfiguration(string name, Input<string> id, ObjectBucketLockConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectBucketLockConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectBucketLockConfiguration Get(string name, Input<string> id, ObjectBucketLockConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectBucketLockConfiguration(name, id, state, options);
        }
    }

    public sealed class ObjectBucketLockConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the Object Lock rule for the specified object.
        /// </summary>
        [Input("rule", required: true)]
        public Input<Inputs.ObjectBucketLockConfigurationRuleArgs> Rule { get; set; } = null!;

        public ObjectBucketLockConfigurationArgs()
        {
        }
        public static new ObjectBucketLockConfigurationArgs Empty => new ObjectBucketLockConfigurationArgs();
    }

    public sealed class ObjectBucketLockConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket's name or regional ID.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the Object Lock rule for the specified object.
        /// </summary>
        [Input("rule")]
        public Input<Inputs.ObjectBucketLockConfigurationRuleGetArgs>? Rule { get; set; }

        public ObjectBucketLockConfigurationState()
        {
        }
        public static new ObjectBucketLockConfigurationState Empty => new ObjectBucketLockConfigurationState();
    }
}
