// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"os/exec"
	"strings"
	
	"github.com/sirupsen/logrus"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	// Goals defines the Ant targets/goals to execute.
	Goals string `envconfig:"PLUGIN_GOALS"`
}

var execCommand = exec.Command

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {

	// Split the goals into individual targets
	goals := strings.Fields(args.Goals)

	// Run `ant` command with specified goals
	antCmd := execCommand("ant", goals...)
	antOutput, antErr := antCmd.CombinedOutput()
	logrus.Info("Output of 'ant " + args.Goals + "': " + string(antOutput))

	if antErr != nil {
		logrus.WithError(antErr).Errorf("Error running 'ant %s'", args.Goals)
		return antErr
	}

	return nil
}
