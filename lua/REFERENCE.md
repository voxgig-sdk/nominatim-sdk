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
| `address` | `table` | No |  |
| `boundingbox` | `table` | No |  |
| `class` | `string` | No |  |
| `display_name` | `string` | No |  |
| `importance` | `number` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `number` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `number` | No |  |
| `type` | `string` | No |  |

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
| `class` | `string` | No |  |
| `country_code` | `string` | No |  |
| `errormessage` | `string` | No |  |
| `name` | `string` | No |  |
| `osm_id` | `number` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `number` | No |  |
| `type` | `string` | No |  |
| `updated` | `string` | No |  |

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
| `addresstag` | `table` | No |  |
| `admin_level` | `number` | No |  |
| `calculated_importance` | `number` | No |  |
| `calculated_postcode` | `string` | No |  |
| `calculated_wikipedia` | `string` | No |  |
| `category` | `string` | No |  |
| `centroid` | `table` | No |  |
| `country_code` | `string` | No |  |
| `extratag` | `table` | No |  |
| `geometry` | `table` | No |  |
| `housenumber` | `string` | No |  |
| `importance` | `number` | No |  |
| `indexed_date` | `string` | No |  |
| `isarea` | `boolean` | No |  |
| `localname` | `string` | No |  |
| `name` | `table` | No |  |
| `osm_id` | `number` | No |  |
| `osm_type` | `string` | No |  |
| `parent_place_id` | `number` | No |  |
| `place_id` | `number` | No |  |
| `rank_address` | `number` | No |  |
| `rank_search` | `number` | No |  |
| `type` | `string` | No |  |

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
| `address` | `table` | No |  |
| `boundingbox` | `table` | No |  |
| `display_name` | `string` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `number` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `number` | No |  |

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
| `address` | `table` | No |  |
| `boundingbox` | `table` | No |  |
| `class` | `string` | No |  |
| `display_name` | `string` | No |  |
| `icon` | `string` | No |  |
| `importance` | `number` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `number` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `number` | No |  |
| `type` | `string` | No |  |

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
| `data_updated` | `string` | No |  |
| `database_version` | `string` | No |  |
| `message` | `string` | No |  |
| `software_version` | `string` | No |  |
| `status` | `number` | No |  |

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

