package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/govinda-attal/hello-world/handler"

	"github.com/govinda-attal/hello-world/pkg/core/status"
	"github.com/govinda-attal/hello-world/pkg/hw"
	"github.com/govinda-attal/hello-world/test/mocks"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("Hello World Handler's", func() {

	Describe("Basic Behaviour", func() {
		var (
			testRouter = func(path string, handler http.HandlerFunc) *mux.Router {
				r := mux.NewRouter()
				r.HandleFunc(path, handler)
				return r
			}
		)
		BeforeEach(func() {

		})

		Context("When requested hello with a given name X", func() {
			greetings := []byte(`{"Greetings":"Hello X !"}`)
			mock := &mocks.HelloWorldMock{}
			mock.HelloCall.Returns.HelloRs = &hw.HelloRs{}
			mock.HelloCall.Returns.HelloRs.Greetings = "Hello X !"
			It("must return greetings to X and no error", func() {
				rr := httptest.NewRecorder()
				hwhandler := NewHelloHandler(mock)
				handler := WrapperHandler(hwhandler.Hello)
				req, _ := http.NewRequest("GET", "/?name=X", nil)
				router := testRouter("/", handler)
				router.ServeHTTP(rr, req)
				Expect(rr.Code).To(Equal(http.StatusOK))
				respBytes, err := ioutil.ReadAll(rr.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(respBytes).To(MatchJSON(greetings))
			})
		})

		Context("When requested hello with a given name 'error'", func() {
			errStatusMsg := []byte(`{"code":400,"message":"must not be error"}`)
			mock := &mocks.HelloWorldMock{}
			mock.HelloCall.Returns.HelloRs = &hw.HelloRs{}
			mock.HelloCall.Returns.Error = status.ErrBadRequest.WithMessage("must not be error")
			It("must return an error", func() {
				rr := httptest.NewRecorder()
				hwhandler := NewHelloHandler(mock)
				handler := WrapperHandler(hwhandler.Hello)
				req, _ := http.NewRequest("GET", "/?name=error", nil)
				router := testRouter("/", handler)
				router.ServeHTTP(rr, req)
				Expect(rr.Code).To(Equal(http.StatusBadRequest))
				respBytes, err := ioutil.ReadAll(rr.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(respBytes).To(MatchJSON(errStatusMsg))
			})
		})

	})
})
