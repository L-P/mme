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
        <a
          role="button"
          aria-label="menu"
          aria-expanded="false"
          data-target="navbar"
          :class="{ 'navbar-burger': true, burger: true, 'is-active': burgerOpen }"
          @click="burgerOpen = !burgerOpen"
        >
          <span aria-hidden="true" />
          <span aria-hidden="true" />
          <span aria-hidden="true" />
        </a>
      </div>

      <div class="navbar-start">
        <div
          id="navbar"
          :class="{ 'navbar-menu': true, 'is-active': burgerOpen }"
        >
          <RouterLink
            class="navbar-item"
            to="/colormap"
          >
            Color map
          </RouterLink>
          <RouterLink
            class="navbar-item"
            to="/files"
          >
            Files
          </RouterLink>
          <RouterLink
            class="navbar-item"
            to="/scenes"
          >
            Scenes
          </RouterLink>
          <RouterLink
            class="navbar-item"
            to="/messages"
          >
            Messages
          </RouterLink>
        </div>
      </div>
    </div> <!-- navbar -->

    <RouterView />
  </div>
</template>

<script>
export default {
  data() {
    return {
      burgerOpen: false,
    };
  },
};
</script>

<style lang="scss">
@import '~bulma/css/bulma';
</style>
