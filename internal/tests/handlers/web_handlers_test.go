package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/buelbuel/gowc/handlers"
	"github.com/buelbuel/gowc/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestStartPage(t *testing.T) {
	echo := echo.New()

	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	projectRoot := filepath.Join(workingDir, "../../../")
	err = os.Chdir(projectRoot)
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	renderer := utils.NewTemplates()
	echo.Renderer = renderer

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	context := echo.NewContext(req, rec)

	if assert.NoError(t, handlers.StartPageHandler(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		responseBody := rec.Body.String()
		t.Logf("Response Body: %s", responseBody)
		assert.Contains(t, responseBody, "Start")
	}
}
