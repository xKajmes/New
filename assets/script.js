istnieje = false;
licznik = 1;
istnieje2 = false


function DeleteStudent(id) {

    if (confirm("Jesteś pewny żeby usunąć studenta o id = " + id)) {

        $.ajax({
            type: "DELETE",
            url: "/panel/student/" + id,
            // data: {},
            dataType: "json",
            success: function() {
                alertbs('success', '#jsoninfo', 'Usunałeś studetna o ' + id)
                GetAllStudents()
            }
        });
    }
}
function GetMyClass() {
    if (istnieje == true) {
        $(d).remove();
        istnieje = false
    }
    //alert("witaj");
    var url = $(location).attr('href');
    var parts = url.split("/");
    panelornot = parts[parts.length - 1]
    console.log(parts[parts.length - 1])
    $.ajax({
        type: "GET",
        url: "/teacher/lookclass",
        dataType: "json",
        success: function(results) {
            console.log("SUKCES")
            console.log("SUKCES")
            istnieje = true;
            d = document.createElement('table');
            $(d).addClass('table table-striped table-light')
            $(d).insertAfter($("#jsoninfo")) //main div
             $(d).html("<thead class='thead-light'> <tr><th scope='col'>#</th><th scope='col'>Imie</th><th scope='col'>Nazwisko</th><th scope='col'>Wiek</th><th scope='col'>Oceny</th></tr></thead><tbody>"); //<th scope='col'>Usuń</th><th scope='col'>Edytuj</th>
             for (var i in results) {
                                // "<a  onclick='DeleteStudent(" + results[i].ID + ");'>Usuń</a>" + "</td><td> " + "<a  onclick='EditStudent(" + results[i].id + ");'>Edytuj</a>"
                $(d).append("<tr><th scope='row'>" + results[i].ID + "</th><td> " + results[i].name + "</td><td> " + results[i].surname + " </td><td>" + results[i].age + " </td>  <td>" + "<a  onclick='GetGrades(" + results[i].ID + ");'>Zobacz</a>" + "</td></tr>");
            }
            }
    })
}
function GetNameStudent(id){
    $.ajax({
        type: "GET",
        url: "/panel/student/" + id,

        dataType: "json",
        success: function(results) {
            console.log(results.name+" "+ results.surname)
            document.getElementById("getname").innerHTML = results.name+" "+ results.surname
        }
})
}

function AddGradeClick(id){
    //id = $("#studentid").val()
    console.log($("#grade").val())
    console.log($("#comment").val())
    
    idteacher = MyID()
    console.log(idteacher)
    idstudent = id
    console.log($("#idstudent").val())
    //console.log(MyID())
    $.ajax({
        type: "POST",
        url: "/teacher/grades",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify({
            number: $("#grade").val(),
            comment: $("#comment").val(),
            id_student: $("#idstudent").val(),
            id_teacher: MyID()
        }),
        dataType: "json",
        success: function() {
            console.log('sukces')
            $("#EditModal").modal('hide');
            alertbs('success', '#jsoninfo', 'Pomyślnie zedytowano studenta o id = ' + id)
            //GetAllStudents()

        },
        error: function() {
            console.log('alert')
            $("#EditModal").modal('hide');
            alertbs('danger', '#jsoninfo', 'Coś poszło nie tak! ')

        }
    });
} 

