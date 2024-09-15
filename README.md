# Golang y Vue3js

## Server

El servidor esta hecho con el lenguaje de programación Go (Golang), este es un pequeño proyecto donde sirve para poder exponer una API Rest y que el cliente (hecho con Vue) pueda consumirla.

## Comandos

```bash
#Comando principal para que funcione el servidor, pero ¿Y los datos?
go run main.go

#Con este comando se llenará automaticamente la base de datos con los datos de los clientes.
go run main.go artisan --seed

#Con este comando se vaciará la tabla clientes
go run main.go artisan --clear

#Con este coamndo se rellenará la tabla cliente pero solo ciertos registros (El json donde coge los datos tiene 998 registros, esto ayudará a probar la paginación)
go run main.go artisan --records=<número de registros>
```

Pronto se hará más autoatizaciones por el lado del servidor.

## Vista

La vista (frontend) esta hecho con Vue, TypeScript y Boostrap


## Docker y Nginx (Pronto)



