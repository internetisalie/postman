// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-01-04 17:21:26.918888 -0800 PST m=+0.000341635
// 
// This contains Go HTTP testing rendered from a postman JSON export

package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// HTTPTest contains all the parameters for a HTTP Unit Test
type HTTPTest struct {
    Name string
    URL string
    Method string
    Body string
    ExpectedStatus int
    ExpectedContains []string
}

// RunHTTPTest accepts a HTTPTest type to execute the HTTP request
func RunHTTPTest(test HTTPTest) (*http.Response, error) {
    req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.Body))
    if err != nil {
        return nil, err
    }
    rr := httptest.NewRecorder()
    Router().ServeHTTP(rr, req)
    return rr.Result(), err
}

// TestPostmanMain runs all the Postman tests within the Main group.
func TestPostmanMain(t *testing.T) {

	tests := []HTTPTest{
		{
            Name: "Statup Details",
            URL: "https://api.elementchains.com/api",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        }}

    for _, v := range tests {
        t.Run(v.Name, func(t *testing.T) {
            run, err := RunHTTPTest(v)
            assert.Nil(t, err)
            body, err := ioutil.ReadAll(run.Body)
            assert.Nil(t, err)
            assert.Equal(t, v.ExpectedStatus, run.StatusCode)
            if v.ExpectedContains != nil {
                stringBody := string(body)
                for _, c := range v.ExpectedContains {
                    assert.Contains(t, stringBody, c)
                }
            }
            t.Logf("Test %v got: %v\n", v.Name, string(body))
        })
	}

}

// TestPostmanServices runs all the Postman tests within the Services group.
func TestPostmanServices(t *testing.T) {

	tests := []HTTPTest{
		{
            Name: "All Services",
            URL: "https://api.elementchains.com/api/services",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View Service",
            URL: "https://api.elementchains.com/api/services/1",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View Service Chart Data",
            URL: "https://api.elementchains.com/api/services/1/data",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Reorder Services",
            URL: "https://api.elementchains.com/api/services/reorder",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Create Service",
            URL: "https://api.elementchains.com/api/services",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Update Service",
            URL: "https://api.elementchains.com/api/services/{{service_id}}",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Delete Service",
            URL: "https://api.elementchains.com/api/services/{{service_id}}",
            Method: "DELETE",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        }}

    for _, v := range tests {
        t.Run(v.Name, func(t *testing.T) {
            run, err := RunHTTPTest(v)
            assert.Nil(t, err)
            body, err := ioutil.ReadAll(run.Body)
            assert.Nil(t, err)
            assert.Equal(t, v.ExpectedStatus, run.StatusCode)
            if v.ExpectedContains != nil {
                stringBody := string(body)
                for _, c := range v.ExpectedContains {
                    assert.Contains(t, stringBody, c)
                }
            }
            t.Logf("Test %v got: %v\n", v.Name, string(body))
        })
	}

}

// TestPostmanGroups runs all the Postman tests within the Groups group.
func TestPostmanGroups(t *testing.T) {

	tests := []HTTPTest{
		{
            Name: "All Groups",
            URL: "https://api.elementchains.com/api/groups",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View Group",
            URL: "https://api.elementchains.com/api/groups/1",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Create Group",
            URL: "https://api.elementchains.com/api/groups",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Delete Group",
            URL: "https://api.elementchains.com/api/groups/{{group_id}}",
            Method: "DELETE",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        }}

    for _, v := range tests {
        t.Run(v.Name, func(t *testing.T) {
            run, err := RunHTTPTest(v)
            assert.Nil(t, err)
            body, err := ioutil.ReadAll(run.Body)
            assert.Nil(t, err)
            assert.Equal(t, v.ExpectedStatus, run.StatusCode)
            if v.ExpectedContains != nil {
                stringBody := string(body)
                for _, c := range v.ExpectedContains {
                    assert.Contains(t, stringBody, c)
                }
            }
            t.Logf("Test %v got: %v\n", v.Name, string(body))
        })
	}

}

