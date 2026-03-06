package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Server is starting...")

	http.HandleFunc("/demo", demoHandler)

	err := http.ListenAndServe(":8386", nil)
	if err != nil {
		log.Fatal("Lost connection", err)
		os.Exit(0)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)
	if r.Method != http.MethodGet {
		http.Error(w, "Phuong thuc nay khong duoc ho tro", http.StatusMethodNotAllowed)
	}
	res := map[string]string{
		"hasData": "true",
		"message": "Hello Tuan",
	}
	// Set content type cho response
	w.Header().Set("Content-type", "application/json")
	// marshall convert json ra byte
	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Loi ma hoa JSON", http.StatusInternalServerError)
	}
    // hàm write chỉ nhận []byte nên convert ra byte
	w.Write(data)

}
