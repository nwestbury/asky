package main

import (
	"errors"
)

func create_card(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.CardData == "" {
		resp.Message = "Missing card data"
		resp.Success = false
		return errors.New("No title")
	}

	var card_id int = -1
	var card_data string = rec.CardData
	tx, err := db.Begin()
	
	err = tx.QueryRow(`INSERT INTO main_schema.card (contents, creation_time)
                      VALUES ($1, NOW()) RETURNING id`, card_data).Scan(&card_id)

	if err != nil || card_id == -1 {
		resp.Message = "Failed Card Creation"
		resp.Success = false
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO main_schema.card_diff (id, card_id, user_id, index, length, action, data, creation_time) VALUES
                      (0, $1, $2, 0, $3, 'insert', $4, NOW())`, card_id, user_id, len(card_data), card_data)

	if err != nil {
		resp.Message = "Failed card diff"
		resp.Success = false
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(`INSERT INTO main_schema.card_owner (user_id, card_id, creation_time) VALUES
                      ($1, $2, NOW())`, user_id, card_id)

	if err != nil {
		resp.Message = "Failed card owner"
		resp.Success = false
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	resp.Message = "Created card"
	resp.Success = true
	resp.CardID  = card_id

	return nil
}

func create_deck(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.DeckName == "" {
		resp.Message = "Missing deck name"
		resp.Success = false
		return errors.New("No name")
	}

	var deck_id int = -1
	var deck_name string = rec.DeckName
	tx, err := db.Begin()
	
	err = tx.QueryRow(`INSERT INTO main_schema.deck (name, creation_time)
                      VALUES ($1, NOW()) RETURNING id`, deck_name).Scan(&deck_id)

	if err != nil || deck_id == -1 {
		resp.Message = "Failed Deck Creation"
		resp.Success = false
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO main_schema.deck_diff (id, deck_id, user_id, action, data, creation_time) VALUES
                      (0, $1, $2, 'rename', $3, NOW())`, deck_id, user_id, deck_name)

	if err != nil {
		resp.Message = "Failed deck rename"
		resp.Success = false
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(`INSERT INTO main_schema.deck_owner (user_id, deck_id, creation_time) VALUES
                      ($1, $2, NOW())`, user_id, deck_id)

	if err != nil {
		resp.Message = "Failed deck owner"
		resp.Success = false
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	resp.Message = "Created deck"
	resp.Success = true
	resp.DeckID  = deck_id

	return nil
}

func create_collection(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.CollectionName == "" {
		resp.Message = "Missing collection name"
		resp.Success = false
		return errors.New("No name")
	}

	var collection_id int = -1
	var collection_name string = rec.CollectionName
	tx, err := db.Begin()
	
	err = tx.QueryRow(`INSERT INTO main_schema.collection (name, creation_time)
                      VALUES ($1, NOW()) RETURNING id`, collection_name).Scan(&collection_id)

	if err != nil || collection_id == -1 {
		resp.Message = "Failed collection creation"
		resp.Success = false
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO main_schema.collection_diff (id, collection_id, user_id, action, data, creation_time) VALUES
                      (0, $1, $2, 'rename', $3, NOW())`, collection_id, user_id, collection_name)

	if err != nil {
		resp.Message = "Failed collection rename"
		resp.Success = false
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(`INSERT INTO main_schema.collection_owner (user_id, collection_id, creation_time) VALUES
                      ($1, $2, NOW())`, user_id, collection_id)

	if err != nil {
		resp.Message = "Failed collection owner"
		resp.Success = false
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	resp.Message = "Created collection"
	resp.Success = true
	resp.CollectionID = collection_id

	return nil
}

func list_cards(user_id int, rec *Recieve, resp *Response) (error) {
	return nil
}

func list_decks(user_id int, rec *Recieve, resp *Response) (error) {
	return nil
}

func list_collections(user_id int, rec *Recieve, resp *Response) (error) {
	return nil
}
