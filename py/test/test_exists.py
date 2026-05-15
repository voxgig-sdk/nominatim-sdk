# ProjectName SDK exists test

import pytest
from nominatim_sdk import NominatimSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = NominatimSDK.test(None, None)
        assert testsdk is not None
