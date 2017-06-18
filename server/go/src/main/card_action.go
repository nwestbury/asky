package main

import (
	"log"
	"errors"
)

func check_card_owner(user_id int, card_id int) (bool, error) {
	var card_owner bool
	err := db.QueryRow(`SELECT EXISTS(SELECT from main_schema.card_owner
                       WHERE user_id = $1 AND card_id = $2)`, user_id, card_id).Scan(&card_owner)
	return card_owner, err
}

func check_deck_owner(user_id int, deck_id int) (bool, error) {
	var deck_owner bool
	err := db.QueryRow(`SELECT EXISTS(SELECT from main_schema.deck_owner
                       WHERE user_id = $1 AND deck_id = $2)`, user_id, deck_id).Scan(&deck_owner)
	return deck_owner, err
}

func check_collection_owner(user_id int, collection_id int) (bool, error) {
	var collection_owner bool
	err := db.QueryRow(`SELECT EXISTS(SELECT from main_schema.collection_owner
                        WHERE user_id = $1 AND collection_id = $2)`, user_id, collection_id).Scan(&collection_owner)
	return collection_owner, err
}


func replace_card(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.CardID <= 0 || rec.CardData == ""{
		resp.Message = "Please provide card_id and card_data"
		resp.Success = false
		return errors.New("Missing fields")
	}

	if owner, err := check_card_owner(user_id, rec.CardID); err != nil || !owner {
		log.Print(err)
		resp.Message = "Good try mofo"
		resp.Success = false
		return err
	}

	tx, err := db.Begin()

	_, err = tx.Exec(`UPDATE main_schema.card SET contents = $1 WHERE id = $2`, rec.CardData, rec.CardID)

	if err != nil{
		resp.Message = "Can't update"
		resp.Success = false
		tx.Rollback()
		return err
	}

	var card_diff_id int = 0
	err = db.QueryRow(`SELECT MAX(id)+1 FROM main_schema.card_diff WHERE card_id = $1`, rec.CardID).Scan(&card_diff_id)
	_, err = tx.Exec(`INSERT INTO main_schema.card_diff  (id, card_id, user_id, index, length, action, data, creation_time) VALUES
                      ($1, $2, $3, 0, $4, 'replace', $5, NOW())`, card_diff_id, rec.CardID, user_id, len(rec.CardData), rec.CardData)

	if err != nil{
		log.Print(err)
		resp.Message = "Can't update diff"
		resp.Success = false
		tx.Rollback()
		return err
	}
	
	_, err = tx.Exec(`UPDATE main_schema.card SET contents = $1 WHERE id = $2`, rec.CardData, rec.CardID)

	if err != nil{
		resp.Message = "Can't update card"
		resp.Success = false
		tx.Rollback()
		return err
	}
	
	err = tx.Commit()

	resp.Message = "Replace card contents"
	resp.Success = true
	resp.CardID  = rec.CardID

	return nil
}

func rename_deck(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.DeckID <= 0 || rec.DeckName == ""{
		resp.Message = "Please provide deck_id and deck_name"
		resp.Success = false
		return errors.New("Missing fields")
	}

	if owner, err := check_deck_owner(user_id, rec.DeckID); err != nil || !owner {
		log.Print(err)
		resp.Message = "Good try mofo"
		resp.Success = false
		return err
	}

	tx, err := db.Begin()

	_, err = tx.Exec(`UPDATE main_schema.deck SET name = $1 WHERE id = $2`, rec.DeckName, rec.DeckID)

	if err != nil{
		resp.Message = "Can't update"
		resp.Success = false
		tx.Rollback()
		return err
	}

	var deck_diff_id int = 0
	err = db.QueryRow(`SELECT MAX(id)+1 FROM main_schema.deck_diff WHERE deck_id = $1`, rec.DeckID).Scan(&deck_diff_id)
	_, err = tx.Exec(`INSERT INTO main_schema.deck_diff  (id, deck_id, user_id, action, data, creation_time) VALUES
                      ($1, $2, $3, 'rename', $4, NOW())`, deck_diff_id, rec.DeckID, user_id, rec.DeckName)

	if err != nil{
		log.Print(err)
		resp.Message = "Can't update diff"
		resp.Success = false
		tx.Rollback()
		return err
	}
	
	_, err = tx.Exec(`UPDATE main_schema.deck SET name = $1 WHERE id = $2`, rec.DeckName, rec.DeckID)

	if err != nil{
		resp.Message = "Can't update deck"
		resp.Success = false
		tx.Rollback()
		return err
	}
	
	err = tx.Commit()

	resp.Message = "Renamed deck"
	resp.Success = true
	resp.DeckID  = rec.DeckID

	return nil
}

