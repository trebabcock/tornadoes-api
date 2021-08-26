package model

type Tornado struct {
	Id       int     `json:"id"`
	Geometry string  `json:"geometry"`
	Om       int     `json:"om"`
	Yr       int     `json:"yr"`
	Mo       int     `json:"mo"`
	Dy       int     `json:"dy"`
	Date     string  `json:"date"`
	Time     string  `json:"time"`
	Tz       int     `json:"tz"`
	St       string  `json:"st"`
	Stf      int     `json:"stf"`
	Stn      int     `json:"stn"`
	Mag      int     `json:"mag"`
	Inj      int     `json:"inj"`
	Fat      int     `json:"fat"`
	Loss     float32 `json:"loss"`
	Closs    float32 `json:"closs"`
	Slat     float32 `json:"slat"`
	Slon     float32 `json:"slon"`
	Elat     float32 `json:"elat"`
	Elon     float32 `json:"elon"`
	Len      float32 `json:"len"`
	Wid      int     `json:"wid"`
	Fc       int     `json:"fc"`
}

type RequestBody struct {
	YearStart int    `json:"ys"`
	YearEnd   int    `json:"ye"`
	Magnitude []int  `json:"mag"`
	State     string `json:"st"`
}

/*
0|id|INTEGER|0||1
1|geometry|TEXT|0||0
2|om|INTEGER|0||0
3|yr|INTEGER|0||0
4|mo|INTEGER|0||0
5|dy|INTEGER|0||0
6|date|TEXT|0||0
7|time|TEXT|0||0
8|tz|INTEGER|0||0
9|st|TEXT|0||0
10|stf|INTEGER|0||0
11|stn|INTEGER|0||0
12|mag|INTEGER|0||0
13|inj|INTEGER|0||0
14|fat|INTEGER|0||0
15|loss|FLOAT|0||0
16|closs|FLOAT|0||0
17|slat|FLOAT|0||0
18|slon|FLOAT|0||0
19|elat|FLOAT|0||0
20|elon|FLOAT|0||0
21|len|FLOAT|0||0
22|wid|INTEGER|0||0
23|fc|INTEGER|0||0
*/
