package usuario

import (
	"context"
	"net/http"

	"github.com/Eliot120191/Banbif.DeudaTecnica.GoDomain/meta"
	"github.com/Eliot120191/Banbif.DeudaTecnica.GoDomain/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		LoginAsync Controller
	}

	LoginAsyncRequest struct {
		Correo string `json:"correo"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}

	Config struct {
		LimPageDef string
	}
)

func MakeEndpoints(s Service, config Config) Endpoints {

	return Endpoints{
		LoginAsync: makeLoginAsyncEndpoint(s),
	}

}

/*
	func makeCreateEndpoint(s Service) Controller {
		return func(ctx context.Context, request interface{}) (interface{}, error) {

			req := request.(CreateReq)

			if req.FirstName == "" {
				return nil, response.BadRequest(ErrFirstNameRequired.Error())
			}

			if req.LastName == "" {
				return nil, response.BadRequest(ErrLastNameRequired.Error())
			}

			user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email, req.Phone)
			if err != nil {
				return nil, response.InternalServerError(err.Error())
			}

			return response.Created("success", user, nil), nil
		}
	}
*/
func makeLoginAsyncEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginAsyncRequest)

		res, err := s.LoginAsync(ctx, req.Correo)
		if err != nil {
			return response.Error(err.Error(), http.StatusInternalServerError, nil), nil
		}

		return response.Success(0, "", res), nil
	}
}

/*
func makeGetAllEndpoint(s Service, config Config) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetAllReq)

		filters := Filters{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		}

		count, err := s.Count(ctx, filters)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		meta, err := meta.New(req.Page, req.Limit, count, config.LimPageDef)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		users, err := s.GetAll(ctx, filters, meta.Offset(), meta.Limit())
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("success", users, meta), nil
	}
}
func makeGetEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetReq)

		user, err := s.Get(ctx, req.ID)
		if err != nil {

			if errors.As(err, &ErrNotFound{}) {
				return nil, response.NotFound(err.Error())
			}

			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("success", user, nil), nil
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateReq)

		err := s.Update(ctx, req.ID, req.Correo)
		if err != nil {

			if errors.As(err, &ErrNotFound{}) {
				return nil, response.NotFound(err.Error())
			}

			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("success", nil, nil), nil
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(DeleteReq)

		err := s.Delete(ctx, req.ID)

		if err != nil {

			if errors.As(err, &ErrNotFound{}) {
				return nil, response.NotFound(err.Error())
			}
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("success", nil, nil), nil
	}
}
*/
