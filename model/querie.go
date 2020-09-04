package model

import (
	"encoding/json"

	s "challenge/api/service"

	"github.com/rafael180496/libcore/database"
)

var (
	/*Queries : quieries del proyecto */
	Queries = map[string]string{
		"sql01": `
		select  coalesce(count(*),0) as cont
		from users
		where username=trim(:USER)
		and pass=trim(:PASS)`,
		"sql02": `
		select clientId,name,platformId,segment1,segment2,segment3,segment4
		from clients
		order by :ORDER
		limit :CANT offset :INI
		`,
		"sql03": `
		select  coalesce(count(*),0) as cont
		from users AS  U,tokens AS T
		where u.username=t.username
		and t.token=trim(:TOKEN)
		`,
		"ins01": `
		insert into tokens (username, token)
		values (:USER,:TOKEN)
		`,
		"del01": `delete from tokens where username=trim(:USER)`,
		"del02": `delete from tokens where token=trim(:TOKEN)`,
	}
)

/*SetQuerie : queries backup para la base de datos*/
func SetQuerie() {
	s.DbCx.SetBackupScript(`
	create table clients
(
    clientId   int     not null
        constraint clients_pk
            primary key,
    platformId int     not null,
    name       text    not null,
    segment1   boolean not null,
    segment2   boolean not null,
    segment3   boolean not null,
    segment4   boolean not null
);

create table tokens
(
    username text,
    token    text not null,
    creacion text default current_timestamp not null,
    constraint tokens_pk
        unique (token, username)
);

create table users
(
    username text not null
        constraint users_pk
            primary key,
    pass     text not null
);
	
	insert into users (username, pass)
	values ('prueba','abc123');
	`)
}

/*SendSQL : envia un sql para la base de datos */
func SendSQL(cod string) string {
	return Queries[cod]
}

/*SendDB : envia una conexion ala base de datos valida */
func SendDB() database.StConect {
	return s.DbCx
}

/*ExecSQLMsj : ejecuta y captura un mensaje generico de la data*/
func ExecSQLMsj(info interface{}, sql string) ([]byte, string) {
	var (
		result []database.StSQLData
	)
	cnx := SendDB()
	datain, err := json.Marshal(&info)
	if err != nil {
		return nil, "PET58"
	}
	err = cnx.QueryStruct(&result, database.StQuery{
		Querie: sql,
		Args: map[string]interface{}{
			"DATA": datain,
		},
	}, false)
	if len(result) <= 0 || err != nil {
		return nil, "PET59"
	}
	if result[0].Data == nil {
		return nil, "PET59"
	}
	data, erraux := result[0].ToMap()
	_, ok := data[0]["msj"]
	if erraux != nil {
		return nil, "PET60"
	}
	if ok {
		return nil, data[0]["msj"].(string)
	}
	return result[0].Data, ""
}
