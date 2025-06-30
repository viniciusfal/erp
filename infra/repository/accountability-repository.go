package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/viniciusfal/erp/infra/model"
)

type AccountabilityRepository struct {
	connection *sql.DB
}

func NewAccountabilityRepository(connection *sql.DB) AccountabilityRepository {
	return AccountabilityRepository{
		connection: connection,
	}
}

func (ar *AccountabilityRepository) CreateAcc(acc *model.Accountability) (*string, error) {
	var id *string
	acc_uuid := uuid.NewString()
	acc_created_at := time.Now()

	query, err := ar.connection.Prepare("INSERT INTO accountability" +
		"(id, send_date, resp_id, deb, pix, coin, total_of_day, vias, guiche, created_at, resp_name)" +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id")
	if err != nil {
		fmt.Printf("erro ao preparar a criação de uma nova accountability: %s", err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(
		acc_uuid,
		acc.Send_date,
		acc.Resp_id,
		acc.Deb,
		acc.PIX,
		acc.Coin,
		acc.Total_of_Day,
		acc.Vias,
		acc.Guiche,
		acc_created_at,
		acc.Resp_name,
	).Scan(&id)
	if err != nil {
		fmt.Printf("erro ao executar a criação de uma nova accountability: %s", err)
		return nil, err
	}

	return id, nil
}

func (ar *AccountabilityRepository) GetAccBydate(startDate, endDate time.Time) ([]*model.Accountability, error) {

	var accs []*model.Accountability
	query := "SELECT * FROM accountability WHERE send_date BETWEEN $1 AND $2"

	rows, err := ar.connection.Query(query, startDate, endDate)
	if err != nil {
		fmt.Printf("Erro ao fazer a query de accountability por data: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var acc model.Accountability

		err = rows.Scan(
			&acc.ID,
			&acc.Send_date,
			&acc.Resp_id,
			&acc.Deb,
			&acc.PIX,
			&acc.Coin,
			&acc.Total_of_Day,
			&acc.Vias,
			&acc.Guiche,
			&acc.Created_at,
			&acc.Resp_name,
		)

		if err != nil {
			fmt.Printf("Erro ao interar a lista de accountabilities: %s", err)
			return nil, err
		}

		accs = append(accs, &acc)
	}

	if rows.Err(); err != nil {
		fmt.Printf("Erro na interação 2: %S", err)
		return nil, err
	}

	return accs, nil
}

func (ar *AccountabilityRepository) GetAccById(id string) (*model.Accountability, error) {
	var acc model.Accountability

	query := `
        SELECT 
            id, send_date, resp_id, deb, pix, coin,
            total_of_day, vias, guiche, created_at, resp_name
        FROM accountability 
        WHERE id = $1`

	err := ar.connection.QueryRow(query, id).Scan(
		&acc.ID,
		&acc.Send_date,
		&acc.Resp_id,
		&acc.Deb,
		&acc.PIX,
		&acc.Coin,
		&acc.Total_of_Day,
		&acc.Vias,
		&acc.Guiche,
		&acc.Created_at,
		&acc.Resp_name,
	)

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar accountability: %w", err)
	}

	return &acc, nil
}

func (ar *AccountabilityRepository) GetAccByUser(startDate, endDate time.Time, respId string) ([]*model.Accountability, error) {

	var accs []*model.Accountability
	query := "SELECT * FROM accountability WHERE resp_id = $1 AND send_date BETWEEN $2 AND $3"

	rows, err := ar.connection.Query(query, respId, startDate, endDate)
	if err != nil {
		fmt.Printf("Erro ao fazer a query de accountability por data: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var acc model.Accountability

		err = rows.Scan(
			&acc.ID,
			&acc.Send_date,
			&acc.Resp_id,
			&acc.Deb,
			&acc.PIX,
			&acc.Coin,
			&acc.Total_of_Day,
			&acc.Vias,
			&acc.Guiche,
			&acc.Created_at,
			&acc.Resp_name,
		)

		if err != nil {
			fmt.Printf("Erro ao interar a lista de accountabilities: %s", err)
			return nil, err
		}

		accs = append(accs, &acc)
	}

	if rows.Err(); err != nil {
		fmt.Printf("Erro na interação 2: %S", err)
		return nil, err
	}

	return accs, nil
}

func (ar *AccountabilityRepository) SetAcc(acc *model.Accountability) (*model.Accountability, error) {

	query, err := ar.connection.Prepare(`
		UPDATE accountability
		SET
			send_date = $1,
			resp_id = $2,
			deb = $3,
			pix = $4,
			coin = $5,
			total_of_day = $6,
			vias = $7,
			guiche = $8,
			resp_name=$9,
		WHERE
			id = $10
	`)
	if err != nil {
		fmt.Printf("erro ao preparar a atualização da accountability: %s", err)
		return nil, err
	}

	defer query.Close()

	_, err = query.Exec(
		acc.Send_date, acc.Resp_id, acc.Deb, acc.PIX, acc.Coin, acc.Total_of_Day,
		acc.Vias, acc.Guiche, acc.Resp_name, acc.ID,
	)
	if err != nil {
		fmt.Printf("erro ao Executar a atualização da accountability: %s", err)
		return nil, err
	}

	return acc, nil
}

func (ar *AccountabilityRepository) ChangeRequest(acr *model.AccountabilityChangeRequest) (*model.AccountabilityChangeRequest, error) {
	// 1. Validações iniciais
	if acr.OriginalAccountabilityID == "" {
		return nil, fmt.Errorf("original_accountability_id é obrigatório")
	}

	// 2. Obter accountability original com tratamento de erro detalhado
	oldAcc, err := ar.GetAccById(acr.OriginalAccountabilityID)
	if err != nil {
		return nil, fmt.Errorf("falha ao buscar accountability original (ID: %s): %w", acr.OriginalAccountabilityID, err)
	}

	// 3. Serialização com verificação de nil
	var oldAccJSON []byte
	if oldAcc != nil {
		oldAccJSON, err = json.Marshal(oldAcc)
		if err != nil {
			return nil, fmt.Errorf("falha na serialização: %w", err)
		}
	} else {
		oldAccJSON = []byte("{}") // JSON vazio se for nil
	}

	// 4. Configuração de valores padrão
	acr.ID = uuid.NewString()
	createdAt := time.Now()
	if acr.Status == "" {
		acr.Status = "pendente"
	}

	// 5. Query com formatação mais legível
	query := `
        INSERT INTO accountability_change_requests (
            id, original_accountability_id, requested_by, reviewed_by,
            send_date, new_deb, new_pix, new_coin, new_total_of_day,
            new_vias, new_guiche, status, request_reason,
            rejection_reason, created_at, reviewed_at, old_accountability
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
                 $11, $12, $13, $14, $15, $16, $17)`

	// 6. Execução com tratamento detalhado de erros
	_, err = ar.connection.Exec(query,
		acr.ID,
		acr.OriginalAccountabilityID,
		acr.RequestedBy,
		acr.ReviewedBy,
		acr.SendDate,
		acr.NewDeb,
		acr.NewPIX,
		acr.NewCoin,
		acr.NewTotalOfDay,
		acr.NewVias,
		acr.NewGuiche,
		strings.ToLower(acr.Status), // Garante minúsculas
		acr.RequestReason,
		acr.RejectionReason,
		createdAt,
		acr.ReviewedAt,
		oldAccJSON,
	)

	if err != nil {
		// Log detalhado do erro do banco
		fmt.Printf("ERRO NO BANCO: %v\n", err)

		// Verifica se é erro de constraint
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return nil, fmt.Errorf("dados referenciados não existem: %w", err)
		}
		if strings.Contains(err.Error(), "violates not-null constraint") {
			return nil, fmt.Errorf("campo obrigatório não informado: %w", err)
		}

		return nil, fmt.Errorf("erro ao criar solicitação: %w", err)
	}

	// 7. Atualiza os campos gerados
	acr.CreatedAt = createdAt
	acr.OldAccountability = oldAcc

	return acr, nil
}
func (ar *AccountabilityRepository) GetPendingRequests(status string) ([]model.AccountabilityChangeRequest, error) {
	var acrs []model.AccountabilityChangeRequest

	query := "SELECT * FROM accountability_change_requests WHERE status = $1"

	rows, err := ar.connection.Query(query, status)
	if err != nil {
		fmt.Printf("Erro ao buscar acountability_change_request com status pendente: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var acr model.AccountabilityChangeRequest
		var oldAccJSON []byte

		err = rows.Scan(
			&acr.ID,
			&acr.OriginalAccountabilityID,
			&acr.RequestedBy,
			&acr.ReviewedBy,
			&acr.SendDate,
			&acr.NewDeb,
			&acr.NewPIX,
			&acr.NewCoin,
			&acr.NewTotalOfDay,
			&acr.NewVias,
			&acr.NewGuiche,
			&acr.Status,
			&acr.RequestReason,
			&acr.RejectionReason,
			&acr.CreatedAt,
			&acr.ReviewedAt,
			&oldAccJSON,
		)

		if err != nil {
			fmt.Printf("Erro ao interar a lista de AccountabilityChangeRequest: %s", err)
			return nil, err
		}

		if len(oldAccJSON) > 0 {
			var oldAcc model.Accountability
			if err := json.Unmarshal(oldAccJSON, &oldAcc); err != nil {
				return nil, fmt.Errorf("erro ao desserializar old accountability: %v", err)
			}
			acr.OldAccountability = &oldAcc
		}

		acrs = append(acrs, acr)
	}

	if rows.Err(); err != nil {
		fmt.Printf("Erro na interação 2: %S", err)
		return nil, err
	}

	return acrs, nil
}

func (ar *AccountabilityRepository) ApproveAccountabilityChangeRequest(requestID, adminID string) (*model.AccountabilityChangeRequest, error) {
	tx, err := ar.connection.Begin()
	if err != nil {
		return nil, fmt.Errorf("erro ao iniciar transação: %w", err)
	}
	defer tx.Rollback()

	// 1. Buscar a solicitação de alteração
	var request model.AccountabilityChangeRequest
	var oldAccJSON []byte

	err = tx.QueryRow(`
        SELECT 
            id,
            original_accountability_id,
            requested_by,
            send_date,
            new_deb,
            new_pix,
            new_coin,
            new_total_of_day,
            new_vias,
            new_guiche,
            status,
            request_reason,
            created_at,
						old_accountability
        FROM accountability_change_requests
        WHERE id = $1 AND status = 'pendente'`, requestID).Scan(
		&request.ID,
		&request.OriginalAccountabilityID,
		&request.RequestedBy,
		&request.SendDate,
		&request.NewDeb,
		&request.NewPIX,
		&request.NewCoin,
		&request.NewTotalOfDay,
		&request.NewVias,
		&request.NewGuiche,
		&request.Status,
		&request.RequestReason,
		&request.CreatedAt,
		&oldAccJSON,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar solicitação: %w", err)
	}

	if len(oldAccJSON) > 0 {
		var oldAcc model.Accountability
		if err := json.Unmarshal(oldAccJSON, &oldAcc); err != nil {
			return nil, fmt.Errorf("erro ao desserializar old accountability: %w", err)
		}
		request.OldAccountability = &oldAcc
	}

	_, err = tx.Exec(`
        UPDATE accountability SET
            deb = $1,
            pix = $2,
            coin = $3,
            total_of_day = $4,
            vias = $5,
            guiche = $6
        WHERE id = $7`,
		request.NewDeb,
		request.NewPIX,
		request.NewCoin,
		request.NewTotalOfDay,
		request.NewVias,
		request.NewGuiche,
		request.OriginalAccountabilityID)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar accountability: %w", err)
	}

	// 3. Marcar a solicitação como aprovada
	_, err = tx.Exec(`
        UPDATE accountability_change_requests SET
            status = 'Aprovado',
            reviewed_by = $1,
            reviewed_at = NOW()
        WHERE id = $2`,
		adminID, requestID)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar status da solicitação: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("erro ao commitar transação: %w", err)
	}

	request.ReviewedBy = &adminID
	request.ReviewedAt = time.Now()
	request.Status = "aprovado"

	return &request, nil
}

func (ar *AccountabilityRepository) RejectAccountabilityChangeRequest(requestID, adminID, rejectionReason string) (*model.AccountabilityChangeRequest, error) {
	tx, err := ar.connection.Begin()
	if err != nil {
		return nil, fmt.Errorf("erro ao iniciar transação: %w", err)
	}
	defer tx.Rollback()

	// 1. Buscar a solicitação de alteração
	var request model.AccountabilityChangeRequest
	var oldAccJSON []byte

	err = tx.QueryRow(`
        SELECT 
            id,
            original_accountability_id,
            requested_by,
            send_date,
            new_deb,
            new_pix,
            new_coin,
            new_total_of_day,
            new_vias,
            new_guiche,
            status,
            request_reason,
            created_at,
						old_accountability
        FROM accountability_change_requests
        WHERE id = $1 AND status = 'pendente'`, requestID).Scan(
		&request.ID,
		&request.OriginalAccountabilityID,
		&request.RequestedBy,
		&request.SendDate,
		&request.NewDeb,
		&request.NewPIX,
		&request.NewCoin,
		&request.NewTotalOfDay,
		&request.NewVias,
		&request.NewGuiche,
		&request.Status,
		&request.RequestReason,
		&request.CreatedAt,
		&oldAccJSON,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar solicitação: %w", err)
	}

	if len(oldAccJSON) > 0 {
		var oldAcc model.Accountability
		if err := json.Unmarshal(oldAccJSON, &oldAcc); err != nil {
			return nil, fmt.Errorf("erro ao desserializar old accountability: %w", err)
		}
		request.OldAccountability = &oldAcc
	}

	// 2. Apenas atualizar o status para rejeitado (NÃO atualiza a accountability original)
	_, err = tx.Exec(`
        UPDATE accountability_change_requests SET
            status = 'Rejeitado',
            reviewed_by = $1,
            rejection_reason = $2,
            reviewed_at = NOW()
        WHERE id = $3`,
		adminID, rejectionReason, requestID)
	if err != nil {
		return nil, fmt.Errorf("erro ao rejeitar solicitação: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("erro ao commitar transação: %w", err)
	}

	// Atualizar os campos no objeto retornado
	request.ReviewedBy = &adminID
	request.ReviewedAt = time.Now()
	request.RejectionReason = rejectionReason
	request.Status = "rejeitado"

	return &request, nil
}
