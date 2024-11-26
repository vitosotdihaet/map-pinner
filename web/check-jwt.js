notInAuth = !window.location.href.includes('auth')
notInMap = !window.location.href.includes('map')

const userToken = localStorage.getItem("jwt")

document.addEventListener("DOMContentLoaded", async () => {
    if (userToken === null && notInAuth) {
        window.location.href = "/static/auth.html"
        return
    } 

    try {
        const response = await UserFetch.validateToken()

        if (!response.ok && notInAuth) {
            window.location.href = "/static/auth.html"
        } else if (response.ok && notInMap) {
            window.location.href = "/static/map.html"
        }
    } catch (error) {
        console.error("Error:", error)
        window.location.href = "/static/auth.html"
    }
})
