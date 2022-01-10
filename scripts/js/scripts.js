
class Techniciens {
  constructor(_FirstName, _LastName, _Login, _Number, _Password, _Level) {
    this._FirstName = _FirstName;
    this._LastName = _LastName;
    this._Login = _Login;
    this._Number = _Number;
    this._Password = _Password;
    this._Level = _Level;
  }
};

class Equipements {
  constructor(_CaracTech, _Marque, _Modele, _EtatAcha, _PrixAcha, _DateAcha, _EtatUti, _FourniId) {
    this._CaracTech = _CaracTech;
    this._Marque = _Marque;
    this._Modele = _Modele;
    this._EtatAcha = _EtatAcha;
    this._PrixAcha = _PrixAcha;
    this._DateAcha = _DateAcha;
    this._EtatUti = _EtatUti;
    this._FourniId = _FourniId;
  }
};

class Prestataires {
  constructor(_Name, _Number) {
    this._Name = _Name;
    this._Number = _Number;
  }
};

class Utilisateurs {
  constructor(_FirstName, _LastName, _Number) {
    this._FirstName = _FirstName;
    this._LastName = _LastName;
    this._Number = _Number;
  }
};


class Fournisseurs {
  constructor(_Name, _Adresse, _Number) {
    this._Name = _Name;
    this._Adresse = _Adresse;
    this._Number = _Number;
  }
};

class Affectation {
  constructor(_EquipId, _UserId, _DateDebut, _DateFin) {
    this._EquipId = _EquipId;
    this._UserId = _UserId;
    this._DateDebut = _DateDebut;
    this._DateFin = _DateFin;
  };
}

class Services {
  constructor(_DateExp, _Ref, _Type, _PrestataireID) {
    this._DateExp = _DateExp;
    this._Ref = _Ref;
    this._Type = _Type;
    this._PrestataireID = _PrestataireID;
  };
}

const urlcourante = document.location.href;
const debut = urlcourante.substring(0, urlcourante.indexOf("/", 7))


function createTech() { // fonction de création d'un technicien
  var nom = document.getElementById('nom').value;
  var prenom = document.getElementById('prenom').value;
  var login = document.getElementById('login').value;
  var pwd = document.getElementById('pwd').value;
  var numero = document.getElementById('numero').value;
  var lvl = 0;

  technicien = new Techniciens(nom, prenom, login, numero, pwd, lvl);

  url = debut + ":8090/techniciens";

  fetch(url, {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    },
    body: JSON.stringify(technicien)
  })
    .then(response => {
      if (response.ok) {
        window.location.href = "equipement.php"
      }
    })
    .catch(data => {
      message = document.getElementById('response');
      message.innerHTML = res.result
    })
}




function createEquip() { //creer uniquement l'equipement

  marque = document.getElementById('mequip').value;
  modele = document.getElementById('moequip').value;
  prix = document.getElementById('pequip').value;
  etatAcha = document.getElementById('etequip').value;
  fourniId = document.getElementById('efourni').value;
  caracTech = document.getElementById('caract').value;
  dateAcha = document.getElementById('dequip').value;

  prix = parseInt(prix);
  equipement = new Equipements(caracTech, marque, modele, etatAcha, prix, dateAcha, "", fourniId);

  fetch('http://localhost:8090/equipement', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    },
    body: JSON.stringify(equipement)
  })
    .then(response => {
      if (response.ok) {
        alert("equipement " + marque + " " + modele + " ajouté");
        window.location.reload();
      }
    })
    .catch(data => {
      console.log(data)
    })
}
function getEquip(id) { // récupère un équipement en particulier
  //code
}
function createEquiFourni() { // créer équipement et fournisseur en même temps info de l'équipement

  this.preventDefault();

  marque = document.getElementById('mequipe').value;
  modele = document.getElementById('moequip').value;
  prix = document.getElementById('pequip').value;
  etatAcha = document.getElementById('etequip').value;
  fourniId = document.getElementById('efourni').value;
  caracTech = document.getElementById('caract').value;
  dateAcha = document.getElementById('dequip').value;

  //info fournisseur
  nom = document.getElementById("nomf").value;
  adresse = document.getElementById("adresse").value;
  number = document.getElementById("numero").value;

  equipement = new Equipements(caracTech, marque, modele, etatAcha, prix, dateAcha, fourniId);
  fournisseur = new Fournisseurs(nom, adresse, number);

  fetch('http://localhost:8090/equipements', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    },
    body: {
      equi: JSON.stringify(equipement),
      four: JSON.stringify(fournisseur)
    }
  })
    .then(res => {
      if (res.ok) {
        alert("equipement " + marque + " " + modele + " ajouté");
      }
    })
    .catch(data => {
      console.log(data)
    })
}

