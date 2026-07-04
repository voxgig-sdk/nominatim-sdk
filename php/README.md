# Nominatim PHP SDK



The PHP SDK for the Nominatim API — an entity-oriented client using PHP conventions.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nominatim-sdk/releases](https://github.com/voxgig-sdk/nominatim-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'nominatim_sdk.php';

$client = new NominatimSDK();
```

### 2. List addresslookups

```php
try {
    $result = $client->addresslookup()->list();
    if (is_array($result)) {
        foreach ($result as $item) {
            $d = $item->data_get();
            echo $d["id"] . " " . $d["name"] . "\n";
        }
    }
} catch (\Exception $err) {
    echo "Error: " . $err->getMessage();
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    echo "Error: " . $result["err"]->getMessage();
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = NominatimSDK::test();

$result = $client->addresslookup()->load(["id" => "test01"]);
// $result contains mock response data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new NominatimSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
NOMINATIM_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### NominatimSDK

```php
require_once 'nominatim_sdk.php';
$client = new NominatimSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = NominatimSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### NominatimSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `AddressLookup` | `($data): AddressLookupEntity` | Create a AddressLookup entity instance. |
| `Administrative` | `($data): AdministrativeEntity` | Create a Administrative entity instance. |
| `Debug` | `($data): DebugEntity` | Create a Debug entity instance. |
| `Reverse` | `($data): ReverseEntity` | Create a Reverse entity instance. |
| `Search` | `($data): SearchEntity` | Create a Search entity instance. |
| `ServerStatus` | `($data): ServerStatusEntity` | Create a ServerStatus entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `($reqmatch, $ctrl): array` | List entities matching the criteria. |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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
error is returned to the caller as the second element in the return array.

### Features and hooks

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── nominatim_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`nominatim_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```php
$addresslookup = $client->addresslookup();
$addresslookup->load(["id" => "example_id"]);

// $addresslookup->dataGet() now returns the loaded addresslookup data
// $addresslookup->matchGet() returns the last match criteria
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
