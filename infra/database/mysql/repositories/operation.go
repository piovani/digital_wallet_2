package repositories

import (
	"database/sql"

	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database/mysql"
)

type OperationRepository struct {
	db *sql.DB
}

func NewOperationRepository() *OperationRepository {
	conn, _ := mysql.GetClient()
	return &OperationRepository{
		db: conn,
	}
}

func (o OperationRepository) Insert(opt domain.Operation) error {
	statement, err := o.db.Prepare("INSERT INTO operations () VALUE ()")
	if err != nil {
		return err
	}

	defer statement.Close()

	// r, err := statement.Exec(opt.)

	// resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	// if erro != nil {
	// 	return 0, erro
	// }

	// ultimoIdInserido, erro := resultado.LastInsertId()
	// if erro != nil {
	// 	return 0, erro
	// }

	// return uint64(ultimoIdInserido), nil

	// fmt.Println(operation)
	return nil
}
