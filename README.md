## Task Terakhir VIX Fullstack Developer BTPNS

### Getting Started With Go Backend App
1. Lakukan konfigurasi pada folder ```condig/db.go```
2. Jalankan command ```go run main.go```

### Tech Stack
1. Go
2. Gorilla Mux
3. Gorm
4. MySQL

<hr>

### Documentation
**Documentation**

#### User
1. ##### Get All Book ```METHOD GET```
    endpoint : ```/book```    
    json response body: 
    ```
    [
        {
            "id": 1,
            "name": "One Piece",
            "author": "Oda",
            "publication": "Gramedia"
        },
        {
            "id": 2,
            "name": "Hujan",
            "author": "Tere Liye",
            "publication": "Gramedia"
        },
        {
            "id": 3,
            "name": "Harry Potter",
            "author": "J. K. Rowling",
            "publication": "Kompas"
        },
        {
            "id": 4,
            "name": "Percy Jackson",
            "author": "Rick Riordan",
            "publication": "Tempo"
        }
    ]
    ```

2. ##### Create Book ```METHOD POST```
    endpoint : ```/book```    
    json request body: 
    ```
    {
        "Name":"Hujan",
        "Author":"Tere Liye",
        "Publication":"Gramedia"
    }
    ```
    json response body: 
    ```
    {"id":2,"name":"Hujan","author":"Tere Liye","publication":"Gramedia"}
    ```

3. ##### Get Book by Id ```METHOD GET```
    endpoint : ```/book/4```    
    json response body:
    ``` 
    {
        "id": 4,
        "name": "Percy Jackson",
        "author": "Rick Riordan",
        "publication": "Tempo"
    }
    ```
4. ##### Update Book by Id ```METHOD PUT```
    endpoint : ```/book/4```    
    json request body:
    ``` 
    {
        "name": "Percy Jackson & the Olympians",
        "author": "Rick Riordan",
        "publication": "Kompas"
    }
    ```
    json response body:
    ```
    {
        "id": 4,
        "name": "Percy Jackson & the Olympians",
        "author": "Rick Riordan",
        "publication": "Kompas"
    }
    ```
5. ##### Delete Book by Id ```METHOD DELETE```
    endpoint : ```/book/3```    
    json response body:
    ``` 
    {
        "id": 0,
        "name": "",
        "author": "",
        "publication": ""
    }
    ```