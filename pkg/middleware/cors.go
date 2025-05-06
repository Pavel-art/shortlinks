package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// получаем название домена из хэдера фронта
		origin := r.Header.Get("Origin")
		// если origin ( название домена из хэдера фронта) пустой, пропускаем дальше как обычно
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}
		// Если Origin есть → разрешает доступ этому домену
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", origin)
		header.Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", "POST, GET, PUT,DELETE,PATCH,HEAD,OPTIONS")
			header.Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, content-length")
			header.Set("Access-Control-Max-Age", "86400")

		}
		next.ServeHTTP(w, r)
	})
}

//Браузер видит, что фронтенд (например, https://мой-сайт.ru) пытается запросить данные с бэкенда (например, https://api.site.com).
//
//Без CORS браузер скажет: "Стоп! Ты (мой-сайт.ru) не можешь трогать api.site.com!" и заблокирует запрос.
//
//С CORS бэкенд кричит браузеру:
//"Эй, браузер! Я (api.site.com) разрешаю мой-сайт.ru получать мои данные!"
// header.Set("Access-Control-Allow-Origin", "https://мой-сайт.ru")
// Браузер видит этот заголовок и пропускает запрос
