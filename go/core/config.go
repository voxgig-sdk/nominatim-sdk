package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Nominatim",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://nominatim.openstreetmap.org",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"address_lookup": map[string]any{},
				"administrative": map[string]any{},
				"debug": map[string]any{},
				"reverse": map[string]any{},
				"search": map[string]any{},
				"server_status": map[string]any{},
			},
		},
		"entity": map[string]any{
			"address_lookup": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "address",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "boundingbox",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "class",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "display_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "importance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "lat",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "licence",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lon",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
				},
				"name": "address_lookup",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "accept_language",
											"orig": "accept_language",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "addressdetail",
											"orig": "addressdetail",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "extratag",
											"orig": "extratag",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "jsonv2",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "namedetail",
											"orig": "namedetail",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "osm_id",
											"orig": "osm_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_geojson",
											"orig": "polygon_geojson",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_kml",
											"orig": "polygon_kml",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_svg",
											"orig": "polygon_svg",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_text",
											"orig": "polygon_text",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/lookup",
								"parts": []any{
									"lookup",
								},
								"select": map[string]any{
									"exist": []any{
										"accept_language",
										"addressdetail",
										"extratag",
										"format",
										"namedetail",
										"osm_id",
										"polygon_geojson",
										"polygon_kml",
										"polygon_svg",
										"polygon_text",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"administrative": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "class",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country_code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "errormessage",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated",
						"type": "`$STRING`",
					},
				},
				"name": "administrative",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": 7,
											"kind": "query",
											"name": "day",
											"orig": "day",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "json",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/polygons",
								"parts": []any{
									"polygons",
								},
								"select": map[string]any{
									"exist": []any{
										"day",
										"format",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "json",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/deletable",
								"parts": []any{
									"deletable",
								},
								"select": map[string]any{
									"exist": []any{
										"format",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"debug": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "addresstags",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "admin_level",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "calculated_importance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "calculated_postcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "calculated_wikipedia",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "centroid",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "country_code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "extratags",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "geometry",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "housenumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "importance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "indexed_date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isarea",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "localname",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "names",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "osm_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parent_place_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "place_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "rank_address",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "rank_search",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
				},
				"name": "debug",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "addressdetail",
											"orig": "addressdetail",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "class",
											"orig": "class",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "html",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "group_hierarchy",
											"orig": "group_hierarchy",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "keyword",
											"orig": "keyword",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "osmid",
											"orig": "osmid",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "osmtype",
											"orig": "osmtype",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "place_id",
											"orig": "place_id",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_geojson",
											"orig": "polygon_geojson",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/details",
								"parts": []any{
									"details",
								},
								"select": map[string]any{
									"exist": []any{
										"addressdetail",
										"class",
										"format",
										"group_hierarchy",
										"keyword",
										"osmid",
										"osmtype",
										"place_id",
										"polygon_geojson",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"reverse": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "address",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "boundingbox",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "display_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lat",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "licence",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lon",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"type": "`$INTEGER`",
					},
				},
				"name": "reverse",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "accept_language",
											"orig": "accept_language",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "addressdetail",
											"orig": "addressdetail",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "extratag",
											"orig": "extratag",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "jsonv2",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "lat",
											"orig": "lat",
											"reqd": true,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "lon",
											"orig": "lon",
											"reqd": true,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "namedetail",
											"orig": "namedetail",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_geojson",
											"orig": "polygon_geojson",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_kml",
											"orig": "polygon_kml",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_svg",
											"orig": "polygon_svg",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_text",
											"orig": "polygon_text",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 18,
											"kind": "query",
											"name": "zoom",
											"orig": "zoom",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/reverse",
								"parts": []any{
									"reverse",
								},
								"select": map[string]any{
									"exist": []any{
										"accept_language",
										"addressdetail",
										"extratag",
										"format",
										"lat",
										"lon",
										"namedetail",
										"polygon_geojson",
										"polygon_kml",
										"polygon_svg",
										"polygon_text",
										"zoom",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "address",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "boundingbox",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "class",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "display_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "icon",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "importance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "lat",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "licence",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lon",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "accept_language",
											"orig": "accept_language",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "addressdetail",
											"orig": "addressdetail",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "bounded",
											"orig": "bounded",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "city",
											"orig": "city",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "country",
											"orig": "country",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "countrycode",
											"orig": "countrycode",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "county",
											"orig": "county",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "dedupe",
											"orig": "dedupe",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "extratag",
											"orig": "extratag",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "jsonv2",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 10,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "namedetail",
											"orig": "namedetail",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_geojson",
											"orig": "polygon_geojson",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_kml",
											"orig": "polygon_kml",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_svg",
											"orig": "polygon_svg",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "polygon_text",
											"orig": "polygon_text",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "postalcode",
											"orig": "postalcode",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "street",
											"orig": "street",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "viewbox",
											"orig": "viewbox",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/search",
								"parts": []any{
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"accept_language",
										"addressdetail",
										"bounded",
										"city",
										"country",
										"countrycode",
										"county",
										"dedupe",
										"extratag",
										"format",
										"limit",
										"namedetail",
										"polygon_geojson",
										"polygon_kml",
										"polygon_svg",
										"polygon_text",
										"postalcode",
										"q",
										"state",
										"street",
										"viewbox",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"server_status": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data_updated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "database_version",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "message",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "software_version",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$INTEGER`",
					},
				},
				"name": "server_status",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "text",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/status",
								"parts": []any{
									"status",
								},
								"select": map[string]any{
									"exist": []any{
										"format",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
