package workerapi

// We should use go generate to run:
//   thrift --gen go:package_prefix=github.com/scootdev/scoot/workerapi/gen-go/ worker.thrift
//   find gen-go/ -name '*.go' -exec sed -i'' 's,git.apache.org/thrift.git/lib/go/thrift,github.com/apache/thrift/lib/go/thrift,' {} \;
