document.addEventListener('DOMContentLoaded', function() {
    // Переключение видимости пароля
    document.querySelectorAll('.password-toggle').forEach(toggle => {
        toggle.addEventListener('click', function() {
            const input = this.previousElementSibling;
            const type = input.getAttribute('type') === 'password' ? 'text' : 'password';
            input.setAttribute('type', type);
            this.classList.toggle('fa-eye-slash');
        });
    });

    // Валидация формы регистрации
    const registerForm = document.getElementById('registerForm');
    if (registerForm) {
        registerForm.addEventListener('submit', async function(e) {
            e.preventDefault();

            // Получаем значения полей
            const name = document.getElementById('name').value.trim();
            const login = document.getElementById('login').value
            const email = document.getElementById('email').value.trim();
            const password = document.getElementById('password').value;
            const confirmPassword = document.getElementById('confirmPassword').value;

            // Валидация
            if (!validateRegistration(name, email, password, confirmPassword)) {
                return;
            }

            try {
                // Отправка данных на сервер
                const response = await fetch('/account/register', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        "login": login,
                        "name": name,
                        "email": email,
                        "password": password
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || 'Ошибка регистрации');
                }

                const user = await response.json();

                // Перенос гостевой корзины
                mergeGuestCartWithUserCart(user.id);

                // Перенаправление
                window.location.assign('/account/auth');
            } catch (error) {
                showError(error.message);
            }
        });
    }
});

// Функция валидации
function validateRegistration(name, email, password, confirmPassword) {
    let isValid = true;

    // Очистка предыдущих ошибок
    document.querySelectorAll('.error-message').forEach(el => {
        el.style.display = 'none';
    });
    document.querySelectorAll('.input-error').forEach(el => {
        el.classList.remove('input-error');
    });

    // Проверка имени
    if (name.length < 2) {
        showFieldError('name', 'Имя должно содержать минимум 2 символа');
        isValid = false;
    }

    // Проверка email
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
        showFieldError('email', 'Введите корректный email');
        isValid = false;
    }

    // Проверка пароля
    if (password.length < 6) {
        showFieldError('password', 'Пароль должен содержать минимум 6 символов');
        isValid = false;
    }

    // Проверка подтверждения пароля
    if (password !== confirmPassword) {
        showFieldError('confirmPassword', 'Пароли не совпадают');
        isValid = false;
    }

    return isValid;
}

// Показать ошибку для конкретного поля
function showFieldError(fieldId, message) {
    const field = document.getElementById(fieldId);
    field.classList.add('input-error');

    let errorElement = field.nextElementSibling;
    if (!errorElement || !errorElement.classList.contains('error-message')) {
        errorElement = document.createElement('div');
        errorElement.className = 'error-message';
        field.parentNode.insertBefore(errorElement, field.nextSibling);
    }

    errorElement.textContent = message;
    errorElement.style.display = 'block';
}

// Показать общую ошибку
function showError(message) {
    const errorElement = document.createElement('div');
    errorElement.className = 'error-message';
    errorElement.textContent = message;
    errorElement.style.display = 'block';
    errorElement.style.textAlign = 'center';
    errorElement.style.margin = '1rem 0';

    const form = document.getElementById('registerForm') || document.getElementById('loginForm');
    form.prepend(errorElement);

    setTimeout(() => {
        errorElement.style.opacity = '0';
        setTimeout(() => errorElement.remove(), 300);
    }, 5000);
}

// Объединение корзин (как в предыдущем примере)
function mergeGuestCartWithUserCart(userId) {
    const guestCart = JSON.parse(localStorage.getItem('guest_cart')) || [];
    const userCartKey = `cart_${userId}`;
    const userCart = JSON.parse(localStorage.getItem(userCartKey)) || [];

    const mergedCart = [...userCart];

    guestCart.forEach(guestItem => {
        const existingItem = mergedCart.find(item => item.id === guestItem.id);
        if (existingItem) {
            existingItem.quantity += guestItem.quantity;
        } else {
            mergedCart.push(guestItem);
        }
    });

    localStorage.setItem(userCartKey, JSON.stringify(mergedCart));
    localStorage.removeItem('guest_cart');
}