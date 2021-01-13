// Copyright 2020. Akamai Technologies, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terminal

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestWrite(t *testing.T) {
	out, err := ioutil.TempFile("", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(out.Name()) // clean up

	term := New(out, os.Stdin, os.Stderr)

	term.Write(t.Name())

	out.Seek(0, 0)

	data, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, t.Name(), string(data), "they should be equal")
}

func TestWriteErr(t *testing.T) {
	out, err := ioutil.TempFile("", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(out.Name()) // clean up

	term := New(os.Stdin, os.Stdin, out)

	term.WriteError(t.Name())

	out.Seek(0, 0)

	data, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, t.Name(), string(data), "they should be equal")
}

func TestPrompt(t *testing.T) {
	content := []byte("Tom\r\n")
	in, err := ioutil.TempFile("", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(in.Name()) // clean up

	if _, err := in.Write(content); err != nil {
		t.Fatal(err)
	}
	in.Seek(0, 0)

	term := New(os.Stdout, in, os.Stderr)

	name, err := term.Prompt("What is your name")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, name, "Tom", "they should be equal")
}

func TestPromptOptions(t *testing.T) {
	content := []byte("yellow\r\n")
	in, err := ioutil.TempFile("", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(in.Name()) // clean up

	if _, err := in.Write(content); err != nil {
		t.Fatal(err)
	}
	in.Seek(0, 0)

	term := New(os.Stdout, in, os.Stderr)

	color, err := term.Prompt("What is your favorite color", "yellow", "red", "blue")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, color, "yellow", "they should be equal")
}
