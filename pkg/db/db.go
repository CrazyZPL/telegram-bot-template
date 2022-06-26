package db

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/telegram-bot-template/pkg/config"
	"github.com/telegram-bot-template/pkg/db/ent"
	"golang.org/x/xerrors"
)

func InitMysqlClient(mysqlConfig config.MysqlConfig) (*ent.Client, error) {
	client, err := ent.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database))
	if err != nil {
		return nil, xerrors.Errorf("Failed openning mysql client: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, xerrors.Errorf("Failed creating schema resources: %v", err)
	}

	return client, nil
}
