document.addEventListener('DOMContentLoaded', function() {
    const products = document.querySelectorAll('.product-card');
    const filterButtons = document.querySelectorAll('.filter-btn');

    filterButtons.forEach(button => {
        button.addEventListener('click', function() {
            // Активация кнопки фильтра
            filterButtons.forEach(b => b.classList.remove('active'));
            this.classList.add('active');

            // Фильтрация товаров
            const category = this.dataset.category;
            products.forEach(product => {
                product.style.display = category === 'all' ||
                product.dataset.category === category ? 'block' : 'none';
            });
        });
    });
});