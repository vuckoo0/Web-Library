const token = localStorage.getItem('token');
const username = localStorage.getItem('name');

const statusDot = document.querySelector('#status-dot');
const statusUsername = document.querySelector('#status-username');

if (token) {
    statusDot.style.backgroundColor = 'green';
    statusUsername.textContent = username;
} else {
    statusDot.style.backgroundColor = 'red';
    statusUsername.textContent = 'Not logged in';
}