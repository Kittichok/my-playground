import React, { useState, useEffect } from "react"
import * as api from "../services/api"

const TodoCard = ({ isDone, text, id, onChange }) => {
    return (
        <div
            className="p-4 max-w-xs mx-auto bg-white rounded-xl shadow-md m-1"
            key={id}
            id={id}>
            <label className="flex items-center space-x-3">
                <input type="checkbox"
                    checked={isDone}
                    className="appearance-none h-6 w-6 border border-gray-300 rounded-md checked:bg-blue-600 checked:border-transparent focus:outline-none"
                    onChange={onChange}
                />
                <span className="text-gray-900 font-medium">{`${text}`}</span>
            </label>
        </div>
    )
}

const Todo = () => {
    const [todoList, setTodoList] = useState([])

    const fetchData = async () => {
        const res = await api.getTodo();
        setTodoList(res.data.todoList)
    }

    useEffect(() => {
        fetchData()
    }, [])

    const renderTodoList = (list) => {
        const cardList = list.map(item =>
            <TodoCard isDone={item.isDone} text={item.text} id={item.id} key={item.id} onChange={
                (e) => {
                    api.updateTodo({ id: item.id, isDone: e.target.checked, text: item.text });
                    setTodoList(todoList.map(x => {
                        if (x.id === item.id) {
                            return {
                                text: item.text,
                                isDone: e.target.checked,
                                id: item.id,
                            }
                        }
                        return x
                    }))
                }

            } />);
        return cardList;
    }

    const CreateTodo = () => {
        return (
            <div>
                <form className="flex">
                    <input className="flex-1 rounded-l-lg p-4 border-t mr-0 border-b border-l text-gray-800 border-gray-200 bg-white" placeholder="i want todo" type="text" name="text" id="todo-input" />
                    <button
                        className="px-6 rounded-r-lg bg-green-400 text-gray-800 font-bold p-4 uppercase border-green-500 border-t border-b border-r"
                        onClick={async () => {
                            const text = document.getElementById('todo-input').value
                            api.createTodo(text).then((resp) => {
                                const addedList = Object.assign([], todoList)
                                addedList.push(resp.data.createTodo)
                                setTodoList(addedList)
                            })
                        }}>Add</button>
                </form>
            </div>
        )
    }

    // TODO loading screen
    return (
        <>
            <title>todo</title>
            <div className="p-4 max-w-xs mx-auto bg-white rounded-xl shadow-md m-3">
                <div className="mx-auto w-1/4 p-2">
                    <h1 >TODO</h1>
                </div>
                <CreateTodo />
                {
                    todoList && todoList.length > 0
                        ? renderTodoList(todoList)
                        : <div>Empty</div>
                }
            </div>
        </>
    )
}

export default Todo
