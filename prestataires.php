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
            <h1>GesTechno - liste de vos prestataires et leurs services</h1>
            <p>Info : Cliquez sur le nom pour afficher les services associés.</p>
        </div>

        <div class="row m-4">

            <div class="overflow-auto fournisseurs my-3">
                <table class="table fixed table-hover table-inverse table-responsive ">
                    <caption>Tableau qui liste les prestataires</caption>
                    <thead class="thead-inverse">
                        <tr style=" position: sticky; top: 0; z-index: 1;">
                            <th># </th>
                            <th>Nom </th>
                            <th>Contact </th>
                            <th>Date d'ajout</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody id="pres">
                        <tr>
                            <td scope="row">#2</td>
                            <td>test 2</td>
                            <td>test 2</td>
                            <td>test 2</td>
                            <td class="mod"> <button type="button" name="" id="" class="btn btn-primary btn-sm btn-block">Modifier </button> <button type="button" name="" id="" class="btn btn-danger btn-sm btn-block">Supprimer</button></td>

                        </tr>

                    </tbody>
                </table>
            </div>

            <div class="modal fade" id="modelId" tabindex="-1" role="dialog" aria-labelledby="modelTitleId" aria-hidden="true">
                <div class="modal-dialog" role="document">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h5 class="modal-title">Ajout d'un prestataire</h5>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <div class="modal-body">
                            <form action="" method="post" class="mt-4" id="addprestataire">
                                <div class="row">
                                    <div class="col-6">
                                        <label for="nompres">Nom du prestataire</label>
                                        <input list="nompress" class="form-control" placeholder="Kapersky..." required name="nompres" id="nompres">
                                        <datalist id="nompress">
                                            <option value="Kapersky">
                                            <option value="Microland">
                                            <option value="Mantenancier">
                                            <option value="Space tech">
                                        </datalist>
                                    </div>
                                    <div class="col-6">
                                        <label for="numpres">Numéro </label>
                                        <input type="tel" class="form-control" placeholder="22969696969" name="numpres" id="numpres" required>
                                    </div>
                                </div>

                        </div>
                        <div class="modal-footer">
                            <div class="row mt-3">
                                <div class="col-6">
                                    <button class="btn btn-danger btn-sm" data-dismiss="modal" type="reset"> Annuler</button>
                                </div>
                                <div class="col-6">
                                    <button class="btn btn-warning btn-sm" type="submit"> Enregistrer</button>
                                </div>
                            </div>
                            </form>
                            <!-- <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                            <button type="button" class="btn btn-primary">Save</button> -->
                        </div>
                    </div>
                </div>
            </div>



            <!-- Modal -->
            <div class="modal fade" id="modelId2" tabindex="-1" role="dialog" aria-labelledby="modelTitleId2" aria-hidden="true">
                <div class="modal-dialog" role="document">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h5 class="modal-title">Ajout d'un service</h5>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <div class="modal-body">
                            <form action="" method="post" class="mt-4">
                                <div class="row">
                                    <div class="col-12">
                                        <label for="tservice">Type de service</label>
                                        <select class="form-control form-control-sm" name="tservice" id="tservice" required>
                                            <option value="Licence Logiciel" selected>Licence Logiciel</option>
                                            <option value="Garantie">Garantie</option>
                                            <option value="Contrat de maintenance">Contrat de maintenance</option>
                                        </select>
                                    </div>
                                </div>
                                <div class="row ">
                                    <div class="col-6">
                                        <label for="rservice">Référence </label>
                                        <input type="text" class="form-control" placeholder="REF0***" id="rservice" name="rservice" required>
                                    </div>
                                    <div class="col-6">
                                        <label for="datexp">Date d'expiration</label>
                                        <input type="date" class="form-control" name="datexp" required id="dservice">
                                    </div>
                                </div>
                                <div class="row">
                                    <div class="col-6">
                                        <label for="sequip">Équipement lié</label>
                                        <select class="form-control" name="sequip" id="sequip"  >
                                        <!-- <datalist id="sequips" readonly> -->
                                            <option value="1">lapto hp</option>
                                            <option value="2">imprimante HD</option>
                                            <option value="3">Scanner</option>
                                            <option value="4">appareil photo</option>
                                        <!-- </datalist> -->
                                        </select>
                                    </div>
                                    <div class="col-6">
                                        <label for="pservice">Prestataire</label>
                                        <select class="form-control form-control-sm" name="pservice" id="pservice" required>
                                            <option value="1" selected>Microland</option>
                                            <option value="2">Kaparsky</option>
                                            <option value="3">Go maintenance</option>
                                        </select>
                                    </div>
                                </div>
                        </div>
                        <div class="modal-footer">
                            <div class="row mt-3">
                                <div class="col-6">
                                    <button class="btn btn-dark btn-sm" data-dismiss="modal" type="reset"> Annuler</button>
                                </div>
                                <div class="col-6">
                                    <button class="btn btn-warning btn-sm" type="button" onclick="createService();"> Enregistrer</button>
                                </div>
                            </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col-3">
                    <!-- Button trigger modal -->
                    <button type="button" class="btn btn-primary btn-sm mb-2" data-toggle="modal" data-target="#modelId">
                        Ajouter d'un prestataire
                    </button>
                </div>
                <div class="col-3">
                    <!-- Button trigger modal -->
                    <button type="button" class="btn btn-dark btn-sm" data-toggle="modal" data-target="#modelId2">
                        Ajout d'un service
                    </button>
                </div>
            </div>
            <div class="row">
            <div class="overflow-auto fournisseurs my-3">
                <table class="table fixed table-hover table-inverse table-responsive ">
                    <caption>Tableau qui liste les prestataires</caption>
                    <thead class="thead-inverse">
                        <tr style=" position: sticky; top: 0; z-index: 1;">
                            <th>Ref </th>
                            <th>Type </th>
                           
                            <th>Date de fin</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody id="tserv">
                        <tr>
                            <td scope="row">#2</td>
                            <td>test 2</td>
                         
                            <td>test 2</td>
                            <td class="mod"> <button type="button" name="" id="" class="btn btn-primary btn-sm btn-block">Modifier </button> <button type="button" name="" id="" class="btn btn-danger btn-sm btn-block">Supprimer</button></td>

                        </tr>

                    </tbody>
                </table>
            </div>
            </div>
        </div>
    </section>
    <?php
    include_once('./scripts/php/elements/footer.php')
    ?>
    <script>
        document.getElementById("addprestataire").onsubmit = function(e) {
            e.preventDefault();
            createPrestataire();
        }
        window.onload = function(e) {
            getPrestataires();
            equipSe("sequip");
            presSe("pservice");
        }
        $(document).ready(function(e){
            getServices();
        })
    </script>
</body>

</html>