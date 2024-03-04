/*
  - Harness NextGen Software Delivery Platform API Reference
    *
  - This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of

the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject:
<security-definitions> -->

	*
	* API version: 3.0
	* Contact: contact@harness.io
	* Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
*/
package nextgen

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ServiceOverridesApiService service

/*
ServiceOverridesApiService Create an ServiceOverride Entity
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
Passed from http.Request or context.Background().
 * @param accountIdentifier
 * @param optional nil or *ServiceOverridesApiCreateServiceOverrideV2Opts - Optional Parameters:
     * @param "Body" (optional.Interface of ServiceOverrideRequestDtov2) -
     * @param "Branch" (optional.String) - 
     * @param "FilePath" (optional.String) - 
     * @param "CommitMsg" (optional.String) - 
     * @param "IsNewBranch" (optional.Bool) - 
     * @param "BaseBranch" (optional.String) - 
     * @param "ConnectorRef" (optional.String) - 
     * @param "StoreType" (optional.String) - 
     * @param "RepoName" (optional.String) - 
     * @param "IsHarnessCodeRepo" (optional.Bool) - 
@return ResponseServiceOverridesResponseDtov2
*/

type ServiceOverridesApiCreateServiceOverrideV2Opts struct {
	Body optional.Interface
	Branch optional.String
    FilePath optional.String
    CommitMsg optional.String
    IsNewBranch optional.Bool
    BaseBranch optional.String
    ConnectorRef optional.String
    StoreType optional.String
    RepoName optional.String
    IsHarnessCodeRepo optional.Bool
}

