package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"leckomio.dev/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/EUR" // weird or no? We have a const but we set some value dynamically... WTF???

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters are required: %d received", len(currency))
	}

	upperCaseCurrency := strings.ToUpper(currency) // strings is a package and a string variable has no methods because it is no struct (remember: do not say object because there are not objects in go!), so you can not do: mystring.ToUpper
	// ToUpper is a fuction here ^

	// Get is not a method. It is a function because
	// http is not a struct but a package
	// and .Get is a function that is beeing exported (aka Capital-Letter)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCaseCurrency)) // notice that this is synchronous code... see *
	// Sprintf is returning a string and not printing to the poutput

	if err != nil {
		// to return nil for the Rate I have to return a pointer to the Rate not the Rate itself
		// see response type: (*datatypes.Rate, error)
		return nil, err
	}

	// Candy: Status codes are evolving over time
	// There is a new status code 103 - Early hints
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/103
	if res.StatusCode == http.StatusTeapot {
		fmt.Println("It is teatime!")
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != err {
			return nil, err
		}
		// json := string(bodyBytes)
		err2 := json.Unmarshal(bodyBytes, &response)
		if err2 != nil {
			fmt.Println(err2)
			return nil, err2
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}

	// structs have constructors by default
	rate := datatypes.Rate{Currency: currency, Price: response.Ask}
	return &rate, nil
}

// * In go we typically have not async calls
// in go we use threads aka goroutines for this, muahahahaha 😍
