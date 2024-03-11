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
    /// Creates and manages Scaleway Container.
    /// 
    /// For more information consult the [documentation](https://www.scaleway.com/en/docs/faq/serverless-containers/).
    /// 
    /// For more details about the limitation check [containers-limitations](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/).
    /// 
    /// You can check also our [containers guide](https://www.scaleway.com/en/docs/compute/containers/concepts/).
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
    ///     var mainContainerNamespace = new Scaleway.ContainerNamespace("mainContainerNamespace", new()
    ///     {
    ///         Description = "test container",
    ///     });
    /// 
    ///     var mainContainer = new Scaleway.Container("mainContainer", new()
    ///     {
    ///         Description = "environment variables test",
    ///         NamespaceId = mainContainerNamespace.Id,
    ///         RegistryImage = mainContainerNamespace.RegistryEndpoint.Apply(registryEndpoint =&gt; $"{registryEndpoint}/alpine:test"),
    ///         Port = 9997,
    ///         CpuLimit = 140,
    ///         MemoryLimit = 256,
    ///         MinScale = 3,
    ///         MaxScale = 5,
    ///         Timeout = 600,
    ///         MaxConcurrency = 80,
    ///         Privacy = "private",
    ///         Protocol = "http1",
    ///         Deploy = true,
    ///         EnvironmentVariables = 
    ///         {
    ///             { "foo", "var" },
    ///         },
    ///         SecretEnvironmentVariables = 
    ///         {
    ///             { "key", "secret" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Protocols
    /// 
    /// The supported protocols are:
    /// 
    /// * `h2c`: HTTP/2 over TCP.
    /// * `http1`: Hypertext Transfer Protocol.
    /// 
    /// **Important:** For details about the protocols check [this](https://httpd.apache.org/docs/2.4/howto/http2.html)
    /// 
    /// ## Privacy
    /// 
    /// By default, creating a container will make it `public`, meaning that anybody knowing the endpoint could execute it.
    /// A container can be made `private` with the privacy parameter.
    /// 
    /// Please check our [authentication](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) section
    /// 
    /// ## Memory and vCPUs configuration
    /// 
    /// The vCPU represents a portion or share of the underlying, physical CPU that is assigned to a particular virtual machine (VM).
    /// 
    /// You may decide how much computing resources to allocate to each container.
    /// The `memory_limit` (in MB) must correspond with the right amount of vCPU.
    /// 
    /// **Important:** The right choice for your container's resources is very important, as you will be billed based on compute usage over time and the number of Containers executions.
    /// 
    /// Please check our [price](https://www.scaleway.com/en/docs/faq/serverless-containers/#prices) section for more details.
    /// 
    /// | Memory (in MB) | vCPU |
    /// |----------------|------|
    /// | 128            | 70m  |
    /// | 256            | 140m |
    /// | 512            | 280m |
    /// | 1024           | 560m |
    /// 
    /// **Note:** 560mCPU accounts roughly for half of one CPU power of a Scaleway General Purpose instance
    /// 
    /// ## Import
    /// 
    /// Container can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/container:Container main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/container:Container")]
    public partial class Container : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
        /// </summary>
        [Output("cpuLimit")]
        public Output<int> CpuLimit { get; private set; } = null!;

        /// <summary>
        /// The cron status of the container.
        /// </summary>
        [Output("cronStatus")]
        public Output<string> CronStatus { get; private set; } = null!;

        /// <summary>
        /// Boolean controlling whether the container is on a production environment.
        /// 
        /// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
        /// </summary>
        [Output("deploy")]
        public Output<bool?> Deploy { get; private set; } = null!;

        /// <summary>
        /// The description of the container.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The native domain name of the container
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
        /// </summary>
        [Output("environmentVariables")]
        public Output<ImmutableDictionary<string, string>> EnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// The error message of the container.
        /// </summary>
        [Output("errorMessage")]
        public Output<string> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// HTTP traffic configuration
        /// </summary>
        [Output("httpOption")]
        public Output<string?> HttpOption { get; private set; } = null!;

        /// <summary>
        /// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
        /// </summary>
        [Output("maxConcurrency")]
        public Output<int> MaxConcurrency { get; private set; } = null!;

        /// <summary>
        /// The maximum of number of instances this container can scale to. Default to 20.
        /// </summary>
        [Output("maxScale")]
        public Output<int> MaxScale { get; private set; } = null!;

        /// <summary>
        /// The memory computing resources in MB to allocate to each container. Defaults to 256.
        /// </summary>
        [Output("memoryLimit")]
        public Output<int> MemoryLimit { get; private set; } = null!;

        /// <summary>
        /// The minimum of running container instances continuously. Defaults to 0.
        /// </summary>
        [Output("minScale")]
        public Output<int> MinScale { get; private set; } = null!;

        /// <summary>
        /// The unique name of the container name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The container namespace ID of the container.
        /// 
        /// &gt; **Important** Updates to `name` will recreate the container.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// The port to expose the container. Defaults to 8080.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
        /// </summary>
        [Output("privacy")]
        public Output<string?> Privacy { get; private set; } = null!;

        /// <summary>
        /// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// (Defaults to provider `region`) The region in which the container was created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
        /// </summary>
        [Output("registryImage")]
        public Output<string> RegistryImage { get; private set; } = null!;

        /// <summary>
        /// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
        /// </summary>
        [Output("registrySha256")]
        public Output<string?> RegistrySha256 { get; private set; } = null!;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
        /// </summary>
        [Output("secretEnvironmentVariables")]
        public Output<ImmutableDictionary<string, string>?> SecretEnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// The container status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a Container resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Container(string name, ContainerArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/container:Container", name, args ?? new ContainerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Container(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/container:Container", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "secretEnvironmentVariables",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Container resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Container Get(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
        {
            return new Container(name, id, state, options);
        }
    }

    public sealed class ContainerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
        /// </summary>
        [Input("cpuLimit")]
        public Input<int>? CpuLimit { get; set; }

        /// <summary>
        /// Boolean controlling whether the container is on a production environment.
        /// 
        /// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
        /// </summary>
        [Input("deploy")]
        public Input<bool>? Deploy { get; set; }

        /// <summary>
        /// The description of the container.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// HTTP traffic configuration
        /// </summary>
        [Input("httpOption")]
        public Input<string>? HttpOption { get; set; }

        /// <summary>
        /// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
        /// </summary>
        [Input("maxConcurrency")]
        public Input<int>? MaxConcurrency { get; set; }

        /// <summary>
        /// The maximum of number of instances this container can scale to. Default to 20.
        /// </summary>
        [Input("maxScale")]
        public Input<int>? MaxScale { get; set; }

        /// <summary>
        /// The memory computing resources in MB to allocate to each container. Defaults to 256.
        /// </summary>
        [Input("memoryLimit")]
        public Input<int>? MemoryLimit { get; set; }

        /// <summary>
        /// The minimum of running container instances continuously. Defaults to 0.
        /// </summary>
        [Input("minScale")]
        public Input<int>? MinScale { get; set; }

        /// <summary>
        /// The unique name of the container name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The container namespace ID of the container.
        /// 
        /// &gt; **Important** Updates to `name` will recreate the container.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        /// <summary>
        /// The port to expose the container. Defaults to 8080.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
        /// </summary>
        [Input("privacy")]
        public Input<string>? Privacy { get; set; }

        /// <summary>
        /// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// (Defaults to provider `region`) The region in which the container was created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
        /// </summary>
        [Input("registryImage")]
        public Input<string>? RegistryImage { get; set; }

        /// <summary>
        /// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
        /// </summary>
        [Input("registrySha256")]
        public Input<string>? RegistrySha256 { get; set; }

        [Input("secretEnvironmentVariables")]
        private InputMap<string>? _secretEnvironmentVariables;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
        /// </summary>
        public InputMap<string> SecretEnvironmentVariables
        {
            get => _secretEnvironmentVariables ?? (_secretEnvironmentVariables = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _secretEnvironmentVariables = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The container status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public ContainerArgs()
        {
        }
        public static new ContainerArgs Empty => new ContainerArgs();
    }

    public sealed class ContainerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
        /// </summary>
        [Input("cpuLimit")]
        public Input<int>? CpuLimit { get; set; }

        /// <summary>
        /// The cron status of the container.
        /// </summary>
        [Input("cronStatus")]
        public Input<string>? CronStatus { get; set; }

        /// <summary>
        /// Boolean controlling whether the container is on a production environment.
        /// 
        /// Note that if you want to use your own configuration, you must consult our configuration [restrictions](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/#configuration-restrictions) section.
        /// </summary>
        [Input("deploy")]
        public Input<bool>? Deploy { get; set; }

        /// <summary>
        /// The description of the container.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The native domain name of the container
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The [environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#environment-variables) variables of the container.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The error message of the container.
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// HTTP traffic configuration
        /// </summary>
        [Input("httpOption")]
        public Input<string>? HttpOption { get; set; }

        /// <summary>
        /// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
        /// </summary>
        [Input("maxConcurrency")]
        public Input<int>? MaxConcurrency { get; set; }

        /// <summary>
        /// The maximum of number of instances this container can scale to. Default to 20.
        /// </summary>
        [Input("maxScale")]
        public Input<int>? MaxScale { get; set; }

        /// <summary>
        /// The memory computing resources in MB to allocate to each container. Defaults to 256.
        /// </summary>
        [Input("memoryLimit")]
        public Input<int>? MemoryLimit { get; set; }

        /// <summary>
        /// The minimum of running container instances continuously. Defaults to 0.
        /// </summary>
        [Input("minScale")]
        public Input<int>? MinScale { get; set; }

        /// <summary>
        /// The unique name of the container name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The container namespace ID of the container.
        /// 
        /// &gt; **Important** Updates to `name` will recreate the container.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// The port to expose the container. Defaults to 8080.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The privacy type define the way to authenticate to your container. Please check our dedicated [section](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8).
        /// </summary>
        [Input("privacy")]
        public Input<string>? Privacy { get; set; }

        /// <summary>
        /// The communication [protocol](https://developers.scaleway.com/en/products/containers/api/#protocol-9dd4c8) http1 or h2c. Defaults to http1.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// (Defaults to provider `region`) The region in which the container was created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The registry image address. e.g: **"rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE"**.
        /// </summary>
        [Input("registryImage")]
        public Input<string>? RegistryImage { get; set; }

        /// <summary>
        /// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
        /// </summary>
        [Input("registrySha256")]
        public Input<string>? RegistrySha256 { get; set; }

        [Input("secretEnvironmentVariables")]
        private InputMap<string>? _secretEnvironmentVariables;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the container.
        /// </summary>
        public InputMap<string> SecretEnvironmentVariables
        {
            get => _secretEnvironmentVariables ?? (_secretEnvironmentVariables = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _secretEnvironmentVariables = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The container status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public ContainerState()
        {
        }
        public static new ContainerState Empty => new ContainerState();
    }
}
