<template>
  <div class="container">
    <div class="columns">

      <div class="column is-one-quarter">
        <img src="assets/mme.svg">
      </div>

      <div class="column">
        <table class="table">
          <tbody>
            <tr v-for="v, k in this.rom" :key="k">
              <td>{{k}}</td>
              <td>{{v}}</td>
            </tr>
          </tbody>
        </table>
      </div>

    </div>
  </div>
</template>

<style scoped>
table {
  width: 100%;
}
</style>

<script src="./Home.js"></script>
