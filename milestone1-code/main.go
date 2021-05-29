package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type projectConfig struct {
	disk         string
	name         string
	repository   string
	staticAssets bool
}

func main() {

	// Flags parse
	conf, err := setupParseFlags(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", conf)
		os.Exit(1)
	}

	// Validate configuration
	errors := validateConf(conf)
	if len(errors) > 0 {
		for i := 0; i < len(errors); i++ {
			fmt.Println(errors[i].Error())
		}
		os.Exit(1)
	}

	// generate scaffold with configuration
	generateScaffold(os.Stdout, conf)

}

func setupParseFlags(w io.Writer, args []string) (projectConfig, error) {

	var conf projectConfig
	fs := flag.NewFlagSet("fs", flag.ContinueOnError)
	fs.StringVar(&conf.disk, "d", "", "Project location on disk")
	fs.StringVar(&conf.name, "n", "", "Project name")
	fs.StringVar(&conf.repository, "r", "", "Project remote repository URL")
	fs.BoolVar(&conf.staticAssets, "s", false, "Project will have static assets or not")
	fs.Usage = func() {
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		return conf, err
	}

	return conf, nil
}

func validateConf(conf projectConfig) []error {

	var allErrors []error
	if len(conf.name) == 0 {
		allErrors = append(allErrors, errors.New("Project name cannot be empty"))
	}
	if len(conf.disk) == 0 {
		allErrors = append(allErrors, errors.New("Project path cannot be empty"))
	}
	if len(conf.repository) == 0 {
		allErrors = append(allErrors, errors.New("Project repository URL cannot be empty"))
	}

	return allErrors
}

func generateScaffold(w io.Writer, conf projectConfig) {

	fmt.Fprintf(w, "Generating scaffold for project %s in %s", conf.name, conf.disk)
}
