# frozen_string_literal: true

# Typed models for the Nominatim SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# AddressLookup entity data model.
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] boundingbox
#   @return [Array, nil]
#
# @!attribute [rw] class
#   @return [String, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] importance
#   @return [Float, nil]
#
# @!attribute [rw] lat
#   @return [String, nil]
#
# @!attribute [rw] licence
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
AddressLookup = Struct.new(
  :address,
  :boundingbox,
  :class,
  :display_name,
  :importance,
  :lat,
  :licence,
  :lon,
  :osm_id,
  :osm_type,
  :place_id,
  :type,
  keyword_init: true
)

# Request payload for AddressLookup#list.
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] boundingbox
#   @return [Array, nil]
#
# @!attribute [rw] class
#   @return [String, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] importance
#   @return [Float, nil]
#
# @!attribute [rw] lat
#   @return [String, nil]
#
# @!attribute [rw] licence
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
AddressLookupListMatch = Struct.new(
  :address,
  :boundingbox,
  :class,
  :display_name,
  :importance,
  :lat,
  :licence,
  :lon,
  :osm_id,
  :osm_type,
  :place_id,
  :type,
  keyword_init: true
)

# Administrative entity data model.
#
# @!attribute [rw] class
#   @return [String, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] errormessage
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] updated
#   @return [String, nil]
Administrative = Struct.new(
  :class,
  :country_code,
  :errormessage,
  :name,
  :osm_id,
  :osm_type,
  :place_id,
  :type,
  :updated,
  keyword_init: true
)

# Request payload for Administrative#list.
#
# @!attribute [rw] class
#   @return [String, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] errormessage
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] updated
#   @return [String, nil]
AdministrativeListMatch = Struct.new(
  :class,
  :country_code,
  :errormessage,
  :name,
  :osm_id,
  :osm_type,
  :place_id,
  :type,
  :updated,
  keyword_init: true
)

# Debug entity data model.
#
# @!attribute [rw] addresstag
#   @return [Hash, nil]
#
# @!attribute [rw] admin_level
#   @return [Integer, nil]
#
# @!attribute [rw] calculated_importance
#   @return [Float, nil]
#
# @!attribute [rw] calculated_postcode
#   @return [String, nil]
#
# @!attribute [rw] calculated_wikipedia
#   @return [String, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] centroid
#   @return [Hash, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] extratag
#   @return [Hash, nil]
#
# @!attribute [rw] geometry
#   @return [Hash, nil]
#
# @!attribute [rw] housenumber
#   @return [String, nil]
#
# @!attribute [rw] importance
#   @return [Float, nil]
#
# @!attribute [rw] indexed_date
#   @return [String, nil]
#
# @!attribute [rw] isarea
#   @return [Boolean, nil]
#
# @!attribute [rw] localname
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [Hash, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] parent_place_id
#   @return [Integer, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] rank_address
#   @return [Integer, nil]
#
# @!attribute [rw] rank_search
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
Debug = Struct.new(
  :addresstag,
  :admin_level,
  :calculated_importance,
  :calculated_postcode,
  :calculated_wikipedia,
  :category,
  :centroid,
  :country_code,
  :extratag,
  :geometry,
  :housenumber,
  :importance,
  :indexed_date,
  :isarea,
  :localname,
  :name,
  :osm_id,
  :osm_type,
  :parent_place_id,
  :place_id,
  :rank_address,
  :rank_search,
  :type,
  keyword_init: true
)

