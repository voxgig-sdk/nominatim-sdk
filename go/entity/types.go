// Typed models for the Nominatim SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/nominatim-sdk/go/core"
)

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

// AddressLookupListMatch is the typed request payload for AddressLookup.ListTyped.
type AddressLookupListMatch struct {
	AcceptLanguage *string `json:"accept_language,omitempty"`
	Addressdetail *int `json:"addressdetail,omitempty"`
	Extratag *int `json:"extratag,omitempty"`
	Format *string `json:"format,omitempty"`
	Namedetail *int `json:"namedetail,omitempty"`
	OsmId string `json:"osm_id"`
	PolygonGeojson *int `json:"polygon_geojson,omitempty"`
	PolygonKml *int `json:"polygon_kml,omitempty"`
	PolygonSvg *int `json:"polygon_svg,omitempty"`
	PolygonText *int `json:"polygon_text,omitempty"`
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

// AdministrativeListMatch is the typed request payload for Administrative.ListTyped.
type AdministrativeListMatch struct {
	Day *int `json:"day,omitempty"`
	Format *string `json:"format,omitempty"`
}

// Debug is the typed data model for the debug entity.
type Debug struct {
	Addresstags *map[string]any `json:"addresstags,omitempty"`
	AdminLevel *int `json:"admin_level,omitempty"`
	CalculatedImportance *float64 `json:"calculated_importance,omitempty"`
	CalculatedPostcode *string `json:"calculated_postcode,omitempty"`
	CalculatedWikipedia *string `json:"calculated_wikipedia,omitempty"`
	Category *string `json:"category,omitempty"`
	Centroid *map[string]any `json:"centroid,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Extratags *map[string]any `json:"extratags,omitempty"`
	Geometry *map[string]any `json:"geometry,omitempty"`
	Housenumber *string `json:"housenumber,omitempty"`
	Importance *float64 `json:"importance,omitempty"`
	IndexedDate *string `json:"indexed_date,omitempty"`
	Isarea *bool `json:"isarea,omitempty"`
	Localname *string `json:"localname,omitempty"`
	Names *map[string]any `json:"names,omitempty"`
	OsmId *int `json:"osm_id,omitempty"`
	OsmType *string `json:"osm_type,omitempty"`
	ParentPlaceId *int `json:"parent_place_id,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	RankAddress *int `json:"rank_address,omitempty"`
	RankSearch *int `json:"rank_search,omitempty"`
	Type *string `json:"type,omitempty"`
}

// DebugLoadMatch is the typed request payload for Debug.LoadTyped.
type DebugLoadMatch struct {
	Addressdetail *int `json:"addressdetail,omitempty"`
	Class *string `json:"class,omitempty"`
	Format *string `json:"format,omitempty"`
	GroupHierarchy *int `json:"group_hierarchy,omitempty"`
	Keyword *int `json:"keyword,omitempty"`
	Osmid *int `json:"osmid,omitempty"`
	Osmtype *string `json:"osmtype,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	PolygonGeojson *int `json:"polygon_geojson,omitempty"`
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

// ReverseListMatch is the typed request payload for Reverse.ListTyped.
type ReverseListMatch struct {
	AcceptLanguage *string `json:"accept_language,omitempty"`
	Addressdetail *int `json:"addressdetail,omitempty"`
	Extratag *int `json:"extratag,omitempty"`
	Format *string `json:"format,omitempty"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Namedetail *int `json:"namedetail,omitempty"`
	PolygonGeojson *int `json:"polygon_geojson,omitempty"`
	PolygonKml *int `json:"polygon_kml,omitempty"`
	PolygonSvg *int `json:"polygon_svg,omitempty"`
	PolygonText *int `json:"polygon_text,omitempty"`
	Zoom *int `json:"zoom,omitempty"`
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

// SearchListMatch is the typed request payload for Search.ListTyped.
type SearchListMatch struct {
	AcceptLanguage *string `json:"accept_language,omitempty"`
	Addressdetail *int `json:"addressdetail,omitempty"`
	Bounded *int `json:"bounded,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	Countrycode *string `json:"countrycode,omitempty"`
	County *string `json:"county,omitempty"`
	Dedupe *int `json:"dedupe,omitempty"`
	Extratag *int `json:"extratag,omitempty"`
	Format *string `json:"format,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Namedetail *int `json:"namedetail,omitempty"`
	PolygonGeojson *int `json:"polygon_geojson,omitempty"`
	PolygonKml *int `json:"polygon_kml,omitempty"`
	PolygonSvg *int `json:"polygon_svg,omitempty"`
	PolygonText *int `json:"polygon_text,omitempty"`
	Postalcode *string `json:"postalcode,omitempty"`
	Q *string `json:"q,omitempty"`
	State *string `json:"state,omitempty"`
	Street *string `json:"street,omitempty"`
	Viewbox *string `json:"viewbox,omitempty"`
}

// ServerStatus is the typed data model for the server_status entity.
type ServerStatus struct {
	DataUpdated *string `json:"data_updated,omitempty"`
	DatabaseVersion *string `json:"database_version,omitempty"`
	Message *string `json:"message,omitempty"`
	SoftwareVersion *string `json:"software_version,omitempty"`
	Status *int `json:"status,omitempty"`
}

// ServerStatusLoadMatch is the typed request payload for ServerStatus.LoadTyped.
type ServerStatusLoadMatch struct {
	Format *string `json:"format,omitempty"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
