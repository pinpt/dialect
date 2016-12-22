package copyright

import (
	"testing"

	"github.com/pinpt/dialect/pkg/copyright"
)

func TestCopyright1(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright (c) 2010 by Jeff Haynie")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "Jeff Haynie" {
		t.Fatalf("result name should have been Jeff Haynie, but was %s", result.Name)
	}
}

func TestCopyright2(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright (c) 2010-2012 by Jeff Haynie")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 2 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Dates[1] != "2012" {
		t.Fatalf("result date should be 2012, but was %s", result.Dates[0])
	}
	if result.Name != "Jeff Haynie" {
		t.Fatalf("result name should have been Jeff Haynie, but was %s", result.Name)
	}
}

func TestCopyright3(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright 2010-2012 by Jeff Haynie")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 2 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Dates[1] != "2012" {
		t.Fatalf("result date should be 2012, but was %s", result.Dates[0])
	}
	if result.Name != "Jeff Haynie" {
		t.Fatalf("result name should have been Jeff Haynie, but was %s", result.Name)
	}
}

func TestCopyright4(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright 2010-2012 Jeff Haynie")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 2 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Dates[1] != "2012" {
		t.Fatalf("result date should be 2012, but was %s", result.Dates[0])
	}
	if result.Name != "Jeff Haynie" {
		t.Fatalf("result name should have been Jeff Haynie, but was %s", result.Name)
	}
}

func TestCopyright5(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright &copy; 2010 Jeff Haynie")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "Jeff Haynie" {
		t.Fatalf("result name should have been Jeff Haynie, but was %s", result.Name)
	}
}

func TestCopyright6(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright (c) 2010")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright7(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright © 2010")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright8(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright &#169; 2010")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright9(t *testing.T) {
	result, err := copyright.ParseCopyright("copyright &#xa9; 2010")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright10(t *testing.T) {
	result, err := copyright.ParseCopyright("// copyright 2010")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright11(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright 2010")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright12(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright 2010, 2012")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 2 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Dates[1] != "2012" {
		t.Fatalf("result date should be 2012, but was %s", result.Dates[1])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright13(t *testing.T) {
	result, err := copyright.ParseCopyright("COPYRIGHT 2010, 2012")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 2 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Dates[1] != "2012" {
		t.Fatalf("result date should be 2012, but was %s", result.Dates[1])
	}
	if result.Name != "" {
		t.Fatalf("result name should have been empty, but was %s", result.Name)
	}
}

func TestCopyright14(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright 2010 The Go Authors.  All rights reserved.")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2010" {
		t.Fatalf("result date should be 2010, but was %s", result.Dates[0])
	}
	if result.Name != "The Go Authors" {
		t.Fatalf("result name should have been \"The Go Authors\", but was %s", result.Name)
	}
}

func TestCopyright15(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright (C) 2012 by Nick Craig-Wood http://www.craig-wood.com/nick/")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2012" {
		t.Fatalf("result date should be 2012, but was %s", result.Dates[0])
	}
	if result.Name != "Nick Craig-Wood http://www.craig-wood.com/nick/" {
		t.Fatalf("result name should have been \"Nick Craig-Wood http://www.craig-wood.com/nick/\", but was %s", result.Name)
	}
}

func TestCopyright16(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright [yyyy] [name of copyright owner]")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found {
		t.Fatal("result was found but should not have been")
	}
}

func TestCopyright17(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright (c) 2014 Colin Marc (colinmarc@gmail.com)")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2014" {
		t.Fatalf("result date should be 2014, but was %s", result.Dates[0])
	}
	if result.Name != "Colin Marc (colinmarc@gmail.com)" {
		t.Fatalf("result name should have been \"Colin Marc (colinmarc@gmail.com)\", but was %s", result.Name)
	}
}

func TestCopyright18(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright (c) 2016 A Pinpoint PBC")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 1 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2016" {
		t.Fatalf("result date should be 2016, but was %s", result.Dates[0])
	}
	if result.Name != "A Pinpoint PBC" {
		t.Fatalf("result name should have been \"A Pinpoint PBC\", but was %s", result.Name)
	}
	str := result.String()
	expected := "Copyright © 2016 by A Pinpoint PBC. All Rights Reserved."
	if str != expected {
		t.Fatalf("expected \"%s\", was: \"%s\"", expected, str)
	}
}

func TestCopyright19(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright (c) 2016-2017 A Pinpoint PBC")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 2 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2016" {
		t.Fatalf("result date should be 2016, but was %s", result.Dates[0])
	}
	if result.Dates[1] != "2017" {
		t.Fatalf("result date should be 2017, but was %s", result.Dates[1])
	}
	if result.Name != "A Pinpoint PBC" {
		t.Fatalf("result name should have been \"A Pinpoint PBC\", but was %s", result.Name)
	}
	str := result.String()
	expected := "Copyright © 2016-2017 by A Pinpoint PBC. All Rights Reserved."
	if str != expected {
		t.Fatalf("expected \"%s\", was: \"%s\"", expected, str)
	}
}

func TestCopyright20(t *testing.T) {
	result, err := copyright.ParseCopyright("Copyright (c) 2016-2017 A Pinpoint PBC")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Found == false {
		t.Fatal("result not found")
	}
	if len(result.Dates) != 2 {
		t.Fatal("result date not found")
	}
	if result.Dates[0] != "2016" {
		t.Fatalf("result date should be 2016, but was %s", result.Dates[0])
	}
	if result.Dates[1] != "2017" {
		t.Fatalf("result date should be 2017, but was %s", result.Dates[1])
	}
	if result.Name != "A Pinpoint PBC" {
		t.Fatalf("result name should have been \"A Pinpoint PBC\", but was %s", result.Name)
	}
	str := result.HTMLString()
	expected := "Copyright &copy; 2016-2017 by A Pinpoint PBC. All Rights Reserved."
	if str != expected {
		t.Fatalf("expected \"%s\", was: \"%s\"", expected, str)
	}
}
