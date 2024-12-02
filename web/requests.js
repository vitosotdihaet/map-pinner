async function getJSON(url) {
    const response = await getFetch(url)
    if (!response.ok) { throw "invalid get request" }
    return response.json()
}

async function postJSON(url, body) {
    const response = await postFetch(url, body)
    if (!response.ok) { throw "invalid post request" }
    return response.json()
}

async function deleteJSON(url, body) {
    const response = await deleteFetch(url, body)
    if (!response.ok) { throw "invalid delete request" }
    return response.json()
}

async function putJSON(url, body) {
    const response = await putFetch(url, body)
    if (!response.ok) { throw "invalid put request" }
    return response.json()
}

async function getFetch(url) {
    return await fetch(url, {
        headers: { "Authorization": `Bearer ${userToken}` },
    })
}

async function postFetch(url, body) {
    return await fetch(url, {
        method: "POST",
        body: body,
        headers: { "Authorization": `Bearer ${userToken}` },
    })
}

async function deleteFetch(url, body) {
    return await fetch(url, {
        method: "DELETE",
        body: body,
        headers: { "Authorization": `Bearer ${userToken}` },
    })
}

async function putFetch(url, body) {
    return await fetch(url, {
        method: "PUT",
        body: body,
        headers: { "Authorization": `Bearer ${userToken}` },
    })
}