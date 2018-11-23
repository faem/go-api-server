package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Profile struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Company  string  `json:"company"`
	Position string  `json:"position"`
	Skill    []Skill `json:"skill"`
}

type Skill struct {
	Name            string `json:"name"`
	NoOfEndorsement int    `json:"noOfEndorsement"`
}

type Admin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

//map for our demo database, used map key as ID
var profilesDB map[string]Profile

//used to get all the profile info using GET request
func GetProfiles(w http.ResponseWriter, r *http.Request) {
	if !BasicAuth() {
		return
	}
	var profiles []Profile

	for _, profile := range profilesDB {
		profiles = append(profiles, profile)
	}
	//shows error if no content in the DB
	if profiles == nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(fmt.Sprintf("No content found!")))
		return
	}

	json.NewEncoder(w).Encode(profiles)
	//w.Write([]byte(fmt.Sprintf("%v", profiles)))
}

//used to get a specific profile info using GET request
func GetProfile(w http.ResponseWriter, r *http.Request) {
	if !BasicAuth() {
		return
	}

	params := mux.Vars(r)

	if _, flag := profilesDB[params["id"]]; !flag {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Profile of " + string(params["id"]) + " not found!")))
		return
	}

	json.NewEncoder(w).Encode(profilesDB[params["id"]])
}

//used to delete a profile using DELETE request
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	if !BasicAuth() {
		return
	}

	params := mux.Vars(r)
	if _, flag := profilesDB[params["id"]]; !flag {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Profile of " + string(params["id"]) + " not found!")))
		return
	}

	delete(profilesDB, params["id"])
}

//used to add new profile using POST request
func AddProfile(w http.ResponseWriter, r *http.Request) {
	if !BasicAuth() {
		return
	}

	var profile Profile
	json.NewDecoder(r.Body).Decode(&profile)

	if _, flag := profilesDB[profile.ID]; !flag {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(fmt.Sprintf("Username " + string(profile.ID) + " already exists!")))
		return
	}

	profilesDB[profile.ID] = profile
}

//used to update a profile using PUT request
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if !BasicAuth() {
		return
	}

	params := mux.Vars(r)

	if _, flag := profilesDB[params["id"]]; !flag {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Profile of " + string(params["id"]) + " not found!")))
		return
	}

	var profile Profile
	json.NewDecoder(r.Body).Decode(&profile)
	profilesDB[params["id"]] = profile
}

//basic authentication function
func BasicAuth() bool {
	return true
}

//this function creates a demo DB for our server
func CreateDemoDB() {
	profilesDB = make(map[string]Profile)
	profilesDB["fahim-abrar"] = Profile{
		"fahim-abrar",
		"Mohammad Fahim Abrar",
		"AppsCode Inc.",
		"Software Engineer",
		[]Skill{
			{
				"C++",
				3,
			},
			{
				"Android",
				4,
			},
		}}

	profilesDB["masud-rahman"] = Profile{
		"masud-rahman",
		"Masudur Rahman",
		"AppsCode Inc.",
		"Software Engineer",
		[]Skill{
			{
				"C",
				3,
			},
			{
				"C++",
				4,
			},
		}}

	profilesDB["mohan"] = Profile{
		"mohan",
		"Tahsin Rahman",
		"AppsCode Inc.",
		"Software Engineer",
		[]Skill{
			{
				"C",
				100,
			},
			{
				"C++",
				110,
			},
			{
				"Linux",
				100,
			},
		}}
}

func main() {
	router := mux.NewRouter()
	CreateDemoDB()

	router.HandleFunc("/in", GetProfiles).Methods("GET")
	router.HandleFunc("/in/{id}", GetProfile).Methods("GET")
	router.HandleFunc("/in/{id}", UpdateProfile).Methods("PUT")
	router.HandleFunc("/in", AddProfile).Methods("POST")
	router.HandleFunc("/in/{id}", DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
