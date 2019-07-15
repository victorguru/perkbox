package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/victorguru/perkbox/types"
)

// Helper holds DB methods
type Helper int

// createNewCoupon will create a new coupon in the db
func (s *Helper) createNewCoupon(db *sql.DB, name string, brand string, value float64, expiry string) (int64, error) {
	query := `INSERT INTO coupons (name, brand, value, expiry)
			VALUES (?, ?, ?, ?);`

	stm, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	res, err := stm.Exec(name, brand, value, expiry)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return lastInsertedID, nil
}

// getCoupons will get all or one coupon
func (s *Helper) getCoupons(db *sql.DB, id int, limit int, offset int) ([]types.Coupon, error) {
	// Building the SQL
	query := `SELECT * FROM coupons`
	// In case we are looking for an id
	if id != 0 {
		query += ` WHERE id = ?`
	}
	// Making sure the order is always the same
	query += ` ORDER BY id ASC`
	// In case we want to limit/offset the results
	if limit != 0 {
		query += ` LIMIT ? OFFSET ?`
	}

	stm, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rows *sql.Rows

	if id == 0 && offset == 0 {
		// In this case we want to take all the coupons
		rows, err = stm.Query()
	} else if id != 0 {
		// In this case we want to take one coupon
		rows, err = stm.Query(id)
	} else {
		// In this case we want to take a limited and offset number of cupons
		rows, err = stm.Query(limit, offset)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result []types.Coupon
	for rows.Next() {
		var row types.Coupon
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Brand,
			&row.Value,
			&row.CreatedAt,
			&row.Expiry,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}

// updateCoupon will create a new coupon in the db
func (s *Helper) updateCoupon(db *sql.DB, id int, name string, brand string, value float64, expiry string) (int64, error) {
	query := `UPDATE coupons
				SET name = ?, brand = ?, value = ?, expiry = ?
				WHERE id = ?;`

	stm, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	res, err := stm.Exec(name, brand, value, expiry, id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return rowsAffected, nil
}
