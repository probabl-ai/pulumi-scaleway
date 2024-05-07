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
    ///     var main = new Scaleway.IotHub("main", new()
    ///     {
    ///         ProductPlan = "plan_shared",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IoT Hubs can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/iotHub:IotHub hub01 fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/iotHub:IotHub")]
    public partial class IotHub : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The current number of connected devices in the Hub.
        /// </summary>
        [Output("connectedDeviceCount")]
        public Output<int> ConnectedDeviceCount { get; private set; } = null!;

        /// <summary>
        /// The date and time the Hub was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Wether to enable the device auto provisioning or not
        /// </summary>
        [Output("deviceAutoProvisioning")]
        public Output<bool?> DeviceAutoProvisioning { get; private set; } = null!;

        /// <summary>
        /// The number of registered devices in the Hub.
        /// </summary>
        [Output("deviceCount")]
        public Output<int> DeviceCount { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the hub events or not
        /// </summary>
        [Output("disableEvents")]
        public Output<bool?> DisableEvents { get; private set; } = null!;

        /// <summary>
        /// Wether the IoT Hub instance should be enabled or not.
        /// 
        /// &gt; **Important:** Updates to `enabled` will disconnect eventually connected devices.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The MQTT network endpoint to connect MQTT devices to.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Topic prefix for the hub events
        /// </summary>
        [Output("eventsTopicPrefix")]
        public Output<string?> EventsTopicPrefix { get; private set; } = null!;

        /// <summary>
        /// Custom user provided certificate authority
        /// </summary>
        [Output("hubCa")]
        public Output<string?> HubCa { get; private set; } = null!;

        /// <summary>
        /// Challenge certificate for the user provided hub CA
        /// </summary>
        [Output("hubCaChallenge")]
        public Output<string?> HubCaChallenge { get; private set; } = null!;

        /// <summary>
        /// The MQTT certificat content
        /// </summary>
        [Output("mqttCa")]
        public Output<string> MqttCa { get; private set; } = null!;

        /// <summary>
        /// The MQTT ca url
        /// </summary>
        [Output("mqttCaUrl")]
        public Output<string> MqttCaUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        /// 
        /// &gt; **Important:** Updates to `product_plan` will recreate the IoT Hub Instance.
        /// </summary>
        [Output("productPlan")]
        public Output<string> ProductPlan { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the IoT Hub Instance is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The current status of the Hub.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The date and time the Hub resource was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a IotHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotHub(string name, IotHubArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/iotHub:IotHub", name, args ?? new IotHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotHub(string name, Input<string> id, IotHubState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iotHub:IotHub", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IotHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotHub Get(string name, Input<string> id, IotHubState? state = null, CustomResourceOptions? options = null)
        {
            return new IotHub(name, id, state, options);
        }
    }

    public sealed class IotHubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wether to enable the device auto provisioning or not
        /// </summary>
        [Input("deviceAutoProvisioning")]
        public Input<bool>? DeviceAutoProvisioning { get; set; }

        /// <summary>
        /// Whether to enable the hub events or not
        /// </summary>
        [Input("disableEvents")]
        public Input<bool>? DisableEvents { get; set; }

        /// <summary>
        /// Wether the IoT Hub instance should be enabled or not.
        /// 
        /// &gt; **Important:** Updates to `enabled` will disconnect eventually connected devices.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Topic prefix for the hub events
        /// </summary>
        [Input("eventsTopicPrefix")]
        public Input<string>? EventsTopicPrefix { get; set; }

        /// <summary>
        /// Custom user provided certificate authority
        /// </summary>
        [Input("hubCa")]
        public Input<string>? HubCa { get; set; }

        /// <summary>
        /// Challenge certificate for the user provided hub CA
        /// </summary>
        [Input("hubCaChallenge")]
        public Input<string>? HubCaChallenge { get; set; }

        /// <summary>
        /// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        /// 
        /// &gt; **Important:** Updates to `product_plan` will recreate the IoT Hub Instance.
        /// </summary>
        [Input("productPlan", required: true)]
        public Input<string> ProductPlan { get; set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the IoT Hub Instance is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public IotHubArgs()
        {
        }
        public static new IotHubArgs Empty => new IotHubArgs();
    }

    public sealed class IotHubState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The current number of connected devices in the Hub.
        /// </summary>
        [Input("connectedDeviceCount")]
        public Input<int>? ConnectedDeviceCount { get; set; }

        /// <summary>
        /// The date and time the Hub was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Wether to enable the device auto provisioning or not
        /// </summary>
        [Input("deviceAutoProvisioning")]
        public Input<bool>? DeviceAutoProvisioning { get; set; }

        /// <summary>
        /// The number of registered devices in the Hub.
        /// </summary>
        [Input("deviceCount")]
        public Input<int>? DeviceCount { get; set; }

        /// <summary>
        /// Whether to enable the hub events or not
        /// </summary>
        [Input("disableEvents")]
        public Input<bool>? DisableEvents { get; set; }

        /// <summary>
        /// Wether the IoT Hub instance should be enabled or not.
        /// 
        /// &gt; **Important:** Updates to `enabled` will disconnect eventually connected devices.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The MQTT network endpoint to connect MQTT devices to.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Topic prefix for the hub events
        /// </summary>
        [Input("eventsTopicPrefix")]
        public Input<string>? EventsTopicPrefix { get; set; }

        /// <summary>
        /// Custom user provided certificate authority
        /// </summary>
        [Input("hubCa")]
        public Input<string>? HubCa { get; set; }

        /// <summary>
        /// Challenge certificate for the user provided hub CA
        /// </summary>
        [Input("hubCaChallenge")]
        public Input<string>? HubCaChallenge { get; set; }

        /// <summary>
        /// The MQTT certificat content
        /// </summary>
        [Input("mqttCa")]
        public Input<string>? MqttCa { get; set; }

        /// <summary>
        /// The MQTT ca url
        /// </summary>
        [Input("mqttCaUrl")]
        public Input<string>? MqttCaUrl { get; set; }

        /// <summary>
        /// The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization_id you want to attach the resource to
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        /// 
        /// &gt; **Important:** Updates to `product_plan` will recreate the IoT Hub Instance.
        /// </summary>
        [Input("productPlan")]
        public Input<string>? ProductPlan { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the IoT Hub Instance is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The current status of the Hub.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date and time the Hub resource was updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public IotHubState()
        {
        }
        public static new IotHubState Empty => new IotHubState();
    }
}