function userSe(id) {

  //liste user
  myEl = document.getElementById(id);

  url = debut + ":8090/users";
  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEl.innerHTML = "";
        response.json()
          .then((users) => {
            console.log(users);

            for (let le of users) {
              let listItem = document.createElement('option');
              listItem.value = le._id;
              listItem.innerHTML = `${le._FirstName} ${le._LastName}`;

              myEl.appendChild(listItem);
            }
          })
      }
    })
    .catch(error => {
      console.log(error)
    })
}
function equipSe(id) {

  //liste equipement dans select
  myEle = document.getElementById(id);

  url = debut + ":8090/equipement";
  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEle.innerHTML = "";
        response.json()
          .then((equip) => {
            console.log(equip);
            for (let le of equip) {
              let listItem = document.createElement('option');
              listItem.value = le._id;
              listItem.innerHTML = `${le._Marque} ${le._Modele}`;

              myEle.appendChild(listItem);
            }
          })
      }
    })
    .catch(error => {
      console.log(error)
    })
}
function presSe(id) { //liste des prestataires dans select
  myEle = document.getElementById(id);

  url = debut + ":8090/prestataires";
  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEle.innerHTML = "";
        response.json()
          .then((equip) => {
            console.log(equip);
            for (let le of equip) {
              let listItem = document.createElement('option');
              listItem.value = le._id;
              listItem.innerHTML = `${le._Name}`;

              myEle.appendChild(listItem);
            }
          })
      }
    })
    .catch(error => {
      console.log(error)
    })
}

function getEquips() { // liste des équipements et chargement dans la page

  myEquip = document.getElementById("equip");
  url = debut + ":8090/equipement"
  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEquip.innerHTML = '';
        response.json()
          .then((data) => {
            console.log(data);
            console.log("response");

            for (let el of data) {
              let listItem = document.createElement('div');
              //  listItem.addClass("card");
              listItem.innerHTML = `
              <div class="card">
              <div class="row">
            <div class="col-md-3" style="align-items: center;display: inline-grid;">

              <img height="100vh" src="./assets/media/img1.jpg" alt="Card image cap">
              <div class="card-body">
                <h5 class="card-title">Utilisateur X</h5>
              </div>
            </div>
            <div class="col-md-5">
              <h5>${el._Marque} : ${el._Modele} </h5>
              <ul>
                <li>Prix : ${el._PrixAcha} </li>
                <li>Etat achat: ${el._EtatAcha} </li>
                <li>Date achat : ${el._DateAcha} </li>
                <li>Etat d'utilisation : ${el._EtatUti} </li>
                </ul>
                <p>Service appliqué</p>
                <p>Description : ${el._CaracTech} </p>
            </div>
            <div class="col-md-4 p-4">
              <form>
              <div class="form-group">
                <label for="etat">Changer état : </label>
                <select class="form-control" name="etat" id="etat">
                  <option selected>Sélectionner...</option>
                  <option value="">En cours</option>
                  <option value="">En panne</option>
                  <option value="">Hors parc</option>
                </select>
              </div>
              <div class="form-group pt-2">
                <input class="form-control" type="text" placeholder="raison">
              </div>
              <div class="form-group pt-2">
                <input class="form-control" type="submit" value="Valider">
              </div>

            </form>
            </div>
          </div>
          </div>`;

              myEquip.appendChild(listItem);
              userSe(el._id);

            }

            // myList.appendChild(aff)
            // myList.innerHTML = JSON.stringify(data)
          })
          .catch(error => console.log(error));
      }
    })
    .catch(data => {
      console.log(data)
    })
}



