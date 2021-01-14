import React, { useState, useEffect } from "react"
import * as api from "../services/api"

const todoCard = {

}

const container = {
    marginRight: 'auto', /* 1 */
    marginLeft: 'auto', /* 1 */

    maxWidth: '960px', /* 2 */

    paddingRight: '10px', /* 3 */
    paddingLeft: '10px' /* 3 */,
    boader: '1px',
}

// TODO action on click checkbox
// TODO Decorate
const TodoCard = ({ isDone, text, id }) => {
    return (
        <div id={id}>{isDone ? '/' : 'x'} : {`${text}`}</div>
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
            <TodoCard isDone={item.isDone} text={item.text} />);
        return cardList;
    }

    // TODO loading screen
    return (
        <>
            <title>todo</title>
            <h1>TODO</h1>
            <div style={container}>
                {
                    todoList && todoList.length > 0
                        ? renderTodoList(todoList)
                        : <div>Empty Data</div>
                }
            </div>
        </>
    )
}

export default Todo
