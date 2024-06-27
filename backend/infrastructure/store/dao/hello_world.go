package dao

import "myapp/domain/model"

// HelloWorldTable はハローワールドのテーブルを表したモデル
type HelloWorldTable struct {
	Lang    string `db:"lang"`
	Message string `db:"message"`
}

// ConvertHelloWorldTableToDomainHelloWorld はテーブルのモデルからドメインモデルに変換する
func ConvertHelloWorldTableToDomainHelloWorld(ht HelloWorldTable) model.HelloWorld {
	return model.HelloWorld{
		Lang:    ht.Lang,
		Message: ht.Message,
	}
}
