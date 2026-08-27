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
			"slug": "nominatim",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
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
						"short": "Address breakdown",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "boundingbox",
						"short": "Bounding box coordinates",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "class",
						"short": "Main OSM tag key",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "display_name",
						"short": "Full comma-separated address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "importance",
						"short": "Computed importance rank",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "lat",
						"short": "Latitude",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "licence",
						"short": "License information",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lon",
						"short": "Longitude",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"short": "OSM object ID",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"short": "OSM type (node, way, relation)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"short": "Unique identifier for the place",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"short": "Main OSM tag value",
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
						"short": "Main OSM tag key",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country_code",
						"short": "Country code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "errormessage",
						"short": "Error message describing the polygon issue",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "Name of the object",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"short": "OSM object ID",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"short": "OSM type (way, relation)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"short": "Unique identifier for the place",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"short": "Main OSM tag value",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated",
						"short": "Last update timestamp",
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
						"short": "Address tags",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "admin_level",
						"short": "Administrative level",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "calculated_importance",
						"short": "Calculated importance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "calculated_postcode",
						"short": "Calculated postcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "calculated_wikipedia",
						"short": "Wikipedia reference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "category",
						"short": "Main OSM tag key",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "centroid",
						"short": "Centroid coordinates",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "country_code",
						"short": "Country code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "extratags",
						"short": "Extra OSM tags",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "geometry",
						"short": "Geometry information",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "housenumber",
						"short": "House number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "importance",
						"short": "Computed importance rank",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "indexed_date",
						"short": "Date when the object was indexed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isarea",
						"short": "Whether the object is an area",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "localname",
						"short": "Local name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "names",
						"short": "All available names",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "osm_id",
						"short": "OSM object ID",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"short": "OSM type (node, way, relation)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parent_place_id",
						"short": "Parent place ID",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "place_id",
						"short": "Unique identifier for the place",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "rank_address",
						"short": "Address rank",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "rank_search",
						"short": "Search rank",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"short": "Main OSM tag value",
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
						"short": "Address breakdown",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "boundingbox",
						"short": "Bounding box coordinates",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "display_name",
						"short": "Full comma-separated address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lat",
						"short": "Latitude",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "licence",
						"short": "License information",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lon",
						"short": "Longitude",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"short": "OSM object ID",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"short": "OSM type (node, way, relation)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"short": "Unique identifier for the place",
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
						"short": "Address breakdown",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "boundingbox",
						"short": "Bounding box coordinates",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "class",
						"short": "Main OSM tag key",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "display_name",
						"short": "Full comma-separated address",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "icon",
						"short": "URL of icon representing the place",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "importance",
						"short": "Computed importance rank",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "lat",
						"short": "Latitude",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "licence",
						"short": "License information",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lon",
						"short": "Longitude",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "osm_id",
						"short": "OSM object ID",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "osm_type",
						"short": "OSM type (node, way, relation)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "place_id",
						"short": "Unique identifier for the place",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"short": "Main OSM tag value",
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
						"short": "Timestamp when the database was last updated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "database_version",
						"short": "Database version",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "message",
						"short": "Status message",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "software_version",
						"short": "Nominatim software version",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "Status code (0 = OK)",
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
