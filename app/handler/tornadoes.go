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
	mss := r.URL.Query()["ms"][0]
	dss := r.URL.Query()["ds"][0]
	yes := r.URL.Query()["ye"][0]
	mes := r.URL.Query()["me"][0]
	des := r.URL.Query()["de"][0]
	mags := r.URL.Query()["mag"][0]
	st := r.URL.Query()["st"][0]

	ys, _ := strconv.Atoi(yss)
	ms, _ := strconv.Atoi(mss)
	ds, _ := strconv.Atoi(dss)
	ye, _ := strconv.Atoi(yes)
	me, _ := strconv.Atoi(mes)
	de, _ := strconv.Atoi(des)
	magss := strings.Split(mags, ",")

	mag := []int{}

	for _, m := range magss {
		mi, _ := strconv.Atoi(m)
		mag = append(mag, mi)
	}

	tornadoes := []model.Tornado{}
	if st == "ALL" {
		db.Where("mag IN ?", mag).Where("st = ?", st).Where(db.Where("yr >= ?", ys).Where("mo >= ?", ms)).Where(db.Where("yr <= ?", ye).Where("mo <= ?", me)).Find(&tornadoes)
	} else {
		db.Where("mag IN ?", mag).Where("st = ?", st).Where(db.Where("yr >= ?", ys).Where("mo >= ?", ms)).Where(db.Where("yr <= ?", ye).Where("mo <= ?", me)).Find(&tornadoes)

	}

	responseTornadoes := []model.ResponseTornado{}

	for _, t := range tornadoes {
		tor := model.ResponseTornado{
			Id:       t.Id,
			Geometry: t.Geometry,
			Yr:       t.Yr,
			Mo:       t.Mo,
			Dy:       t.Dy,
			St:       t.St,
			Mag:      t.Mag,
			Inj:      t.Inj,
			Fat:      t.Fat,
		}
		if tor.Mo == ms && tor.Dy >= ds {
			responseTornadoes = append(responseTornadoes, tor)
		} else if tor.Mo == me && tor.Dy <= de {
			responseTornadoes = append(responseTornadoes, tor)
		}
	}

	RespondJSON(w, http.StatusOK, responseTornadoes)
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

	responseTornadoes := []model.ResponseTornado{}

	for _, t := range tornadoes {
		tor := model.ResponseTornado{
			Id:       t.Id,
			Geometry: t.Geometry,
			Yr:       t.Yr,
			Mo:       t.Mo,
			Dy:       t.Dy,
			St:       t.St,
			Mag:      t.Mag,
			Inj:      t.Inj,
			Fat:      t.Fat,
		}
		responseTornadoes = append(responseTornadoes, tor)
	}

	RespondJSON(w, http.StatusOK, responseTornadoes)
}
