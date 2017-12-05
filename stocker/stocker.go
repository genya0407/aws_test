package stocker

import (
    "gopkg.in/doug-martin/goqu.v4"
    _ "gopkg.in/doug-martin/goqu.v4/adapters/postgres"
    "log"
    "errors"
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
    ItemId int `db:"id"`
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
        _, err := stmtAvailableItemNameAndAmountByName(tx, name).ScanStruct(&itemDTO)
        return err
    })

    return itemDTO, err
}

func (st *Stocker) checkAllStock() ([]ItemDTO, error) {
    var itemDTOs []ItemDTO
    err := st.withTx(func(tx *goqu.TxDatabase) error {
        err := stmtAvailableItemNameAndAmount(tx).ScanStructs(&itemDTOs)
        return err
    })

    return itemDTOs, err
}

func (st *Stocker) Sell(name string, amount int, price int) error {
    err := st.withTx(func(tx *goqu.TxDatabase) error {
        var itemDTO ItemDTO
        _, err := stmtAvailableItemNameAndAmountByName(tx, name).ScanStruct(&itemDTO)
        if err != nil {
            return err
        }
        if itemDTO.Amount < amount {
            return errors.New("short of amount")
        }

        updateTargets := tx.From("stocks").
            Select("id").
            Where(goqu.Ex{"item_id": itemDTO.ItemId, "sold": false}).
            Limit(uint(amount))
        _, err = tx.From("stocks").
            Where(goqu.I("id").In(updateTargets)).
            Update(goqu.Record{"sold": true, "price": price}).
            Exec()

        return err
    })

    return err
}

func stmtAvailableItemNameAndAmountByName(tx *goqu.TxDatabase, name string) *goqu.Dataset {
    return stmtAvailableItemNameAndAmount(tx).Where(goqu.Ex{"items.name": name})
}

func stmtAvailableItemNameAndAmount(tx *goqu.TxDatabase) *goqu.Dataset {
    return tx.From("stocks").
        Select("items.id", "name", goqu.COUNT("stocks.id")).
        InnerJoin(goqu.I("items"), goqu.On(goqu.I("items.id").Eq(goqu.I("stocks.item_id")))).
        GroupBy("items.id", "items.name").
        Where(goqu.Ex{"stocks.sold": false})
}