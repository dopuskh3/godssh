package gos

import (
  "github.com/bmizerany/assert"
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

func TestLoadConfig(t *testing.T) {
  LoadConfig([]byte(sample_config))
}

func TestGetExistingGroup(t *testing.T) {
  g, err := GetGroup("group1")
  assert.NotEqual(t, g, nil)
  assert.Equal(t, err, nil)
  assert.Equal(t, g.User, "francois")
  assert.Equal(t, g.Password, "password")
  assert.Equal(t, g.Hosts, []string{"h1", "h2", "h3"})
  t.Logf("Loaded :%#v", g)
}

func TestGetNonExistingGroup(t *testing.T) {
  _, err := GetGroup("Do-not-exists")
  assert.NotEqual(t, err, nil)
}