func rename_collection(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.CollectionID <= 0 || rec.CollectionName == ""{
		resp.Message = "Please provide collection_id and collection_name"
		resp.Success = false
		return errors.New("Missing fields")
	}

	if owner, err := check_collection_owner(user_id, rec.CollectionID); err != nil || !owner {
		log.Print(err)
		resp.Message = "Good try mofo"
		resp.Success = false
		return err
	}

	tx, err := db.Begin()

	_, err = tx.Exec(`UPDATE main_schema.collection SET name = $1 WHERE id = $2`, rec.CollectionName, rec.CollectionID)

	if err != nil{
		resp.Message = "Can't update"
		resp.Success = false
		tx.Rollback()
		return err
	}

	var collection_diff_id int = 0
	err = db.QueryRow(`SELECT MAX(id)+1 FROM main_schema.collection_diff WHERE collection_id = $1`, rec.CollectionID).Scan(&collection_diff_id)
	_, err = tx.Exec(`INSERT INTO main_schema.collection_diff  (id, collection_id, user_id, action, data, creation_time) VALUES
                      ($1, $2, $3, 'rename', $4, NOW())`, collection_diff_id, rec.CollectionID, user_id, rec.CollectionName)

	if err != nil{
		log.Print(err)
		resp.Message = "Can't update diff"
		resp.Success = false
		tx.Rollback()
		return err
	}
	
	_, err = tx.Exec(`UPDATE main_schema.collection SET name = $1 WHERE id = $2`, rec.CollectionName, rec.CollectionID)

	if err != nil{
		resp.Message = "Can't update collection"
		resp.Success = false
		tx.Rollback()
		return err
	}
	
	err = tx.Commit()

	resp.Message = "Renamed collection"
	resp.Success = true
	resp.CollectionID  = rec.CollectionID

	return err
}

func add_card_to_deck(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.CardID <= 0 || rec.DeckID <= 0 {
		resp.Message = "Please provide card_id and deck_id"
		resp.Success = false
		return errors.New("Missing fields")
	}

	if owner, err := check_card_owner(user_id, rec.CardID); err != nil || !owner {
		log.Print(err)
		resp.Message = "Good try mofo"
		resp.Success = false
		return err
	}

	if owner, err := check_deck_owner(user_id, rec.DeckID); err != nil || !owner {
		log.Print(err)
		resp.Message = "Good try mofo"
		resp.Success = false
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		resp.Message = "TX Failed"
		resp.Success = false
		return err
	}

	_, err = tx.Exec(`INSERT into main_schema.deck_cards (card_id, deck_id) VALUES ($1, $2)`, rec.CardID, rec.DeckID)
	if err != nil {
		resp.Message = "Card already in deck?"
		resp.Success = false
		tx.Rollback()
		return err
	}

	var deck_diff_id int = 0
	err = db.QueryRow(`SELECT MAX(id)+1 FROM main_schema.deck_diff WHERE deck_id = $1`, rec.DeckID).Scan(&deck_diff_id)
	_, err = tx.Exec(`INSERT INTO main_schema.deck_diff (id, deck_id, user_id, card_id, action, creation_time) VALUES
                      ($1, $2, $3, $4, 'insert_card', NOW())`, deck_diff_id, rec.DeckID, user_id, rec.CardID)
	err = tx.Commit()
	
	resp.Message = "Added card to deck"
	resp.Success = true
	resp.DeckID = rec.DeckID
	resp.CardID = rec.CardID
	return err
}

func add_deck_to_collection(user_id int, rec *Recieve, resp *Response) (error) {
	if rec.CollectionID <= 0 || rec.DeckID <= 0 {
		resp.Message = "Please provide collection_id and deck_id"
		resp.Success = false
		return errors.New("Missing fields")
	}

	if owner, err := check_collection_owner(user_id, rec.CollectionID); err != nil || !owner {
		log.Print(err)
		resp.Message = "Good try mofo"
		resp.Success = false
		return err
	}

	if owner, err := check_deck_owner(user_id, rec.DeckID); err != nil || !owner {
		log.Print(err)
		resp.Message = "Good try mofo"
		resp.Success = false
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		resp.Message = "TX Failed"
		resp.Success = false
		return err
	}

	_, err = tx.Exec(`INSERT into main_schema.collection_decks (collection_id, deck_id) VALUES ($1, $2)`, rec.CollectionID, rec.DeckID)
	if err != nil {
		resp.Message = "Deck already in collection?"
		resp.Success = false
		tx.Rollback()
		return err
	}

	var collection_diff_id int = 0
	err = db.QueryRow(`SELECT MAX(id)+1 FROM main_schema.collection_diff WHERE collection_id = $1`, rec.CollectionID).Scan(&collection_diff_id)
	_, err = tx.Exec(`INSERT INTO main_schema.collection_diff (id, collection_id, user_id, deck_id, action, creation_time) VALUES
                      ($1, $2, $3, $4, 'insert_deck', NOW())`, collection_diff_id, rec.CollectionID, user_id, rec.DeckID)
	if err != nil {
		resp.Message = "Deck diff failed"
		resp.Success = false
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	
	resp.Message = "Added deck to collection"
	resp.Success = true
	resp.CollectionID = rec.CollectionID
	resp.DeckID = rec.DeckID
	return err
}