# Request payload for Debug#load.
#
# @!attribute [rw] addresstag
#   @return [Hash, nil]
#
# @!attribute [rw] admin_level
#   @return [Integer, nil]
#
# @!attribute [rw] calculated_importance
#   @return [Float, nil]
#
# @!attribute [rw] calculated_postcode
#   @return [String, nil]
#
# @!attribute [rw] calculated_wikipedia
#   @return [String, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] centroid
#   @return [Hash, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] extratag
#   @return [Hash, nil]
#
# @!attribute [rw] geometry
#   @return [Hash, nil]
#
# @!attribute [rw] housenumber
#   @return [String, nil]
#
# @!attribute [rw] importance
#   @return [Float, nil]
#
# @!attribute [rw] indexed_date
#   @return [String, nil]
#
# @!attribute [rw] isarea
#   @return [Boolean, nil]
#
# @!attribute [rw] localname
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [Hash, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] parent_place_id
#   @return [Integer, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] rank_address
#   @return [Integer, nil]
#
# @!attribute [rw] rank_search
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
DebugLoadMatch = Struct.new(
  :addresstag,
  :admin_level,
  :calculated_importance,
  :calculated_postcode,
  :calculated_wikipedia,
  :category,
  :centroid,
  :country_code,
  :extratag,
  :geometry,
  :housenumber,
  :importance,
  :indexed_date,
  :isarea,
  :localname,
  :name,
  :osm_id,
  :osm_type,
  :parent_place_id,
  :place_id,
  :rank_address,
  :rank_search,
  :type,
  keyword_init: true
)

# Reverse entity data model.
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] boundingbox
#   @return [Array, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] lat
#   @return [String, nil]
#
# @!attribute [rw] licence
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
Reverse = Struct.new(
  :address,
  :boundingbox,
  :display_name,
  :lat,
  :licence,
  :lon,
  :osm_id,
  :osm_type,
  :place_id,
  keyword_init: true
)

# Request payload for Reverse#list.
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] boundingbox
#   @return [Array, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] lat
#   @return [String, nil]
#
# @!attribute [rw] licence
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
ReverseListMatch = Struct.new(
  :address,
  :boundingbox,
  :display_name,
  :lat,
  :licence,
  :lon,
  :osm_id,
  :osm_type,
  :place_id,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] boundingbox
#   @return [Array, nil]
#
# @!attribute [rw] class
#   @return [String, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] icon
#   @return [String, nil]
#
# @!attribute [rw] importance
#   @return [Float, nil]
#
# @!attribute [rw] lat
#   @return [String, nil]
#
# @!attribute [rw] licence
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
Search = Struct.new(
  :address,
  :boundingbox,
  :class,
  :display_name,
  :icon,
  :importance,
  :lat,
  :licence,
  :lon,
  :osm_id,
  :osm_type,
  :place_id,
  :type,
  keyword_init: true
)

# Request payload for Search#list.
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] boundingbox
#   @return [Array, nil]
#
# @!attribute [rw] class
#   @return [String, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] icon
#   @return [String, nil]
#
# @!attribute [rw] importance
#   @return [Float, nil]
#
# @!attribute [rw] lat
#   @return [String, nil]
#
# @!attribute [rw] licence
#   @return [String, nil]
#
# @!attribute [rw] lon
#   @return [String, nil]
#
# @!attribute [rw] osm_id
#   @return [Integer, nil]
#
# @!attribute [rw] osm_type
#   @return [String, nil]
#
# @!attribute [rw] place_id
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
SearchListMatch = Struct.new(
  :address,
  :boundingbox,
  :class,
  :display_name,
  :icon,
  :importance,
  :lat,
  :licence,
  :lon,
  :osm_id,
  :osm_type,
  :place_id,
  :type,
  keyword_init: true
)

# ServerStatus entity data model.
#
# @!attribute [rw] data_updated
#   @return [String, nil]
#
# @!attribute [rw] database_version
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] software_version
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [Integer, nil]
ServerStatus = Struct.new(
  :data_updated,
  :database_version,
  :message,
  :software_version,
  :status,
  keyword_init: true
)

# Request payload for ServerStatus#load.
#
# @!attribute [rw] data_updated
#   @return [String, nil]
#
# @!attribute [rw] database_version
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] software_version
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [Integer, nil]
ServerStatusLoadMatch = Struct.new(
  :data_updated,
  :database_version,
  :message,
  :software_version,
  :status,
  keyword_init: true
)

