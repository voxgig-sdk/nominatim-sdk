# Nominatim SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

NominatimUtility.registrar = ->(u) {
  u.clean = NominatimUtilities::Clean
  u.done = NominatimUtilities::Done
  u.make_error = NominatimUtilities::MakeError
  u.feature_add = NominatimUtilities::FeatureAdd
  u.feature_hook = NominatimUtilities::FeatureHook
  u.feature_init = NominatimUtilities::FeatureInit
  u.fetcher = NominatimUtilities::Fetcher
  u.make_fetch_def = NominatimUtilities::MakeFetchDef
  u.make_context = NominatimUtilities::MakeContext
  u.make_options = NominatimUtilities::MakeOptions
  u.make_request = NominatimUtilities::MakeRequest
  u.make_response = NominatimUtilities::MakeResponse
  u.make_result = NominatimUtilities::MakeResult
  u.make_point = NominatimUtilities::MakePoint
  u.make_spec = NominatimUtilities::MakeSpec
  u.make_url = NominatimUtilities::MakeUrl
  u.param = NominatimUtilities::Param
  u.prepare_auth = NominatimUtilities::PrepareAuth
  u.prepare_body = NominatimUtilities::PrepareBody
  u.prepare_headers = NominatimUtilities::PrepareHeaders
  u.prepare_method = NominatimUtilities::PrepareMethod
  u.prepare_params = NominatimUtilities::PrepareParams
  u.prepare_path = NominatimUtilities::PreparePath
  u.prepare_query = NominatimUtilities::PrepareQuery
  u.result_basic = NominatimUtilities::ResultBasic
  u.result_body = NominatimUtilities::ResultBody
  u.result_headers = NominatimUtilities::ResultHeaders
  u.transform_request = NominatimUtilities::TransformRequest
  u.transform_response = NominatimUtilities::TransformResponse
}
