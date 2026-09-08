const bookTable = document.querySelector('.book-table tbody');
const searchForm = document.querySelector('form');

const startPageButton = document.querySelector('#start-page-button');
const adminPanelButton = document.querySelector('#admin-panel-button');

const API_URL = `http://${window.location.hostname}:8080`;

function showWarning(message) {
    const warning = document.querySelector('.warning');
    warning.textContent = message;
    warning.style.display = 'block';
}

function hideWarning() {
    document.querySelector('.warning').style.display = 'none';
}

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

        const response = await fetch(`${API_URL}/books/search?title=${encodeURIComponent(book.title)}`, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
        });
        const booksWithTitle = await response.json();

        bookTable.innerHTML = '';
        booksWithTitle.forEach(bookwt => addBookToTable(bookwt));

    } catch (error) {
        console.log("[-] Error: ", error);
    }
}

async function loadBooksFromDB() {
    
    try {

        const response = await fetch(`${API_URL}/books`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
        });
        const books = await response.json();

        bookTable.innerHTML = '';
        books.forEach(book => addBookToTable(book));

    } catch (error) {
        console.log("[-] Error:", error);
    }
}

startPageButton.addEventListener('click', event => {
    window.location.href = 'index.html';
})

adminPanelButton.addEventListener('click', event => {
    window.location.href = 'adminPanel.html';
})

searchForm.querySelector('#refresh').addEventListener('click', () => {
    document.querySelector('#search').value = '';
    hideWarning();
    loadBooksFromDB();
});

searchForm.addEventListener('submit', event => {

    event.preventDefault();
    hideWarning();

    try {

        const book = {
            title: document.querySelector('#search').value.trim()
        };

        if (book.title == '') {
            showWarning('Invalid book title!');
        }

        loadBooksWithTitle(book);
        
    } catch (error) {
        console.log("[-] Error: ", error);
    }
});

loadBooksFromDB();