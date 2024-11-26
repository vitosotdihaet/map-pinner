document.getElementById("loginButton").addEventListener("click", async () => {
    const userName = document.getElementById("userName").value
    const password = document.getElementById("password").value

    try {
        const response = await UserFetch.getByUsernamePassword(userName, password);

        if (response.token) {
            localStorage.setItem("jwt", response.token)
            window.location.href = "/static/map.html"
        } else {
            document.getElementById("errorMessage").style.display = "block"
        }
    } catch (error) {
        console.error("Error during login:", error)
    }
});

document.getElementById("registerButton").addEventListener("click", async () => {
    const userName = document.getElementById("userName").value
    const password = document.getElementById("password").value

    try {
        const response = await UserFetch.createWithNamePassword(userName, password);

        if (response.id) {
            location.reload()
        } else {
            document.getElementById("errorMessage").style.display = "block"
        }
    } catch (error) {
        console.error("Error during registration:", error)
    }
});
