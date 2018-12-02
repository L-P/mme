<template>
  <div class="container">

    <div class="columns">
      <div class="column">
        <h2 class="title">{{this.room.SceneName}} - room #{{this.room.ID}}</h2>
      </div>
      <div class="column">
        <RouterLink
          class="button is-primary"
          :to="{name: 'SceneDetail', params: {start: room.SceneVROMStart}}"
        >Scene</RouterLink>
      </div>
    </div>

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
