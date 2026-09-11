const startMenuButton = document.querySelector('#start-page-button');
const signUpForm = document.querySelector('#sign-up-form');
const logInForm = document.querySelector('#log-in-form');

const API_URL = `http://${window.location.hostname}:8080`;

startMenuButton.addEventListener('click', event => {
    window.location.href = 'index.html';
});

signUpForm.addEventListener('submit', async event => {

    event.preventDefault();

    try {

        const newUser = {
            name: document.querySelector('#sign-up-user-name').value.trim(),
            password: document.querySelector('#sign-up-password').value.trim()
        };

        if (newUser.name == '') {
            throw new Error('Enter a valid username!');
        }

        if (newUser.password == '') {
            throw new Error('Enter a valid password!');
        }

        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json'
            },
            body: JSON.stringify(newUser)
        });

        if (!response.ok) {
            throw new Error(`Server error: ${response.error}`)
        }

        alert('Sign Up succsessful')
        window.location.reload();
        
    } catch (error) {
        alert(`An error ocured!: ${error}`)
        console.log(error)
        signUpForm.reset();
    }
});

logInForm.addEventListener('submit', async event => {

    event.preventDefault();

    try {

        const loggingUser = {
            name: document.querySelector('#log-in-user-name').value.trim(),
            password: document.querySelector('#log-in-user-password').value.trim()
        };

        if (loggingUser.name == '') {
            throw new Error('Enter a valid username!');
        }

        if (loggingUser.password == '') {
            throw new Error('Enter a valid password!');
        }

        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(loggingUser)
        });

        if (!response.ok) {
            throw new Error(`Server error: ${response.error}`)
        }

        const data = await response.json();
        sessionStorage.setItem('token', data.token);
        sessionStorage.setItem('name', data.name);
        window.location.reload();

    } catch (error) {
        alert(`An error ocured!: ${error}`);
        logInForm.reset();
    }
});