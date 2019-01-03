
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Dwarves []Dwarve `json:"dwarves"`
}

type Dwarve struct {
	Name string `json:"name"`
}

// func main() {
	

// }


func main() {
	response, err := http.Get("https://thedwarves.pusherplatform.io/api/dwarves")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var responseObject Response	
	// json.Unmarshal([]byte(responseData), &dwarves)
	json.Unmarshal(responseData, &responseObject)
// fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)



	// fmt.Println(string(dwarves.Name))

	for i := 0; i < len(responseObject.Dwarves); i++ {
		fmt.Println(responseObject.Dwarves[i].Name)
	}

		b, err := json.Marshal(responseObject)
    if err != nil {
        fmt.Println(err)
        return
    }

    http.HandleFunc("/api/dwarves", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s", b)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}