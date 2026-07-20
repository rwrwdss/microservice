package config

import "testing"

func TestLoadUsesEnvOverrides(t *testing.T) {
	t.Setenv("DB_HOST", "db.internal")
	t.Setenv("DB_PORT", "5433")
	t.Setenv("DB_USER", "service")
	t.Setenv("DB_PASSWORD", "secret")
	t.Setenv("DB_NAME", "appdb")
	t.Setenv("DB_SSLMODE", "require")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if cfg.DBHost != "db.internal" {
		t.Fatalf("expected DBHost to be db.internal, got %s", cfg.DBHost)
	}

	if cfg.DBPort != "5433" {
		t.Fatalf("expected DBPort to be 5433, got %s", cfg.DBPort)
	}

	if cfg.DBUser != "service" {
		t.Fatalf("expected DBUser to be service, got %s", cfg.DBUser)
	}

	if cfg.DBPassword != "secret" {
		t.Fatalf("expected DBPassword to be secret, got %s", cfg.DBPassword)
	}

	if cfg.DBName != "appdb" {
		t.Fatalf("expected DBName to be appdb, got %s", cfg.DBName)
	}

	if cfg.DBSSLMode != "require" {
		t.Fatalf("expected DBSSLMode to be require, got %s", cfg.DBSSLMode)
	}
}
