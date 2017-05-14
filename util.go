// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/FactomProject/factom"
	"errors"
)

// exidCollector accumulates the external ids from the command line -e and -E
// flags
var exidCollector [][]byte

// extids will be a flag receiver for adding chains and entries
// In ASCII
type extidsASCII []string

func (e *extidsASCII) String() string {
	return fmt.Sprint(*e)
}

func (e *extidsASCII) Set(s string) error {
	*e = append(*e, s)
	exidCollector = append(exidCollector[:], []byte(s))
	return nil
}

// extids will be a flag receiver for adding chains and entries
// In HEX
type extidsHex []string

func (e *extidsHex) String() string {
	return fmt.Sprint(*e)
}

func (e *extidsHex) Set(s string) error {
	*e = append(*e, s)
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	exidCollector = append(exidCollector[:], b)
	return nil
}

// nameCollector accumulates the components of a chain name from the command
// line -n and -N flags
var nameCollector [][]byte

// namesASCII will be a flag receiver for ASCII chain names.
type namesASCII []string

func (n *namesASCII) String() string {
	return fmt.Sprint(*n)
}

func (n *namesASCII) Set(s string) error {
	*n = append(*n, s)
	nameCollector = append(nameCollector[:], []byte(s))
	return nil
}

// namesHex will be a flag receiver for HEX encoded chain names.
type namesHex []string

func (n *namesHex) String() string {
	return fmt.Sprint(*n)
}

func (n *namesHex) Set(s string) error {
	*n = append(*n, s)
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	nameCollector = append(nameCollector[:], b)
	return nil
}

func factoshiToFactoid(v interface{}) string {
	value, err := strconv.Atoi(fmt.Sprint(v))
	if err != nil {
		return ""
	}
	d := value / 1e8
	r := value % 1e8
	ds := fmt.Sprintf("%d", d)
	rs := fmt.Sprintf("%08d", r)
	rs = strings.TrimRight(rs, "0")
	if len(rs) > 0 {
		ds = ds + "."
	}
	return fmt.Sprintf("%s%s", ds, rs)
}

// nametoid computes a chainid from the chain name components
func nametoid(name [][]byte) string {
	hs := sha256.New()
	for _, v := range name {
		h := sha256.Sum256(v)
		hs.Write(h[:])
	}
	return hex.EncodeToString(hs.Sum(nil))
}

// waitOnFctAck blocks while waiting for a factom ack message and returns the
// ack status or times out after 10 seconds.
func waitOnFctAck(txid string) (string, error) {
		for i :=0; i < 120; i++ {
			s, err := factom.FactoidACK(txid, "")
			if err != nil {
				time.Sleep(time.Second / 2)
				continue
			}
			if (s.Status != "Unknown") && (s.Status != "NotConfirmed") {
				return s.Status, nil
			}
			time.Sleep(time.Second / 2)
		}

		return "", errors.New("timeout: no acknowledgement found")
}

// waitOnCommitAck blocks while waiting for an ack message for an Entry Commit
// and returns the ack status or times out after 10 seconds.
func waitOnCommitAck(txid string) (string, error) {
	for i :=0; i < 120; i++ {
			s, err := factom.EntryACK(txid, "")
			if err != nil {
				time.Sleep(time.Second / 2)
				continue
			}
			if (s.CommitData.Status != "Unknown") && (s.CommitData.Status != "NotConfirmed") {
				return s.CommitData.Status, nil
			}
			time.Sleep(time.Second / 2)
		}
		return "", errors.New("timeout: no acknowledgement found")
}

// waitOnRevealAck blocks while waiting for an ack message for an Entry Reveal
// and returns the ack status or times out after 10 seconds.
func waitOnRevealAck(txid string) (string, error) {
	for i :=0; i < 120; i++ {
			s, err := factom.EntryACK(txid, "")
			if err != nil {
				time.Sleep(time.Second / 2)
				continue
			}
			if (s.EntryData.Status != "Unknown") && (s.EntryData.Status != "NotConfirmed") {
				return s.EntryData.Status, nil
			}
			time.Sleep(time.Second / 2)
		}
		return "", errors.New("timeout: no acknowledgement found")
}
