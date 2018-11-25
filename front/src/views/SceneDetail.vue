<template>
  <div>
    <table class="table">
      <tbody>
        <tr v-for="v, k in this.scene" :key="k">
          <td>{{k}}</td>
          <td>{{v | maybeHex}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script src="./SceneDetail.js"></script>
