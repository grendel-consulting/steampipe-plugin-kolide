#!/usr/bin/env bash

define_test_results(){
    # Unexpected behaviour in Kolide API under K2; this endpoint returns an existing active users
    export EXPECTED_COUNT="2"
    # export NAME=""
    # export EMAIL=""
}
