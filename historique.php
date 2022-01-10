<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<?php 
include_once("./scripts/php/elements/header.php")
?>
<title>GesTechno - Historique</title>
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
            <h1>GesTechno - Historique de gestion du parc</h1>
            <p>Info : Utilisez les en-têtes pour naviguer dans l'historique.</p>
        </div>
  
<div class="container mt-3">


  <!-- Nav tabs -->
  <ul class="nav nav-tabs">
    <li class="nav-item">
      <a class="nav-link active" href="#home">Affections</a>
    </li>
    <li class="nav-item">
      <a class="nav-link" href="#menu1">Problèmes </a>
    </li>
    <li class="nav-item">
      <a class="nav-link" href="#menu2">Solutions</a>
    </li>
    <li class="nav-item">
      <a class="nav-link" href="#menu3">Mise hors parc</a>
    </li>
  </ul>

  <!-- Tab panes -->
  <div class="tab-content border mb-3">
    <div id="home" class="container tab-pane active"><br>
      <h3>Affections</h3>
      <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
    </div>
    <div id="menu1" class="container tab-pane fade"><br>
      <h3>Problèmes</h3>
      <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
    </div>
    <div id="menu2" class="container tab-pane fade"><br>
      <h3>Solutions</h3>
      <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam.</p>
    </div>
    <div id="menu3" class="container tab-pane fade"><br>
      <h3>Mise hors parc</h3>
      <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam.</p>
    </div>
  </div>
  <!-- <p class="act"><b>Active Tab</b>: <span></span></p>
  <p class="prev"><b>Pre    vious Tab</b>: <span></span></p> -->
</div>
  <?php
    include_once('./scripts/php/elements/footer.php')
    ?>
<script>
// $(document).ready(function(){
//   $(".nav-tabs a").click(function(){
//     $(this).tab('show');
//   });
//   $('.nav-tabs a').on('shown.bs.tab', function(event){
//     var x = $(event.target).text();         // active tab
//     var y = $(event.relatedTarget).text();  // previous tab
//     $(".act span").text(x);
//     $(".prev span").text(y);
//   });
});
</script>  
    </section>
  
</body>
</html>