package configure

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"path/filepath"
)

// Default values for the configuration
const (
	defaultNoopValue            = false
	defaultLogLevel             = "info"
	defaultModel                = "mistral"
	defaultServer               = "http://localhost:11434"
	defaultRetriesCommitMessage = 3
	defaultPromptTemplate       = `Commit changes:
{{ .Diff }}

In an impersonal way, write a commit message that explains what the commit is for. Use conventional commits
and the imperative mood in the first line. The first line should start with: feat, fix, refactor, docs, style,
build, perf, ci, style, test or chore. Set the file name and the changes made in the body. Only one subject
line is allowed. An example of commit message is:

feat(file or class): Add user authentication

- Implement user sign-up and login functionality
- Add password hashing for security
- Integrate with authentication API

Add line breaks to separate subject from body.`
)

// Config represents the configuration for the ai-commit service
type Config struct {
	Noop                        bool          `yaml:"noop,omitempty"`
	LogLevel                    string        `yaml:"logLevel,omitempty"`
	OllamaServer                string        `yaml:"ollamaServer,omitempty"`
	Model                       string        `yaml:"model,omitempty"`
	DefaultPromptTemplate       LiteralString `yaml:"defaultPromptTemplate,omitempty"`
	DefaultRetriesCommitMessage int           `yaml:"defaultRetriesCommitMessage,omitempty"`
}

// Cfg represents the configuration for the ai-commit service
var Cfg Config

var (
	configFilePath = os.Getenv("HOME") + "/.config/ai-commit/config.yaml" // Path to the configuration file
)

// LoadConfig loads the configuration
func LoadConfig() error {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		slog.Info("Config file not found. Creating config file...")
		err := createConfigFile()
		if err != nil {
			return err
		}
	}

	err := loadConfigFile()
	if err != nil {
		return err
	}

	return nil
}

// loadConfigFile loads the configuration from a yaml file
func loadConfigFile() error {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		return err
	}

	return nil
}

// createConfigFile creates a default config file
func createConfigFile() error {
	defaultConfig := Config{
		Noop:                        defaultNoopValue,
		LogLevel:                    defaultLogLevel,
		OllamaServer:                defaultServer,
		Model:                       defaultModel,
		DefaultRetriesCommitMessage: defaultRetriesCommitMessage,
		DefaultPromptTemplate:       LiteralString(defaultPromptTemplate),
	}

	data, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		return err
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(configFilePath), 0755)
		if err != nil {
			return err
		}
		err = os.WriteFile(configFilePath, data, 0644)
		if err != nil {
			return err
		} else {
			slog.Info("Config file created successfully.")
		}
	} else {
		slog.Warn("Config file already exists.")
	}

	return nil
}
