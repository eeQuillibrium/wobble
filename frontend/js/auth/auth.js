document.getElementById('loginForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('/account/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "login": username,
                "password": password
            }),
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.message || 'Login failed');
        }

        // Сохраняем токен в localStorage
        localStorage.setItem('jwtToken', data.token);

        // Перенаправляем на защищенную страницу
        window.location.href = '/';

    } catch (error) {
        alert(error.message);
        console.error('Error:', error);
    }
});