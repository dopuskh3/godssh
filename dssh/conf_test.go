package dssh

import (
  "github.com/bmizerany/assert"
  "io/ioutil"
  "syscall"
  "testing"
)

var sample_config = `
group1:
  keyfile: /foo/var
  user: francois
  password: password
  hosts:
    - h1
    - h2
    - h3
group2:
  user: francois
  keyfile: /foo/bar
  hosts:
    - h1
    - h2
`

func unsetConfig() {
  _registry = nil
}

func TestLoadConfig(t *testing.T) {
  defer unsetConfig()
  LoadConfig([]byte(sample_config))
}

func TestGetExistingGroup(t *testing.T) {
  LoadConfig([]byte(sample_config))
  defer unsetConfig()
  g, err := GetGroup("group1")
  assert.NotEqual(t, g, nil)
  assert.Equal(t, err, nil)
  assert.Equal(t, g.User, "francois")
  assert.Equal(t, g.Password, "password")
  assert.Equal(t, g.Hosts, []string{"h1", "h2", "h3"})
  t.Logf("Loaded :%#v", g)
}

func TestLoadConfigFromFile(t *testing.T) {
  f, err := ioutil.TempFile("", "config")
  defer syscall.Unlink(f.Name())
  defer unsetConfig()
  if err != nil {
    panic(err)
  }
  ioutil.WriteFile(f.Name(), []byte(sample_config), 0644)
  LoadConfigFromFile(f.Name())
  assert.NotEqual(t, _registry, nil)
}

func TestLoadConfigFromNonExistingFile(t *testing.T) {
  err := LoadConfigFromFile("non-existent-file")
  assert.NotEqual(t, err, nil)
}

func TestLoadInvalidConfig(t *testing.T) {
  invalidConfig := `
{ 
  "asdaidsasd
asdasdasda--d-adasd
  `
  err := LoadConfig([]byte(invalidConfig))
  assert.NotEqual(t, err, nil)
}

func TestGetNonExistingGroup(t *testing.T) {
  LoadConfig([]byte(sample_config))
  defer unsetConfig()
  _, err := GetGroup("Do-not-exists")
  assert.NotEqual(t, err, nil)
}
