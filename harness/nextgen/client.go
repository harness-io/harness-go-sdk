/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Harness NextGen Software Delivery Platform API Reference API v3.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	AccountId string
	ApiKey    string
	Endpoint  string

	// API Services

	APIKeysApi *APIKeysApiService

	AccessControlListApi *AccessControlListApiService

	AccountSettingApi *AccountSettingApiService

	AccountsApi *AccountsApiService

	AgentApi *AgentsApiService

	ApiKeyApi *ApiKeyApiService

	ApplicationsApiService *ApplicationsApiService

	AuditApi *AuditApiService

	AuditFiltersApi *AuditFiltersApiService

	AuthenticationSettingsApi *AuthenticationSettingsApiService

	ClustersApi *ClustersApiService

	CloudCostAnomaliesApi *CloudCostAnomaliesApiService

	CloudCostAutoStoppingFixedSchedulesApi *CloudCostAutoStoppingFixedSchedulesApiService

	CloudCostAutoStoppingLoadBalancersApi *CloudCostAutoStoppingLoadBalancersApiService

	CloudCostAutoStoppingRulesApi *CloudCostAutoStoppingRulesApiService

	CloudCostAutoStoppingRulesV2Api *CloudCostAutoStoppingRulesV2ApiService

	CloudCostBudgetsApi *CloudCostBudgetsApiService

	CloudCostDetailsApi *CloudCostDetailsApiService

	CloudCostPerspectiveReportsApi *CloudCostPerspectiveReportsApiService

	CloudCostPerspectivesApi *CloudCostPerspectivesApiService

	CloudCostRecommendationsApi *CloudCostRecommendationsApiService

	CloudCostRecommendationsDetailsApi *CloudCostRecommendationsDetailsApiService

	ConnectorsApi *ConnectorsApiService

	DelegateGroupTagsResourceApi *DelegateGroupTagsResourceApiService

	DelegateTokenResourceApi *DelegateTokenResourceApiService

	EnvironmentsApi *EnvironmentsApiService

	EnvironmentGroupApi *EnvironmentGroupApiService

	ExecuteApi *ExecuteApiService

	ExecutionDetailsApi *ExecutionDetailsApiService

	FeatureFlagsApi *FeatureFlagsApiService

	FileStoreApi *FileStoreApiService

	FilterApi *FilterApiService

	FreezeCRUDApi *FreezeCRUDApiService

	GitBranchesApi *GitBranchesApiService

	GitFullSyncApi *GitFullSyncApiService

	GitSyncApi *GitSyncApiService

	GitSyncErrorsApi *GitSyncErrorsApiService

	GitSyncSettingsApi *GitSyncSettingsApiService

	GnuPGPKeysApi *GnuPGPKeysApiService

	HarnessResourceGroupApi *HarnessResourceGroupApiService

	HarnessResourceTypeApi *HarnessResourceTypeApiService

	HostsApi *HostsApiService

	InfrastructuresApi *InfrastructuresApiService

	InputSetsApi *InputSetsApiService

	InviteApi *InviteApiService

	LicensesApi *LicensesApiService

	MonitoredServiceApi *MonitoredServiceApiService

	OrganizationApi *OrganizationApiService

	PermissionsApi *PermissionsApiService

	PipelinesApi *PipelinesApiService

	PipelinesDashboardApi *PipelinesDashboardApiService

	ProjectApi *ProjectApiService

	ProjectGitOpsApi *ProjectsApiService

	ProjectMappingsApi *ProjectMappingsApiService

	RepositoriesApiService *RepositoriesApiService

	RepositoryCertificatesApi *RepositoryCertificatesApiService

	RepositoryCredentialsApi *RepositoryCredentialsApiService

	RoleAssignmentsApi *RoleAssignmentsApiService

	RolesApi *RolesApiService

	RuleApi *RuleApiService

	RuleEnforcementApi *RuleEnforcementApiService

	SCIMApi *SCIMApiService

	SCMApi *SCMApiService

	SMTPApi *SMTPApiService

	SecretManagersApi *SecretManagersApiService

	SecretsApi *SecretsApiService

	ServiceAccountApi *ServiceAccountApiService

	ServicesApi *ServicesApiService

	ServiceOverridesApi *ServiceOverridesApiService
	OverridesApi        *OverridesApiService
	ProviderApi         *ProviderApiService

	SettingsApi *SettingsApiService

	SloApi *SloApiService

	SrmNotificationApiService *SrmNotificationApiService

	SourceCodeManagerApi *SourceCodeManagerApiService

	TargetGroupsApi *TargetGroupsApiService

	TargetsApi *TargetsApiService

	TokenApi *TokenApiService

	TriggersApi *TriggersApiService

	UsageApi *UsageApiService

	UserApi *UserApiService

	UserGroupApi *UserGroupApiService

	ValidateHostApi *ValidateHostApiService

	VariablesApi *VariablesApiService

	WebhookEventHandlerApi *WebhookEventHandlerApiService

	WebhookTriggersApi *WebhookTriggersApiService

	WorkspaceApi                  *WorkspacesApiService
	GitXWebhooksApiService        *GitXWebhooksApiService
	ProjectGitxWebhooksApiService *ProjectGitxWebhooksApiService
	OrgGitxWebhooksApiService     *OrgGitxWebhooksApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// Api Config
	c.ApiKey = cfg.ApiKey
	c.AccountId = cfg.AccountId
	c.Endpoint = cfg.BasePath

	// API Services
	c.APIKeysApi = (*APIKeysApiService)(&c.common)
	c.AccessControlListApi = (*AccessControlListApiService)(&c.common)
	c.AccountSettingApi = (*AccountSettingApiService)(&c.common)
	c.AccountsApi = (*AccountsApiService)(&c.common)
	c.AgentApi = (*AgentsApiService)(&c.common)
	c.ApiKeyApi = (*ApiKeyApiService)(&c.common)
	c.ApplicationsApiService = (*ApplicationsApiService)(&c.common)
	c.AuditApi = (*AuditApiService)(&c.common)
	c.AuditFiltersApi = (*AuditFiltersApiService)(&c.common)
	c.AuthenticationSettingsApi = (*AuthenticationSettingsApiService)(&c.common)
	c.ClustersApi = (*ClustersApiService)(&c.common)
	c.CloudCostAnomaliesApi = (*CloudCostAnomaliesApiService)(&c.common)
	c.CloudCostAutoStoppingFixedSchedulesApi = (*CloudCostAutoStoppingFixedSchedulesApiService)(&c.common)
	c.CloudCostAutoStoppingLoadBalancersApi = (*CloudCostAutoStoppingLoadBalancersApiService)(&c.common)
	c.CloudCostAutoStoppingRulesApi = (*CloudCostAutoStoppingRulesApiService)(&c.common)
	c.CloudCostAutoStoppingRulesV2Api = (*CloudCostAutoStoppingRulesV2ApiService)(&c.common)
	c.CloudCostBudgetsApi = (*CloudCostBudgetsApiService)(&c.common)
	c.CloudCostDetailsApi = (*CloudCostDetailsApiService)(&c.common)
	c.CloudCostPerspectiveReportsApi = (*CloudCostPerspectiveReportsApiService)(&c.common)
	c.CloudCostPerspectivesApi = (*CloudCostPerspectivesApiService)(&c.common)
	c.CloudCostRecommendationsApi = (*CloudCostRecommendationsApiService)(&c.common)
	c.CloudCostRecommendationsDetailsApi = (*CloudCostRecommendationsDetailsApiService)(&c.common)
	c.ConnectorsApi = (*ConnectorsApiService)(&c.common)
	c.DelegateGroupTagsResourceApi = (*DelegateGroupTagsResourceApiService)(&c.common)
	c.DelegateTokenResourceApi = (*DelegateTokenResourceApiService)(&c.common)
	c.EnvironmentsApi = (*EnvironmentsApiService)(&c.common)
	c.EnvironmentGroupApi = (*EnvironmentGroupApiService)(&c.common)
	c.ExecuteApi = (*ExecuteApiService)(&c.common)
	c.ExecutionDetailsApi = (*ExecutionDetailsApiService)(&c.common)
	c.FeatureFlagsApi = (*FeatureFlagsApiService)(&c.common)
	c.FileStoreApi = (*FileStoreApiService)(&c.common)
	c.FilterApi = (*FilterApiService)(&c.common)
	c.FreezeCRUDApi = (*FreezeCRUDApiService)(&c.common)
	c.GitBranchesApi = (*GitBranchesApiService)(&c.common)
	c.GitFullSyncApi = (*GitFullSyncApiService)(&c.common)
	c.GitSyncApi = (*GitSyncApiService)(&c.common)
	c.GitSyncErrorsApi = (*GitSyncErrorsApiService)(&c.common)
	c.GitSyncSettingsApi = (*GitSyncSettingsApiService)(&c.common)
	c.GnuPGPKeysApi = (*GnuPGPKeysApiService)(&c.common)
	c.HarnessResourceGroupApi = (*HarnessResourceGroupApiService)(&c.common)
	c.HarnessResourceTypeApi = (*HarnessResourceTypeApiService)(&c.common)
	c.HostsApi = (*HostsApiService)(&c.common)
	c.InfrastructuresApi = (*InfrastructuresApiService)(&c.common)
	c.InputSetsApi = (*InputSetsApiService)(&c.common)
	c.InviteApi = (*InviteApiService)(&c.common)
	c.LicensesApi = (*LicensesApiService)(&c.common)
	c.MonitoredServiceApi = (*MonitoredServiceApiService)(&c.common)
	c.OrganizationApi = (*OrganizationApiService)(&c.common)
	c.PermissionsApi = (*PermissionsApiService)(&c.common)
	c.PipelinesApi = (*PipelinesApiService)(&c.common)
	c.PipelinesDashboardApi = (*PipelinesDashboardApiService)(&c.common)
	c.ProjectApi = (*ProjectApiService)(&c.common)
	c.ProjectGitOpsApi = (*ProjectsApiService)(&c.common)
	c.ProjectMappingsApi = (*ProjectMappingsApiService)(&c.common)
	c.RepositoriesApiService = (*RepositoriesApiService)(&c.common)
	c.RepositoryCertificatesApi = (*RepositoryCertificatesApiService)(&c.common)
	c.RepositoryCredentialsApi = (*RepositoryCredentialsApiService)(&c.common)
	c.RoleAssignmentsApi = (*RoleAssignmentsApiService)(&c.common)
	c.RolesApi = (*RolesApiService)(&c.common)
	c.RuleApi = (*RuleApiService)(&c.common)
	c.RuleEnforcementApi = (*RuleEnforcementApiService)(&c.common)
	c.SCIMApi = (*SCIMApiService)(&c.common)
	c.SCMApi = (*SCMApiService)(&c.common)
	c.SMTPApi = (*SMTPApiService)(&c.common)
	c.SecretManagersApi = (*SecretManagersApiService)(&c.common)
	c.SecretsApi = (*SecretsApiService)(&c.common)
	c.ServiceAccountApi = (*ServiceAccountApiService)(&c.common)
	c.ServicesApi = (*ServicesApiService)(&c.common)
	c.ServiceOverridesApi = (*ServiceOverridesApiService)(&c.common)
	c.OverridesApi = (*OverridesApiService)(&c.common)
	c.ProviderApi = (*ProviderApiService)(&c.common)
	c.SettingsApi = (*SettingsApiService)(&c.common)
	c.SloApi = (*SloApiService)(&c.common)
	c.SrmNotificationApiService = (*SrmNotificationApiService)(&c.common)
	c.SourceCodeManagerApi = (*SourceCodeManagerApiService)(&c.common)
	c.TargetGroupsApi = (*TargetGroupsApiService)(&c.common)
	c.TargetsApi = (*TargetsApiService)(&c.common)
	c.TokenApi = (*TokenApiService)(&c.common)
	c.TriggersApi = (*TriggersApiService)(&c.common)
	c.UsageApi = (*UsageApiService)(&c.common)
	c.UserApi = (*UserApiService)(&c.common)
	c.UserGroupApi = (*UserGroupApiService)(&c.common)
	c.ValidateHostApi = (*ValidateHostApiService)(&c.common)
	c.VariablesApi = (*VariablesApiService)(&c.common)
	c.WebhookEventHandlerApi = (*WebhookEventHandlerApiService)(&c.common)
	c.WebhookTriggersApi = (*WebhookTriggersApiService)(&c.common)
	c.WorkspaceApi = (*WorkspacesApiService)(&c.common)
	c.GitXWebhooksApiService = (*GitXWebhooksApiService)(&c.common)
	c.ProjectGitxWebhooksApiService = (*ProjectGitxWebhooksApiService)(&c.common)
	c.OrgGitxWebhooksApiService = (*OrgGitxWebhooksApiService)(&c.common)
	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *retryablehttp.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *retryablehttp.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}
		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = strings.ReplaceAll(query.Encode(), "+", "%20")

	// Generate a new request
	if body != nil {
		localVarRequest, err = retryablehttp.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = retryablehttp.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest.Request)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	if e.model == nil {
		return e.error
	}

	failure, ok := e.model.(Failure)
	if ok {
		if failure.Message != "" {
			return failure.Message
		}
	} else {
		failure, ok := e.model.(AuthzFailure)
		if ok {
			if failure.Message != "" {
				return failure.Message
			}
		} else {
			failure, ok := e.model.(ModelError)
			if ok {
				if failure.Message != "" {
					return failure.Message
				}
			}
		}
	}

	// if len(failure.ResponseMessages) > 0 {
	// 	return failure.ResponseMessages[0].Message
	// }

	if len(failure.Errors) > 0 {
		return fmt.Sprintf("%s %s", failure.Errors[0].FieldId, failure.Errors[0].Error_)
	}

	return failure.Code
}

func (e GenericSwaggerError) Code() ErrorCode {
	return ErrorCode(e.model.(Failure).Code)
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
