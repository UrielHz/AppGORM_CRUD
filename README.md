# AppGORM_CRUD

Nombre alumno: Oscar Uriel Pasillas Hernández
Matricula: 200456
Grupo: IDGS 10-A

Proyecto en GO usando ORM para realizar un API REST CRUD para MySQL (en mi caso, MySQL Workbench)
Para probarlo utilice POSTMAN para hacer las peticiones con la siguiente escrtuctura del json, sea para crear nuevo registro o actualizar:

{
    "Title": "Nueva Tarea",
    "Completed": false
}

Al realizar un DELETE, se elimina el formato json en POSTMAN o revisar en http://localhost:8080/api/v1/todos/ 
Pero en la base de datos solo se registra la fecha de eliminación y el registro permanece.
