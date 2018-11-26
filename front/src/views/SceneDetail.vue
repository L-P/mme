<template>
  <div>
    <div class="columns">

      <div class="column">
        <h2 class="title">{{this.scene.Name}} - {{this.scene.EntranceMessage}}</h2>
        <table class="table">
          <tbody>
            <tr v-for="v, k in scene" :key="k" v-if="k !== 'Rooms'">
              <td>{{k}}</td>
              <td>{{v | maybeHex}}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="column">
        <h2 class="title">Rooms</h2>
        <table class="table">
          <thead>
            <tr>
              <th>ID</th>
              <th>VROMStart</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="room in scene.Rooms" :key="room.ID">
              <td>{{room.ID | hex(2) }}</td>
              <td>{{room.VROMStart | hex(8) }}</td>
              <td><RouterLink
                class="button is-primary"
                :to="{name: 'RoomDetail', params: {start: room.VROMStart}}"
                >Details</RouterLink></td>
            </tr>
          </tbody>
        </table>
      </div>

    </div>
  </div>
</template>

<script src="./SceneDetail.js"></script>
