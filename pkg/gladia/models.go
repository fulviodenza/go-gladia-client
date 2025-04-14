package gladia

// TranscriptionRequest represents a request to the Gladia transcription API
type TranscriptionRequest struct {
	AudioURL            string             `json:"audio_url"`
	Diarization         bool               `json:"diarization,omitempty"`
	DiarizationConfig   *DiarizationConfig `json:"diarization_config,omitempty"`
	Translation         bool               `json:"translation,omitempty"`
	TranslationConfig   *TranslationConfig `json:"translation_config,omitempty"`
	Subtitles           bool               `json:"subtitles,omitempty"`
	SubtitlesConfig     *SubtitlesConfig   `json:"subtitles_config,omitempty"`
	DetectLanguage      bool               `json:"detect_language,omitempty"`
	EnableCodeSwitching bool               `json:"enable_code_switching,omitempty"`
	Callback            bool               `json:"callback,omitempty"`
	CallbackConfig      *CallbackConfig    `json:"callback_config,omitempty"`
}

// DiarizationConfig contains settings for speaker diarization
type DiarizationConfig struct {
	NumberOfSpeakers int `json:"number_of_speakers,omitempty"`
	MinSpeakers      int `json:"min_speakers,omitempty"`
	MaxSpeakers      int `json:"max_speakers,omitempty"`
}

// TranslationConfig contains settings for translation
type TranslationConfig struct {
	Model           string   `json:"model,omitempty"`
	TargetLanguages []string `json:"target_languages,omitempty"`
}

// SubtitlesConfig contains settings for subtitle generation
type SubtitlesConfig struct {
	Formats []string `json:"formats,omitempty"`
}

// CallbackConfig contains settings for callback notifications
type CallbackConfig struct {
	URL    string `json:"url"`
	Method string `json:"method,omitempty"` // Default is POST
}
