const taskInput = document.getElementById('new-task');
const addTaskButton = document.getElementById('add-task');
const taskList = document.getElementById('task-list');
const taskCount = document.getElementById('task-count');

let tasks = JSON.parse(localStorage.getItem('tasks')) || [];

function updateLocalStorage() {
    localStorage.setItem('tasks', JSON.stringify(tasks));
}

function renderTasks() {
    taskList.innerHTML = '';
    tasks.forEach((task, index) => {
        const taskItem = document.createElement('li');
        taskItem.className = `task-item ${task.completed ? 'Completed' : ''}`;

        taskItem.innerHTML = `
            <span>${task.name}</span>
            <div class="task-actions">
                <button class="complete" onclick="toggleComplete(${index})">&#10004;</button>
                <button class="edit" onclick="editTask(${index})">&#9998;</button>
                <button class="delete" onclick="deleteTask(${index})">&#10006;</button>
            </div>
        `;

        taskList.appendChild(taskItem);
    });
    updateTaskCount();
}

function updateTaskCount() {
    const pendingTasks = tasks.filter(task => !task.completed).length;
    taskCount.textContent = `Pending tasks: ${pendingTasks}`;
}

function addTask() {
    const taskName = taskInput.value.trim();
    if (taskName) {
        tasks.push({ name: taskName, completed: false });
        taskInput.value = '';
        updateLocalStorage();
        renderTasks();
    }
}

function editTask(index) {
    const newTaskName = prompt('Edit task:', tasks[index].name);
    if (newTaskName !== null) {
        tasks[index].name = newTaskName.trim();
        updateLocalStorage();
        renderTasks();
    }
}

function deleteTask(index) {
    tasks.splice(index, 1);
    updateLocalStorage();
    renderTasks();
}

function toggleComplete(index) {
    tasks[index].completed = !tasks[index].completed;
    updateLocalStorage();
    renderTasks();
}

addTaskButton.addEventListener('click', addTask);
taskInput.addEventListener('keypress', (e) => {
    if (e.key === 'Enter') {
        addTask();
    }
});

renderTasks();