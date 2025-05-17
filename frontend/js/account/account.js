document.addEventListener('DOMContentLoaded', function() {
    // Конфигурация
    const config = {
        dadataToken: "", // Замените на реальный ключ
        endpoints: {
            createOrder: '/account/CreateOrder',
            getUser: '/account/GetUser',
            getOrders: '/account/GetOrders',
            logout: '/account/logout',
            addProduct: '/store/AddProduct'
        }
    };

    // Состояние приложения
    const state = {
        cart: JSON.parse(localStorage.getItem('cart')) || [],
        selectedAddress: '',
        modal: null,
        isAdmin: false
    };

    // Инициализация
    function init() {
        checkRequiredElements();
        loadAccountData();
        setupEventListeners();
        setupCartControls();
        addModalStyles();
        setupAdminPanel();
    }

    // Проверка обязательных элементов DOM
    function checkRequiredElements() {
        const requiredIds = [
            'user-name', 'user-email', 'user-login',
            'orders-list', 'account-cart-items', 'account-cart-total',
            'logout-btn'
        ];

        const missing = requiredIds.filter(id => !document.getElementById(id));
        if (missing.length > 0) {
            console.error('Missing elements:', missing);
            showError('Ошибка загрузки страницы. Обновите страницу.');
            return false;
        }
        return true;
    }

    // Настройка обработчиков событий
    function setupEventListeners() {
        // Переключение вкладок
        document.querySelectorAll('[data-tab]').forEach(tab => {
            tab.addEventListener('click', handleTabSwitch);
        });

        // Кнопка оформления заказа
        document.querySelector('.checkout-btn')?.addEventListener('click', openOrderModal);

        // Выход из аккаунта
        document.getElementById('logout-btn')?.addEventListener('click', handleLogout);
    }

    // Управление корзиной
    function setupCartControls() {
        document.addEventListener('click', function(e) {
            const target = e.target;

            if (target.classList.contains('quantity-btn') && target.dataset.action === 'increase') {
                adjustQuantity(target.dataset.id, 1);
            }

            if (target.classList.contains('quantity-btn') && target.dataset.action === 'decrease') {
                adjustQuantity(target.dataset.id, -1);
            }

            if (target.classList.contains('remove-btn')) {
                removeFromCart(target.dataset.id);
            }
        });
    }

    // Инициализация админ-панели
    function setupAdminPanel() {
        const adminTab = document.querySelector('[data-tab="admin-panel"]');
        if (adminTab) {
            adminTab.addEventListener('click', function(e) {
                e.preventDefault();
                switchTab('admin-panel');
            });
        }

        const addProductForm = document.getElementById('add-product-form');
        if (addProductForm) {
            addProductForm.addEventListener('submit', async function(e) {
                e.preventDefault();
                await handleAddProduct();
            });

            // Превью изображения
            document.getElementById('product-image').addEventListener('change', function() {
                const preview = document.getElementById('image-preview');
                preview.innerHTML = '';

                if (this.files && this.files[0]) {
                    const reader = new FileReader();
                    reader.onload = function(e) {
                        const img = document.createElement('img');
                        img.src = e.target.result;
                        img.style.maxWidth = '200px';
                        img.style.maxHeight = '200px';
                        preview.appendChild(img);
                    };
                    reader.readAsDataURL(this.files[0]);
                }
            });
        }
    }

    // Обработка добавления товара
    async function handleAddProduct() {
        const submitBtn = document.querySelector('#add-product-form button[type="submit"]');
        submitBtn.disabled = true;

        try {
            const formData = new FormData();
            formData.append('name', document.getElementById('product-name').value);
            formData.append('price', document.getElementById('product-price').value);
            formData.append('category', document.getElementById('product-category').value);
            formData.append('description', document.getElementById('product-description').value);
            formData.append('amount', document.getElementById('product-amount').value)
            formData.append('image', document.getElementById('product-image').files[0]);

            const response = await fetch(config.endpoints.addProduct, {
                method: 'POST',
                credentials: 'include',
                body: formData
            });

            if (!response.ok) throw new Error('Ошибка добавления товара');

            const result = await response.json();
            if (!result.success) throw new Error(result.message || 'Неизвестная ошибка');

            showSuccess('Товар успешно добавлен!');
            document.getElementById('add-product-form').reset();
            document.getElementById('image-preview').innerHTML = '';
        } catch (error) {
            showError(error.message || 'Ошибка при добавлении товара');
        } finally {
            submitBtn.disabled = false;
        }
    }

    // Загрузка данных аккаунта
    async function loadAccountData() {
        try {
            showLoading();

            // Загрузка данных пользователя
            const userData = await fetchData(config.endpoints.getUser);
            updateUserInfo(userData);

            // Проверка роли администратора
            if (userData.role === 'admin') {
                state.isAdmin = true;
                document.getElementById('admin-menu-item').style.display = 'block';
            }

            // Загрузка истории заказов
            const ordersData = await fetchData(config.endpoints.getOrders);
            renderOrders(ordersData);

            // Обновление корзины
            updateAccountCart();

        } catch (error) {
            console.error('Error loading account data:', error);
            showError(error.message || 'Ошибка загрузки данных');
        } finally {
            hideLoading();
        }
    }

    // Модальное окно оформления заказа
    function openOrderModal() {
        if (state.cart.length === 0) {
            showError('Корзина пуста. Добавьте товары перед оформлением.');
            return;
        }

        if (!state.modal) createOrderModal();
        state.modal.style.display = 'flex';
        document.body.style.overflow = 'hidden';
    }

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
        state.modal = document.getElementById('order-modal');

        initAddressAutocomplete();
        setupModalEvents();
    }

    function initAddressAutocomplete() {
        const input = document.getElementById('address-input');
        const container = document.getElementById('address-suggestions');

        input.addEventListener('input', debounce(async () => {
            const query = input.value.trim();

            if (query.length < 3) {
                container.innerHTML = '';
                return;
            }

            try {
                const suggestions = await searchAddress(query);
                renderSuggestions(suggestions, container, input);
            } catch (error) {
                console.error('Address search error:', error);
                container.innerHTML = '<div class="error">Ошибка поиска адреса</div>';
            }
        }, 300));
    }

    async function searchAddress(query) {
        const response = await fetch('https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Authorization': 'Token ' + config.dadataToken
            },
            body: JSON.stringify({
                query: query,
                count: 5,
                locations: [{ country: 'Россия' }]
            })
        });

        if (!response.ok) throw new Error('Ошибка сервиса поиска адресов');
        return (await response.json()).suggestions;
    }

    function renderSuggestions(suggestions, container, input) {
        container.innerHTML = '';

        suggestions.forEach(item => {
            const div = document.createElement('div');
            div.className = 'suggestion-item';
            div.textContent = item.value;
            div.addEventListener('click', () => {
                input.value = item.value;
                state.selectedAddress = item.value;
                container.innerHTML = '';
            });
            container.appendChild(div);
        });
    }

    function setupModalEvents() {
        state.modal.querySelector('.close-modal').addEventListener('click', closeOrderModal);
        state.modal.querySelector('#confirm-order').addEventListener('click', submitOrder);
        state.modal.addEventListener('click', (e) => e.target === state.modal && closeOrderModal());
    }

    function closeOrderModal() {
        state.modal.style.display = 'none';
        document.body.style.overflow = 'auto';
    }

    async function submitOrder() {
        if (!state.selectedAddress) {
            showError('Пожалуйста, укажите адрес доставки');
            return;
        }

        try {
            showLoading();

            const orderData = {
                items: state.cart.map(item => ({
                    id: parseInt(item.id),
                    name: item.name,
                    price: item.price,
                    quantity: item.quantity,
                    imageUrl: item.image
                })),
                totalAmount: calculateTotal(),
                deliveryAddress: state.selectedAddress
            };

            const result = await fetchData(config.endpoints.createOrder, {
                method: 'POST',
                body: orderData
            });

            if (!result.success) throw new Error(result.message || 'Ошибка сервера');

            // Очистка после успешного заказа
            state.cart = [];
            localStorage.removeItem('cart');
            updateAccountCart();
            closeOrderModal();

            // Обновление интерфейса
            switchTab('orders');
            await loadAccountData();

            showSuccess('Заказ успешно оформлен!');

        } catch (error) {
            showError(error.message || 'Ошибка при оформлении заказа');
        } finally {
            hideLoading();
        }
    }

    // Вспомогательные функции
    function debounce(func, delay) {
        let timeout;
        return function() {
            clearTimeout(timeout);
            timeout = setTimeout(() => func.apply(this, arguments), delay);
        };
    }

    async function fetchData(endpoint, options = {}) {
        const defaultOptions = {
            credentials: 'include',
            headers: options.body instanceof FormData ? {} : { 'Content-Type': 'application/json' }
        };

        const response = await fetch(endpoint, {
            ...defaultOptions,
            ...options,
            body: options.body instanceof FormData ? options.body : JSON.stringify(options.body)
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        return await response.json();
    }

    function calculateTotal() {
        return state.cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
    }

    function adjustQuantity(itemId, delta) {
        const item = state.cart.find(i => i.id === itemId);
        if (item) {
            item.quantity = Math.max(1, item.quantity + delta);
            saveCart();
            updateAccountCart();
        }
    }

    function removeFromCart(itemId) {
        state.cart = state.cart.filter(i => i.id !== itemId);
        saveCart();
        updateAccountCart();
    }

    function saveCart() {
        localStorage.setItem('cart', JSON.stringify(state.cart));
    }

    function updateUserInfo(userData) {
        safeSetContent('user-name', userData.name);
        safeSetContent('user-email', userData.email);
        safeSetContent('user-login', userData.login);
    }

    function updateAccountCart() {
        const itemsContainer = document.getElementById('account-cart-items');
        const totalElement = document.getElementById('account-cart-total');

        if (!itemsContainer || !totalElement) return;

        itemsContainer.innerHTML = state.cart.length > 0
            ? state.cart.map(renderCartItem).join('')
            : '<p class="empty-message">Ваша корзина пуста</p>';

        totalElement.textContent = calculateTotal().toFixed(2);
    }

    function renderCartItem(item) {
        const total = (item.price * item.quantity).toFixed(2);
        return `
        <div class="cart-item">
            <img src="${item.image || '/img/no-image.png'}" alt="${item.name}">
            <div class="cart-item-info">
                <h4>${item.name}</h4>
                <p>${item.price.toFixed(2)} ₽ × ${item.quantity}</p>
                <div class="cart-item-controls">
                    <button class="quantity-btn" data-id="${item.id}" data-action="decrease">-</button>
                    <span>${item.quantity}</span>
                    <button class="quantity-btn" data-id="${item.id}" data-action="increase">+</button>
                    <button class="remove-btn" data-id="${item.id}">Удалить</button>
                </div>
            </div>
        </div>
        `;
    }

    function renderOrders(orders) {
        const container = document.getElementById('orders-list');
        if (!container) return;

        container.innerHTML = !orders || orders.length === 0
            ? '<p class="empty-message">У вас пока нет заказов</p>'
            : orders.map(renderOrder).join('');
    }

    function renderOrder(order) {
        const items = Array.isArray(order.items) ? order.items : [];
        return `
        <div class="order-item">
            <div class="order-header">
                <span>Заказ #${order.id || 'N/A'}</span>
                <span>${formatDate(order.orderDate)}</span>
                <span>${order.totalAmount?.toFixed(2) || '0.00'} ₽</span>
                <span class="status ${getStatusClass(order.status)}">${getStatusText(order.status)}</span>
            </div>
            ${items.length > 0 ? `
            <div class="order-products">
                ${items.map(renderOrderItem).join('')}
            </div>` : '<p>Нет информации о товарах</p>'}
            <div class="address">
                ${order.deliveryAddress}
            </div>
        </div>
        `;
    }

    function renderOrderItem(item) {
        const total = (item.quantity * item.price).toFixed(2);
        return `
        <div class="order-product">
            <img src="${item.imageUrl || '/img/no-image.png'}" alt="${item.name}" width="50">
            <div class="order-product-info">
                <span class="product-name">${item.name}</span>
                <span class="product-quantity">${item.quantity} × ${item.price?.toFixed(2) || '0.00'} ₽</span>
                <span class="product-total">${total} ₽</span>
            </div>
        </div>
        `;
    }

    async function handleLogout(e) {
        e.preventDefault();
        try {
            await fetch(config.endpoints.logout, {
                method: 'POST',
                credentials: 'include'
            });
            localStorage.removeItem('cart');
            window.location.href = '/';
        } catch (error) {
            showError('Ошибка при выходе из системы');
        }
    }

    function handleTabSwitch(e) {
        e.preventDefault();
        const tabId = this.getAttribute('data-tab');
        switchTab(tabId);
    }

    function switchTab(tabId) {
        // Переключение вкладок
        document.querySelectorAll('.tab-content').forEach(c => c.classList.remove('active'));
        document.getElementById(tabId)?.classList.add('active');

        // Обновление активного пункта меню
        document.querySelectorAll('.account-menu li').forEach(item => {
            item.classList.remove('active');
        });

        const activeItem = document.querySelector(`.account-menu li a[data-tab="${tabId}"]`);
        if (activeItem) {
            activeItem.parentElement.classList.add('active');
        }

        // Обновление корзины при необходимости
        if (tabId === 'cart') updateAccountCart();
    }

    function addModalStyles() {
        const style = document.createElement('style');
        style.textContent = `
        .modal-overlay {
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: rgba(0,0,0,0.5);
            display: none;
            justify-content: center;
            align-items: center;
            z-index: 1000;
        }
        .modal-content {
            background: white;
            padding: 2rem;
            border-radius: 8px;
            width: 90%;
            max-width: 500px;
            position: relative;
        }
        .close-modal {
            position: absolute;
            top: 1rem;
            right: 1rem;
            background: none;
            border: none;
            font-size: 1.5rem;
            cursor: pointer;
        }
        .suggestions-container {
            border: 1px solid #ddd;
            max-height: 200px;
            overflow-y: auto;
            margin-top: 0.5rem;
        }
        .suggestion-item {
            padding: 0.5rem;
            cursor: pointer;
        }
        .suggestion-item:hover {
            background: #f5f5f5;
        }
        .error {
            color: #ff4d4f;
            padding: 0.5rem;
        }
        #address-input {
            width:70%;
            height: 3.5rem;
        }
        #confirm-order {
            padding: 0.5rem 0.5rem;
        }
        /* Стили для админ-панели */
        #admin-panel {
            padding: 20px;
        }
        .admin-panel-content {
            max-width: 800px;
            margin: 0 auto;
        }
        #add-product-form {
            display: grid;
            gap: 1.5rem;
        }
        #add-product-form .form-group {
            margin-bottom: 1rem;
        }
        #add-product-form label {
            display: block;
            margin-bottom: 0.5rem;
            font-weight: 500;
        }
        #add-product-form input,
        #add-product-form select,
        #add-product-form textarea {
            width: 100%;
            padding: 0.8rem;
            border: 1px solid #ddd;
            border-radius: 4px;
            font-family: inherit;
        }
        #add-product-form textarea {
            min-height: 100px;
            resize: vertical;
        }
        #image-preview {
            margin-top: 1rem;
        }
        #image-preview img {
            border-radius: 4px;
            border: 1px solid #eee;
            max-width: 200px;
            max-height: 200px;
        }
        .error-message {
            display: flex;
            justify-content: space-between;
            align-items: center;
            color: #ff4d4f;
            padding: 0.5rem;
            background: #fff2f0;
            border: 1px solid #ffccc7;
            border-radius: 4px;
            margin: 0.5rem 0;
            transition: opacity 0.3s ease;
        }
        .close-error {
            background: none;
            border: none;
            font-size: 1.2rem;
            cursor: pointer;
            color: inherit;
            padding: 0 0 0 1rem;
        }
        .fade-out {
            opacity: 0;
        }
        `;
        document.head.appendChild(style);
    }

    function safeSetContent(elementId, value, defaultValue = 'Не указано') {
        const element = document.getElementById(elementId);
        if (element) element.textContent = value || defaultValue;
    }

    function formatDate(dateString) {
        try {
            return new Date(dateString).toLocaleDateString('ru-RU', {
                year: 'numeric',
                month: 'long',
                day: 'numeric'
            });
        } catch {
            return dateString || 'Дата не указана';
        }
    }

    function getStatusClass(status) {
        const map = {
            'completed': 'delivered',
            'delivered': 'delivered',
            'processing': 'processing',
            'shipped': 'processing',
            'cancelled': 'cancelled'
        };
        return map[status?.toLowerCase()] || 'unknown';
    }

    function getStatusText(status) {
        const map = {
            'completed': 'Доставлен',
            'delivered': 'Доставлен',
            'processing': 'В обработке',
            'shipped': 'Отправлен',
            'cancelled': 'Отменен'
        };
        return map[status?.toLowerCase()] || status || 'Статус неизвестен';
    }

    function showLoading() {
        document.getElementById('loading-overlay')?.style?.removeProperty('display');
    }

    function hideLoading() {
        const loader = document.getElementById('loading-overlay');
        if (loader) loader.style.display = 'none';
    }

    function showError(message) {
        const error = document.createElement('div');
        error.className = 'error-message';
        error.innerHTML = `
        <span>${message}</span>
        <button class="close-error">&times;</button>
    `;

        document.querySelectorAll('.tab-content').forEach(el => {
            if (el) {
                const errorClone = error.cloneNode(true);
                el.appendChild(errorClone);

                // Закрытие по кнопке
                errorClone.querySelector('.close-error').addEventListener('click', () => {
                    errorClone.classList.add('fade-out');
                    setTimeout(() => {
                        if (errorClone.parentNode) {
                            errorClone.parentNode.removeChild(errorClone);
                        }
                    }, 300);
                });

                // Автоматическое закрытие
                setTimeout(() => {
                    if (errorClone.parentNode) { // Проверяем, не удалено ли уже сообщение
                        errorClone.classList.add('fade-out');
                        setTimeout(() => {
                            if (errorClone.parentNode) {
                                errorClone.parentNode.removeChild(errorClone);
                            }
                        }, 300);
                    }
                }, 1000);
            }
        });
    }

    function showSuccess(message) {
        const toast = document.getElementById('toast');
        if (!toast) return;

        const toastMessage = toast.querySelector('.toast-message');
        toastMessage.textContent = message;
        toast.classList.add('active');

        setTimeout(() => {
            toast.classList.remove('active');
            setTimeout(() => {
                toast.querySelector('.toast-progress').style.animation = 'none';
                setTimeout(() => {
                    toast.querySelector('.toast-progress').style.animation = '';
                }, 10);
            }, 300);
        }, 3000);
    }

    // Запуск приложения
    init();
});