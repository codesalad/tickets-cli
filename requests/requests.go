package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Get(url string, headers map[string]string) (string, int) {

	bufferedReader := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", url, bufferedReader)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.Body != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Sprintf("Response body: %s", string(bodyBytes)), resp.StatusCode
		} else {
			log.Println("Response body is nil")
		}
	}

	_, err = bufferedReader.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	return bufferedReader.String(), resp.StatusCode
}

func Post(url string, payload interface{}, headers map[string]string) (string, int) {

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return err.Error(), -1
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err.Error(), -1
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err.Error(), -1
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code")
		if resp.Body != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Sprintf("Response body: %s", string(bodyBytes)), resp.StatusCode
		} else {
			fmt.Println("Response body is nil")
		}
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err.Error(), -1
	}

	return buf.String(), resp.StatusCode
}

func Patch(url string, payload interface{}, headers map[string]string) (string, int) {

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return err.Error(), -1
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err.Error(), -1
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err.Error(), -1
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code")
		if resp.Body != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Sprintf("Response body: %s", string(bodyBytes)), resp.StatusCode
		} else {
			fmt.Println("Response body is nil")
		}
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err.Error(), -1
	}

	return buf.String(), resp.StatusCode
}


func Delete(url string, headers map[string]string) (string, int) {

	bufferedReader := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", url, bufferedReader)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.Body != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Sprintf("Response body: %s", string(bodyBytes)), resp.StatusCode
		} else {
			log.Println("Response body is nil")
		}
	}

	_, err = bufferedReader.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	return bufferedReader.String(), resp.StatusCode
}


