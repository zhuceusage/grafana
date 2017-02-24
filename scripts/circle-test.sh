#!/bin/bash
function exit_if_fail {
    command=$@
    echo "Executing '$command'"
    eval $command
    rc=$?
    if [ $rc -ne 0 ]; then
        echo "'$command' returned $rc."
        exit $rc
    fi
}

cd /home/ubuntu/.go_workspace/src/github.com/grafana/grafana

rm -rf node_modules
npm install -g yarn --quiet
yarn install --pure-lockfile --no-progress

exit_if_fail npm test

echo "running go fmt"
exit_if_fail test -z "$(gofmt -s -l ./pkg | tee /dev/stderr)"

echo "running go vet"
exit_if_fail test -z "$(go vet ./pkg/... | tee /dev/stderr)"

echo "running go test"
exit_if_fail go test -v ./pkg/...

#exit_if_fail go run build.go build
if [ "$CIRCLE_TAG" != "" ]; then
  echo "Building a release from tag $CIRCLE_TAG"
  exit_if_fail go run build.go -buildNumber=${CIRCLE_BUILD_NUM} -includeBuildNumber=false build
else
  echo "Building incremental build for $CIRCLE_BRANCH"
  exit_if_fail go run build.go -buildNumber=${CIRCLE_BUILD_NUM} build
fi
