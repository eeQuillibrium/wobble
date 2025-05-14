document.addEventListener('DOMContentLoaded', function() {
    // Инициализация корзины
    let cart = JSON.parse(localStorage.getItem('cart')) || [];

    // Загрузка данных пользователя и заказов
    loadAccountData();

    // Переключение вкладок
    document.querySelectorAll('[data-tab]').forEach(tab => {
        tab.addEventListener('click', function(e) {
            e.preventDefault();
            const tabId = this.getAttribute('data-tab');

            // Скрыть все вкладки
            document.querySelectorAll('.tab-content').forEach(content => {
                content.classList.remove('active');
            });

            // Показать выбранную вкладку
            document.getElementById(tabId).classList.add('active');

            // Обновить активный пункт меню
            document.querySelectorAll('.account-menu li').forEach(item => {
                item.classList.remove('active');
            });
            this.parentElement.classList.add('active');

            // При переключении на корзину обновить её содержимое
            if (tabId === 'cart') {
                updateAccountCart();
            }
        });
    });

    // Обработка выхода
    document.getElementById('logout-btn').addEventListener('click', function(e) {
        e.preventDefault();
        logout();
    });

    // Функция загрузки данных аккаунта
    function loadAccountData() {
        // Здесь будет запрос к API для получения данных пользователя и заказов
        // Пока используем mock-данные
        const userData = {
            name: "Иван Иванов",
            email: "user@example.com"
        };

        const orders = [
            {
                id: 1001,
                date: "2023-05-15",
                total: 3500,
                status: "Доставлен",
                items: [
                    { name: "Лосось", price: 1500, quantity: 1, image: "/img/products/salmon.jpg" },
                    { name: "Форель", price: 1000, quantity: 2, image: "/img/products/trout.jpg" }
                ]
            },
            {
                id: 1002,
                date: "2023-06-20",
                total: 4200,
                status: "В обработке",
                items: [
                    { name: "Сёмга", price: 1800, quantity: 1, image: "/img/products/salmon2.jpg" },
                    { name: "Икра", price: 2400, quantity: 1, image: "/img/products/caviar.jpg" }
                ]
            }
        ];

        // Заполняем данные пользователя
        document.getElementById('user-name').textContent = userData.name;
        document.getElementById('user-email').textContent = userData.email;

        // Заполняем историю заказов
        renderOrders(orders);

        // Обновляем корзину
        updateAccountCart();
    }

    // Функция отображения заказов
    function renderOrders(orders) {
        const ordersList = document.getElementById('orders-list');

        ordersList.innerHTML = orders.map(order => `
            <div class="order-item">
                <div class="order-header">
                    <span>Заказ #${order.id}</span>
                    <span>${order.date}</span>
                    <span>${order.total} ₽</span>
                    <span class="status ${order.status.toLowerCase()}">${order.status}</span>
                </div>
                <div class="order-products">
                    ${order.items.map(item => `
                        <div class="order-product">
                            <img src="${item.image}" alt="${item.name}" width="50">
                            <span>${item.name} - ${item.price} ₽ × ${item.quantity}</span>
                        </div>
                    `).join('')}
                </div>
            </div>
        `).join('');
    }

    // Функция обновления корзины в аккаунте
    function updateAccountCart() {
        const cartItems = document.getElementById('account-cart-items');
        const cartTotal = document.getElementById('account-cart-total');

        let total = 0;

        cartItems.innerHTML = cart.map(item => {
            total += item.price * item.quantity;
            return `
                <div class="cart-item">
                    <img src="${item.image}" alt="${item.name}">
                    <div class="cart-item-info">
                        <h4>${item.name}</h4>
                        <p>${item.price} ₽ × ${item.quantity}</p>
                        <div class="cart-item-controls">
                            <button class="quantity-btn" data-id="${item.id}" data-action="decrease">-</button>
                            <span>${item.quantity}</span>
                            <button class="quantity-btn" data-id="${item.id}" data-action="increase">+</button>
                            <button class="remove-btn" data-id="${item.id}">Удалить</button>
                        </div>
                    </div>
                </div>
            `;
        }).join('') || '<p>Ваша корзина пуста</p>';

        cartTotal.textContent = total.toFixed(2);

        // Добавляем обработчики для кнопок корзины
        addCartEventListeners();
    }

    // Функция добавления обработчиков для корзины
    function addCartEventListeners() {
        document.querySelectorAll('.quantity-btn').forEach(btn => {
            btn.addEventListener('click', function() {
                const id = this.getAttribute('data-id');
                const action = this.getAttribute('data-action');
                const item = cart.find(i => i.id === id);

                if (item) {
                    if (action === 'increase') {
                        item.quantity++;
                    } else if (action === 'decrease') {
                        item.quantity--;
                        if (item.quantity < 1) item.quantity = 1;
                    }

                    localStorage.setItem('cart', JSON.stringify(cart));
                    updateAccountCart();
                }
            });
        });

        document.querySelectorAll('.remove-btn').forEach(btn => {
            btn.addEventListener('click', function() {
                const id = this.getAttribute('data-id');
                cart = cart.filter(i => i.id !== id);
                localStorage.setItem('cart', JSON.stringify(cart));
                updateAccountCart();
            });
        });
    }

    // Функция выхода
    function logout() {
        // Здесь будет запрос к API для выхода
        localStorage.removeItem('cart');
        window.location.href = '/';
    }
});