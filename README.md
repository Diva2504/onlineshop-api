# Final Project 4 (ONLINESHOP-API)
Project ini adalah sebuah backend service dengan interface berupa API endpoint. Data disimpan pada database dengan migrasi automatis menggunakan framework gorm. 
Pada project ini menggunakan fungsi authentication untuk mengenali user yang login, dan authorization untuk memberikan akses atau kebijakan tertentu untuk user dan juga 
admin.

## Deployment url
https://onlineshop-api-hacktiv8.herokuapp.com

## API Endpoint
### users
- {POST} /users/login (memberikan data kepada backend yang selanjutnya dicocokan dengan data user yang sudah ada pada database atau sudah registrasi)
- {POST} /users/register (memberikan data kepada backend yang bertujuan untuk membuat user baru dan disimpan ke dalam database)
- {PATCH} /users/topup (memberikan/mengubah salah satu kolom pada user yaitu kolom topup, bertujuan untuk mengisi ulang saldo dompet digital dari user)

### products
- {GET} /products/ (mengambil semua data pada tabel products yang di keluarkan dalam bentuk json)
- {GET} /products/:id (mengambil data spesifik pada tabel products berdasarkan id yang di keluarkan dalam bentuk json)
- {POST} /products/ (memberikan / menulis data ke dalam backend atau database untuk membuat product baru pada tabel product)
- {PUT} /products/:id (mengubah data pada tabel product berdasarkan id)
- {DELETE} /products/:id (menghapus data pada tabel product berdasarkan id)

### categories
- {GET} /categories/ (mengambil semua data pada tabel categories yang di keluarkan dalam bentuk json)
- {POST} /categories/ (memberikan / menulis data ke dalam backend atau database untuk membuat category baru)
- {PATCH} /categories/:id (mengubah data salah satu kolom pada tabel category berdasarkan id)
- {DELETE} /categories/:id (menghapus data pada tabel category berdasarkan id)

### transaction
- {POST} /transactions/ (user yang sudah login dapat memberikan/menulis data pada tabel transaction yang digunakan untuk pembelian barang oleh user
                         dilakukan beberapa pengecekan terlebih dahulu yaitu mengecek data product ada atau tidak, pengecekan stock produk quantity tidak boleh
                         melebihi stock product, user harus mempunyai balance yang cukup)
- {GET} /transactions/my-transaction (user yang sudah login dapat mengakses semua data transaction pada database yang dimiliki oleh user tersebut)
- {GET} /transaction/:user_id (endpoint hanya digunakan oleh admin untuk melihat semua data transaction oleh user berdasarkan user_id)


## TABLE MODEL
- User
------------      ------------          -------------
     id             primary key             -
    full_name       string                  required
    email           string                  Valid email, unique, required
    password        string                  Required, min length 6
    role            string                  required, enum(admin/customer)
    balance         integer                 Required, max 100.000.000, min 0
    created_at      date                    -
    updated_at      date                    -

- Product
------------      ---------------       -------------
     id             primary key             -
    title           string                  required
    price           integer                 Required, min 0, max 50.000.000
    stock           integer                 Required, min 5
    category_id     FK Category             -
    created_at      date                    -
    updated_at      date                    -

- Category
------------                  ---------------       -------------
     id                         primary key             -
    type                        string                  required
    sold_product_amount         integer                 -
    created_at                  date                    -
    updated_at                  date                    -

- Transaction
------------      ------------          -------------
     id             primary key             -
    product_id      FK Product              -
    user_id         FK User                 -
    quantity        integer                 required
    total_price     integer                 required
    created_at      date                    -
    updated_at      date                    -
    
    
## DATABASE
  - PostgreSQL
  
## Request and Response
### users
- {POST} /users/login
    + request
    ```
    {
      “email”:       "string",
      “password”:    "string",
    }
    ```
    + response
    ```
      {
        “token”:     "string",
      }
    ```
    
- {POST} /users/register
   + request
    ```
    {
      “email”:       "string",
      “password”:    "string",
      “full_name”:   "string",

    }
    ```
    + response
    ```
      {
        “id”:         "integer", 
      “full_name”:    "string", 
      “email”:        "string",
      “password”:     "string",
      “balance”:      "integer",
      “created_at”:   "date", 
      }
    ```
    
 - {PATCH} /users/topup   
     + request
    ```
    Header: Authorization
    Body:
    {
          “balance”:           "string",
    }

    ```
    + response
    ```
      {
        “message”:    "string",
      }
    ```

### categories
- {GET} /categories/
    + request
    ```
      Header: Authorization
    ```
    + response
    ```
      [
        {
          “id”:                   "integer", 
          “type”:                 "string", 
          “sold_product_amount”:  "integer",
          “created_at”:           "date", 
        }
      ]
    ```
    
