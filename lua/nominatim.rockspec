package = "voxgig-sdk-nominatim"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/nominatim-sdk.git"
}
description = {
  summary = "Nominatim SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["nominatim_sdk"] = "nominatim_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
