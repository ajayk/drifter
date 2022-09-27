#!/bin/bash

#####
# This script is created to use drifter with a few basic checks in a kind cluster.
# In the long run I hope that we can use it to run test cases.
#####

# Set NS
NS=namespace-test
name=namespace-test

# Create namespace
kubectl create namespace $NS

sleep 2

echo ${PWD}
ls -al