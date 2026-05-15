<?php
declare(strict_types=1);

// Nominatim SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

NominatimUtility::setRegistrar(function (NominatimUtility $u): void {
    $u->clean = [NominatimClean::class, 'call'];
    $u->done = [NominatimDone::class, 'call'];
    $u->make_error = [NominatimMakeError::class, 'call'];
    $u->feature_add = [NominatimFeatureAdd::class, 'call'];
    $u->feature_hook = [NominatimFeatureHook::class, 'call'];
    $u->feature_init = [NominatimFeatureInit::class, 'call'];
    $u->fetcher = [NominatimFetcher::class, 'call'];
    $u->make_fetch_def = [NominatimMakeFetchDef::class, 'call'];
    $u->make_context = [NominatimMakeContext::class, 'call'];
    $u->make_options = [NominatimMakeOptions::class, 'call'];
    $u->make_request = [NominatimMakeRequest::class, 'call'];
    $u->make_response = [NominatimMakeResponse::class, 'call'];
    $u->make_result = [NominatimMakeResult::class, 'call'];
    $u->make_point = [NominatimMakePoint::class, 'call'];
    $u->make_spec = [NominatimMakeSpec::class, 'call'];
    $u->make_url = [NominatimMakeUrl::class, 'call'];
    $u->param = [NominatimParam::class, 'call'];
    $u->prepare_auth = [NominatimPrepareAuth::class, 'call'];
    $u->prepare_body = [NominatimPrepareBody::class, 'call'];
    $u->prepare_headers = [NominatimPrepareHeaders::class, 'call'];
    $u->prepare_method = [NominatimPrepareMethod::class, 'call'];
    $u->prepare_params = [NominatimPrepareParams::class, 'call'];
    $u->prepare_path = [NominatimPreparePath::class, 'call'];
    $u->prepare_query = [NominatimPrepareQuery::class, 'call'];
    $u->result_basic = [NominatimResultBasic::class, 'call'];
    $u->result_body = [NominatimResultBody::class, 'call'];
    $u->result_headers = [NominatimResultHeaders::class, 'call'];
    $u->transform_request = [NominatimTransformRequest::class, 'call'];
    $u->transform_response = [NominatimTransformResponse::class, 'call'];
});
