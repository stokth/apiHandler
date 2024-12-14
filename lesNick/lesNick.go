package lesNick

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, world!")     // принт в консоль
		d, err := ioutil.ReadAll(r.Body) // читаем тело запроса

		// Проверяем ошибки чтения тела запроса
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusBadRequest)
			return
		}

		// Если все хорошо, печатаем тело запроса в консоль
		fmt.Printf("Received request with body: %s", d)
	})

	http.ListenAndServe("localhost:8080", nil)
}
