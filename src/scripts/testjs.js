async function logincheck(username, password) {
    const URL = "http://localhost:5050/users/" + username;

    const response = await fetch(URL, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json();
    if (password === myJson.password) {
        location.replace("./main.html");
        localStorage.setItem("username", username);
    }
    else{
        alert("Username or password is incorect!");
    }

}

