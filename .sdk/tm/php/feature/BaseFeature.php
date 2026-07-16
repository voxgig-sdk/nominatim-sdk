<?php
declare(strict_types=1);

// Nominatim SDK base feature

class NominatimBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(NominatimContext $ctx, array $options): void {}
    public function PostConstruct(NominatimContext $ctx): void {}
    public function PostConstructEntity(NominatimContext $ctx): void {}
    public function SetData(NominatimContext $ctx): void {}
    public function GetData(NominatimContext $ctx): void {}
    public function GetMatch(NominatimContext $ctx): void {}
    public function SetMatch(NominatimContext $ctx): void {}
    public function PrePoint(NominatimContext $ctx): void {}
    public function PreSpec(NominatimContext $ctx): void {}
    public function PreRequest(NominatimContext $ctx): void {}
    public function PreResponse(NominatimContext $ctx): void {}
    public function PreResult(NominatimContext $ctx): void {}
    public function PreDone(NominatimContext $ctx): void {}
    public function PreUnexpected(NominatimContext $ctx): void {}
}
