package main

import "github.com/google/uuid"

// Config holds all configuration values
type Config struct {
	//Api holds all API configuration values
	Api API
	//Project holds all project configuration values
	Project Project
	// BomPath holds the path to the BOM to process
	BomPath string
	// ViolationThreshold holds the rules for each threshold type
	ViolationThresholds ViolationThresholds
}

// API holds all API configuration values
type API struct {
	// Url holds the base URL for all API calls
	Url string
	// Token holds the API token requirec for Authentication/Authorization
	Token string
	// Timeout holds the API call timeout in seconds
	Timeout int
}

// Project holds all project configuration values
type Project struct {
	// Name holds the project name to send BOM to
	Name string
	// Version holds the project version to send BOM to
	Version string
	// AutoCreate toggles automatic project creation if project does not exist, requires PORTFOLIO_MANAGEMENT permission
	AutoCreate bool
	//Uuid holds the retrieved project UUID from dependency-track
	Uuid *uuid.UUID
}

// ViolationThreshold holds the rules for each threshold type
type ViolationThresholds struct {
	// Critical holds the threshold on which critical violations fail the analysis
	Critical int
	// High holds the threshold on which high violations fail the analysis
	High int
	// Medium holds the threshold on which medium violations fail the analysis
	Medium int
	// Low holds the threshold on which low violations fail the analysis
	Low int
	// Unassigned holds the threshold on which unassigned violations fail the analysis
	Unassigned int
}

// AnalysisResult holds information about an analysis result
type AnalysisResult struct {
	// Passed determines whether the analysis passed according to the entered rules
	Passed bool
	// CriticalVulnerabilities counts the number of critical vulnerabilities found in the BOM
	CriticalVulnerabilities int
	// HighVulnerabilities counts the number of high vulnerabilities found in the BOM
	HighVulnerabilities int
	// MediumVulnerabilities counts the number of medium vulnerabilities found in the BOM
	MediumVulnerabilities int
	// LowVulnerabilities counts the number of low vulnerabilities found in the BOM
	LowVulnerabilities int
	// UnassignedVulnerabilities counts the number of unassigned vulnerabilities found in the BOM
	UnassignedVulnerabilities int
}
