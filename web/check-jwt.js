inAuth = window.location.href.includes('auth.html')

const userToken = localStorage.getItem("jwt")

document.addEventListener("DOMContentLoaded", async () => {
    if (userToken === null && !inAuth) {
        window.location.href = "/static/auth.html"
        return
    }

    try {
        response = await UserFetch.getCurrent()

        if (!response.ok && !inAuth) {
            window.location.href = "/static/auth.html"
        }
        
        if (response.ok) {
            response = await response.json()
            localStorage.setItem("jwt", response.token)
            localStorage.setItem("user", JSON.stringify(response.user))

            if (inAuth) {
                window.location.href = "/static/group.html"
            }
        }
    } catch (error) {
        console.error("Error:", error)
        window.location.href = "/static/auth.html"
    }
})
