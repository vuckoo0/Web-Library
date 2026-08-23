const bookTable = document.querySelector('.book-table tbody');
const searchForm = document.querySelector('form');

function addBookToTable(book) {
    
    const row = document.createElement('tr');

    ['id', 'title', 'author', 'isbn'].forEach(field => {
        const td = document.createElement('td');
        td.textContent = book[field];
        row.appendChild(td);
    });

    bookTable.appendChild(row);
}

async function loadBooksWithTitle(book) {
    
    try {

        const response = await fetch(`http://localhost:8080/books/search?title=${encodeURIComponent(book.title)}`);
        const booksWithTitle = await response.json();

        bookTable.innerHTML = '';
        booksWithTitle.forEach(bookwt => addBookToTable(bookwt));

    } catch (error) {
        console.log("[-] Error: ", error);
    }
}

async function loadBooksFromDB() {
    
    try {

        const response = await fetch('http://localhost:8080/books');
        const books = await response.json();

        bookTable.innerHTML = '';
        books.forEach(book => addBookToTable(book));

    } catch (error) {
        console.log("[-] Error:", error);
    }
}

searchForm.querySelector('#refresh').addEventListener('click', () => {
    document.querySelector('#search').value = '';
    loadBooksFromDB();
});

searchForm.addEventListener('submit', event => {

    event.preventDefault();

    try {

        const book = {
            title: document.querySelector('#search').value.trim()
        };

        loadBooksWithTitle(book);
        
    } catch (error) {
        console.log("[-] Error: ", error);
    }
});

loadBooksFromDB();