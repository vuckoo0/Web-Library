const booksButton = document.querySelector('#books-button');
const accountButton = document.querySelector('#account-button');

booksButton.addEventListener('click', event => {
    window.location.href = 'books.html';
})

accountButton.addEventListener('click', event => {
    window.location.href = 'login.html';
})