package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Eliot120191/Banbif.DeudaTecnica.GoDomain/response"
	"github.com/Eliot120191/Banbif.DeudaTecnica.GoMobileUsers/internal/usuario"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewUserHTTPServer(ctx context.Context, endpoints usuario.Endpoints) http.Handler {

	r := mux.NewRouter()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Handle("/appmobile-users/v1/usuario", httptransport.NewServer(
		endpoint.Endpoint(endpoints.LoginAsync),
		decodeLoginAsync, encodeResponse,
		opts...,
	)).Methods("POST")
	/*
		r.Handle("/users", httptransport.NewServer(
			endpoint.Endpoint(endpoints.Create),
			decodeCreateUser, encodeResponse,
			opts...,
		)).Methods("POST")

		r.Handle("/users", httptransport.NewServer(
			endpoint.Endpoint(endpoints.GetAll),
			decodeGetAllUser,
			encodeResponse,
			opts...,
		)).Methods("GET")

		r.Handle("/users/{id}", httptransport.NewServer(
			endpoint.Endpoint(endpoints.Get),
			decodeGetUser,
			encodeResponse,
			opts...,
		)).Methods("GET")

		r.Handle("/users/{id}", httptransport.NewServer(
			endpoint.Endpoint(endpoints.Update),
			decodeUpdateUser,
			encodeResponse,
			opts...,
		)).Methods("PATCH")

		r.Handle("/users/{id}", httptransport.NewServer(
			endpoint.Endpoint(endpoints.Delete),
			decodeDeleteUser,
			encodeResponse,
			opts...,
		)).Methods("DELETE")
	*/
	return r
}

func decodeLoginAsync(_ context.Context, r *http.Request) (interface{}, error) {

	var req usuario.LoginAsyncRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

/*
func decodeCreateUser(_ context.Context, r *http.Request) (interface{}, error) {

	var req users.CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}

func decodeGetUser(_ context.Context, r *http.Request) (interface{}, error) {
	p := mux.Vars(r)
	intVal1, _ := strconv.ParseInt(p["id"], 0, 64)

	req := users.GetReq{
		ID: intVal1,
	}

	return req, nil
}

func decodeGetAllUser(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	limit, _ := strconv.Atoi(v.Get("limit"))
	page, _ := strconv.Atoi(v.Get("page"))

	req := users.GetAllReq{
		FirstName: v.Get("first_name"),
		LastName:  v.Get("last_name"),
		Limit:     limit,
		Page:      page,
	}

	return req, nil
}

func decodeUpdateUser(_ context.Context, r *http.Request) (interface{}, error) {
	var req users.UpdateReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	path := mux.Vars(r)
	intVal1, _ := strconv.ParseInt(path["id"], 0, 64)
	req.ID = intVal1

	return req, nil
}

func decodeDeleteUser(_ context.Context, r *http.Request) (interface{}, error) {

	path := mux.Vars(r)
	intVal1, _ := strconv.ParseInt(path["id"], 0, 64)
	req := users.DeleteReq{
		ID: intVal1,
	}

	return req, nil
}*/

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(response.Response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(r.StatusCode())
	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := err.(response.Response)
	w.WriteHeader(resp.StatusCode())
	_ = json.NewEncoder(w).Encode(resp)
}
