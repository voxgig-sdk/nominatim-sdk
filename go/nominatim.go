package voxgignominatimsdk

import (
	"github.com/voxgig-sdk/nominatim-sdk/core"
	"github.com/voxgig-sdk/nominatim-sdk/entity"
	"github.com/voxgig-sdk/nominatim-sdk/feature"
	_ "github.com/voxgig-sdk/nominatim-sdk/utility"
)

// Type aliases preserve external API.
type NominatimSDK = core.NominatimSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type NominatimEntity = core.NominatimEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type NominatimError = core.NominatimError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAddressLookupEntityFunc = func(client *core.NominatimSDK, entopts map[string]any) core.NominatimEntity {
		return entity.NewAddressLookupEntity(client, entopts)
	}
	core.NewAdministrativeEntityFunc = func(client *core.NominatimSDK, entopts map[string]any) core.NominatimEntity {
		return entity.NewAdministrativeEntity(client, entopts)
	}
	core.NewDebugEntityFunc = func(client *core.NominatimSDK, entopts map[string]any) core.NominatimEntity {
		return entity.NewDebugEntity(client, entopts)
	}
	core.NewReverseEntityFunc = func(client *core.NominatimSDK, entopts map[string]any) core.NominatimEntity {
		return entity.NewReverseEntity(client, entopts)
	}
	core.NewSearchEntityFunc = func(client *core.NominatimSDK, entopts map[string]any) core.NominatimEntity {
		return entity.NewSearchEntity(client, entopts)
	}
	core.NewServerStatusEntityFunc = func(client *core.NominatimSDK, entopts map[string]any) core.NominatimEntity {
		return entity.NewServerStatusEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewNominatimSDK = core.NewNominatimSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
