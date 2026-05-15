# Nominatim SDK exists test

require "minitest/autorun"
require_relative "../Nominatim_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = NominatimSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
