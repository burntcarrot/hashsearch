package config

import (
	"errors"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var (
	k = koanf.New(".")

	DB_URL string // Database URL.

	SERVER_ADDR string // Server address.

	FILES_DIR string // Directory where the images would be saved.

	CORS_ALLOW_ORIGIN string // CORS Allow-Origin value.

	LOGGING_FILE string // Log file path.
)

// LoadConfig loads config from file and environment variables.
func LoadConfig(configPath string) error {
	// Load config from YAML file.
	err := k.Load(file.Provider(configPath), yaml.Parser())
	if err != nil {
		return err
	}

	// Load config from environment variables.
	err = k.Load(env.Provider("PERHASH_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "PERHASH_")), "_", ".", -1)
	}), nil)
	if err != nil {
		return err
	}

	// Read values.

	// Read database config values.
	DB_URL = k.String("db.url")

	// Read database config values.
	SERVER_ADDR = k.String("server.addr")

	// Read database config values.
	FILES_DIR = k.String("files.dir")

	// Read database config values.
	CORS_ALLOW_ORIGIN = k.String("cors.allow_origin")

	// Read database config values.
	LOGGING_FILE = k.String("logging.file")

	return nil
}

// Validate validates the config values.
func Validate() error {
	if DB_URL == "" {
		return errors.New("no database url provided")
	}

	if SERVER_ADDR == "" {
		return errors.New("no server address provided")
	}

	if FILES_DIR == "" {
		return errors.New("no files directory provided")
	}

	if CORS_ALLOW_ORIGIN == "" {
		return errors.New("no CORS allow origin provided")
	}

	if LOGGING_FILE == "" {
		return errors.New("no logging file provided")
	}

	return nil
}
