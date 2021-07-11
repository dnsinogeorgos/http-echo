package http_echo

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
)

type Request struct {
	Method           string
	URL              *url.URL
	Proto            string
	Header           http.Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Host             string
	Form             url.Values
	PostForm         url.Values
	MultipartForm    *multipart.Form
	Trailer          http.Header
	RemoteAddr       string
	RequestURI       string
	Response         *http.Response
}

func DumpRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = r.ParseMultipartForm(131072)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req := &Request{
		Method:           r.Method,
		URL:              r.URL,
		Proto:            r.Proto,
		Header:           r.Header,
		Body:             r.Body,
		ContentLength:    r.ContentLength,
		TransferEncoding: r.TransferEncoding,
		Host:             r.Host,
		Form:             r.Form,
		PostForm:         r.PostForm,
		MultipartForm:    r.MultipartForm,
		Trailer:          r.Trailer,
		RemoteAddr:       r.RemoteAddr,
		RequestURI:       r.RequestURI,
		Response:         r.Response,
	}

	render.JSON(w, r, req)
}
