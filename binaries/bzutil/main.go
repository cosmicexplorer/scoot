package main

// Command line tool to enable testing for Bazel requests in Scoot
// We use this for client-side operations Scoot does not implement,
// and are difficult to reproduce using generic tools like grpc_cli.
// Not part of production deployment
//
// Supports subcommands (--help for usage):
// * remoteexecution.Command protobuf data structure insertion to a CAS
// * longrunning.GetOperation polling of operation/scootjob by name and pretty print of result

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	remoteexecution "google.golang.org/genproto/googleapis/devtools/remoteexecution/v1test"

	"github.com/twitter/scoot/bazel/cas"
	"github.com/twitter/scoot/bazel/execution"
	"github.com/twitter/scoot/common/dialer"
	"github.com/twitter/scoot/common/log/hooks"
	scootproto "github.com/twitter/scoot/common/proto"
	"github.com/twitter/scoot/scootapi"
)

var uploadCmdStr string = "upload_command"
var getOpCmdStr string = "get_operation"
var supportedCommands map[string]bool = map[string]bool{uploadCmdStr: true, getOpCmdStr: true}

func main() {
	log.AddHook(hooks.NewContextHook())

	// TODO additional commands
	// - upload Directory (tbd)
	// - run Execute (tbd)

	// Subcommands
	uploadCommand := flag.NewFlagSet(uploadCmdStr, flag.ExitOnError)
	getCommand := flag.NewFlagSet(getOpCmdStr, flag.ExitOnError)

	// Upload Command flags
	uploadAddr := uploadCommand.String("cas_addr", scootapi.DefaultApiBundlestore_GRPC, "'host:port' of grpc CAS server")
	uploadArgv := uploadCommand.String("args", "", "space-separated command arguments, i.e. \"./run arg1 arg2\"")
	uploadEnv := uploadCommand.String("env", "", "comma-separated command environment variables, i.e. \"key1=val1,key2=val2\"")

	// Get Operation flags
	getAddr := getCommand.String("grpc_addr", scootapi.DefaultSched_GRPC, "'host:port' of grpc Exec server")
	getName := getCommand.String("name", "", "Operation name to query")

	if len(os.Args) < 2 {
		printSupported()
		os.Exit(1)
	}
	switch os.Args[1] {
	case uploadCmdStr:
		uploadCommand.Parse(os.Args[2:])
	case getOpCmdStr:
		getCommand.Parse(os.Args[2:])
	default:
		printSupported()
		os.Exit(1)
	}

	if uploadCommand.Parsed() {
		if *uploadArgv == "" {
			log.Fatalf("args required for %s", uploadCmdStr)
		}
		uploadBzCommand(*uploadAddr, *uploadArgv, *uploadEnv)
	} else if getCommand.Parsed() {
		if *getName == "" {
			log.Fatalf("name required for %s", getOpCmdStr)
		}
		getOperation(*getAddr, *getName)
	} else {
		log.Fatal("No expected commands parsed")
	}
}

func uploadBzCommand(casAddr, argv, env string) {
	cmdArgs := strings.Split(argv, " ")
	envMap := make(map[string]string)
	for _, pair := range strings.Split(env, ",") {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			continue
		}
		envMap[kv[0]] = kv[1]
	}
	log.Infof("Using argv: %q env: %s", cmdArgs, envMap)

	// create Command struct from inputs
	cmdEnvVars := []*remoteexecution.Command_EnvironmentVariable{}
	for k, v := range envMap {
		cmdEnvVars = append(cmdEnvVars, &remoteexecution.Command_EnvironmentVariable{Name: k, Value: v})
	}
	cmd := &remoteexecution.Command{
		Arguments:            cmdArgs,
		EnvironmentVariables: cmdEnvVars,
	}

	// serialize and get hash/size
	bytes, err := proto.Marshal(cmd)
	if err != nil {
		log.Fatalf("Error serializing command message: %s", err)
	}
	hash, size, err := scootproto.GetSha256(cmd)
	if err != nil {
		log.Fatalf("Error serializing command message: %s", err)
	}

	// upload command to CAS
	r := dialer.NewConstantResolver(casAddr)
	digest := &remoteexecution.Digest{Hash: hash, SizeBytes: size}
	err = cas.ByteStreamWrite(r, digest, bytes)
	if err != nil {
		log.Fatalf("Error writing to CAS: %s", err)
	}

	log.Info("Wrote to CAS successfully")
	fmt.Printf("%s %d\n", hash, size)
}

func getOperation(execAddr, opName string) {
	r := dialer.NewConstantResolver(execAddr)
	operation, err := execution.GetOperation(r, opName)
	if err != nil {
		log.Fatalf("Error making GetOperation request: %s", err)
	}

	log.Info(execution.ExecuteOperationToStr(operation))
}

func printSupported() {
	cmds := make([]string, 0, len(supportedCommands))
	for k := range supportedCommands {
		cmds = append(cmds, k)
	}
	fmt.Printf("Supported commands: %s\n", strings.Join(cmds, ", "))
}
