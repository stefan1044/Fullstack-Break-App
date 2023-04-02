function App() {
  

    const header = document.createElement('header')
    header.innerHTML = `
    <div id="LoginDiv">
        <label for="fname">Userame:</label>
        <input type="text" id="username" name="username"><br><br>
        <label for="lname">Password:</label>
        <input type="password" id="password" name="password"><br><br>
        <button id="loginbutton" onclick="logincheck(username.value, password.value)">Login</button>
    </div>
  `;

    return header.cloneNode(true);
}

export default App;