import { getUser } from "./auth"

export const getTodo = async () => {
    const body = {
        "query": "{ todoList { text, isDone, id }}"
    }
    const url = process.env.GATSBY_TODO_DOMAIN
    const resultData = await callAPI({ url, body })
    return resultData
}

export const createTodo = async (text) => {
    const body = {
        "query": `mutation { createTodo(text: "${text}") { text, isDone, id }}`
    }
    const url = process.env.GATSBY_TODO_DOMAIN
    const resultData = await callAPI({ url, body })
    return resultData
}

export const updateTodo = async ({ id, isDone, text }) => {
    const body = {
        "query": `mutation { updateTodo(id: ${id}, data: { text: "${text}", isDone: ${isDone} }) { text, isDone, id }}`
    }
    const url = process.env.GATSBY_TODO_DOMAIN
    const resultData = await callAPI({ url, body })
    return resultData
}

// export const deleteTodo = async (id) => {
// }

export const login = async ({ username, password }) => {
    const body = {
        username,
        password
    }
    const url = `${process.env.GATSBY_AUTH_DOMAIN}v1/api/signin`
    const resultData = await callAPI({ url, body })
    return resultData
}

const callAPI = async ({ url, body }) => {
    let method = 'GET'
    if (body) {
        method = 'POST'
    }
    const headers = {
        'Content-Type': 'application/json'
    }
    const user = getUser();
    if (!!user.token) {
        headers["Authorization"] = `Bearer ${user.token}`
    }
    const result = await fetch(url, {
        method,
        mode: 'cors',
        headers,
        body: JSON.stringify(body),
    })
    return await result.json()
}