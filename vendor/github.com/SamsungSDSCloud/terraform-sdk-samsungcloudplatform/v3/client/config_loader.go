package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

// LoadDefaultConfigFromHome reads ~/.scp/config.json + ~/.scp/credentials.json
// and returns a populated *Configuration ready to be passed to per-service
// NewAPIClient constructors.
//
// Field-name mapping (kebab-case in JSON → PascalCase in Configuration):
//
//	config.json:        host → Host (via ServiceHost), user-id → UserId,
//	                    email → Email, project-id → ProjectId, login-id → LoginId
//	credentials.json:   auth-method → AuthMethod,
//	                    access-key → Credentials.AccessKey,
//	                    secret-key → Credentials.SecretKey
//
// The provider stores Host on Configuration.BasePath at NewDefaultConfig time
// (joined with the per-service path component). This loader returns a raw
// Configuration where BasePath == Host; callers that need a per-service path
// segment should clone and tack it on (or call ChangeBasePath later).
func LoadDefaultConfigFromHome() (*Configuration, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dir := filepath.Join(home, ".scp")
	cfg, err := loadConfig(filepath.Join(dir, "config.json"))
	if err != nil {
		return nil, err
	}
	creds, err := loadCredentials(filepath.Join(dir, "credentials.json"))
	if err != nil {
		return nil, err
	}
	cfg.Credentials = creds.toCredentials()
	cfg.AuthMethod = creds.AuthMethod
	cfg.Token = creds.Token
	cfg.HTTPClient = &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment, // OUR SDK respects HTTPS_PROXY/HTTP_PROXY
	}}
	cfg.UserAgent = "scpsdk-recovered/3.16.1"
	if cfg.BasePath == "" {
		cfg.BasePath = cfg.Host
	}
	return cfg, nil
}

type scpConfigJSON struct {
	Host      string `json:"host"`
	UserID    string `json:"user-id"`
	Email     string `json:"email"`
	ProjectID string `json:"project-id"`
	LoginID   string `json:"login-id"`
	Target    string `json:"target"`
}

type scpCredsJSON struct {
	AuthMethod string `json:"auth-method"`
	AccessKey  string `json:"access-key"`
	SecretKey  string `json:"secret-key"`
	Token      string `json:"token"`
}

func loadConfig(path string) (*Configuration, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s scpConfigJSON
	if err := json.Unmarshal(raw, &s); err != nil {
		return nil, err
	}
	if s.Host == "" {
		return nil, errors.New("scpsdk: config.json missing 'host'")
	}
	return &Configuration{
		Host:      s.Host,
		BasePath:  s.Host,
		ProjectId: s.ProjectID,
		UserId:    s.UserID,
		Email:     s.Email,
		LoginId:   s.LoginID,
	}, nil
}

func loadCredentials(path string) (*scpCredsJSON, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s scpCredsJSON
	if err := json.Unmarshal(raw, &s); err != nil {
		return nil, err
	}
	if s.AuthMethod == "" {
		s.AuthMethod = "access-key"
	}
	return &s, nil
}

func (c *scpCredsJSON) toCredentials() *Credentials {
	return &Credentials{AccessKey: c.AccessKey, SecretKey: c.SecretKey}
}

// CloneWithServicePath returns a copy of cfg with `/path` appended to BasePath.
// Mirrors the upstream provider's NewDefaultConfig pattern.
func (c *Configuration) CloneWithServicePath(servicePath string) *Configuration {
	clone := *c
	if servicePath != "" {
		clone.BasePath = c.Host + "/" + servicePath
	}
	return &clone
}
