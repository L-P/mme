<template>
  <div class="container">

    <div class="columns">
      <div class="column">
        <h2 class="title">{{this.room.SceneName}} - room #{{this.room.ID}}</h2>
      </div>
      <div class="column is-one-fifth">
        <!-- have to v-if here to avoid a warning because SceneVROMStart
             does not exist on first render -->
        <RouterLink
          v-if="room.SceneVROMStart"
          class="button is-primary"
          :to="{name: 'SceneDetail', params: {start: room.SceneVROMStart}}"
        >Scene</RouterLink>
      </div>
    </div>

    <b-tabs>

      <b-tab-item label="Configuration">
        <table class="table">
          <tbody>
            <tr v-for="v, k in room" :key="k">
              <td>{{k}}</td>
              <td>{{v | maybeHex}}</td>
            </tr>
          </tbody>
        </table>
      </b-tab-item>

      <b-tab-item label="Actors">
        <ActorList :actors=this.room.ActorList></ActorList>
      </b-tab-item>

      <b-tab-item label="Objects">
      </b-tab-item>

    </b-tabs>
  </div>
</template>

<script src="./RoomDetail.js"></script>
