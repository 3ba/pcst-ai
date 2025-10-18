package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pcst-ai/backend/models"
	"pcst-ai/backend/services"
)

func HandleCorrosion(c *gin.Context) {
	var req models.CorrosionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All process parameters are required"})
		return
	}

	ctx := context.Background()
	gemini, err := services.NewGeminiService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer gemini.Close()

	prompt := fmt.Sprintf(`You are a Corrosion Engineering AI analyzing process equipment. Assess the corrosion risk based on the following parameters:

MATERIAL: %s
OPERATING TEMPERATURE: %.1f°C
pH LEVEL: %.1f
OPERATING PRESSURE: %.1f bar
FLUID VELOCITY: %.1f m/s

Provide your assessment in this exact format:

CORROSION RISK:
[HIGH/MEDIUM/LOW]

CORROSION RATE:
[Rate in mm/year]

CORROSION MECHANISMS:
- [Mechanism 1]
- [Mechanism 2]
- [Mechanism 3]

RECOMMENDATIONS:
- [Recommendation 1]
- [Recommendation 2]
- [Recommendation 3]

ESTIMATED LIFE:
[Equipment lifetime estimate]

Base your analysis on industry standards, material properties, and process conditions. Be specific and technical.`,
		req.Material, req.Temperature, req.PH, req.Pressure, req.Velocity)

	responseText, err := gemini.Generate(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	riskLevel, corrosionRate, mechanisms, recommendations, estimatedLife := services.ParseCorrosionResponse(responseText)

	c.JSON(http.StatusOK, models.CorrosionResponse{
		Success: true,
		Response: models.CorrosionDetails{
			RiskLevel:       riskLevel,
			CorrosionRate:   corrosionRate,
			Mechanisms:      mechanisms,
			Recommendations: recommendations,
			EstimatedLife:   estimatedLife,
		},
	})
}
