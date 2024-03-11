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
    /// Create and manage Scaleway DocumentDB database privilege.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
    /// 
    /// ## Example Usage
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
    ///     var main = new Scaleway.DocumentdbPrivilege("main", new()
    ///     {
    ///         DatabaseName = "my-db-name",
    ///         InstanceId = "11111111-1111-1111-1111-111111111111",
    ///         Permission = "all",
    ///         UserName = "my-db-user",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/documentdbPrivilege:DocumentdbPrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/documentdbPrivilege:DocumentdbPrivilege")]
    public partial class DocumentdbPrivilege : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the database (e.g. `my-db-name`).
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        /// </summary>
        [Output("permission")]
        public Output<string> Permission { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Name of the user (e.g. `my-db-user`).
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentdbPrivilege resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentdbPrivilege(string name, DocumentdbPrivilegeArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbPrivilege:DocumentdbPrivilege", name, args ?? new DocumentdbPrivilegeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentdbPrivilege(string name, Input<string> id, DocumentdbPrivilegeState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbPrivilege:DocumentdbPrivilege", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DocumentdbPrivilege resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentdbPrivilege Get(string name, Input<string> id, DocumentdbPrivilegeState? state = null, CustomResourceOptions? options = null)
        {
            return new DocumentdbPrivilege(name, id, state, options);
        }
    }

    public sealed class DocumentdbPrivilegeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database (e.g. `my-db-name`).
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Name of the user (e.g. `my-db-user`).
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public DocumentdbPrivilegeArgs()
        {
        }
        public static new DocumentdbPrivilegeArgs Empty => new DocumentdbPrivilegeArgs();
    }

    public sealed class DocumentdbPrivilegeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database (e.g. `my-db-name`).
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Name of the user (e.g. `my-db-user`).
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public DocumentdbPrivilegeState()
        {
        }
        public static new DocumentdbPrivilegeState Empty => new DocumentdbPrivilegeState();
    }
}
