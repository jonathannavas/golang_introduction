package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	User User `json:"data"`
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserBody struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Job       string     `json:"job"`
	Age       int64      `json:"age,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

const (
	base = "https://reqres.in/api"
)

func main() {
	p := fmt.Println
	p("Package net")

	req, err := GetReqExample(fmt.Sprintf("%s/users/2", base))
	if err != nil {
		log.Fatal(err)
	}
	p(string(req))

	p()

	req, err = Get(fmt.Sprintf("%s/users/2", base))
	if err != nil {
		log.Fatal(err)
	}
	p(string(req))

	p()
	res, err := GetUser("2")
	if err != nil {
		log.Fatal(err)
	}

	p(res)
	p("Id: ", res.ID)
	p("email: ", res.Email)
	p("Firstname: ", res.FirstName)
	p("Lastname: ", res.LastName)

	p()
	usr, err := CreateUser("Jonathan", "Software Engineer")
	if err != nil {
		log.Fatal(err)
	}
	p(usr)
	p("id:", usr.ID)
	p("name:", usr.Name)
	p("job:", usr.Job)
	p("createdAt:", usr.CreatedAt)
	p("age:", usr.Age)
	p("deletedAt:", usr.DeletedAt)

	userConverted, err := json.Marshal(usr)
	if err != nil {
		log.Fatal(err)
	}

	p(userConverted)
	p()
	p(string(userConverted))
}

func GetReqExample(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil

}

func GetUser(userID string) (*User, error) {

	req, err := Get(fmt.Sprintf("%s/users/%s", base, userID))
	if err != nil {
		log.Fatal(err)
	}

	var dataResponse Data

	if err := json.Unmarshal(req, &dataResponse); err != nil {
		return nil, err
	}

	return &dataResponse.User, nil
}

func Post(url string, data interface{}) ([]byte, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil

}

func CreateUser(name string, job string) (*UserBody, error) {
	user := &UserBody{
		Name: name,
		Job:  job,
	}

	req, err := Post(fmt.Sprintf("%s/users", base), user)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(req, user); err != nil {
		return nil, err
	}

	// user.Age = 27
	// user.DeletedAt = time.Now()

	return user, nil

}
