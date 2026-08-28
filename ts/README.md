# Nominatim TypeScript SDK



The TypeScript SDK for the Nominatim API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.AddressLookup()` — each with a small set of operations (`list`, `load`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nominatim-sdk/releases](https://github.com/voxgig-sdk/nominatim-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { NominatimSDK } from '@voxgig-sdk/nominatim'

const client = new NominatimSDK()
```

### 2. List addresslookup records

`list()` resolves to an array of AddressLookup ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const addresslookups = await client.AddressLookup().list({ osm_id: "example" })

for (const addresslookup of addresslookups) {
  console.log(addresslookup)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const reverses = await client.Reverse().list()
  console.log(reverses)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = NominatimSDK.test()

const reverse = await client.Reverse().list()
// reverse is the entity, populated with mock response data
// — call reverse.data() for the record itself
console.log(reverse)
```

You can also use the instance method:

```ts
const client = new NominatimSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Reverse()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new NominatimSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NOMINATIM_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### NominatimSDK

#### Constructor

```ts
new NominatimSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `AddressLookup(data?)` | `AddressLookupEntity` | Create an AddressLookup entity instance. |
| `Administrative(data?)` | `AdministrativeEntity` | Create an Administrative entity instance. |
| `Debug(data?)` | `DebugEntity` | Create a Debug entity instance. |
| `Reverse(data?)` | `ReverseEntity` | Create a Reverse entity instance. |
| `Search(data?)` | `SearchEntity` | Create a Search entity instance. |
| `ServerStatus(data?)` | `ServerStatusEntity` | Create a ServerStatus entity instance. |
| `tester(testopts?, sdkopts?)` | `NominatimSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `NominatimSDK.test(testopts?, sdkopts?)` | `NominatimSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): NominatimSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` resolves to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### AddressLookup

| Field | Description |
| --- | --- |
| `address` | Address breakdown |
| `boundingbox` | Bounding box coordinates |
| `class` | Main OSM tag key |
| `display_name` | Full comma-separated address |
| `importance` | Computed importance rank |
| `lat` | Latitude |
| `licence` | License information |
| `lon` | Longitude |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `place_id` | Unique identifier for the place |
| `type` | Main OSM tag value |

Operations: list.

API path: `/lookup`

#### Administrative

| Field | Description |
| --- | --- |
| `class` | Main OSM tag key |
| `country_code` | Country code |
| `errormessage` | Error message describing the polygon issue |
| `name` | Name of the object |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (way, relation) |
| `place_id` | Unique identifier for the place |
| `type` | Main OSM tag value |
| `updated` | Last update timestamp |

Operations: list.

API path: `/polygons`

#### Debug

| Field | Description |
| --- | --- |
| `addresstags` | Address tags |
| `admin_level` | Administrative level |
| `calculated_importance` | Calculated importance |
| `calculated_postcode` | Calculated postcode |
| `calculated_wikipedia` | Wikipedia reference |
| `category` | Main OSM tag key |
| `centroid` | Centroid coordinates |
| `country_code` | Country code |
| `extratags` | Extra OSM tags |
| `geometry` | Geometry information |
| `housenumber` | House number |
| `importance` | Computed importance rank |
| `indexed_date` | Date when the object was indexed |
| `isarea` | Whether the object is an area |
| `localname` | Local name |
| `names` | All available names |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `parent_place_id` | Parent place ID |
| `place_id` | Unique identifier for the place |
| `rank_address` | Address rank |
| `rank_search` | Search rank |
| `type` | Main OSM tag value |

Operations: load.

API path: `/details`

#### Reverse

| Field | Description |
| --- | --- |
| `address` | Address breakdown |
| `boundingbox` | Bounding box coordinates |
| `display_name` | Full comma-separated address |
| `lat` | Latitude |
| `licence` | License information |
| `lon` | Longitude |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `place_id` | Unique identifier for the place |

Operations: list.

API path: `/reverse`

#### Search

| Field | Description |
| --- | --- |
| `address` | Address breakdown |
| `boundingbox` | Bounding box coordinates |
| `class` | Main OSM tag key |
| `display_name` | Full comma-separated address |
| `icon` | URL of icon representing the place |
| `importance` | Computed importance rank |
| `lat` | Latitude |
| `licence` | License information |
| `lon` | Longitude |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `place_id` | Unique identifier for the place |
| `type` | Main OSM tag value |

Operations: list.

API path: `/search`

#### ServerStatus

| Field | Description |
| --- | --- |
| `data_updated` | Timestamp when the database was last updated |
| `database_version` | Database version |
| `message` | Status message |
| `software_version` | Nominatim software version |
| `status` | Status code (0 = OK) |

Operations: load.

API path: `/status`



## Entities


### AddressLookup

Create an instance: `const address_lookup = client.AddressLookup()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `Record<string, any>` | Address breakdown |
| `boundingbox` | `any[]` | Bounding box coordinates |
| `class` | `string` | Main OSM tag key |
| `display_name` | `string` | Full comma-separated address |
| `importance` | `number` | Computed importance rank |
| `lat` | `string` | Latitude |
| `licence` | `string` | License information |
| `lon` | `string` | Longitude |
| `osm_id` | `number` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `place_id` | `number` | Unique identifier for the place |
| `type` | `string` | Main OSM tag value |

#### Example: List

```ts
const address_lookups = await client.AddressLookup().list({ osm_id: "example" })
```


### Administrative

Create an instance: `const administrative = client.Administrative()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `class` | `string` | Main OSM tag key |
| `country_code` | `string` | Country code |
| `errormessage` | `string` | Error message describing the polygon issue |
| `name` | `string` | Name of the object |
| `osm_id` | `number` | OSM object ID |
| `osm_type` | `string` | OSM type (way, relation) |
| `place_id` | `number` | Unique identifier for the place |
| `type` | `string` | Main OSM tag value |
| `updated` | `string` | Last update timestamp |

#### Example: List

```ts
const administratives = await client.Administrative().list()
```


### Debug

Create an instance: `const debug = client.Debug()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresstags` | `Record<string, any>` | Address tags |
| `admin_level` | `number` | Administrative level |
| `calculated_importance` | `number` | Calculated importance |
| `calculated_postcode` | `string` | Calculated postcode |
| `calculated_wikipedia` | `string` | Wikipedia reference |
| `category` | `string` | Main OSM tag key |
| `centroid` | `Record<string, any>` | Centroid coordinates |
| `country_code` | `string` | Country code |
| `extratags` | `Record<string, any>` | Extra OSM tags |
| `geometry` | `Record<string, any>` | Geometry information |
| `housenumber` | `string` | House number |
| `importance` | `number` | Computed importance rank |
| `indexed_date` | `string` | Date when the object was indexed |
| `isarea` | `boolean` | Whether the object is an area |
| `localname` | `string` | Local name |
| `names` | `Record<string, any>` | All available names |
| `osm_id` | `number` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `parent_place_id` | `number` | Parent place ID |
| `place_id` | `number` | Unique identifier for the place |
| `rank_address` | `number` | Address rank |
| `rank_search` | `number` | Search rank |
| `type` | `string` | Main OSM tag value |

#### Example: Load

```ts
const debug = await client.Debug().load()
```


### Reverse

Create an instance: `const reverse = client.Reverse()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `Record<string, any>` | Address breakdown |
| `boundingbox` | `any[]` | Bounding box coordinates |
| `display_name` | `string` | Full comma-separated address |
| `lat` | `string` | Latitude |
| `licence` | `string` | License information |
| `lon` | `string` | Longitude |
| `osm_id` | `number` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `place_id` | `number` | Unique identifier for the place |

#### Example: List

```ts
const reverses = await client.Reverse().list({ lat: 1, lon: 1 })
```


### Search

Create an instance: `const search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `Record<string, any>` | Address breakdown |
| `boundingbox` | `any[]` | Bounding box coordinates |
| `class` | `string` | Main OSM tag key |
| `display_name` | `string` | Full comma-separated address |
| `icon` | `string` | URL of icon representing the place |
| `importance` | `number` | Computed importance rank |
| `lat` | `string` | Latitude |
| `licence` | `string` | License information |
| `lon` | `string` | Longitude |
| `osm_id` | `number` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `place_id` | `number` | Unique identifier for the place |
| `type` | `string` | Main OSM tag value |

#### Example: List

```ts
const searchs = await client.Search().list()
```


### ServerStatus

Create an instance: `const server_status = client.ServerStatus()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data_updated` | `string` | Timestamp when the database was last updated |
| `database_version` | `string` | Database version |
| `message` | `string` | Status message |
| `software_version` | `string` | Nominatim software version |
| `status` | `number` | Status code (0 = OK) |

#### Example: Load

```ts
const server_status = await client.ServerStatus().load()
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
nominatim/
├── src/
│   ├── NominatimSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { NominatimSDK } from '@voxgig-sdk/nominatim'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const reverse = client.Reverse()
await reverse.list()

// reverse.data() now returns the reverse data from the last `list`
// reverse.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