function GetAllStudents() {
    
    // if (istnieje == true) {
    //     $(d).remove();
    //     istnieje = false
    // }
    //alert("witaj");
    var url = $(location).attr('href');
    var parts = url.split("/");
    panelornot = parts[parts.length - 1]
    console.log(parts[parts.length - 1])
    // if(panelornot = "class"){
    //     GetMyClass();
    //     return
    // }
    $.ajax({
        type: "GET",
        url: "/panel/student/",
        dataType: "json",
        success: function(results) {
            console.log("SUKCES")
            istnieje = true;
            d = document.createElement('table');
            $(d).addClass('table table-striped table-light')
            $(d).insertAfter($("#jsoninfo")) //main div
            
            if (panelornot == "panel" ) {
                
                $(d).html("<thead class='thead-light'> <tr><th scope='col'>#</th><th scope='col'>Imie</th><th scope='col'>Nazwisko</th><th scope='col'>Wiek</th><th scope='col'>Rola</th></tr></thead><tbody>");
                //$("#jsoninfo").append(" <tbody>")
                for (var i in results) {
                    
                    if(results[i].name == "" || results[i].role == "admin" || results[i].role == "dziekan" || results[i].role == "nauczyciel"){
                        
                        return
                        
                    }
                    // if(results[i].class == myclass){

                    // $(d).append("<tr><th scope='row'>" + results[i].ID + "</th><td> " + results[i].name + "</td><td> " + results[i].surname + " </td><td>" + results[i].age + " </td> <td>" + results[i].role + "</td></tr>");
                    // }
                }
            }
            if (panelornot == "dziekan" ) {

                $(d).html("<thead class='thead-light'> <tr><th scope='col'>#</th><th scope='col'>Imie</th><th scope='col'>Nazwisko</th><th scope='col'>Wiek</th><th scope='col'>Rola</th><th scope='col'>Usuń</th><th scope='col'>Edytuj</th></tr></thead><tbody>");
                //$("#jsoninfo").append(" <tbody>")
                for (var i in results) {
                    
                    if(results[i].name == "" || results[i].role == "admin" || results[i].role == "dziekan" || results[i].role == "nauczyciel"){
                        
                        return
                        
                    }

                    $(d).append("<tr><th scope='row'>" + results[i].ID + "</th><td> " + results[i].name + "</td><td> " + results[i].surname + " </td><td>" + results[i].age + " </td> <td>" + results[i].role +" </td> <td>" + "<a  onclick='DeleteStudent(" + results[i].ID + ");'>Usuń</a>" + "</td><td> " + "<a  onclick='EditStudent(" + results[i].id + ");'>Edytuj</a>" + "</td></tr>");
                }
            } else {
                console.log("EJJ")
                $(d).html("<thead class='thead-light'> <tr><th scope='col'>#</th><th scope='col'>Imie</th><th scope='col'>Nazwisko</th><th scope='col'>Wiek</th></tr></thead><tbody>");
                //$("#jsoninfo").append(" <tbody>")
                for (var i in results) {
                    if(results[i].name == "" || results[i].role == "admin" || results[i].role == "dziekan" || results[i].role == "nauczyciel"){
                        return
                    }

                    $(d).append("<tr><th scope='row'>" + results[i].ID + "</th><td> " + results[i].name + "</td><td> " + results[i].surname + " </td><td>" + results[i].age + " </td></tr>");
                }
            }
            $(d).append(" </tbody></table>")
        },
        error: function(results) {
            console.log(results.responseText)
        }

    });
}

function EditStudent(id) {

    $('#EditModal').modal();
    $.ajax({
        type: "GET",
        url: "/panel/student/" + id,
        dataType: "json",
        success: function(results) {
            console.log(results.name);
            $('input[name="name2"]').val(results.name);
            $('input[name="surname2"]').val(results.surname);
            $('input[name="age2"]').val(results.age);
            $('input[name="studentid"]').val(id);
        }
    });
}
function CompleteStudent(){
    var userid
    $.ajax({
        type: "GET",
        url: "/panel/myid/",
        dataType: "json",
        success: function(results) {
            userid = results.id;
            EditStudent(userid);
            console.log(userid);
        }
    });
    
}

function GetOneStudent() {
    $.ajax({
        type: "GET",
        url: "/panel/student/",
        dataType: "json",
        success: function(results) {
            for (var i in results) {
                $("#jsoninfo").append(results[i].id + " " + results[i].name + " " + "<a  onclick='deletestudent(" + results[i].id + ");'>Usuń</a>" + " " + "<a  onclick='editstudent(" + results[i].id + ");'>Edytuj</a>" + "<br>");
            }
        }
    });
}

function EditStudentClick() {
    id = $("#studentid").val()
    $.ajax({
        type: "PUT",
        url: "/panel/student/" + id,
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify({
            Name: $("#name2").val(),
            Surname: $("#surname2").val(),
            Age: $("#age2").val(),
        }),
        dataType: "json",
        success: function() {
            console.log('sukces')
            $("#EditModal").modal('hide');
            alertbs('success', '#jsoninfo', 'Pomyślnie zedytowano studenta o id = ' + id)
            GetAllStudents()

        },
        error: function() {
            console.log('alert')
            $("#EditModal").modal('hide');
            alertbs('danger', '#jsoninfo', 'Coś poszło nie tak! ')

        }
    });

}


