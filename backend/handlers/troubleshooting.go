package handlers

import (
	"context"
	"fmt"
	"net/http"

	"pcst-ai/backend/models"
	"pcst-ai/backend/services"

	"github.com/gin-gonic/gin"
)

func HandleSearch(c *gin.Context) {
	var req models.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Equipment and problem description are required"})
		return
	}

	ctx := context.Background()
	gemini, err := services.NewGeminiService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer gemini.Close()

	errorCodeText := "None provided"
	if req.ErrorCode != "" {
		errorCodeText = req.ErrorCode
	}

	prompt := fmt.Sprintf(`You are an expert Process Control System Technician at Aramco.

Analyze the following troubleshooting request and provide a clear, structured response:

EQUIPMENT: %s
PROBLEM: %s
ERROR CODE: %s

IMPORTANT SAFETY GUIDELINES:
- Always prioritize safety over production
- Follow Aramco safety protocols
- Verify equipment isolation before maintenance
- Use proper PPE and safety equipment
- Never bypass safety systems
- Follow lockout/tagout procedures

Please provide your response in this exact format:

ANALYSIS:
[Your analysis of the problem]

POSSIBLE CAUSES:
1. [Cause 1]
2. [Cause 2]
3. [Cause 3]

TROUBLESHOOTING STEPS:
1. [Step 1 - include safety precautions]
2. [Step 2 - include safety precautions]
3. [Step 3 - include safety precautions]
4. [Continue as needed]

SAFETY WARNINGS:
- [Important safety warning 1]
- [Important safety warning 2]
- [Important safety warning 3]

EQUIPMENT NOTES:
[Specific considerations for this equipment type]

Provide your response in a clear, structured format that a technician can follow safely.`, req.Equipment, req.Problem, errorCodeText)

	responseText, err := gemini.Generate(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	analysis, causes, steps, safetyWarnings, equipmentNotes := services.ParseTroubleshootingResponse(responseText)

	c.JSON(http.StatusOK, models.SearchResponse{
		Success:   true,
		Equipment: req.Equipment,
		Response: models.ResponseSections{
			Analysis:       analysis,
			Causes:         causes,
			Steps:          steps,
			SafetyWarnings: safetyWarnings,
			EquipmentNotes: equipmentNotes,
		},
	})
}

func HandleGetEquipment(c *gin.Context) {
	equipmentList := []string{
		"Control Valve",
		"Pressure Gauge",
		"Displacement Level Transmitter",
		"Flow Transmitter",
		"Pressure Transmitter",
		"Temperature Transmitter",
		"Level Transmitter",
		"Resistance Temperature Detector (RTD)",
		"Thermocouple",
		"Proximity Transducer",
		"Control Panel",
		"SCADA System",
		"PLC",
		"DCS",
		"Safety Instrumented System (SIS)",
		"Emergency Shutdown System (ESD)",
		"Fire and Gas Detection System",
		"Compressor",
		"Pump",
		"Turbine",
		"Heat Exchanger",
		"Separator",
		"Storage Tank",
		"Pipeline",
		"Valve Actuator",
	}

	c.JSON(http.StatusOK, gin.H{"equipment": equipmentList})
}
