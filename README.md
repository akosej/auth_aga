# AGA ![aga](./aga.ico) AUTH

## Sistema de Autenticación Centralizada de la UHO

El Sistema de Autenticación Centralizada (**AGA_AUHT**) permite que todas las aplicaciones y sistemas desarrollados e implementados en la Universidad de Holguín (UHO) utilicen un registro único de usuarios. Esto garantiza una uniformidad en el uso de credenciales y mejora la seguridad de los datos personales.

## Modo de Uso

1. Solicitar el token de su aplicacion. `Debe de enviar el dominio al que responde su aplicacion`
2. Agregar en los headers de la Aplicacion  `'Authorization': 'AGA token'`

## Urls disponibles

+ `/`
    + Documentacion del AGA_AUTH
        + Login de prueba
+ `/login`
    + Esta url es para iniciar sesión en el sistema, debe enviar BODY en formato raw con application/json por método POST:

        + `{"username": "$val1", "password": "$val2"}`
            + `$val1` _Specify the user_
            + `$val2` _Type the password in plain text_
    + La salida es un json:

      + ```json
            {
        "OK": true,
        "activeUser": {
            "status": 200,
            "account_state": "TRUE|FALSE",
            "uid": "usuario",
            "personal_information": {
              "dni": "",
              "cn": "Nombre completo",
              "given_name": "Nombre",
              "sn": "Apellidos",
              "PersonalPhoto": "base64",
	            "Overlapping":   "base64"
            },
            "account_info": {
              "user_type": "Institucional|Trabajadores|Estudianes",
              "expire_date": "",
              "description": "",
              "title": "",
              "initials": "update",
              "create_user": "[user]",
              "create_date": "2022-10-22 02:02:12 am",
              "modify_user": "user",
              "modify_data": "2022-10-22 03:35:08 am",
              "password": {
              "user_password_set": "2022-10-22 03:35:08 am",
              "pass_valid": "Expirado",
              "pass_set": "191"
            },
        },
        "message": "success"
        }
        ```
    + Estados
  
        + ``200`` _success_
        + ``401`` _Credenciales incorrectas
        + ``409`` _Format error in user or password_
        + ``500`` _could not login_
+ `/logout`
  + Esta URL es para cerrar sesión en el sistema, método POST
  + Salida
  
    + 
      ```json 
        {
          "message": "success|false",
        }
        ``` 
    
  + Status
    + `200` _Success_
    + `401` _Error_
+ `/active`
  + Devuelve los datos del usuario conectado con la estructura de (activeUser)