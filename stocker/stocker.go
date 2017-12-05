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

type ItemDTO struct {
    Name string `db:"name"`
    Amount int `db:"count"`
}
func (st *Stocker) CheckStock(name string) ([]ItemDTO, error) {
    if name == "" {
        return st.checkAllStock()
    } else {
        itemDTO, err := st.checkStockByName(name)
        return []ItemDTO{ itemDTO }, err
    }
}

func (st *Stocker) checkStockByName(name string) (ItemDTO, error) {
    var itemDTO ItemDTO
    err := st.withTx(func(tx *goqu.TxDatabase) error {
        _, err := stmtItemNameAndAmount(tx).Where(goqu.Ex{"items.name": name}).ScanStruct(&itemDTO)
        return err
    })

    return itemDTO, err
}

func (st *Stocker) checkAllStock() ([]ItemDTO, error) {
    var itemDTOs []ItemDTO
    err := st.withTx(func(tx *goqu.TxDatabase) error {
        err := stmtItemNameAndAmount(tx).ScanStructs(&itemDTOs)
        return err
    })

    return itemDTOs, err
}

func stmtItemNameAndAmount(tx *goqu.TxDatabase) *goqu.Dataset {
    return tx.From("stocks").
        Select("name", goqu.COUNT("stocks.id")).
        InnerJoin(goqu.I("items"), goqu.On(goqu.I("items.id").Eq(goqu.I("stocks.item_id")))).
        GroupBy("items.name").
        Where(goqu.Ex{"stocks.sold": false})
}