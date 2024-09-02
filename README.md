# AGA ![aga](./aga.ico)

Main API of the AGA system, its function is to handle all requests and interact with the different databases. The
system accepts the integration of modules compiled with an extension **.aga** and plugins with an extension
**name_url.plugins** for those that respond by url and **name.Plugins** for those that perform internal functions.

## Base template for plugins

```go
package main

import (
	"github.com/gofiber/fiber/v2"
)

var (
	// -- Specify the permissions that the plugin will use
	Permissions = []string{"read", "write"...}
)

// Function for AGA to extract plugin permissions
func GetPermissions() []string {
	return Permissions
}

// Function for AGA to extract the routes of the plugin
// If you use this function in the compilation you must specify it in the name of the plugin (name_url.plugin)
func GetRoutes(app *fiber.App) error {
	plugin := app.Group("/name")
	plugin.Get(path
	string, handlers ...Handler)
	plugin.Post(path
	string, handlers ...Handler)
	plugin.Put(path
	string, handlers ...Handler)
	plugin.Delete(path
	string, handlers ...Handler)
	return nil
}

//
func CustomFunc(arg ...interface{}) interface{} {
	...
	return interface {}
}
```

### Compile plugins

`go build -buildmode=plugin -trimpath -o name_url.plugins ./main.go `

The binary must be copied to the **plugins** folder of the system

`> cp ./name_url.plugins path_aga/plugins/`

## Function to use static plugins

```go
/*
   ExtractFunctionsPlugins(plugins, function string, args ...interface{}) interface{}
   _ = system.ExtractFunctionsPlugins(
        plugins:  "namePlugin",
        function: "nameFunction",
        args:    ...interface{})
*/

type Oup struct {
Name string
}

result := ExtractFunctionsPlugins("ldap", "connectLdap", ...args) // you can pass any type of data in the args, var, struct...

b, _ := json.Marshal(&result)

var oup Oup

_ = json.Unmarshal(b, &m)

fmt.Println(oup.Name)

```

## Manual compilation

