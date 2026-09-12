const statusDot = document.querySelector('#status-dot');
const statusDropdownMenuButton = document.querySelector('#status-username-button');
const statusDropdownMenu = document.querySelector('#status-username-dropdown');

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
    return payload.exp && payload.exp * 1000 <= Date.now();
}

const token = sessionStorage.getItem('token');
const isLoggedIn = token && !isTokenExpired(token);

function addDropdownButton(label, onClick) {
    const button = document.createElement('button');
    button.type = 'button';
    button.textContent = label;
    button.addEventListener('click', onClick);
    statusDropdownMenu.appendChild(button);
}

if (isLoggedIn) {
    statusDot.style.backgroundColor = 'green';
    statusDropdownMenuButton.textContent = sessionStorage.getItem('name');
    addDropdownButton('Log Out', () => {
        removeToken();
        window.location.reload();
    });

} else {
    removeToken();
    statusDot.style.backgroundColor = 'red';
    statusDropdownMenuButton.textContent = 'Not logged in';
    addDropdownButton('Log In', () => {
        window.location.href = 'login.html';
    });
    addDropdownButton('Sign Up', () => {
        window.location.href = 'signup.html';
    });
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
