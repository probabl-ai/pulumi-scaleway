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
    public static class GetDatabase
    {
        /// <summary>
        /// Gets information about a RDB database.
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
        ///     var myDb = Scaleway.GetDatabase.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///         Name = "foobar",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("scaleway:index/getDatabase:getDatabase", args ?? new GetDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a RDB database.
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
        ///     var myDb = Scaleway.GetDatabase.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///         Name = "foobar",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseResult> Invoke(GetDatabaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseResult>("scaleway:index/getDatabase:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the RDB instance.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("region")]
        public string? Region { get; set; }

        public GetDatabaseArgs()
        {
        }
        public static new GetDatabaseArgs Empty => new GetDatabaseArgs();
    }

    public sealed class GetDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the RDB instance.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetDatabaseInvokeArgs()
        {
        }
        public static new GetDatabaseInvokeArgs Empty => new GetDatabaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Whether the database is managed or not.
        /// </summary>
        public readonly bool Managed;
        public readonly string Name;
        /// <summary>
        /// The name of the owner of the database.
        /// </summary>
        public readonly string Owner;
        public readonly string? Region;
        /// <summary>
        /// Size of the database (in bytes).
        /// </summary>
        public readonly string Size;

        [OutputConstructor]
        private GetDatabaseResult(
            string id,

            string instanceId,

            bool managed,

            string name,

            string owner,

            string? region,

            string size)
        {
            Id = id;
            InstanceId = instanceId;
            Managed = managed;
            Name = name;
            Owner = owner;
            Region = region;
            Size = size;
        }
    }
}
