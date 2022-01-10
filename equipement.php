<!DOCTYPE html>
<html lang="fr">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
  <link rel="stylesheet" href="./bootstrap/css/bootstrap.css">
  <link rel="stylesheet" href="./assets/css/sidebars.css">
  <link rel="stylesheet" href="./assets/css/style.css">
  <title>GesTechno</title>
</head>



<body>
  <div class="main">
    <header>
      <nav class="navbar navbar-expand-sm navbar-dark bg-dark fixed-top ">
        <div class="container-fluid">
          <a class="navbar-brand" href="#">GesTechno</a>
          <button class="navbar-toggler d-lg-none" type="button" data-toggle="collapse" data-target="#collapsibleNavId" aria-controls="collapsibleNavId" aria-expanded="false" aria-label="Toggle navigation"></button>
          <div class="collapse navbar-collapse" id="collapsibleNavId">
            <ul class="navbar-nav me-auto">
              <li class="nav-item active">
                <a class="nav-link" href="equipement.php">Accueil <span class="sr-only">(actuelle)</span></a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="fournisseurs.php">Fournisseurs</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="prestataires.php">Prestataire</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="historique.php">historique</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="utilisateurs.php">Utilisateurs</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" id="deco" href="#">Déconnexion</a>
              </li>
              <!-- <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="#" id="dropdownId" data-toggle="dropdown" aria-haspopup="true"
                  aria-expanded="false">Actions</a>
                <div class="dropdown-menu" aria-labelledby="dropdownId">
                  <a class="dropdown-item" href="#">Déconnexion</a>
                  <a class="dropdown-item" href="#">Action 2</a>
                </div>
              </li> -->
            </ul>
            <form class="d-flex">
              <input class="form-control me-2" type="text" placeholder="Chercher">
              <button class="btn btn-primary" type="button">chercher</button>
            </form>
          </div>
        </div>
      </nav>
    </header>
  </div>
  <section class="pt-4">


    <div class="p-5 bg-primary text-white text-center">
      <h1>GesTechno - Maintenez votre parc au top</h1>
      <p>Une application web de gestion de parc informatique flexible et responsive.</p>
    </div>


    <div class="container mt-5">
      <div class="row">
        <div class="col-sm-3">
          <div class="d-flex flex-column flex-shrink-0 p-3 text-white bg-dark" style="width: 280px;">
            <!-- <a href="/" class="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none">
              <svg class="bi me-2" width="40" height="32">
                <use xlink:href="#bootstrap"></use>
              </svg>
              <span class="fs-4">Filtrer : </span>
            </a> -->
            <hr>


            <!-- Modal -->
            <div class="modal fade" id="affectmodal" tabindex="-1" role="dialog" aria-labelledby="modelTitleId" aria-hidden="true">
              <div class="modal-dialog" role="document">
                <div class="modal-content">
                  <div class="modal-header">
                    <h5 class="modal-title">Modal title</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                      <span aria-hidden="true">&times;</span>
                    </button>
                  </div>
                  <div class="modal-body">
                    <form action="" method="post" id="afform">
                      <label for="useraff" class="text-dark">Utilisateurs</label>
                      <select class="form-control" name="useraf" id="useraff">
                        <option></option>
                      </select>
                      <label for="equipaff" class="text-dark">Equipements</label>
                      <select class="form-control" name="equipaf" id="equipaff">
                        <option></option>
                      </select>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Annuler</button>
                    <button type="submit" class="btn btn-primary">Enregistrer</button>
                  </div>
                  </form>
                </div>
              </div>
            </div>
            <ul class="nav nav-pills flex-column mb-auto">

              <li class="nav-item">
                <!-- Button trigger affectation modal -->
                <a href="#" class="nav-link active" aria-current="page" data-toggle="modal" data-target="#affectmodal" onclick="userSe('useraff');equipSe('equipaff');">
                  <svg class="bi me-2" width="16" height="16">
                    <use xlink:href="#home"></use>
                  </svg>
                  Affecter
                </a>
              </li>

              <li>
                <div class="dropdown">
                  <nav class="nav-item dropdown" aria-label="affection ">
                    <a class="nav-link dropdown-toggle" href="#" id="dropdownId" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <strong>Affectation</strong>
                    </a>
                    <div class="dropdown-menu" aria-labelledby="dropdownId">
                      <a class="dropdown-item" onclick="getAffectations()" href="#">Affectés</a>
                      <a class="dropdown-item" href="#">Non affectés</a>
                    </div>
                  </nav>
                </div>
              </li>


              <li>
                <div class="dropdown">
                  <nav class="nav-item dropdown" aria-label="Etat utilisation">
                    <a class="nav-link dropdown-toggle" href="#" id="dropdownId" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <strong>Etat</strong>
                    </a>
                    <div class="dropdown-menu" aria-labelledby="dropdownId">
                      <a class="dropdown-item" href="#">En panne</a>
                      <a class="dropdown-item" href="#">En cours</a>
                      <a class="dropdown-item" href="#">Hors parc</a>
                    </div>
                  </nav>
                </div>
              </li>
              <hr>

              <li>
                <form action="" method="post">
                  <div class="form-group">
                    <label for="utilisateur">Par utilisateurs : </label>
                    <select class="form-control" name="utilisateur" id="utilisateur">
                      <option selected>Tous...</option>
                      <option value="">Utilisateur 1</option>
                      <option value="">Utilisateur 2</option>
                    </select>
                  </div>
                </form>
              </li>
              <li>
                <form action="" method="post">
                  <div class="form-group">
                    <label for="ffourni">Par Fournisseurs : </label>
                    <select class="form-control" name="Fournisseur" id="ffourni">
                      <option selected>Tous...</option>
                    </select>
                  </div>
                </form>
              </li>
            </ul>


          </div>
        </div>

        <div class="col-sm-9">
          <div class="row mb-3">
            <div class="col-sm-9 ">
              <!-- Button trigger modal -->
              <button type="button" class="btn btn-primary btn-sm" data-toggle="modal" data-target="#modelId">
                Ajouter un équipement
              </button>
            </div>
          </div>

          <div id="equip">
            <div class="card">
              <div class="row">
                <div class="col-md-3" style="align-items: center;display: inline-grid;">

                  <img height="100vh" src="./assets/media/img1.jpg" alt="Card image cap">
                  <div class="card-body">
                    <h5 class="card-title">Utilisateur X</h5>
                  </div>
                </div>
                <div class="col-md-5">
                  <h2>Nom equipement</h2>
                  <h5>Marque et modèle</h5>
                  <p>Service appliqué</p>
                  <p>Description : Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do
                    eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
                    exercitation ullamco.</p>
                </div>
                <div class="col-md-4 p-4">
                  <form action="" method="post">
                    <div class="form-group">
                      <label for="utilisateur">Affecter à : </label>
                      <select class="form-control" name="utilisateur" id="utilisateur">
                        <option selected>Sélectionner...</option>
                        <option value="">Utilisateur 1</option>
                        <option value="">Utilisateur 2</option>
                      </select>
                    </div>
                    <div class="form-group pt-2">
                      <input class="form-control" type="submit" value="Valider">
                    </div>
                  </form>
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
            </div>
            <div class="card">
              <div class="row">
                <div class="col-md-3" style="align-items: center;display: inline-grid;">

                  <img height="100vh" src="./assets/media/img1.jpg" alt="Card image cap">
                  <div class="card-body">
                    <h5 class="card-title">Utilisateur X</h5>
                  </div>
                </div>
                <div class="col-md-5">
                  <h2>Nom equipement</h2>
                  <h5>Marque et modèle</h5>
                  <p>Service appliqué</p>
                  <p>Description : Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do
                    eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
                    exercitation ullamco.</p>
                </div>
                <div class="col-md-4 p-4">
                  <form action="" method="post">
                    <div class="form-group">
                      <label for="utilisateur">Affecter à : </label>
                      <select class="form-control" name="utilisateur" id="utilisateur">
                        <option selected>Sélectionner...</option>
                        <option value="">Utilisateur 1</option>
                        <option value="">Utilisateur 2</option>
                      </select>
                    </div>
                    <div class="form-group pt-2">
                      <input class="form-control" type="submit" value="Valider">
                    </div>
                  </form>
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
            </div>
          </div>
          <!-- Modal -->
          <div class="modal fade" id="modelId" tabindex="-1" role="dialog" aria-labelledby="modelTitleId" aria-hidden="true">
            <div class="modal-dialog" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">Ajout d'un équipement</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <form action="" method="post" class="mt-4" id="cequip">
                    <div class="row">
                      <div class="col-6">
                        <label for="mequip">Marque </label>
                        <input type="text" class="form-control" placeholder="HP" name="mequip" id="mequip" required>
                      </div>
                      <div class="col-6">
                        <label for="moequip">Modèle</label>
                        <input type="text" class="form-control" placeholder="Laser 9700" name="moequip" id="moequip" required>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-6">
                        <label for="pequip">Prix d'achat </label>
                        <input type="number" class="form-control" placeholder="15000" name="pequip" id="pequip" aria-describedby="helpId" required min="1">
                        <small id="helpId" class="form-text text-muted">Prix en F CFA</small>
                      </div>
                      <div class="col-6">
                        <label for="etequip">Etat achat</label>
                        <select class="form-control form-control-sm" name="etequip" id="etequip" required>
                          <option value="neuf" selected>Neuf</option>
                          <option value="occasion">Occasion</option>
                        </select>
                      </div>
                    </div>

                    <div class="row">
                      <div class="col-12">
                        <label for="caract">Caractéristiques Technique</label>
                        <textarea name="caract" class="form-control" id="caract" cols="35" rows="2"></textarea>
                      </div>
                    </div>
                    <div class="row ">
                      <div class="col-6">
                        <label for="efourni">Founisseur</label>
                        <select class="form-control form-control-sm" name="efourni" id="efourni" required>
                          <option value="0" selected> Selectionner...</option>
                        </select>
                      </div>
                      <div class="col-6">
                        <label for="dequip">Date d'achat</label>
                        <input type="date" class="form-control" name="dequip" id="dequip" required>
                      </div>
                    </div>
                    <div class="row">
                      <!-- Bouton d'ajout d'un nouveau fournisseur  en même temps que l'équipement -->
                      <div class="col-12">
                        <label for="newf">Nouveau fournisseur ?</label>
                        <input type="checkbox" data-toggle="collapse" class="btn btn-primary my-2" data-target="#demo" name="newf" id="newf">

                        <div id="demo" class="collapse">
                          <div class="row">
                            <div class="col-12">
                              <input type="text" class="form-control" placeholder="John" name="nomf" id="nomf">
                            </div>
                          </div>
                          <div class="row pt-2">
                            <div class="col-6">
                              <input type="tel" class="form-control" placeholder="229 69 69 69 69" name="numero" id="numero">
                            </div>
                            <div class="col-6">
                              <input type="text" class="form-control" placeholder="Pays, ville, quartier" name="adresse" id="adresse">
                            </div>
                          </div>

                        </div>
                      </div>
                    </div>

                </div>
                <div class="modal-footer">
                  <div class="row mt-3">
                    <div class="col-6">
                      <button class="btn btn-danger btn-sm" type="reset" data-dismiss="modal"> Annuler</button>
                    </div>
                    <div class="col-6">
                      <button class="btn btn-primary btn-sm" type="submit"> Enregistrer</button>
                    </div>
                  </div>

                  </form>
                  <!-- <button type="button" class="btn btn-secondary" data-dismiss="modal">Fermer</button>
                    <button type="button" class="btn btn-primary">Save</button> -->
                </div>
              </div>
            </div>
          </div>


        </div>
      </div>
    </div>
  </section>
  </div>

  <script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>
  <script src="./scripts/js/scripts.js"></script>
  <script>
    function getFourniSelect() { // fonction de récupération des fournisseurs
      // et charge le select des fournisseurs
      select = document.getElementById("efourni");
      selectf = document.getElementById("ffourni");

      url = debut + ":8090/fournisseurs"
      fetch(url, {
          method: "GET",
          headers: {
            'Accept': 'application/json',
            'Content-type': 'application/json'
          }
        })
        .then(response => {
          if (response.ok) {
            select.innerHTML = '';
            response.json()
              .then(data => {
                for (let el of data) {
                  //  let listItem = document.createElement('option');
                  opt = document.createElement('option')
                  opt.value = `${el._id}`;
                  opt.textContent = `${el._Name}`;
                  select.appendChild(opt);
                }
                for (let le of data) {
                  //  let listItem = document.createElement('option');
                  opt = document.createElement('option')
                  opt.value = `${le._id}`;
                  opt.textContent = `${le._Name}`;
                  selectf.appendChild(opt);
                }
              })
          }
        })
        .catch(error => console.log(error))
    }

    document.getElementById("cequip").onsubmit = function(e) {
      e.preventDefault();
      createEquip();
    }
    document.getElementById("afform").onsubmit = function(e) {
      e.preventDefault();
      equipid = document.getElementById("equipaff").value;
      userid = document.getElementById("useraff").value;
      affecter(equipid, userid);
    }
    window.onload = function() {
      getFourniSelect();
      getEquips();
    }
  </script>
</body>

</html>