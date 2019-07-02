package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/xwb1989/sqlparser"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/task1", Task1Handler)
	r.HandleFunc("/task2", Task2Handler)
	r.HandleFunc("/task3", Task3Handler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
}
func Task1Handler(w http.ResponseWriter, r *http.Request) {
	emailList := returnEmails()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emailList)

}
func Task2Handler(w http.ResponseWriter, r *http.Request) {

}
func Task3Handler(w http.ResponseWriter, r *http.Request) {
	var s Result
	keys, ok := r.URL.Query()["query"]
	if ok {
		stmt, _ := sqlparser.Parse(keys[0])
		var tables []string
		tables = getTableNames(reflect.Indirect(reflect.ValueOf(stmt)), tables, 0, false, keys[0])
		s.Status = true
		s.Text = tables
	} else {
		s.Status = false
		s.Text[0] = "No Param detected"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

type Result struct {
	Status bool     `json:"status"`
	Text   []string `json:"text"`
}
type Emails struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	About   string `json:"about"`
}

func returnEmails() []Emails {
	var emails []Emails
	var email Emails
	email.About = "First Mail after conversation"
	email.Subject = "Thank you for the conversation"
	email.Body = "Hey Umesh,\nIt was amazing speaking with you today. There were some great insights about your vision for American Express and some striking similarities between our product and your expectations.\n\n\nJust wanted to quickly summarise my understandings from today's conversation : \n1. To increase the growth of AmEx, you are looking to make a data lake, but the challenge here is that the data is distributed amongst various sources.\n2. In the previous financial year, Amex integrated the data sources generating employee's data\n3. In the current financial year, Amex is planning to integrate the customer's data\n4. Both 2 and 3 take time to structure the data as the sources are differently structured which makes it difficult to maintain the quality of the data\n5. There are instances where the data collected in incomplete and sometimes various versions are found\n6. There is a tool that you have developed 5 years back but it needs an update to cater all of your data requirements\n\nOn the basis of my understanding, I would like to pitch to you the product which I introduced to you during our conversation.\n\n1. Workflows: This drag and drop interface enables you to merge data from various sources and will migrate the resultant set into the said destination. All you need to do is set up the sources on our panel and connect the sources through the UI tool. \nIn your use case, integrating the employee's data as well as the customer's data and forming a data lake would be a work of a week and you can start populating intelligence from the integrated data from day 1\n\n2. Discovery: The ultimate tool to democratize data is something that you will surely need in times of updating the access management as well as to monitor any changes or issues that are raised by any of the stakeholders.\nI understand that a lot of time and effort has been spent by your team already to integrate your data, but I ensure that once you have our back, you can totally focus on how you can channelize the key factors that can drive your growth.\nIt will be great if I can meet your data team along with you so that I can get feedback as to how do they find our product. Also, would love to know in detail about the entire data flow requirements so that we can plan the best solution of your data management.\n\n\nRegards\nNikhil"
	emails = append(emails, email)
	email.About = "First follow-up"
	email.Subject = "Atlan wishes you a very happy Pongal"
	email.Body = "Hey Umesh,\n\nI hope all is fine at your end. A very happy Pongal to you.\nOur product has upgraded this week and now it has even more features to fulfil your data requirements.\n\nThe key additions to our product are:\n1. Multiple vendors integration\n2. Precise monitoring\n3. Log access for your data engineering team\n4. Slack notifications to keep you posted on the go\n\nOur product team has worked exceptionally well to get all of these features in such a small time. I am really excited to demonstrate the upgraded Atlan to you.\n\nPlease let me know a good time for the next set of conversations so that we can take this forward.\n\nRegards\nNikhil"
	emails = append(emails, email)
	email.About = "Second follow-up"
	email.Subject = "Follow up from Atlan"
	email.Body = "Hey Umesh,\n\nI hope all is fine at your end. \n\nJust wanted to quickly follow up with you if you and your team have any free slots in the upcoming weeks so that we can schedule a demo of Atlan.\n\nRegards\nNikhil"
	emails = append(emails, email)
	email.About = "Third follow-up"
	email.Subject = "Atlan updates"
	email.Body = "Hey Umesh,\n\nI hope all is fine at your end.\n\nIt gives me immense please to share with you that Atlan is currently been used by most of the Banks in South-East Asia and they are very convinced with the results.\n\nSome of the key impacts post-Atlan integration on these banks are:\n1. Loan repayment success ration has enhanced by 5%\n2. Appropriate IVR messages formulated based on the customers' data, thus, less turnaround time to resolve issues\n3. Quick insights on the employees' history helped in attrition\n\nYou can find some of the video testimonials here.\n\nWould love to have your team a look at the product that we have developed so that AmEx gets even better results and achieve what you are looking for.\n\nPlease let me know a good time when I can visit your office and demonstrate the product to you and your team.\n\nRegards\nNikhil"
	emails = append(emails, email)
	email.About = "Fourth follow-up"
	email.Subject = "Atlan Wishes you very happy Vesak day"
	email.Body = "Hey Umesh,\n\nI hope all is fine at your end. A very happy belated Vesak day to you.\n\nIt seems that organizing and democratizing data is a very exhausting activity and it becomes more painful when there are multiple updates and upgrades in the schemas where the distributed data is stored.\n\nMany of our clients have mentioned that they used to face the same situation in their pre-Atlan era.\n\nWhat if I tell you that in Atlan's latest release, you can set up scripts and schedule syncing of data. These sync scripts also have a monitor which notifies you whenever there is anything wrong with any of your data instances and makes sure that your data lake is safe and atomic.\n\nWould love to have a long term business relationship with your prestigious financial institution and thus, requesting to provide a slot with you and your team where I can demonstrate this amazing product.\n\n\nRegards\nNikhil"
	emails = append(emails, email)
	email.About = "Fifth follow-up"
	email.Subject = "Atlan wishes you happy Deepavali"
	email.Body = "Hey Umesh,\n\nI hope all is fine at your end. A very happy Deepavali to you.\n\nMay this day bring a lot of joy to you and your family.\n\nI am travelling to Singapore for a fin-tech conference in the next few days so I was wondering if you have any free slots available with your team so that we can have a demo of our tool, Atlan.\n\nHoping to hear back from you.\n\n\nRegards\nNikhil"
	emails = append(emails, email)
	return emails
}
func getTableNames(v reflect.Value, tables []string, level int, isTable bool, statement string) []string {
	switch v.Kind() {
	case reflect.Struct:
		if v.Type().Name() == "TableIdent" {
			// if this is a TableIdent struct, extract the table name
			tableName := v.FieldByName("v").String()
			if tableName != "" && isTable {
				tables = AppendIfMissing(tables, tableName)
			}
		} else {
			// otherwise enumerate all fields of the struct and process further
			for i := 0; i < v.NumField(); i++ {
				tables = getTableNames(reflect.Indirect(v.Field(i)), tables, level+1, isTable, "")
			}
		}
	case reflect.Array, reflect.Slice:

		for i := 0; i < v.Len(); i++ {
			// enumerate all elements of an array/slice and process further
			tables = getTableNames(reflect.Indirect(v.Index(i)), tables, level+1, isTable, "")
		}
	case reflect.Interface:
		if v.Type().Name() == "SimpleTableExpr" {
			isTable = true
		}
		// get the actual object that satisfies an interface and process further
		tables = getTableNames(reflect.Indirect(reflect.ValueOf(v.Interface())), tables, level+1, isTable, "")

	}

	return tables
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
