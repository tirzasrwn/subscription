package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"subscription/data"
	"testing"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHtml       string
}{
	{
		name:               "home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
	},
	{
		name:               "login page",
		url:                "/login",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.LoginPage,
		expectedHtml:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "logout",
		url:                "/logout",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func Test_Pages(t *testing.T) {
	pathToTemplate = "./templates"

	for _, page := range pageTests {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", page.url, nil)

		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(page.sessionData) > 0 {
			for key, value := range page.sessionData {
				testApp.Session.Put(ctx, key, value)
			}
		}
		page.handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("%s failed: expected %d but got %d", page.name, page.expectedStatusCode, rr.Code)
		}

		if len(page.expectedHtml) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, page.expectedHtml) {
				t.Errorf("%s failed: expected to find %s, but did not", page.name, page.expectedHtml)
			}
		}

	}

}
