<!DOCTYPE html>
<html lang="en">

<head>
    <?php
    include_once('./scripts/php/elements/header.php')
    ?>
    <title>GesTechno - Fournisseurs</title>
</head>

<body>
    <div class="main">
        <header>
            <?php
            include_once('./scripts/php/elements/menu.php')
            ?>
        </header>
    </div>
    <section class="pt-4">
        <div class="p-5 bg-primary text-white text-center">
            <h1>GesTechno - liste de vos utilisateurs et les informations les concernants</h1>
            <p>Info : Cliquez sur le nom pour afficher en dessous les informations.</p>
        </div>

        <div class="container mt-3">
            <h2>Liste de vos utilisateurs <button type="button" class="btn btn-primary btn-sm" data-toggle="modal" data-target="#modelId">
                    Ajouter <i class="fa fa-plus-circle" aria-hidden="true"></i>
                </button>
            </h2>
            <!-- <p><strong>Note:</strong> Ci-dessous la list.</p> -->
            <!-- Button trigger modal -->

            <div id="accordion">
                <div class="card">
                    <div class="card-header">
                        <a class="btn" data-toggle="collapse" href="#collapseOne">
                            Nom + Prénom #1
                        </a>
                    </div>
                    <div id="collapseOne" class="collapse show" data-parent="#accordion">
                        <div class="card-body">
                            <ul>
                                <li><strong>Numéro </strong> : +266 6699 69 </li>
                                <li><strong>Id </strong> : Parakou</li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="card">
                    <div class="card-header">
                        <a class="collapsed btn" data-toggle="collapse" href="#collapseThree">
                            Nom + Prénom #3
                        </a>
                    </div>
                    <div id="collapseThree" class="collapse" data-parent="#accordion">
                        <div class="card-body">
                            <ul>
                                <li><strong>Numéro </strong> : +266 6699 69 </li>
                                <li><strong>Id </strong> : Parakou</li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- Modal -->
        <div class="modal fade" id="modelId" tabindex="-1" role="dialog" aria-labelledby="modelTitleId" aria-hidden="true">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">Ajout d'un utilisateur</h5>
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                    </div>
                    <div class="modal-body">
                        <form action="" method="post" class="mt-4" id="addutilisateur">
                            <div class="row">
                                <div class="col-6">
                                    <label for="nom">Nom </label>
                                    <input type="text" class="form-control" placeholder="John" name="nom" id="nom" required>
                                </div>
                                <div class="col-6">
                                    <label for="prenom">Prénom </label>
                                    <input type="text" class="form-control" placeholder="Doe" name="prenom" id="prenom" required>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-12">
                                    <label for="numero">Numero</label>
                                    <input type="text" class="form-control" placeholder="+2299852700" name="numero" id="numero" required>
                                </div>

                            </div>

                            <!-- <div class="row mt-3">
                    <div class="col-2">
                        <button class="btn btn-danger btn-sm" type="reset"> Annuler</button>
                    </div>
                    <div class="col-2">
                        <button class="btn btn-primary btn-sm" type="submit"> Enregistrer</button>
                    </div>
                </div> -->

                      
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-dismiss="modal">Fermer</button>
                        <button type="submit" id="saveU" class="btn btn-primary">Enregistrer</button>
                    </div>
                    </form>
                </div>
            </div>
        </div>
    </section>

    <?php
    include_once('./scripts/php/elements/footer.php')
    ?>
    <script>
        document.getElementById("addutilisateur").onsubmit = function(e) {
            e.preventDefault();
            createUser();
        }

        window.onload = function() {
            getUsers();
        }
    </script>
</body>

</html>