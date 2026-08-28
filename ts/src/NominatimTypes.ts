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
  accept_language?: string
  addressdetail?: number
  extratag?: number
  format?: string
  namedetail?: number
  osm_id: string
  polygon_geojson?: number
  polygon_kml?: number
  polygon_svg?: number
  polygon_text?: number
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
  day?: number
  format?: string
}

export interface Debug {
  addresstags?: Record<string, any>
  admin_level?: number
  calculated_importance?: number
  calculated_postcode?: string
  calculated_wikipedia?: string
  category?: string
  centroid?: Record<string, any>
  country_code?: string
  extratags?: Record<string, any>
  geometry?: Record<string, any>
  housenumber?: string
  importance?: number
  indexed_date?: string
  isarea?: boolean
  localname?: string
  names?: Record<string, any>
  osm_id?: number
  osm_type?: string
  parent_place_id?: number
  place_id?: number
  rank_address?: number
  rank_search?: number
  type?: string
}

export interface DebugLoadMatch {
  addressdetail?: number
  class?: string
  format?: string
  group_hierarchy?: number
  keyword?: number
  osmid?: number
  osmtype?: string
  place_id?: number
  polygon_geojson?: number
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
  accept_language?: string
  addressdetail?: number
  extratag?: number
  format?: string
  lat: number
  lon: number
  namedetail?: number
  polygon_geojson?: number
  polygon_kml?: number
  polygon_svg?: number
  polygon_text?: number
  zoom?: number
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
  accept_language?: string
  addressdetail?: number
  bounded?: number
  city?: string
  country?: string
  countrycode?: string
  county?: string
  dedupe?: number
  extratag?: number
  format?: string
  limit?: number
  namedetail?: number
  polygon_geojson?: number
  polygon_kml?: number
  polygon_svg?: number
  polygon_text?: number
  postalcode?: string
  q?: string
  state?: string
  street?: string
  viewbox?: string
}

export interface ServerStatus {
  data_updated?: string
  database_version?: string
  message?: string
  software_version?: string
  status?: number
}

export interface ServerStatusLoadMatch {
  format?: string
}

