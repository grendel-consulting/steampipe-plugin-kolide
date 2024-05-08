#!/usr/bin/env bash

define_test_results(){
    # Unexpected behaviour in Kolide API under K2; this endpoint returns an empty list
    export EXPECTED_COUNT="0"
    # export TITLE=""
    # export DETECTED_AT=""
    # export BLOCKS_DEVICE_AT=""
    # export RESOLVED_AT=""
    # export EXEMPTED=""
}