- {POST} /categories/
    + request
    ```
   Header: Authorization
   Body:
    {
      “type”:           "string",
    }

    ```
    + response
    ```
      [
        {
          “id”:                   "integer", 
          “type”:                 "string", 
          “sold_product_amount”:  "integer",
          “created_at”:           "date", 
        }
      ]
    ```
    
- {PATCH} /categories/:id
   + request
    ```
    Header: Authorization
    params: id
    Body:
    {
          “type”:  "string",
    }

    ```
    + response
    ```
      {
        “id”:                  "integer",
        "type":                "string",
        "sold_product_amount": "integer",
        "update_at":           "date", 
      }
    ```
    
- {DELETE} /categories/:id
    + request
    ```
    Header: Authorization
    params: id
    ```
    + response
    ```
      {
        “message”:   "category has been successfully deleted",
      }
    ```

### products
- {GET} /products/
    + request
    ```
    Header: Authorization
    ```
    + response
    ```
      [
        {
            "id":           "integer",
            "title":        "string",
            "price":        "integer",
            "stock":        "integer",
            "category_id":  "integer",
            "created_at"    "date",
         }
      ]  
    ```
- {GET} /products/:id
    + request
    ```
    Header: Authorization
    params:id
    ```
    + response
    ```
        {
            "id":           "integer",
            "title":        "string",
            "price":        "integer",
            "stock":        "integer",
            "category_id":  "integer",
            "created_at"    "date",
         }  
    ```
- {POST} /products/
    + request
    ```
    Header: Authorization
    Body: 
        {
          "title":       "string",
          "price":       "integer",
          "stock":       "integer",
          "category_id": "integer",
        }
    ```
    + response
    ```
        {
            "id":           "integer",
            "title":        "string",
            "price":        "integer",
            "stock":        "integer",
            "category_id":  "integer",
            "created_at"    "date",
         }  
    ```
   
- {PUT} /products/:id
    + request
    ```
    Header: Authorization
    params:id
    Body: 
        {
          "title":       "string",
          "price":       "integer",
          "stock":       "integer",
          "category_id": "integer",
        }
    ```
    + response
    ```
        {
            "id":           "integer",
            "title":        "string",
            "price":        "integer",
            "stock":        "integer",
            "category_id":  "integer",
            "created_at"    "date",
         }  
    ```  
- {DELETE} /products/:id
    + request
    ```
    Header: Authorization
    params:id
    ```
    + response
    ```
        {
           “message”:   "product has been successfully deleted", 
         }  
    ```
### transaction    
- {POST} /transactions/
    + request
    ```
    Header: Authorization
    Body: 
          {
              "product_id": "integer",
              "quantity":   "integer",
          }
    ```
    + response
    ```
        {
           “message”:   "you have successfully purchased the product",
           "transaction_bill":{
                 "total_price":   "integer",
                 "quantity":      "integer",
                 "product_title": "string", 
           }
         }  
    ```     
    
- {GET} /transactions/my-transaction
    + request
    ```
    Header: Authorization
    ```
    + response
    ```
     [
        {
           “id”:            "integer",
           "product_id":    "integer",
           "user_id":       "integer",
           "quantity":      "integer",
           "total_price":   "integer",
           "Product": {
                "id":           "integer",
                "title":        "string",
                "price":        "integer",
                "stock":        "integer",
                "category_id":  "integer",
                "created_at":   "date",
                "updated_at":   "date",
           }
         } 
      ]   
    ``` 
    
- {GET} /transactions/:user_id
    + request
    ```
    Header: Authorization
    ```
    + response
    ```
     [
        {
           “id”:            "integer",
           "product_id":    "integer",
           "user_id":       "integer",
           "quantity":      "integer",
           "total_price":   "integer",
           "Product": {
                "id":           "integer",
                "title":        "string",
                "price":        "integer",
                "stock":        "integer",
                "category_id":  "integer",
                "created_at":   "date",
                "updated_at":   "date",
           },
           "User": {
                 "id":         "integer", 
                 "full_name":    "string", 
                 "email":        "string",
                 "password":     "string",
                 "balance":      "integer",
                 "created_at":   "date", 
                 "updated_at":   "date,"
           }
         } 
      ]   
    ``` 
   
## Cara Menjalankan Code
Pastikan telah menginstal Go dan Docker pada komputer.
Instal Image Postgresql dari website hub.docker.com
gunakan perintah 
```shell
    docker pull postgres
```
Download Repository ini dengan 
```shell
    git clone https://github.com/takadev15/onlineshop-api.git
```
Menuju folder direktori onlineshop-api
kemudian jalankan terlebih dahulu file docker-compose.yml untuk konfigurasi database dengan perintah
```shell
    docker-compose up
```
kemudian untuk menjalankan aplikasi gunakan perintah 
```shell
    go run main.go
```
Server sudah berjalan di port "3030"
endpoint dapat dicoba menggunakan aplikasi REST CLIENT seperti Postman
