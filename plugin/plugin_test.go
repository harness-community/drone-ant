// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestVerifyEnvs(t *testing.T) {
	type args struct {
		ctx  context.Context
		args Args
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid goals",
			args: args{
				ctx: context.TODO(),
				args: Args{
					Goals: "clean compile",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid goals",
			args: args{
				ctx: context.TODO(),
				args: Args{
					Goals: "invalid-goal",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Exec(tt.args.ctx, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// Mock exec.Command to avoid actual command execution during tests
	execCommand = func(name string, arg ...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperProcess", "--", name}
		cs = append(cs, arg...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
		return cmd
	}

	code := m.Run()

	// Restore exec.Command after tests
	execCommand = exec.Command

	os.Exit(code)
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// Output expected based on args
	if strings.Contains(strings.Join(os.Args, " "), "invalid-goal") {
		os.Exit(1)
	}
	os.Exit(0)
}
