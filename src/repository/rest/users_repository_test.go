package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredential(t *testing.T) {

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}
