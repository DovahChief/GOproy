function valida() {
    var nombre      = document.getElementById('nombre').value;
    var apellido    = document.getElementById('apellido').value;
    var username    = document.getElementById('username').value;
    var email       = document.getElementById('email').value;

    if (nombre === "null" || nombre === "undefined") {
        alert("se debe ingresar un nombre");
    }

}
