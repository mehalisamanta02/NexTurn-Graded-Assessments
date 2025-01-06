const amountInput = document.getElementById('amount');
const descriptionInput = document.getElementById('description');
const categoryInput = document.getElementById('category');
const addExpenseButton = document.getElementById('add-expense');
const expenseTableBody = document.querySelector('#expense-table tbody');
const categorySummary = document.getElementById('category-summary');
const expenseChart = document.getElementById('expense-chart');

let expenses = JSON.parse(localStorage.getItem('expenses')) || [];

function updateLocalStorage() {
    localStorage.setItem('expenses', JSON.stringify(expenses));
}

function renderExpenses() {
    expenseTableBody.innerHTML = '';
    expenses.forEach((expense, index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>Rs.${expense.amount}</td>
            <td>${expense.description}</td>
            <td><span class="category-badge ${expense.category}">${expense.category}</span></td>
            <td><button onclick="deleteExpense(${index})">Delete</button></td>
        `;
        expenseTableBody.appendChild(row);
    });
    renderCategorySummary();
    renderChart();
}

function renderCategorySummary() {
    const categoryTotals = expenses.reduce((totals, expense) => {
        totals[expense.category] = (totals[expense.category] || 0) + expense.amount;
        return totals;
    }, {});

    categorySummary.innerHTML = Object.entries(categoryTotals)
        .map(([category, total]) => `${category}: Rs.${total.toFixed(2)}`)
        .join(' | ');
}

function renderChart() {
    const ctx = expenseChart.getContext('2d');
    const categoryTotals = expenses.reduce((totals, expense) => {
        totals[expense.category] = (totals[expense.category] || 0) + expense.amount;
        return totals;
    }, {});

    const categories = Object.keys(categoryTotals);
    const totals = Object.values(categoryTotals);

    if (window.myChart) {
        window.myChart.destroy();
    }

    window.myChart = new Chart(ctx, {
        type: 'pie',
        data: {
            labels: categories,
            datasets: [{
                data: totals,
                backgroundColor: ['blue', 'yellow', 'green', 'grey'],
            }],
        },
    });
}

function addExpense() {
    const amount = parseFloat(amountInput.value);
    const description = descriptionInput.value.trim();
    const category = categoryInput.value;

    if (!amount || !description) {
        alert('Please enter a valid amount and description!');
        return;
    }

    expenses.push({ amount, description, category });
    updateLocalStorage();
    renderExpenses();

    amountInput.value = '';
    descriptionInput.value = '';
    categoryInput.value = 'Food';
}

function deleteExpense(index) {
    expenses.splice(index, 1);
    updateLocalStorage();
    renderExpenses();
}

addExpenseButton.addEventListener('click', addExpense);
document.addEventListener('DOMContentLoaded', renderExpenses);
