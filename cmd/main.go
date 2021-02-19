// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

// DO NOT MODIFY THIS FILE DIRECTLY

package main

import (
	"os"

	"github.com/drone-plugins/drone-plugin-lib/pkg/urfave"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/KazanExpress/drone-envtpl/pkg/plugin"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "envtpl plugin"
	app.Usage = "fills jinja templates with env vars"
	app.Action = run

	// Create the flags for the application
	flags := settingsFlags()
	flags = append(flags, urfave.Flags()...)

	app.Flags = modifyFlags(flags)

	// Run the application
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

// Why just dont use github api to commit generated file?

func run(c *cli.Context) error {
	var pipeline = urfave.PipelineFromContext(c)
	config := plugin.Config{
		Build:    pipeline.Build,
		Repo:     pipeline.Repo,
		Commit:   pipeline.Commit,
		Stage:    pipeline.Stage,
		Step:     pipeline.Step,
		SemVer:   pipeline.SemVer,
		Settings: settingsFromContext(c),
	}

	plugin := plugin.New(config)

	// Validate the settings
	if err := plugin.Validate(); err != nil {
		return errors.Wrap(err, "validation failed")
	}

	// Run the plugin
	if err := plugin.Exec(); err != nil {
		return errors.Wrap(err, "exec failed")
	}

	return nil
}
