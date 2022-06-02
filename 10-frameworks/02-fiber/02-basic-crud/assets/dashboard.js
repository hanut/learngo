var actionDlg, fname, lname, age, address, password, role;

console.log("Loaded dashboard");
actionDlg = document.querySelector("#actionDlg");
fname = document.querySelector("#firstName");
lname = document.querySelector("#lastName");
age = document.querySelector("#age");
address = document.querySelector("#address");
password = document.querySelector("#password");
role = document.querySelector("#role");


function viewUser() {
  disableInputs(true);
  actionDlg.classList.remove("hide");  
}

function editUser() {
  disableInputs(false);
}

function removeUser() {

}

function dismissDlg() {
  actionDlg.classList.add("hide");
}

function disableInputs(state = true){
  fname.disabled = state;
  lname.disabled = state;
  age.disabled = state;
  address.disabled = state;
  password.disabled = state;
  password.role = state;
}