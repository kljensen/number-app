package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func numberHandler(w http.ResponseWriter, r *http.Request) {
	myRandomNumber := rand.Intn(10)
	numbers, ok := r.URL.Query()["userNumber"]
	userNumber := "0"
	if ok {
		userNumber = numbers[0]
	}
	myHTML := `
        <html>
            <head>
                <style>
                    body {
                        background-color: #C2DFFF;
                        color: red;
                        font-size: 2rem;
                    }
                    img {
                        width: 50%%;
                    }
                </style>
            </head>
            <body>
                <h1>Welcome to the guessing game!</h1>
                <p>Your guess was %s</p>
                <p>The computer selected %d</p>
                 <form>
                    <input name="userNumber" type="number" min="0" max="10" />
                    <input type="submit" />
                 </form>
                 <div>
                 <img src="https://cataas.com/cat/gif">
                 </div>
            </body>
        </html>
    `
	fmt.Fprintf(w, myHTML, userNumber, myRandomNumber)
}

func main() {
	http.HandleFunc("/", numberHandler)
	http.ListenAndServe(":8080", nil)
}
