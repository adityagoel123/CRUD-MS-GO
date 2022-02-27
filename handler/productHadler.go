package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/adityagoel/product-api/data"
)

type Products struct {
	thisLogger *log.Logger
}

func NewProducts(thisLogger *log.Logger) *Products {
	return &Products{thisLogger}
}

func (h *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		h.getProducts(responseWriter, request)
		return
	}

	if request.Method == http.MethodPost {
		h.addProduct(responseWriter, request)
		return
	}

	if request.Method == http.MethodPut {
		h.thisLogger.Println("The request.URL.Path received in Request is :", request.URL.Path)
		regexReceived := regexp.MustCompile(`/([0-9]+)`)

		g := regexReceived.FindAllStringSubmatch(request.URL.Path, -1)

		h.thisLogger.Println("The group received in Request, Single Param is :", g[0])
		h.thisLogger.Println("The group received in Request, Single's ONE Param is :", g[0][1])

		if len(g) != 1 {
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(responseWriter, "Invalid CaptureGroup", http.StatusBadRequest)
			return
		}

		idString := g[0][1]

		h.thisLogger.Println("The idString received in Request is :", idString)

		id, errWhileConvertingStrToInt := strconv.Atoi(idString)

		if errWhileConvertingStrToInt != nil {
			http.Error(responseWriter, "Error while converting String to Integer.", http.StatusBadRequest)
			return
		}

		h.thisLogger.Println("The id received in Request is :", id)

		// Updating this prooduct to to our Temporary DataStore Now.
		h.updateSingleProduct(id, responseWriter, request)
		return
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Products) updateSingleProduct(id int, responseWriter http.ResponseWriter, request *http.Request) {

	h.thisLogger.Println("Handle PUT Products")

	thisProd := &data.Product{}

	errorWhileUnMarshalling := thisProd.FromJson(request.Body)

	if errorWhileUnMarshalling != nil {
		http.Error(responseWriter, "Unable to unmarshall the JSON", http.StatusBadRequest)
	}

	// Adding this prooduct to tour Temporary DataStore Now.
	err := data.UpdateSingleProduct(id, thisProd)

	if err != nil {
		http.Error(responseWriter, "Unable to update this product to DataStore", http.StatusBadRequest)
	}

	h.thisLogger.Printf("Prod: %#v", thisProd)
}

func (h *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {

	h.thisLogger.Println("Handle Get Products")

	listOfProducts := data.GetProducts()
	errorWhileEncoding := listOfProducts.ToJson(responseWriter)

	if errorWhileEncoding != nil {
		http.Error(responseWriter, "Unable to encode the JSON", http.StatusInternalServerError)
	}
}

func (h *Products) addProduct(responseWriter http.ResponseWriter, request *http.Request) {

	h.thisLogger.Println("Handle Post Product")

	thisProd := &data.Product{}

	errorWhileUnMarshalling := thisProd.FromJson(request.Body)

	if errorWhileUnMarshalling != nil {
		http.Error(responseWriter, "Unable to unmarshall the JSON", http.StatusBadRequest)
	}

	// Adding this prooduct to tour Temporary DataStore Now.
	data.AddSingleProduct(thisProd)

	h.thisLogger.Printf("Prod: %#v", thisProd)
}
