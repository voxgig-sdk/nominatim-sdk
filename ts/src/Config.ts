
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'Nominatim',
        slug: "nominatim",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     test:     {
      "options": {
        "active": false
      },
      "transport": "base"
    },

  }


  options = {
    base: "https://nominatim.openstreetmap.org",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      address_lookup: {
      },

      administrative: {
      },

      debug: {
      },

      reverse: {
      },

      search: {
      },

      server_status: {
      },

    }
  }


  entity = {
    "address_lookup": {
      "fields": [
        {
          "name": "address",
          "short": "Address breakdown",
          "type": "`$OBJECT`"
        },
        {
          "name": "boundingbox",
          "short": "Bounding box coordinates",
          "type": "`$ARRAY`"
        },
        {
          "name": "class",
          "short": "Main OSM tag key",
          "type": "`$STRING`"
        },
        {
          "name": "display_name",
          "short": "Full comma-separated address",
          "type": "`$STRING`"
        },
        {
          "name": "importance",
          "short": "Computed importance rank",
          "type": "`$NUMBER`"
        },
        {
          "name": "lat",
          "short": "Latitude",
          "type": "`$STRING`"
        },
        {
          "name": "licence",
          "short": "License information",
          "type": "`$STRING`"
        },
        {
          "name": "lon",
          "short": "Longitude",
          "type": "`$STRING`"
        },
        {
          "name": "osm_id",
          "short": "OSM object ID",
          "type": "`$INTEGER`"
        },
        {
          "name": "osm_type",
          "short": "OSM type (node, way, relation)",
          "type": "`$STRING`"
        },
        {
          "name": "place_id",
          "short": "Unique identifier for the place",
          "type": "`$INTEGER`"
        },
        {
          "name": "type",
          "short": "Main OSM tag value",
          "type": "`$STRING`"
        }
      ],
      "name": "address_lookup",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "accept_language",
                    "orig": "accept_language",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "addressdetail",
                    "orig": "addressdetail",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "extratag",
                    "orig": "extratag",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": "jsonv2",
                    "kind": "query",
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "namedetail",
                    "orig": "namedetail",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "osm_id",
                    "orig": "osm_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_geojson",
                    "orig": "polygon_geojson",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_kml",
                    "orig": "polygon_kml",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_svg",
                    "orig": "polygon_svg",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_text",
                    "orig": "polygon_text",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/lookup",
              "parts": [
                "lookup"
              ],
              "select": {
                "exist": [
                  "accept_language",
                  "addressdetail",
                  "extratag",
                  "format",
                  "namedetail",
                  "osm_id",
                  "polygon_geojson",
                  "polygon_kml",
                  "polygon_svg",
                  "polygon_text"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "administrative": {
      "fields": [
        {
          "name": "class",
          "short": "Main OSM tag key",
          "type": "`$STRING`"
        },
        {
          "name": "country_code",
          "short": "Country code",
          "type": "`$STRING`"
        },
        {
          "name": "errormessage",
          "short": "Error message describing the polygon issue",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "short": "Name of the object",
          "type": "`$STRING`"
        },
        {
          "name": "osm_id",
          "short": "OSM object ID",
          "type": "`$INTEGER`"
        },
        {
          "name": "osm_type",
          "short": "OSM type (way, relation)",
          "type": "`$STRING`"
        },
        {
          "name": "place_id",
          "short": "Unique identifier for the place",
          "type": "`$INTEGER`"
        },
        {
          "name": "type",
          "short": "Main OSM tag value",
          "type": "`$STRING`"
        },
        {
          "name": "updated",
          "short": "Last update timestamp",
          "type": "`$STRING`"
        }
      ],
      "name": "administrative",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": 7,
                    "kind": "query",
                    "name": "day",
                    "orig": "day",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": "json",
                    "kind": "query",
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/polygons",
              "parts": [
                "polygons"
              ],
              "select": {
                "exist": [
                  "day",
                  "format"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            },
            {
              "args": {
                "query": [
                  {
                    "example": "json",
                    "kind": "query",
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/deletable",
              "parts": [
                "deletable"
              ],
              "select": {
                "exist": [
                  "format"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "debug": {
      "fields": [
        {
          "name": "addresstags",
          "short": "Address tags",
          "type": "`$OBJECT`"
        },
        {
          "name": "admin_level",
          "short": "Administrative level",
          "type": "`$INTEGER`"
        },
        {
          "name": "calculated_importance",
          "short": "Calculated importance",
          "type": "`$NUMBER`"
        },
        {
          "name": "calculated_postcode",
          "short": "Calculated postcode",
          "type": "`$STRING`"
        },
        {
          "name": "calculated_wikipedia",
          "short": "Wikipedia reference",
          "type": "`$STRING`"
        },
        {
          "name": "category",
          "short": "Main OSM tag key",
          "type": "`$STRING`"
        },
        {
          "name": "centroid",
          "short": "Centroid coordinates",
          "type": "`$OBJECT`"
        },
        {
          "name": "country_code",
          "short": "Country code",
          "type": "`$STRING`"
        },
        {
          "name": "extratags",
          "short": "Extra OSM tags",
          "type": "`$OBJECT`"
        },
        {
          "name": "geometry",
          "short": "Geometry information",
          "type": "`$OBJECT`"
        },
        {
          "name": "housenumber",
          "short": "House number",
          "type": "`$STRING`"
        },
        {
          "name": "importance",
          "short": "Computed importance rank",
          "type": "`$NUMBER`"
        },
        {
          "name": "indexed_date",
          "short": "Date when the object was indexed",
          "type": "`$STRING`"
        },
        {
          "name": "isarea",
          "short": "Whether the object is an area",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "localname",
          "short": "Local name",
          "type": "`$STRING`"
        },
        {
          "name": "names",
          "short": "All available names",
          "type": "`$OBJECT`"
        },
        {
          "name": "osm_id",
          "short": "OSM object ID",
          "type": "`$INTEGER`"
        },
        {
          "name": "osm_type",
          "short": "OSM type (node, way, relation)",
          "type": "`$STRING`"
        },
        {
          "name": "parent_place_id",
          "short": "Parent place ID",
          "type": "`$INTEGER`"
        },
        {
          "name": "place_id",
          "short": "Unique identifier for the place",
          "type": "`$INTEGER`"
        },
        {
          "name": "rank_address",
          "short": "Address rank",
          "type": "`$INTEGER`"
        },
        {
          "name": "rank_search",
          "short": "Search rank",
          "type": "`$INTEGER`"
        },
        {
          "name": "type",
          "short": "Main OSM tag value",
          "type": "`$STRING`"
        }
      ],
      "name": "debug",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "addressdetail",
                    "orig": "addressdetail",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "class",
                    "orig": "class",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "html",
                    "kind": "query",
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "group_hierarchy",
                    "orig": "group_hierarchy",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "keyword",
                    "orig": "keyword",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "osmid",
                    "orig": "osmid",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "osmtype",
                    "orig": "osmtype",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "place_id",
                    "orig": "place_id",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_geojson",
                    "orig": "polygon_geojson",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/details",
              "parts": [
                "details"
              ],
              "select": {
                "exist": [
                  "addressdetail",
                  "class",
                  "format",
                  "group_hierarchy",
                  "keyword",
                  "osmid",
                  "osmtype",
                  "place_id",
                  "polygon_geojson"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "reverse": {
      "fields": [
        {
          "name": "address",
          "short": "Address breakdown",
          "type": "`$OBJECT`"
        },
        {
          "name": "boundingbox",
          "short": "Bounding box coordinates",
          "type": "`$ARRAY`"
        },
        {
          "name": "display_name",
          "short": "Full comma-separated address",
          "type": "`$STRING`"
        },
        {
          "name": "lat",
          "short": "Latitude",
          "type": "`$STRING`"
        },
        {
          "name": "licence",
          "short": "License information",
          "type": "`$STRING`"
        },
        {
          "name": "lon",
          "short": "Longitude",
          "type": "`$STRING`"
        },
        {
          "name": "osm_id",
          "short": "OSM object ID",
          "type": "`$INTEGER`"
        },
        {
          "name": "osm_type",
          "short": "OSM type (node, way, relation)",
          "type": "`$STRING`"
        },
        {
          "name": "place_id",
          "short": "Unique identifier for the place",
          "type": "`$INTEGER`"
        }
      ],
      "name": "reverse",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "accept_language",
                    "orig": "accept_language",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "addressdetail",
                    "orig": "addressdetail",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "extratag",
                    "orig": "extratag",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": "jsonv2",
                    "kind": "query",
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "lat",
                    "orig": "lat",
                    "reqd": true,
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "lon",
                    "orig": "lon",
                    "reqd": true,
                    "type": "`$NUMBER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "namedetail",
                    "orig": "namedetail",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_geojson",
                    "orig": "polygon_geojson",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_kml",
                    "orig": "polygon_kml",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_svg",
                    "orig": "polygon_svg",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_text",
                    "orig": "polygon_text",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 18,
                    "kind": "query",
                    "name": "zoom",
                    "orig": "zoom",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/reverse",
              "parts": [
                "reverse"
              ],
              "select": {
                "exist": [
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
                  "zoom"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "search": {
      "fields": [
        {
          "name": "address",
          "short": "Address breakdown",
          "type": "`$OBJECT`"
        },
        {
          "name": "boundingbox",
          "short": "Bounding box coordinates",
          "type": "`$ARRAY`"
        },
        {
          "name": "class",
          "short": "Main OSM tag key",
          "type": "`$STRING`"
        },
        {
          "name": "display_name",
          "short": "Full comma-separated address",
          "type": "`$STRING`"
        },
        {
          "name": "icon",
          "short": "URL of icon representing the place",
          "type": "`$STRING`"
        },
        {
          "name": "importance",
          "short": "Computed importance rank",
          "type": "`$NUMBER`"
        },
        {
          "name": "lat",
          "short": "Latitude",
          "type": "`$STRING`"
        },
        {
          "name": "licence",
          "short": "License information",
          "type": "`$STRING`"
        },
        {
          "name": "lon",
          "short": "Longitude",
          "type": "`$STRING`"
        },
        {
          "name": "osm_id",
          "short": "OSM object ID",
          "type": "`$INTEGER`"
        },
        {
          "name": "osm_type",
          "short": "OSM type (node, way, relation)",
          "type": "`$STRING`"
        },
        {
          "name": "place_id",
          "short": "Unique identifier for the place",
          "type": "`$INTEGER`"
        },
        {
          "name": "type",
          "short": "Main OSM tag value",
          "type": "`$STRING`"
        }
      ],
      "name": "search",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "accept_language",
                    "orig": "accept_language",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "addressdetail",
                    "orig": "addressdetail",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "bounded",
                    "orig": "bounded",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "city",
                    "orig": "city",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "country",
                    "orig": "country",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "countrycode",
                    "orig": "countrycode",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "county",
                    "orig": "county",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "dedupe",
                    "orig": "dedupe",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "extratag",
                    "orig": "extratag",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": "jsonv2",
                    "kind": "query",
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "namedetail",
                    "orig": "namedetail",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_geojson",
                    "orig": "polygon_geojson",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_kml",
                    "orig": "polygon_kml",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_svg",
                    "orig": "polygon_svg",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "polygon_text",
                    "orig": "polygon_text",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "postalcode",
                    "orig": "postalcode",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "state",
                    "orig": "state",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "street",
                    "orig": "street",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "viewbox",
                    "orig": "viewbox",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/search",
              "parts": [
                "search"
              ],
              "select": {
                "exist": [
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
                  "viewbox"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "server_status": {
      "fields": [
        {
          "name": "data_updated",
          "short": "Timestamp when the database was last updated",
          "type": "`$STRING`"
        },
        {
          "name": "database_version",
          "short": "Database version",
          "type": "`$STRING`"
        },
        {
          "name": "message",
          "short": "Status message",
          "type": "`$STRING`"
        },
        {
          "name": "software_version",
          "short": "Nominatim software version",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "short": "Status code (0 = OK)",
          "type": "`$INTEGER`"
        }
      ],
      "name": "server_status",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "text",
                    "kind": "query",
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/status",
              "parts": [
                "status"
              ],
              "select": {
                "exist": [
                  "format"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config
}

