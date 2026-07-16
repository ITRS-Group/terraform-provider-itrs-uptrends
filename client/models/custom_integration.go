package client

// CustomIntegration represents a custom integration in the Uptrends API.
// The same shape is used for the request body (POST/PUT) and the response body
// (GET). Fields that are server-generated use omitempty so they are not sent on
// create.
type CustomIntegration struct {
	IntegrationGuid string `json:"IntegrationGuid,omitempty"`
	Name            string `json:"Name"`
	IsActive        bool   `json:"IsActive"`
	// Notes, HttpStepDefinitions and IntegrationVariables are required by the API
	// on both POST and PUT and must always be present in the request body
	// (the lists serialized as non-null arrays, []  when empty). PUT performs no
	// server-side defaulting, so the full object is always sent.
	Notes                string                `json:"Notes"`
	HttpStepDefinitions  []HttpStepDefinition  `json:"HttpStepDefinitions"`
	IntegrationVariables []IntegrationVariable `json:"IntegrationVariables"`
}

// HttpStepDefinition groups the HTTP steps, user defined functions and usages
// of a custom integration.
type HttpStepDefinition struct {
	HttpStepDefinitionUsageGuid string                `json:"HttpStepDefinitionUsageGuid,omitempty"`
	Steps                       []HttpStep            `json:"Steps,omitempty"`
	UserDefinedFunctions        []UserDefinedFunction `json:"UserDefinedFunctions,omitempty"`
	Usages                      []string              `json:"Usages,omitempty"`
}

// HttpStep represents a single HTTP request step of a custom integration.
type HttpStep struct {
	Url                        string                  `json:"Url,omitempty"`
	Method                     string                  `json:"Method,omitempty"`
	Body                       string                  `json:"Body,omitempty"`
	BodyType                   string                  `json:"BodyType,omitempty"`
	VaultFileId                string                  `json:"VaultFileId,omitempty"`
	MultiPartForm              []MultiPartFormItem     `json:"MultiPartForm,omitempty"`
	RequestHeaders             []IntegrationKeyValue   `json:"RequestHeaders,omitempty"`
	Variables                  []HttpStepVariable      `json:"Variables,omitempty"`
	Assertions                 []HttpStepAssertion     `json:"Assertions,omitempty"`
	Name                       string                  `json:"Name,omitempty"`
	UseFixedClientCertificate  bool                    `json:"UseFixedClientCertificate"`
	Authentication             *HttpStepAuthentication `json:"Authentication,omitempty"`
	IgnoreCertificateErrors    bool                    `json:"IgnoreCertificateErrors"`
	ClientCertificateVaultItem string                  `json:"ClientCertificateVaultItem,omitempty"`
	Delay                      int                     `json:"Delay"`
	StepType                   string                  `json:"StepType,omitempty"`
	RetryUntilSuccessful       bool                    `json:"RetryUntilSuccessful"`
	MaxAttempts                int                     `json:"MaxAttempts"`
	RetryWaitMilliseconds      int                     `json:"RetryWaitMilliseconds"`
	PreRequestScript           string                  `json:"PreRequestScript,omitempty"`
	PostResponseScript         string                  `json:"PostResponseScript,omitempty"`
	CalculatedContentType      string                  `json:"CalculatedContentType,omitempty"`
	AllowedTlsVersions         []string                `json:"AllowedTlsVersions,omitempty"`
	MsaStepHttpVersion         string                  `json:"MsaStepHttpVersion,omitempty"`
}

// MultiPartFormItem represents an entry of a multipart form body.
type MultiPartFormItem struct {
	Type  string `json:"Type,omitempty"`
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

// IntegrationKeyValue is a generic key/value pair (used for request headers).
type IntegrationKeyValue struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

// HttpStepVariable captures a value from a response into a named variable.
type HttpStepVariable struct {
	Source    string   `json:"Source,omitempty"`
	Property  string   `json:"Property,omitempty"`
	Name      string   `json:"Name,omitempty"`
	Arguments []string `json:"Arguments,omitempty"`
}

// HttpStepAssertion validates a property of a response.
type HttpStepAssertion struct {
	Source      string `json:"Source,omitempty"`
	Property    string `json:"Property,omitempty"`
	Comparison  string `json:"Comparison,omitempty"`
	TargetValue string `json:"TargetValue,omitempty"`
}

// HttpStepAuthentication holds the authentication settings of a step.
type HttpStepAuthentication struct {
	Id                 string `json:"Id,omitempty"`
	AuthenticationType string `json:"AuthenticationType,omitempty"`
	UserName           string `json:"UserName,omitempty"`
	Password           string `json:"Password,omitempty"`
	PasswordSpecified  bool   `json:"PasswordSpecified"`
}

// UserDefinedFunction represents a reusable transformation function.
type UserDefinedFunction struct {
	Name          string                `json:"Name,omitempty"`
	Type          string                `json:"Type,omitempty"`
	Mappings      []IntegrationKeyValue `json:"Mappings,omitempty"`
	Regex         string                `json:"Regex,omitempty"`
	JwtSigningKey string                `json:"JwtSigningKey,omitempty"`
	JwtAlgorithm  string                `json:"JwtAlgorithm,omitempty"`
	HashKey       string                `json:"HashKey,omitempty"`
	HashAlgorithm string                `json:"HashAlgorithm,omitempty"`
}

// IntegrationVariable represents a variable that can be referenced by the
// integration steps.
type IntegrationVariable struct {
	Name                        string `json:"Name,omitempty"`
	Value                       string `json:"Value,omitempty"`
	IsValueSetInEscalationLevel bool   `json:"IsValueSetInEscalationLevel"`
}
