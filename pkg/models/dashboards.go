package models

import (
	"encoding/json"
	"io"
	"regexp"
	"strings"
	"time"
)

type Dashboard struct {
	Id                   string `gorethink:"id,omitempty"`
	Slug                 string
	AccountId            int
	LastModifiedByUserId string
	LastModifiedByDate   time.Time
	CreatedDate          time.Time

	Title string
	Tags  []string
	Data  map[string]interface{}
}

type CollaboratorLink struct {
	AccountId  int
	Role       string
	ModifiedOn time.Time
	CreatedOn  time.Time
}

type UserAccount struct {
	Id              int `gorethink:"id"`
	UserName        string
	Login           string
	Email           string
	Password        string
	NextDashboardId int
	UsingAccountId  int
	Collaborators   []CollaboratorLink
	CreatedOn       time.Time
	ModifiedOn      time.Time
}

type UserContext struct {
	UserId    string
	AccountId string
}

type SearchResult struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

func NewDashboard(title string) *Dashboard {
	dash := &Dashboard{}
	dash.Id = ""
	dash.LastModifiedByDate = time.Now()
	dash.CreatedDate = time.Now()
	dash.LastModifiedByUserId = "123"
	dash.Data = make(map[string]interface{})
	dash.Data["title"] = title
	dash.Title = title
	dash.UpdateSlug()

	return dash
}

func NewFromJson(reader io.Reader) (*Dashboard, error) {
	dash := NewDashboard("temp")
	jsonParser := json.NewDecoder(reader)

	if err := jsonParser.Decode(&dash.Data); err != nil {
		return nil, err
	}

	return dash, nil
}

func (dash *Dashboard) GetString(prop string) string {
	return dash.Data[prop].(string)
}

func (dash *Dashboard) UpdateSlug() {
	title := strings.ToLower(dash.Data["title"].(string))
	re := regexp.MustCompile("[^\\w ]+")
	re2 := regexp.MustCompile("\\s")
	dash.Slug = re2.ReplaceAllString(re.ReplaceAllString(title, ""), "-")
}

func (account *UserAccount) AddCollaborator(accountId int) {
	account.Collaborators = append(account.Collaborators, CollaboratorLink{
		AccountId:  accountId,
		Role:       "admin",
		CreatedOn:  time.Now(),
		ModifiedOn: time.Now(),
	})
}
