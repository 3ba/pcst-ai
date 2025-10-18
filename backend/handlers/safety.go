package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pcst-ai/backend/models"
	"pcst-ai/backend/services"
)

func HandleSafety(c *gin.Context) {
	var req models.SafetyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task description is required"})
		return
	}

	ctx := context.Background()
	gemini, err := services.NewGeminiService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer gemini.Close()

	prompt := fmt.Sprintf(`You are an AI Safety Advisor for oil and gas operations. Analyze the following job task and provide a comprehensive safety assessment.

JOB TASK:
%s

Provide your assessment in this exact format:

HAZARD LEVEL:
[HIGH/MEDIUM/LOW]

IDENTIFIED HAZARDS:
- [Hazard 1]
- [Hazard 2]
- [Hazard 3]

RECOMMENDED MITIGATIONS:
- [Mitigation 1]
- [Mitigation 2]
- [Mitigation 3]

RELEVANT STANDARDS:
- [Standard 1]
- [Standard 2]
- [Standard 3]

Be thorough and specific. Include industry best practices and Aramco safety standards.`, req.Task)

	responseText, err := gemini.Generate(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hazardLevel, hazards, mitigations, standards := services.ParseSafetyResponse(responseText)

	hazardDetails := make([]models.HazardDetail, 0, len(hazards))
	for _, h := range hazards {
		hazardDetails = append(hazardDetails, models.HazardDetail{
			Name:        h["name"],
			Severity:    h["severity"],
			Probability: h["probability"],
		})
	}

	c.JSON(http.StatusOK, models.SafetyResponse{
		Success: true,
		Response: models.SafetyDetails{
			HazardLevel: hazardLevel,
			Hazards:     hazardDetails,
			Mitigations: mitigations,
			Standards:   standards,
		},
	})
}
