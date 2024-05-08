#!/usr/bin/env bash

define_test_results(){
    # Unexpected behaviour in Kolide API under K2; this endpoint returns an empty list
    export EXPECTED_COUNT="0"
    # export ID=""
    # export NAME=""
    # export EMAIL=""
    # export HAS_REGISTERED_DEVICE=""
}
