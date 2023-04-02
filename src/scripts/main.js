// ---------horizontal-navbar-menu-----------
var tabsNewAnim = $('#navbar-animmenu');
var selectorNewAnim = $('#navbar-animmenu').find('li').length;
//var selectorNewAnim = $(".tabs").find(".selector");
var activeItemNewAnim = tabsNewAnim.find('.active');
var activeWidthNewAnimWidth = activeItemNewAnim.innerWidth();
var itemPosNewAnimLeft = activeItemNewAnim.position();
$(".hori-selector").css({
    "left": itemPosNewAnimLeft.left + "px",
    "width": activeWidthNewAnimWidth + "px"
});
$("#navbar-animmenu").on("click", "li", function (e) {
    $('#navbar-animmenu ul li').removeClass("active");
    $(this).addClass('active');
    var activeWidthNewAnimWidth = $(this).innerWidth();
    var itemPosNewAnimLeft = $(this).position();
    $(".hori-selector").css({
        "left": itemPosNewAnimLeft.left + "px",
        "width": activeWidthNewAnimWidth + "px"
    });
});

async function getprofile() {

    const URL = "http://localhost:5050/profiles/" + localStorage.getItem("username");

    const response = await fetch(URL, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json();
    //let jsonobject = JSON.stringify(myJson).split(', ');

    //alert(myJson.username);
    document.getElementById("maincontent").innerHTML = '<div id="profilesection"><p>Nume: ' + myJson.firstname + '</p><p>Prenume: ' + myJson.password + '</p><p>Username: ' + myJson.username + '</p><p>Departament: ' + myJson.departament + '</p><p>Office Name: ' + myJson.officename + '</p><p>Team Name: ' + myJson.teamname + '</p></div>';
}

function getprofiletest() {
    document.getElementById("maincontent").innerHTML = '<div id="profilesection"><p>Nume: ' + "Test" + '</p><p>Prenume: ' + "Test" + '</p><p>Username: ' + "Test" + '</p><p>Departament: ' + "Test" + '</p><p>Office Name: ' + "Test" + '</p><p>Team Name: ' + "Testasdasdasdasd" + '</p></div>';
}

async function getcalendar() {
    let disabledinterval = "tomato";
    let enabledinterval = "springgreen";

    const URL = "http://localhost:5050/freetime/" + localStorage.getItem("username");
    const response = await fetch(URL, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json();
    console.log(myJson);

    let freetimearray = myJson.freetime.slice(1, myJson.freetime.length - 1).split(',');
    console.log(freetimearray);
    /*
    var stringjson = JSON.stringify(myJson, null, 2);
    alert(stringjson);*/

    document.getElementById("maincontent").innerHTML = '';
    let nrid = 1;
    let buttoncolor = disabledinterval;
    for (let i = 8; i < 20; i++, nrid++) {
        let iasstring = i + "";
        let iplus1asstring = (i + 1) + "";
        if (i < 10)
            iasstring = i + "0";
        if (i < 9)
            iplus1asstring = (i + 1) + "0";

        if (freetimearray.includes(nrid.toString()))
            buttoncolor = enabledinterval;
        else
            buttoncolor = disabledinterval;
        console.log("DEBUG:" + nrid + buttoncolor)
        document.getElementById("maincontent").innerHTML += '<button class="calendarbutton" id="interval' + nrid + '" onclick="togglecalendarbutton(this.id)" style="background-color:' + buttoncolor + ';">' + i + ':00 - ' + i + ':30</button>';

        nrid++;

        if (freetimearray.includes(nrid.toString()))
            buttoncolor = enabledinterval;
        else
            buttoncolor = disabledinterval;
        document.getElementById("maincontent").innerHTML += '<button class="calendarbutton" id="interval' + nrid + '" onclick="togglecalendarbutton(this.id)" style="background-color:' + buttoncolor + ';">' + i + ':30 - ' + iplus1asstring + ':00</button>';
    }
}

async function togglecalendarbutton(idbutton) {

    const URL = "http://localhost:5050/freetime/" + localStorage.getItem("username");
    const response = await fetch(URL, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json();

    console.log(myJson);

    let freetimearray = myJson.freetime.slice(1, myJson.freetime.length - 1).split(',');

    let disabledinterval = "tomato";
    let enabledinterval = "springgreen";

    let button = document.getElementById(idbutton);
    //alert("am selectat " + idbutton);
    if (button.style.backgroundColor === disabledinterval)
        button.style.backgroundColor = enabledinterval;
    else
        button.style.backgroundColor = disabledinterval;

    let half = 0;
    let idindex = idbutton.slice(8);

    let jsontosend = myJson;
    let numbers = jsontosend.freetime.slice(1, jsontosend.freetime.length - 1).split(",").map(e => parseInt(e));

    let include = false;
    for (let i = 0; i < numbers.length; i++) {
        if (numbers[i] === parseInt(idindex)) {
            include = true;
        }
    }
    console.log(include);
    if (include === true) {
        numbers.splice(numbers.indexOf(parseInt(idindex)), 1);
    }
    else {
        numbers.push(parseInt(idindex));
    }
    jsontosend.freetime = numbers.map(e => String(e));
    jsontosend.freetime = jsontosend.freetime.join();
    jsontosend.freetime = `[` + jsontosend.freetime + ']';

    /*let jsontosend = { username: localStorage.getItem("username"), freetime: freetimearray.toString() };*/
    //jsontosend = JSON.stringify({ Username: jsontosend.username, Freetime: jsontosend.freetime });


    const requestOptions = {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body:
            '{"Username":"' + jsontosend.username + '", "Freetime":"' + jsontosend.freetime + '"}'

    };

    const putResponse = await fetch(URL, requestOptions);
    const myJson2 = await putResponse.json();

    console.log("DEBUG2: " + JSON.stringify(myJson2))

    if (idindex % 2 == 0)
        half = 1, idindex--;
    //alert("ora: " + ((idindex - 1) / 2 + 8) + "half " + half);
}

function logout() {
    location.replace("./index.html");
    localStorage.removeItem("username");
}

let searchUsers = [];

async function searchuser(username) {

    //searchUsers = {document.getElementById("option4").innerHTML, document.getElementById("option3").innerHTML, document.getElementById("option2").innerHTML, document.getElementById("option1").innerHTML}
    document.getElementById("option5").innerHTML = document.getElementById("option4").innerHTML;
    document.getElementById("option4").innerHTML = document.getElementById("option3").innerHTML;
    document.getElementById("option3").innerHTML = document.getElementById("option2").innerHTML;
    document.getElementById("option2").innerHTML = document.getElementById("option1").innerHTML;
    document.getElementById("option1").innerHTML = username;
    alert("il caut pe " + username);
    document.getElementById("maincontent").innerHTML = '';


    const URL2 = "http://localhost:5050/profiles/" + username;

    const response2 = await fetch(URL2, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson2 = await response2.json();
    //alert(myJson2);

    //let jsonobject = JSON.parse(myJson2);

    //alert(myJson2.firstname);
    document.getElementById("maincontent").innerHTML = '<div id="profilesection"><p>Nume: ' + myJson2.firstname + '</p><p>Prenume: ' + myJson2.password + '</p><p>Username: ' + myJson2.username + '</p><p>Departament: ' + myJson2.departament + '</p><p>Office Name: ' + myJson2.officename + '</p><p>Team Name: ' + myJson2.teamname + '</p></div>';


    let disabledinterval = "tomato";
    let enabledinterval = "springgreen";

    const URL = "http://localhost:5050/freetime/" + username;
    const response = await fetch(URL, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json();
    console.log(JSON.stringify(myJson, null, 2));
    console.log(myJson);

    let freetimearray = myJson.freetime.slice(1, myJson.freetime.length - 1).split(',');
    console.log(freetimearray);
    /*
    var stringjson = JSON.stringify(myJson, null, 2);
    alert(stringjson);*/

    let nrid = 1;
    let buttoncolor = disabledinterval;
    for (let i = 8; i < 20; i++, nrid++) {
        let iasstring = i + "";
        let iplus1asstring = (i + 1) + "";
        if (i < 10)
            iasstring = i + "0";
        if (i < 9)
            iplus1asstring = (i + 1) + "0";

        if (freetimearray.includes(nrid.toString()))
            buttoncolor = enabledinterval;
        else
            buttoncolor = disabledinterval;
        console.log("DEBUG:" + nrid + buttoncolor)
        document.getElementById("maincontent").innerHTML += '<button class="calendarbutton" id="interval' + nrid + '" style="background-color:' + buttoncolor + ';">' + i + ':00 - ' + i + ':30</button>';

        nrid++;

        if (freetimearray.includes(nrid.toString()))
            buttoncolor = enabledinterval;
        else
            buttoncolor = disabledinterval;
        document.getElementById("maincontent").innerHTML += '<button class="calendarbutton" id="interval' + (nrid + 1) + '" style="background-color:' + buttoncolor + ';">' + i + ':30 - ' + iplus1asstring + ':00</button>';
    }
}


var allusers = [];




async function createnewmeeting() {
    let userswithfreetime = [];
    let nousers = localStorage.getItem("nousers");
    let usersformeeting = [];
    let jsontosend = {};

    for (let i = 0; i < nousers; i++) {
        userswithfreetime.push([]);
        if (document.getElementById(("cbx" + (i + 1))).checked) {
            usersformeeting.push({ "username": allusers[i] });


            const URL = "http://localhost:5050/freetime/" + allusers[i];
            console.log(allusers[i]);
            const response = await fetch(URL, {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            const myJson = await response.json();
            console.log(userswithfreetime);
            userswithfreetime[userswithfreetime.length - 1].push(myJson.freetime);
        }
    }
    jsontosend.users = usersformeeting;

    let j = 0;
    console.log(userswithfreetime[0][0]);
    for (let i = 1; i <= 24; i++) {
        for (j = 0; j < userswithfreetime.length; j++) {
            console.log("j=" + j + "  " + userswithfreetime[j][0]);
            if (!userswithfreetime[j][0].includes(i.toString()))
                break;
        }
        console.log("jul" + j);
        if (j > 1) {
            alert("Meeting has been created at interval " + i);
            return;
        }
    }

    alert("No available time for a break for all selected users");
}

async function newmeeting() {
    const URL = "http://localhost:5050/users";
    const response = await fetch(URL, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json();
    let objjson = myJson;
    let objList = [];
    objjson.map(e => (objList.push(e)));
    console.log(objList);
    let userdb = [];
    objList.map(e => userdb.push(e.username));
    console.log(userdb);
    document.getElementById("maincontent").innerHTML = "";
    let cbxid = 1;
    allusers = userdb.slice();
    userdb.forEach(user => {
        document.getElementById("maincontent").innerHTML += '<p>' + user + '</p><input type="checkbox" id="cbx' + cbxid + '" name="horns"><br>';
        cbxid++;
    });
    localStorage.setItem("nousers", (cbxid - 1).toString());
    document.getElementById("maincontent").innerHTML += '<button style="background-color:blue;padding:50px;" onclick="createnewmeeting()">CREATE NEW MEETING</button>';
}