package main

import (
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type Recieve struct {
	Action string    `json:"action"`
	Token string     `json:"token"`
	TokenType string `json:"token_type"`
	CardData string  `json:"card_data,omitempty"`
	DeckName string  `json:"deck_name,omitempty"`
	CollectionName string  `json:"collection_name,omitempty"`
	IDs []int        `json:"ids,omitempty"`
}

type Response struct {
    Message string `json:"msg"`
	Success bool   `json:"success"`
    CardID int     `json:"card_id,omitempty"`
	Cards string   `json:"cards,omitempty"`
	DeckID int     `json:"card_id,omitempty"`
	CollectionID int `json:"collection_id,omitempty"`
}

type FBResponse struct {
	ID string                  `json:"id"`
	Name string                `json:"name"`
	X map[string]interface{}   `json:"-"`
}

var failed_parse_resp, _ = json.Marshal(Response{
	Message: "Failed to parse input",
	Success: false,
})

func handle_action(msg []byte, c *Client) {
	rec := Recieve{}
	err := json.Unmarshal(msg, &rec)

	if err != nil {
		c.send <- failed_parse_resp
		return
	}

	resp := Response{
		Message: "Something went wrong",
		Success: false,
	}

	log.Print("Got the following action ", rec.Action)
	
	user_id := check_login(&rec, &resp, c)
	
	if user_id != -1 {
		uc := UserClient{
			client: c,
			user_id: user_id,
		}
		c.hub.login <- &uc

		switch action := rec.Action; action {
		case "login":
		case "create_card":
			err = create_card(user_id, &rec, &resp)
		case "create_deck":
			err = create_deck(user_id, &rec, &resp)
		case "create_collection":
			err = create_collection(user_id, &rec, &resp)
		case "list_cards":
			err = list_cards(user_id, &rec, &resp)
		case "list_decks":
			err = list_decks(user_id, &rec, &resp)
		case "list_collections":
			err = list_collections(user_id, &rec, &resp)
		default:
			resp.Message = "Unknown action"
			resp.Success = false
		}
		
	}

	if err != nil {
		log.Print(err)
	}

	// c.hub.broadcast <- msg

	byte_msg, _ := json.Marshal(resp)
	c.send <- byte_msg
}

func check_login(rec *Recieve, resp *Response, c *Client) int{

	var user_id int = -1
	
	if rec.TokenType == "fb" {
		http_req, err := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(rec.Token))

		if err != nil {
			resp.Message = "Could not verify token"
			resp.Success = false
			return -1
		}

		http_resp, err := ioutil.ReadAll(http_req.Body)

		if err != nil {
			log.Print(err)
			resp.Message = "Weird read error"
			resp.Success = false
			return -1
		}

		fbresp := FBResponse{}
		if err := json.Unmarshal(http_resp, &fbresp); err != nil {
			log.Print(err)
			resp.Message = "Invalid login"
			resp.Success = false
			return -1
		}
		if err := json.Unmarshal(http_resp, &fbresp.X); err != nil {
			log.Print(err)
			resp.Message = "Invalid login"
			resp.Success = false
			return -1
		}

		if fbresp.ID != "" {
			log.Print("Got userid ", fbresp.ID)
			err := register_if_not_exists("fb", fbresp.ID)
			if err == nil {
				user_id = get_user_id("fb", fbresp.ID)
			}
		}else{
			resp.Message = "Facebook Authentication Failed"
			resp.Success = false
			return -1
		}
	}

	if user_id != -1 {
		resp.Message = "Login Successful"
		resp.Success = true
	}else{
		resp.Message = "Failed to login"
		resp.Success = false
	}

	return user_id
}

func register_if_not_exists(api_type string, api_id string) (error){
	stmt, err := db.Prepare(`INSERT INTO main_schema.user (type, api_id, creation_time)
                             VALUES ($1, $2, NOW()) ON CONFLICT DO NOTHING`)
	if err != nil {
		log.Print("Register failed", err)
		return err
	}
	
	_, err = stmt.Exec(api_type, api_id)
	if err != nil {
		log.Print("Register failed exec", err)
		return err
	}

	return err
}

func get_user_id(api_type string, api_id string) int{
	var id int = -1

	err := db.QueryRow("SELECT id FROM main_schema.user WHERE type=$1 AND api_id=$2 LIMIT 1", api_type, api_id).Scan(&id)
	
	if err != nil {
		log.Print("Fetch user id failed", err)
		return -1
	}
	
	return id
}
