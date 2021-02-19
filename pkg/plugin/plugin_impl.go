// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/pkg/errors"
)

var errFileUploadConflict = errors.New("file upload conflict")

// Settings for the Plugin.
type Settings struct {
	// Fill in the data structure with appropriate values
	TemplatePath   string
	OutputFilePath string
}

func (p *pluginImpl) Validate() error {
	// Validate the Config and return an error if there are issues.
	if p.config.Settings.TemplatePath == "" {
		return errors.New("settings param template is empty")
	}

	if p.config.Settings.OutputFilePath == "" {
		return errors.New("settings param output_file is empty")
	}

	return nil
}

func (p *pluginImpl) Exec() error {

	var generatedFile, err = p.generateFromTemplate()

	if err != nil {
		return errors.Wrap(err, "template generation failure")
	}

	err = p.saveFile(generatedFile)

	return err
}

func (p *pluginImpl) generateFromTemplate() ([]byte, error) {
	var template, err = ioutil.ReadFile(p.config.Settings.TemplatePath)
	if err != nil {
		return nil, err
	}
	return fillJinjaTemplate(template)
}

func (p *pluginImpl) saveFile(fileContent []byte) error {
	var settings = p.config.Settings

	if settings.OutputFilePath == "" {
		settings.OutputFilePath = path.Join(p.config.Repo.Name, fmt.Sprintf("deploy.%s.yaml", p.config.Commit.Branch))
	}

	ioutil.WriteFile(settings.OutputFilePath, fileContent, 0644)

	return nil
}

func fillJinjaTemplate(template []byte) ([]byte, error) {
	var output bytes.Buffer

	var cmd = exec.Command("envtpl")
	cmd.Stdin = bytes.NewBuffer(template)
	cmd.Stderr = os.Stderr
	cmd.Stdout = &output

	var err = cmd.Run()

	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
