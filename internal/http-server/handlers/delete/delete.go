package delete 

import ("errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage")

type URLDelete interface{
	DeleteURL(alias string) error 
}
func Delete(log *slog.Logger, urlGetter URLDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
      const op = "handlers.url.delete.Delete"
	  log := log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	alias := chi.URLParam(r, "alias")
	if alias == "" {
		log.Info("alias is empty")

		render.JSON(w, r, resp.Error("invalid request"))

		return
	}

	 err := urlGetter.DeleteURL(alias)
	if errors.Is(err, storage.ErrURLNotFound) {
		log.Info("url not found", "alias", alias)

		render.JSON(w, r, resp.Error("not found"))

		return
	}
	if err != nil {
		log.Error("failed to delete url", sl.Err(err))

		render.JSON(w, r, resp.Error("internal error"))

		return
	}

	log.Info("delete url")

	 w.WriteHeader(http.StatusOK)
	}}