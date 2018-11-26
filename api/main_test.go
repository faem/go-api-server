package api

import (
	"encoding/base64"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Reqs struct {
	Method             string
	URL                string
	Body               io.Reader
	ExpectedStatusCode int
	ExpectedResponse   string
}

func TestGetProfile(t *testing.T) {
	reqs := make([]Reqs, 4)
	reqs[0] = Reqs{
		"GET",
		"http://localhost:8080/in/fahim-abrar",
		nil,
		200,
		`{"id":"fahim-abrar","name":"Mohammad Fahim Abrar","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C++","noOfEndorsement":3},{"name":"Android","noOfEndorsement":4}]}`+"\n",
	}

	reqs[1] = Reqs{
		"GET",
		"http://localhost:8080/in/masud-rahman",
		nil,
		200,
		`{"id":"masud-rahman","name":"Masudur Rahman","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":3},{"name":"C++","noOfEndorsement":4}]}`+"\n",
	}

	reqs[2] = Reqs{
		"GET",
		"http://localhost:8080/in/mohan",
		nil,
		200,
		`{"id":"mohan","name":"Tahsin Rahman","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":100},{"name":"C++","noOfEndorsement":110},{"name":"Linux","noOfEndorsement":100}]}`+"\n",
	}

	reqs[3] = Reqs{
		"GET",
		"http://localhost:8080/in/mohand",
		nil,
		404,
		"Profile of mohand not found!",
	}

	processRequest(t, reqs)
}

//var expectedStsCode []int
func TestGetProfiles(t *testing.T) {
	reqs := make([]Reqs, 2)
	reqs[0] = Reqs{
		"GET",
		"http://localhost:8080/in",
		nil,
		200,
		`[{"id":"fahim-abrar","name":"Mohammad Fahim Abrar","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C++","noOfEndorsement":3},{"name":"Android","noOfEndorsement":4}]},{"id":"masud-rahman","name":"Masudur Rahman","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":3},{"name":"C++","noOfEndorsement":4}]},{"id":"mohan","name":"Tahsin Rahman","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":100},{"name":"C++","noOfEndorsement":110},{"name":"Linux","noOfEndorsement":100}]}]`,
	}

	reqs[1] = Reqs{
		"GET",
		"http://localhost:8080/in/",
		nil,
		404,
		"404 page not found\n",
	}

	processRequest(t, reqs)
}

func TestAddProfile(t *testing.T) {
	reqs := make([]Reqs, 2)
	reqs[0] = Reqs{
		"POST",
		"http://localhost:8080/in",
		strings.NewReader(`{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":100},{"name":"C++","noOfEndorsement":100}]}`),
		200,
		"",
	}

	reqs[1] = Reqs{
		"POST",
		"http://localhost:8080/in",
		strings.NewReader(`{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":100},{"name":"C++","noOfEndorsement":101}]}`),
		409,
		"Username kfoozminus already exists!",
	}

	processRequest(t,reqs)
}

func TestDeleteProfile(t *testing.T) {
	reqs := make([]Reqs, 3)
	reqs[0] = Reqs{
		"DELETE",
		"http://localhost:8080/in/fahim-abrar",
		nil,
		200,
		"",
	}

	reqs[1] = Reqs{
		"DELETE",
		"http://localhost:8080/in/fahim-abrar",
		nil,
		404,
		"Profile of fahim-abrar not found!",
	}

	reqs[2] = Reqs{
		"DELETE",
		"http://localhost:8080/in/mohand",
		nil,
		404,
		"Profile of mohand not found!",
	}

	processRequest(t, reqs)
}

func TestUpdateProfile(t *testing.T) {
	reqs := make([]Reqs, 2)
	reqs[0] = Reqs{
		"PUT",
		"http://localhost:8080/in/mohan",
		strings.NewReader(`{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":100},{"name":"C++","noOfEndorsement":100}]}`),
		200,
		"",
	}

	reqs[1] = Reqs{
		"PUT",
		"http://localhost:8080/in/kfoozminus",
		strings.NewReader(`{"id":"kfoozminus","name":"Jannatul Ferdous","company":"AppsCode Inc.","position":"Software Engineer","skill":[{"name":"C","noOfEndorsement":100},{"name":"C++","noOfEndorsement":101}]}`),
		404,
		"Profile of kfoozminus not found!\n",
	}

	processRequest(t,reqs)
}

func processRequest(t *testing.T, reqs []Reqs) {
	for _, req := range reqs {
		r, _ := http.NewRequest(req.Method, req.URL, req.Body)
		r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("fahim:1234")))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		if w.Code != req.ExpectedStatusCode{
			t.Error("\nExpected Status Code\t= "+cast.ToString(req.ExpectedStatusCode)+"\nFound Status Code\t\t= "+cast.ToString(w.Code)+"\n")

		}

		if cast.ToString(w.Body) != req.ExpectedResponse{
			t.Error("\nExpected Response\t= "+req.ExpectedResponse+"\nFound Response\t\t= "+cast.ToString(w.Body)+"\n")
		}
	}
}