package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pcst-ai/backend/models"
	"pcst-ai/backend/services"
)

func HandleVCRA(c *gin.Context) {
	var req models.VCRARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Control room logs are required"})
		return
	}

	ctx := context.Background()
	gemini, err := services.NewGeminiService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer gemini.Close()

	prompt := fmt.Sprintf(`You are a Virtual Control Room Advisor for an oil and gas facility. Analyze the following control room logs and provide a detailed incident analysis.

CONTROL ROOM LOGS:
%s

Provide your analysis in this exact format:

ROOT CAUSE:
[Identify the root cause of the incident]

RISK LEVEL:
[HIGH/MEDIUM/LOW]

CONFIDENCE:
[Confidence percentage, e.g., 85%%]

IMMEDIATE ACTIONS:
- [Action 1]
- [Action 2]
- [Action 3]

RECOVERY TIMELINE:
- [Timeline step 1]
- [Timeline step 2]
- [Timeline step 3]

Be specific and actionable. Focus on immediate response and safety.`, req.Logs)

	responseText, err := gemini.Generate(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rootCause, riskLevel, confidence, actions, timeline := services.ParseVCRAResponse(responseText)

	c.JSON(http.StatusOK, models.VCRAResponse{
		Success: true,
		Response: models.VCRADetails{
			RootCause:  rootCause,
			RiskLevel:  riskLevel,
			Confidence: confidence,
			Actions:    actions,
			Timeline:   timeline,
		},
	})
}
