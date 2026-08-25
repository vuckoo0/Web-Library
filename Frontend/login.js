const startMenuButton = document.querySelector('#start-page-button');
const signUpForm = document.querySelector('#sign-up-form');

startMenuButton.addEventListener('click', event => {
    window.location.href = 'index.html';
});

signUpForm.addEventListener('submit', async event => {

    event.preventDefault();

    try {

        const newUser = {
            name: document.querySelector('sign-up-user-name').value.trim(),
            password: document.querySelector('sign-up-password').value.trim()
        };

        if (newUser.name == '') {
            alert('Enter a valid username!');
        }

        if (newUser.password == '') {
            allert('Enter a valid password!');
        }

        const response = await fetch("http://localhost:8080/register", {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json'
            },
            body: JSON.stringify(newUser)
        });

        if (!response.ok) {
            throw new Error(`Server error: ${response.error}`)
        }

        const addedUser = await response.json();

        alert('Sign Up succsessful')
        signUpForm.reset();
        
    } catch (error) {
        alert(`An error ocured!: ${error}`)
        console.log(error)
    }
});