package stocker

import (
    "gopkg.in/doug-martin/goqu.v4"
    _ "gopkg.in/doug-martin/goqu.v4/adapters/postgres"
    "log"
)

type Stocker struct {
    DB *goqu.Database
}

type WithTxFunction func(tx *goqu.TxDatabase) error

func (st *Stocker) withTx(f WithTxFunction) error {
    tx, err := st.DB.Begin()
    if err != nil {
        log.Println(err)
        tx.Rollback()
        return err
    }

    err = f(tx)
    if err != nil {
        tx.Rollback()
        return err
    } else {
        tx.Commit()
        return nil
    }
}

func (st *Stocker) AddStock(name string, amount int) error {
    return st.withTx(func(tx *goqu.TxDatabase) error {
        var itemId int
        _, err := tx.From("items").
                Returning(goqu.I("id")).
                Insert(goqu.Record{
                    "name": name,
                }).ScanVal(&itemId)
        if err != nil {
            log.Println(err)
            return err
        }

        stockRecords := []goqu.Record{}
        for i := 0; i < amount; i++ {
            stockRecords = append(stockRecords, goqu.Record{"item_id": itemId})
        }
        _, err = tx.From("stocks").
                Insert(stockRecords).
                Exec()
        if err != nil {
            log.Println(err)
            return err
        }

        return nil
    })
}

func (st *Stocker) DeleteAll() error {
    return st.withTx(func(tx *goqu.TxDatabase) error {
        _, err := tx.From("stocks").Delete().Exec()
        if err != nil {
            return err
        }
        _, err = tx.From("items").Delete().Exec()
        if err != nil {
            return err
        }

        return nil
    })
}