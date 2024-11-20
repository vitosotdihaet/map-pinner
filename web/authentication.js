document.getElementById("loginButton").addEventListener("click", async () => {
    const userName = document.getElementById("userName").value;
    const password = document.getElementById("password").value;

    try {
        const response = await UserFetch.getByUsernamePassword(userName, password);

        if (response.id) {
            const data = await response.json();
            const token = data.token; // Assuming the server responds with a JWT token.

            // Store token in localStorage
            localStorage.setItem("jwt", token);

            // Redirect to map.html
            window.location.href = "static/map.html";
        } else {
            // Show error message
            document.getElementById("errorMessage").style.display = "block";
        }
    } catch (error) {
        console.error("Error during login:", error);
    }
});

document.getElementById("registerButton").addEventListener("click", async () => {
    const userName = document.getElementById("userName").value;
    const password = document.getElementById("password").value;

    try {
        const response = await UserFetch.createWithNamePassword(userName, password);

        if (response.ok) {
            const data = await response.json();
            const token = data.token; // Assuming the server responds with a JWT token.

            // Store token in localStorage
            localStorage.setItem("jwt", token);

            // Redirect to map.html
            window.location.href = "map.html";
        } else {
            // Show error message
            document.getElementById("errorMessage").style.display = "block";
        }
    } catch (error) {
        console.error("Error during registration:", error);
    }
});
