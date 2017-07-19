package controllers

import (
	"krypton-server/utils"
	"net/http"
	"testing"
	"time"

	"github.com/golib/assert"
)

var (
	testTitle   = "testTitle"
	testContent = "testContent"
)

func Test_Toto(t *testing.T) {
	// creating todo item
	params := &CreateTodoParams{
		Title:   testTitle,
		Content: testContent,
		Due:     time.Now(),
	}

	r := utils.RequestJson(http.MethodPost, "/todo", params)
	resp := mockRequest(r)
	assertion := assert.New(t)
	assertion.Equal(200, resp.Status)
}