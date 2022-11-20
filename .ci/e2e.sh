#!/bin/bash



#####
# This script is created to use drifter with a few basic checks in a kind cluster.
# In the long run I hope that we can use it to run test cases.
#####

# Set NS
NS=namespace-test
name=namespace-test
kubeconfigpath=/home/runner/.kube/config
# Create namespace
kubectl create namespace $NS
sleep 2
echo ${PWD}
go mod download
go build

kubectl apply -f ${PWD}/.ci/nginx.yaml

kubectl create secret generic test-secret \
  --from-literal=username=devuser  -n kube-system

kubectl create configmap test-configmap \
 --from-literal=username=devuser  -n kube-system

helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --version 5.11.0

sleep 2
./drifter version

./drifter check  -k  ${kubeconfigpath} -c  ${PWD}/.ci/check.yaml
if [ $? != 0 ]
then
    echo "E2E Tests Failed ... Should have returned exit code 0"
    exit 1
fi

./drifter check  -k  ${kubeconfigpath} -c  ${PWD}/.ci/check-helm-pass.yaml
if [ $? != 0 ]
then
    echo "E2E Tests Failed ... Should have returned exit code 0"
    exit 1
fi

for failureTest in "${PWD}"/.ci/*-fail.yaml; do
  echo "Running ${failureTest}"
  ./drifter check -k ${kubeconfigpath} -c "${failureTest}"
  if [ $? != 2 ]
  then
      echo "${failureTest} test Failed ... Should have returned exit code 2"
      exit 1
  fi
done

helm delete kubernetes-dashboard

rm -rf drifter
