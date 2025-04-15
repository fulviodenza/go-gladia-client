package gladia

import "time"

// TranscriptionRequest represents a request to the Gladia transcription API
type TranscriptionRequest struct {
	AudioURL                 string                          `json:"audio_url"`
	Diarization              bool                            `json:"diarization,omitempty"`
	DiarizationConfig        *DiarizationConfig              `json:"diarization_config,omitempty"`
	Translation              bool                            `json:"translation,omitempty"`
	TranslationConfig        *TranslationConfig              `json:"translation_config,omitempty"`
	Subtitles                bool                            `json:"subtitles,omitempty"`
	SubtitlesConfig          *SubtitlesConfig                `json:"subtitles_config,omitempty"`
	DetectLanguage           bool                            `json:"detect_language,omitempty"`
	EnableCodeSwitching      bool                            `json:"enable_code_switching,omitempty"`
	CodeSwitchingConfig      *CodeSwitchingConfig            `json:"code_switching_config,omitempty"`
	Callback                 bool                            `json:"callback,omitempty"`
	CallbackURL              string                          `json:"callback_url,omitempty"`
	CallbackConfig           *CallbackConfig                 `json:"callback_config,omitempty"`
	Language                 string                          `json:"language,omitempty"`
	ContextPrompt            string                          `json:"context_prompt,omitempty"`
	CustomVocabulary         bool                            `json:"custom_vocabulary,omitempty"`
	CustomVocabularyConfig   *CustomVocabularyConfig         `json:"custom_vocabulary_config,omitempty"`
	Summarization            bool                            `json:"summarization,omitempty"`
	SummarizationConfig      *SummarizationConfig            `json:"summarization_config,omitempty"`
	Moderation               bool                            `json:"moderation,omitempty"`
	NamedEntityRecognition   bool                            `json:"named_entity_recognition,omitempty"`
	Chapterization           bool                            `json:"chapterization,omitempty"`
	NameConsistency          bool                            `json:"name_consistency,omitempty"`
	CustomSpelling           bool                            `json:"custom_spelling,omitempty"`
	CustomSpellingConfig     *CustomSpellingConfig           `json:"custom_spelling_config,omitempty"`
	StructuredDataExtraction bool                            `json:"structured_data_extraction,omitempty"`
	StructuredDataExtrConfig *StructuredDataExtractionConfig `json:"structured_data_extraction_config,omitempty"`
	SentimentAnalysis        bool                            `json:"sentiment_analysis,omitempty"`
	AudioToLLM               bool                            `json:"audio_to_llm,omitempty"`
	AudioToLLMConfig         *AudioToLLMConfig               `json:"audio_to_llm_config,omitempty"`
	Sentences                bool                            `json:"sentences,omitempty"`
	DisplayMode              bool                            `json:"display_mode,omitempty"`
	PunctuationEnhanced      bool                            `json:"punctuation_enhanced,omitempty"`
}

// DiarizationConfig contains settings for speaker diarization
type DiarizationConfig struct {
	NumberOfSpeakers int  `json:"number_of_speakers,omitempty"`
	MinSpeakers      int  `json:"min_speakers,omitempty"`
	MaxSpeakers      int  `json:"max_speakers,omitempty"`
	Enhanced         bool `json:"enhanced,omitempty"`
}

// TranslationConfig contains settings for translation
type TranslationConfig struct {
	Model                   string   `json:"model,omitempty"`
	TargetLanguages         []string `json:"target_languages,omitempty"`
	MatchOriginalUtterances bool     `json:"match_original_utterances,omitempty"`
}

// SubtitlesConfig contains settings for subtitle generation
type SubtitlesConfig struct {
	Formats                 []string `json:"formats,omitempty"`
	MinimumDuration         float64  `json:"minimum_duration,omitempty"`
	MaximumDuration         float64  `json:"maximum_duration,omitempty"`
	MaximumCharactersPerRow int      `json:"maximum_characters_per_row,omitempty"`
	MaximumRowsPerCaption   int      `json:"maximum_rows_per_caption,omitempty"`
	Style                   string   `json:"style,omitempty"`
}

// CodeSwitchingConfig contains settings for code switching
type CodeSwitchingConfig struct {
	Languages []string `json:"languages,omitempty"`
}

// CallbackConfig contains settings for callback notifications
type CallbackConfig struct {
	URL    string `json:"url"`
	Method string `json:"method,omitempty"` // Default is POST
}