1. Download or clone the repository from [github](https://github.com/agaUHO/backend).
2. Navigate to the root directory of the product and run
3. `> go mod tidy`
4. `> go mod vendor`
5. `> CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -a -ldflags "-extldflags '-static' -s -w" -trimpath -o aga ./main.go`

## Config .env

Change the value of the variables as appropriate in the .env file

```dotenv
#-----------------------------------------------------|
#                         43DG4R                      |
#                       464   3J                      |
#                      64      C                      |
#              3      46              3               |
#              4D     36   3JPH      D4               |
#              646    64     464    646               |
#             3J  4    3J     3J   4  J3              |
#            464   R    64   UH   R   464             |
#           464R46D34J    3JP   4646464646            |
#          464        3J      PH        464           |
#        64              4   4             46         |
#       4                  6                 4        |
#--------------------------------------------------------------------------|
#-- System configuration file                                              |                                            
#-- Created by Team AGA                                                    |
#-- Warning: All password variables must be base64 encoded, not plain text |
#--------------------------------------------------------------------------|
# -- INTERNAL
aga.Domain="uho.edu.cu"
aga.AllowOrigins = "https://aga.uho.edu.cu, https://sgu.uho.edu.cu"
aga.AllowHeaders = "Origin, Content-Type, Accept, Authorization"
aga.Port="8000"
aga.Language="es"
aga.PasswordExpirationThreshold="90"

# -- RPC
aga.rpc.Server="127.0.0.1"
aga.rpc.Port="5300"
# -- JWT
UserKeyAGA="aga_jwt"
SecretKeyAGA="7NvcnBpv255Nzc4NjkKv"

# -- LDAP MASTER
aga.ldap.Host="ip"
aga.ldap.Port="389"
aga.ldap.User="cn=admin,dc=uho,dc=edu,dc=cu"
aga.ldap.Password=""
aga.ldap.BaseDn="dc=uho,dc=edu,dc=cu"

# -- MYSQL
aga.mysql.Host="127.0.0.1"
aga.mysql.Port="3306"
aga.mysql.Username="user"
aga.mysql.Password="pass"
aga.mysql.DB="aga"
aga.AdminUserId=1

# -- MAIL
aga.mail.Domain="uho.edu.cu"
aga.mail.Imap="imap.uho.edu.cu"
aga.mail.Smtp="smtp.uho.edu.cu"
aga.mail.Username="user"
aga.mail.Password="pass"

# -- OrcID
aga.OrcIdClientID="aga.OrcIdClientID"
aga.OrcIdClientSecret="aga.OrcIdClientSecret"

#--------------------------------------------------------------------------|
#-- Module variables                                                       |
#-- Warning: Each module must have its variables, in case some variables   |
#-- are not specified, the system ones will be assumedt                    |
#--------------------------------------------------------------------------|
```

## Start up

1. Give it execution permission `> chmod +x ./aga`
2. Run the compiled binary `>./aga`
3. Open a browser and enter the url [http://127.0.0.1:8000](http://127.0.0.1:8000)

## System routes

+ `/`
    + Output is a html
        + Returns the system documentation
+ `/aga/login`
    + This url is to login to the system, you must send BODY in raw format with application/json by POST method:

        + `{"username": "$val1", "password": "$val2", "securitycode": "$val3"}`
            + `$val1` _Specify the user_
            + `$val2` _Type the password in plain text_
            + `$val3` _Specify the 6-digit 2fa code_
    + Output is a json:

      + ```json
            {
        "OK": true,
        "activeUser": {
            "status": 200,
            "account_state": "TRUE",
            "uid": "akos",
            "personal_information": {
              "dni": "92082643541",
              "cn": "EDGAR JAVIER PEÑA HERNANDEZ",
              "given_name": "EDGAR JAVIER",
              "sn": "PEÑA HERNANDEZ"
            },
            "account_info": {
              "user_type": "Institucional",
              "expire_date": "",
              "description": "",
              "title": "",
              "initials": "update",
              "create_user": "[ej]",
              "create_date": "2022-10-22 02:02:12 am",
              "modify_user": "akos",
              "modify_data": "2022-10-22 03:35:08 am",
              "password": {
              "user_password_set": "2022-10-22 03:35:08 am",
              "pass_valid": "Expirado",
              "pass_set": "191"
            },
            "synchro": {
              "hash_mod_ldap": "",
              "hash_mod_api": "",
              "sync_required": false
            },
            "ou": "corporatives"
            },
            "permissions": [],
            "other_accounts": null,
            "preference": {
              "color": "",
              "theme_dark": "",
              "orc_id": "",
              "facebook": "",
              "twitter": "",
              "instagram": "",
              "telegram": "",
              "whats_app": "",
              "research_gate": "",
              "linked_in": "",
              "google": ""
            },
            "service": {
              "mail": {
                "enabled": "TRUE",
                "access_in": "internac-in",
                "access_out": "internac-out",
                "quota": "1048576"
              },
              "proxy": {
                "enabled": "TRUE",
                "access": "internac",
                "quota": "0"
              },
              "nextcloud": {
                "enabled": "FALSE",
                "quota": "0"
              },
              "Remote": {
                "enabled": "FALSE",
                "phone_number": ""
              },
              "wifi": {
                "enabled": "TRUE",
                "device_number": "1",
                "macaddress": []
              }
            },
            "assets": {
              "area": "",
              "department_number": "",
              "category": "",
              "position": "",
              "category_name": "",
              "subcategory_name": "",
              "profession": ""
            },
            "sigenu": {
              "country": "",
              "student_type": "",
              "carrer": "",
              "faculty": "",
              "course_type": "",
              "scholastic_origin": "",
              "town_university": "",
              "matriculation_end_date": "",
              "re_matriculation_end_date": "",
              "student_status": "",
              "student_year": ""
            },
            "two_factor": {
              "status": "false",
              "qr": ""
            }
        },
        "message": "success"
        }
        ```
    + Status
  
        + ``200`` _success_
        + ``401`` _Credenciales incorrectas | 2fa incorrect_
        + ``409`` _Format error in user or password_
        + ``500`` _could not login_
+ `/aga/logout`
  + This url is to log out of the system, POST method
  + Output
  
    + 
      ```json 
        {
          "message": "success|false",
        }
        ``` 
    
  + Status
    + `200` _Success_
    + `401` _Error_
+ `/aga/active`
  + Returns the data of the logged in user with the structure of (activeUser)
+ `/aga/hasPasswordBeenUsed`
  + This url is to check current or previously used password, you need to send BODY in raw format with application/json by POST method:
    + `{"password": "$val1"}`
        + `$val1` _Type the password in plain text_
  + Output is a json:
    ```json 
      {
        used: true | false
        message: "Coincide con las anteriores" 
                 "No coincide con las anteriores"
      }
    ```
  + Status
  
    + ``200`` _success_
    + ``400`` _The server could not interpret the request_
    + ``401`` _password is not used_
+ `/aga/changePassword`
  + This url is to change password, you need to send BODY in raw format with application/json by PUT method:
    + `{"password": "$val1"}`
        + `$val1` _Type the password in plain text_
  + Output is a json:
    ```json 
      {
        message: "Se cambio la contraseña" 
                 "Esta proporcionando la contraseña vigente"
                 "El servidor no pudo interpretar la solicitud debido a que el cliente empleó una sintaxis no válida"
      }
    ```
  + Status
  
    + ``200`` _success_
    + ``400`` _The server could not interpret the request_
    + ``409`` _password is current_
