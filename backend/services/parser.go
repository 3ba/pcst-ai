package services

import (
	"regexp"
	"strconv"
	"strings"
)

// TODO: make this readable
func ParseTroubleshootingResponse(text string) (string, []string, []string, []string, string) {
	var analysis string
	var causes []string
	var steps []string
	var safetyWarnings []string
	var equipmentNotes string

	lines := strings.Split(text, "\n")
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		upperLine := strings.ToUpper(line)

		if strings.Contains(upperLine, "ANALYSIS:") {
			currentSection = "analysis"
			analysis = strings.TrimSpace(strings.Replace(line, "ANALYSIS:", "", 1))
			continue
		} else if strings.Contains(upperLine, "POSSIBLE CAUSES:") {
			currentSection = "causes"
			continue
		} else if strings.Contains(upperLine, "TROUBLESHOOTING STEPS:") {
			currentSection = "steps"
			continue
		} else if strings.Contains(upperLine, "SAFETY WARNINGS:") {
			currentSection = "safety_warnings"
			continue
		} else if strings.Contains(upperLine, "EQUIPMENT NOTES:") {
			currentSection = "equipment_notes"
			continue
		}

		if strings.HasPrefix(line, "1.") || strings.HasPrefix(line, "2.") ||
			strings.HasPrefix(line, "3.") || strings.HasPrefix(line, "4.") ||
			strings.HasPrefix(line, "5.") || strings.HasPrefix(line, "6.") ||
			strings.HasPrefix(line, "7.") || strings.HasPrefix(line, "8.") ||
			strings.HasPrefix(line, "9.") {
			if currentSection == "causes" {
				causes = append(causes, line)
			} else if currentSection == "steps" {
				steps = append(steps, line)
			}
		} else if strings.HasPrefix(line, "-") && currentSection == "safety_warnings" {
			safetyWarnings = append(safetyWarnings, line)
		} else if currentSection == "equipment_notes" && line != "" {
			if equipmentNotes != "" {
				equipmentNotes += " "
			}
			equipmentNotes += line
		} else if currentSection == "analysis" && analysis == "" {
			analysis = line
		}
	}

	return analysis, causes, steps, safetyWarnings, equipmentNotes
}

func ParseVCRAResponse(text string) (string, string, float64, []string, []string) {
	rootCause := ""
	riskLevel := "MEDIUM"
	confidence := 0.75
	var actions []string
	var timeline []string

	lines := strings.Split(text, "\n")
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		upperLine := strings.ToUpper(line)

		if strings.Contains(upperLine, "ROOT CAUSE:") {
			currentSection = "root"
			rootCause = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			continue
		} else if strings.Contains(upperLine, "RISK LEVEL:") {
			currentSection = ""
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				level := strings.ToUpper(strings.TrimSpace(parts[1]))
				if strings.Contains(level, "HIGH") {
					riskLevel = "HIGH"
				} else if strings.Contains(level, "LOW") {
					riskLevel = "LOW"
				} else {
					riskLevel = "MEDIUM"
				}
			}
			continue
		} else if strings.Contains(upperLine, "CONFIDENCE:") {
			currentSection = ""
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				confStr := strings.TrimSpace(parts[1])
				confStr = strings.Trim(confStr, "%")
				if val, err := strconv.ParseFloat(confStr, 64); err == nil {
					if val > 1 {
						confidence = val / 100
					} else {
						confidence = val
					}
				}
			}
			continue
		} else if strings.Contains(upperLine, "IMMEDIATE ACTIONS:") || strings.Contains(upperLine, "RECOMMENDED ACTIONS:") {
			currentSection = "actions"
			continue
		} else if strings.Contains(upperLine, "TIMELINE:") || strings.Contains(upperLine, "RECOVERY TIMELINE:") {
			currentSection = "timeline"
			continue
		}

		if currentSection == "root" && rootCause == "" {
			rootCause = line
		} else if currentSection == "actions" && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•") || regexp.MustCompile(`^\d+\.`).MatchString(line)) {
			actions = append(actions, line)
		} else if currentSection == "timeline" && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•") || regexp.MustCompile(`^\d+\.`).MatchString(line)) {
			timeline = append(timeline, line)
		}
	}

	return rootCause, riskLevel, confidence, actions, timeline
}

