package rotakey

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestCredentialsMissing(t *testing.T) {
	file, _ := ioutil.TempFile(".", "missing-*")
	fname := file.Name()
	os.Remove(fname)
	_, err := credentialsFile(fname)
	if err == nil {
		t.Error("TestCredentialsMissing failed to notice missing file")
	}
	log.Warningf("TestCredentialsMissing :: err=%v", err)
	if !strings.Contains(err.Error(), " does not exist") {
		t.Errorf("TestCredentialsMissing returned wrong err (%s)", err.Error())
	}
}

func TestCredentialsReadMissing(t *testing.T) {
	fname := "TestCredentialsReadMissing"
	os.Create(fname)
	defer os.Remove(fname)
	os.Chmod(fname, 0222)
	_, err := credentialsFile(fname)
	if err == nil {
		t.Error("TestCredentialsReadMissing failed to notice missing read permission")
	}
	log.Infof("fname=%s, err=%v", fname, err)
	if !strings.Contains(err.Error(), "no read access to ") {
		t.Errorf("TestCredentialsReadMissing returned wrong err (%s)", err.Error())
	}
}

func TestCredentialsWriteMissing(t *testing.T) {
	fname := "TestCredentialsWriteMissing"
	os.Create(fname)
	defer os.Remove(fname)
	os.Chmod(fname, 0444)
	_, err := credentialsFile(fname)
	if err == nil {
		t.Error("TestCredentialsWriteMissing failed to notice missing write permission")
	}
	if !strings.Contains(err.Error(), "no write access to ") {
		t.Errorf("TestCredentialsWriteMissing returned wrong err (%s)", err.Error())
	}
}
