package dssh

import (
  "code.google.com/p/go.crypto/ssh"
  "github.com/gokyle/sshkey"
)

func LoadAuthMethod(g *Group) (ssh.AuthMethod, error) {
  rawKey, _, err := sshkey.LoadPrivateKeyFile(g.Keyfile)
  if err != nil {
    return nil, err
  }
  keySigner, err := ssh.NewSignerFromKey(rawKey)
  if err != nil {
    return nil, err
  }
  return ssh.PublicKeys(keySigner), nil
}

func LoadClientConfig(group string) (*ssh.ClientConfig, error) {
  g, err := GetGroup(group)
  if err != nil {
    return nil, err
  }
  auth, err := LoadAuthMethod(g)
  if err != nil {
    return nil, err
  }

  auths := []ssh.AuthMethod{
    auth,
  }

  clientConfig := &ssh.ClientConfig{
    User: g.User,
    Auth: auths,
  }
  return clientConfig, nil
}
