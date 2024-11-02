package repository

import (
    "github.com/jackc/pgx/v5/pgxpool"
)


type Database interface{
    connect(url string)(*pgxpool.Pool,error)
}

type ConnectionDb struct{
    EstablishedConnection *pgxpool.Pool
}

func (cd *ConnectionDb)NewConnection(db Database,url string)error{
     dbpool,err := db.connect(url)
     if err != nil{
        return err
     }
     cd.EstablishedConnection = dbpool
     return nil
}



// func ConnectPool(route string)(*pgxpool.Pool, error){
//  dbpool,err := pgxpool.New(context.Background(), route)
//  if err != nil {
// 	return nil, fmt.Errorf("don't connect to data base: %w",err)
//  }
// //  defer dbpool.Close()



//  fmt.Println("Conexion establecida")
//  return dbpool,nil;
// }


