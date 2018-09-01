<template>
  <div id="app">
    <!-- The Top Navbar -> TEMP Should be Fixed -->
    <nav class="navbar is-info" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="#">
          <img src="https://bulma.io/images/bulma-logo.png" alt="Bulma: a modern CSS framework based on Flexbox" width="112" height="28">
        </a>

        <!--
        <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>-->
      </div>
      <!--
      <a class="navbar-item has-text-white">
        Home
      </a>

      <a class="navbar-item has-text-white">
        Testing2
      </a>-->
    </nav>

    <!-- Left Side Menu -->
    <div class="columns is-fullheight">
      <div class="column is-2 is-sidebar-menu is-hidden-mobile">
        <aside class="menu">
      <p class="menu-label">
        General
      </p>
      <ul class="menu-list">
        <li><a>Dashboard</a></li>
      </ul>
      <p class="menu-label">
        Content
      </p>
      <ul class="menu-list">
        <li><a>Policies</a></li>
        <li><a>Profiles</a></li>
        <li><a>Applications</a></li>
        <li><a>Updates</a></li>
      </ul>
      <p class="menu-label">
        Groups
      </p>
      <ul class="menu-list">
        <li><a>Groups</a></li>
        <li><a>Scooped Groups</a></li>
        <li><a>Prestage Enrollment</a></li>
      <p class="menu-label">
        Settings
      </p>
      <ul class="menu-list">
        <li><a>Mattrax Settings</a></li>
      </ul>
    </ul>
    </aside>
      </div>
      <div class="column is-main-content"> <!-- container column is-8 body-content -->






        <!--<div class="container column is-8 body-content">-->
          <!-- Breadcrumbs -->
          <nav class="breadcrumb" aria-label="breadcrumbs">
            <ul>
              <li><a href="#">Devices</a></li>
              <li class="is-active"><a href="#" aria-current="page">Oscar's iPhone</a></li>
            </ul>
          </nav>

          <!-- Page Tabs -->
          <div class="tabs">
            <ul>
              <li class="is-active"><a>Device Overview</a></li>
              <li><a>Policies</a></li>

              etc

              <li><a>Settings</a></li>
            </ul>
          </div>

          <!-- Page Content (TEMP This Will Be On React Router For SubViews) -->
          <div class="columns">
            <div class="column is-3 has-text-centered">
              <img src="iPhone10.png"></img>
              <p>iPhone 10, 64GB</p>
            </div>
            <div class="column is-9">



              <!-- The Device's Information -->
              <!--<p class="para-title">The Devices Details:</p>-->
              <table class="table is-responsive is-fullwidth">
                <h5 class="title is-5">Device Details</h5>
                <tr>
                  <td>Model:</td>
                  <td>iPhone 10, 64GB</td>
                </tr>
                <tr>
                  <td>Operating System:</td>
                  <td>iOS 11.3.4 (15A5278f)</td>
                </tr>
                <tr>
                  <td>Last Checkin:</td>
                  <td>Feb 26, 2018 4:23AM</td>
                </tr>
                <tr>
                  <td>Device ID:</td>
                  <td>040e19a7-0344-41b5-9da2-20f04d162fe9</td>
                </tr>
                <tr>
                  <td>Supervised:</td>
                  <td>Yes</td>
                </tr>
              </table>

              <h5 class="title is-5">Device Actions</h5>
              <a class="button">
                <span class="icon is-small">
                  <i class="fas fa-sync"></i>
                </span>
                <span>Force Checkin</span>
              </a>

              <a class="button">
                <span class="icon is-small">
                  <i class="fas fa-lock"></i>
                </span>
                <span>Lock</span>
              </a>

              <a class="button">
                <span class="icon is-small">
                  <i class="fas fa-retweet"></i>
                </span>
                <span>Restart</span>
              </a>

              <a class="button is-danger">
                <span class="icon is-small">
                  <i class="fas fa-power-off"></i>
                </span>
                <span>Power Off</span>
              </a>


            </div>
          </div>
        <!--</div>-->










      </div>
    </div>





    <!--<div id="nav">
      <router-link to="/">Home</router-link> |
    </div>
    <router-view/>-->




  </div>
</template>

<style lang="scss">
$breadcrumb-item-padding-vertical: 0.5em;

.breadcrumb:not(:last-child) {
  margin-bottom: 0 !important;
}

@import "bulma"; //TODO: Extend The Bulma Theme
</style>

<style lang="scss">
  @charset "utf-8";

  // Customization

  // You can easily customize Bulma with your own variables.
  // Just uncomment the following block to see the result.

  /*
  // 1. Import the initial variables
  @import "../node_modules/bulma/sass/utilities/initial-variables";

  // 2. Set your own initial variables
  // Update the blue shade, used for links
  $blue: #06bcef;
  // Add pink and its invert
  $pink: #ff8080;
  $pink-invert: #fff;
  // Update the sans-serif font family
  $family-sans-serif: "Helvetica", "Arial", sans-serif;

  // 3. Set the derived variables
  // Use the new pink as the primary color
  $primary: $pink;
  $primary-invert: $pink-invert;

  // 4. Import the rest of Bulma
  */

  //TODO: Shows Good Use of Varibles For Themeing: https://codepen.io/wikiki/pen/OjdNKM


  //@import "../node_modules/bulma/bulma";

  $grey-dark:    hsl(0, 0%, 29%) !default; //TODO: MOve To To If Code Is Kept
  $navbar-height: 52px;


  ///////
  // The Login Page
  ///////

  /*#login { //TODO: The Login Page
  .login-form {
    width: 400px;
    margin: 0 auto; //Center It
  }
}*/


  .body-content {
    margin-top: 10px;
  }

  .info-bar {
    padding-bottom: 20px;
  }

  .para-title {
    padding-bottom: 10px;
  }

  .sidebar-menu {
    //margin-top: 52px;
    position: absolute;
    //right: 0px;
    left: 0px;
    top: 52px;
    padding: 10px 0 0 10px;
  }

  /* Custom Testing View */
  .is-sidebar-menu {
    padding: 2.5rem;
    background: $grey-dark;

    li > a {
      color: white;
    }
  }

  .columns {
    &.is-fullheight {
      min-height: calc(100vh - ( #{$navbar-height} - .75rem ) );
      max-height: calc(100vh - ( #{$navbar-height} - .75rem ) );
      height: calc(100vh - ( #{$navbar-height} - .75rem ) );
      display: flex;
      flex-direction: row;
    justify-content: stretch;
      .column {
        overflow-y: auto
      }
    }
  }

  .breadcrumb {
    margin-bottom: 0px;
  }

  .policy-type-list {
    tr > td {
      padding: 10px 10px 20px 20px;
    }

    .subtitle {
      display: inline-block;
    }
  }

</style>
