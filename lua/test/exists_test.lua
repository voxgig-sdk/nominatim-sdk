-- Nominatim SDK exists test

local sdk = require("nominatim_sdk")

describe("NominatimSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
