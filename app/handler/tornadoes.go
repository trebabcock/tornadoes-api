package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"tornadoes/app/model"

	"gorm.io/gorm"
)

// GetTornadoesByRange gets all tornadoes of a specified rating, state, and year range
func GetTornadoesByRange(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	yss := r.URL.Query()["ys"][0]
	yes := r.URL.Query()["ye"][0]
	mags := r.URL.Query()["mag"][0]
	st := r.URL.Query()["st"][0]

	ys, _ := strconv.Atoi(yss)
	ye, _ := strconv.Atoi(yes)
	magss := strings.Split(mags, ",")

	mag := []int{}

	for _, m := range magss {
		mi, _ := strconv.Atoi(m)
		mag = append(mag, mi)
	}

	tornadoes := []model.Tornado{}
	// db.Where("pizza = ?", "pepperoni").Where
	//db.Where(model.Tornado{Mag: magInt}).Find(&tornadoes)
	if st == "ALL" {
		db.Where("yr >= ?", ys).Where("yr <= ?", ye).Where("mag IN ?", mag).Find(&tornadoes)
	} else {
		db.Where("yr >= ?", ys).Where("yr <= ?", ye).Where("mag IN ?", mag).Where("st = ?", st).Find(&tornadoes)
	}

	RespondJSON(w, http.StatusOK, tornadoes)
}

// GetTornadoesByDate gets all tornadoes of a specified rating, state, and date
func GetTornadoesByDate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	years := r.URL.Query()["year"][0]
	months := r.URL.Query()["month"][0]
	days := r.URL.Query()["day"][0]
	mags := r.URL.Query()["mag"][0]
	st := r.URL.Query()["st"][0]

	year, _ := strconv.Atoi(years)
	month, _ := strconv.Atoi(months)
	day, _ := strconv.Atoi(days)

	magss := strings.Split(mags, ",")

	mag := []int{}

	for _, m := range magss {
		mi, _ := strconv.Atoi(m)
		mag = append(mag, mi)
	}

	tornadoes := []model.Tornado{}
	// db.Where("pizza = ?", "pepperoni").Where
	//db.Where(model.Tornado{Mag: magInt}).Find(&tornadoes)
	if st == "ALL" {
		db.Where("yr = ?", year).Where("mo = ?", month).Where("dy = ?", day).Where("mag IN ?", mag).Find(&tornadoes)
	} else {
		db.Where("yr = ?", year).Where("mo = ?", month).Where("dy = ?", day).Where("mag IN ?", mag).Where("st = ?", st).Find(&tornadoes)
	}

	RespondJSON(w, http.StatusOK, tornadoes)
}
