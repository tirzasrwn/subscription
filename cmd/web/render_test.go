package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(r)
	r = r.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "flash")
	testApp.Session.Put(ctx, "error", "error")
	testApp.Session.Put(ctx, "warning", "warning")

	td := testApp.AddDefaultData(&TemplateData{}, r)
	if td.Flash != "flash" {
		t.Error("failed to get flash data")
	}
	if td.Error != "error" {
		t.Error("failed to get error data")
	}
	if td.Warning != "warning" {
		t.Error("failed to get warning data")
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(r)
	r = r.WithContext(ctx)

	auth := testApp.IsAuthenticated(r)
	if auth {
		t.Error("return true for authenticated, when it should be false")
	}

	testApp.Session.Put(ctx, "userID", 1)

	auth = testApp.IsAuthenticated(r)
	if !auth {
		t.Error("return false for authenticated, when it should be true")
	}
}

func TestConfing_render(t *testing.T) {
	pathToTemplate = "./templates"
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(r)
	r = r.WithContext(ctx)
	testApp.render(rr, r, "home.page.gohtml", &TemplateData{})
	if rr.Code != 200 {
		t.Error("failed to render page")
	}
}
