package internal

import (
	"encoding/json"
	"runtime"
	"testing"
)

func TestMarshalEnvironment(t *testing.T) {
	js, err := json.Marshal(&sampleEnvironment)
	if nil != err {
		t.Fatal(err)
	}
	expect := `[["Compiler","comp"],["GOARCH","arch"],["GOOS","goos"],["Version","vers"]]`
	if string(js) != expect {
		t.Fatal(string(js))
	}
}

func TestEnvironmentFields(t *testing.T) {
	env := newEnvironment()
	if env.Compiler != runtime.Compiler {
		t.Error(env.Compiler, runtime.Compiler)
	}
	if env.GOARCH != runtime.GOARCH {
		t.Error(env.GOARCH, runtime.GOARCH)
	}
	if env.GOOS != runtime.GOOS {
		t.Error(env.GOOS, runtime.GOOS)
	}
	if env.Version != runtime.Version() {
		t.Error(env.Version, runtime.Version())
	}
}
