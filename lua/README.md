# Nominatim Lua SDK



The Lua SDK for the Nominatim API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:AddressLookup()` — each with the same small set of operations (`list`, `load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nominatim-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("nominatim_sdk")

local client = sdk.new()
```

### 2. List addresslookup records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local addresslookups, err = client:AddressLookup():list()
if err then error(err) end

for _, item in ipairs(addresslookups) do
  print(item["class"])
end
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local addresslookups, err = client:AddressLookup():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:AddressLookup():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NOMINATIM_TEST_LIVE=TRUE
```

Then run:

```bash
cd lua && busted test/
```


## Reference

### NominatimSDK

```lua
local sdk = require("nominatim_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NominatimSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
| `AddressLookup` | `(data) -> AddressLookupEntity` | Create an AddressLookup entity instance. |
| `Administrative` | `(data) -> AdministrativeEntity` | Create an Administrative entity instance. |
| `Debug` | `(data) -> DebugEntity` | Create a Debug entity instance. |
| `Reverse` | `(data) -> ReverseEntity` | Create a Reverse entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |
| `ServerStatus` | `(data) -> ServerStatusEntity` | Create a ServerStatus entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local address_lookup, err = client:AddressLookup():load()
    if err then error(err) end
    -- address_lookup is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### AddressLookup

| Field | Description |
| --- | --- |
| `address` |  |
| `boundingbox` |  |
| `class` |  |
| `display_name` |  |
| `importance` |  |
| `lat` |  |
| `licence` |  |
| `lon` |  |
| `osm_id` |  |
| `osm_type` |  |
| `place_id` |  |
| `type` |  |

Operations: List.

API path: `/lookup`

#### Administrative

| Field | Description |
| --- | --- |
| `class` |  |
| `country_code` |  |
| `errormessage` |  |
| `name` |  |
| `osm_id` |  |
| `osm_type` |  |
| `place_id` |  |
| `type` |  |
| `updated` |  |

Operations: List.

API path: `/polygons`

#### Debug

| Field | Description |
| --- | --- |
| `addresstag` |  |
| `admin_level` |  |
| `calculated_importance` |  |
| `calculated_postcode` |  |
| `calculated_wikipedia` |  |
| `category` |  |
| `centroid` |  |
| `country_code` |  |
| `extratag` |  |
| `geometry` |  |
| `housenumber` |  |
| `importance` |  |
| `indexed_date` |  |
| `isarea` |  |
| `localname` |  |
| `name` |  |
| `osm_id` |  |
| `osm_type` |  |
| `parent_place_id` |  |
| `place_id` |  |
| `rank_address` |  |
| `rank_search` |  |
| `type` |  |

Operations: Load.

API path: `/details`

#### Reverse

| Field | Description |
| --- | --- |
| `address` |  |
| `boundingbox` |  |
| `display_name` |  |
| `lat` |  |
| `licence` |  |
| `lon` |  |
| `osm_id` |  |
| `osm_type` |  |
| `place_id` |  |

Operations: List.

API path: `/reverse`

#### Search

| Field | Description |
| --- | --- |
| `address` |  |
| `boundingbox` |  |
| `class` |  |
| `display_name` |  |
| `icon` |  |
| `importance` |  |
| `lat` |  |
| `licence` |  |
| `lon` |  |
| `osm_id` |  |
| `osm_type` |  |
| `place_id` |  |
| `type` |  |

Operations: List.

API path: `/search`

#### ServerStatus

| Field | Description |
| --- | --- |
| `data_updated` |  |
| `database_version` |  |
| `message` |  |
| `software_version` |  |
| `status` |  |

Operations: Load.

API path: `/status`



## Entities


### AddressLookup

Create an instance: `local address_lookup = client:AddressLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `table` |  |
| `boundingbox` | `table` |  |
| `class` | `string` |  |
| `display_name` | `string` |  |
| `importance` | `number` |  |
| `lat` | `string` |  |
| `licence` | `string` |  |
| `lon` | `string` |  |
| `osm_id` | `number` |  |
| `osm_type` | `string` |  |
| `place_id` | `number` |  |
| `type` | `string` |  |

#### Example: List

```lua
local address_lookups, err = client:AddressLookup():list()
```


### Administrative

Create an instance: `local administrative = client:Administrative(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `class` | `string` |  |
| `country_code` | `string` |  |
| `errormessage` | `string` |  |
| `name` | `string` |  |
| `osm_id` | `number` |  |
| `osm_type` | `string` |  |
| `place_id` | `number` |  |
| `type` | `string` |  |
| `updated` | `string` |  |

#### Example: List

```lua
local administratives, err = client:Administrative():list()
```


### Debug

Create an instance: `local debug = client:Debug(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresstag` | `table` |  |
| `admin_level` | `number` |  |
| `calculated_importance` | `number` |  |
| `calculated_postcode` | `string` |  |
| `calculated_wikipedia` | `string` |  |
| `category` | `string` |  |
| `centroid` | `table` |  |
| `country_code` | `string` |  |
| `extratag` | `table` |  |
| `geometry` | `table` |  |
| `housenumber` | `string` |  |
| `importance` | `number` |  |
| `indexed_date` | `string` |  |
| `isarea` | `boolean` |  |
| `localname` | `string` |  |
| `name` | `table` |  |
| `osm_id` | `number` |  |
| `osm_type` | `string` |  |
| `parent_place_id` | `number` |  |
| `place_id` | `number` |  |
| `rank_address` | `number` |  |
| `rank_search` | `number` |  |
| `type` | `string` |  |

#### Example: Load

```lua
local debug, err = client:Debug():load()
```


### Reverse

Create an instance: `local reverse = client:Reverse(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `table` |  |
| `boundingbox` | `table` |  |
| `display_name` | `string` |  |
| `lat` | `string` |  |
| `licence` | `string` |  |
| `lon` | `string` |  |
| `osm_id` | `number` |  |
| `osm_type` | `string` |  |
| `place_id` | `number` |  |

#### Example: List

```lua
local reverses, err = client:Reverse():list()
```


### Search

Create an instance: `local search = client:Search(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `table` |  |
| `boundingbox` | `table` |  |
| `class` | `string` |  |
| `display_name` | `string` |  |
| `icon` | `string` |  |
| `importance` | `number` |  |
| `lat` | `string` |  |
| `licence` | `string` |  |
| `lon` | `string` |  |
| `osm_id` | `number` |  |
| `osm_type` | `string` |  |
| `place_id` | `number` |  |
| `type` | `string` |  |

#### Example: List

```lua
local searchs, err = client:Search():list()
```


### ServerStatus

Create an instance: `local server_status = client:ServerStatus(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data_updated` | `string` |  |
| `database_version` | `string` |  |
| `message` | `string` |  |
| `software_version` | `string` |  |
| `status` | `number` |  |

#### Example: Load

```lua
local server_status, err = client:ServerStatus():load()
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── nominatim_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`nominatim_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local addresslookup = client:AddressLookup()
addresslookup:list()

-- addresslookup:data_get() now returns the addresslookup data from the last list
-- addresslookup:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
