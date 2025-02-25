package ised

import (
	"testing"

	"github.com/pcunning/hamcall/data"
)

func TestBasic(t *testing.T) {
	line := "VA2TRW;Tripp;Sanders;;;QC;;A;;;;;;;;;;"
	call := ProcessLine(line)

	if call.Class != "Basic" {
		t.Fatalf(`call.Class = %s, not 'Basic'`, call.Class)
	}
}

func TestBasicWithHonors(t *testing.T) {
	line := "VA2TRW;Tripp;Sanders;;;QC;;;;;;E;;;;;;"
	call := ProcessLine(line)

	if call.Class != "Basic with Honors" {
		t.Fatalf(`call.Class = %s, not 'Basic With Honors'`, call.Class)
	}
}

func TestAdvanced(t *testing.T) {
	line := "VA2TRW;Tripp;Sanders;;;QC;;;;;D;;;;;;;"
	call := ProcessLine(line)

	if call.Class != "Advanced" {
		t.Fatalf(`call.Class = %s, not 'Advanced'`, call.Class)
	}
}

func TestWithMorse(t *testing.T) {
	var line string
	var call data.HamCall

	line = "VA2TRW;Tripp;Sanders;;;QC;;A;B;;;;;;;;;"
	call = ProcessLine(line)

	if call.Class != "Basic (w/ Morse)" {
		t.Fatalf(`call.Class = %s, not 'Basic (w/ Morse)'`, call.Class)
	}

	line = "VA2TRW;Tripp;Sanders;;;QC;;;B;;;E;;;;;;"
	call = ProcessLine(line)

	if call.Class != "Basic with Honors (w/ Morse)" {
		t.Fatalf(`call.Class = %s, not 'Basic with Honors (w/ Morse)'`, call.Class)
	}

	line = "VA2TRW;Tripp;Sanders;;;QC;;;B;;D;;;;;;;"
	call = ProcessLine(line)

	if call.Class != "Advanced (w/ Morse)" {
		t.Fatalf(`call.Class = %s, not 'Advanced (w/ Morse)'`, call.Class)
	}
}

func TestFileImport(t *testing.T) {
	calls := make(map[string]data.HamCall)
	Process(&calls, "./ised_test_data.txt")

	if len(calls) != 5 {
		t.Fatalf(`len(calls) = %d, not 5`, len(calls))
	}

	if calls["VA4AAA"].State != "MB" {
		t.Fatalf(`call.State = %s, not 'MB'`, calls["VA4AAA"].State)
	}

	if calls["VA4AAA"].Class != "Advanced (w/ Morse)" {
		t.Fatalf(`call.Class = %s, not 'Advanced (w/ Morse)'`, calls["VA4AAA"].Class)
	}
}