func ParseSafetyResponse(text string) (string, []map[string]string, []string, []string) {
	hazardLevel := "MEDIUM"
	var hazards []map[string]string
	var mitigations []string
	var standards []string

	lines := strings.Split(text, "\n")
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		upperLine := strings.ToUpper(line)

		if strings.Contains(upperLine, "HAZARD LEVEL:") || strings.Contains(upperLine, "OVERALL RISK:") {
			currentSection = ""
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				level := strings.ToUpper(strings.TrimSpace(parts[1]))
				if strings.Contains(level, "HIGH") {
					hazardLevel = "HIGH"
				} else if strings.Contains(level, "LOW") {
					hazardLevel = "LOW"
				} else {
					hazardLevel = "MEDIUM"
				}
			}
			continue
		} else if strings.Contains(upperLine, "IDENTIFIED HAZARDS:") || strings.Contains(upperLine, "HAZARDS:") {
			currentSection = "hazards"
			continue
		} else if strings.Contains(upperLine, "MITIGATIONS:") || strings.Contains(upperLine, "RECOMMENDED MITIGATIONS:") {
			currentSection = "mitigations"
			continue
		} else if strings.Contains(upperLine, "STANDARDS:") || strings.Contains(upperLine, "RELEVANT STANDARDS:") {
			currentSection = "standards"
			continue
		}

		if currentSection == "hazards" && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•") || regexp.MustCompile(`^\d+\.`).MatchString(line)) {
			hazard := map[string]string{
				"name":        line,
				"severity":    "Medium",
				"probability": "Medium",
			}
			if strings.Contains(strings.ToLower(line), "high") || strings.Contains(strings.ToLower(line), "critical") {
				hazard["severity"] = "High"
			} else if strings.Contains(strings.ToLower(line), "low") || strings.Contains(strings.ToLower(line), "minor") {
				hazard["severity"] = "Low"
			}
			hazards = append(hazards, hazard)
		} else if currentSection == "mitigations" && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•") || regexp.MustCompile(`^\d+\.`).MatchString(line)) {
			mitigations = append(mitigations, line)
		} else if currentSection == "standards" && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•") || regexp.MustCompile(`^\d+\.`).MatchString(line)) {
			standards = append(standards, line)
		}
	}

	return hazardLevel, hazards, mitigations, standards
}

func ParseCorrosionResponse(text string) (string, float64, []string, []string, string) {
	riskLevel := "MEDIUM"
	corrosionRate := 0.5
	var mechanisms []string
	var recommendations []string
	estimatedLife := "10-15 years"

	lines := strings.Split(text, "\n")
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		upperLine := strings.ToUpper(line)

		if strings.Contains(upperLine, "RISK LEVEL:") || strings.Contains(upperLine, "CORROSION RISK:") {
			currentSection = ""
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				level := strings.ToUpper(strings.TrimSpace(parts[1]))
				if strings.Contains(level, "HIGH") {
					riskLevel = "HIGH"
				} else if strings.Contains(level, "LOW") {
					riskLevel = "LOW"
				} else {
					riskLevel = "MEDIUM"
				}
			}
			continue
		} else if strings.Contains(upperLine, "CORROSION RATE:") {
			currentSection = ""
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				rateStr := strings.TrimSpace(parts[1])
				re := regexp.MustCompile(`[\d.]+`)
				if match := re.FindString(rateStr); match != "" {
					if val, err := strconv.ParseFloat(match, 64); err == nil {
						corrosionRate = val
					}
				}
			}
			continue
		} else if strings.Contains(upperLine, "ESTIMATED LIFE:") || strings.Contains(upperLine, "EQUIPMENT LIFE:") {
			currentSection = ""
			parts := strings.SplitN(line, ":", 2)
			if len(parts) > 1 {
				estimatedLife = strings.TrimSpace(parts[1])
			}
			continue
		} else if strings.Contains(upperLine, "MECHANISMS:") || strings.Contains(upperLine, "CORROSION MECHANISMS:") {
			currentSection = "mechanisms"
			continue
		} else if strings.Contains(upperLine, "RECOMMENDATIONS:") {
			currentSection = "recommendations"
			continue
		}

		if currentSection == "mechanisms" && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•") || regexp.MustCompile(`^\d+\.`).MatchString(line)) {
			mechanisms = append(mechanisms, line)
		} else if currentSection == "recommendations" && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•") || regexp.MustCompile(`^\d+\.`).MatchString(line)) {
			recommendations = append(recommendations, line)
		}
	}

	return riskLevel, corrosionRate, mechanisms, recommendations, estimatedLife
}
