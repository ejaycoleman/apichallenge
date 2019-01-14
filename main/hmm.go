
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Reponse should be an array of the structure Dwarve, stored in the json as 'dwarves'
type Response struct {
	Dwarves []Dwarve `json:"dwarves"`
}

// One dwarve should have one attribute: Name - stored in the json as 'name'
type Dwarve struct {
	Name string `json:"name"`
}

// Our second type of request should consist of an array of DwarveFulls, again stored in the json as 'dwarves'
type ResponseTwo struct {
	Dwarves []DwarveFull `json:"dwarves"`
}

// Here we can pull the parameters from the json: name, birth, death, and culture
type DwarveFull struct {
	Name string `json:"name"`
	Birth string `json:"birth"`
	Death string `json:"death"`
	Culture string `json:"culture"`
}

func main() {
	// GET request to the source URL, store in variable response
	response, err := http.Get("https://thedwarves.pusherplatform.io/api/dwarves")

	// print any errors that occured during the HTTP GET request and exit
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	// store the body into the variable responseData
	responseData, err := ioutil.ReadAll(response.Body)

	// again, print any errors
	if err != nil {
		log.Fatal(err)
	}

	// initialse responseObject with the stucture of Response
	var responseObject Response	

	// I believe this puts the JSON into the object based on the structure of Response
	json.Unmarshal(responseData, &responseObject)

	// first API endpoint: GET /api/dwarves
  http.HandleFunc("/api/dwarves", func(w http.ResponseWriter, r *http.Request) {
  	// We can now put the object we created back into a JSON form, which we can display
  	// The Unmarshal and Marshal steps are required to filter out biirth, death and culture
  	b, err := json.Marshal(responseObject)
	  if err != nil {
	      fmt.Println(err)
	      return
	  }

	  // Display the JSON in the GET request body
		fmt.Fprintf(w, "%s", b)
  })

  // our second type of request, using the form ResponseTwo
  var responseObjectFull ResponseTwo
  // We populate the responseObjectFull
	json.Unmarshal(responseData, &responseObjectFull)

	// second API endpoint: /api/dwarves/
	http.HandleFunc("/api/dwarves/", func(w http.ResponseWriter, r *http.Request) {
		// loop through all the Dwarves in the responseObjectFull
		for i := 0; i < len(responseObjectFull.Dwarves); i++ {
			// get the name of the dwarve. Remove first 13 characters because '/api/dwarves/'
			// I didnt know how to get the parameter only
			// if a dwarve with this name exists...
			if responseObjectFull.Dwarves[i].Name == r.URL.Path[13:] {
				// we put the Dwarveback into a JSON form
				b, err := json.Marshal(responseObjectFull.Dwarves[i])
		    if err != nil {
		        fmt.Println(err)
		        return
		    }
		    // and display it in the GET request body
				fmt.Fprintf(w, "%s" , b)
			}
		}

  })


	// Start the server on port 8080
  log.Fatal(http.ListenAndServe(":8080", nil))

}