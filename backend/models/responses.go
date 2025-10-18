package models

type SearchResponse struct {
	Success   bool             `json:"success"`
	Equipment string           `json:"equipment"`
	Response  ResponseSections `json:"response"`
}

type ResponseSections struct {
	Analysis       string   `json:"analysis"`
	Causes         []string `json:"causes"`
	Steps          []string `json:"steps"`
	SafetyWarnings []string `json:"safety_warnings"`
	EquipmentNotes string   `json:"equipment_notes"`
}

type VCRAResponse struct {
	Success  bool        `json:"success"`
	Response VCRADetails `json:"response"`
}

type VCRADetails struct {
	RootCause  string   `json:"rootCause"`
	RiskLevel  string   `json:"riskLevel"`
	Confidence float64  `json:"confidence"`
	Actions    []string `json:"actions"`
	Timeline   []string `json:"timeline"`
}

type SafetyResponse struct {
	Success  bool          `json:"success"`
	Response SafetyDetails `json:"response"`
}

type SafetyDetails struct {
	HazardLevel  string         `json:"hazardLevel"`
	Hazards      []HazardDetail `json:"hazards"`
	Mitigations  []string       `json:"mitigations"`
	Standards    []string       `json:"standards"`
}

type HazardDetail struct {
	Name        string `json:"name"`
	Severity    string `json:"severity"`
	Probability string `json:"probability"`
}

type CorrosionResponse struct {
	Success  bool             `json:"success"`
	Response CorrosionDetails `json:"response"`
}

type CorrosionDetails struct {
	RiskLevel       string   `json:"riskLevel"`
	CorrosionRate   float64  `json:"corrosionRate"`
	Mechanisms      []string `json:"mechanisms"`
	Recommendations []string `json:"recommendations"`
	EstimatedLife   string   `json:"estimatedLife"`
}