func (a *ServiceOverridesApiService) CreateServiceOverrideV2(ctx context.Context, accountIdentifier string, localVarOptionals *ServiceOverridesApiCreateServiceOverrideV2Opts) (ResponseServiceOverridesResponseDtov2, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseServiceOverridesResponseDtov2
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/serviceOverrides"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilePath.IsSet() {
		localVarQueryParams.Add("filePath", parameterToString(localVarOptionals.FilePath.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CommitMsg.IsSet() {
		localVarQueryParams.Add("commitMsg", parameterToString(localVarOptionals.CommitMsg.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsNewBranch.IsSet() {
		localVarQueryParams.Add("isNewBranch", parameterToString(localVarOptionals.IsNewBranch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseBranch.IsSet() {
		localVarQueryParams.Add("baseBranch", parameterToString(localVarOptionals.BaseBranch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConnectorRef.IsSet() {
		localVarQueryParams.Add("connectorRef", parameterToString(localVarOptionals.ConnectorRef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StoreType.IsSet() {
		localVarQueryParams.Add("storeType", parameterToString(localVarOptionals.StoreType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoName.IsSet() {
		localVarQueryParams.Add("repoName", parameterToString(localVarOptionals.RepoName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsHarnessCodeRepo.IsSet() {
		localVarQueryParams.Add("isHarnessCodeRepo", parameterToString(localVarOptionals.IsHarnessCodeRepo.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody,
			localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResponseServiceOverridesResponseDtov2
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ServiceOverridesApiService Delete a Service Override entity
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
Passed from http.Request or context.Background().
 * @param identifier
 * @param accountIdentifier
 * @param optional nil or *ServiceOverridesApiDeleteServiceOverrideV2Opts - Optional Parameters:
     * @param "OrgIdentifier" (optional.String) -
     * @param "ProjectIdentifier" (optional.String) -
@return ResponseBoolean
*/

type ServiceOverridesApiDeleteServiceOverrideV2Opts struct {
	OrgIdentifier     optional.String
	ProjectIdentifier optional.String
}

func (a *ServiceOverridesApiService) DeleteServiceOverrideV2(ctx context.Context, identifier string, accountIdentifier string, localVarOptionals *ServiceOverridesApiDeleteServiceOverrideV2Opts) (ResponseDtoBoolean, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoBoolean
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/serviceOverrides/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", fmt.Sprintf("%v",
		identifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier",
			parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier",
			parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody,
			localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResponseDtoBoolean
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ServiceOverridesApiService Gets Service Overrides by Identifier
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
Passed from http.Request or context.Background().
 * @param identifier
 * @param accountIdentifier
 * @param optional nil or *ServiceOverridesApiGetServiceOverridesV2Opts - Optional Parameters:
     * @param "OrgIdentifier" (optional.String) -
     * @param "ProjectIdentifier" (optional.String) -
     * @param "Branch" (optional.String) - 
     * @param "RepoName" (optional.String) - 
     * @param "LoadFromCache" (optional.String) - 
     * @param "LoadFromFallbackBranch" (optional.Bool) - 
     * @param "GetMetadataOnly" (optional.Bool) - 
@return ResponseServiceOverridesResponseDtov2
*/

type ServiceOverridesApiGetServiceOverridesV2Opts struct {
	OrgIdentifier     optional.String
	ProjectIdentifier optional.String
	Branch optional.String
    RepoName optional.String
    LoadFromCache optional.String
    LoadFromFallbackBranch optional.Bool
    GetMetadataOnly optional.Bool
}

func (a *ServiceOverridesApiService) GetServiceOverridesV2(ctx context.Context, identifier string, accountIdentifier string, localVarOptionals *ServiceOverridesApiGetServiceOverridesV2Opts) (ResponseServiceOverridesResponseDtov2, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseServiceOverridesResponseDtov2
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/serviceOverrides/get-with-yaml/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", fmt.Sprintf("%v",
		identifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier",
			parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier",
			parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoName.IsSet() {
		localVarQueryParams.Add("repoName", parameterToString(localVarOptionals.RepoName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LoadFromFallbackBranch.IsSet() {
		localVarQueryParams.Add("loadFromFallbackBranch", parameterToString(localVarOptionals.LoadFromFallbackBranch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetMetadataOnly.IsSet() {
		localVarQueryParams.Add("getMetadataOnly", parameterToString(localVarOptionals.GetMetadataOnly.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody,
			localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResponseServiceOverridesResponseDtov2
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ServiceOverridesApiService Update an ServiceOverride Entity
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountIdentifier
 * @param optional nil or *ServiceOverridesApiUpdateServiceOverrideV2Opts - Optional Parameters:
     * @param "Body" (optional.Interface of ServiceOverrideRequestDtov2) - 
     * @param "Branch" (optional.String) - 
     * @param "RepoIdentifier" (optional.String) - 
     * @param "RootFolder" (optional.String) - 
     * @param "FilePath" (optional.String) - 
     * @param "CommitMsg" (optional.String) - 
     * @param "LastObjectId" (optional.String) - 
     * @param "ResolvedConflictCommitId" (optional.String) - 
     * @param "BaseBranch" (optional.String) - 
     * @param "ConnectorRef" (optional.String) - 
     * @param "StoreType" (optional.String) - 
     * @param "LastCommitId" (optional.String) - 
     * @param "IsNewBranch" (optional.Bool) - 
     * @param "IsHarnessCodeRepo" (optional.Bool) - 
@return ResponseServiceOverridesResponseDtov2
*/

type ServiceOverridesApiUpdateServiceOverrideV2Opts struct {
	Body optional.Interface
	Branch optional.String
    FilePath optional.String
    CommitMsg optional.String
    LastObjectId optional.String
    BaseBranch optional.String
    ConnectorRef optional.String
    StoreType optional.String
    LastCommitId optional.String
    IsNewBranch optional.Bool
    IsHarnessCodeRepo optional.Bool
}

func (a *ServiceOverridesApiService) UpdateServiceOverrideV2(ctx context.Context, accountIdentifier string, localVarOptionals *ServiceOverridesApiUpdateServiceOverrideV2Opts) (ResponseServiceOverridesResponseDtov2, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseServiceOverridesResponseDtov2
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/serviceOverrides"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilePath.IsSet() {
		localVarQueryParams.Add("filePath", parameterToString(localVarOptionals.FilePath.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CommitMsg.IsSet() {
		localVarQueryParams.Add("commitMsg", parameterToString(localVarOptionals.CommitMsg.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastObjectId.IsSet() {
		localVarQueryParams.Add("lastObjectId", parameterToString(localVarOptionals.LastObjectId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseBranch.IsSet() {
		localVarQueryParams.Add("baseBranch", parameterToString(localVarOptionals.BaseBranch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConnectorRef.IsSet() {
		localVarQueryParams.Add("connectorRef", parameterToString(localVarOptionals.ConnectorRef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StoreType.IsSet() {
		localVarQueryParams.Add("storeType", parameterToString(localVarOptionals.StoreType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LastCommitId.IsSet() {
		localVarQueryParams.Add("lastCommitId", parameterToString(localVarOptionals.LastCommitId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsNewBranch.IsSet() {
		localVarQueryParams.Add("isNewBranch", parameterToString(localVarOptionals.IsNewBranch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsHarnessCodeRepo.IsSet() {
		localVarQueryParams.Add("isHarnessCodeRepo", parameterToString(localVarOptionals.IsHarnessCodeRepo.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody,
			localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResponseServiceOverridesResponseDtov2
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody,
				localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ServiceOverridesApiService import Service Overrides from remote
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ServiceOverridesApiImportServiceOverridesOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ServiceOverrideImportRequestDto) - 
     * @param "AccountIdentifier" (optional.String) - 
     * @param "ConnectorRef" (optional.String) - 
     * @param "RepoName" (optional.String) - 
     * @param "Branch" (optional.String) - 
     * @param "FilePath" (optional.String) - 
     * @param "IsForceImport" (optional.Bool) - 
     * @param "IsHarnessCodeRepo" (optional.Bool) - 
@return ResponseServiceOverrideImportResponseDto
*/

type ServiceOverridesApiImportServiceOverridesOpts struct {
    Body optional.Interface
    AccountIdentifier optional.String
    ConnectorRef optional.String
    RepoName optional.String
    Branch optional.String
    FilePath optional.String
    IsForceImport optional.Bool
    IsHarnessCodeRepo optional.Bool
}

func (a *ServiceOverridesApiService) ImportServiceOverrides(ctx context.Context, accountIdentifier string, localVarOptionals *ServiceOverridesApiImportServiceOverridesOpts) (ResponseServiceOverrideImportResponseDto, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseServiceOverrideImportResponseDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/serviceOverrides/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConnectorRef.IsSet() {
		localVarQueryParams.Add("connectorRef", parameterToString(localVarOptionals.ConnectorRef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoName.IsSet() {
		localVarQueryParams.Add("repoName", parameterToString(localVarOptionals.RepoName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilePath.IsSet() {
		localVarQueryParams.Add("filePath", parameterToString(localVarOptionals.FilePath.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsForceImport.IsSet() {
		localVarQueryParams.Add("isForceImport", parameterToString(localVarOptionals.IsForceImport.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsHarnessCodeRepo.IsSet() {
		localVarQueryParams.Add("isHarnessCodeRepo", parameterToString(localVarOptionals.IsHarnessCodeRepo.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResponseServiceOverrideImportResponseDto
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
