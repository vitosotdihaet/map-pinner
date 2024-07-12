async function getData(url) {
    const response = await fetch(url)
    return await response.json()
}

async function postData(url, body) {
    const response = await fetch(url, {
        method: "POST",
        body: body,
    });
    return await response.json()
}

async function deleteData(url, body) {
    const response = await fetch(url, {
        method: "DELETE",
        body: body,
    });
    return await response.json()
}

async function putData(url, body) {
    const response = await fetch(url, {
        method: "PUT",
        body: body,
    });
    return await response.json()
}