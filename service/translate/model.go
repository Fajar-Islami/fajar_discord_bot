package translate

type (
	DetectLang struct {
		Q string `form:"q"`
	}

	RespDetectLang struct {
		Data struct {
			Detections [][]struct {
				Confidence float64 `json:"confidence"`
				Language   string  `json:"language"`
				IsReliable bool    `json:"isReliable"`
			} `json:"detections"`
		} `json:"data"`
	}
)
