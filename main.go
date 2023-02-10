package main

import (
	"encoding/json"
	"log"

	dtrack "github.com/DependencyTrack/client-go"
)

func main() {
	log.Println("dtrack-auditor-go started...")
	log.Println("validating config before run...")

	config := validateConfig()

	log.Println("config is valid, creating dependency-track client...")

	client, err := createDtrackClient(config)

	log.Println("client created successfully, retrieving project info...")

	if err != nil {
		log.Fatalf("Error creating dependency-track client: %v", err)
	}

	project, err := getProjectInfo(config, client)

	if err != nil {
		if err.(*dtrack.APIError).StatusCode == 404 {
			if !config.Project.AutoCreate {
				log.Fatalf("Error retrieving project info and project autocreate is not enabled: %v", err)
			} else {
				log.Println("Project does not exist, will autocreate during upload...")
			}
		} else {
			log.Fatalf("Error retrieving project info: %v", err)
		}
	} else {
		config.Project.Uuid = &project.UUID
	}

	log.Println("project info loaded successfully or project to be created during upload, processing BOM...")

	doneChan, errChan := sendBOM(config, client)

	select {
	case <-doneChan:
		log.Println("BOM processed successfully, priting thresholds...")

		log.Printf("Thresholds: %v", prettyPrint(config.ViolationThresholds))

		log.Println("Analyzing results...")

		result, err := analyzeResults(config, client)

		if err != nil {
			log.Fatalf("Error analyzing results: %v", err)
		}

		if !result.Passed {
			log.Fatalf("Analysis failed, details: %v", prettyPrint(result))
		} else {
			log.Printf("Analysis passed, details: %v", prettyPrint(result))
		}

	case err := <-errChan:
		log.Fatalf("Error processing BOM: %v", err)
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
