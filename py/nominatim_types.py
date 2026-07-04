# Typed models for the Nominatim SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class AddressLookup:
    address: Optional[dict] = None
    boundingbox: Optional[list] = None
    display_name: Optional[str] = None
    importance: Optional[float] = None
    lat: Optional[str] = None
    licence: Optional[str] = None
    lon: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None
    type: Optional[str] = None


@dataclass
class AddressLookupListMatch:
    address: Optional[dict] = None
    boundingbox: Optional[list] = None
    display_name: Optional[str] = None
    importance: Optional[float] = None
    lat: Optional[str] = None
    licence: Optional[str] = None
    lon: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None
    type: Optional[str] = None


@dataclass
class Administrative:
    country_code: Optional[str] = None
    errormessage: Optional[str] = None
    name: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None
    type: Optional[str] = None
    updated: Optional[str] = None


@dataclass
class AdministrativeListMatch:
    country_code: Optional[str] = None
    errormessage: Optional[str] = None
    name: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None
    type: Optional[str] = None
    updated: Optional[str] = None


@dataclass
class Debug:
    addresstag: Optional[dict] = None
    admin_level: Optional[int] = None
    calculated_importance: Optional[float] = None
    calculated_postcode: Optional[str] = None
    calculated_wikipedia: Optional[str] = None
    category: Optional[str] = None
    centroid: Optional[dict] = None
    country_code: Optional[str] = None
    extratag: Optional[dict] = None
    geometry: Optional[dict] = None
    housenumber: Optional[str] = None
    importance: Optional[float] = None
    indexed_date: Optional[str] = None
    isarea: Optional[bool] = None
    localname: Optional[str] = None
    name: Optional[dict] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    parent_place_id: Optional[int] = None
    place_id: Optional[int] = None
    rank_address: Optional[int] = None
    rank_search: Optional[int] = None
    type: Optional[str] = None


@dataclass
class DebugLoadMatch:
    addresstag: Optional[dict] = None
    admin_level: Optional[int] = None
    calculated_importance: Optional[float] = None
    calculated_postcode: Optional[str] = None
    calculated_wikipedia: Optional[str] = None
    category: Optional[str] = None
    centroid: Optional[dict] = None
    country_code: Optional[str] = None
    extratag: Optional[dict] = None
    geometry: Optional[dict] = None
    housenumber: Optional[str] = None
    importance: Optional[float] = None
    indexed_date: Optional[str] = None
    isarea: Optional[bool] = None
    localname: Optional[str] = None
    name: Optional[dict] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    parent_place_id: Optional[int] = None
    place_id: Optional[int] = None
    rank_address: Optional[int] = None
    rank_search: Optional[int] = None
    type: Optional[str] = None


@dataclass
class Reverse:
    address: Optional[dict] = None
    boundingbox: Optional[list] = None
    display_name: Optional[str] = None
    lat: Optional[str] = None
    licence: Optional[str] = None
    lon: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None


@dataclass
class ReverseListMatch:
    address: Optional[dict] = None
    boundingbox: Optional[list] = None
    display_name: Optional[str] = None
    lat: Optional[str] = None
    licence: Optional[str] = None
    lon: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None


@dataclass
class Search:
    address: Optional[dict] = None
    boundingbox: Optional[list] = None
    display_name: Optional[str] = None
    icon: Optional[str] = None
    importance: Optional[float] = None
    lat: Optional[str] = None
    licence: Optional[str] = None
    lon: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None
    type: Optional[str] = None


@dataclass
class SearchListMatch:
    address: Optional[dict] = None
    boundingbox: Optional[list] = None
    display_name: Optional[str] = None
    icon: Optional[str] = None
    importance: Optional[float] = None
    lat: Optional[str] = None
    licence: Optional[str] = None
    lon: Optional[str] = None
    osm_id: Optional[int] = None
    osm_type: Optional[str] = None
    place_id: Optional[int] = None
    type: Optional[str] = None


@dataclass
class ServerStatus:
    data_updated: Optional[str] = None
    database_version: Optional[str] = None
    message: Optional[str] = None
    software_version: Optional[str] = None
    status: Optional[int] = None


@dataclass
class ServerStatusLoadMatch:
    data_updated: Optional[str] = None
    database_version: Optional[str] = None
    message: Optional[str] = None
    software_version: Optional[str] = None
    status: Optional[int] = None

