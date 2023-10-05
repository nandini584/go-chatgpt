package main
import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/go-resty/resty/v2"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
	apiKey = "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"
)

func main(){
	client := resty.New()
	response, err:= client.R().
	SetHeader("Authorization","Bearer "+apiKey).
	SetHeader("Content-Type","application/json").
	SetBody(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []interface{}{map[string]interface{}{"role":"system","content":"Hi can you generate a description of CNCF organization that comes in LFX programs in about 200 words?"}},
		"max_tokens": 200,
	}).Post(apiEndpoint)

	if err!= nil {
		log.Fatalf("Error while sending the response, Be patient")
	}

	body:=response.Body()
	fmt.Println(string(body))
	var data map[string]interface{}
	err= json.Unmarshal(body, &data)
	if err!= nil {
		fmt.Println("Error while generating your response", err)
		return
	}
	choices, ok:=data["choices"].([]interface{})
	if !ok || len(choices) ==0{
		fmt.Println("No response choices found")
		return
	}
	content := choices[0].
	(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	fmt.Println(content)
}