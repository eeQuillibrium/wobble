document.addEventListener('DOMContentLoaded', function() {
    // Проверка наличия необходимых DOM элементов
    const requiredElements = [
        'user-name', 'user-email', 'user-login',
        'orders-list', 'account-cart-items', 'account-cart-total',
        'logout-btn'
    ];

    const missingElements = requiredElements.filter(id => !document.getElementById(id));
    if (missingElements.length > 0) {
        console.error('Missing required elements:', missingElements);
        showError('Ошибка загрузки страницы. Пожалуйста, обновите страницу.');
        return;
    }

    // Инициализация корзины
    let cart = JSON.parse(localStorage.getItem('cart')) || [];

    // Добавляем в начало файла (после объявления cart)
    let modal = null;
    let selectedAddress = '';

    // Функция создания модального окна
    function createOrderModal() {
        const modalHTML = `
    <div class="modal-overlay" id="order-modal">
        <div class="modal-content">
            <button class="close-modal">&times;</button>
            <h2>Оформление заказа</h2>
            <div class="form-group">
                <label for="address-input">Адрес доставки</label>
                <input type="text" id="address-input" placeholder="Введите ваш адрес...">
                <div id="address-suggestions" class="suggestions-container"></div>
            </div>
            <button id="confirm-order" class="btn-primary">Подтвердить заказ</button>
        </div>
    </div>
    `;

        document.body.insertAdjacentHTML('beforeend', modalHTML);
        modal = document.getElementById('order-modal');

        // Инициализация поиска адреса (аналогично предыдущему примеру)
        initAddressAutocomplete();

        // Закрытие модального окна
        modal.querySelector('.close-modal').addEventListener('click', closeModal);
        modal.querySelector('#confirm-order').addEventListener('click', submitOrder);

        // Закрытие при клике вне окна
        modal.addEventListener('click', (e) => {
            if (e.target === modal) closeModal();
        });
    }

    function initAddressAutocomplete() {
        const addressInput = document.getElementById('address-input');
        const suggestionsContainer = document.getElementById('address-suggestions');

        addressInput.addEventListener('input', async function() {
            const query = this.value.trim();

            if (query.length < 3) {
                suggestionsContainer.innerHTML = '';
                return;
            }

            try {
                // Используем API DaData (замените YOUR_TOKEN)
                const response = await fetch('https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Accept': 'application/json',
                        'Authorization': 'Token YOUR_DADATA_API_KEY'
                    },
                    body: JSON.stringify({
                        query: query,
                        count: 5,
                        locations: [{ country: 'Россия' }]
                    })
                });

                const data = await response.json();
                suggestionsContainer.innerHTML = '';

                data.suggestions.forEach(suggestion => {
                    const div = document.createElement('div');
                    div.className = 'suggestion-item';
                    div.textContent = suggestion.value;
                    div.addEventListener('click', () => {
                        addressInput.value = suggestion.value;
                        selectedAddress = suggestion.value;
                        suggestionsContainer.innerHTML = '';
                    });
                    suggestionsContainer.appendChild(div);
                });
            } catch (error) {
                console.error('Address search error:', error);
            }
        });
    }

    function openModal() {
        if (!modal) createOrderModal();
        modal.style.display = 'flex';
        document.body.style.overflow = 'hidden';
    }

    function closeModal() {
        modal.style.display = 'none';
        document.body.style.overflow = 'auto';
    }

    async function submitOrder() {
        if (!selectedAddress) {
            showError('Пожалуйста, укажите адрес доставки');
            return;
        }

        try {
            showLoading();

            const orderItems = cart.map(item => ({
                id: parseInt(item.id),
                name: item.name,
                price: item.price,
                quantity: item.quantity,
                imageUrl: item.image
            }));

            const totalAmount = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);

            const response = await fetch('/CreateOrder', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    items: orderItems,
                    totalAmount: totalAmount,
                    deliveryAddress: selectedAddress
                })
            });

            if (!response.ok) throw new Error('Ошибка при оформлении заказа');

            const result = await response.json();
            if (!result.success) throw new Error(result.message || 'Ошибка сервера');

            // Очищаем корзину после успешного заказа
            cart = [];
            localStorage.removeItem('cart');
            updateAccountCart();
            closeModal();

            // Показываем уведомление об успехе
            alert('Заказ успешно оформлен!');

            // Переключаем на вкладку заказов
            document.querySelector('[data-tab="orders"]').click();
            await loadAccountData();

        } catch (error) {
            showError(error.message || 'Ошибка при оформлении заказа');
        } finally {
            hideLoading();
        }
    }


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
            const activeTab = document.getElementById(tabId);
            if (activeTab) {
                activeTab.classList.add('active');
            }

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
    const logoutBtn = document.getElementById('logout-btn');
    if (logoutBtn) {
        logoutBtn.addEventListener('click', function(e) {
            e.preventDefault();
            logout();
        });
    }

    // Функция загрузки данных аккаунта
    async function loadAccountData() {
        try {
            showLoading();

            // Загружаем данные пользователя
            const userResponse = await fetch('/account/GetUser', {
                credentials: 'include'
            });

            if (!userResponse.ok) {
                throw new Error('Ошибка загрузки данных пользователя');
            }

            const userData = await userResponse.json();
            console.log('User data received:', userData);

            // Заполняем данные пользователя
            safeSetContent('user-name', userData.name);
            safeSetContent('user-email', userData.email);
            safeSetContent('user-login', userData.login);

            // Загружаем историю заказов
            const ordersResponse = await fetch('/account/GetOrders', {
                credentials: 'include'
            });

            if (!ordersResponse.ok) {
                throw new Error('Ошибка загрузки истории заказов');
            }

            const ordersData = await ordersResponse.json();
            console.log('Orders data received:', ordersData);

            // Заполняем историю заказов
            renderOrders(ordersData);

            // Обновляем корзину
            updateAccountCart();

        } catch (error) {
            console.error('Error loading account data:', error);
            showError(error.message || 'Произошла ошибка при загрузке данных');
        } finally {
            hideLoading();
        }
    }

    // Функция безопасного установки содержимого элемента
    function safeSetContent(elementId, value, defaultValue = 'Не указано') {
        const element = document.getElementById(elementId);
        if (element) {
            element.textContent = value || defaultValue;
        } else {
            console.warn(`Element with id ${elementId} not found`);
        }
    }

    // Функция отображения заказов
    function renderOrders(orders) {
        const ordersList = document.getElementById('orders-list');
        if (!ordersList) return;

        if (!orders || !Array.isArray(orders) || orders.length === 0) {
            ordersList.innerHTML = '<p class="empty-message">У вас пока нет заказов</p>';
            return;
        }


        ordersList.innerHTML = orders.map(order => {
            // Проверяем наличие items и что это массив
            const items = Array.isArray(order.items) ? order.items : [];

            return `
                <div class="order-item">
                    <div class="order-header">
                        <span>Заказ #${order.id || 'N/A'}</span>
                        <span>${formatDate(order.orderDate)}</span>
                        <span>${order.totalAmount ? order.totalAmount.toFixed(2) : '0.00'} ₽</span>
                        <span class="status ${getStatusClass(order.status)}">${getStatusText(order.status)}</span>
                    </div>
                    ${items.length > 0 ? `
                    <div class="order-products">
                        ${items.map(item => `
                            <div class="order-product">
                                <img src="${item.imageUrl || '/img/no-image.png'}" alt="${item.name || 'Товар'}" width="50">
                                <div class="order-product-info">
                                    <span class="product-name">${item.name || 'Неизвестный товар'}</span>
                                    <span class="product-quantity">${item.quantity || 1} × ${item.price ? item.price.toFixed(2) : '0.00'} ₽</span>
                                    <span class="product-total">${(item.quantity * item.price).toFixed(2)} ₽</span>
                                </div>
                            </div>
                        `).join('')}
                    </div>
                    ` : '<p>Нет информации о товарах</p>'}
                </div>
            `;
        }).join('');
    }

    // Функция обновления корзины в аккаунте
    function updateAccountCart() {
        const cartItems = document.getElementById('account-cart-items');
        const cartTotal = document.getElementById('account-cart-total');

        if (!cartItems || !cartTotal) return;

        let total = 0;

        cartItems.innerHTML = cart.length > 0 ? cart.map(item => {
            const itemTotal = (item.price || 0) * (item.quantity || 1);
            total += itemTotal;

            return `
                <div class="cart-item">
                    <img src="${item.image || '/img/no-image.png'}" alt="${item.name || 'Товар'}">
                    <div class="cart-item-info">
                        <h4>${item.name || 'Неизвестный товар'}</h4>
                        <p>${(item.price || 0).toFixed(2)} ₽ × ${item.quantity || 1}</p>
                        <div class="cart-item-controls">
                            <button class="quantity-btn" data-id="${item.id}" data-action="decrease">-</button>
                            <span>${item.quantity || 1}</span>
                            <button class="quantity-btn" data-id="${item.id}" data-action="increase">+</button>
                            <button class="remove-btn" data-id="${item.id}">Удалить</button>
                        </div>
                    </div>
                </div>
            `;
        }).join('') : '<p class="empty-message">Ваша корзина пуста</p>';

        cartTotal.textContent = total.toFixed(2);

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
                        item.quantity = (item.quantity || 0) + 1;
                    } else if (action === 'decrease') {
                        item.quantity = Math.max(1, (item.quantity || 1) - 1);
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
    async function logout() {
        try {
            localStorage.removeItem('cart');

            const response = await fetch('/account/logout', {
                method: 'POST',
                credentials: 'include'
            });

            if (response.ok) {
                window.location.href = '/';
            } else {
                showError('Не удалось выйти. Попробуйте снова.');
            }
        } catch (error) {
            console.error('Logout error:', error);
            showError('Ошибка соединения. Попробуйте снова.');
        }
    }

    // Вспомогательные функции
    function formatDate(dateString) {
        try {
            const options = { year: 'numeric', month: 'long', day: 'numeric' };
            return new Date(dateString).toLocaleDateString('ru-RU', options);
        } catch (e) {
            console.error('Error formatting date:', e);
            return dateString || 'Дата не указана';
        }
    }

    function getStatusClass(status) {
        if (!status) return 'unknown';

        const statusMap = {
            'completed': 'delivered',
            'delivered': 'delivered',
            'processing': 'processing',
            'shipped': 'processing',
            'cancelled': 'cancelled'
        };
        return statusMap[status.toLowerCase()] || 'unknown';
    }

    function getStatusText(status) {
        if (!status) return 'Статус неизвестен';

        const statusTextMap = {
            'completed': 'Доставлен',
            'delivered': 'Доставлен',
            'processing': 'В обработке',
            'shipped': 'Отправлен',
            'cancelled': 'Отменен'
        };
        return statusTextMap[status.toLowerCase()] || status;
    }

    function showLoading() {
        const overlay = document.getElementById('loading-overlay');
        if (overlay) {
            overlay.style.display = 'flex'; // или 'block', если без flex
        }
    }
    function hideLoading() {
        const overlay = document.getElementById('loading-overlay');
        if (overlay) {
            overlay.style.display = 'none';
        }
    }

    function showError(message) {
        const errorElement = document.createElement('div');
        errorElement.className = 'error-message';
        errorElement.textContent = message;

        document.querySelectorAll('.tab-content').forEach(el => {
            if (el) {
                el.appendChild(errorElement.cloneNode(true));
            }
        });
    }

    // ===== Обработка оформления заказа =====
    const checkoutBtn = document.querySelector('.checkout-btn');
    if (checkoutBtn) {
        checkoutBtn.addEventListener('click', async function() {
            if (cart.length === 0) {
                showError('Корзина пуста. Добавьте товары перед оформлением заказа.');
                return;
            }

            try {
                showLoading();

                const orderItems = cart.map(item => ({
                    id: parseInt(item.id),       // Преобразуем string в int
                    name: item.name,
                    price: item.price,
                    quantity: item.quantity
                }));


                const response = await fetch('/account/CreateOrder', {
                    method: 'POST',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ items: orderItems })
                });

                if (!response.ok) {
                    throw new Error('Ошибка при оформлении заказа');
                }

                const result = await response.json();

                if (result.success) {
                    // Очищаем корзину после успешного оформления
                    cart = [];
                    localStorage.removeItem('cart');
                    updateAccountCart();  // Обновляем отображение корзины

                    // Переключаемся на вкладку "История заказов"
                    document.querySelector('[data-tab="orders"]').click();

                    // Обновляем список заказов
                    await loadAccountData();
                } else {
                    throw new Error(result.message || 'Неизвестная ошибка при оформлении заказа');
                }
            } catch (error) {
                console.error('Checkout error:', error);
                showError(error.message || 'Произошла ошибка при оформлении заказа');
            } finally {
                hideLoading();
            }
        });
    }

});