package models

/*{
"name": "Attend Artificial Inteligence Laboratory",
"description": "",
"project": "University",
"id": "693cc8a2-75f0-4c09-9b42-c5a3f6a2c9ab",
"date": "1/9/2019",
"time": "13:00 PM",
"priority": "3"
},*/

type Todo struct{
	Id          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Project     string `json:"project"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Priority    int `json:"priority"`
}
