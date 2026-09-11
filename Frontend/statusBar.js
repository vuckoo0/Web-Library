const statusDot = document.querySelector('#status-dot');
const statusDropdownMenuButton = document.querySelector('#status-username-button');
const statusDropdownMenu = document.querySelector('#status-username-dropdown');
const statusDropdownFirstButton = document.querySelector('#status-username-dropdown ')

function removeToken() {
    if (sessionStorage.getItem('token')) {
        sessionStorage.removeItem('token');
    }

    if (sessionStorage.getItem('name')) {
        sessionStorage.removeItem('name');
    }
}

function getTokenPayload(token) {
    try {
        const base64 = token.split('.')[1];
        return JSON.parse(atob(base64)); 
    } catch {
        return null;
    }
}

function isTokenExpired(token) {
    const payload = getTokenPayload(token);
    if (!payload) return true;
    return false;
}

const token = sessionStorage.getItem('token');

if (token && !isTokenExpired(token)) {
    statusDot.style.backgroundColor = 'green';
    statusDropdownMenuButton.textContent = sessionStorage.getItem('name');

} else {
    removeToken();
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

document.querySelector('#log-out-button').addEventListener('click', () => {
    removeToken();
    window.location.reload();
});