function AddStudent() {

    $.ajax({
        type: "POST",
        url: "/panel/student/",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify({
            Name: $('#name1').val(),
            Surname: $('#surname1').val(),
            Age: $('#age1').val(),
        }),
        dataType: "json",
        success: function() {
            console.log($('#age1').val())
            $('#age1').val("")
            $('#name1').val("")
            $('#surname1').val("")
            $("#AddModal").modal('hide');
            alertbs('success', '#jsoninfo', 'Pomyślnie dodano studenta!')
            GetAllStudents()
        },
        error: function() {
            $("#AddModal").modal('hide');
            alertbs('danger', '#jsoninfo', 'Coś poszło nie tak! ')
        }
    });

}

function Login() {
    $('#LoginButton').attr("disabled", true);
    $('#LoginButton').html("<span class='spinner-border spinner-border-sm' role='status' aria-hidden='true'></span> Logowanie ")
        $.ajax({
            type: "POST",
            url: "/auth/login",
            contentType: 'application/json; charset=utf-8',
            data: JSON.stringify({
                Login: $('#login').val(),
                Password: $('#password').val(),
            }),
            dataType: "json",
            success: function(data) {
                alertbs('success', '#jsoninfo', 'Zalogowano!')
                $('#LoginButton').attr("disabled", false);
                $('#LoginButton').html("ZALOGOWANO")
                token = data.token;
                
                console.log(token)
                iduser = data.userid
                console.log(" ")
                console.log(iduser)
                location.href='/home';
                // // document.cookie = "token="+token;
                // // document.cookie = 'cookie=ok;expires='+now.toUTCString()+';path=/';
                // var now = new Date();
                // var time = now.getTime();
                // var expireTime = time + 250*36000;
                // now.setTime(expireTime);
                // document.cookie = 'token='+token+';expires='+now.toUTCString()+';path=/';
                // console.log(document.cookie);  // 'Wed, 31 Oct 2012 08:50:17 UTC'
                // //location.href("/home")
                // var decodejwt = parseJwt(token);
                // console.log('Zdekodowany jwt')
                // console.log(decodejwt)
                // console.log('USER ID:')
                // console.log(decodejwt.user_id)
                //location.href = "home";
               // SessionStart(nick)
            },
            error: function(data) {
                console.log(data.responseText);
                console.log(data[1]);
                $('#LoginButton').attr("disabled", false);
                $('#LoginButton').html("Login")
                alertbs('danger', '#jsoninfo', "Podałeś nieprawidłowe dane!")
            }
        });
}

function Register() {
    $('#RegisterButton').attr("disabled", true);
    $('#RegisterButton').html("<span class='spinner-border spinner-border-sm' role='status' aria-hidden='true'></span> Rejestrowanie ")
    p1 = $('#password2').val()
    p2 = $('#password2').val()
    nick = $('#login').val()
    if (p1 == p2) {
        console.log('Zgadzają się!')
        $.ajax({
            type: "POST",
            url: "/auth/register",
            contentType: 'application/json; charset=utf-8',
            data: JSON.stringify({
                Login: nick,
                Password: p1,
                
            }),
            dataType: "json",
            done: function(data) {
                console.log('sukces');
                console.log(data.responseText);
                //alert('success', '#jsoninfo', '<a href="google.com">KLIKNIJ TUTAJ</a>')
                $('#RegisterButton').attr("disabled", false);
                $('#RegisterButton').html("Zarejestrowano!");
                console.log('http://localhost:8080/register')
                //location.href = "login?register=succes";
                
            },
            error: function(data) {
                
                alert(data.responseText );
                console.log('cosnietak');
                console.log(data.responseText);
                alertbs('danger', '#jsoninfo', 'Podany login już istnieje! ')
                $('#RegisterButton').attr("disabled", false);
                $('#RegisterButton').html("Zarejestruj ")
            }
        });
    } else {
        alertbs('danger', '#jsoninfo', 'Hasła się nie zgadzają! ')
        $('#RegisterButton').attr("disabled", false);
        $('#RegisterButton').html("Zarejestruj ")
    }
}

