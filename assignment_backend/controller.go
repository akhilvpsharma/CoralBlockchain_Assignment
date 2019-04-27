package main
import
(
	//"fmt"
	"log"
	"strings"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func SaveController(w http.ResponseWriter, r *http.Request){
		// Setting return type
		w.Header().Set("Content-Type", "application/json")

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
		var org WebForm
		err = json.Unmarshal(reqBody, &org)
		//fmt.Println(string(reqBody))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
		err=nil
		success, message:=SaveService(org)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
	
		// Status code
		w.WriteHeader(http.StatusOK)
		
		// Body
		data := simplejson.New()
		data.Set("success", success)
		data.Set("message", message)
		payload, err := data.MarshalJSON()
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(getErrorPayload(err))
			return
		}
		// Return response
		w.Write(payload)
		return
}

func SearchController(w http.ResponseWriter, r *http.Request) {

	// Setting return type
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	emailId := vars["emailId"]
	
	webform, success, message, err:=SearchService(emailId)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}

	// Status code
	w.WriteHeader(http.StatusOK)
	
	// Body
	data := simplejson.New()
	if(success=="true"){
		data.Set("Email Id", webform.EmailID)
		data.Set("User Name", webform.UserName)
		data.Set("Phone Number", webform.PhoneNumber)
		data.Set("Date Time", webform.DateTime)
	}else{
		data.Set("success", success)
		data.Set("message", message)	
	}
	payload, err := data.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}
	// Return response
	w.Write(payload)
	return
}
func DeleteController(w http.ResponseWriter, r *http.Request) {

	// Setting return type
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	emailId := vars["emailId"]
	
	success, message, err:=DeleteService(emailId)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}

	// Status code
	w.WriteHeader(http.StatusOK)
	
	// Body
	data := simplejson.New()
	data.Set("success", success)
	data.Set("message", message)	
	payload, err := data.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorPayload(err))
		return
	}
	// Return response
	w.Write(payload)
	return
}

func getErrorPayload(err error) []byte {
	var data = simplejson.New()
	errorParts := strings.Split(string(err.Error()), "Description: ")

	// Body
	data.Set("success", false)
	data.Set("error", errorParts[len(errorParts)-1])

	payload, err := data.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	return payload
}
