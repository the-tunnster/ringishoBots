package handlers

import (
	"botStuff/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"googlemaps.github.io/maps"
)

const KEY string = ""
const AlgoliaAdminKey string = "5d4aa09d1df84fb62ff793d7fcb1b16b"
const AlgoliaApplicationID string = "2Z19YUPYM0"

func BotSupervisor(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit BotSupervisorPage Post Endpoint.")

	var req models.Question
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := models.Answer{
		UserID:     req.UserID,
		QuestionID: req.QuestionID,
	}

	json.NewEncoder(w).Encode(res)

}

func TravelBot(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit TravelBotPage Post Endpoint.")

	var req models.Question
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	location := req.BotParameters

	googleClient, err := maps.NewClient(maps.WithAPIKey(KEY))
	if err != nil {
		log.Fatal("Error creating maps client")
	}

	geocodeReq := &maps.GeocodingRequest{
		Address: location,
	}

	geocodeRes, err := googleClient.Geocode(context.Background(), geocodeReq)
	if err != nil {
		log.Fatal("Wack.")
	}

	nearbyReq := &maps.NearbySearchRequest{
		Location: &geocodeRes[0].Geometry.Location,
		Radius:   10000,
		Type:     "tourist_attraction",
	}

	nearbyRes, err := googleClient.NearbySearch(context.Background(), nearbyReq)
	if err != nil {
		log.Println("ded")
	}

	var results []string

	for _, v := range nearbyRes.Results {
		results = append(results, v.Name)
	}

	json.NewEncoder(w).Encode(results)
}

func AlgoliaBot(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit AlgoliaBotPage Post Endpoint.")

	client := search.NewClient(AlgoliaApplicationID, AlgoliaAdminKey)
	index := client.InitIndex("Dev_5TTC")

	params := []interface{}{
		opt.HitsPerPage(10),
	}

	res, err := index.Search("Education", params...)

	var results []interface{}

	if err != nil {
		log.Println(err)
	} else {
		for _, v := range res.Hits {
			results = append(results, v["Title"])
		}
	}

	json.NewEncoder(w).Encode(results)
}