// VocabularyEntry represents a vocabulary entry with optional pronunciation
type VocabularyEntry struct {
	Value          string   `json:"value,omitempty"`
	Pronunciations []string `json:"pronunciations,omitempty"`
	Intensity      float64  `json:"intensity,omitempty"`
	Language       string   `json:"language,omitempty"`
}

// RealtimeProcessing contains real-time processing configurations
type RealtimeProcessing struct {
	CustomVocabulary       bool                   `json:"custom_vocabulary,omitempty"`
	CustomVocabularyConfig *InnerVocabularyConfig `json:"custom_vocabulary_config,omitempty"`
}

// InnerVocabularyConfig contains vocabulary configuration settings
type InnerVocabularyConfig struct {
	Vocabulary       []any   `json:"vocabulary,omitempty"` // Can be string or VocabularyEntry
	DefaultIntensity float64 `json:"default_intensity,omitempty"`
}

// VocabularyConfig contains vocabulary settings
type VocabularyConfig struct {
	RealtimeProcessing RealtimeProcessing `json:"realtime_processing,omitempty"`
}

// CustomVocabularyConfig for vocabulary enhancement
type CustomVocabularyConfig struct {
	Vocabulary       VocabularyConfig `json:"vocabulary,omitempty"`
	DefaultIntensity float64          `json:"default_intensity,omitempty"`
}

// SummarizationConfig contains settings for summarization
type SummarizationConfig struct {
	Type string `json:"type,omitempty"`
}

// CustomSpellingConfig contains custom spelling configuration
type CustomSpellingConfig struct {
	SpellingDictionary map[string][]string `json:"spelling_dictionary,omitempty"`
}

// StructuredDataExtractionConfig contains structured data extraction settings
type StructuredDataExtractionConfig struct {
	Classes []string `json:"classes,omitempty"`
}

// AudioToLLMConfig contains audio to LLM configuration
type AudioToLLMConfig struct {
	Prompts []string `json:"prompts,omitempty"`
}

// TranscriptionResponse represents the initial response from the Gladia transcription API
type TranscriptionResponse struct {
	ID        string `json:"id"`
	ResultURL string `json:"result_url"`
}

// GladiaFile represents a file in the Gladia API
type GladiaFile struct {
	ID               string  `json:"id"`
	Filename         string  `json:"filename"`
	Source           string  `json:"source"`
	AudioDuration    float64 `json:"audio_duration"`
	NumberOfChannels int     `json:"number_of_channels"`
}

// ErrorInfo contains error details
type ErrorInfo struct {
	StatusCode int    `json:"status_code"`
	Exception  string `json:"exception,omitempty"`
	Message    string `json:"message,omitempty"`
}

// Word represents a single transcribed word
type Word struct {
	Word       string  `json:"word"`
	Start      float64 `json:"start"`
	End        float64 `json:"end"`
	Confidence float64 `json:"confidence"`
}

// Utterance represents a complete utterance by a speaker
type Utterance struct {
	Language   string  `json:"language,omitempty"`
	Start      float64 `json:"start"`
	End        float64 `json:"end"`
	Confidence float64 `json:"confidence"`
	Channel    int     `json:"channel,omitempty"`
	Speaker    int     `json:"speaker,omitempty"`
	Words      []Word  `json:"words,omitempty"`
	Text       string  `json:"text"`
}

// Subtitle represents subtitle information
type Subtitle struct {
	Format    string `json:"format"`
	Subtitles string `json:"subtitles"`
}

// SentenceResult represents sentence processing results
type SentenceResult struct {
	Success  bool       `json:"success"`
	IsEmpty  bool       `json:"is_empty"`
	ExecTime int        `json:"exec_time"`
	Error    *ErrorInfo `json:"error,omitempty"`
	Results  []string   `json:"results,omitempty"`
}

// TranscriptionMetadata contains metadata about transcription
type TranscriptionMetadata struct {
	AudioDuration            float64 `json:"audio_duration"`
	NumberOfDistinctChannels int     `json:"number_of_distinct_channels"`
	BillingTime              float64 `json:"billing_time"`
	TranscriptionTime        float64 `json:"transcription_time"`
}

// TranscriptionData contains the actual transcription data
type TranscriptionData struct {
	FullTranscript string           `json:"full_transcript,omitempty"`
	Languages      []string         `json:"languages,omitempty"`
	Sentences      []SentenceResult `json:"sentences,omitempty"`
	Subtitles      []Subtitle       `json:"subtitles,omitempty"`
	Utterances     []Utterance      `json:"utterances,omitempty"`
}

