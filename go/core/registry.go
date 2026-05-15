package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAddressLookupEntityFunc func(client *NominatimSDK, entopts map[string]any) NominatimEntity

var NewAdministrativeEntityFunc func(client *NominatimSDK, entopts map[string]any) NominatimEntity

var NewDebugEntityFunc func(client *NominatimSDK, entopts map[string]any) NominatimEntity

var NewReverseEntityFunc func(client *NominatimSDK, entopts map[string]any) NominatimEntity

var NewSearchEntityFunc func(client *NominatimSDK, entopts map[string]any) NominatimEntity

var NewServerStatusEntityFunc func(client *NominatimSDK, entopts map[string]any) NominatimEntity

