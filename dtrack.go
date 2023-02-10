package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	dtrack "github.com/DependencyTrack/client-go"
)

func createDtrackClient(config *Config) (*dtrack.Client, error) {
	return dtrack.NewClient(config.Api.Url, dtrack.WithAPIKey(config.Api.Token), dtrack.WithTimeout(time.Duration(config.Api.Timeout)))
}

func getProjectInfo(config *Config, client *dtrack.Client) (dtrack.Project, error) {
	return client.Project.Lookup(context.TODO(), config.Project.Name, config.Project.Version)
}

func sendBOM(config *Config, client *dtrack.Client) (<-chan bool, <-chan error) {
	bomContent, e := os.ReadFile(config.BomPath)

	if e != nil {
		log.Fatalf("Error reading BOM contents: %s", e)
	}

	var uploadToken dtrack.BOMUploadToken

	if config.Project.Uuid == nil {
		uploadToken, e = client.BOM.Upload(context.TODO(), dtrack.BOMUploadRequest{
			ProjectName:    config.Project.Name,
			ProjectVersion: config.Project.Version,
			AutoCreate:     true,
			BOM:            base64.StdEncoding.EncodeToString(bomContent),
		})
	} else {
		uploadToken, e = client.BOM.Upload(context.TODO(), dtrack.BOMUploadRequest{
			ProjectUUID: config.Project.Uuid,
			BOM:         base64.StdEncoding.EncodeToString(bomContent),
		})
	}

	if e != nil {
		log.Fatalf("Error sending BOM to dependency-track: %s", e)
	}

	var (
		done    = make(chan bool)
		err     = make(chan error)
		ticker  = time.NewTicker(1 * time.Second)
		timeout = time.After(time.Duration(config.Api.Timeout) * time.Second)
	)

	go func() {
		defer func() {
			close(done)
			close(err)
		}()

		for {
			select {
			case <-ticker.C:
				processing, e := client.BOM.IsBeingProcessed(context.TODO(), uploadToken)
				if e != nil {
					err <- e
					return
				}
				if !processing {
					project, _ := client.Project.Lookup(context.TODO(), config.Project.Name, config.Project.Version)

					config.Project.Uuid = &project.UUID

					done <- true
					return
				}
			case <-timeout:
				err <- fmt.Errorf("timeout exceeded")
				return
			}
		}
	}()

	return done, err
}

func analyzeResults(config *Config, client *dtrack.Client) (AnalysisResult, error) {
	metrics, err := client.Metrics.LatestProjectMetrics(context.TODO(), *config.Project.Uuid)

	if err != nil {
		// dependency-track takes a while to process metrics for a new project, we wait for 10 seconds before trying again
		if err.Error() == "EOF" {
			time.Sleep(10 * time.Second)
			metrics, err = client.Metrics.LatestProjectMetrics(context.TODO(), *config.Project.Uuid)
		} else {
			return AnalysisResult{}, err
		}
	}

	result := AnalysisResult{
		Passed: true,
	}

	result.CriticalVulnerabilities = metrics.Critical
	result.HighVulnerabilities = metrics.High
	result.MediumVulnerabilities = metrics.Medium
	result.LowVulnerabilities = metrics.Low
	result.UnassignedVulnerabilities = metrics.Unassigned

	if config.ViolationThresholds.Critical != 0 && (result.CriticalVulnerabilities >= config.ViolationThresholds.Critical) {
		result.Passed = false
	}

	if config.ViolationThresholds.High != 0 && (result.HighVulnerabilities >= config.ViolationThresholds.High) {
		result.Passed = false
	}

	if config.ViolationThresholds.Medium != 0 && (result.MediumVulnerabilities >= config.ViolationThresholds.Medium) {
		result.Passed = false
	}

	if config.ViolationThresholds.Low != 0 && (result.LowVulnerabilities >= config.ViolationThresholds.Low) {
		result.Passed = false
	}

	if config.ViolationThresholds.Unassigned != 0 && (result.UnassignedVulnerabilities >= config.ViolationThresholds.Unassigned) {
		result.Passed = false
	}

	return result, err
}
