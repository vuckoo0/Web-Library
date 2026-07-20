const form = document.querySelector('form')
console.log(form)

const bookTable = document.querySelector('.book-table tbody')
console.log(bookTable)

function addBookToTable(book) {
    const row = document.createElement('tr');

    row.innerHTML = `
        <td>${book.title}</td>
        <td>${book.author}</td>
        <td>${book.isbn}</td>
    `;

    bookTable.appendChild(row);
}

async function saveBookToDB(book) {
    
    const response = await fetch('http://192.168.1.237:8080/books', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(book)
    });

    const savedBook = await response.json();
    return savedBook;
}

async function loadBooksFromDB() {
    
    try {

        const response = await fetch('http://192.168.0.14:8080/books');
        const books = await response.json();

        bookTable.innerHTML = '';
        books.forEach(book => addBookToTable(book));

    } catch (error) {
        console.log("[-] Error:", error);
    }
}

form.addEventListener('submit', async (event) => {

    event.preventDefault();

    try {

        const savedBook = await saveBookToDB({
            title: document.querySelector('#title').value,
            author: document.querySelector('#author').value,
            isbn: document.querySelector('#isbn').value
        });

        addBookToTable(savedBook);
        form.reset();
    } catch (error) {
        console.log("[-] Error: ", error)
    }
});

loadBooksFromDB();