# Nominatim SDK

Geocode and reverse-geocode addresses using OpenStreetMap data

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Nominatim API

[Nominatim](https://nominatim.openstreetmap.org) is the geocoding service maintained by the [OpenStreetMap](https://www.openstreetmap.org) community. It indexes OSM data to let you turn place names and addresses into coordinates (forward geocoding) and turn coordinates back into addresses (reverse geocoding).

What you get from the API:

- `search` — find OSM objects by name, address fragments, or structured query parameters.
- `reverse` — find the nearest OSM object to a given latitude/longitude.
- `lookup` — fetch address details for one or more known OSM object IDs.
- `details` — internal object details, intended for debugging.
- `status` — report on the health and freshness of the server.

Responses are returned as JSON (XML and other formats are also supported by the upstream service). The public instance at `https://nominatim.openstreetmap.org` is operated by the OpenStreetMap Foundation and is subject to a strict usage policy: a maximum of 1 request per second, a valid `User-Agent` or `Referer` identifying the application, and visible attribution to OpenStreetMap contributors. For heavier workloads, self-host Nominatim or use a commercial provider.

## Try it

**TypeScript**
```bash
npm install nominatim
```

**Python**
```bash
pip install nominatim-sdk
```

**PHP**
```bash
composer require voxgig/nominatim-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/nominatim-sdk/go
```

**Ruby**
```bash
gem install nominatim-sdk
```

**Lua**
```bash
luarocks install nominatim-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { NominatimSDK } from 'nominatim'

const client = new NominatimSDK({})

// List all addresslookups
const addresslookups = await client.AddressLookup().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o nominatim-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "nominatim": {
      "command": "/abs/path/to/nominatim-mcp"
    }
  }
}
```

## Entities

The API exposes 6 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **AddressLookup** | Fetch address details for one or more known OSM object IDs via `/lookup`. | `/lookup` |
| **Administrative** | Administrative boundaries and related metadata exposed through the search and lookup endpoints. | `/polygons` |
| **Debug** | Internal object details intended for debugging via `/details`. | `/details` |
| **Reverse** | Reverse geocoding: find the OSM object closest to a given latitude/longitude via `/reverse`. | `/reverse` |
| **Search** | Forward geocoding: find OSM places by free-form query or structured address fields via `/search`. | `/search` |
| **ServerStatus** | Server health and data freshness via `/status`. | `/status` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from nominatim_sdk import NominatimSDK

client = NominatimSDK({})

# List all addresslookups
addresslookups, err = client.AddressLookup(None).list(None, None)
```

### PHP

```php
<?php
require_once 'nominatim_sdk.php';

$client = new NominatimSDK([]);

// List all addresslookups
[$addresslookups, $err] = $client->AddressLookup(null)->list(null, null);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/nominatim-sdk/go"

client := sdk.NewNominatimSDK(map[string]any{})

// List all addresslookups
addresslookups, err := client.AddressLookup(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "Nominatim_sdk"

client = NominatimSDK.new({})

# List all addresslookups
addresslookups, err = client.AddressLookup(nil).list(nil, nil)
```

### Lua

```lua
local sdk = require("nominatim_sdk")

local client = sdk.new({})

-- List all addresslookups
local addresslookups, err = client:AddressLookup(nil):list(nil, nil)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = NominatimSDK.test()
const result = await client.AddressLookup().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = NominatimSDK.test(None, None)
result, err = client.AddressLookup(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = NominatimSDK::test(null, null);
[$result, $err] = $client->AddressLookup(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.AddressLookup(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = NominatimSDK.test(nil, nil)
result, err = client.AddressLookup(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:AddressLookup(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Nominatim API

- Upstream: [https://nominatim.openstreetmap.org](https://nominatim.openstreetmap.org)
- API docs: [https://nominatim.org/release-docs/latest/api/Overview/](https://nominatim.org/release-docs/latest/api/Overview/)

- Geocoding results are derived from OpenStreetMap data, licensed under the [Open Database License (ODbL)](https://opendatacommons.org/licenses/odbl/).
- Attribution to OpenStreetMap contributors is required when displaying results.
- The Nominatim software itself is licensed under the GPLv2.
- The public instance at `nominatim.openstreetmap.org` is governed by the [OSM Foundation Nominatim Usage Policy](https://operations.osmfoundation.org/policies/nominatim/).

---

Generated from the Nominatim API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