function Logout() {
    location.href = "/logoutnow"
    

}

function CheckGrades(id){
    licznikocen = 0;
    suma = 0;

    $.ajax({
        type: "GET",
        url: "/teacher/grades/" + id,
        contentType: 'application/json; charset=utf-8',
        dataType: "json",
        success: function(results) {
            console.log(results.responseText)
            
            check = document.createElement('table');
            $(check).addClass('table table-striped table-light')
            $(check).insertAfter($("#checkgrades")) //main div
            $(check).html("<thead class='thead-light'> <tr><th scope='col'>Ocena</th><th scope='col'>Data</th><th scope='col'>Komentarz</th><th scope='col'>Usuń</th></thead><tbody>");
            for (var i in results) {

                full = results[i].CreatedAt.split('.')[0]
                day = full.split('T')[0]
                hour = full.split('T')[1]
                hour = hour.split(':')[0] + ":" + hour.split(':')[1]
                $(check).append("<tr><th scope='row'>" + results[i].number + "</th><td>" + day + " " + hour + "</td><td> " + results[i].comment + "</td> <td>" + "<a  onclick='DelGrade(" + results[i].ID + ");'>Usuń</a></td></tr>");
                licznikocen++;
                suma+=results[i].number;
            }
            if(licznikocen>0){
            avg = suma/licznikocen;
            
            avg = Round(avg,2)
            console.log(avg)
            $(check).append("<tr class='table-success'><th scope='row'>" + avg + "</th><td colspan='3'>" + "Srednia" + "</tr>");
            }
        },
        error: function(results) {
            console.log(results.responseText)

        }
    })
}

function GetGrades(id){
    location.href = "/teacher/grades?id="+id;

}
function DelGrade(id){
    if (confirm("Napewno chcesz usunąć tą ocene? ")) {
    $.ajax({
        type: "DELETE",
        url: "/teacher/grades/" + id,
        contentType: 'application/json; charset=utf-8',
        dataType: "json",
        success: function(results) {
            alertbs('success', '#jsoninfo', 'Usunięto ocene!')
        },
        error: function(results) {
            console.log(results.responseText)

        }

})
    }
}
function MyID(){

    $.ajax({
        type: "GET",
        url: "/panel/myid/",
        contentType: 'application/json; charset=utf-8',
        dataType: "json",
        success: function(results) {
           idteacher = results.id;
           console.log("ID TEACHER:"+idteacher)
           return idteacher;
        },
        error: function(results) {
            console.log(results.responseText)

        }

})
    
}
// }

// function SessionStart(login) {
//     location.href = "/home"
//     console.log('WITAJ WITAJ' + login)
// }

function alertbs($class, $append, $tresc) {
    numeralert = 'numeralert' + licznik;
    if ($class == "success") {

        x = document.createElement('div');
        $(x).addClass('fixa alert alert-success alert-dismissible container ' + numeralert);
        $(x).appendTo($($append)) //main div
        $(x).html('<button type="button" class="close" data-dismiss="alert">×</button> <strong>Sukces!</strong> ' + $tresc);
        licznik++
        closealert(numeralert)
    }
    if ($class == "danger") {
        y = document.createElement('div');
        $(y).addClass('fixa alert alert-danger alert-dismissible container ' + numeralert);
        $(y).appendTo($($append)) //main div
        $(y).html('<button type="button" class="close" data-dismiss="alert">×</button> <strong>Uwaga!</strong> ' + $tresc);
        licznik++
        closealert(numeralert)
    }

}

function closealert(string) {
    setTimeout(function() {

        $('.' + numeralert).remove();
    }, 1000);
}
function parseJwt (token) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};
function registersucces(){
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    var product = urlParams.get('register')
    console.log(product)
    if(product == "succes"){
        console.log('UDANY')
        alertbs('success', '#jsoninfo', 'Zarejestrowano pomyslnie teraz sie zaloguj!')
    }
}
function Round(n, k)
{
    var factor = Math.pow(10, k);
    return Math.round(n*factor)/factor;
}


    
