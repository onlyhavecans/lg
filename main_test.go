package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func Test_run(t *testing.T) {
	dateFmt := "--date=\"%s\" new "
	dateLayout := "2006/01/02 15:04:05"
	now := time.Now()
	dateStr := fmt.Sprintf(dateFmt, now.Format(dateLayout))
	hourAgoStr := fmt.Sprintf(dateFmt, now.Add(-time.Hour).Format(dateLayout))

	tests := []struct {
		name       string
		args       []string
		wantStdout string
		wantErr    bool
	}{
		{"no args", []string{"lg"}, "", true},
		{"entry", []string{"lg", "entry"}, dateStr + "entry\n", false},
		{"time shift", []string{"lg", "[1 hour ago] entry"}, hourAgoStr + "entry\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			d := DayOne{
				DayOneCmd: exec.Command("echo"),
				stdout:    stdout,
			}
			err := run(tt.args, d)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("run() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
		})
	}
}
