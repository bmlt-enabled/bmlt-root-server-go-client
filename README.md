# Go API client for bmlt

BMLT Admin API Documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import bmlt "github.com/bmlt-enabled/bmlt-root-server-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), bmlt.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), bmlt.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), bmlt.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), bmlt.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8000/main_server*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*RootServerApi* | [**AuthLogout**](docs/RootServerApi.md#authlogout) | **Post** /api/v1/auth/logout | Revokes a token
*RootServerApi* | [**AuthRefresh**](docs/RootServerApi.md#authrefresh) | **Post** /api/v1/auth/refresh | Revokes and issues a new token
*RootServerApi* | [**AuthToken**](docs/RootServerApi.md#authtoken) | **Post** /api/v1/auth/token | Creates a token
*RootServerApi* | [**CreateFormat**](docs/RootServerApi.md#createformat) | **Post** /api/v1/formats | Creates a format
*RootServerApi* | [**CreateMeeting**](docs/RootServerApi.md#createmeeting) | **Post** /api/v1/meetings | Creates a meeting
*RootServerApi* | [**CreateServiceBody**](docs/RootServerApi.md#createservicebody) | **Post** /api/v1/servicebodies | Creates a service body
*RootServerApi* | [**CreateUser**](docs/RootServerApi.md#createuser) | **Post** /api/v1/users | Creates a user
*RootServerApi* | [**DeleteFormat**](docs/RootServerApi.md#deleteformat) | **Delete** /api/v1/formats/{formatId} | Deletes a format
*RootServerApi* | [**DeleteMeeting**](docs/RootServerApi.md#deletemeeting) | **Delete** /api/v1/meetings/{meetingId} | Deletes a meeting
*RootServerApi* | [**DeleteServiceBody**](docs/RootServerApi.md#deleteservicebody) | **Delete** /api/v1/servicebodies/{serviceBodyId} | Deletes a service body
*RootServerApi* | [**DeleteUser**](docs/RootServerApi.md#deleteuser) | **Delete** /api/v1/users/{userId} | Deletes a user
*RootServerApi* | [**GetFormat**](docs/RootServerApi.md#getformat) | **Get** /api/v1/formats/{formatId} | Retrieves a format
*RootServerApi* | [**GetFormats**](docs/RootServerApi.md#getformats) | **Get** /api/v1/formats | Retrieves formats
*RootServerApi* | [**GetMeeting**](docs/RootServerApi.md#getmeeting) | **Get** /api/v1/meetings/{meetingId} | Retrieves a meeting
*RootServerApi* | [**GetMeetings**](docs/RootServerApi.md#getmeetings) | **Get** /api/v1/meetings | Retrieves meetings
*RootServerApi* | [**GetServiceBodies**](docs/RootServerApi.md#getservicebodies) | **Get** /api/v1/servicebodies | Retrieves service bodies
*RootServerApi* | [**GetServiceBody**](docs/RootServerApi.md#getservicebody) | **Get** /api/v1/servicebodies/{serviceBodyId} | Retrieves a service body
*RootServerApi* | [**GetUser**](docs/RootServerApi.md#getuser) | **Get** /api/v1/users/{userId} | Retrieves a single user
*RootServerApi* | [**GetUsers**](docs/RootServerApi.md#getusers) | **Get** /api/v1/users | Retrieves users
*RootServerApi* | [**PartialUpdateUser**](docs/RootServerApi.md#partialupdateuser) | **Patch** /api/v1/users/{userId} | Patches a user
*RootServerApi* | [**PatchFormat**](docs/RootServerApi.md#patchformat) | **Patch** /api/v1/formats/{formatId} | Patches a format
*RootServerApi* | [**PatchMeeting**](docs/RootServerApi.md#patchmeeting) | **Patch** /api/v1/meetings/{meetingId} | Patches a meeting
*RootServerApi* | [**PatchServiceBody**](docs/RootServerApi.md#patchservicebody) | **Patch** /api/v1/servicebodies/{serviceBodyId} | Patches a service body
*RootServerApi* | [**UpdateFormat**](docs/RootServerApi.md#updateformat) | **Put** /api/v1/formats/{formatId} | Updates a format
*RootServerApi* | [**UpdateMeeting**](docs/RootServerApi.md#updatemeeting) | **Put** /api/v1/meetings/{meetingId} | Updates a meeting
*RootServerApi* | [**UpdateServiceBody**](docs/RootServerApi.md#updateservicebody) | **Put** /api/v1/servicebodies/{serviceBodyId} | Updates a Service Body
*RootServerApi* | [**UpdateUser**](docs/RootServerApi.md#updateuser) | **Put** /api/v1/users/{userId} | Update single user


## Documentation For Models

 - [AuthenticationError](docs/AuthenticationError.md)
 - [AuthorizationError](docs/AuthorizationError.md)
 - [Format](docs/Format.md)
 - [FormatAllOf](docs/FormatAllOf.md)
 - [FormatBase](docs/FormatBase.md)
 - [FormatCreate](docs/FormatCreate.md)
 - [FormatPartialUpdate](docs/FormatPartialUpdate.md)
 - [FormatTranslation](docs/FormatTranslation.md)
 - [FormatUpdate](docs/FormatUpdate.md)
 - [Meeting](docs/Meeting.md)
 - [MeetingBase](docs/MeetingBase.md)
 - [MeetingCreate](docs/MeetingCreate.md)
 - [MeetingPartialUpdate](docs/MeetingPartialUpdate.md)
 - [MeetingUpdate](docs/MeetingUpdate.md)
 - [NotFoundError](docs/NotFoundError.md)
 - [ServiceBody](docs/ServiceBody.md)
 - [ServiceBodyBase](docs/ServiceBodyBase.md)
 - [ServiceBodyCreate](docs/ServiceBodyCreate.md)
 - [ServiceBodyPartialUpdate](docs/ServiceBodyPartialUpdate.md)
 - [ServiceBodyUpdate](docs/ServiceBodyUpdate.md)
 - [Token](docs/Token.md)
 - [TokenCredentials](docs/TokenCredentials.md)
 - [User](docs/User.md)
 - [UserBase](docs/UserBase.md)
 - [UserCreate](docs/UserCreate.md)
 - [UserCreateAllOf](docs/UserCreateAllOf.md)
 - [UserPartialUpdate](docs/UserPartialUpdate.md)
 - [UserUpdate](docs/UserUpdate.md)
 - [ValidationError](docs/ValidationError.md)


## Documentation For Authorization



### bmltToken


- **Type**: OAuth
- **Flow**: password
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


