// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os/exec"
)

func pj(v interface{}) {
	buf, err := json.MarshalIndent(v, "", "    ")
	ce(err)
	p(string(buf))
}

func ce(err error) {
	if err != nil {
		panic(err)
	}
}

func px(v interface{}) {
	buf, err := xml.MarshalIndent(v, "", "    ")
	ce(err)
	fmt.Println(string(buf))
}

func format(file string) {
	cmd := exec.Command("go", "fmt", file)
	err := cmd.Start()
	ce(err)
	err = cmd.Wait()
	ce(err)
}

func fp(w io.Writer, v ...interface{}) {
	fmt.Fprint(w, v...)
	fmt.Fprintln(w)
}

func c(err error) {
	if err != nil {
		panic(err)
	}
}

func check(ch <-chan error) {
	for err := range ch {
		c(err)
	}
}
