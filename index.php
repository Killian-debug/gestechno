<?php

session_start()

?>

<!DOCTYPE html>
<html lang="fr">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <meta name="Description" content="Enter your description here" />
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
    <link rel="stylesheet" href="assets/css/style.css">
    <title>GesTechno - Connexion</title>
</head>

<body>
    <div class="container-fluid">
        <div class="p-4 bg-primary text-white text-center">
            <h1>GesTechno - Maintenez votre parc au top</h1>
            <p>Une application web de gestion de parc informatique flexible et responsive.</p>
        </div>
        <form id="connexion" action="" method="POST">
            <h3>Connexion</h3>
            <div class="row">
                <div class="col-md-3">
                    <div class="form-group">
                        <label for="identifiant">Identifiant</label>
                        <input type="text" class="form-control" name="identifiant" id="identifiant" aria-describedby="helpId" placeholder="John" required>
                        <small id="helpId" class="form-text text-muted">Votre identifiant de connexion</small>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="form-group">
                        <label for="mdp">Mot de passe</label>
                        <input type="password" class="form-control" name="mdp" id="mdp" aria-describedby="helpId" placeholder="Doe" required>
                        <small id="helpId" class="form-text text-muted">Votre mot de passe</small>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col-md-3">
                    <input class="btn btn-danger btn-sm" type="reset" value="Annuler">
                </div>
                <div class="col-md-3">
                    <input class="btn btn-primary btn-sm" type="submit" value="Connexion">
                </div>
            </div>
        </form>
        <form id="inscription" action="" method="post">
            <h3>Inscription</h3>
            <div class="row">
                <div class="col-md-3">
                    <div class="form-group">
                        <label for="nom">Nom</label>
                        <input type="text" class="form-control" name="nom" id="nom" aria-describedby="helpId" placeholder="John" required>
                        <small id="helpId" class="form-text text-muted">Votre nom</small>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="form-group">
                        <label for="prenom">Prénom</label>
                        <input type="text" class="form-control" name="prenom" id="prenom" aria-describedby="helpId" placeholder="Doe" required>
                        <small id="helpId" class="form-text text-muted">Votre prénom</small>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col-md-3">
                    <div class="form-group">
                        <label for="login">Identifiant</label>
                        <input type="text" class="form-control" name="login" id="login" aria-describedby="helpId" placeholder="nomprenom12..." required>
                        <small id="helpId" class="form-text text-muted">Identifiant de connexion</small>
                    </div>
                </div>
                <div class="col-md-3">
                    <div class="form-group">
                        <label for="numero">Numero</label>
                        <input type="text" class="form-control" name="numero" id="numero" aria-describedby="helpId" placeholder="+22969696969" required>
                        <small id="helpId" class="form-text text-muted">Votre numero de téléphone</small>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col-md-6">
                    <div class="form-group">
                        <label for="pwd">Mot de passe</label>
                        <input type="password" class="form-control" name="pwd" id="pwd" aria-describedby="helpId" required>
                        <small id="helpId" class="form-text text-muted">Le mot de passe</small>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col-md-3">
                    <button class="btn btn-danger btn-sm" type="reset"> Annuler</button>
                </div>
                <div class="col-md-3">
                    <button class="btn btn-primary btn-sm" onclick="createTech()" type="button"> Enregistrer</button>
                </div>
            </div>
        </form>
        <span id="response"></span>

    </div>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
    <script src="./scripts/js/scripts.js"></script>

</body>

</html>