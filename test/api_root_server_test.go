/*
BMLT

Testing RootServerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package bmlt

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_bmlt_RootServerApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RootServerApiService AuthLogout", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.AuthLogout(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService AuthRefresh", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.AuthRefresh(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService AuthToken", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.AuthToken(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService CreateFormat", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.CreateFormat(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService CreateMeeting", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.CreateMeeting(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService CreateServiceBody", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.CreateServiceBody(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService CreateUser", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.CreateUser(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService DeleteFormat", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var formatId int64

        resp, httpRes, err := apiClient.RootServerApi.DeleteFormat(context.Background(), formatId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService DeleteMeeting", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var meetingId int64

        resp, httpRes, err := apiClient.RootServerApi.DeleteMeeting(context.Background(), meetingId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService DeleteServiceBody", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceBodyId int64

        resp, httpRes, err := apiClient.RootServerApi.DeleteServiceBody(context.Background(), serviceBodyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService DeleteUser", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var userId int64

        resp, httpRes, err := apiClient.RootServerApi.DeleteUser(context.Background(), userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetFormat", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var formatId int64

        resp, httpRes, err := apiClient.RootServerApi.GetFormat(context.Background(), formatId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetFormats", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.GetFormats(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetMeeting", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var meetingId int64

        resp, httpRes, err := apiClient.RootServerApi.GetMeeting(context.Background(), meetingId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetMeetings", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.GetMeetings(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetServiceBodies", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.GetServiceBodies(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetServiceBody", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceBodyId int64

        resp, httpRes, err := apiClient.RootServerApi.GetServiceBody(context.Background(), serviceBodyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetUser", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var userId int64

        resp, httpRes, err := apiClient.RootServerApi.GetUser(context.Background(), userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService GetUsers", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RootServerApi.GetUsers(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService PartialUpdateUser", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var userId int64

        resp, httpRes, err := apiClient.RootServerApi.PartialUpdateUser(context.Background(), userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService PatchFormat", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var formatId int64

        resp, httpRes, err := apiClient.RootServerApi.PatchFormat(context.Background(), formatId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService PatchMeeting", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var meetingId int64

        resp, httpRes, err := apiClient.RootServerApi.PatchMeeting(context.Background(), meetingId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService PatchServiceBody", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceBodyId int64

        resp, httpRes, err := apiClient.RootServerApi.PatchServiceBody(context.Background(), serviceBodyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService UpdateFormat", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var formatId int64

        resp, httpRes, err := apiClient.RootServerApi.UpdateFormat(context.Background(), formatId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService UpdateMeeting", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var meetingId int64

        resp, httpRes, err := apiClient.RootServerApi.UpdateMeeting(context.Background(), meetingId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService UpdateServiceBody", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceBodyId int64

        resp, httpRes, err := apiClient.RootServerApi.UpdateServiceBody(context.Background(), serviceBodyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RootServerApiService UpdateUser", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var userId int64

        resp, httpRes, err := apiClient.RootServerApi.UpdateUser(context.Background(), userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
