// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDocumentdbDatabaseResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String instanceId;
    /**
     * @return Whether the database is managed or not.
     * 
     */
    private Boolean managed;
    private @Nullable String name;
    /**
     * @return The name of the owner of the database.
     * 
     */
    private String owner;
    private String projectId;
    private @Nullable String region;
    /**
     * @return Size of the database (in bytes).
     * 
     */
    private String size;

    private GetDocumentdbDatabaseResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return Whether the database is managed or not.
     * 
     */
    public Boolean managed() {
        return this.managed;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The name of the owner of the database.
     * 
     */
    public String owner() {
        return this.owner;
    }
    public String projectId() {
        return this.projectId;
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    /**
     * @return Size of the database (in bytes).
     * 
     */
    public String size() {
        return this.size;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDocumentdbDatabaseResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String instanceId;
        private Boolean managed;
        private @Nullable String name;
        private String owner;
        private String projectId;
        private @Nullable String region;
        private String size;
        public Builder() {}
        public Builder(GetDocumentdbDatabaseResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.managed = defaults.managed;
    	      this.name = defaults.name;
    	      this.owner = defaults.owner;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.size = defaults.size;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder managed(Boolean managed) {
            this.managed = Objects.requireNonNull(managed);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder size(String size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        public GetDocumentdbDatabaseResult build() {
            final var o = new GetDocumentdbDatabaseResult();
            o.id = id;
            o.instanceId = instanceId;
            o.managed = managed;
            o.name = name;
            o.owner = owner;
            o.projectId = projectId;
            o.region = region;
            o.size = size;
            return o;
        }
    }
}
