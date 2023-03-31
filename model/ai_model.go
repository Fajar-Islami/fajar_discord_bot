package model

type AIReq struct {
	EnableGoogleResults bool   `json:"enable_google_results"`
	EnableMemory        bool   `json:"enable_memory"`
	InputText           string `json:"input_text"`
}

type AIRes struct {
	Message   string        `json:"message,omitempty"`
	Detail    string        `json:"detail,omitempty"`
	ImageUrls []interface{} `json:"image_urls,omitempty"`
}
