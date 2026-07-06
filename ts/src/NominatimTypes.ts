// Typed models for the Nominatim SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface AddressLookup {
  address?: Record<string, any>
  boundingbox?: any[]
  class?: string
  display_name?: string
  importance?: number
  lat?: string
  licence?: string
  lon?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
  type?: string
}

export interface AddressLookupListMatch {
  address?: Record<string, any>
  boundingbox?: any[]
  class?: string
  display_name?: string
  importance?: number
  lat?: string
  licence?: string
  lon?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
  type?: string
}

export interface Administrative {
  class?: string
  country_code?: string
  errormessage?: string
  name?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
  type?: string
  updated?: string
}

export interface AdministrativeListMatch {
  class?: string
  country_code?: string
  errormessage?: string
  name?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
  type?: string
  updated?: string
}

export interface Debug {
  addresstag?: Record<string, any>
  admin_level?: number
  calculated_importance?: number
  calculated_postcode?: string
  calculated_wikipedia?: string
  category?: string
  centroid?: Record<string, any>
  country_code?: string
  extratag?: Record<string, any>
  geometry?: Record<string, any>
  housenumber?: string
  importance?: number
  indexed_date?: string
  isarea?: boolean
  localname?: string
  name?: Record<string, any>
  osm_id?: number
  osm_type?: string
  parent_place_id?: number
  place_id?: number
  rank_address?: number
  rank_search?: number
  type?: string
}

export interface DebugLoadMatch {
  addresstag?: Record<string, any>
  admin_level?: number
  calculated_importance?: number
  calculated_postcode?: string
  calculated_wikipedia?: string
  category?: string
  centroid?: Record<string, any>
  country_code?: string
  extratag?: Record<string, any>
  geometry?: Record<string, any>
  housenumber?: string
  importance?: number
  indexed_date?: string
  isarea?: boolean
  localname?: string
  name?: Record<string, any>
  osm_id?: number
  osm_type?: string
  parent_place_id?: number
  place_id?: number
  rank_address?: number
  rank_search?: number
  type?: string
}

export interface Reverse {
  address?: Record<string, any>
  boundingbox?: any[]
  display_name?: string
  lat?: string
  licence?: string
  lon?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
}

export interface ReverseListMatch {
  address?: Record<string, any>
  boundingbox?: any[]
  display_name?: string
  lat?: string
  licence?: string
  lon?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
}

export interface Search {
  address?: Record<string, any>
  boundingbox?: any[]
  class?: string
  display_name?: string
  icon?: string
  importance?: number
  lat?: string
  licence?: string
  lon?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
  type?: string
}

export interface SearchListMatch {
  address?: Record<string, any>
  boundingbox?: any[]
  class?: string
  display_name?: string
  icon?: string
  importance?: number
  lat?: string
  licence?: string
  lon?: string
  osm_id?: number
  osm_type?: string
  place_id?: number
  type?: string
}

export interface ServerStatus {
  data_updated?: string
  database_version?: string
  message?: string
  software_version?: string
  status?: number
}

export interface ServerStatusLoadMatch {
  data_updated?: string
  database_version?: string
  message?: string
  software_version?: string
  status?: number
}

