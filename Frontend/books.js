const bookTable = document.querySelector('.book-table tbody')
console.log(bookTable)

const searchForm = document.querySelector('form')
console.log(searchForm)

searchForm.addEventListener('submit', event => {

    event.preventDefault()

    try {

        const book = {
            title: document.querySelector('#title').value.trim()
        }

        loadBooksWithTitle(book)
        
    } catch (error) {
        console.log("[-] Error: ", error)
    }
});

function addBookToTable(book) {
    
    const row = document.createElement('tr');

    ['title', 'author', 'isbn'].forEach(field => {
        const td = document.createElement('td');
        td.textContent = book[field];
        row.appendChild(td);
    });

    bookTable.appendChild(row)
}

async function loadBooksWithTitle(book) {
    
    try {

        const response = await fetch('http://localhost:8080/books/find', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(book.title)
        });
        const booksWithTitle = await response.json()

        bookTable.innerHTML = ''
        booksWithTitle.forEach(bookwt => addBookToTable(bookwt))

    } catch (error) {
        console.log("[-] Error: ", error)
    }
}

async function loadBooksFromDB() {
    
    try {

        const response = await fetch('http://localhost:8080/books')
        const books = await response.json()

        bookTable.innerHTML = ''
        books.forEach(book => addBookToTable(book))

    } catch (error) {
        console.log("[-] Error:", error)
    }
}

loadBooksFromDB()