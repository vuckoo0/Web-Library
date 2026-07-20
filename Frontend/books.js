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

loadBooksFromDB();