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
go mod download
go build
kubectl get pods --v=6
kubectl get ds -n kube-system
kubectl get deployments -A
kubectl get sc -A
kubectl apply -f ${PWD}/.ci/nginx.yaml

helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --version 5.11.0

./drifter version
./drifter check  -k  /home/runner/.kube/config -c  ${PWD}/.ci/check.yaml

if [ $? != 0 ]
then
    echo "E2E Tests Failed ... Should have returned exit code 0"
    exit 1
fi

./drifter check  -k  /home/runner/.kube/config -c  ${PWD}/.ci/check-fail.yaml
if [ $? != 2 ]
then
    echo "E2E Tests Failed ... Should have returned exit code 2"
    exit 1
fi

./drifter check  -k  /home/runner/.kube/config -c  ${PWD}/.ci/check-helm-pass.yaml
if [ $? != 0 ]
then
    echo "E2E Tests Failed ... Should have returned exit code 0"
    exit 1
fi

./drifter check  -k  /home/runner/.kube/config -c  ${PWD}/.ci/check-helm-fail.yaml
if [ $? != 2 ]
then
    echo "E2E Tests Failed ... Should have returned exit code 0"
    exit 1
fi


helm delete kubernetes-dashboard

rm -rf drifter
