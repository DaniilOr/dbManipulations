package service

import (
	"context"
	"errors"
	"github.com/DaniilOr/dbManipulations/src/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)
var (
	errNoEnv = errors.New("Envirinment was not set")
)
type ServiceInterface interface {
	GetCards(int64) ([]models.CardDTO, error)
	GetTransactions(int64) ([]models.TransactionsDTO, error)
	GetMostSpent(int64)(string, int64, error)
	GetMostVisited(int64)(string, int64, error)
}
type Service struct {
	db  *pgxpool.Pool
	ctx context.Context
}


func CreateNewService() (s*Service, err error){
	dsn, ok := os.LookupEnv("dsn")
	if !ok{
		log.Println(errNoEnv)
		return nil, errNoEnv
	}
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	ctx := context.Background()
	return &Service{db, ctx}, nil
}

func (s*Service) GetCards(uid int64) ([]models.CardDTO, error){
	rows, err := s.db.Query(s.ctx, `
	SELECT id, issuer, type, number FROM cards
	WHERE owner_id=$1
	LIMIT 50
`,
	uid)
	if err != nil{
		log.Println(err)
		return []models.CardDTO{}, err
	}
	cardDTOs := []models.CardDTO{}
	for rows.Next() {
		var card models.CardDTO
		rows.Scan(
			&card.Id,
			&card.Issuer,
			&card.Type,
			&card.Number,
		)
		cardDTOs = append(cardDTOs, card)
	}
	if rows.Err() != nil{
		log.Println(rows.Err())
		return []models.CardDTO{}, rows.Err()
	}
	return cardDTOs, nil
}

func (s*Service) GetTransactions(cid int64) ([]models.TransactionsDTO, error){
	rows, err := s.db.Query(s.ctx, `
	SELECT id, mcc, icon_id, amount, date FROM transactions
	WHERE card=$1
	LIMIT 50
`, cid)
	if err != nil{
		log.Println(err)
		return []models.TransactionsDTO{}, err
	}
	transactions := []models.TransactionsDTO{}
	for rows.Next(){
		var transaction models.TransactionsDTO
		rows.Scan(
			&transaction.Id,
			&transaction.Mcc,
			&transaction.IconId,
			&transaction.Amount,
			&transaction.Date,
			)
		transactions = append(transactions, transaction)
	}
	if rows.Err() != nil{
		log.Println(rows.Err())
		return []models.TransactionsDTO{}, err
	}
	return transactions, nil
}
func (s*Service) GetMostSpent(cid int64) (string, int64, error){
	rows, err := s.db.Query(s.ctx, `
	SELECT mcc, amount FROM transactions
	WHERE card=$1 AND AMOUNT < 0
	LIMIT 50
`, cid)
	if err != nil{
		log.Println(err)
		return "", 0, err
	}
	mapper := make(map[string]int64)
	for rows.Next(){
		var transaction models.TransactionsDTO
		rows.Scan(
			&transaction.Mcc,
			&transaction.Amount,
			)
		mapper[transaction.Mcc] += -transaction.Amount

	}
	if rows.Err() != nil{
		log.Println(rows.Err())
		return "", 0, rows.Err()
	}
	var maxKey string
	max := int64(0)
 	for key, value := range mapper{
 		if value > max{
 			max = value
 			maxKey = key
		}
	}
	return  maxKey, max, nil
}
func (s*Service) GetMostVisited(cid int64) (string, int64, error){
	rows, err := s.db.Query(s.ctx, `
	SELECT mcc, amount FROM transactions
	WHERE card=$1 AND AMOUNT < 0
	LIMIT 50
`, cid)
	if err != nil{
		log.Println(err)
		return "", 0, err
	}
	mapper := make(map[string]int64)
	for rows.Next(){
		var transaction models.TransactionsDTO
		rows.Scan(
			&transaction.Mcc,
			&transaction.Amount,
		)
		mapper[transaction.Mcc]++

	}
	if rows.Err() != nil{
		log.Println(rows.Err())
		return "", 0, rows.Err()
	}
	var maxKey string
	max := int64(0)
	for key, value := range mapper{
		if value > max{
			max = value
			maxKey = key
		}
	}
	return  maxKey, max, nil
}
