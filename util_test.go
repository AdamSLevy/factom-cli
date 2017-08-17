// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"testing"
)

func TestGettersAndSetters(t *testing.T) {
	var eHex extidsHex
	testHex := "761dc23d4fdd8be932107d6a03cead5a3edb69685c8b1f6641fb98608b22b4d6"
	err := eHex.Set(testHex)
	if err != nil {
		t.Error(err)
	}

	eHexString := eHex.String()
	if eHexString != fmt.Sprintf("[%s]", testHex) {
		t.Errorf("Hex string does not match: expected %s, got %s", eHexString, testHex)
	}

	var namesText namesASCII
	testNamesASCII := "123test"
	err = namesText.Set(testNamesASCII)
	if err != nil {
		t.Error(err)
	}

	namesTextString := namesText.String()
	if namesTextString != fmt.Sprintf("[%s]", testNamesASCII) {
		t.Errorf("Hex string does not match: expected %s, got %s", namesTextString, testNamesASCII)
	}

	var namesHexText namesHex
	testNamesHex := "761dc23d4fdd8be932107d6a03cead5a3edb69685c8b1f6641fb98608b22b4d6"
	err = namesHexText.Set(testNamesHex)
	if err != nil {
		t.Error(err)
	}

	namesHexTextString := namesHexText.String()
	if namesHexTextString != fmt.Sprintf("[%s]", testNamesHex) {
		t.Errorf("Hex string does not match: expected %s, got %s", namesHexTextString, testNamesHex)
	}
}
