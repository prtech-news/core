#/bin/bash
echo "Setup dev environment"
export GOPATH=$PWD
echo "export GOPATH = $GOPATH"
echo "export GOPRIVATE = github.com/prtech-news/*"
echo "export GONOPROXY=none"
echo "export GONOSUMDB=github.com/prtech-news/*"
echo "export GONOPROXY=github.com/prtech-news/*"
echo "export GOPRIVATE=github.com/prtech-news/common"
echo "Finished"
