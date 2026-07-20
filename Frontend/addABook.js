const form = document.querySelector('form')
console.log(form)

async function saveBookToDB(book) {
    
    const response = await fetch('http://localhost:8080/books', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(book)
    });

    const savedBook = await response.json();
    return savedBook;
}

form.addEventListener('submit', async (event) => {

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

        form.reset();
    } catch (error) {
        console.log("[-] Error: ", error)
    }
});