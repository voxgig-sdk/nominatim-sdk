# Nominatim Ruby SDK Reference

Complete API reference for the Nominatim Ruby SDK.


## NominatimSDK

### Constructor

```ruby
require_relative 'Nominatim_sdk'

client = NominatimSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NominatimSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = NominatimSDK.test
```


### Instance Methods

#### `AddressLookup(data = nil)`

Create a new `AddressLookup` entity instance. Pass `nil` for no initial data.

#### `Administrative(data = nil)`

Create a new `Administrative` entity instance. Pass `nil` for no initial data.

#### `Debug(data = nil)`

Create a new `Debug` entity instance. Pass `nil` for no initial data.

#### `Reverse(data = nil)`

Create a new `Reverse` entity instance. Pass `nil` for no initial data.

#### `Search(data = nil)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `ServerStatus(data = nil)`

Create a new `ServerStatus` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AddressLookupEntity

```ruby
address_lookup = client.AddressLookup
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Hash` | No | Address breakdown |
| `boundingbox` | `Array` | No | Bounding box coordinates |
| `class` | `String` | No | Main OSM tag key |
| `display_name` | `String` | No | Full comma-separated address |
| `importance` | `Float` | No | Computed importance rank |
| `lat` | `String` | No | Latitude |
| `licence` | `String` | No | License information |
| `lon` | `String` | No | Longitude |
| `osm_id` | `Integer` | No | OSM object ID |
| `osm_type` | `String` | No | OSM type (node, way, relation) |
| `place_id` | `Integer` | No | Unique identifier for the place |
| `type` | `String` | No | Main OSM tag value |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.AddressLookup.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## AdministrativeEntity

```ruby
administrative = client.Administrative
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `class` | `String` | No | Main OSM tag key |
| `country_code` | `String` | No | Country code |
| `errormessage` | `String` | No | Error message describing the polygon issue |
| `name` | `String` | No | Name of the object |
| `osm_id` | `Integer` | No | OSM object ID |
| `osm_type` | `String` | No | OSM type (way, relation) |
| `place_id` | `Integer` | No | Unique identifier for the place |
| `type` | `String` | No | Main OSM tag value |
| `updated` | `String` | No | Last update timestamp |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Administrative.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DebugEntity

```ruby
debug = client.Debug
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstags` | `Hash` | No | Address tags |
| `admin_level` | `Integer` | No | Administrative level |
| `calculated_importance` | `Float` | No | Calculated importance |
| `calculated_postcode` | `String` | No | Calculated postcode |
| `calculated_wikipedia` | `String` | No | Wikipedia reference |
| `category` | `String` | No | Main OSM tag key |
| `centroid` | `Hash` | No | Centroid coordinates |
| `country_code` | `String` | No | Country code |
| `extratags` | `Hash` | No | Extra OSM tags |
| `geometry` | `Hash` | No | Geometry information |
| `housenumber` | `String` | No | House number |
| `importance` | `Float` | No | Computed importance rank |
| `indexed_date` | `String` | No | Date when the object was indexed |
| `isarea` | `Boolean` | No | Whether the object is an area |
| `localname` | `String` | No | Local name |
| `names` | `Hash` | No | All available names |
| `osm_id` | `Integer` | No | OSM object ID |
| `osm_type` | `String` | No | OSM type (node, way, relation) |
| `parent_place_id` | `Integer` | No | Parent place ID |
| `place_id` | `Integer` | No | Unique identifier for the place |
| `rank_address` | `Integer` | No | Address rank |
| `rank_search` | `Integer` | No | Search rank |
| `type` | `String` | No | Main OSM tag value |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Debug.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DebugEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ReverseEntity

```ruby
reverse = client.Reverse
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Hash` | No | Address breakdown |
| `boundingbox` | `Array` | No | Bounding box coordinates |
| `display_name` | `String` | No | Full comma-separated address |
| `lat` | `String` | No | Latitude |
| `licence` | `String` | No | License information |
| `lon` | `String` | No | Longitude |
| `osm_id` | `Integer` | No | OSM object ID |
| `osm_type` | `String` | No | OSM type (node, way, relation) |
| `place_id` | `Integer` | No | Unique identifier for the place |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Reverse.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ReverseEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SearchEntity

```ruby
search = client.Search
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Hash` | No | Address breakdown |
| `boundingbox` | `Array` | No | Bounding box coordinates |
| `class` | `String` | No | Main OSM tag key |
| `display_name` | `String` | No | Full comma-separated address |
| `icon` | `String` | No | URL of icon representing the place |
| `importance` | `Float` | No | Computed importance rank |
| `lat` | `String` | No | Latitude |
| `licence` | `String` | No | License information |
| `lon` | `String` | No | Longitude |
| `osm_id` | `Integer` | No | OSM object ID |
| `osm_type` | `String` | No | OSM type (node, way, relation) |
| `place_id` | `Integer` | No | Unique identifier for the place |
| `type` | `String` | No | Main OSM tag value |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Search.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ServerStatusEntity

```ruby
server_status = client.ServerStatus
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data_updated` | `String` | No | Timestamp when the database was last updated |
| `database_version` | `String` | No | Database version |
| `message` | `String` | No | Status message |
| `software_version` | `String` | No | Nominatim software version |
| `status` | `Integer` | No | Status code (0 = OK) |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.ServerStatus.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = NominatimSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

