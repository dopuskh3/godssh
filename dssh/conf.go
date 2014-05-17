package gos

import (
  "errors"
  "gopkg.in/yaml.v1"
  "io/ioutil"
  "os"
)

type Group struct {
  Keyfile  string
  Password string `,omitempty`
  User     string
  Hosts    []string
}

var _registry map[string]*Group

func LoadConfigFromFile(path string) error {
  fi, err := ioutil.ReadFile(os.ExpandEnv(path))
  if err != nil {
    return err
  }
  LoadConfig(fi)
  return nil
}

func LoadConfig(conf []byte) error {
  registry := make(map[string]*Group)
  err := yaml.Unmarshal(conf, &registry)
  if err != nil {
    return err
  }
  _registry = registry
  return nil
}

func GetGroup(group string) (*Group, error) {
  value, ok := _registry[group]
  if !ok {
    return nil, errors.New("Cannot find group")
  }
  return value, nil
}
