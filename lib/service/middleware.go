package service

import "net/http"

type Middleware func(serviceHandler) serviceHandler

func (s *Service) logMiddleware(next serviceHandler) serviceHandler {

	return func(w http.ResponseWriter, r *http.Request) error {
		s.Logger.Info("Logging before")
		if err := next(w, r); err != nil {
			return err
		}
		s.Logger.Info("Looging after")
		return nil
	}
}
