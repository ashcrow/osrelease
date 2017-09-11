package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/ashcrow/osrelease"
	"github.com/spf13/cobra"
)

// version of the build
var version string

// latest git commit in this build
var gitCommit string

// when the build occured
var buildInfo string

// main is the main entry point for executing the CLI
func main() {
	// The OSRelease instance to use in commands
	or, err := osrelease.New()
	if err != nil {
		log.Fatalln(err)
	}

	// root command which doesn't do anything but show help
	var rootCommand = &cobra.Command{
		Use:   "osrelease",
		Short: "osrelease parses and provides osrelease data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("Please specify an output format.")
		},
	}

	// The version number of this build
	var versionCommand = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of osrelease",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	// The version number of this build
	var buildInfoCommand = &cobra.Command{
		Use:   "buildinfo",
		Short: "Print the build information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", version)
			fmt.Printf("GitCommit: %s\n", gitCommit)
			fmt.Printf("BuildInfo: %s\n", buildInfo)
		},
	}

	// command for outputting in json
	var jsonCommand = &cobra.Command{
		Use:   "json",
		Short: "Parse and print the file in json format",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := json.MarshalIndent(or, "", " ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(string(data))
		},
	}

	// command for outputting in YAML
	var yamlCommand = &cobra.Command{
		Use:   "yaml",
		Short: "Parse and print the file in yaml format",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := yaml.Marshal(or)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(string(data))
		},
	}

	// command for outputting a single field
	var fieldCommand = &cobra.Command{
		Use:   "field FIELDNAME",
		Short: "Parse and print a single field",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				return nil
			}
			return errors.New("Requires a single field to be provided")
		},
		Run: func(cmd *cobra.Command, args []string) {
			data, err := or.GetField(args[0])
			if err != nil {
				os.Exit(1)
			}
			fmt.Printf(string(data))
		},
	}

	// Add our commands
	rootCommand.AddCommand(versionCommand, buildInfoCommand, jsonCommand, yamlCommand, fieldCommand)

	// Execute
	rootCommand.Execute()
}
