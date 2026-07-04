# Nominatim Python SDK



The Python SDK for the Nominatim API — an entity-oriented client following Pythonic conventions.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nominatim-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
from nominatim_sdk import NominatimSDK

client = NominatimSDK()
```

### 2. List addresslookup records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    addresslookups = client.AddressLookup().list({})
    for addresslookup in addresslookups:
        print(addresslookup)
except Exception as err:
    print(f"list failed: {err}")
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    print(result["err"])     # error value
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = NominatimSDK.test()

# Entity ops return the bare record and raise on error.
addresslookup = client.AddressLookup().load({"id": "test01"})
# addresslookup contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = NominatimSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### NominatimSDK

```python
from nominatim_sdk import NominatimSDK

client = NominatimSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = NominatimSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### NominatimSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the bare result data (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `address_lookup = client.AddressLookup()`

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

```python
address_lookups = client.AddressLookup().list({})
```


### Administrative

Create an instance: `administrative = client.Administrative()`

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

```python
administratives = client.Administrative().list({})
```


### Debug

Create an instance: `debug = client.Debug()`

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

```python
debug = client.Debug().load({"id": "debug_id"})
```


### Reverse

Create an instance: `reverse = client.Reverse()`

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

```python
reverses = client.Reverse().list({})
```


### Search

Create an instance: `search = client.Search()`

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

```python
searchs = client.Search().list({})
```


### ServerStatus

Create an instance: `server_status = client.ServerStatus()`

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

```python
server_status = client.ServerStatus().load({"id": "server_status_id"})
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
error is returned to the caller as the second element in the return tuple.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── nominatim_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`nominatim_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
addresslookup = client.AddressLookup()
addresslookup.load({"id": "example_id"})

# addresslookup.data_get() now returns the loaded addresslookup data
# addresslookup.match_get() returns the last match criteria
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
