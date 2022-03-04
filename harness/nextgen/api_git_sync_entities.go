/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
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

type GitSyncEntitiesApiService service

/*
GitSyncEntitiesApiService Lists Git Sync Entity filtered by their Type for the given scope
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entityType Entity Type
 * @param optional nil or *GitSyncEntitiesApiListGitSyncEntitiesByTypeOpts - Optional Parameters:
     * @param "ProjectIdentifier" (optional.String) -  Project Identifier for the Entity
     * @param "OrgIdentifier" (optional.String) -  Organization Identifier for the Entity
     * @param "AccountIdentifier" (optional.String) -  Account Identifier for the Entity
     * @param "GitSyncConfigId" (optional.String) -  Git Sync Config Id
     * @param "Branch" (optional.String) -  Branch Name
     * @param "Page" (optional.Int32) -  Indicates the number of pages. Results for these pages will be retrieved.
     * @param "Size" (optional.Int32) -  The number of the elements to fetch
     * @param "ModuleType" (optional.String) -  Module Type
@return ResponseDtoPageResponseGitSyncEntityList
*/

type GitSyncEntitiesApiListGitSyncEntitiesByTypeOpts struct {
	ProjectIdentifier optional.String
	OrgIdentifier     optional.String
	AccountIdentifier optional.String
	GitSyncConfigId   optional.String
	Branch            optional.String
	Page              optional.Int32
	Size              optional.Int32
	ModuleType        optional.String
}

func (a *GitSyncEntitiesApiService) ListGitSyncEntitiesByType(ctx context.Context, entityType string, localVarOptionals *GitSyncEntitiesApiListGitSyncEntitiesByTypeOpts) (ResponseDtoPageResponseGitSyncEntityList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoPageResponseGitSyncEntityList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/git-sync-entities/{entityType}"
	localVarPath = strings.Replace(localVarPath, "{"+"entityType"+"}", fmt.Sprintf("%v", entityType), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier", parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier", parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GitSyncConfigId.IsSet() {
		localVarQueryParams.Add("gitSyncConfigId", parameterToString(localVarOptionals.GitSyncConfigId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ModuleType.IsSet() {
		localVarQueryParams.Add("moduleType", parameterToString(localVarOptionals.ModuleType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml", "text/yaml", "text/html"}

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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoPageResponseGitSyncEntityList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
GitSyncEntitiesApiService Lists Git Sync Entity by product for the given Repo, Branch and list of Entity Types
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body This filters the Git Sync Entity based on multiple parameters
 * @param branch Branch Name
 * @param optional nil or *GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndBranchOpts - Optional Parameters:
     * @param "ProjectIdentifier" (optional.String) -  Project Identifier for the Entity
     * @param "OrgIdentifier" (optional.String) -  Organization Identifier for the Entity
     * @param "AccountIdentifier" (optional.String) -  Account Identifier for the Entity
     * @param "Size" (optional.Int32) -  The number of the elements to fetch
     * @param "GitSyncConfigId" (optional.String) -  Git Sync Config Id
@return ResponseDtoListGitSyncEntityList
*/

type GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndBranchOpts struct {
	ProjectIdentifier optional.String
	OrgIdentifier     optional.String
	AccountIdentifier optional.String
	Size              optional.Int32
	GitSyncConfigId   optional.String
}

func (a *GitSyncEntitiesApiService) ListGitSyncEntitiesSummaryForRepoAndBranch(ctx context.Context, body GitEntityBranchSummaryFilter, branch string, localVarOptionals *GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndBranchOpts) (ResponseDtoListGitSyncEntityList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoListGitSyncEntityList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/git-sync-entities/branch/{branch}"
	localVarPath = strings.Replace(localVarPath, "{"+"branch"+"}", fmt.Sprintf("%v", branch), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier", parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier", parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GitSyncConfigId.IsSet() {
		localVarQueryParams.Add("gitSyncConfigId", parameterToString(localVarOptionals.GitSyncConfigId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/yaml", "text/html"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml", "text/yaml", "text/html"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoListGitSyncEntityList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
GitSyncEntitiesApiService Lists Git Sync Entity by product for the given list of Repos and Entity Types
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body Filter Git Sync Entity based on multiple parameters
 * @param optional nil or *GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndTypesOpts - Optional Parameters:
     * @param "ProjectIdentifier" (optional.String) -  Project Identifier for the Entity
     * @param "OrgIdentifier" (optional.String) -  Organization Identifier for the Entity
     * @param "AccountIdentifier" (optional.String) -  Account Identifier for the Entity
     * @param "Size" (optional.Int32) -  The number of the elements to fetch
@return ResponseDtoGitSyncRepoFilesList
*/

type GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndTypesOpts struct {
	ProjectIdentifier optional.String
	OrgIdentifier     optional.String
	AccountIdentifier optional.String
	Size              optional.Int32
}

func (a *GitSyncEntitiesApiService) ListGitSyncEntitiesSummaryForRepoAndTypes(ctx context.Context, body GitEntitySummaryFilter, localVarOptionals *GitSyncEntitiesApiListGitSyncEntitiesSummaryForRepoAndTypesOpts) (ResponseDtoGitSyncRepoFilesList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ResponseDtoGitSyncRepoFilesList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/git-sync-entities/summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ProjectIdentifier.IsSet() {
		localVarQueryParams.Add("projectIdentifier", parameterToString(localVarOptionals.ProjectIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrgIdentifier.IsSet() {
		localVarQueryParams.Add("orgIdentifier", parameterToString(localVarOptionals.OrgIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountIdentifier.IsSet() {
		localVarQueryParams.Add("accountIdentifier", parameterToString(localVarOptionals.AccountIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/yaml", "text/html"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml", "text/yaml", "text/html"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoGitSyncRepoFilesList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
