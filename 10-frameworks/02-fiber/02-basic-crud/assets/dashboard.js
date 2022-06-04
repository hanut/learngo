var actionDlg, fname, lname, age, address, password, role;

console.log("Loaded dashboard");
var actionDlg = document.querySelector("#actionDlg"),
fname = document.querySelector("#firstName"),
lname = document.querySelector("#lastName"),
age = document.querySelector("#age"),
address = document.querySelector("#address"),
password = document.querySelector("#password"),
role = document.querySelector("#role"),
userForm = document.querySelector("#userForm"),
formLoader = document.querySelector("#formLoader"),
closeBtn = document.querySelector("#closeBtn"),
saveBtn = document.querySelector("#saveBtn");

var mode = "";

const token = (() => {
  try {
    // Retrieve the auth data from the localstorage
    let authData = JSON.parse(localStorage.getItem("authData"));
    let timeLeft = ((authData.expiry*1000) - Date.now())/1000;
    if (timeLeft < 30) {
      throw new Error("Token expired");
    }
    return authData.token;
  } catch(error) {
    console.warn(error);
    window.location.replace("/webapp");
  }
})()
console.log("Token:", token);

function addUser() {
  mode = "add";
  disableInputs(false);
  showFormLoader(false);
  actionDlg.classList.remove("hide");
}

function viewUser(userId) {
  mode = "view";
  disableInputs(true);
  showFormLoader(true);
  fetch("/users/" + userId, {
    headers: {
      "Authorization": `Bearer ${token}`
    }
  })
    .then((r) => r.json())
    .then((r) => {
      fillValues(r);
      showFormLoader(false);
      actionDlg.classList.remove("hide");
    })
    .catch(console.warn);
}

function editUser() {
  mode = "view";
  disableInputs(false);

}

function removeUser() {}

function save() {
  showFormLoader(true);
  const user = {
    FirstName: fname.value,
    LastName: lname.value,
    Password: password.value,
    Age: age.value,
    Address: address.value,
    Role: role.value,
  };
  console.log(user);
  try {
    // save the user
  } catch(error) {
    console.warn(error)
  }
}

function closeDlg() {
  actionDlg.classList.add("hide");
  userForm.reset();
}

function fillValues(user) {
  fname.value = user.FirstName;
  lname.value = user.LastName;
  age.value = user.Age;
  address.value = user.Address;
  password.value = user.Password;
  role.value = user.Role;
}

function disableInputs(state = true) {
  fname.disabled = state;
  lname.disabled = state;
  age.disabled = state;
  address.disabled = state;
  password.disabled = state;
  role.disabled = state;
}

function showFormLoader(shown = false) {
  if (shown) {
    formLoader.style.display = "block";
    userForm.style.display = "none";
  } else {
    formLoader.style.display = "none";
    userForm.style.display = "block";
  }
}
