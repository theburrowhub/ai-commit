package configure

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

// Default values for the configuration
var (
	defaultLogLevel             = "info"
	defaultModel                = "mistral"
	defaultServer               = "http://localhost:11434"
	defaultRetriesCommitMessage = 3
	defaultSystemPrompt         = "In an impersonal way, write a commit message that explains what the commit\n" +
		"is for. Use conventional commits and the imperative mood in the first line.\n" +
		"The first line should start with: feat, fix, refactor, docs, style, build,\n" +
		"perf, ci, style, test or chore. Set the file name and the changes made in\n" +
		"the body. Only one subject line is allowed.\n\n" +
		"An example of commit message is:\n\n" +
		"feat(file or class): Add user authentication\n\n" +
		"- Implement user sign-up and login functionality\n" +
		"- Add password hashing for security\n" +
		"- Integrate with authentication API\n\n" +
		"Add line breaks to separate subject from body."
	defaultNumCtx      = 4096
	defaultTemperature = float32(0)
	defaultNumKeep     = 512

	configFilePath = os.Getenv("HOME") + "/.config/ai-commit/config.yaml" // Path to the configuration file
)

// Config represents the configuration for the ai-commit service
type Config struct {
	Noop                 *bool          `yaml:"noop,omitempty"`
	LogLevel             *string        `yaml:"logLevel,omitempty"`
	OllamaServer         *string        `yaml:"ollamaServer,omitempty"`
	Model                *string        `yaml:"model,omitempty"`
	SystemPrompt         *LiteralString `yaml:"systemPrompt,omitempty"`
	RetriesCommitMessage *int           `yaml:"retriesCommitMessage,omitempty"`
	NumCtx               *int           `yaml:"numCtx,omitempty"`
	NumKeep              *int           `yaml:"numKeep,omitempty"`
	Temperature          *float32       `yaml:"temperature,omitempty"`
}

// Cfg represents the configuration for the ai-commit service
var Cfg Config

// LoadConfig loads the configuration
func LoadConfig() error {
	err := loadAndUpdateConfigFile()
	if err != nil {
		return err
	}

	return nil
}

// loadAndUpdateConfigFile updates the config file with new fields and removes deprecated fields
func loadAndUpdateConfigFile() error {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		slog.Info("Config file not found. Creating config file...")
		err = os.WriteFile(configFilePath, []byte{}, 0644)
		if err != nil {
			return err
		}
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		return err
	}

	// Set default values for new fields
	setSomeDefaultValue := false

	setSomeDefaultValue = setDefault(&Cfg.Noop, false)
	setSomeDefaultValue = setDefault(&Cfg.LogLevel, defaultLogLevel) || setSomeDefaultValue
	setSomeDefaultValue = setDefault(&Cfg.OllamaServer, defaultServer) || setSomeDefaultValue
	setSomeDefaultValue = setDefault(&Cfg.Model, defaultModel) || setSomeDefaultValue
	setSomeDefaultValue = setDefault(&Cfg.RetriesCommitMessage, defaultRetriesCommitMessage) || setSomeDefaultValue
	setSomeDefaultValue = setDefault(&Cfg.SystemPrompt, LiteralString(defaultSystemPrompt)) || setSomeDefaultValue
	setSomeDefaultValue = setDefault(&Cfg.NumCtx, defaultNumCtx) || setSomeDefaultValue
	setSomeDefaultValue = setDefault(&Cfg.NumKeep, defaultNumKeep) || setSomeDefaultValue
	setSomeDefaultValue = setDefault(&Cfg.Temperature, defaultTemperature) || setSomeDefaultValue

	if setSomeDefaultValue {
		updatedData, err := yaml.Marshal(&Cfg)
		if err != nil {
			return err
		}

		err = os.WriteFile(configFilePath, updatedData, 0644)
		if err != nil {
			return err
		}

		slog.Info("Config file updated successfully.")
	}

	return nil
}

// setDefault sets the default value for a field if it is nil
// It returns true if the field was set to the default value
func setDefault[T any](field **T, defaultValue T) bool {
	if *field == nil {
		*field = new(T)
		**field = defaultValue

		return true
	}

	return false
}
