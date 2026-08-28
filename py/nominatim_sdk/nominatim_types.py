# Typed models for the Nominatim SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class AddressLookup(TypedDict, total=False):
    address: dict
    boundingbox: list
    display_name: str
    importance: float
    lat: str
    licence: str
    lon: str
    osm_id: int
    osm_type: str
    place_id: int
    type: str


class AddressLookupListMatchRequired(TypedDict):
    osm_id: str


class AddressLookupListMatch(AddressLookupListMatchRequired, total=False):
    accept_language: str
    addressdetail: int
    extratag: int
    format: str
    namedetail: int
    polygon_geojson: int
    polygon_kml: int
    polygon_svg: int
    polygon_text: int


class Administrative(TypedDict, total=False):
    country_code: str
    errormessage: str
    name: str
    osm_id: int
    osm_type: str
    place_id: int
    type: str
    updated: str


class AdministrativeListMatch(TypedDict, total=False):
    day: int
    format: str


class Debug(TypedDict, total=False):
    addresstags: dict
    admin_level: int
    calculated_importance: float
    calculated_postcode: str
    calculated_wikipedia: str
    category: str
    centroid: dict
    country_code: str
    extratags: dict
    geometry: dict
    housenumber: str
    importance: float
    indexed_date: str
    isarea: bool
    localname: str
    names: dict
    osm_id: int
    osm_type: str
    parent_place_id: int
    place_id: int
    rank_address: int
    rank_search: int
    type: str


class DebugLoadMatch(TypedDict, total=False):
    addressdetail: int
    format: str
    group_hierarchy: int
    keyword: int
    osmid: int
    osmtype: str
    place_id: int
    polygon_geojson: int


class Reverse(TypedDict, total=False):
    address: dict
    boundingbox: list
    display_name: str
    lat: str
    licence: str
    lon: str
    osm_id: int
    osm_type: str
    place_id: int


class ReverseListMatchRequired(TypedDict):
    lat: float
    lon: float


class ReverseListMatch(ReverseListMatchRequired, total=False):
    accept_language: str
    addressdetail: int
    extratag: int
    format: str
    namedetail: int
    polygon_geojson: int
    polygon_kml: int
    polygon_svg: int
    polygon_text: int
    zoom: int


class Search(TypedDict, total=False):
    address: dict
    boundingbox: list
    display_name: str
    icon: str
    importance: float
    lat: str
    licence: str
    lon: str
    osm_id: int
    osm_type: str
    place_id: int
    type: str


class SearchListMatch(TypedDict, total=False):
    accept_language: str
    addressdetail: int
    bounded: int
    city: str
    country: str
    countrycode: str
    county: str
    dedupe: int
    extratag: int
    format: str
    limit: int
    namedetail: int
    polygon_geojson: int
    polygon_kml: int
    polygon_svg: int
    polygon_text: int
    postalcode: str
    q: str
    state: str
    street: str
    viewbox: str


class ServerStatus(TypedDict, total=False):
    data_updated: str
    database_version: str
    message: str
    software_version: str
    status: int


class ServerStatusLoadMatch(TypedDict, total=False):
    format: str
