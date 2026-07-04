<?php
declare(strict_types=1);

// Administrative entity test

require_once __DIR__ . '/../nominatim_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class AdministrativeEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NominatimSDK::test(null, null);
        $ent = $testsdk->Administrative(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = administrative_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "administrative." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NOMINATIM_TEST_ADMINISTRATIVE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $administrative_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.administrative")));
        $administrative_ref01_data = null;
        if (count($administrative_ref01_data_raw) > 0) {
            $administrative_ref01_data = Helpers::to_map($administrative_ref01_data_raw[0][1]);
        }

        // LIST
        $administrative_ref01_ent = $client->Administrative(null);
        $administrative_ref01_match = [];

        $administrative_ref01_list_result = $administrative_ref01_ent->list($administrative_ref01_match, null);
        $this->assertIsArray($administrative_ref01_list_result);

    }
}

function administrative_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/administrative/AdministrativeTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NominatimSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["administrative01", "administrative02", "administrative03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NOMINATIM_TEST_ADMINISTRATIVE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NOMINATIM_TEST_ADMINISTRATIVE_ENTID" => $idmap,
        "NOMINATIM_TEST_LIVE" => "FALSE",
        "NOMINATIM_TEST_EXPLAIN" => "FALSE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NOMINATIM_TEST_ADMINISTRATIVE_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["NOMINATIM_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
            ],
            $extra ?? [],
        ]);
        $client = new NominatimSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["NOMINATIM_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["NOMINATIM_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
