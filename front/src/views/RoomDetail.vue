<template>
  <div class="container">
    <table class="table">
      <tbody>
        <tr v-for="v, k in room" :key="k">
          <td>{{k}}</td>
          <td>{{v | maybeHex}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script src="./RoomDetail.js"></script>
