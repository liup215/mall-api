docker run --rm --env GOPROXY=https://goproxy.io -v $PWD:/usr/src/app -v gopath:/go -w /usr/src/app golang:latest ./build.sh
