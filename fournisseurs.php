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
            <h1>GesTechno - liste de vos fournisseurs et leurs équipements</h1>
            <p>Info : Cliquez sur le nom pour afficher les équipements associés.</p>
        </div>
    
        <div class="row m-3">
            <div class="overflow-auto fournisseurs" >
            <table class="table fixed table-hover table-inverse table-responsive ">
                <caption>Tableau qui liste les fournisseurs</caption>
                <thead class="thead-inverse" >
                    <tr style=" position: sticky; top: 0; z-index: 1;" >
                        <!-- <th># </th> -->
                        <th>Nom </th>
                        <th>Contact </th>
                        <th>Adresse</th>
                        <!-- <th>Date d'ajout</th> -->
                        <th></th>
                    </tr>
                </thead>
                <tbody id="mylist" >
                    <tr>
                        <td scope="row">test 2</td>
                        <td>test 2</td>
                        <td>test 2</td>
                        <td>test 2</td>
                        <td class="mod"> <button type="button" name="" id="" class="btn btn-primary btn-sm btn-block">Modifier </button> <button type="button" name="" id="" class="btn btn-danger btn-sm btn-block">Supprimer</button></td>
                    </tr>
                    
                </tbody>
            </table>
            </div>
            
            <form action="" method="post" class="mt-3" id="addfournisseur" >
                <h3>Ajout d'un fournisseur</h3>
                <div class="row">
                    <div class="col-6">
                        <input type="text" id="nom" class="form-control" placeholder="John" name="email" required>
                    </div>
                    
                </div>
                <div class="row pt-2">
                    <div class="col-3">
                        <input type="tel" class="form-control" placeholder="229 69 69 69 69" name="numero" id="numero" required>
                    </div>
                    <div class="col-3">
                        <input type="text" class="form-control" placeholder="Pays, ville, quartier" name="adresse" id="adresse" required>
                    </div>
                </div>
                <div class="row mt-3">
                    <div class="col-2">
                        <button class="btn btn-danger btn-sm" type="reset"> Annuler</button>
                    </div>
                    <div class="col-2">
                        <button class="btn btn-primary btn-sm" type="submit"> Enregistrer</button>
                </div>
                </div>
            </form>
        </div>
    </section>
    <?php
    include_once('./scripts/php/elements/footer.php')
    ?>
    <script>
         function getFournisseurs(){ // fonction de récupération des fournisseur
    // et chargement dans le tableau
    myList = document.getElementById("mylist");

    url = debut +":8090/fournisseurs";
    fetch(url, {
      method : "GET",
      headers: {
        'Accept': 'application/json',
        'Content-type':'application/json'
      }
    })
    .then( response => {
      if(response.ok){
        myList.innerHTML= '';
        response.json()
        .then((data) => {
          for( let el of data ) {
            let listItem = document.createElement('tr');
            listItem.appendChild(
              document.createElement('td')
            ).textContent = `${el._Name}`;
            
            listItem.appendChild(
              document.createElement('td')
            ).textContent = `${el._Number}`;

            listItem.appendChild(
              document.createElement('td')
            ).textContent = `${el._Adresse}`;

            listItem.appendChild(
              document.createElement('td')
            ).innerHTML = '<button type="button" onclick="modFourni(el._id);" class="btn btn-primary btn-sm btn-block">Modifier </button> <button type="button" onclik="delFourni(el._id)" class="btn btn-danger btn-sm btn-block">Supprimer</button>';
            myList.appendChild(listItem);
          }
        
        //  myList.appendChild(aff)
       //  myList.innerHTML = JSON.stringify(data)
       })
       .catch(console.error);
      }
    })
    .catch(data => {
      console.log(data.result)
    })
  };
   // création d'un fournisseur lors du submit du formulaire
   document.getElementById("addfournisseur").onsubmit = function(e){
    e.preventDefault();
    nom = document.getElementById("nom").value;
    adresse = document.getElementById("adresse").value;
    number = document.getElementById("numero").value;

    fournisseur = new Fournisseurs(nom, adresse, number);

    fetch("http://localhost:8090/fournisseurs", {
      method : "POST",
      headers: {
        'Accept': 'application/json',
        'Content-type':'application/json'
      },
      body: JSON.stringify(fournisseur)
    })
    .then( response => {
      if(response.ok){
        alert("Fournisseur : " + nom  + " ajouté !");
        getFournisseur();
      }
    })
    .catch(data => {
      console.log(data.result)
    })
  }

        window.onload = function(e) {
            getFournisseurs();
        }
    </script>
</body>
</html>