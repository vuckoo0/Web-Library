const booksButton = document.querySelector('#books-button');
const accountButton = document.querySelector('#account-button');

booksButton.addEventListener('click', event => {
    window.location.href = 'books.html';
})

accountButton.addEventListener('click', event => {
    window.location.href = 'login.html';
})

if (sessionStorage.getItem('isRunning') != '1') {

    if (sessionStorage.getItem('token')) {
        sessionStorage.removeItem('token');
    }

    if (sessionStorage.getItem('name')) {
        sessionStorage.removeItem('name');
    }

    sessionStorage.setItem('isRunning', '1');
    console.log('Token deleted from local storage!');
}