function getUsers() { // récupère la liste des utilisateurs 
  // et charge les accordeons
  myEls = document.getElementById("accordion");
  url = debut + ":8090/users";

  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEls.innerHTML = "";
        response.json()
          .then((data) => {
            console.log(data);
            for (let el of data) {
              let listItem = document.createElement('div');
              listItem.innerHTML = `
            <div class="card">
            <div class="card-header">
                <a class="btn" data-toggle="collapse" href="#collapseOne">
                    ${el._FirstName} ${el._LastName}
                </a>
            </div>
            <div id="collapseOne" class="collapse show" data-parent="#accordion">
                <div class="card-body">
                    <ul>
                        <li><strong>Numéro </strong> : ${el._Number} </li>
                        <li><strong>Id </strong> : ${el._id} </li>
                    </ul>
                </div>
            </div>
        </div>`;
              //   listItem.addClass("card");
              myEls.appendChild(listItem);
            }
          })
      }
    })
    .catch(data => {
      console.log(data)
    })
}
function createUser() { // créer un utilisateur
  nom = document.getElementById("nom").value;
  prenom = document.getElementById("prenom").value;
  number = document.getElementById("numero").value;

  console.log("test");

  utilisateur = new Utilisateurs(nom, prenom, number);

  url = debut + ":8090/users";

  fetch(url, {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    },
    body: JSON.stringify(utilisateur)
  })
    .then(response => {
      if (response.ok) {
        alert("Utilisateur : " + nom + " ajouté !");
        getUsers();
      }
    })
    .catch(data => {
      console.log(data.result)
    })

}
function getUser(id) { //récupérer un utilisateur de la bdd avec son id 
  fetch('http://localhost:8090/utilisateur?id=' + id, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(user => {
      return user
    })
    .catch(error => console.log(error))
}

function getPrestataires() { //récupere la liste des prestataires et charge le tableau
  // de la page prestatairesphp
  myEls = document.getElementById("pres");
  url = debut + ":8090/prestataires";

  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEls.innerHTML = '';
        response.json()
          .then((data) => {
            for (let el of data) {
              let listItem = document.createElement('tr');
              listItem.appendChild(
                document.createElement('td')
              ).textContent = `${el._id}`;

              listItem.appendChild(
                document.createElement('td')
              ).textContent = `${el._Name}`;

              listItem.appendChild(
                document.createElement('td')
              ).textContent = `${el._Number}`;
             
              listItem.appendChild(
                document.createElement('td')
              ).innerHTML = `<button type="button" onclick="modPres(el._id);" class="btn btn-primary btn-sm btn-block">Modifier </button> <input type="text" name="desaff" readonly value="${el._id}" class="form-control" onclick="delPres(this.value)" btn-sm btn-block">cliquez pour désaffecter</input>`;
              myEls.appendChild(listItem);
            }
          })
      }
    })
    .catch(data => {
      console.log(data)
    })
}
function createPrestataire() { //créer un prestataire
  var nom = document.getElementById('nompres').value;
  var numero = document.getElementById('numpres').value;

  prestataire = new Prestataires(nom, numero);
  url = debut + ":8090/prestataires"
  fetch(url, {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    },
    body: JSON.stringify(prestataire)
  })
    .then(response => {
      if (response.ok) {
        alert("Prestataire : " + nom + " ajouté !");
        getPrestataires();
      }
    })
    .catch(data => {
      console.log(data.result)
    })
}
function delPres(id){
  url = debut +":8090/prestataires/" + id;
  fetch(url, {
    method: "DELETE",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
  .then(response => {
    if(response.ok){
      alert("Prestataire supprimé.");
      window.location.reload();
    }
  })
  .catch(data => {
    console.log(data.result)
  })
}

function getAffectation(id) { //récupérer une affectation avec l'id de l'équipement 
  fetch('http://localhost:8090/affectation?id=' + id, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(equip => {
      return equip
    })
    .catch(error => console.log(error))
}
function affecter(idequip, iduser) { //création des affectations
  affectation = new Affectation(idequip, iduser, '', '');

  url = debut + ":8090/affectations";

  fetch(url, {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    },
    body: JSON.stringify(affectation)
  })
    .then((response) => {
      if (response.ok) {
        alert("Affection créée avec succès ");
        window.location.reload();
      }
    })
    .catch(error => {
      console.log(error);
      alert(error);
    })
}
function getAffectations(){
  myEquip = document.getElementById("equip");
  
  url = debut + ":8090/affectations";

  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEquip.innerHTML = '';
        response.json()
          .then((data) => {
           console.log(data);
            for (let el of data) {
              let listItem = document.createElement('div');
              //  listItem.addClass("card");
              listItem.innerHTML = `
              <div class="card">
              <div class="row">
            <div class="col-md-3" style="align-items: center;display: inline-grid;">

              <img height="100vh" src="./assets/media/img1.jpg" alt="Card image cap">
              <div class="card-body">
                <h5 class="card-title"> Utilisateur : ${el._affectationData._User} </h5>
              </div>
            </div>
            <div class="col-md-5">
              <p>Marque et modèle</p>
              <h5>${el._affectationData._EquipMarq} : ${el._affectationData._EquipMod} </h5>
              
                <p>Service appliqué</p>
            </div>
            <div class="col-md-4 p-4">
              <input type="text" name="desaff" readonly value="${el._affectation._id}" class="form-control" onclick="desAff(this.value)" btn-sm btn-block">cliquez pour désaffecter</input>
            </div>
          </div>
          </div>`;

              myEquip.appendChild(listItem);
              userSe(el._id);
            }
          })
      }
    })
    .catch(data => {
      console.log(data)
    })
}
function desAff(id){ // supprimer une affectation

  url = debut +":8090/affectations/" + id;
  fetch(url, {
    method: "DELETE",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
  .then(response => {
    if(response.ok){
      alert("Désaffecté.");
     // window.location.reload();
     getAffectations();
    }
  })
  .catch(data => {
    console.log(data.result)
  })
}

function createService(){ // creation de services
  // var nom = document.getElementById('nomserv').value;
  var type = document.getElementById('tservice').value;
  var ref = document.getElementById('rservice').value;
  var date = document.getElementById('dservice').value;
  //var equilie = document.getElementById('sequip').value;  
  var pserv = document.getElementById('pservice').value;  

  service = new Services(date,ref,type,pserv);
  url = debut + ":8090/services";
  fetch(url, {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    },
    body: JSON.stringify(service)
  })
    .then(response => {
      if (response.ok) {
        alert("Service : " + nom + " ajouté !");
        //getServices();
      //  window.location.reload();
      }
    })
    .catch(data => {
      console.log(data.result)
    }) 
}
function getServices(){ // récuperation des services 
  myEli = document.getElementById("tserv");
  url = debut + ":8090/services";

  fetch(url, {
    method: "GET",
    headers: {
      'Accept': 'application/json',
      'Content-type': 'application/json'
    }
  })
    .then(response => {
      if (response.ok) {
        myEli.innerHTML = '';
        response.json()
          .then((data) => {
            for (let el of data) {
              let listItem = document.createElement('tr');
              listItem.appendChild(
                document.createElement('td')
              ).textContent = `${el._Ref}`;

              listItem.appendChild(
                document.createElement('td')
              ).textContent = `${el._Type}`;

              // listItem.appendChild(
              //   document.createElement('td')
              // ).textContent = `${el._Prestataire}`;

              listItem.appendChild(
                document.createElement('td')
              ).textContent = `${el._DateExp}`;
         
              listItem.appendChild(
                document.createElement('td')
              ).innerHTML = '<button type="button" onclick="modServ(el._id);" class="btn btn-primary btn-sm btn-block">Modifier </button> <button type="button" onclik="delServ(`${el._id}`)" class="btn btn-danger btn-sm btn-block">Supprimer</button>';
              myEli.appendChild(listItem);
            }
          })
      }
    })
    .catch(data => {
      console.log(data)
    })
}