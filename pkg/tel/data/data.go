package data

type TraceInfo struct {
	TraceId string `json:"trace.id"`
	SpanId  string `json:"span.id"`
}

type HttpRequestInfo struct {
	ClientService  string `json:"client_service"`
	IPAddress      string `json:"ip_address"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	UserAgent      string `json:"user_agent"`
	Origin         string `json:"origin,omitempty"`
	XforwardedHost string `json:"xforwarded_host,omitempty"`
	RawQuery       string `json:"raw_query,omitempty"`

	SchoolId  string `json:"school_id,omitempty"`
	AccountId string `json:"account_id,omitempty"`
	UserId    string `json:"user_id,omitempty"`
}

type HttpResponseInfo struct {
	StatusCode int    `json:"status_code"`
	ErrorMsg   string `json:"error,omitempty"`
	Body       string `json:"body,omitempty"` // only if there is some error
}

type CustomDataType = map[string]any

type TLogLevel = string

const (
	LevelInfo  string = "Info"
	LevelError string = "Error"
	LevelWarn  string = "Warn"
	LevelDebug string = "Debug"
)

type Account struct {
	Id           int  `json:"account_id,omitempty"`
	SchoolId     int  `json:"school_id,omitempty"`
	UserId       int  `json:"user_id,omitempty"`
	Impersonated bool `json:"impersonated,omitempty"`
}
