const addBookForm = document.querySelector('.book-form');
const editBookButton = document.querySelector('.edit-book button')
const editBookForm = document.querySelector('.edit-book-form');
const editBookSelector = document.querySelector('.book-field-selector');

async function saveBookToDB(book) {
    
    const response = await fetch('http://localhost:8080/books', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(book)
    });

    if (!response.ok) {
        throw new Error(`Server error: ${response.error}`)
    }

    const savedBook = await response.json();
    return savedBook;
}

async function editBookFromDB(newBook) {
    
    const response = await fetch(`http://localhost:8080/books?id=${newBook.id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            field: newBook.field,
            new_value: newBook.new_value
        })
    });

    if (!response.ok) {
        throw new Error(`Server error: ${response.error}`);
    }

    const editedBook = await response.json();
    return editedBook;
}

editBookButton.addEventListener('click', async (event) => {

    event.preventDefault();

    try {

        const bookChange = {
            id: document.querySelector('#change-book-title').value,
            field: document.querySelector('#field').value,
            new_value: document.querySelector('#field-value').value.trim()
        };

        if (bookChange.id <= 0) {
            throw new Error('Enter a valid book id!');
        }

        if (bookChange.new_value == '') {
            throw new Error('Enter a valid new value!');
        }

        const newBook = editBookFromDB(bookChange);

        if (!newBook.ok) {
            console.log(newBook.error);
        }

        editBookForm.reset();
        console.log(bookChange);
        alert('Field succsessfully changed!')

    } catch (error) {
        alert(`An error ocured: ${error}`);
        console.log(error);
    }
});

addBookForm.addEventListener('submit', async (event) => {

    event.preventDefault();

    try {

        const newBook = {
            title: document.querySelector('#title').value.trim(),
            author: document.querySelector('#author').value.trim(),
            isbn: document.querySelector('#isbn').value.trim()
        };

        if ((newBook.author == "") || (newBook.title == "") || (newBook.isbn == "")) {
            alert('You must enter a valid book!');
            return;
        }

        const savedBook = await saveBookToDB(newBook);

        alert('Book added!')
        addBookForm.reset();

    } catch (error) {
        console.log("[-] Error: ", error)
    }
});