
/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type EnvironmentGroupApiService service
/*
EnvironmentGroupApiService Delete en Environment Group by Identifier
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envGroupIdentifier Environment Group Identifier for the entity
 * @param accountIdentifier Account Identifier for the Entity.
 * @param orgIdentifier Organization Identifier for the Entity.
 * @param projectIdentifier Project Identifier for the Entity.
 * @param optional nil or *EnvironmentGroupApiDeleteEnvironmentGroupOpts - Optional Parameters:
     * @param "IfMatch" (optional.String) - 
     * @param "Branch" (optional.String) -  Name of the branch.
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Id.
     * @param "RootFolder" (optional.String) -  Path to the root folder of the Entity.
     * @param "FilePath" (optional.String) -  File Path of the Entity.
     * @param "CommitMsg" (optional.String) -  Commit Message to use for the merge commit.
     * @param "LastObjectId" (optional.String) -  Last Object Id
@return ResponseDtoEnvironmentGroupDelete
*/

type EnvironmentGroupApiDeleteEnvironmentGroupOpts struct {
    IfMatch optional.String
    Branch optional.String
    RepoIdentifier optional.String
    RootFolder optional.String
    FilePath optional.String
    CommitMsg optional.String
    LastObjectId optional.String
}

func (a *EnvironmentGroupApiService) DeleteEnvironmentGroup(ctx context.Context, envGroupIdentifier string, accountIdentifier string, orgIdentifier string, projectIdentifier string, localVarOptionals *EnvironmentGroupApiDeleteEnvironmentGroupOpts) (ResponseDtoEnvironmentGroupDelete, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoEnvironmentGroupDelete
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/environmentGroup/{envGroupIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"envGroupIdentifier"+"}", fmt.Sprintf("%v", envGroupIdentifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	localVarQueryParams.Add("orgIdentifier", parameterToString(orgIdentifier, ""))
	localVarQueryParams.Add("projectIdentifier", parameterToString(projectIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RootFolder.IsSet() {
		localVarQueryParams.Add("rootFolder", parameterToString(localVarOptionals.RootFolder.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.IfMatch.IsSet() {
		localVarHeaderParams["If-Match"] = parameterToString(localVarOptionals.IfMatch.Value(), "")
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
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoEnvironmentGroupDelete
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
/*
EnvironmentGroupApiService Gets an Environment Group by identifier
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param envGroupIdentifier Environment Group Identifier for the entity
 * @param accountIdentifier Account Identifier for the Entity.
 * @param orgIdentifier Organization Identifier for the Entity.
 * @param projectIdentifier Project Identifier for the Entity.
 * @param optional nil or *EnvironmentGroupApiGetEnvironmentGroupOpts - Optional Parameters:
     * @param "Deleted" (optional.Bool) -  Specify whether Environment is deleted or not
     * @param "Branch" (optional.String) -  Name of the branch.
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Id.
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
@return ResponseDtoEnvironmentGroup
*/

type EnvironmentGroupApiGetEnvironmentGroupOpts struct {
    Deleted optional.Bool
    Branch optional.String
    RepoIdentifier optional.String
    GetDefaultFromOtherRepo optional.Bool
}

func (a *EnvironmentGroupApiService) GetEnvironmentGroup(ctx context.Context, envGroupIdentifier string, accountIdentifier string, orgIdentifier string, projectIdentifier string, localVarOptionals *EnvironmentGroupApiGetEnvironmentGroupOpts) (ResponseDtoEnvironmentGroup, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoEnvironmentGroup
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/environmentGroup/{envGroupIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"envGroupIdentifier"+"}", fmt.Sprintf("%v", envGroupIdentifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	localVarQueryParams.Add("orgIdentifier", parameterToString(orgIdentifier, ""))
	localVarQueryParams.Add("projectIdentifier", parameterToString(projectIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.Deleted.IsSet() {
		localVarQueryParams.Add("deleted", parameterToString(localVarOptionals.Deleted.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
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
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoEnvironmentGroup
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
/*
EnvironmentGroupApiService Gets Environment Group list for a Project
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountIdentifier Account Identifier for the Entity.
 * @param orgIdentifier Organization Identifier for the Entity.
 * @param projectIdentifier Project Identifier for the Entity.
 * @param optional nil or *EnvironmentGroupApiGetEnvironmentGroupListOpts - Optional Parameters:
     * @param "Body" (optional.Interface of FilterProperties) -  This is the body for the filter properties for listing Environment Groups
     * @param "EnvGroupIdentifiers" (optional.Interface of []string) - 
     * @param "SearchTerm" (optional.String) -  The word to be searched and included in the list response
     * @param "Page" (optional.Int32) -  Page Index of the results to fetch.Default Value: 0
     * @param "Size" (optional.Int32) -  Results per page
     * @param "Sort" (optional.Interface of []string) -  Sort criteria for the elements.
     * @param "FilterIdentifier" (optional.String) -  Filter identifier
     * @param "Branch" (optional.String) -  Name of the branch.
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Id.
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
@return ResponseDtoPageResponseEnvironmentGroup
*/

type EnvironmentGroupApiGetEnvironmentGroupListOpts struct {
    Body optional.Interface
    EnvGroupIdentifiers optional.Interface
    SearchTerm optional.String
    Page optional.Int32
    Size optional.Int32
    Sort optional.Interface
    FilterIdentifier optional.String
    Branch optional.String
    RepoIdentifier optional.String
    GetDefaultFromOtherRepo optional.Bool
}

func (a *EnvironmentGroupApiService) GetEnvironmentGroupList(ctx context.Context, accountIdentifier string, orgIdentifier string, projectIdentifier string, localVarOptionals *EnvironmentGroupApiGetEnvironmentGroupListOpts) (ResponseDtoPageResponseEnvironmentGroup, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoPageResponseEnvironmentGroup
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/environmentGroup/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	localVarQueryParams.Add("orgIdentifier", parameterToString(orgIdentifier, ""))
	localVarQueryParams.Add("projectIdentifier", parameterToString(projectIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.EnvGroupIdentifiers.IsSet() {
		localVarQueryParams.Add("envGroupIdentifiers", parameterToString(localVarOptionals.EnvGroupIdentifiers.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("searchTerm", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.FilterIdentifier.IsSet() {
		localVarQueryParams.Add("filterIdentifier", parameterToString(localVarOptionals.FilterIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
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
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoPageResponseEnvironmentGroup
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
/*
EnvironmentGroupApiService Create an Environment Group
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountIdentifier Account Identifier for the Entity.
 * @param optional nil or *EnvironmentGroupApiPostEnvironmentGroupOpts - Optional Parameters:
     * @param "Body" (optional.Interface of EnvironmentGroupRequest) -  Details of the Environment Group to be created
     * @param "Branch" (optional.String) -  Name of the branch.
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Id.
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
@return ResponseDtoEnvironmentGroup
*/

type EnvironmentGroupApiPostEnvironmentGroupOpts struct {
    Body optional.Interface
    Branch optional.String
    RepoIdentifier optional.String
    GetDefaultFromOtherRepo optional.Bool
}

func (a *EnvironmentGroupApiService) PostEnvironmentGroup(ctx context.Context, accountIdentifier string, localVarOptionals *EnvironmentGroupApiPostEnvironmentGroupOpts) (ResponseDtoEnvironmentGroup, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoEnvironmentGroup
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/environmentGroup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
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
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoEnvironmentGroup
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
/*
EnvironmentGroupApiService Update an Environment Group by Identifier
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountIdentifier Account Identifier for the Entity.
 * @param envGroupIdentifier Environment Group Identifier for the entity
 * @param optional nil or *EnvironmentGroupApiUpdateEnvironmentGroupOpts - Optional Parameters:
     * @param "Body" (optional.Interface of EnvironmentGroupRequest) -  Details of the Environment Group to be updated
     * @param "IfMatch" (optional.String) - 
     * @param "Branch" (optional.String) -  Name of the branch.
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Id.
     * @param "RootFolder" (optional.String) -  Path to the root folder of the Entity.
     * @param "FilePath" (optional.String) -  Path to the root folder of the Entity.
     * @param "CommitMsg" (optional.String) -  Commit Message to use for the merge commit.
     * @param "LastObjectId" (optional.String) -  Last Object Id
     * @param "ResolvedConflictCommitId" (optional.String) -  If the entity is git-synced, this parameter represents the commit id against which file conflicts are resolved
     * @param "BaseBranch" (optional.String) -  Name of the default branch.
     * @param "ConnectorRef" (optional.String) -  Identifier of Connector needed for CRUD operations on the respective Entity
@return ResponseDtoEnvironmentGroup
*/

type EnvironmentGroupApiUpdateEnvironmentGroupOpts struct {
    Body optional.Interface
    IfMatch optional.String
    Branch optional.String
    RepoIdentifier optional.String
    RootFolder optional.String
    FilePath optional.String
    CommitMsg optional.String
    LastObjectId optional.String
    ResolvedConflictCommitId optional.String
    BaseBranch optional.String
    ConnectorRef optional.String
}

func (a *EnvironmentGroupApiService) UpdateEnvironmentGroup(ctx context.Context, accountIdentifier string, envGroupIdentifier string, localVarOptionals *EnvironmentGroupApiUpdateEnvironmentGroupOpts) (ResponseDtoEnvironmentGroup, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoEnvironmentGroup
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ng/api/environmentGroup/{envGroupIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"envGroupIdentifier"+"}", fmt.Sprintf("%v", envGroupIdentifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RootFolder.IsSet() {
		localVarQueryParams.Add("rootFolder", parameterToString(localVarOptionals.RootFolder.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.ResolvedConflictCommitId.IsSet() {
		localVarQueryParams.Add("resolvedConflictCommitId", parameterToString(localVarOptionals.ResolvedConflictCommitId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BaseBranch.IsSet() {
		localVarQueryParams.Add("baseBranch", parameterToString(localVarOptionals.BaseBranch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConnectorRef.IsSet() {
		localVarQueryParams.Add("connectorRef", parameterToString(localVarOptionals.ConnectorRef.Value(), ""))
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
	if localVarOptionals != nil && localVarOptionals.IfMatch.IsSet() {
		localVarHeaderParams["If-Match"] = parameterToString(localVarOptionals.IfMatch.Value(), "")
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
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoEnvironmentGroup
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
