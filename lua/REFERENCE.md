# Nominatim Lua SDK Reference

Complete API reference for the Nominatim Lua SDK.


## NominatimSDK

### Constructor

```lua
local sdk = require("nominatim_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `AddressLookup(data)`

Create a new `AddressLookup` entity instance. Pass `nil` for no initial data.

#### `Administrative(data)`

Create a new `Administrative` entity instance. Pass `nil` for no initial data.

#### `Debug(data)`

Create a new `Debug` entity instance. Pass `nil` for no initial data.

#### `Reverse(data)`

Create a new `Reverse` entity instance. Pass `nil` for no initial data.

#### `Search(data)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `ServerStatus(data)`

Create a new `ServerStatus` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AddressLookupEntity

```lua
local address_lookup = client:AddressLookup(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `table` | No | Address breakdown |
| `boundingbox` | `table` | No | Bounding box coordinates |
| `class` | `string` | No | Main OSM tag key |
| `display_name` | `string` | No | Full comma-separated address |
| `importance` | `number` | No | Computed importance rank |
| `lat` | `string` | No | Latitude |
| `licence` | `string` | No | License information |
| `lon` | `string` | No | Longitude |
| `osm_id` | `number` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `place_id` | `number` | No | Unique identifier for the place |
| `type` | `string` | No | Main OSM tag value |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:AddressLookup():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## AdministrativeEntity

```lua
local administrative = client:Administrative(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `class` | `string` | No | Main OSM tag key |
| `country_code` | `string` | No | Country code |
| `errormessage` | `string` | No | Error message describing the polygon issue |
| `name` | `string` | No | Name of the object |
| `osm_id` | `number` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (way, relation) |
| `place_id` | `number` | No | Unique identifier for the place |
| `type` | `string` | No | Main OSM tag value |
| `updated` | `string` | No | Last update timestamp |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Administrative():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DebugEntity

```lua
local debug = client:Debug(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstags` | `table` | No | Address tags |
| `admin_level` | `number` | No | Administrative level |
| `calculated_importance` | `number` | No | Calculated importance |
| `calculated_postcode` | `string` | No | Calculated postcode |
| `calculated_wikipedia` | `string` | No | Wikipedia reference |
| `category` | `string` | No | Main OSM tag key |
| `centroid` | `table` | No | Centroid coordinates |
| `country_code` | `string` | No | Country code |
| `extratags` | `table` | No | Extra OSM tags |
| `geometry` | `table` | No | Geometry information |
| `housenumber` | `string` | No | House number |
| `importance` | `number` | No | Computed importance rank |
| `indexed_date` | `string` | No | Date when the object was indexed |
| `isarea` | `boolean` | No | Whether the object is an area |
| `localname` | `string` | No | Local name |
| `names` | `table` | No | All available names |
| `osm_id` | `number` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `parent_place_id` | `number` | No | Parent place ID |
| `place_id` | `number` | No | Unique identifier for the place |
| `rank_address` | `number` | No | Address rank |
| `rank_search` | `number` | No | Search rank |
| `type` | `string` | No | Main OSM tag value |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Debug():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DebugEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ReverseEntity

```lua
local reverse = client:Reverse(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `table` | No | Address breakdown |
| `boundingbox` | `table` | No | Bounding box coordinates |
| `display_name` | `string` | No | Full comma-separated address |
| `lat` | `string` | No | Latitude |
| `licence` | `string` | No | License information |
| `lon` | `string` | No | Longitude |
| `osm_id` | `number` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `place_id` | `number` | No | Unique identifier for the place |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Reverse():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReverseEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SearchEntity

```lua
local search = client:Search(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `table` | No | Address breakdown |
| `boundingbox` | `table` | No | Bounding box coordinates |
| `class` | `string` | No | Main OSM tag key |
| `display_name` | `string` | No | Full comma-separated address |
| `icon` | `string` | No | URL of icon representing the place |
| `importance` | `number` | No | Computed importance rank |
| `lat` | `string` | No | Latitude |
| `licence` | `string` | No | License information |
| `lon` | `string` | No | Longitude |
| `osm_id` | `number` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `place_id` | `number` | No | Unique identifier for the place |
| `type` | `string` | No | Main OSM tag value |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ServerStatusEntity

```lua
local server_status = client:ServerStatus(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data_updated` | `string` | No | Timestamp when the database was last updated |
| `database_version` | `string` | No | Database version |
| `message` | `string` | No | Status message |
| `software_version` | `string` | No | Nominatim software version |
| `status` | `number` | No | Status code (0 = OK) |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ServerStatus():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
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

