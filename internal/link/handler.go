package link

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"shortlinks/configs"
	"shortlinks/pkg/middleware"
	"shortlinks/pkg/req"
	"shortlinks/pkg/res"
	"strconv"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
}
type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())

}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получаем тело запроса от пользователя
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		// проверяем не совпадает ли hash с существующим в базе
		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		// записываем в базу новый объект
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// возвращаем ответ пользователю
		res.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}
		// получаем тело запроса от пользователя
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		//получаем динамический параметр из запроса
		idString := r.PathValue("id")
		// приводим строку к типу uint
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// обновляем данные в базе данных
		link, err := handler.LinkRepository.Update(
			&Link{
				Model: gorm.Model{ID: uint(id)},
				Url:   body.Url,
				Hash:  body.Hash,
			})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// возвращаем ответ пользователю
		res.Json(w, link, http.StatusOK)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//получаем динамический параметр из запроса
		idString := r.PathValue("id")
		// приводим строку к типу uint
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// проверям по id существует ли запись в базе и сели нет, возвращаем статус ошибки пользователю.
		_, err = handler.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// удаляем запись из базы
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// возвращаем ответ пользователю
		res.Json(w, nil, http.StatusOK)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//получаем динамический параметр из запроса
		hash := r.PathValue("hash")
		// находим ссылку по hash
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// перенапрявляем пользователя по ссылке на сайт
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
