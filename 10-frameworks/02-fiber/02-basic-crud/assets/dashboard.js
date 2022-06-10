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
  saveBtn = document.querySelector("#saveBtn"),
  formError = document.querySelector("#formError");

var mode = "";

const token = (() => {
  try {
    // Retrieve the auth data from the localstorage
    let authData = JSON.parse(localStorage.getItem("authData"));
    let timeLeft = (authData.expiry * 1000 - Date.now()) / 1000;
    if (timeLeft < 30) {
      throw new Error("Token expired");
    }
    return authData.token;
  } catch (error) {
    console.warn(error);
    window.location.replace("/webapp");
  }
})();

function addUser() {
  mode = "add";
  saveBtn.style.display = "inline-block";
  disableInputs(false);
  showFormLoader(false);
  actionDlg.classList.remove("hide");
}

function viewUser(userId) {
  mode = "view";
  saveBtn.style.display = "none";
  showFormLoader(true);
  actionDlg.classList.remove("hide");
  fetch("/users/" + userId, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then((r) => {
      if (r.status !== 200) {
        throw new Error(r.statusText);
      }
      return r.json();
    })
    .then((r) => {
      fillValues(r);
      disableInputs(true);
      showFormLoader(false);
    })
    .catch((error) => {
      flashError(error).then(() => {
        actionDlg.classList.add("hide");
      });
    });
}

function editUser(userId) {
  mode = "edit";
  saveBtn.style.display = "inline-block";
  showFormLoader(true);
  actionDlg.classList.remove("hide");

  fetch("/users/" + userId, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then((r) => {
      if (r.status !== 200) {
        throw new Error(r.statusText);
      }
      return r.json();
    })
    .then((r) => {
      fillValues(r);
      disableInputs(false);
      showFormLoader(false);
    })
    .catch((error) => {
      flashError(error).then(() => {
        actionDlg.classList.add("hide");
      });
    });
}

function removeUser(userId) {
  const response = confirm("Are you sure you want to delete that user ?");
  if (!response) return;
  fetch("/users/" + userId, {
    method: "delete",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
    .then((r) => {
      console.log(r.status, r.statusText);
      return r.text();
    })
    .then((r) => window.location.reload())
    .catch(flashError);
}

function save(e) {
  console.log("event", e);
  e.preventDefault();
  showFormLoader(true);
  const user = {
    FirstName: fname.value,
    LastName: lname.value,
    Password: password.value,
    Age: age.value,
    Address: address.value,
    Role: role.value,
  };
  try {
    // save the user
    if (mode === "add") {
      fetch("/users", {
        method: "post",
        body: JSON.stringify(user),
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
    } else {
      alert("TODO");
    }
  } catch (error) {
    console.warn(error);
  }
}

function closeDlg() {
  console.log(mode);
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

function flashError(errorMsg, time = 3) {
  formError.innerHTML = errorMsg;
  return new Promise((resolve) => {
    setTimeout(() => {
      formError.innerHTML = "";
      resolve();
    }, time * 1000);
  });
}
