package model

import (

)


type Inspection struct {
    Jadwal Jadwal
    Inspector []Inspector
    Condition []Condition

}

type Jadwal struct {
        Area      int    	`json:"area"`
        Hari      string 	`json:"hari"`
        Tanggal   string 	`json:"tanggal"`
        Waktu     string 	`json:"waktu"`
        Judul     string 	`json:"judul"`
        Lokasi    string 	`json:"lokasi"`
        SubLokasi string 	`json:"sub_lokasi"`
        Detail    string 	`json:"detail"`
    } 

type Inspector struct {
	Name	string	`json:"name"`
}
type InspectorArray struct {
	Name	[]string	
}

type Condition struct {
    Chart 	int    	`json:"chart"`
    Action 	string  	`json:"action"`
    Date   	string  	`json:"date"`
    Image  	string  	`json:"image"`
}
type ConditionArray struct {
	Chart	[]int
	Action	string
	Date	string
	Image	[]string
}

