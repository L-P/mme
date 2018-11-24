<template>
  <div>

    <div
      class="navbar is-light"
      role="navigation"
      aria-label="main navigation"
    >

      <div class="navbar-brand">
        <div class="navbar-item">
            <RouterLink
              class="navbar-item"
              to="/"
            >
              <img
                alt="logo"
                src="assets/mme.svg"
              >
              Majora's Mask Explorer
            </RouterLink>
        </div>
      </div>

      <div class="navbar-start">
        <div class="navbar-menu is-active" id="navbar">
          <RouterLink class="navbar-item" to="/colormap">Color map</RouterLink>
        </div>
      </div>

    </div> <!-- navbar -->

    <RouterView />
  </div>
</template>

<style lang="scss">
@import '~bulma/css/bulma';
</style>
