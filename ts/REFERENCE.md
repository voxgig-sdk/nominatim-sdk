# Nominatim TypeScript SDK Reference

Complete API reference for the Nominatim TypeScript SDK.


## NominatimSDK

### Constructor

```ts
new NominatimSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NominatimSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = NominatimSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `NominatimSDK` instance in test mode.


### Instance Methods

#### `AddressLookup(data?: object)`

Create a new `AddressLookup` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AddressLookupEntity` instance.

#### `Administrative(data?: object)`

Create a new `Administrative` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AdministrativeEntity` instance.

#### `Debug(data?: object)`

Create a new `Debug` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DebugEntity` instance.

#### `Reverse(data?: object)`

Create a new `Reverse` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReverseEntity` instance.

#### `Search(data?: object)`

Create a new `Search` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SearchEntity` instance.

#### `ServerStatus(data?: object)`

Create a new `ServerStatus` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ServerStatusEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `NominatimSDK.test()`.

**Returns:** `NominatimSDK` instance in test mode.


---

## AddressLookupEntity

```ts
const address_lookup = client.AddressLookup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Record<string, any>` | No | Address breakdown |
| `boundingbox` | `any[]` | No | Bounding box coordinates |
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.AddressLookup().list({ osm_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AdministrativeEntity

```ts
const administrative = client.Administrative()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Administrative().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DebugEntity

```ts
const debug = client.Debug()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstags` | `Record<string, any>` | No | Address tags |
| `admin_level` | `number` | No | Administrative level |
| `calculated_importance` | `number` | No | Calculated importance |
| `calculated_postcode` | `string` | No | Calculated postcode |
| `calculated_wikipedia` | `string` | No | Wikipedia reference |
| `category` | `string` | No | Main OSM tag key |
| `centroid` | `Record<string, any>` | No | Centroid coordinates |
| `country_code` | `string` | No | Country code |
| `extratags` | `Record<string, any>` | No | Extra OSM tags |
| `geometry` | `Record<string, any>` | No | Geometry information |
| `housenumber` | `string` | No | House number |
| `importance` | `number` | No | Computed importance rank |
| `indexed_date` | `string` | No | Date when the object was indexed |
| `isarea` | `boolean` | No | Whether the object is an area |
| `localname` | `string` | No | Local name |
| `names` | `Record<string, any>` | No | All available names |
| `osm_id` | `number` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `parent_place_id` | `number` | No | Parent place ID |
| `place_id` | `number` | No | Unique identifier for the place |
| `rank_address` | `number` | No | Address rank |
| `rank_search` | `number` | No | Search rank |
| `type` | `string` | No | Main OSM tag value |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Debug().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DebugEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReverseEntity

```ts
const reverse = client.Reverse()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Record<string, any>` | No | Address breakdown |
| `boundingbox` | `any[]` | No | Bounding box coordinates |
| `display_name` | `string` | No | Full comma-separated address |
| `lat` | `string` | No | Latitude |
| `licence` | `string` | No | License information |
| `lon` | `string` | No | Longitude |
| `osm_id` | `number` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `place_id` | `number` | No | Unique identifier for the place |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Reverse().list({ lat: 1, lon: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReverseEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SearchEntity

```ts
const search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Record<string, any>` | No | Address breakdown |
| `boundingbox` | `any[]` | No | Bounding box coordinates |
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Search().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ServerStatusEntity

```ts
const server_status = client.ServerStatus()
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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ServerStatus().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new NominatimSDK({
  feature: {
    test: { active: true },
  }
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

