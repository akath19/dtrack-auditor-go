package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func validateConfig() *Config {
	apiUrl, err := getEnvVar("API_URL")

	if err != nil {
		log.Fatalf("%s", err)
	}

	apiToken, err := getEnvVar("API_TOKEN")

	if err != nil {
		log.Fatalf("%s", err)
	}

	timeoutEnv, err := getEnvVar("API_TIMEOUT")

	if err != nil {
		log.Fatalf("%s", err)
	}

	timeout, err := strconv.Atoi(timeoutEnv)

	if err != nil {
		log.Fatalf("Error parsing timeout variable: %s", err)
	}

	projectName, err := getEnvVar("PROJECT_NAME")

	if err != nil {
		log.Fatalf("%s", err)
	}

	projectVersion, err := getEnvVar("PROJECT_VERSION")

	if err != nil {
		log.Fatalf("%s", err)
	}

	autoCreateEnv, err := getEnvVar("PROJECT_AUTOCREATE")

	var autoCreate bool

	if err != nil {
		autoCreate = false
	} else {
		autoCreate, err = strconv.ParseBool(autoCreateEnv)

		if err != nil {
			log.Fatalf("Error parsing autoCreate variable: %s", err)
		}
	}

	bomPath, err := getEnvVar("BOM_PATH")

	if err != nil {
		log.Fatalf("%s", err)
	}

	criticalEnv, err := getEnvVar("CRITICAL_THRESHOLD")

	if err != nil {
		log.Fatalf("%s", err)
	}

	critical, err := strconv.Atoi(criticalEnv)

	if err != nil {
		log.Fatalf("Error parsing critical threshold variable: %s", err)
	}

	highEnv, err := getEnvVar("HIGH_THRESHOLD")

	if err != nil {
		log.Fatalf("%s", err)
	}

	high, err := strconv.Atoi(highEnv)

	if err != nil {
		log.Fatalf("Error parsing high threshold variable: %s", err)
	}

	mediumEnv, err := getEnvVar("MEDIUM_THRESHOLD")

	if err != nil {
		log.Fatalf("%s", err)
	}

	medium, err := strconv.Atoi(mediumEnv)

	if err != nil {
		log.Fatalf("Error parsing medium threshold variable: %s", err)
	}

	lowEnv, err := getEnvVar("LOW_THRESHOLD")

	if err != nil {
		log.Fatalf("%s", err)
	}

	low, err := strconv.Atoi(lowEnv)

	if err != nil {
		log.Fatalf("Error parsing low threshold variable: %s", err)
	}

	unassignedEnv, err := getEnvVar("UNASSIGNED_THRESHOLD")

	if err != nil {
		log.Fatalf("%s", err)
	}

	unassigned, err := strconv.Atoi(unassignedEnv)

	if err != nil {
		log.Fatalf("Error parsing unassigned threshold variable: %s", err)
	}

	return &Config{
		Api: API{
			Url:     apiUrl,
			Token:   apiToken,
			Timeout: timeout,
		},
		Project: Project{
			Name:       projectName,
			Version:    projectVersion,
			AutoCreate: autoCreate,
		},
		BomPath: bomPath,
		ViolationThresholds: ViolationThresholds{
			Critical:   critical,
			High:       high,
			Medium:     medium,
			Low:        low,
			Unassigned: unassigned,
		},
	}
}

// getEnvVar checks for the existence of an env variable, returns an error if not found or empty
func getEnvVar(name string) (string, error) {
	val := os.Getenv(name)

	if val == "" {
		return "", fmt.Errorf("no %s env variable found", name)
	}

	return val, nil
}
