package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myusername"
	password = "mypassword"
	dbname   = "postgres"
)

type ThinkTank struct {
	gorm.Model
	ProjectDescription string   `json:"projectDescription"`
	ProjectName        string   `json:"projectName"`
	Keywords           []string `json:"Keywords"`
	RelatedLinks       []string `json:"RelatedLinks"`
}

var ThinkTanks = []ThinkTank{

	{
		ProjectDescription: "11gbhghgnbn111",
		ProjectName:        "333gbgbbnb3",
		Keywords:           []string{"ttgngtbghtt"},
		RelatedLinks:       []string{"444gngbgbn4444"},
	},
}

func CreateThinkTank(w http.ResponseWriter, r *http.Request) {
	//  READ MY  BODY
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var thinkTank ThinkTank
	json.Unmarshal(body, &thinkTank)

	ThinkTanks = append(ThinkTanks, thinkTank)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content-Type", "applicatio/json")
	json.NewEncoder(w).Encode("Created")

}

func GetThinkTank(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "applicatio/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ThinkTanks)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            //  *  means select all                                                      // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func main() {
	fmt.Println("HERE WEBPAGE")
	// router := mux.New()

	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/thinkTanks", GetThinkTank).Methods("GET")
	router.HandleFunc("/thinkTank", CreateThinkTank).Methods("POST")
	// json.NewEncoder(w).Encode("Hello world")

	log.Println("API  IS RUNNING")
	http.ListenAndServe(":3000", router)

	//CONNECTION OF DATABASE

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected!")
}
