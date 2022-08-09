package configs

import (
	"fmt"
	"log"
	"musement/src/internal/configs/environments"
	"musement/src/internal/core/domain/enums"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Scope     enums.Scope
	EnvConfig environments.EnvironmentConfig
}

func NewConfig() *Config {
	instanceConfig := NewInstanceConfig()
	scope := setupScope(instanceConfig.scope)
	envConfig := setupEnvironmentConfig([]byte(environments.Development), []byte(findConfigStringByScope(scope)))

	return &Config{
		Scope:     scope,
		EnvConfig: *envConfig,
	}
}

func setupScope(instanceScope string) enums.Scope {
	fmt.Println("Setting Scope")
	var currentScope enums.Scope
	if strings.EqualFold(instanceScope, "dev") {
		currentScope = enums.DevScope
	} else {
		currentScope = enums.ProdScope
	}
	return currentScope
}

func applyEnvironmentConfigFromString(configBytes []byte) (*environments.EnvironmentConfig, error) {
	environmentConfig := &environments.EnvironmentConfig{}
	err := yaml.Unmarshal(configBytes, environmentConfig)
	return environmentConfig, err
}

func setupEnvironmentConfig(defaultEnvConfigBytes, scopeEnvConfigBytes []byte) *environments.EnvironmentConfig {
	envCfg, err := applyEnvironmentConfigFromString(defaultEnvConfigBytes)
	if err != nil {
		log.Panicf("error with default configs: '%v'", err)
	}

	err = yaml.Unmarshal(scopeEnvConfigBytes, envCfg)
	if err != nil {
		log.Panicf("error with Scope configs: '%v'", err)
	}

	return envCfg
}

func findConfigStringByScope(scope enums.Scope) string {
	var configsByScopeMap = map[enums.Scope]string{
		enums.DevScope:  environments.Development,
		enums.ProdScope: environments.Production,
	}

	configString, ok := configsByScopeMap[scope]
	if !ok {
		log.Panicf("there is no configuration for Scope '%s'", scope)
	}

	return configString
}
