# Nominatim SDK feature factory

from nominatim_sdk.feature.base_feature import NominatimBaseFeature
from nominatim_sdk.feature.test_feature import NominatimTestFeature


def _make_feature(name):
    features = {
        "base": lambda: NominatimBaseFeature(),
        "test": lambda: NominatimTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