// TranslationResult contains translation results for a specific language
type TranslationResult struct {
	Error          *ErrorInfo       `json:"error,omitempty"`
	FullTranscript string           `json:"full_transcript,omitempty"`
	Languages      []string         `json:"languages,omitempty"`
	Sentences      []SentenceResult `json:"sentences,omitempty"`
	Subtitles      []Subtitle       `json:"subtitles,omitempty"`
	Utterances     []Utterance      `json:"utterances,omitempty"`
}

// ProcessingResult represents a generic processing result
type ProcessingResult struct {
	Success  bool       `json:"success"`
	IsEmpty  bool       `json:"is_empty"`
	ExecTime int        `json:"exec_time"`
	Error    *ErrorInfo `json:"error,omitempty"`
	Results  any        `json:"results,omitempty"`
}

// AudioToLLMPromptResult represents results for a specific LLM prompt
type AudioToLLMPromptResult struct {
	Success  bool       `json:"success"`
	IsEmpty  bool       `json:"is_empty"`
	ExecTime int        `json:"exec_time"`
	Error    *ErrorInfo `json:"error,omitempty"`
	Results  struct {
		Prompt   string `json:"prompt"`
		Response string `json:"response"`
	} `json:"results,omitempty"`
}

// AudioToLLMResults contains audio to LLM processing results
type AudioToLLMResults struct {
	Success  bool                     `json:"success"`
	IsEmpty  bool                     `json:"is_empty"`
	ExecTime int                      `json:"exec_time"`
	Error    *ErrorInfo               `json:"error,omitempty"`
	Results  []AudioToLLMPromptResult `json:"results,omitempty"`
}

// TranscriptionResultData contains all the processing results
type TranscriptionResultData struct {
	Metadata                 TranscriptionMetadata `json:"metadata"`
	Transcription            TranscriptionData     `json:"transcription"`
	Translation              ProcessingResult      `json:"translation,omitempty"`
	Summarization            ProcessingResult      `json:"summarization,omitempty"`
	Moderation               ProcessingResult      `json:"moderation,omitempty"`
	NamedEntityRecognition   ProcessingResult      `json:"named_entity_recognition,omitempty"`
	NameConsistency          ProcessingResult      `json:"name_consistency,omitempty"`
	CustomSpelling           ProcessingResult      `json:"custom_spelling,omitempty"`
	SpeakerReidentification  ProcessingResult      `json:"speaker_reidentification,omitempty"`
	StructuredDataExtraction ProcessingResult      `json:"structured_data_extraction,omitempty"`
	SentimentAnalysis        ProcessingResult      `json:"sentiment_analysis,omitempty"`
	AudioToLLM               AudioToLLMResults     `json:"audio_to_llm,omitempty"`
	Sentences                ProcessingResult      `json:"sentences,omitempty"`
	DisplayMode              ProcessingResult      `json:"display_mode,omitempty"`
	Chapters                 ProcessingResult      `json:"chapters,omitempty"`
}

// CompletedTranscriptionResult represents the complete transcription result data
type CompletedTranscriptionResult struct {
	ID             string                  `json:"id"`
	RequestID      string                  `json:"request_id"`
	Version        int                     `json:"version"`
	Status         string                  `json:"status"`
	CreatedAt      time.Time               `json:"created_at"`
	CompletedAt    time.Time               `json:"completed_at,omitempty"`
	CustomMetadata map[string]any          `json:"custom_metadata,omitempty"`
	ErrorCode      int                     `json:"error_code,omitempty"`
	Kind           string                  `json:"kind"`
	File           GladiaFile              `json:"file"`
	RequestParams  TranscriptionRequest    `json:"request_params"`
	Result         TranscriptionResultData `json:"result"`
}

// GetTranscriptionStatus contains status information for a transcription
type GetTranscriptionStatus struct {
	ID             string                   `json:"id"`
	RequestID      string                   `json:"request_id"`
	Version        int                      `json:"version"`
	Status         string                   `json:"status"`
	CreatedAt      time.Time                `json:"created_at"`
	CompletedAt    *time.Time               `json:"completed_at,omitempty"`
	CustomMetadata map[string]any           `json:"custom_metadata,omitempty"`
	ErrorCode      *int                     `json:"error_code,omitempty"`
	Kind           string                   `json:"kind"`
	File           *GladiaFile              `json:"file,omitempty"`
	RequestParams  *TranscriptionRequest    `json:"request_params,omitempty"`
	Result         *TranscriptionResultData `json:"result,omitempty"`
}
