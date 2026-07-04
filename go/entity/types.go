// Typed models for the Nominatim SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// AddressLookup is the typed data model for the address_lookup entity.
type AddressLookup struct {
	Address *map[string]any `json:"address,omitempty"`
	Boundingbox *[]any `json:"boundingbox,omitempty"`
	Class *string `json:"class,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Importance *float64 `json:"importance,omitempty"`
	Lat *string `json:"lat,omitempty"`
	Licence *string `json:"licence,omitempty"`
	Lon *string `json:"lon,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// AddressLookupListMatch mirrors the address_lookup fields as an all-optional match
// filter (Go analog of Partial<AddressLookup>).
type AddressLookupListMatch struct {
	Address *map[string]any `json:"address,omitempty"`
	Boundingbox *[]any `json:"boundingbox,omitempty"`
	Class *string `json:"class,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Importance *float64 `json:"importance,omitempty"`
	Lat *string `json:"lat,omitempty"`
	Licence *string `json:"licence,omitempty"`
	Lon *string `json:"lon,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Administrative is the typed data model for the administrative entity.
type Administrative struct {
	Class *string `json:"class,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Errormessage *string `json:"errormessage,omitempty"`
	Name *string `json:"name,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Updated *string `json:"updated,omitempty"`
}

// AdministrativeListMatch mirrors the administrative fields as an all-optional match
// filter (Go analog of Partial<Administrative>).
type AdministrativeListMatch struct {
	Class *string `json:"class,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Errormessage *string `json:"errormessage,omitempty"`
	Name *string `json:"name,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Updated *string `json:"updated,omitempty"`
}

// Debug is the typed data model for the debug entity.
type Debug struct {
	Addresstag *map[string]any `json:"addresstag,omitempty"`
	AdminLevel *int `json:"admin_level,omitempty"`
	CalculatedImportance *float64 `json:"calculated_importance,omitempty"`
	CalculatedPostcode *string `json:"calculated_postcode,omitempty"`
	CalculatedWikipedia *string `json:"calculated_wikipedia,omitempty"`
	Category *string `json:"category,omitempty"`
	Centroid *map[string]any `json:"centroid,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Extratag *map[string]any `json:"extratag,omitempty"`
	Geometry *map[string]any `json:"geometry,omitempty"`
	Housenumber *string `json:"housenumber,omitempty"`
	Importance *float64 `json:"importance,omitempty"`
	IndexedDate *string `json:"indexed_date,omitempty"`
	Isarea *bool `json:"isarea,omitempty"`
	Localname *string `json:"localname,omitempty"`
	Name *map[string]any `json:"name,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	ParentPlaceId *int `json:"parent_place_id,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	RankAddress *int `json:"rank_address,omitempty"`
	RankSearch *int `json:"rank_search,omitempty"`
	Type *string `json:"type,omitempty"`
}

// DebugLoadMatch mirrors the debug fields as an all-optional match
// filter (Go analog of Partial<Debug>).
type DebugLoadMatch struct {
	Addresstag *map[string]any `json:"addresstag,omitempty"`
	AdminLevel *int `json:"admin_level,omitempty"`
	CalculatedImportance *float64 `json:"calculated_importance,omitempty"`
	CalculatedPostcode *string `json:"calculated_postcode,omitempty"`
	CalculatedWikipedia *string `json:"calculated_wikipedia,omitempty"`
	Category *string `json:"category,omitempty"`
	Centroid *map[string]any `json:"centroid,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Extratag *map[string]any `json:"extratag,omitempty"`
	Geometry *map[string]any `json:"geometry,omitempty"`
	Housenumber *string `json:"housenumber,omitempty"`
	Importance *float64 `json:"importance,omitempty"`
	IndexedDate *string `json:"indexed_date,omitempty"`
	Isarea *bool `json:"isarea,omitempty"`
	Localname *string `json:"localname,omitempty"`
	Name *map[string]any `json:"name,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	ParentPlaceId *int `json:"parent_place_id,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	RankAddress *int `json:"rank_address,omitempty"`
	RankSearch *int `json:"rank_search,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Reverse is the typed data model for the reverse entity.
type Reverse struct {
	Address *map[string]any `json:"address,omitempty"`
	Boundingbox *[]any `json:"boundingbox,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Lat *string `json:"lat,omitempty"`
	Licence *string `json:"licence,omitempty"`
	Lon *string `json:"lon,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
}

// ReverseListMatch mirrors the reverse fields as an all-optional match
// filter (Go analog of Partial<Reverse>).
type ReverseListMatch struct {
	Address *map[string]any `json:"address,omitempty"`
	Boundingbox *[]any `json:"boundingbox,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Lat *string `json:"lat,omitempty"`
	Licence *string `json:"licence,omitempty"`
	Lon *string `json:"lon,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
}

// Search is the typed data model for the search entity.
type Search struct {
	Address *map[string]any `json:"address,omitempty"`
	Boundingbox *[]any `json:"boundingbox,omitempty"`
	Class *string `json:"class,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Importance *float64 `json:"importance,omitempty"`
	Lat *string `json:"lat,omitempty"`
	Licence *string `json:"licence,omitempty"`
	Lon *string `json:"lon,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// SearchListMatch mirrors the search fields as an all-optional match
// filter (Go analog of Partial<Search>).
type SearchListMatch struct {
	Address *map[string]any `json:"address,omitempty"`
	Boundingbox *[]any `json:"boundingbox,omitempty"`
	Class *string `json:"class,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Importance *float64 `json:"importance,omitempty"`
	Lat *string `json:"lat,omitempty"`
	Licence *string `json:"licence,omitempty"`
	Lon *string `json:"lon,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ServerStatus is the typed data model for the server_status entity.
type ServerStatus struct {
	DataUpdated *string `json:"data_updated,omitempty"`
	DatabaseVersion *string `json:"database_version,omitempty"`
	Message *string `json:"message,omitempty"`
	SoftwareVersion *string `json:"software_version,omitempty"`
	Status *int `json:"status,omitempty"`
}

// ServerStatusLoadMatch mirrors the server_status fields as an all-optional match
// filter (Go analog of Partial<ServerStatus>).
type ServerStatusLoadMatch struct {
	DataUpdated *string `json:"data_updated,omitempty"`
	DatabaseVersion *string `json:"database_version,omitempty"`
	Message *string `json:"message,omitempty"`
	SoftwareVersion *string `json:"software_version,omitempty"`
	Status *int `json:"status,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
