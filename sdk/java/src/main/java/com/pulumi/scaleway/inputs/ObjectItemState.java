// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ObjectItemState extends com.pulumi.resources.ResourceArgs {

    public static final ObjectItemState Empty = new ObjectItemState();

    /**
     * The name of the bucket.
     * 
     */
    @Import(name="bucket")
    private @Nullable Output<String> bucket;

    /**
     * @return The name of the bucket.
     * 
     */
    public Optional<Output<String>> bucket() {
        return Optional.ofNullable(this.bucket);
    }

    /**
     * The name of the file to upload, defaults to an empty file
     * 
     */
    @Import(name="file")
    private @Nullable Output<String> file;

    /**
     * @return The name of the file to upload, defaults to an empty file
     * 
     */
    public Optional<Output<String>> file() {
        return Optional.ofNullable(this.file);
    }

    /**
     * Hash of the file, used to trigger upload on file change
     * 
     */
    @Import(name="hash")
    private @Nullable Output<String> hash;

    /**
     * @return Hash of the file, used to trigger upload on file change
     * 
     */
    public Optional<Output<String>> hash() {
        return Optional.ofNullable(this.hash);
    }

    /**
     * The path of the object.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The path of the object.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * Map of metadata used for the object, keys must be lowercase
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,String>> metadata;

    /**
     * @return Map of metadata used for the object, keys must be lowercase
     * 
     */
    public Optional<Output<Map<String,String>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * The Scaleway region this bucket resides in.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The Scaleway region this bucket resides in.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     * 
     */
    @Import(name="storageClass")
    private @Nullable Output<String> storageClass;

    /**
     * @return Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
     * 
     */
    public Optional<Output<String>> storageClass() {
        return Optional.ofNullable(this.storageClass);
    }

    /**
     * Map of tags
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return Map of tags
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Visibility of the object, `public-read` or `private`
     * 
     */
    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    /**
     * @return Visibility of the object, `public-read` or `private`
     * 
     */
    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    private ObjectItemState() {}

    private ObjectItemState(ObjectItemState $) {
        this.bucket = $.bucket;
        this.file = $.file;
        this.hash = $.hash;
        this.key = $.key;
        this.metadata = $.metadata;
        this.region = $.region;
        this.storageClass = $.storageClass;
        this.tags = $.tags;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectItemState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectItemState $;

        public Builder() {
            $ = new ObjectItemState();
        }

        public Builder(ObjectItemState defaults) {
            $ = new ObjectItemState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(@Nullable Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param file The name of the file to upload, defaults to an empty file
         * 
         * @return builder
         * 
         */
        public Builder file(@Nullable Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file The name of the file to upload, defaults to an empty file
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        /**
         * @param hash Hash of the file, used to trigger upload on file change
         * 
         * @return builder
         * 
         */
        public Builder hash(@Nullable Output<String> hash) {
            $.hash = hash;
            return this;
        }

        /**
         * @param hash Hash of the file, used to trigger upload on file change
         * 
         * @return builder
         * 
         */
        public Builder hash(String hash) {
            return hash(Output.of(hash));
        }

        /**
         * @param key The path of the object.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The path of the object.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param metadata Map of metadata used for the object, keys must be lowercase
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,String>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Map of metadata used for the object, keys must be lowercase
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,String> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param region The Scaleway region this bucket resides in.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The Scaleway region this bucket resides in.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param storageClass Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(@Nullable Output<String> storageClass) {
            $.storageClass = storageClass;
            return this;
        }

        /**
         * @param storageClass Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
         * 
         * @return builder
         * 
         */
        public Builder storageClass(String storageClass) {
            return storageClass(Output.of(storageClass));
        }

        /**
         * @param tags Map of tags
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param visibility Visibility of the object, `public-read` or `private`
         * 
         * @return builder
         * 
         */
        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility Visibility of the object, `public-read` or `private`
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public ObjectItemState build() {
            return $;
        }
    }

}
