package models



type Studentth struct {
	Id        int    `json:"size:32;column:id;auto_increment"`
	Username  string `json:"size:512;column:username"`
	Companyid string `json:"size:512;column:companyid"`
	Date      string `json:"size:512;column:date"`
	Text      string `json:"size:1024;column:text"`
	Messageid      int `json:"size:64;column:messageid"`
	Credit Credit


}

type Credit struct {

	Credit  string `json:"size:512;column:credit"`
	Cred string `json:"size:512;column:cred"`




}


