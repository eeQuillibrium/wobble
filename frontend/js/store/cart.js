document.addEventListener('DOMContentLoaded', function() {
    let cart = JSON.parse(localStorage.getItem('cart')) || [];
    updateCartDisplay();

    // Добавление товара с анимацией
    document.querySelectorAll('.add-to-cart').forEach(button => {
        button.addEventListener('click', function() {
            const product = {
                id: this.dataset.id,
                name: this.dataset.name,
                price: parseFloat(this.dataset.price),
                image: this.dataset.image
            };

            animateToCart(this, product.image);
            addToCart(product);
            updateCartDisplay();
        });
    });

    // Анимация летящего элемента
    function animateToCart(button, imageUrl) {
        const flyItem = document.getElementById('fly-item');
        const buttonRect = button.getBoundingClientRect();
        const cartRect = document.getElementById('cart-button').getBoundingClientRect();

        flyItem.style.cssText = `
            display: block;
            left: ${buttonRect.left}px;
            top: ${buttonRect.top}px;
            background-image: url(${imageUrl});
        `;

        requestAnimationFrame(() => {
            flyItem.style.transform = `translate(
                ${cartRect.left - buttonRect.left}px,
                ${cartRect.top - buttonRect.top}px
            ) scale(0.2)`;
            flyItem.style.opacity = '0';
        });

        setTimeout(() => {
            flyItem.removeAttribute('style');
        }, 500);
    }

    // Логика корзины
    function addToCart(product) {
        const existing = cart.find(item => item.id === product.id);
        existing ? existing.quantity++ : cart.push({...product, quantity: 1});
        localStorage.setItem('cart', JSON.stringify(cart));
    }

    function updateCartDisplay() {
        // Обновление счетчика и суммы
        const totalCount = cart.reduce((sum, item) => sum + item.quantity, 0);
        const totalPrice = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);

        document.getElementById('cart-count').textContent = totalCount;
        document.getElementById('cart-total').textContent = totalPrice.toFixed(2);

        // Отрисовка товаров в корзине
        const cartItems = document.getElementById('cart-items');
        cartItems.innerHTML = cart.map(item => `
            <div class="cart-item">
                <div class="cart-item-info">
                    <div class="cart-item-name">${item.name}</div>
                    <div class="cart-item-price">${item.price}₽ × ${item.quantity}</div>
                </div>
                <div class="cart-item-controls">
                    <button class="quantity-btn minus" data-id="${item.id}">-</button>
                    <button class="quantity-btn plus" data-id="${item.id}">+</button>
                    <button class="remove-item" data-id="${item.id}">×</button>
                </div>
            </div>
        `).join('');

        // Обработчики для управления количеством
        document.querySelectorAll('.quantity-btn').forEach(btn => {
            btn.addEventListener('click', function() {
                const item = cart.find(i => i.id === this.dataset.id);
                this.classList.contains('plus') ? item.quantity++ : item.quantity--;
                if(item.quantity < 1) item.quantity = 1;
                localStorage.setItem('cart', JSON.stringify(cart));
                updateCartDisplay();
            });
        });

        // Обработчики удаления
        document.querySelectorAll('.remove-item').forEach(btn => {
            btn.addEventListener('click', function() {
                cart = cart.filter(i => i.id !== this.dataset.id);
                localStorage.setItem('cart', JSON.stringify(cart));
                updateCartDisplay();
            });
        });
    }
});