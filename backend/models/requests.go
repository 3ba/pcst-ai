package models

type SearchRequest struct {
	Equipment string `json:"equipment" binding:"required"`
	Problem   string `json:"problem" binding:"required"`
	ErrorCode string `json:"error_code"`
}

type VCRARequest struct {
	Logs string `json:"logs" binding:"required"`
}

type SafetyRequest struct {
	Task string `json:"task" binding:"required"`
}

type CorrosionRequest struct {
	Material    string  `json:"material" binding:"required"`
	Temperature float64 `json:"temperature" binding:"required"`
	PH          float64 `json:"ph" binding:"required"`
	Pressure    float64 `json:"pressure" binding:"required"`
	Velocity    float64 `json:"velocity" binding:"required"`
}
