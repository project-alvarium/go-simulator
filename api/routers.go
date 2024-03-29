/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"github.com/gorilla/mux"
	"github.com/project-alvarium/go-simulator/iota"
	"net/http"
	"strings"
)

// Route holds info of routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array
type Routes []Route

var subStore *iota.SubStore
var readingStore *iota.ReadingStore

// NewRouter creates new router
func NewRouter(subs *iota.SubStore, readings *iota.ReadingStore) *mux.Router {
	subStore = subs
	readingStore = readings

	router := mux.NewRouter().StrictSlash(true)
	//router.Use(mux.CORSMethodMiddleware(router))
	router.Methods("OPTIONS").HandlerFunc(PreFlight)
	router.HandleFunc("/favicon.ico", HandleFavIcon)
	router.HandleFunc("/api/sensors", AddSensor)
	router.HandleFunc("/api/annotators", AddAnnotator)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}


var routes = Routes{

	Route{
		"AddInsight",
		strings.ToUpper("Post"),
		"/api/insights",
		AddInsight,
	},

	Route{
		"GetInsightById",
		strings.ToUpper("Get"),
		"/api/insights/{id}",
		GetInsightByID,
	},
}

func AddSensor(w http.ResponseWriter, r *http.Request) {

	AddNewSensor(w, r, subStore, readingStore)

}

func HandleFavIcon(w http.ResponseWriter, r *http.Request) {}

/*
func SendSubRequest(subscriber iota.Subscriber, data []byte) {
	subscriber.SendSubscriptionIdToAuthor(configuration.AuthConsoleUrl, data)
}
*/

func AddAnnotator(w http.ResponseWriter, r *http.Request) {
	AddNewAnnotator(w, r, subStore, readingStore)
}
