#!/usr/bin/env bats

load _support/globals
load _support/extensions

@test "kolide_live_query_campaign" {
  run steampipe query "select * from kolide_live_query_campaign order by created_at desc;"
  assert_success
  assert_output --partial "id"
  assert_output --partial "query"
  assert_output --partial "status"
  assert_output --partial "created_at"
  assert_output --partial "updated_at"
}
