**Todolist Buku Sederhana**

**Initiating Your Go Backend Application**
1. Lakukan konfigurasi database pada config/db.go
2. Jalankan command go run main.go

**Tech Stack**
1. Go
2. Gorm
3. Gorilla Mux
4. MySQL

**Documentation**
1. Get All Book (/book, method GET)
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
2. Create Book (/book, method POST)
{
    "Name":"Hujan",
    "Author":"Tere Liye",
    "Publication":"Gramedia"
}
Response
{"id":2,"name":"Hujan","author":"Tere Liye","publication":"Gramedia"}

3.  Get Book by Id (/book/{bookId}, method GET) -> /book/4
{
    "id": 4,
    "name": "Percy Jackson",
    "author": "Rick Riordan",
    "publication": "Tempo"
}
4.  Update Book by Id (/book/{bookId}, method PUT) -> /book/4
{
    "name": "Percy Jackson & the Olympians",
    "author": "Rick Riordan",
    "publication": "Kompas"
}
Response
{
    "id": 4,
    "name": "Percy Jackson & the Olympians",
    "author": "Rick Riordan",
    "publication": "Kompas"
}
5.  Delete Book by Id (/book/{bookId}, method DELETE) -> /book/3
{
    "id": 0,
    "name": "",
    "author": "",
    "publication": ""
}