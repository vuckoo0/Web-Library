const token = localStorage.getItem('token');
const username = localStorage.getItem('name');
const statusDot = document.querySelector('#status-dot');
const statusDropdownMenuButton = document.querySelector('#status-username-button');
const statusDropdownMenu = document.querySelector('#status-username-dropdown');

if (token) {
    statusDot.style.backgroundColor = 'green';
    statusDropdownMenuButton.textContent = username;
} else {
    statusDot.style.backgroundColor = 'red';
    statusDropdownMenuButton.textContent = 'Not logged in';
}

statusDropdownMenuButton.addEventListener('click', () => {
    const isOpen = statusDropdownMenu.style.display === 'block';
    statusDropdownMenu.style.display = isOpen ? 'none' : 'block';
});

document.addEventListener('click', (event) => {
    if (!document.querySelector('#status-username').contains(event.target)) {
        statusDropdownMenu.style.display = 'none';
    }
});