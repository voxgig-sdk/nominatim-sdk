<?php
declare(strict_types=1);

// Nominatim SDK exists test

require_once __DIR__ . '/../nominatim_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = NominatimSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
