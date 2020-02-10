const prettify = json => JSON.stringify(json, null, 2);
                const logJson = json => console.log(prettify(json));

                const selectById = id => document.getElementById(id);
                const selectAll = cssSelector => [].slice.call(document.querySelectorAll(cssSelector));
                const setText = (id, text) => selectById(id).innerHTML = text;


                selectById('todoId').value = 0;


                const rest = async (url, options) => {
                        const response = await fetch(url, options);
                        return await response.json();
                }

                const createTodoHtml = todoJson => {
                        const todoClass = todoJson.priority == 1 ? 'priority-1-todo'
                                :  todoJson.priority == 2 ? 'priority-2-todo'
                                : 'priority-3-todo';

                        return '<div class="' + todoClass + '">'
                                + '<div class="name">' + todoJson.name + '</div>'
                                + '<div>' + todoJson.description + '</div>'
                                + '<div>' + 'ID: ' + todoJson.id + '</div>'
                                + '<div>' + todoJson.date  + ' ' + todoJson.time + '</div>'
                                + '</div>';
                }

                const isTodoValid = () => {
                        if (selectById('todoName').value === ""){
                                alert('Name field is mandatory.');
                                return false;
                        }
                        return true;

                }

		const getTodoJson = () => {

                        const todoJson = selectAll('#form input, #form select').reduce(
                                (acc, input) => {
                                        // id = todoName => key = name
                                        const key = input.id.replace('todo','').toLowerCase();
                                        acc[key] = [ 'id', 'priority' ].includes(key) ?
                                                parseInt(input.value) :
                                                input.value;

                                        return acc;
                                },
                                {}
                        );
                        return JSON.stringify(todoJson);
                };

                const resetTodos = () => selectById('todos').innerHTML = '';

                const apiUrl = 'https://go-gin-todo.herokuapp.com/api/todo';
                //const apiUrl = 'http://localhost:3000/api/todo';
                const actions = {
                        getAllTodos: async () => {
                                const responseJson =  await rest(apiUrl);
                                resetTodos();

                                responseJson.data.map(createTodoHtml)
                                        .forEach(todoHtml => todos.innerHTML = todos.innerHTML + todoHtml);

                                return responseJson.data;
                        },

                        getTodo: async todoId => {
                                const responseJson =  await rest(apiUrl + '/' + todoId);

                                return responseJson.data;
                        },

                        createTodo: async todoJson => {
                                if (!isTodoValid()){
                                        return;
                                }
                                //console.log('body:  ' + getTodoJson());
                                const responseJson =  await rest(
                                        apiUrl,
                                        {
                                                method: 'POST',
                                                headers: { 'Content-Type': 'application/json' },
                                                body: getTodoJson()
                                        }
                                );
                                actions.getAllTodos();
                                return responseJson.data;
                        },

			updateTodo: async (todoId, todoJson) => {
                                if (!isTodoValid()){
                                        return;
                                }
                                const responseJson =  await rest(
                                        apiUrl + '/' + todoId,
                                        {
                                                method: 'PUT',
                                                headers: { 'Content-Type': 'application/json' },
                                                body: getTodoJson()
                                        }
                                );
                                actions.getAllTodos();
                                return responseJson.data;
                        },

                        deleteTodo: async todoId => {
                                const responseJson =  await rest(
                                        apiUrl + '/' + todoId,
                                        { method: 'DELETE' }
                                );
                                actions.getAllTodos();
                                return responseJson;
                        },

                        deleteAllTodos: async () => {
                                const responseJson =  await rest(
                                        apiUrl + '/reset',
                                        { method: 'POST' }
                                );
                                resetTodos();
                                return responseJson;
                        }
                };

                const addTodoActionToButton = buttonElement => {
                        console.log(todo[button.id]);
                        buttonElement.addEventListener('click', todo[button.id]);
                }

                const getTodoJsonBody = () => JSON.parse(selectById('todoBody').value);
                const getTodoId = () => selectById('todoId').value;

		const buttonsClickEventHandlers = {
                        getAllTodos: actions.getAllTodos,
                        getTodo: () => actions.getTodo(getTodoId()),
                        createTodo: () => actions.createTodo(getTodoJson()),
                        updateTodo: () => actions.updateTodo(getTodoId(), getTodoJson()),
                        deleteTodo: () => actions.deleteTodo(getTodoId()),
                        deleteAllTodos: actions.deleteAllTodos
                }

                const bindButtonsActionClickHandlers = () => selectAll('#actions button')
                        .forEach(button => button.addEventListener(
                                'click',
                                buttonsClickEventHandlers[button.id]
                        )
                );

                bindButtonsActionClickHandlers();

