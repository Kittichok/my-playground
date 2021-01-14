

export const getTodo = async () => {
    const url = `http://localhost:4000/graphql`
    const body = {
        "query": "{ todoList { text, isDone }}"
    }
    const result = await fetch(url, {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body),
    })
    const resultData = await result.json()
    return resultData
}