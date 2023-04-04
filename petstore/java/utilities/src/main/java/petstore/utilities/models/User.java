package petstore.utilities.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class User implements AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private Map<String, Object> additionalData;
    /** The email property */
    private String email;
    /** The firstName property */
    private String firstName;
    /** The id property */
    private Long id;
    /** The lastName property */
    private String lastName;
    /** The password property */
    private String password;
    /** The phone property */
    private String phone;
    /** The username property */
    private String username;
    /** User Status */
    private Integer userStatus;
    /**
     * Instantiates a new User and sets the default values.
     * @return a void
     */
    @javax.annotation.Nullable
    public User() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a User
     */
    @javax.annotation.Nonnull
    public static User createFromDiscriminatorValue(@javax.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new User();
    }
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this.additionalData;
    }
    /**
     * Gets the email property value. The email property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getEmail() {
        return this.email;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @javax.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(8);
        deserializerMap.put("email", (n) -> { this.setEmail(n.getStringValue()); });
        deserializerMap.put("firstName", (n) -> { this.setFirstName(n.getStringValue()); });
        deserializerMap.put("id", (n) -> { this.setId(n.getLongValue()); });
        deserializerMap.put("lastName", (n) -> { this.setLastName(n.getStringValue()); });
        deserializerMap.put("password", (n) -> { this.setPassword(n.getStringValue()); });
        deserializerMap.put("phone", (n) -> { this.setPhone(n.getStringValue()); });
        deserializerMap.put("username", (n) -> { this.setUsername(n.getStringValue()); });
        deserializerMap.put("userStatus", (n) -> { this.setUserStatus(n.getIntegerValue()); });
        return deserializerMap;
    }
    /**
     * Gets the firstName property value. The firstName property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getFirstName() {
        return this.firstName;
    }
    /**
     * Gets the id property value. The id property
     * @return a int64
     */
    @javax.annotation.Nullable
    public Long getId() {
        return this.id;
    }
    /**
     * Gets the lastName property value. The lastName property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getLastName() {
        return this.lastName;
    }
    /**
     * Gets the password property value. The password property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getPassword() {
        return this.password;
    }
    /**
     * Gets the phone property value. The phone property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getPhone() {
        return this.phone;
    }
    /**
     * Gets the username property value. The username property
     * @return a string
     */
    @javax.annotation.Nullable
    public String getUsername() {
        return this.username;
    }
    /**
     * Gets the userStatus property value. User Status
     * @return a integer
     */
    @javax.annotation.Nullable
    public Integer getUserStatus() {
        return this.userStatus;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    @javax.annotation.Nonnull
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("email", this.getEmail());
        writer.writeStringValue("firstName", this.getFirstName());
        writer.writeLongValue("id", this.getId());
        writer.writeStringValue("lastName", this.getLastName());
        writer.writeStringValue("password", this.getPassword());
        writer.writeStringValue("phone", this.getPhone());
        writer.writeStringValue("username", this.getUsername());
        writer.writeIntegerValue("userStatus", this.getUserStatus());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setAdditionalData(@javax.annotation.Nullable final Map<String, Object> value) {
        this.additionalData = value;
    }
    /**
     * Sets the email property value. The email property
     * @param value Value to set for the email property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setEmail(@javax.annotation.Nullable final String value) {
        this.email = value;
    }
    /**
     * Sets the firstName property value. The firstName property
     * @param value Value to set for the firstName property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setFirstName(@javax.annotation.Nullable final String value) {
        this.firstName = value;
    }
    /**
     * Sets the id property value. The id property
     * @param value Value to set for the id property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setId(@javax.annotation.Nullable final Long value) {
        this.id = value;
    }
    /**
     * Sets the lastName property value. The lastName property
     * @param value Value to set for the lastName property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setLastName(@javax.annotation.Nullable final String value) {
        this.lastName = value;
    }
    /**
     * Sets the password property value. The password property
     * @param value Value to set for the password property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setPassword(@javax.annotation.Nullable final String value) {
        this.password = value;
    }
    /**
     * Sets the phone property value. The phone property
     * @param value Value to set for the phone property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setPhone(@javax.annotation.Nullable final String value) {
        this.phone = value;
    }
    /**
     * Sets the username property value. The username property
     * @param value Value to set for the username property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setUsername(@javax.annotation.Nullable final String value) {
        this.username = value;
    }
    /**
     * Sets the userStatus property value. User Status
     * @param value Value to set for the userStatus property.
     * @return a void
     */
    @javax.annotation.Nonnull
    public void setUserStatus(@javax.annotation.Nullable final Integer value) {
        this.userStatus = value;
    }
}