// TestPostmanUsers runs all the Postman tests within the Users group.
func TestPostmanUsers(t *testing.T) {

	tests := []HTTPTest{
		{
            Name: "View All Users",
            URL: "https://api.elementchains.com/api/users",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Create User",
            URL: "https://api.elementchains.com/api/users",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View User",
            URL: "https://api.elementchains.com/api/users/{{user_id}}",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Update User",
            URL: "https://api.elementchains.com/api/users/{{user_id}}",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Delete User",
            URL: "https://api.elementchains.com/api/users/{{user_id}}",
            Method: "DELETE",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        }}

    for _, v := range tests {
        t.Run(v.Name, func(t *testing.T) {
            run, err := RunHTTPTest(v)
            assert.Nil(t, err)
            body, err := ioutil.ReadAll(run.Body)
            assert.Nil(t, err)
            assert.Equal(t, v.ExpectedStatus, run.StatusCode)
            if v.ExpectedContains != nil {
                stringBody := string(body)
                for _, c := range v.ExpectedContains {
                    assert.Contains(t, stringBody, c)
                }
            }
            t.Logf("Test %v got: %v\n", v.Name, string(body))
        })
	}

}

// TestPostmanNotifiers runs all the Postman tests within the Notifiers group.
func TestPostmanNotifiers(t *testing.T) {

	tests := []HTTPTest{
		{
            Name: "All Notifiers",
            URL: "https://api.elementchains.com/api/notifiers",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View Notifier",
            URL: "https://api.elementchains.com/api/notifier/mobile",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Update Notifier",
            URL: "https://api.elementchains.com/api/notifier/mobile",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        }}

    for _, v := range tests {
        t.Run(v.Name, func(t *testing.T) {
            run, err := RunHTTPTest(v)
            assert.Nil(t, err)
            body, err := ioutil.ReadAll(run.Body)
            assert.Nil(t, err)
            assert.Equal(t, v.ExpectedStatus, run.StatusCode)
            if v.ExpectedContains != nil {
                stringBody := string(body)
                for _, c := range v.ExpectedContains {
                    assert.Contains(t, stringBody, c)
                }
            }
            t.Logf("Test %v got: %v\n", v.Name, string(body))
        })
	}

}

// TestPostmanMessages runs all the Postman tests within the Messages group.
func TestPostmanMessages(t *testing.T) {

	tests := []HTTPTest{
		{
            Name: "All Messages",
            URL: "https://api.elementchains.com/api/messages",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Create Message",
            URL: "https://api.elementchains.com/api/messages",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View Message",
            URL: "https://api.elementchains.com/api/messages/{{message_id}}",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Update Message",
            URL: "https://api.elementchains.com/api/messages/{{message_id}}",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Delete Message",
            URL: "https://api.elementchains.com/api/messages/{{message_id}}",
            Method: "DELETE",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        }}

    for _, v := range tests {
        t.Run(v.Name, func(t *testing.T) {
            run, err := RunHTTPTest(v)
            assert.Nil(t, err)
            body, err := ioutil.ReadAll(run.Body)
            assert.Nil(t, err)
            assert.Equal(t, v.ExpectedStatus, run.StatusCode)
            if v.ExpectedContains != nil {
                stringBody := string(body)
                for _, c := range v.ExpectedContains {
                    assert.Contains(t, stringBody, c)
                }
            }
            t.Logf("Test %v got: %v\n", v.Name, string(body))
        })
	}

}

// TestPostmanCheckins runs all the Postman tests within the Checkins group.
func TestPostmanCheckins(t *testing.T) {

	tests := []HTTPTest{
		{
            Name: "Create Checkin",
            URL: "https://api.elementchains.com/api/checkin",
            Method: "POST",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View All Checkin's",
            URL: "https://api.elementchains.com/api/checkins",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Run Checkin",
            URL: "https://api.elementchains.com/checkin/{{checkin_id}}",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "View Checkin",
            URL: "https://api.elementchains.com/api/checkin/{{checkin_id}}",
            Method: "GET",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        },
		{
            Name: "Delete Checkin",
            URL: "https://api.elementchains.com/api/checkin/{{checkin_id}}",
            Method: "DELETE",
            Body: nil,
            ExpectedStatus: 200,
            ExpectedContains: nil,
        }}

    for _, v := range tests {
        t.Run(v.Name, func(t *testing.T) {
            run, err := RunHTTPTest(v)
            assert.Nil(t, err)
            body, err := ioutil.ReadAll(run.Body)
            assert.Nil(t, err)
            assert.Equal(t, v.ExpectedStatus, run.StatusCode)
            if v.ExpectedContains != nil {
                stringBody := string(body)
                for _, c := range v.ExpectedContains {
                    assert.Contains(t, stringBody, c)
                }
            }
            t.Logf("Test %v got: %v\n", v.Name, string(body))
        })
	}

}
