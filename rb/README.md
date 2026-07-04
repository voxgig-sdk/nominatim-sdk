# Nominatim Ruby SDK



The Ruby SDK for the Nominatim API — an entity-oriented client using idiomatic Ruby conventions.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nominatim-sdk/releases](https://github.com/voxgig-sdk/nominatim-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Nominatim_sdk"

client = NominatimSDK.new
```

### 2. List addresslookups

```ruby
begin
  result = client.addresslookup.list
  if result.is_a?(Array)
    result.each do |item|
      d = item.data_get
      puts "#{d["id"]} #{d["name"]}"
    end
  end
rescue => err
  warn "list failed: #{err}"
end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  warn result["err"]
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = NominatimSDK.test

result = client.addresslookup.load({ "id" => "test01" })
# result contains mock response data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = NominatimSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### NominatimSDK

```ruby
require_relative "Nominatim_sdk"
client = NominatimSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = NominatimSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NominatimSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `AddressLookup` | `(data) -> AddressLookupEntity` | Create a AddressLookup entity instance. |
| `Administrative` | `(data) -> AdministrativeEntity` | Create a Administrative entity instance. |
| `Debug` | `(data) -> DebugEntity` | Create a Debug entity instance. |
| `Reverse` | `(data) -> ReverseEntity` | Create a Reverse entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |
| `ServerStatus` | `(data) -> ServerStatusEntity` | Create a ServerStatus entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> Array` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `NominatimError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

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

Create an instance: `const address_lookup = client.address_lookup`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$OBJECT`` |  |
| `boundingbox` | ``$ARRAY`` |  |
| `class` | ``$STRING`` |  |
| `display_name` | ``$STRING`` |  |
| `importance` | ``$NUMBER`` |  |
| `lat` | ``$STRING`` |  |
| `licence` | ``$STRING`` |  |
| `lon` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |

#### Example: List

```ts
const address_lookups = await client.address_lookup.list()
```


### Administrative

Create an instance: `const administrative = client.administrative`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `class` | ``$STRING`` |  |
| `country_code` | ``$STRING`` |  |
| `errormessage` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |
| `updated` | ``$STRING`` |  |

#### Example: List

```ts
const administratives = await client.administrative.list()
```


### Debug

Create an instance: `const debug = client.debug`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresstag` | ``$OBJECT`` |  |
| `admin_level` | ``$INTEGER`` |  |
| `calculated_importance` | ``$NUMBER`` |  |
| `calculated_postcode` | ``$STRING`` |  |
| `calculated_wikipedia` | ``$STRING`` |  |
| `category` | ``$STRING`` |  |
| `centroid` | ``$OBJECT`` |  |
| `country_code` | ``$STRING`` |  |
| `extratag` | ``$OBJECT`` |  |
| `geometry` | ``$OBJECT`` |  |
| `housenumber` | ``$STRING`` |  |
| `importance` | ``$NUMBER`` |  |
| `indexed_date` | ``$STRING`` |  |
| `isarea` | ``$BOOLEAN`` |  |
| `localname` | ``$STRING`` |  |
| `name` | ``$OBJECT`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `parent_place_id` | ``$INTEGER`` |  |
| `place_id` | ``$INTEGER`` |  |
| `rank_address` | ``$INTEGER`` |  |
| `rank_search` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |

#### Example: Load

```ts
const debug = await client.debug.load({ id: 'debug_id' })
```


### Reverse

Create an instance: `const reverse = client.reverse`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$OBJECT`` |  |
| `boundingbox` | ``$ARRAY`` |  |
| `display_name` | ``$STRING`` |  |
| `lat` | ``$STRING`` |  |
| `licence` | ``$STRING`` |  |
| `lon` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |

#### Example: List

```ts
const reverses = await client.reverse.list()
```


### Search

Create an instance: `const search = client.search`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$OBJECT`` |  |
| `boundingbox` | ``$ARRAY`` |  |
| `class` | ``$STRING`` |  |
| `display_name` | ``$STRING`` |  |
| `icon` | ``$STRING`` |  |
| `importance` | ``$NUMBER`` |  |
| `lat` | ``$STRING`` |  |
| `licence` | ``$STRING`` |  |
| `lon` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |

#### Example: List

```ts
const searchs = await client.search.list()
```


### ServerStatus

Create an instance: `const server_status = client.server_status`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data_updated` | ``$STRING`` |  |
| `database_version` | ``$STRING`` |  |
| `message` | ``$STRING`` |  |
| `software_version` | ``$STRING`` |  |
| `status` | ``$INTEGER`` |  |

#### Example: Load

```ts
const server_status = await client.server_status.load({ id: 'server_status_id' })
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

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

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller as a second return value.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── Nominatim_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`Nominatim_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
addresslookup = client.addresslookup
addresslookup.load({ "id" => "example_id" })

# addresslookup.data_get now returns the loaded addresslookup data
# addresslookup.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
