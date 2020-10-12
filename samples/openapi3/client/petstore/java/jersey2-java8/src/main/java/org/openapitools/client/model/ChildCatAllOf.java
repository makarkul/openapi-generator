/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import java.util.Map;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import org.openapitools.client.JSON;


/**
 * ChildCatAllOf
 */
@JsonPropertyOrder({
  ChildCatAllOf.JSON_PROPERTY_NAME,
  ChildCatAllOf.JSON_PROPERTY_PET_TYPE
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ChildCatAllOf {
  public static final String JSON_PROPERTY_NAME = "name";
  private String name;

  public static final String JSON_PROPERTY_PET_TYPE = "pet_type";
  private PetTypeEnum petType = PetTypeEnum.CHILDCAT;


  public ChildCatAllOf name(String name) {
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_NAME)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public static final Set<String> PET_TYPE_VALUES = new HashSet<>(Arrays.asList(
    "ChildCat"
  ));

  public ChildCatAllOf petType(PetTypeEnum petType) {
    if (!PET_TYPE_VALUES.contains(petType)) {
      throw new IllegalArgumentException(petType + " is invalid. Possible values for petType: " + String.join(", ", PET_TYPE_VALUES));
    }

    this.petType = petType;
    return this;
  }

   /**
   * Get petType
   * @return petType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")
  @JsonProperty(JSON_PROPERTY_PET_TYPE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public PetTypeEnum getPetType() {
    return petType;
  }


  public void setPetType(PetTypeEnum petType) {
    if (!PET_TYPE_VALUES.contains(petType)) {
      throw new IllegalArgumentException(petType + " is invalid. Possible values for petType: " + String.join(", ", PET_TYPE_VALUES));
    }

    this.petType = petType;
  }


  /**
   * Return true if this ChildCat_allOf object is equal to o.
   */
  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ChildCatAllOf childCatAllOf = (ChildCatAllOf) o;
    return Objects.equals(this.name, childCatAllOf.name) &&
        Objects.equals(this.petType, childCatAllOf.petType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, petType);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ChildCatAllOf {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    petType: ").append(toIndentedString(petType)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

